package egoclient

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EgoNoramlAgentRouter struct{}

// InitEgoNoramlAgentRouter 初始化 EGO普通智能体 路由信息
func (s *EgoNoramlAgentRouter) InitEgoNoramlAgentRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	ENARouter := Router.Group("ENA").Use(middleware.OperationRecord())
	ENARouterWithoutRecord := Router.Group("ENA")
	ENARouterWithoutAuth := PublicRouter.Group("ENA")
	{
		ENARouter.POST("createEgoNoramlAgent", ENAApi.CreateEgoNoramlAgent)             // 新建EGO普通智能体  (admin / user)
		ENARouter.DELETE("deleteEgoNoramlAgent", ENAApi.DeleteEgoNoramlAgent)           // 删除EGO普通智能体 (admin / user)
		ENARouter.DELETE("deleteEgoNoramlAgentByIds", ENAApi.DeleteEgoNoramlAgentByIds) // 批量删除EGO普通智能体
		ENARouter.PUT("updateEgoNoramlAgent", ENAApi.UpdateEgoNoramlAgent)              // 更新EGO普通智能体
	}
	{
		ENARouterWithoutRecord.GET("findEgoNoramlAgent", ENAApi.FindEgoNoramlAgent)               // 根据ID获取EGO普通智能体
		ENARouterWithoutRecord.GET("getEgoNoramlAgentList", ENAApi.GetEgoNoramlAgentList)         // 获取EGO普通智能体列表 (admin)
		ENARouterWithoutRecord.GET("getEgoNoramlAgentListUser", ENAApi.GetEgoNoramlAgentListUser) // 获取EGO普通智能体列表 (user)

	}
	{
		ENARouterWithoutAuth.GET("getEgoNoramlAgentPublic", ENAApi.GetEgoNoramlAgentPublic) // EGO普通智能体开放接口
	}
}
