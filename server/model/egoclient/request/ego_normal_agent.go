package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type EgoNoramlAgentSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	OwnerID        *int        `json:"ownerID" form:"ownerID"`
	request.PageInfo
}

type EgoNoramlAgentSearchUser struct {
	OwnerID uint `json:"ownerID" form:"ownerID"`
	request.PageInfo
}
