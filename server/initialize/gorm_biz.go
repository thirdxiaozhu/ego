package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(egoclient.EgoClientUser{})
	if err != nil {
		return err
	}
	return nil
}
