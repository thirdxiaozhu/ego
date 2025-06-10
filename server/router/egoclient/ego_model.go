package egoclient

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EgoModelRouter struct{}

// InitEgoModelRouter 初始化 模型 路由信息
func (s *EgoModelRouter) InitEgoModelRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	eModelRouter := Router.Group("eModel").Use(middleware.OperationRecord())
	eModelRouterWithoutRecord := Router.Group("eModel")
	eModelRouterWithoutAuth := PublicRouter.Group("eModel")
	{
		eModelRouter.POST("createEgoModel", eModelApi.CreateEgoModel)             // 新建模型
		eModelRouter.DELETE("deleteEgoModel", eModelApi.DeleteEgoModel)           // 删除模型
		eModelRouter.DELETE("deleteEgoModelByIds", eModelApi.DeleteEgoModelByIds) // 批量删除模型
		eModelRouter.PUT("updateEgoModel", eModelApi.UpdateEgoModel)              // 更新模型
	}
	{
		eModelRouterWithoutRecord.GET("findEgoModel", eModelApi.FindEgoModel)       // 根据ID获取模型
		eModelRouterWithoutRecord.GET("getEgoModelList", eModelApi.GetEgoModelList) // 获取模型列表
		eModelRouterWithoutRecord.GET("getEgoModelAll", eModelApi.GetEgoModelAll)   // 获取所有模型
	}
	{
		eModelRouterWithoutAuth.GET("getEgoModelPublic", eModelApi.GetEgoModelPublic) // 模型开放接口
	}
}
