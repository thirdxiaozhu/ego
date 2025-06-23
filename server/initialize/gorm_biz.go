package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
)

func bizModel() (err error) {
	db := global.GVA_DB
	//err = db.SetupJoinTable(egoclient.EgoModel{}, "VipLevels", egoclient.EgoModelLimit{})
	//if err != nil {
	//	return
	//}
	err = db.AutoMigrate(egoclient.EgoClientUser{}, egoclient.EgoDialogue{}, egoclient.EgoModel{}, egoclient.EgoDialogueHistory{}, egoclient.EgoDialogueItem{},
		egoclient.EgoVipStatus{}, egoclient.EgoVipLevel{}, egoclient.EgoModelLimit{})
	if err != nil {
		return
	}

	return
}
