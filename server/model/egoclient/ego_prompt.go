// 自动生成模板EgoPrompt
package egoclient

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/google/uuid"
)

// Ego提示词记忆 结构体  EgoPrompt
type EgoPrompt struct {
	global.GVA_MODEL
	OwnerID uint           `json:"ownerID" form:"ownerID" gorm:"column:owner_id;"` //拥有者
	Owner   system.SysUser `json:"owner" form:"owner" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:OwnerID;"`
	UUID    uuid.UUID      `json:"uuid" form:"uuid" gorm:"column:uuid;"`       //UUID
	Prompt  *string        `json:"prompt" form:"prompt" gorm:"column:prompt;"` //提示词
}

// TableName Ego提示词记忆 EgoPrompt自定义表名 ego_prompt
func (EgoPrompt) TableName() string {
	return "ego_prompt"
}
