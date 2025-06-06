package egoclient

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EgoClientUserRouter struct{}

func (s *EgoClientUserRouter) InitEgoClientUserRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	ECURouter := Router.Group("ECU").Use(middleware.OperationRecord())
	ECURouterWithoutRecord := Router.Group("ECU")
	ECURouterWithoutAuth := PublicRouter.Group("ECU")
	{
		ECURouter.POST("createEgoClientUser", ECUApi.CreateEgoClientUser)
		ECURouter.DELETE("deleteEgoClientUser", ECUApi.DeleteEgoClientUser)
		ECURouter.DELETE("deleteEgoClientUserByIds", ECUApi.DeleteEgoClientUserByIds)
		ECURouter.PUT("updateEgoClientUser", ECUApi.UpdateEgoClientUser)
		ECURouter.PUT("adminChangePassword", ECUApi.AdminChangePassword)
	}
	{
		ECURouterWithoutRecord.GET("findEgoClientUser", ECUApi.FindEgoClientUser)
		ECURouterWithoutRecord.GET("getEgoClientUserList", ECUApi.GetEgoClientUserList)
		ECURouterWithoutRecord.GET("getEgoClientUserPublic", ECUApi.GetEgoClientUserPublic)
		ECURouterWithoutRecord.GET("getUserInfo", ECUApi.GetUserInfo)
	}
	{
		ECURouterWithoutAuth.POST("register", ECUApi.Register)
		ECURouterWithoutAuth.POST("login", ECUApi.Login)
	}
}
