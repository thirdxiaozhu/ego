package egoclient

import (
	"context"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	egoclientReq "github.com/flipped-aurora/gin-vue-admin/server/model/egoclient/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/google/uuid"
	"time"
)

type EgoClientUserService struct{}

// CreateEgoClientUser 创建EGO用户记录
// Author [yourname](https://github.com/yourname)
func (ECUService *EgoClientUserService) CreateEgoClientUser(ctx context.Context, ECU *egoclient.EgoClientUser) (err error) {
	var count int64

	if err = global.GVA_DB.Model(&egoclient.EgoClientUser{}).Where("user_id = ?", ECU.UserID).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已注册")
	}
	pwdHash := utils.BcryptHash(*ECU.Password)
	ECU.Password = &pwdHash
	ECU.UUID, _ = uuid.NewV6()

	ECU.VipStatus = egoclient.EgoVipStatus{
		ActivatedAt: time.Now(),
		VipLevelID:  1,
	}

	err = global.GVA_DB.Create(ECU).Error
	return err
}

// DeleteEgoClientUser 删除EGO用户记录
// Author [yourname](https://github.com/yourname)
func (ECUService *EgoClientUserService) DeleteEgoClientUser(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&egoclient.EgoClientUser{}, "id = ?", ID).Error
	return err
}

// DeleteEgoClientUserByIds 批量删除EGO用户记录
// Author [yourname](https://github.com/yourname)
func (ECUService *EgoClientUserService) DeleteEgoClientUserByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]egoclient.EgoClientUser{}, "id in ?", IDs).Error
	return err
}

// UpdateEgoClientUser 更新EGO用户记录
// Author [yourname](https://github.com/yourname)
func (ECUService *EgoClientUserService) UpdateEgoClientUser(ctx context.Context, ECU egoclient.EgoClientUser) (err error) {

	ECU.Password = nil
	ECU.UserID = nil

	fmt.Println("<UNK>", ECU.VipStatus.VipLevelID)
	//Update方法自动过滤空值
	if err = global.GVA_DB.Model(&egoclient.EgoClientUser{}).Where("id = ?", ECU.ID).Updates(&ECU).Error; err != nil {
		return err
	}
	if err = global.GVA_DB.Model(&egoclient.EgoVipStatus{}).Where("user_id = ?", ECU.ID).Updates(&ECU.VipStatus).Error; err != nil {
		return err
	}

	return err
}

// GetEgoClientUser 根据ID获取EGO用户记录
// Author [yourname](https://github.com/yourname)
func (ECUService *EgoClientUserService) GetEgoClientUser(ctx context.Context, ID string) (ECU egoclient.EgoClientUser, err error) {
	err = global.GVA_DB.Where("id = ?", ID).Preload("EgoVipStatus").Preload("EgoVipStatus.EgoVipLevel").First(&ECU).Error
	return
}

// GetEgoClientUserInfoList 分页获取EGO用户记录
// Author [yourname](https://github.com/yourname)
func (ECUService *EgoClientUserService) GetEgoClientUserInfoList(ctx context.Context, info egoclientReq.EgoClientUserSearch) (list []egoclient.EgoClientUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&egoclient.EgoClientUser{})
	var ECUs []egoclient.EgoClientUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.UserID != nil && *info.UserID != "" {
		db = db.Where("user_id = ?", *info.UserID)
	}
	if info.Password != nil && *info.Password != "" {
		db = db.Where("password = ?", *info.Password)
	}
	if info.Avatar != "" {
		// TODO 数据类型为复杂类型，请根据业务需求自行实现复杂类型的查询业务
	}
	if info.Gender != nil && *info.Gender != "" {
		db = db.Where("gender = ?", *info.Gender)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&ECUs).Error
	return ECUs, total, err
}
func (ECUService *EgoClientUserService) GetEgoClientUserPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// AdminChangePassword 管理员修改密码
// Author [yourname](https://github.com/yourname)
func (ECUService *EgoClientUserService) AdminChangePassword(ctx context.Context, req egoclientReq.AdminChangePasswordReq) (err error) {
	// 请在这里实现自己的业务逻辑

	pwdHash := utils.BcryptHash(*req.Password)
	req.Password = &pwdHash
	db := global.GVA_DB.Model(&egoclient.EgoClientUser{}).Where("user_id = ?", req.UserID).Update("password", req.Password)
	return db.Error
}

// Login 用户登录
// Author [yourname](https://github.com/yourname)
func (ECUService *EgoClientUserService) Login(ctx context.Context, UserID, password *string) (*egoclient.EgoClientUser, error) {
	// 请在这里实现自己的业务逻辑
	var err error
	var user egoclient.EgoClientUser
	err = global.GVA_DB.First(&user, "user_id = ?", UserID).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	if utils.BcryptCheck(*password, *user.Password) == true {
		return &user, nil
	} else {
		return nil, errors.New("密码错误")
	}
}

// GetUserInfo 获取用户信息
// Author [yourname](https://github.com/yourname)
func (ECUService *EgoClientUserService) GetUserInfo(ctx context.Context, id uint) (user egoclient.EgoClientUser, err error) {
	//主键是id
	err = global.GVA_DB.First(&user, id).Error
	return
}

// Logout 登出
// Author [yourname](https://github.com/yourname)
func (ECUService *EgoClientUserService) Logout(ctx context.Context) (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&egoclient.EgoClientUser{})
	return db.Error
}

// GetEgoClientUser 根据ID获取EGO用户记录
// Author [yourname](https://github.com/yourname)
func (ECUService *EgoClientUserService) GetEgoVipLevels() (ECVL []egoclient.EgoVipLevel, err error) {
	err = global.GVA_DB.Model(&egoclient.EgoVipLevel{}).Find(&ECVL).Error
	return
}
