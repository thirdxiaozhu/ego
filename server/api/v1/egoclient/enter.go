package egoclient

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	EgoClientUserApi
	EgoDialogueApi
	EgoModelApi
	EgoNewsApi
	EgoNoramlAgentApi
}

var (
	ECUService = service.ServiceGroupApp.EgoclientServiceGroup.EgoClientUserService
	EDService  = service.ServiceGroupApp.EgoclientServiceGroup.EgoDialogueService
	EMService  = service.ServiceGroupApp.EgoclientServiceGroup.EgoModelService
	ENService  = service.ServiceGroupApp.EgoclientServiceGroup.EgoNewsService
	ENAService = service.ServiceGroupApp.EgoclientServiceGroup.EgoNoramlAgentService
)
