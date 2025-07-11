package egoclient

type EgoClientUserRouter struct{}

//func (s *EgoClientUserRouter) InitEgoClientUserRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
//	ECURouter := Router.Group("ECU").Use(middleware.OperationRecord())
//	ECURouterWithoutRecord := Router.Group("ECU")
//	ECURouterWithoutAuth := PublicRouter.Group("ECU")
//	{
//		ECURouter.POST("createEgoClientUser", ECUApi.CreateEgoClientUser)
//		ECURouter.DELETE("deleteEgoClientUser", ECUApi.DeleteEgoClientUser)
//		ECURouter.DELETE("deleteEgoClientUserByIds", ECUApi.DeleteEgoClientUserByIds)
//		ECURouter.PUT("updateEgoClientUser", ECUApi.UpdateEgoClientUser)
//		ECURouter.PUT("adminChangePassword", ECUApi.AdminChangePassword)
//	}
//	{
//		ECURouterWithoutRecord.GET("findEgoClientUser", ECUApi.FindEgoClientUser)
//		ECURouterWithoutRecord.GET("getEgoClientUserList", ECUApi.GetEgoClientUserList)
//		ECURouterWithoutRecord.GET("getEgoClientUserPublic", ECUApi.GetEgoClientUserPublic)
//		ECURouterWithoutRecord.GET("getUserInfo", ECUApi.GetUserInfo)
//	}
//	{
//		ECURouterWithoutAuth.POST("register", ECUApi.Register)
//		ECURouterWithoutAuth.POST("login", ECUApi.Login)
//		ECURouterWithoutAuth.POST("logout", ECUApi.Logout)
//	}
//}
