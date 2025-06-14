package egoclient

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EgoDialogueRouter struct{}

// InitEgoDialogueRouter 初始化 Ego对话 路由信息
func (s *EgoDialogueRouter) InitEgoDialogueRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	EDRouter := Router.Group("ED").Use(middleware.OperationRecord())
	EDRouterWithoutRecord := Router.Group("ED")
	EDRouterWithoutAuth := PublicRouter.Group("ED")
	{
		EDRouter.POST("createEgoDialogue", EDApi.CreateEgoDialogue) // 新建Ego对话
		EDRouter.POST("postEgoDialogueUserMsg", EDApi.PostEgoDialogueUserMsg)
		EDRouter.DELETE("deleteEgoDialogue", EDApi.DeleteEgoDialogue)           // 删除Ego对话
		EDRouter.DELETE("deleteEgoDialogueByIds", EDApi.DeleteEgoDialogueByIds) // 批量删除Ego对话
		EDRouter.PUT("updateEgoDialogue", EDApi.UpdateEgoDialogue)              // 更新Ego对话
	}
	{
		EDRouterWithoutRecord.GET("findEgoDialogue", EDApi.FindEgoDialogue)       // 根据ID获取Ego对话
		EDRouterWithoutRecord.GET("getEgoDialogueList", EDApi.GetEgoDialogueList) // 获取Ego对话列表
	}
	{
		EDRouterWithoutAuth.GET("getEgoDialoguePublic", EDApi.GetEgoDialoguePublic) // Ego对话开放接口
	}
}
