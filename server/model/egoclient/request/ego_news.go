package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type EgoNewsSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	PublisherID    *int        `json:"publisherID" form:"publisherID"`
	request.PageInfo
}
