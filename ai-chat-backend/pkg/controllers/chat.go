package controllers

import (
	"ai-chat-backend/pkg/config"
	"ai-chat-backend/pkg/log"
	"ai-chat-backend/services"
	ai_chat_service "ai-chat-backend/services/ai-chat-service"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	ai_chat_service_proto "ai-chat-backend/services/ai-chat-service/proto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	openai "github.com/sashabaranov/go-openai"
	"k8s.io/klog/v2"
)

type ChatService struct {
	config *config.Config
	log    log.ILogger
}

type ChatCompletionParams struct {
	Model                 string        `json:"model"`
	MaxTokens             int           `json:"max_tokens,omitempty"`
	Temperature           float32       `json:"temperature,omitempty"`
	PresencePenalty       float32       `json:"presence_penalty,omitempty"`
	FrequencyPenalty      float32       `json:"frequency_penalty,omitempty"`
	ChatSessionTTL        time.Duration `json:"chat_session_ttl"`
	ChatMinResponseTokens int           `json:"chat_min_response_tokens"`
}

type ChatMessageRequest struct {
	Prompt  string                    `json:"prompt"`
	Options ChatMessageRequestOptions `json:"options"`
}

type ChatMessageRequestOptions struct {
	Name            string `json:"name"`
	ParentMessageId string `json:"parentMessageId"`
}

type ChatMessage struct {
	ID              string                                              `json:"id"`
	Text            string                                              `json:"text"`
	Role            string                                              `json:"role"`
	Name            string                                              `json:"name"`
	Delta           string                                              `json:"delta"`
	Detail          *ai_chat_service_proto.ChatCompletionStreamResponse `json:"detail"`
	TokenCount      int                                                 `json:"tokenCount"`
	ParentMessageId string                                              `json:"parentMessageId"`
}

func NewChatService(config *config.Config, log log.ILogger) (*ChatService, error) {
	return &ChatService{
		config: config,
		log:    log,
	}, nil
}

func (chat *ChatService) ChatProcess(ctx *gin.Context) {
	payload := ChatMessageRequest{}
	if err := ctx.BindJSON(&payload); err != nil {
		klog.Error(err)
		ctx.JSON(200, gin.H{
			"status":  "Fail",
			"message": fmt.Sprintf("%v", err),
			"data":    nil,
		})
		return
	}

	messageID := uuid.New().String()

	result := ChatMessage{
		ID:              uuid.New().String(),
		Role:            openai.ChatMessageRoleAssistant,
		Text:            "",
		ParentMessageId: messageID,
	}

	aiChatServicePool := ai_chat_service.GetAiChatServiceClientPool()
	conn := aiChatServicePool.Get()
	defer aiChatServicePool.Put(conn)
	ctx1 := services.AppendBearerTokenToContext(context.Background(), chat.config.DependOn.AiChatService.AccessToken)
	in := &ai_chat_service_proto.ChatCompletionRequest{
		Id:            messageID,
		Message:       payload.Prompt,
		Pid:           payload.Options.ParentMessageId,
		EnableContext: false,
		ChatParam: &ai_chat_service_proto.ChatParam{
			Model:             chat.config.Chat.Model,
			MaxTokens:         int32(chat.config.Chat.MaxTokens),
			Temperature:       chat.config.Chat.Temperature,
			TopP:              chat.config.Chat.TopP,
			PresencePenalty:   chat.config.Chat.PresencePenalty,
			FrequencyPenalty:  chat.config.Chat.FrequencyPenalty,
			BotDesc:           chat.config.Chat.BotDesc,
			ContextTTL:        int32(chat.config.Chat.ContextTTL),
			ContextLen:        int32(chat.config.Chat.ContextLen),
			MinResponseTokens: int32(chat.config.Chat.MinResponseTokens),
		},
	}
	if in.Pid != "" {
		in.EnableContext = true
	}

	aiChatServiceClient := ai_chat_service_proto.NewChatClient(conn)
	stream, err := aiChatServiceClient.ChatCompletionStream(ctx1, in)
	if err != nil {
		chat.log.Error(err)
		ctx.JSON(200, gin.H{
			"status":  "Fail",
			"message": fmt.Sprintf("%v", err),
			"data":    nil,
		})
		return
	}
	defer stream.CloseSend()

	firstChunk := true
	ctx.Header("Content-type", "application/octet-stream")
	for {
		rsp, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return
		}

		if err != nil {
			klog.Error(err)
			ctx.JSON(200, gin.H{
				"status":  "Fail",
				"message": fmt.Sprintf("OpenAI Event Error %v", err),
				"data":    nil,
			})
			return
		}

		if rsp.Id != "" {
			result.ID = rsp.Id
		}

		if len(rsp.Choices) > 0 {
			content := rsp.Choices[0].Delta.Content
			result.Delta = content
			if len(content) > 0 {
				result.Text += content
			}
			result.Detail = rsp
		}

		bts, err := json.Marshal(result)
		if err != nil {
			klog.Error(err)
			ctx.JSON(200, gin.H{
				"status":  "Fail",
				"message": fmt.Sprintf("OpenAI Event Marshal Error %v", err),
				"data":    nil,
			})
			return
		}

		if !firstChunk {
			ctx.Writer.Write([]byte("\n"))
		} else {
			firstChunk = false
		}

		if _, err := ctx.Writer.Write(bts); err != nil {
			klog.Error(err)
			return
		}

		ctx.Writer.Flush()
	}
}
