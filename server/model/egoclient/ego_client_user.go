// 自动生成模板EgoClientUser
package egoclient

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	//"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/google/uuid"
)

//var _ system.Login = (*EgoClientUser)(nil)

// EGO用户 结构体  EgoClientUser
type EgoClientUser struct {
	global.GVA_MODEL
	UUID        uuid.UUID    `json:"uuid" form:"uuid" gorm:"column:uuid;" `                                                             //用户ID
	UserID      *string      `json:"userID" form:"userID" gorm:"column:user_id;" `                                                      //用户ID
	Password    *string      `json:"password" form:"password" gorm:"column:password;" `                                                 //密码
	Username    *string      `json:"username" form:"username" gorm:"default:新用户;column:username;"`                                      //用户名
	Avatar      string       `json:"avatar" form:"avatar" gorm:"column:avatar;"`                                                        //头像
	Gender      *string      `json:"gender" form:"gender" gorm:"column:gender;"`                                                        //性别
	Description *string      `json:"description" form:"description" gorm:"column:description;type:text;"`                               //用户简介
	AuthorityId uint         `json:"authorityId" gorm:"default:999;comment:用户角色ID"`                                                     // 用户角色ID
	VipStatus   EgoVipStatus `json:"vipStatus" form:"vipStatus" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;"` //VIP状态
}

//
//func (s *EgoClientUser) GetUsername() string {
//	return *s.UserID
//}
//
//func (s *EgoClientUser) GetNickname() string {
//	return *s.Username
//}
//
//func (s *EgoClientUser) GetUUID() uuid.UUID {
//	return s.UUID
//}
//
//func (s *EgoClientUser) GetUserId() uint {
//	return s.ID
//}
//
//func (s *EgoClientUser) GetAuthorityId() uint {
//	return s.AuthorityId
//}
//
//func (s *EgoClientUser) GetUserInfo() any {
//	return *s
//}
//
//func (EgoClientUser) TableName() string {
//	return "ego_client_users"
//}

type EgoVipStatus struct {
	global.GVA_MODEL
	UserID      uint       `json:"userID" gorm:"column:user_id;uniqueIndex"` // 一对一关系
	ActivatedAt time.Time  // VIP激活时间
	ExpiresAt   *time.Time // VIP过期时间
	//VipLevelID  uint        `json:"vipLevelID" gorm:"column:vip_level_id;default:1"` // 外键指向vip_levels表
	//VipLevel    EgoVipLevel `json:"vipLevel" gorm:"foreignKey:VipLevelID"`           // 关联VIP等级
	Points int `json:"points" gorm:"column:points;default:0"`
}

// UnmarshalJSON 自定义反序列化方法，处理字符串到整数的转换
func (v *EgoVipStatus) UnmarshalJSON(data []byte) error {
	type Alias EgoVipStatus
	aux := &struct {
		Points interface{} `json:"points"`
		*Alias
	}{
		Alias: (*Alias)(v),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// 处理 Points 字段的类型转换
	switch val := aux.Points.(type) {
	case float64:
		v.Points = int(val)
	case string:
		if val == "" {
			v.Points = 0
		} else {
			points, err := strconv.Atoi(val)
			if err != nil {
				return err
			}
			v.Points = points
		}
	case nil:
		v.Points = 0
	}

	return nil
}

func (EgoVipStatus) TableName() string {
	return "ego_vip_status"
}

//// EgoVipLevel VIP等级配置表 (可扩展的核心表)
//type EgoVipLevel struct {
//	global.GVA_MODEL
//	Name        string          `json:"name" gorm:"column:name;size:50;uniqueIndex"`      // 等级名称 (如: VIP1, VIP2)
//	Level       int             `json:"level" gorm:"column:level;uniqueIndex"`            // 等级数值 (1,2,3...)
//	Description string          `json:"description" gorm:"column:description;size:255"`   // 等级描述
//	IsDefault   bool            `json:"isDefault" gorm:"column:is_default;default:false"` // 是否默认等级
//	Limits      []EgoModelLimit `json:"limits" form:"limits" gorm:"foreignKey: VipLevelID;"`
//}
//
//func (EgoVipLevel) TableName() string {
//	return "ego_vip_level"
//}
