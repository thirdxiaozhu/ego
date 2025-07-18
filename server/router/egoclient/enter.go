package egoclient

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	EgoClientUserRouter
	EgoDialogueRouter
	EgoModelRouter
	EgoNewsRouter
	EgoNoramlAgentRouter
	EgoPromptRouter
}

var (
	ECUApi    = api.ApiGroupApp.EgoclientApiGroup.EgoClientUserApi
	EDApi     = api.ApiGroupApp.EgoclientApiGroup.EgoDialogueApi
	eModelApi = api.ApiGroupApp.EgoclientApiGroup.EgoModelApi
	ENApi     = api.ApiGroupApp.EgoclientApiGroup.EgoNewsApi
	ENAApi    = api.ApiGroupApp.EgoclientApiGroup.EgoNoramlAgentApi
	EPApi     = api.ApiGroupApp.EgoclientApiGroup.EgoPromptApi
)
