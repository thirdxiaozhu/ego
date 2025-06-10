// 自动生成模板EgoDialogue
package egoclient

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/google/uuid"
)

// Ego对话 结构体  EgoDialogue
type EgoDialogue struct {
	global.GVA_MODEL
	UUID    uuid.UUID `json:"uuid" form:"uuid" gorm:"column:uuid;"`       //对话UUID
	User    *int      `json:"user" form:"user" gorm:"column:user;"`       //所属用户
	ModelID int       `json:"model-id" form:"model" gorm:"column:model;"` //模型
	Model   EgoModel  `json:"model" gorm:"foreignKey:ID;references:ModelID;"`
}

// TableName Ego对话 EgoDialogue自定义表名 ego_dialogue
func (EgoDialogue) TableName() string {
	return "ego_dialogue"
}
