// 自动生成模板EgoNews
package egoclient

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Ego新闻推送 结构体  EgoNews
type EgoNews struct {
	global.GVA_MODEL
	ReleaseTime *time.Time `json:"releaseTime" form:"releaseTime" gorm:"column:release_time;"` //发布时间
	PublisherID *int       `json:"publisherID" form:"publisherID" gorm:"column:publisher_id;"` //发布者
	Content     *string    `json:"content" form:"content" gorm:"column:content;"type:text;"`   //内容
}

// TableName Ego新闻推送 EgoNews自定义表名 ego_news
func (EgoNews) TableName() string {
	return "ego_news"
}
