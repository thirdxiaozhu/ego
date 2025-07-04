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
	Limits        []EgoModelLimit  `json:"limits" form:"limits" gorm:"foreignKey: ModelID;"`
}

// TableName 模型 EgoModel自定义表名 ego-model
func (EgoModel) TableName() string {
	return "ego_model"
}

type EgoModelLimit struct {
	ID         uint `gorm:"primarykey" json:"ID"` // 主键ID
	ModelID    uint `json:"modelID" gorm:"primaryKey"`
	VipLevelID uint `json:"levelID" gorm:"primaryKey"`
	CallLimits int  `json:"callLimits" form:"callLimits" gorm:"column:call_limits;"`
}

//type EgoModelRecord struct {
//	//ID        uint      `gorm:"primarykey" json:"ID"` // 主键ID
//	UserID    uint      `json:"userID" gorm:"primaryKey"`
//	ModelID   uint      `json:"modelID" gorm:"primaryKey"`
//	Date      time.Time `gorm:"type:date;primaryKey"` // 只存储日期部分
//	CallTimes uint      `json:"callTimes" form:"callTimes" gorm:"column:call_times;"`
//
//	Model EgoModel `json:"model" form:"model" gorm:"foreignKey: ModelID;"`
//}
