package egoclient

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EgoPromptRouter struct{}

// InitEgoPromptRouter 初始化 Ego提示词记忆 路由信息
func (s *EgoPromptRouter) InitEgoPromptRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	EPRouter := Router.Group("EP").Use(middleware.OperationRecord())
	EPRouterWithoutRecord := Router.Group("EP")
	EPRouterWithoutAuth := PublicRouter.Group("EP")
	{
		EPRouter.POST("createEgoPrompt", EPApi.CreateEgoPrompt)             // 新建Ego提示词记忆
		EPRouter.DELETE("deleteEgoPrompt", EPApi.DeleteEgoPrompt)           // 删除Ego提示词记忆
		EPRouter.DELETE("deleteEgoPromptByIds", EPApi.DeleteEgoPromptByIds) // 批量删除Ego提示词记忆
		EPRouter.PUT("updateEgoPrompt", EPApi.UpdateEgoPrompt)              // 更新Ego提示词记忆
	}
	{
		EPRouterWithoutRecord.GET("findEgoPrompt", EPApi.FindEgoPrompt)               // 根据ID获取Ego提示词记忆
		EPRouterWithoutRecord.GET("findEgoPromptByOwner", EPApi.FindEgoPromptByOwner) // 根据Owner获取Ego提示词记忆
		EPRouterWithoutRecord.GET("getEgoPromptList", EPApi.GetEgoPromptList)         // 获取Ego提示词记忆列表
	}
	{
		EPRouterWithoutAuth.GET("getEgoPromptPublic", EPApi.GetEgoPromptPublic) // Ego提示词记忆开放接口
	}
}
