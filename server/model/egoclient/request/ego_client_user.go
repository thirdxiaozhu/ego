package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type EgoClientUserSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	UserID         *string     `json:"userID" form:"userID"`
	Password       *string     `json:"password" form:"password"`
	Avatar         string      `json:"avatar" form:"avatar"`
	Gender         *string     `json:"gender" form:"gender"`
	request.PageInfo
}

type AdminChangePasswordReq struct {
	UserID   string  `json:"userID" form:"userID"`
	Password *string `json:"password" form:"password"`
}

type UserAction struct {
	UserID   *string `json:"userID" form:"userID"`
	Password *string `json:"password" form:"password"`
	Username *string `json:"username" form:"username"`
}
