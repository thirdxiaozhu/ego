package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type EgoPromptSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	OwnerID        *int        `json:"ownerID" form:"ownerID"`
	UUID           *string     `json:"uuid" form:"uuid"`
	request.PageInfo
}
