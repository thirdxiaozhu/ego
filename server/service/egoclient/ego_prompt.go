package egoclient

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	egoclientReq "github.com/flipped-aurora/gin-vue-admin/server/model/egoclient/request"
	"github.com/google/uuid"
)

type EgoPromptService struct{}

// CreateEgoPrompt 创建Ego提示词记忆记录
// Author [yourname](https://github.com/yourname)
func (EPService *EgoPromptService) CreateEgoPrompt(ctx context.Context, EP *egoclient.EgoPrompt) (err error) {
	EP.UUID, _ = uuid.NewV6()
	err = global.GVA_DB.Create(EP).Error
	return err
}

// DeleteEgoPrompt 删除Ego提示词记忆记录
// Author [yourname](https://github.com/yourname)
func (EPService *EgoPromptService) DeleteEgoPrompt(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&egoclient.EgoPrompt{}, "id = ?", ID).Error
	return err
}

// DeleteEgoPromptByIds 批量删除Ego提示词记忆记录
// Author [yourname](https://github.com/yourname)
func (EPService *EgoPromptService) DeleteEgoPromptByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]egoclient.EgoPrompt{}, "id in ?", IDs).Error
	return err
}

// UpdateEgoPrompt 更新Ego提示词记忆记录
// Author [yourname](https://github.com/yourname)
func (EPService *EgoPromptService) UpdateEgoPrompt(ctx context.Context, EP egoclient.EgoPrompt) (err error) {
	err = global.GVA_DB.Model(&egoclient.EgoPrompt{}).Where("id = ?", EP.ID).Updates(&EP).Error
	return err
}

// GetEgoPrompt 根据ID获取Ego提示词记忆记录
// Author [yourname](https://github.com/yourname)
func (EPService *EgoPromptService) GetEgoPrompt(ctx context.Context, ID string) (EP egoclient.EgoPrompt, err error) {
	err = global.GVA_DB.Where("id = ?", ID).Preload("Owner").First(&EP).Error
	return
}

// GetEgoPromptByOwner 根据ID获取Ego提示词记忆记录
// Author [yourname](https://github.com/yourname)
func (EPService *EgoPromptService) GetEgoPromptByOwner(ctx context.Context, ownerID uint) (EP egoclient.EgoPrompt, err error) {
	err = global.GVA_DB.Where("owner_id = ?", ownerID).Preload("Owner").First(&EP).Error
	return
}

// GetEgoPromptInfoList 分页获取Ego提示词记忆记录
// Author [yourname](https://github.com/yourname)
func (EPService *EgoPromptService) GetEgoPromptInfoList(ctx context.Context, info egoclientReq.EgoPromptSearch) (list []egoclient.EgoPrompt, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&egoclient.EgoPrompt{})
	var EPs []egoclient.EgoPrompt
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.OwnerID != nil {
		db = db.Where("owner_id = ?", *info.OwnerID)
	}
	if info.UUID != nil && *info.UUID != "" {
		db = db.Where("uuid = ?", *info.UUID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Preload("Owner").Find(&EPs).Error
	return EPs, total, err
}
func (EPService *EgoPromptService) GetEgoPromptPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
