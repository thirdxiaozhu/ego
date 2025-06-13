// 自动生成模板EgoClientUser
package egoclient

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/google/uuid"
)

var _ system.Login = (*EgoClientUser)(nil)

// EGO用户 结构体  EgoClientUser
type EgoClientUser struct {
	global.GVA_MODEL
	UUID        uuid.UUID `json:"uuid" form:"uuid" gorm:"column:uuid;" `                                //用户ID
	UserID      *string   `json:"userID" form:"userID" gorm:"column:user_id;" binding:"required"`       //用户ID
	Password    *string   `json:"password" form:"password" gorm:"column:password;" binding:"required"`  //密码
	Username    *string   `json:"username" form:"username" gorm:"default:新用户;column:username;"`         //用户名
	Avatar      string    `json:"avatar" form:"avatar" gorm:"column:avatar;"`                           //头像
	Gender      *string   `json:"gender" form:"gender" gorm:"column:gender;"`                           //性别
	Description *string   `json:"description" form:"description" gorm:"column:description;"type:text;"` //用户简介
	AuthorityId uint      `json:"authorityId" gorm:"default:999;comment:用户角色ID"`                        // 用户角色ID
}

func (s *EgoClientUser) GetUsername() string {
	return *s.UserID
}

func (s *EgoClientUser) GetNickname() string {
	return *s.Username
}

func (s *EgoClientUser) GetUUID() uuid.UUID {
	return s.UUID
}

func (s *EgoClientUser) GetUserId() uint {
	return s.ID
}

func (s *EgoClientUser) GetAuthorityId() uint {
	return s.AuthorityId
}

func (s *EgoClientUser) GetUserInfo() any {
	return *s
}
