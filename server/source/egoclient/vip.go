package egoclient

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderVipLevels = system.InitOrderExternal + 1

type initVipLevels struct{}

// auto run
func init() {
	system.RegisterInit(initOrderVipLevels, &initVipLevels{})
}

func (i *initVipLevels) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&egoclient.EgoVipLevel{})
}

func (i *initVipLevels) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&egoclient.EgoVipLevel{})
}

func (i *initVipLevels) InitializerName() string {
	return egoclient.EgoVipLevel{}.TableName()
}

func (i *initVipLevels) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []egoclient.EgoVipLevel{
		{
			Name:        "VIP0",
			Level:       0,
			Description: "VIP0",
			IsDefault:   true,
		},
		{
			Name:        "VIP1",
			Level:       1,
			Description: "VIP1",
			IsDefault:   false,
		},
		{
			Name:        "VIP2",
			Level:       2,
			Description: "VIP2",
			IsDefault:   false,
		},
	}

	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysDictionary{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initVipLevels) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var record egoclient.EgoVipLevel
	if errors.Is(db.Where("name = ?", "VIP0").First(&record).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
