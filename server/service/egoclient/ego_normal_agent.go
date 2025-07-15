package egoclient

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	egoclientReq "github.com/flipped-aurora/gin-vue-admin/server/model/egoclient/request"
	"github.com/google/uuid"
)

type EgoNoramlAgentService struct{}

// CreateEgoNoramlAgent 创建EGO普通智能体记录
// Author [yourname](https://github.com/yourname)
func (ENAService *EgoNoramlAgentService) CreateEgoNoramlAgent(ctx context.Context, ENA *egoclient.EgoNoramlAgent) (err error) {
	ENA.UUID, _ = uuid.NewV6()
	err = global.GVA_DB.Create(ENA).Error
	return err
}

// DeleteEgoNoramlAgent 删除EGO普通智能体记录
// Author [yourname](https://github.com/yourname)
func (ENAService *EgoNoramlAgentService) DeleteEgoNoramlAgent(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&egoclient.EgoNoramlAgent{}, "id = ?", ID).Error
	return err
}

// DeleteEgoNoramlAgentByIds 批量删除EGO普通智能体记录
// Author [yourname](https://github.com/yourname)
func (ENAService *EgoNoramlAgentService) DeleteEgoNoramlAgentByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]egoclient.EgoNoramlAgent{}, "id in ?", IDs).Error
	return err
}

// UpdateEgoNoramlAgent 更新EGO普通智能体记录
// Author [yourname](https://github.com/yourname)
func (ENAService *EgoNoramlAgentService) UpdateEgoNoramlAgent(ctx context.Context, ENA egoclient.EgoNoramlAgent) (err error) {
	err = global.GVA_DB.Model(&egoclient.EgoNoramlAgent{}).Where("id = ?", ENA.ID).Updates(&ENA).Error
	return err
}

// GetEgoNoramlAgent 根据ID获取EGO普通智能体记录
// Author [yourname](https://github.com/yourname)
func (ENAService *EgoNoramlAgentService) GetEgoNoramlAgent(ctx context.Context, ID string) (ENA egoclient.EgoNoramlAgent, err error) {
	err = global.GVA_DB.Where("id = ?", ID).Preload("Owner").First(&ENA).Error
	return
}

// GetEgoNoramlAgentInfoList 分页获取EGO普通智能体记录
// Author [yourname](https://github.com/yourname)
func (ENAService *EgoNoramlAgentService) GetEgoNoramlAgentInfoList(ctx context.Context, info egoclientReq.EgoNoramlAgentSearch) (list []egoclient.EgoNoramlAgent, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&egoclient.EgoNoramlAgent{})
	var ENAs []egoclient.EgoNoramlAgent
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.OwnerID != nil {
		db = db.Where("owner_id = ?", *info.OwnerID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Preload("Owner").Find(&ENAs).Error
	return ENAs, total, err
}

// GetEgoNoramlAgentInfoListUser 分页获取EGO普通智能体记录
// Author [yourname](https://github.com/yourname)
func (ENAService *EgoNoramlAgentService) GetEgoNoramlAgentInfoListUser(ctx context.Context, info egoclientReq.EgoNoramlAgentSearchUser) (list []egoclient.EgoNoramlAgent, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&egoclient.EgoNoramlAgent{})
	var ENAs []egoclient.EgoNoramlAgent

	err = db.Joins("Owner").Where("owner_id = ? OR (authority_id = ? AND is_private = ?)", info.OwnerID, 9991, 0).Count(&total).Error

	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Preload("Owner").Find(&ENAs).Error
	return ENAs, total, err
}
func (ENAService *EgoNoramlAgentService) GetEgoNoramlAgentPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
