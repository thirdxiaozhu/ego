package egoclient

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	egoclientReq "github.com/flipped-aurora/gin-vue-admin/server/model/egoclient/request"
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
	err = global.GVA_DB.Model(&egoclient.EgoModel{}).Where("id = ?", eModel.ID).Updates(&eModel).Error
	return err
}

// GetEgoModel 根据ID获取模型记录
// Author [yourname](https://github.com/yourname)
func (eModelService *EgoModelService) GetEgoModel(ctx context.Context, ID string) (eModel egoclient.EgoModel, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&eModel).Error
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
