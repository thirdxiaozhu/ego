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
