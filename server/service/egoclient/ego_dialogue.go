package egoclient

import (
	"context"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	egoclientReq "github.com/flipped-aurora/gin-vue-admin/server/model/egoclient/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service/egoclient/egoModels"
	"github.com/google/uuid"
)

type EgoDialogueService struct{}

// CreateEgoDialogue 创建Ego对话记录
// Author [yourname](https://github.com/yourname)
func (EDService *EgoDialogueService) CreateEgoDialogue(ctx context.Context, userid uint, ED *egoclient.EgoDialogue) (err error) {
	ED.UserID = userid
	ED.UUID, _ = uuid.NewV6()
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
	err = global.GVA_DB.Where("id = ?", ID).Preload("Model").Preload("User").First(&ED).Error
	return
}

// GetEgoDialogueByUuid 根据ID获取Ego对话记录
// Author [yourname](https://github.com/yourname)
func (EDService *EgoDialogueService) GetEgoDialogueByUuid(ctx context.Context, Uuid string) (ED egoclient.EgoDialogue, err error) {
	err = global.GVA_DB.Where("uuid = ?", Uuid).Preload("Model").Preload("User").First(&ED).Error
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
		db = db.Limit(limit).Offset(offset).Preload("Model").Preload("User")
	}

	err = db.Find(&EDs).Error
	return EDs, total, err
}
func (EDService *EgoDialogueService) GetEgoDialoguePublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// PostEgoDialogueUserMsg 创建Ego对话记录
// Author [yourname](https://github.com/yourname)
func (EDService *EgoDialogueService) PostEgoDialogueUserMsg(ctx context.Context, Req *egoclientReq.EgoDialoguePostUserMsg) error {
	ED, err := EDService.GetEgoDialogueByUuid(ctx, Req.DialogueID)
	if err != nil {
		return errors.New("无法找到对话")
	}

	err = egoModels.AssembleRequest(&ED, Req)
	if err != nil {
		return err
	}

	return nil
}
