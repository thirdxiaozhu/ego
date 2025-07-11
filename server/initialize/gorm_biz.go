package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

func bizModel() (err error) {
	db := global.GVA_DB
	//err = db.AutoMigrate(egoclient.EgoClientUser{}, egoclient.EgoDialogue{}, egoclient.EgoModel{}, egoclient.EgoDialogueHistory{}, egoclient.EgoDialogueItem{}, egoclient.EgoVipStatus{}, egoclient.EgoModelRecord{}, egoclient.EgoNews{}, egoclient.EgoUserNewsView{}, egoclient.EgoNoramlAgent{})
	err = db.AutoMigrate(system.EgoVipStatus{}, egoclient.EgoDialogue{}, egoclient.EgoModel{}, egoclient.EgoDialogueHistory{}, egoclient.EgoDialogueItem{}, egoclient.EgoModelRecord{}, egoclient.EgoNews{}, egoclient.EgoUserNewsView{}, egoclient.EgoNoramlAgent{})
	if err != nil {
		return
	}
	return
}
