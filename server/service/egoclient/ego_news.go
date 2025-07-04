package egoclient

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	egoclientReq "github.com/flipped-aurora/gin-vue-admin/server/model/egoclient/request"
)

type EgoNewsService struct{}

// CreateEgoNews 创建Ego新闻推送记录
// Author [yourname](https://github.com/yourname)
func (ENService *EgoNewsService) CreateEgoNews(ctx context.Context, EN *egoclient.EgoNews) (err error) {
	err = global.GVA_DB.Create(EN).Error
	return err
}

// DeleteEgoNews 删除Ego新闻推送记录
// Author [yourname](https://github.com/yourname)
func (ENService *EgoNewsService) DeleteEgoNews(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&egoclient.EgoNews{}, "id = ?", ID).Error
	return err
}

// DeleteEgoNewsByIds 批量删除Ego新闻推送记录
// Author [yourname](https://github.com/yourname)
func (ENService *EgoNewsService) DeleteEgoNewsByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]egoclient.EgoNews{}, "id in ?", IDs).Error
	return err
}

// UpdateEgoNews 更新Ego新闻推送记录
// Author [yourname](https://github.com/yourname)
func (ENService *EgoNewsService) UpdateEgoNews(ctx context.Context, EN egoclient.EgoNews) (err error) {
	err = global.GVA_DB.Model(&egoclient.EgoNews{}).Where("id = ?", EN.ID).Updates(&EN).Error
	return err
}

// GetEgoNews 根据ID获取Ego新闻推送记录
// Author [yourname](https://github.com/yourname)
func (ENService *EgoNewsService) GetEgoNews(ctx context.Context, ID string) (EN egoclient.EgoNews, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&EN).Error
	return
}

// GetEgoNewsInfoList 分页获取Ego新闻推送记录
// Author [yourname](https://github.com/yourname)
func (ENService *EgoNewsService) GetEgoNewsInfoList(ctx context.Context, info egoclientReq.EgoNewsSearch) (list []egoclient.EgoNews, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&egoclient.EgoNews{})
	var ENs []egoclient.EgoNews
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.PublisherID != nil {
		db = db.Where("publisher_id = ?", *info.PublisherID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&ENs).Error
	return ENs, total, err
}
func (ENService *EgoNewsService) GetEgoNewsPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
