// 自动生成模板EgoNoramlAgent
package egoclient

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/google/uuid"
)

// EGO普通智能体 结构体  EgoNoramlAgent
type EgoNoramlAgent struct {
	global.GVA_MODEL
	UUID         uuid.UUID      `json:"uuid" form:"uuid" gorm:"column:uuid"`
	OwnerID      uint           `json:"ownerID" form:"ownerID" gorm:"column:owner_id;"` //拥有者
	Owner        system.SysUser `json:"owner" form:"owner" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:OwnerID;"`
	Title        *string        `json:"title" form:"title" gorm:"column:title;"`
	SystemPrompt *string        `json:"systemPrompt" form:"systemPrompt" gorm:"column:system_prompt;"` //系统提示
	IsPrivate    bool           `json:"isPrivate" form:"isPrivate" gorm:"column:is_private;default:false;"`
}

// TableName EGO普通智能体 EgoNoramlAgent自定义表名 ego_normal_agent
func (EgoNoramlAgent) TableName() string {
	return "ego_normal_agent"
}
