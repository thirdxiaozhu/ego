package egoclient

import (
	"context"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	egoclientReq "github.com/flipped-aurora/gin-vue-admin/server/model/egoclient/request"
	"gorm.io/gorm"
	"time"
)

type EgoModelService struct{}

// CreateEgoModel 创建模型记录
// Author [yourname](https://github.com/yourname)
func (eModelService *EgoModelService) CreateEgoModel(ctx context.Context, eModel *egoclient.EgoModel) (err error) {
	err = global.GVA_DB.Create(eModel).Error
	return err
}

// DeleteEgoModel 删除模型记录
// Author [yourname](https://github.com/yourname)
func (eModelService *EgoModelService) DeleteEgoModel(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&egoclient.EgoModel{}, "id = ?", ID).Error
	return err
}

// DeleteEgoModelByIds 批量删除模型记录
// Author [yourname](https://github.com/yourname)
func (eModelService *EgoModelService) DeleteEgoModelByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]egoclient.EgoModel{}, "id in ?", IDs).Error
	return err
}

// UpdateEgoModel 更新模型记录
// Author [yourname](https://github.com/yourname)
func (eModelService *EgoModelService) UpdateEgoModel(ctx context.Context, eModel egoclient.EgoModel) (err error) {

	tx := global.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 1. 更新主模型数据 (排除关联字段)
	if err = tx.Model(&egoclient.EgoModel{}).Where("id = ?", eModel.ID).
		Omit("Limits"). // 关键：排除关联字段
		Updates(&eModel).Error; err != nil {
		tx.Rollback()
		return err
	}

	//2. 处理关联的Limits数据
	if len(eModel.Limits) > 0 {
		// 先删除所有旧关联
		fmt.Println(eModel.ID)
		if err = tx.Where("model_id = ?", eModel.ID).
			Delete(&egoclient.EgoModelLimit{}).Error; err != nil {
			tx.Rollback()
			return err
		}

		// 设置关联ID并创建新数据
		for i := range eModel.Limits {
			eModel.Limits[i].ModelID = eModel.ID
		}
		if err = tx.Create(&eModel.Limits).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	//
	// 提交事务
	return tx.Commit().Error
}

// GetEgoModel 根据ID获取模型记录
// Author [yourname](https://github.com/yourname)
func (eModelService *EgoModelService) GetEgoModel(ctx context.Context, ID string) (eModel egoclient.EgoModel, err error) {
	//查询 EgoModel 时需要预加载关联的 Limits 并按 VipLevelID 升序排序。以下是完整实现方案：
	err = global.GVA_DB.Where("id = ?", ID).Preload("Limits", func(db *gorm.DB) *gorm.DB {
		return db.Order("ego_model_limits.vip_level_id ASC")
	}).First(&eModel).Error
	return
}

// GetEgoModelInfoList 分页获取模型记录
// Author [yourname](https://github.com/yourname)
func (eModelService *EgoModelService) GetEgoModelInfoList(ctx context.Context, info egoclientReq.EgoModelSearch) (list []egoclient.EgoModel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&egoclient.EgoModel{})
	var eModels []egoclient.EgoModel
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.ModelProvider != nil && *info.ModelProvider != "" {
		db = db.Where("model_provider = ?", *info.ModelProvider)
	}
	if info.ModelType != nil && *info.ModelType != "" {
		db = db.Where("model_type = ?", *info.ModelType)
	}
	if info.ModelName != nil && *info.ModelName != "" {
		db = db.Where("model_name = ?", *info.ModelName)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&eModels).Error
	return eModels, total, err
}
func (eModelService *EgoModelService) GetEgoModelPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// GetEgoModelInfoAll 分页获取模型记录
// Author [yourname](https://github.com/yourname)
func (eModelService *EgoModelService) GetEgoModelInfoAll(ctx context.Context) (list []egoclient.EgoModel, err error) {
	// 创建db
	db := global.GVA_DB.Model(&egoclient.EgoModel{})
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Find(&list).Error
	return
}

func (eModelService *EgoModelService) CreateCallRecord(EMRS egoclientReq.EgoModelRecordDefine) (err error) {
	//fmt.Println("!!!!!!!!!!!!1")
	//err = global.GVA_DB.Create().Error
	return
}

func (eModelService *EgoModelService) GetCallRecord(EMRS egoclient.EgoModelRecord) (record egoclient.EgoModelRecord, err error) {

	err = global.GVA_DB.Where(EMRS).Attrs(egoclient.EgoModelRecord{CallTimes: 0}).FirstOrCreate(&record).Error

	return
}

type ModelOperation func(*egoclient.EgoDialogue, *egoclientReq.EgoDialoguePostUserMsg) int

func (eModelService *EgoModelService) CanCallModel(ED *egoclient.EgoDialogue, Req *egoclientReq.EgoDialoguePostUserMsg, operation ModelOperation) (b bool, err error) {

	//当前是北京时间，如果未来要改成UTC时间，改为 time.Now().UTC()（国际化）
	now := time.Now()
	currDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	//tx := global.GVA_DB.Begin()
	//defer func() {
	//	if r := recover(); r != nil {
	//		tx.Rollback()
	//	}
	//}()

	record := egoclient.EgoModelRecord{
		UserID:  ED.User.ID,
		ModelID: ED.Model.ID,
		Date:    currDate,
	}
	record, err = eModelService.GetCallRecord(record)

	limits := 0
	for _, item := range ED.Model.Limits {
		if item.VipLevelID == ED.User.VipStatus.VipLevelID {
			limits = item.CallLimits
			break
		}
	}
	if limits == 0 {
		return false, errors.New("模型不可用")
	}

	if limits != -1 && record.CallTimes > uint(limits) {
		return false, errors.New("当日用量已达上限")
	}

	operation(ED, Req)

	//TODO:6.24 更新CallTimes+1
	fmt.Println(record)
	return
}
