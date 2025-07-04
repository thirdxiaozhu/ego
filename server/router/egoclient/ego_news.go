package egoclient

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EgoNewsRouter struct{}

// InitEgoNewsRouter 初始化 Ego新闻推送 路由信息
func (s *EgoNewsRouter) InitEgoNewsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	ENRouter := Router.Group("EN").Use(middleware.OperationRecord())
	ENRouterWithoutRecord := Router.Group("EN")
	ENRouterWithoutAuth := PublicRouter.Group("EN")
	{
		ENRouter.POST("createEgoNews", ENApi.CreateEgoNews)             // 新建Ego新闻推送
		ENRouter.DELETE("deleteEgoNews", ENApi.DeleteEgoNews)           // 删除Ego新闻推送
		ENRouter.DELETE("deleteEgoNewsByIds", ENApi.DeleteEgoNewsByIds) // 批量删除Ego新闻推送
		ENRouter.PUT("updateEgoNews", ENApi.UpdateEgoNews)              // 更新Ego新闻推送
	}
	{
		ENRouterWithoutRecord.GET("findEgoNews", ENApi.FindEgoNews)       // 根据ID获取Ego新闻推送
		ENRouterWithoutRecord.GET("getEgoNewsList", ENApi.GetEgoNewsList) // 获取Ego新闻推送列表
	}
	{
		ENRouterWithoutAuth.GET("getEgoNewsPublic", ENApi.GetEgoNewsPublic) // Ego新闻推送开放接口
	}
}
