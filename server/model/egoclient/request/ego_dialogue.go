package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/liusuxian/go-aisdk/models"
	"time"
)

type EgoDialogueSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	UUID           *string     `json:"uuid" form:"uuid"`
	User           *int        `json:"user" form:"user"`
	request.PageInfo
}

// EgoDialogueMultiModal 对话多模态
type EgoDialogueMultiModal struct {
	Type models.ChatUserMsgPartType `json:"type" form:"type"`
	Text string                     `json:"text" form:"text"`
	Url  string                     `json:"url" form:"url"`
}

// UserMsgChatOption 用户请求选项
type UserMsgChatOption struct {
	Reasoning     bool                    `json:"reasoning" form:"reasoning"`         // 是否开启推理
	ReasoningMode string                  `json:"reasoningMode" form:"reasoningMode"` // 推理模式
	WebSearch     bool                    `json:"webSearch" form:"webSearch"`         // 是否开启网页搜索
	Multimodal    []EgoDialogueMultiModal `json:"multimodal" form:"multimodal"`       // 多模态
}

// UserMsgImageOption 用户图片选项
type UserMsgImageOption struct {
	Size string `json:"size" form:"size"` // 图片大小
}

// EgoDialoguePostRequest 聊天请求
type EgoDialoguePostRequest struct {
	DialogueID  string              `json:"dialogue_id" form:"dialogueId"` // 对话ID
	Text        string              `json:"text" form:"text"`              // 文本
	ChatOption  *UserMsgChatOption  `json:"chat" form:"chat"`              // 聊天选项
	ImageOption *UserMsgImageOption `json:"image" form:"image"`            // 图片选项
}
