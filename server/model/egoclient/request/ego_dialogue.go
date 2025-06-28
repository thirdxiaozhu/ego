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

type EgoDialogueMultiModal struct {
	Type models.ChatUserMsgPartType `json:"type" form:"type"`
	Text string                     `json:"text" form:"text"`
	Url  string                     `json:"url" form:"url"`
}

type UserMsgChatOption struct {
	Reasoning     bool                    `json:"reasoning" form:"reasoning"`
	ReasoningMode string                  `json:"reasoning_mode" form:"reasoningMode"`
	Multimodal    []EgoDialogueMultiModal `json:"multimodal" form:"multimodal"`
}

type UserMsgImageOption struct {
}

type EgoDialoguePostUserMsg struct {
	DialogueID  string              `json:"dialogue_id" form:"dialogueId"`
	Text        string              `json:"text" form:"text"`
	ChatOption  *UserMsgChatOption  `json:"chat" form:"chat"`
	ImageOption *UserMsgImageOption `json:"image" form:"image"`
}
