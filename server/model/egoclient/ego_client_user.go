// 自动生成模板EgoClientUser
package egoclient

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/google/uuid"
	"time"
)

var _ system.Login = (*EgoClientUser)(nil)

// EGO用户 结构体  EgoClientUser
type EgoClientUser struct {
	global.GVA_MODEL
	UUID        uuid.UUID `json:"uuid" form:"uuid" gorm:"column:uuid;" `                               //用户ID
	UserID      *string   `json:"userID" form:"userID" gorm:"column:user_id;" `                        //用户ID
	Password    *string   `json:"password" form:"password" gorm:"column:password;" `                   //密码
	Username    *string   `json:"username" form:"username" gorm:"default:新用户;column:username;"`        //用户名
	Avatar      string    `json:"avatar" form:"avatar" gorm:"column:avatar;"`                          //头像
	Gender      *string   `json:"gender" form:"gender" gorm:"column:gender;"`                          //性别
	Description *string   `json:"description" form:"description" gorm:"column:description;type:text;"` //用户简介
	AuthorityId uint      `json:"authorityId" gorm:"default:999;comment:用户角色ID"`                       // 用户角色ID

	VipStatusID uint      `json:"vipStatusID" gorm:"vip_status_id"`
	VipStatus   VipStatus `json:"vipStatus" form:"vipStatus" gorm:"foreignkey:VipStatusID;"`
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

type VipStatus struct {
	global.GVA_MODEL
	UserID uint `json:"user_id" gorm:"column: user_id"` // 一对一关系

	// VIP时效性字段
	ActivatedAt time.Time // VIP激活时间
	ExpiresAt   time.Time // VIP过期时间

	// 关联关系
	VipLevelID uint     `json:"vip_level_id" gorm:"column:"`            // 外键指向vip_levels表
	VipLevel   VipLevel `json:"vip_level" gorm:"foreignKey:VipLevelID"` // 关联VIP等级
}

func (VipStatus) TableName() string {
	return "ego_vip_status"
}

// VIP等级配置表 (可扩展的核心表)
type VipLevel struct {
	global.GVA_MODEL
	Name        string `json:"name" gorm:"column:name;size:50;uniqueIndex"`   // 等级名称 (如: VIP1, VIP2)
	Level       int    `json:"level" gorm:"column:level;uniqueIndex"`         // 等级数值 (1,2,3...)
	Description string `json:"description" orm:"column:description;size:255"` // 等级描述

	// 可扩展字段
	IsDefault bool `json:"is_default" gorm:"column:is_default;default:false"` // 是否默认等级
}

func (VipLevel) TableName() string {
	return "ego_vip_level"
}
