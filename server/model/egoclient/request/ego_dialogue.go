package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type EgoDialogueSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	UUID           *string     `json:"uuid" form:"uuid"`
	User           *int        `json:"user" form:"user"`
	request.PageInfo
}

type EgoDialogueMultiModal struct {
	Type string `json:"type" form:"type"`
	Url  string `json:"url" form:"url"`
}

type EgoDialoguePostUserMsg struct {
	DialogueID string                  `json:"dialogue_id" form:"dialogueId"`
	Text       string                  `json:"text" form:"text"`
	Reasoning  bool                    `json:"reasoning" form:"reasoning"`
	Multimodal []EgoDialogueMultiModal `json:"multimodal" form:"multimodal"`
}
