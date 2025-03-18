package server

import (
	chat_context "ai-chat-service/chat-server/chat-context"
	"ai-chat-service/pkg/config"
	"ai-chat-service/pkg/log"
	"ai-chat-service/pkg/zerror"
	"ai-chat-service/proto"
	"ai-chat-service/services"
	keywords_filter "ai-chat-service/services/keywords-filter"
	keywords_proto "ai-chat-service/services/keywords-filter/proto"
	"ai-chat-service/services/tokenizer"
	"context"
	"github.com/google/uuid"
	"github.com/sashabaranov/go-openai"
	"time"
)

const ChatPrimedTokens = 2

type openaiConf struct {
	ApiKey            string
	BaseUrl           string
	Model             string
	MaxTokens         int
	Temperature       float32
	TopP              float32
	PresencePenalty   float32
	FrequencyPenalty  float32
	BotDesc           string
	ContextTTL        int
	ContextLen        int
	MinResponseTokens int
}
type app struct {
	openaiConf *openaiConf
	log        log.ILogger
	// TODO 内容上下文对象
	contextCache chat_context.ContextCache
}

func (s *chatService) newApp(in *proto.ChatCompletionRequest, contextCache chat_context.ContextCache) *app {
	conf := &openaiConf{
		ApiKey:            s.config.Chat.ApiKey,
		BaseUrl:           s.config.Chat.BaseUrl,
		Model:             s.config.Chat.Model,
		MaxTokens:         s.config.Chat.MaxTokens,
		Temperature:       s.config.Chat.Temperature,
		TopP:              s.config.Chat.TopP,
		PresencePenalty:   s.config.Chat.PresencePenalty,
		FrequencyPenalty:  s.config.Chat.FrequencyPenalty,
		BotDesc:           s.config.Chat.BotDesc,
		ContextTTL:        s.config.Chat.ContextTTL,
		ContextLen:        s.config.Chat.ContextLen,
		MinResponseTokens: s.config.Chat.MinResponseTokens,
	}
	if in.ChatParam != nil {
		if in.ChatParam.Model != "" {
			conf.Model = in.ChatParam.Model
		}
		if in.ChatParam.TopP != 0 {
			conf.TopP = in.ChatParam.TopP
		}
		if in.ChatParam.FrequencyPenalty != 0 {
			conf.FrequencyPenalty = in.ChatParam.FrequencyPenalty
		}
		if in.ChatParam.PresencePenalty != 0 {
			conf.PresencePenalty = in.ChatParam.PresencePenalty
		}
		if in.ChatParam.Temperature != 0 {
			conf.Temperature = in.ChatParam.Temperature
		}
		if in.ChatParam.BotDesc != "" {
			conf.BotDesc = in.ChatParam.BotDesc
		}
		if in.ChatParam.MaxTokens != 0 {
			conf.MaxTokens = int(in.ChatParam.MaxTokens)
		}
		if in.ChatParam.ContextTTL != 0 {
			conf.ContextTTL = int(in.ChatParam.ContextTTL)
		}
		if in.ChatParam.ContextLen != 0 {
			conf.ContextLen = int(in.ChatParam.ContextLen)
		}
		if in.ChatParam.MinResponseTokens != 0 {
			conf.MinResponseTokens = int(in.ChatParam.MinResponseTokens)
		}
	}
	return &app{
		openaiConf:   conf,
		log:          s.log,
		contextCache: contextCache,
	}
}
func (a *app) getOpenaiClient() *openai.Client {
	accessToken := a.openaiConf.ApiKey
	config := openai.DefaultConfig(accessToken)
	config.BaseURL = a.openaiConf.BaseUrl
	client := openai.NewClientWithConfig(config)
	return client
}
func (a *app) buildChatCompletionRequest(in *proto.ChatCompletionRequest, stream bool) (req openai.ChatCompletionRequest, tokens, currTokens int, currMessage openai.ChatCompletionMessage, err error) {
	//当前消息
	currMessage = openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: in.Message,
	}
	req = openai.ChatCompletionRequest{
		Model: a.openaiConf.Model,
		Messages: []openai.ChatCompletionMessage{
			currMessage,
		},
		MaxTokens:        a.openaiConf.MinResponseTokens,
		Temperature:      a.openaiConf.Temperature,
		TopP:             a.openaiConf.TopP,
		PresencePenalty:  a.openaiConf.PresencePenalty,
		FrequencyPenalty: a.openaiConf.FrequencyPenalty,
		Stream:           stream,
	}
	contextList := make([]*chat_context.ChatMessage, 0)
	if in.EnableContext {
		//从缓存中获取上下文信息
		contextList = a.getContext(in.Pid)
	}
	//重构req.Messages
	tokens, currTokens, req.Messages, err = a.rebuildMessages(contextList, currMessage)
	if err != nil {
		a.log.Error(err)
		return
	}
	req.MaxTokens = a.openaiConf.MaxTokens - tokens
	return
}
func (a *app) rebuildMessages(contextList []*chat_context.ChatMessage, currMessage openai.ChatCompletionMessage) (tokens, currTokens int, messages []openai.ChatCompletionMessage, err error) {
	var sysMessage openai.ChatCompletionMessage
	botTokens := 0
	if a.openaiConf.BotDesc != "" {
		sysMessage = openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleSystem,
			Content: a.openaiConf.BotDesc,
		}
		botTokens, err = tokenizer.GetTokens(&sysMessage, a.openaiConf.Model)
		if err != nil {
			a.log.Error(err)
			return
		}
	}
	messages = []openai.ChatCompletionMessage{currMessage}
	currTokens, err = tokenizer.GetTokens(&currMessage, a.openaiConf.Model)
	if err != nil {
		a.log.Error(err)
		return
	}
	if currTokens > a.openaiConf.MaxTokens-a.openaiConf.MinResponseTokens-botTokens-ChatPrimedTokens {
		err = zerror.NewByMsg("请求消息超限")
		a.log.Error(err)
		return
	}
	tokens = currTokens + botTokens + ChatPrimedTokens
	if contextList != nil {
		for _, item := range contextList {
			if tokens+item.Tokens+ChatPrimedTokens > a.openaiConf.MaxTokens-a.openaiConf.MinResponseTokens {
				break
			}
			messages = append(messages, item.Message)
			tokens += item.Tokens + ChatPrimedTokens
		}
	}
	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}
	if botTokens > 0 {
		messages = append([]openai.ChatCompletionMessage{sysMessage}, messages...)
	}
	return
}
func (a *app) buildChatCompletionResponse(msg string) *proto.ChatCompletionResponse {
	res := &proto.ChatCompletionResponse{
		Id:      uuid.New().String(),
		Object:  "chat.completion",
		Created: time.Now().Unix(),
		Model:   a.openaiConf.Model,
		Choices: []*proto.ChatCompletionChoice{
			{
				Message: &proto.ChatCompletionMessage{
					Role:    openai.ChatMessageRoleAssistant,
					Content: msg,
				},
				FinishReason: "stop",
			},
		},
		Usage: &proto.Usage{
			PromptTokens:     0,
			CompletionTokens: 0,
			TotalTokens:      0,
		},
	}
	return res
}

func (a *app) buildChatCompletionStreamResponse(id, delta, finishReason string) *proto.ChatCompletionStreamResponse {
	res := &proto.ChatCompletionStreamResponse{
		Id:      id,
		Object:  "chat.completion.chunk",
		Created: time.Now().Unix(),
		Model:   a.openaiConf.Model,
		Choices: []*proto.ChatCompletionStreamChoice{
			{
				Index: 0,
				Delta: &proto.ChatCompletionStreamChoiceDelta{
					Content: delta,
					Role:    openai.ChatMessageRoleAssistant,
				},
				FinishReason: finishReason,
			},
		},
	}
	return res
}

func (a *app) buildChatCompletionStreamResponseList(id, msg string) []*proto.ChatCompletionStreamResponse {
	list := make([]*proto.ChatCompletionStreamResponse, 0)
	for _, delta := range msg {
		list = append(list, a.buildChatCompletionStreamResponse(id, string(delta), ""))
	}
	return list
}

func (a *app) getContext(id string) []*chat_context.ChatMessage {
	maxLen := a.openaiConf.ContextLen
	list := make([]*chat_context.ChatMessage, 0, maxLen)
	key := id
	for i := 0; i < maxLen; i++ {
		value, err := a.contextCache.Get(key)
		if err != nil {
			a.log.Error(err)
			return nil
		}
		if value == nil {
			break
		}
		list = append(list, value)
		key = value.PID
	}
	return list
}
func (a *app) saveContext(value *chat_context.ChatMessage) error {
	err := a.contextCache.Set(value.ID, value, a.openaiConf.ContextTTL)
	if err != nil {
		a.log.Error(err)
		return err
	}
	return nil
}
func (a *app) keywords(in *proto.ChatCompletionRequest) []string {
	pool := keywords_filter.GetKeywordsClientPool()
	conn := pool.Get()
	defer pool.Put(conn)
	accessToken := config.GetConfig().DependOn.Keywords.AccessToken
	client := keywords_proto.NewFilterClient(conn)
	ctx := services.AppendBearerTokenToContext(context.Background(), accessToken)
	req := &keywords_proto.FilterReq{
		Text: in.Message,
	}
	res, err := client.FindAll(ctx, req)
	if err != nil {
		a.log.Error(err)
		return []string{}
	}
	return res.Keywords

}
func (a *app) sensitive(in *proto.ChatCompletionRequest) (ok bool, msg string, err error) {
	pool := keywords_filter.GetSensitiveClientPool()
	conn := pool.Get()
	defer pool.Put(conn)
	accessToken := config.GetConfig().DependOn.Sensitive.AccessToken
	client := keywords_proto.NewFilterClient(conn)
	ctx := services.AppendBearerTokenToContext(context.Background(), accessToken)
	req := &keywords_proto.FilterReq{
		Text: in.Message,
	}
	res, err := client.Validate(ctx, req)
	if err != nil {
		a.log.Error(err)
		return false, "", err
	}
	ok = res.Ok
	if !ok {
		msg = "触发到了知识盲区，请换个问题再问"
	}
	return
}
