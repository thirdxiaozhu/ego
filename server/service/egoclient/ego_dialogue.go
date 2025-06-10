package egoclient

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	egoclientReq "github.com/flipped-aurora/gin-vue-admin/server/model/egoclient/request"
)

type EgoDialogueService struct{}

// CreateEgoDialogue 创建Ego对话记录
// Author [yourname](https://github.com/yourname)
func (EDService *EgoDialogueService) CreateEgoDialogue(ctx context.Context, ED *egoclient.EgoDialogue) (err error) {
	err = global.GVA_DB.Create(ED).Error
	return err
}

// DeleteEgoDialogue 删除Ego对话记录
// Author [yourname](https://github.com/yourname)
func (EDService *EgoDialogueService) DeleteEgoDialogue(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&egoclient.EgoDialogue{}, "id = ?", ID).Error
	return err
}

// DeleteEgoDialogueByIds 批量删除Ego对话记录
// Author [yourname](https://github.com/yourname)
func (EDService *EgoDialogueService) DeleteEgoDialogueByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]egoclient.EgoDialogue{}, "id in ?", IDs).Error
	return err
}

// UpdateEgoDialogue 更新Ego对话记录
// Author [yourname](https://github.com/yourname)
func (EDService *EgoDialogueService) UpdateEgoDialogue(ctx context.Context, ED egoclient.EgoDialogue) (err error) {
	err = global.GVA_DB.Model(&egoclient.EgoDialogue{}).Where("id = ?", ED.ID).Updates(&ED).Error
	return err
}

// GetEgoDialogue 根据ID获取Ego对话记录
// Author [yourname](https://github.com/yourname)
func (EDService *EgoDialogueService) GetEgoDialogue(ctx context.Context, ID string) (ED egoclient.EgoDialogue, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&ED).Error
	return
}

// GetEgoDialogueInfoList 分页获取Ego对话记录
// Author [yourname](https://github.com/yourname)
func (EDService *EgoDialogueService) GetEgoDialogueInfoList(ctx context.Context, info egoclientReq.EgoDialogueSearch) (list []egoclient.EgoDialogue, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&egoclient.EgoDialogue{})
	var EDs []egoclient.EgoDialogue
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.UUID != nil && *info.UUID != "" {
		db = db.Where("uuid = ?", *info.UUID)
	}
	if info.User != nil {
		db = db.Where("user = ?", *info.User)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&EDs).Error
	return EDs, total, err
}
func (EDService *EgoDialogueService) GetEgoDialoguePublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
