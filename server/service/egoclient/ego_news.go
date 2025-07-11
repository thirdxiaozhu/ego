package egoclient

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	egoclientReq "github.com/flipped-aurora/gin-vue-admin/server/model/egoclient/request"
	"time"
)

type EgoNewsService struct{}

// CreateEgoNews 创建Ego新闻推送记录
// Author [yourname](https://github.com/yourname)
func (ENService *EgoNewsService) CreateEgoNews(ctx context.Context, EN *egoclient.EgoNews) (err error) {
	EN.CreatedAt = time.Now()
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
	err = global.GVA_DB.Where("id = ?", ID).Preload("Publisher").First(&EN).Error
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

	err = db.Preload("Publisher").Order("release_time desc").Find(&ENs).Error
	return ENs, total, err
}
func (ENService *EgoNewsService) GetEgoNewsPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// GetUserNewsViews 获取未读新闻ID
func (ENService *EgoNewsService) GetUserNewsViews(ctx context.Context, userID uint) (list []egoclient.EgoUserNewsView, err error) {
	db := global.GVA_DB.Model(&egoclient.EgoUserNewsView{})

	err = db.Where("user_id = ?", userID).Find(&list).Error
	return
}

// GetUnreadNewsByUser 用户获取三天内发布的未读新闻列表
func (ENService *EgoNewsService) GetUnreadNewsByUser(ctx context.Context, userID uint) (list []egoclient.EgoNews, err error) {
	var unreadNews []egoclient.EgoNews

	threeDaysAgo := time.Now().AddDate(0, 0, -3)

	err = global.GVA_DB.
		Where("release_time >= ?", threeDaysAgo).
		Where("id NOT IN (?)", global.GVA_DB.
			Model(&egoclient.EgoUserNewsView{}).
			Where("user_id = ?", userID).
			Select("news_id")).
		Preload("Publisher").
		Order("release_time DESC").
		Find(&unreadNews).Error

	return unreadNews, err
}

// GetNewsDetailByUser 用户获取新闻推送详情
func (ENService *EgoNewsService) GetNewsDetailByUser(ctx context.Context) {
	// 首先在连接表插入一条该新闻已被当前用户读过

}
