// 自动生成模板EgoModel
package egoclient

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/liusuxian/go-aisdk/consts"
)

// 模型 结构体  EgoModel
type EgoModel struct {
	global.GVA_MODEL
	ModelProvider consts.Provider  `json:"modelProvider" form:"modelProvider" gorm:"column:model_provider;"` //模型供应商
	ModelType     consts.ModelType `json:"modelType" form:"modelType" gorm:"column:model_type;"`             //服务类型
	ModelName     *string          `json:"modelName" form:"modelName" gorm:"column:model_name;"`             //模型名称
	NormalTimes   int              `json:"normalTimes" form:"normalTimes" gorm:"column:normal_times;default:-1;"`
	Limits        []EgoModelLimits `json:"limits" form:"limits" gorm:"foreignKey: ModelID;"`
}

// TableName 模型 EgoModel自定义表名 ego-model
func (EgoModel) TableName() string {
	return "ego_model"
}

type EgoModelLimits struct {
	global.GVA_MODEL
	ModelID    uint `json:"modelID" grom:"primaryKey"`
	VipLevelID uint `json:"levelID" grom:"primaryKey"`
	CallLimits int  `json:"callLimits" form:"callLimits" gorm:"column:call_limits;"`
}
