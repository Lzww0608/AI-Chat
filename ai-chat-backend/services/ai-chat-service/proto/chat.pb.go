// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: proto/chat.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ChatCompletionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message       string     `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Id            string     `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Pid           string     `protobuf:"bytes,3,opt,name=pid,json=p_id,proto3" json:"pid,omitempty"`
	EnableContext bool       `protobuf:"varint,4,opt,name=enableContext,json=enable_context,proto3" json:"enableContext,omitempty"`
	ChatParam     *ChatParam `protobuf:"bytes,5,opt,name=chatParam,json=chat_param,proto3" json:"chatParam,omitempty"`
}

func (x *ChatCompletionRequest) Reset() {
	*x = ChatCompletionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatCompletionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatCompletionRequest) ProtoMessage() {}

func (x *ChatCompletionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatCompletionRequest.ProtoReflect.Descriptor instead.
func (*ChatCompletionRequest) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{0}
}

func (x *ChatCompletionRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ChatCompletionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChatCompletionRequest) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

func (x *ChatCompletionRequest) GetEnableContext() bool {
	if x != nil {
		return x.EnableContext
	}
	return false
}

func (x *ChatCompletionRequest) GetChatParam() *ChatParam {
	if x != nil {
		return x.ChatParam
	}
	return nil
}

type ChatParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Model             string  `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	MaxTokens         int32   `protobuf:"varint,2,opt,name=maxTokens,json=max_tokens,proto3" json:"maxTokens,omitempty"`
	Temperature       float32 `protobuf:"fixed32,3,opt,name=temperature,proto3" json:"temperature,omitempty"`
	TopP              float32 `protobuf:"fixed32,4,opt,name=topP,json=top_p,proto3" json:"topP,omitempty"`
	PresencePenalty   float32 `protobuf:"fixed32,5,opt,name=presencePenalty,json=presence_penalty,proto3" json:"presencePenalty,omitempty"`
	FrequencyPenalty  float32 `protobuf:"fixed32,6,opt,name=frequencyPenalty,json=frequency_penalty,proto3" json:"frequencyPenalty,omitempty"`
	BotDesc           string  `protobuf:"bytes,7,opt,name=botDesc,json=bot_desc,proto3" json:"botDesc,omitempty"`
	MinResponseTokens int32   `protobuf:"varint,8,opt,name=minResponseTokens,json=min_response_tokens,proto3" json:"minResponseTokens,omitempty"`
	ContextTTL        int32   `protobuf:"varint,9,opt,name=contextTTL,json=context_ttl,proto3" json:"contextTTL,omitempty"`
	ContextLen        int32   `protobuf:"varint,10,opt,name=contextLen,json=context_len,proto3" json:"contextLen,omitempty"`
}

func (x *ChatParam) Reset() {
	*x = ChatParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatParam) ProtoMessage() {}

func (x *ChatParam) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatParam.ProtoReflect.Descriptor instead.
func (*ChatParam) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{1}
}

func (x *ChatParam) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *ChatParam) GetMaxTokens() int32 {
	if x != nil {
		return x.MaxTokens
	}
	return 0
}

func (x *ChatParam) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *ChatParam) GetTopP() float32 {
	if x != nil {
		return x.TopP
	}
	return 0
}

func (x *ChatParam) GetPresencePenalty() float32 {
	if x != nil {
		return x.PresencePenalty
	}
	return 0
}

func (x *ChatParam) GetFrequencyPenalty() float32 {
	if x != nil {
		return x.FrequencyPenalty
	}
	return 0
}

func (x *ChatParam) GetBotDesc() string {
	if x != nil {
		return x.BotDesc
	}
	return ""
}

func (x *ChatParam) GetMinResponseTokens() int32 {
	if x != nil {
		return x.MinResponseTokens
	}
	return 0
}

func (x *ChatParam) GetContextTTL() int32 {
	if x != nil {
		return x.ContextTTL
	}
	return 0
}

func (x *ChatParam) GetContextLen() int32 {
	if x != nil {
		return x.ContextLen
	}
	return 0
}

// 服务响应消息，非流式响应
type ChatCompletionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Object  string                  `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	Created int64                   `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Model   string                  `protobuf:"bytes,4,opt,name=model,proto3" json:"model,omitempty"`
	Choices []*ChatCompletionChoice `protobuf:"bytes,5,rep,name=choices,proto3" json:"choices,omitempty"`
	Usage   *Usage                  `protobuf:"bytes,6,opt,name=usage,proto3" json:"usage,omitempty"`
}

func (x *ChatCompletionResponse) Reset() {
	*x = ChatCompletionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatCompletionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatCompletionResponse) ProtoMessage() {}

func (x *ChatCompletionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatCompletionResponse.ProtoReflect.Descriptor instead.
func (*ChatCompletionResponse) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{2}
}

func (x *ChatCompletionResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChatCompletionResponse) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *ChatCompletionResponse) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *ChatCompletionResponse) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *ChatCompletionResponse) GetChoices() []*ChatCompletionChoice {
	if x != nil {
		return x.Choices
	}
	return nil
}

func (x *ChatCompletionResponse) GetUsage() *Usage {
	if x != nil {
		return x.Usage
	}
	return nil
}

type ChatCompletionChoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index        int32                  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Message      *ChatCompletionMessage `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	FinishReason string                 `protobuf:"bytes,3,opt,name=finishReason,json=finish_reason,proto3" json:"finishReason,omitempty"`
}

func (x *ChatCompletionChoice) Reset() {
	*x = ChatCompletionChoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatCompletionChoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatCompletionChoice) ProtoMessage() {}

func (x *ChatCompletionChoice) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatCompletionChoice.ProtoReflect.Descriptor instead.
func (*ChatCompletionChoice) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{3}
}

func (x *ChatCompletionChoice) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *ChatCompletionChoice) GetMessage() *ChatCompletionMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *ChatCompletionChoice) GetFinishReason() string {
	if x != nil {
		return x.FinishReason
	}
	return ""
}

type ChatCompletionMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role    string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ChatCompletionMessage) Reset() {
	*x = ChatCompletionMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatCompletionMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatCompletionMessage) ProtoMessage() {}

func (x *ChatCompletionMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatCompletionMessage.ProtoReflect.Descriptor instead.
func (*ChatCompletionMessage) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{4}
}

func (x *ChatCompletionMessage) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *ChatCompletionMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ChatCompletionMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Usage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PromptTokens     int32 `protobuf:"varint,1,opt,name=promptTokens,json=prompt_tokens,proto3" json:"promptTokens,omitempty"`
	CompletionTokens int32 `protobuf:"varint,2,opt,name=completionTokens,json=completion_tokens,proto3" json:"completionTokens,omitempty"`
	TotalTokens      int32 `protobuf:"varint,3,opt,name=totalTokens,json=total_tokens,proto3" json:"totalTokens,omitempty"`
}

func (x *Usage) Reset() {
	*x = Usage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Usage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Usage) ProtoMessage() {}

func (x *Usage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Usage.ProtoReflect.Descriptor instead.
func (*Usage) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{5}
}

func (x *Usage) GetPromptTokens() int32 {
	if x != nil {
		return x.PromptTokens
	}
	return 0
}

func (x *Usage) GetCompletionTokens() int32 {
	if x != nil {
		return x.CompletionTokens
	}
	return 0
}

func (x *Usage) GetTotalTokens() int32 {
	if x != nil {
		return x.TotalTokens
	}
	return 0
}

// 服务响应消息，流式响应
type ChatCompletionStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Object  string                        `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	Created int64                         `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Model   string                        `protobuf:"bytes,4,opt,name=model,proto3" json:"model,omitempty"`
	Choices []*ChatCompletionStreamChoice `protobuf:"bytes,5,rep,name=choices,proto3" json:"choices,omitempty"`
}

func (x *ChatCompletionStreamResponse) Reset() {
	*x = ChatCompletionStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatCompletionStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatCompletionStreamResponse) ProtoMessage() {}

func (x *ChatCompletionStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatCompletionStreamResponse.ProtoReflect.Descriptor instead.
func (*ChatCompletionStreamResponse) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{6}
}

func (x *ChatCompletionStreamResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChatCompletionStreamResponse) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *ChatCompletionStreamResponse) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *ChatCompletionStreamResponse) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *ChatCompletionStreamResponse) GetChoices() []*ChatCompletionStreamChoice {
	if x != nil {
		return x.Choices
	}
	return nil
}

type ChatCompletionStreamChoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index        int32                            `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Delta        *ChatCompletionStreamChoiceDelta `protobuf:"bytes,2,opt,name=delta,proto3" json:"delta,omitempty"`
	FinishReason string                           `protobuf:"bytes,3,opt,name=finishReason,json=finish_reason,proto3" json:"finishReason,omitempty"`
}

func (x *ChatCompletionStreamChoice) Reset() {
	*x = ChatCompletionStreamChoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatCompletionStreamChoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatCompletionStreamChoice) ProtoMessage() {}

func (x *ChatCompletionStreamChoice) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatCompletionStreamChoice.ProtoReflect.Descriptor instead.
func (*ChatCompletionStreamChoice) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{7}
}

func (x *ChatCompletionStreamChoice) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *ChatCompletionStreamChoice) GetDelta() *ChatCompletionStreamChoiceDelta {
	if x != nil {
		return x.Delta
	}
	return nil
}

func (x *ChatCompletionStreamChoice) GetFinishReason() string {
	if x != nil {
		return x.FinishReason
	}
	return ""
}

type ChatCompletionStreamChoiceDelta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Role    string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *ChatCompletionStreamChoiceDelta) Reset() {
	*x = ChatCompletionStreamChoiceDelta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatCompletionStreamChoiceDelta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatCompletionStreamChoiceDelta) ProtoMessage() {}

func (x *ChatCompletionStreamChoiceDelta) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatCompletionStreamChoiceDelta.ProtoReflect.Descriptor instead.
func (*ChatCompletionStreamChoiceDelta) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{8}
}

func (x *ChatCompletionStreamChoiceDelta) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ChatCompletionStreamChoiceDelta) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

var File_proto_chat_proto protoreflect.FileDescriptor

var file_proto_chat_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x61, 0x69, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x7a, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x22, 0xc1,
	0x01, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x11, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x5f, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x44, 0x0a, 0x09,
	0x63, 0x68, 0x61, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x61, 0x69, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x7a, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x22, 0xdc, 0x02, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x13, 0x0a, 0x04, 0x74, 0x6f, 0x70, 0x50, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x5f, 0x70, 0x12, 0x29, 0x0a, 0x0f,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x10, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x5f,
	0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x12, 0x2b, 0x0a, 0x10, 0x66, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x79, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x11, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x70, 0x65, 0x6e,
	0x61, 0x6c, 0x74, 0x79, 0x12, 0x19, 0x0a, 0x07, 0x62, 0x6f, 0x74, 0x44, 0x65, 0x73, 0x63, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x12,
	0x2e, 0x0a, 0x11, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x6d, 0x69, 0x6e, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12,
	0x1f, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x54, 0x54, 0x4c, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x74, 0x6c,
	0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x4c, 0x65, 0x6e, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x6c, 0x65,
	0x6e, 0x22, 0xf5, 0x01, 0x0a, 0x16, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x4a, 0x0a, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x61, 0x69, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x7a, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73,
	0x12, 0x37, 0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x61, 0x69, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x7a, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x22, 0x9e, 0x01, 0x0a, 0x14, 0x43, 0x68,
	0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x6f, 0x69,
	0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x4b, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x61, 0x69, 0x5f, 0x63,
	0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x7a, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x59, 0x0a, 0x15, 0x43, 0x68,
	0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x7c, 0x0a, 0x05, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23,
	0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x12, 0x21, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x22, 0xc8, 0x01, 0x0a, 0x1c, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x50, 0x0a, 0x07,
	0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e,
	0x61, 0x69, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x7a, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43,
	0x68, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x22, 0xaa,
	0x01, 0x0a, 0x1a, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x51, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x61, 0x69, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x7a, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x52,
	0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0c, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x4f, 0x0a, 0x1f, 0x43,
	0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x32, 0x87, 0x02, 0x0a,
	0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x77, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x2e, 0x61, 0x69, 0x5f, 0x63, 0x68, 0x61,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x7a, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x61, 0x69, 0x5f,
	0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x7a, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x85,
	0x01, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x31, 0x2e, 0x61, 0x69, 0x5f, 0x63, 0x68, 0x61,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x7a, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x61, 0x69, 0x5f,
	0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x7a, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x17, 0x5a, 0x15, 0x61, 0x69, 0x2d, 0x63, 0x68, 0x61,
	0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_chat_proto_rawDescOnce sync.Once
	file_proto_chat_proto_rawDescData = file_proto_chat_proto_rawDesc
)

func file_proto_chat_proto_rawDescGZIP() []byte {
	file_proto_chat_proto_rawDescOnce.Do(func() {
		file_proto_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_chat_proto_rawDescData)
	})
	return file_proto_chat_proto_rawDescData
}

var file_proto_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_chat_proto_goTypes = []interface{}{
	(*ChatCompletionRequest)(nil),           // 0: ai_chat_service.zvoice.com.ChatCompletionRequest
	(*ChatParam)(nil),                       // 1: ai_chat_service.zvoice.com.ChatParam
	(*ChatCompletionResponse)(nil),          // 2: ai_chat_service.zvoice.com.ChatCompletionResponse
	(*ChatCompletionChoice)(nil),            // 3: ai_chat_service.zvoice.com.ChatCompletionChoice
	(*ChatCompletionMessage)(nil),           // 4: ai_chat_service.zvoice.com.ChatCompletionMessage
	(*Usage)(nil),                           // 5: ai_chat_service.zvoice.com.Usage
	(*ChatCompletionStreamResponse)(nil),    // 6: ai_chat_service.zvoice.com.ChatCompletionStreamResponse
	(*ChatCompletionStreamChoice)(nil),      // 7: ai_chat_service.zvoice.com.ChatCompletionStreamChoice
	(*ChatCompletionStreamChoiceDelta)(nil), // 8: ai_chat_service.zvoice.com.ChatCompletionStreamChoiceDelta
}
var file_proto_chat_proto_depIdxs = []int32{
	1, // 0: ai_chat_service.zvoice.com.ChatCompletionRequest.chatParam:type_name -> ai_chat_service.zvoice.com.ChatParam
	3, // 1: ai_chat_service.zvoice.com.ChatCompletionResponse.choices:type_name -> ai_chat_service.zvoice.com.ChatCompletionChoice
	5, // 2: ai_chat_service.zvoice.com.ChatCompletionResponse.usage:type_name -> ai_chat_service.zvoice.com.Usage
	4, // 3: ai_chat_service.zvoice.com.ChatCompletionChoice.message:type_name -> ai_chat_service.zvoice.com.ChatCompletionMessage
	7, // 4: ai_chat_service.zvoice.com.ChatCompletionStreamResponse.choices:type_name -> ai_chat_service.zvoice.com.ChatCompletionStreamChoice
	8, // 5: ai_chat_service.zvoice.com.ChatCompletionStreamChoice.delta:type_name -> ai_chat_service.zvoice.com.ChatCompletionStreamChoiceDelta
	0, // 6: ai_chat_service.zvoice.com.Chat.ChatCompletion:input_type -> ai_chat_service.zvoice.com.ChatCompletionRequest
	0, // 7: ai_chat_service.zvoice.com.Chat.ChatCompletionStream:input_type -> ai_chat_service.zvoice.com.ChatCompletionRequest
	2, // 8: ai_chat_service.zvoice.com.Chat.ChatCompletion:output_type -> ai_chat_service.zvoice.com.ChatCompletionResponse
	6, // 9: ai_chat_service.zvoice.com.Chat.ChatCompletionStream:output_type -> ai_chat_service.zvoice.com.ChatCompletionStreamResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_chat_proto_init() }
func file_proto_chat_proto_init() {
	if File_proto_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatCompletionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatParam); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatCompletionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatCompletionChoice); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatCompletionMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Usage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_chat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatCompletionStreamResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_chat_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatCompletionStreamChoice); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_chat_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatCompletionStreamChoiceDelta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_chat_proto_goTypes,
		DependencyIndexes: file_proto_chat_proto_depIdxs,
		MessageInfos:      file_proto_chat_proto_msgTypes,
	}.Build()
	File_proto_chat_proto = out.File
	file_proto_chat_proto_rawDesc = nil
	file_proto_chat_proto_goTypes = nil
	file_proto_chat_proto_depIdxs = nil
}
