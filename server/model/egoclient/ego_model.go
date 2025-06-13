// 自动生成模板EgoModel
package egoclient

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 模型 结构体  EgoModel
type EgoModel struct {
	global.GVA_MODEL
	ModelProvider *string `json:"modelProvider" form:"modelProvider" gorm:"column:model_provider;"` //模型供应商
	ModelType     *string `json:"modelType" form:"modelType" gorm:"column:model_type;"`             //服务类型
	ModelName     *string `json:"modelName" form:"modelName" gorm:"column:model_name;"`             //模型名称
}

// TableName 模型 EgoModel自定义表名 ego-model
func (EgoModel) TableName() string {
	return "ego-model"
}
