package system

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type Login interface {
	GetUsername() string
	GetNickname() string
	GetUUID() uuid.UUID
	GetUserId() uint
	GetAuthorityId() uint
	GetUserInfo() any
}

var _ Login = new(SysUser)

type SysUser struct {
	global.GVA_MODEL
	UUID          uuid.UUID      `json:"uuid" gorm:"index;comment:用户UUID"`                                                                   // 用户UUID
	Username      string         `json:"userName" gorm:"index;comment:用户登录名"`                                                                // 用户登录名
	Password      string         `json:"-"  gorm:"comment:用户登录密码"`                                                                           // 用户登录密码
	NickName      string         `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                                          // 用户昵称
	HeaderImg     string         `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"`               // 用户头像
	AuthorityId   uint           `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                                                      // 用户角色ID
	Authority     SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`                        // 用户角色
	Authorities   []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`                                                   // 多用户角色
	Phone         string         `json:"phone"  gorm:"comment:用户手机号"`                                                                        // 用户手机号
	Email         string         `json:"email"  gorm:"comment:用户邮箱"`                                                                         // 用户邮箱
	Enable        int            `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`                                                    //用户是否被冻结 1正常 2冻结
	OriginSetting common.JSONMap `json:"originSetting" form:"originSetting" gorm:"type:text;default:null;column:origin_setting;comment:配置;"` //配置
	VipStatus     EgoVipStatus   `json:"vipStatus" form:"vipStatus" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;"`  //VIP状态
}

func (SysUser) TableName() string {
	return "sys_users"
}

func (s *SysUser) GetUsername() string {
	return s.Username
}

func (s *SysUser) GetNickname() string {
	return s.NickName
}

func (s *SysUser) GetUUID() uuid.UUID {
	return s.UUID
}

func (s *SysUser) GetUserId() uint {
	return s.ID
}

func (s *SysUser) GetAuthorityId() uint {
	return s.AuthorityId
}

func (s *SysUser) GetUserInfo() any {
	return *s
}

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
