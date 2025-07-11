// 自动生成模板EgoNews
package egoclient

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Ego新闻推送 结构体  EgoNews
type EgoNews struct {
	global.GVA_MODEL
	ReleaseTime *time.Time    `json:"releaseTime" form:"releaseTime" gorm:"column:release_time;"` //发布时间
	PublisherID *uint         `json:"publisherID" form:"publisherID" gorm:"column:publisher_id;"` //发布者
	Publisher   EgoClientUser `json:"publisher" form:"publisher" gorm:"foreignKey:ID;references:PublisherID;"`
	Title       *string       `json:"title" form:"title" gorm:"column:title;"`
	Content     *string       `json:"content" form:"content" gorm:"column:content; type:text;"` //内容
}

// TableName Ego新闻推送 EgoNews自定义表名 ego_news
func (EgoNews) TableName() string {
	return "ego_news"
}

type EgoUserNewsView struct {
	UserID    uint      `gorm:"primaryKey"` // 复合主键
	NewsID    uint      `gorm:"primaryKey"`
	CreatedAt time.Time // 自动记录阅读时间
}
