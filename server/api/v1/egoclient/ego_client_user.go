package egoclient

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EgoClientUserApi struct{}

// CreateEgoClientUser 创建EGO用户
// @Tags EgoClientUser
// @Summary 创建EGO用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body egoclient.EgoClientUser true "创建EGO用户"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /ECU/createEgoClientUser [post]
func (ECUApi *EgoClientUserApi) CreateEgoClientUser(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var ECU egoclient.EgoClientUser
	err := c.ShouldBindJSON(&ECU)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = ECUService.CreateEgoClientUser(ctx, &ECU)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteEgoClientUser 删除EGO用户
// @Tags EgoClientUser
// @Summary 删除EGO用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body egoclient.EgoClientUser true "删除EGO用户"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /ECU/deleteEgoClientUser [delete]
func (ECUApi *EgoClientUserApi) DeleteEgoClientUser(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := ECUService.DeleteEgoClientUser(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteEgoClientUserByIds 批量删除EGO用户
// @Tags EgoClientUser
// @Summary 批量删除EGO用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /ECU/deleteEgoClientUserByIds [delete]
func (ECUApi *EgoClientUserApi) DeleteEgoClientUserByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := ECUService.DeleteEgoClientUserByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateEgoClientUser 更新EGO用户
// @Tags EgoClientUser
// @Summary 更新EGO用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body egoclient.EgoClientUser true "更新EGO用户"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /ECU/updateEgoClientUser [put]
func (ECUApi *EgoClientUserApi) UpdateEgoClientUser(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var ECU egoclient.EgoClientUser
	err := c.ShouldBindJSON(&ECU)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = ECUService.UpdateEgoClientUser(ctx, ECU)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindEgoClientUser 用id查询EGO用户
// @Tags EgoClientUser
// @Summary 用id查询EGO用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询EGO用户"
// @Success 200 {object} response.Response{data=egoclient.EgoClientUser,msg=string} "查询成功"
// @Router /ECU/findEgoClientUser [get]
func (ECUApi *EgoClientUserApi) FindEgoClientUser(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reECU, err := ECUService.GetEgoClientUser(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reECU, c)
}

// GetEgoClientUserList 分页获取EGO用户列表
// @Tags EgoClientUser
// @Summary 分页获取EGO用户列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query egoclientReq.EgoClientUserSearch true "分页获取EGO用户列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /ECU/getEgoClientUserList [get]
func (ECUApi *EgoClientUserApi) GetEgoClientUserList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo request.EgoClientUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := ECUService.GetEgoClientUserInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetEgoClientUserPublic 不需要鉴权的EGO用户接口
// @Tags EgoClientUser
// @Summary 不需要鉴权的EGO用户接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ECU/getEgoClientUserPublic [get]
func (ECUApi *EgoClientUserApi) GetEgoClientUserPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	ECUService.GetEgoClientUserPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的EGO用户接口信息",
	}, "获取成功", c)
}

// AdminChangePassword 管理员修改密码
// @Tags EgoClientUser
// @Summary 管理员修改密码
// @Accept application/json
// @Produce application/json
// @Param data query egoclientReq.EgoClientUserSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /ECU/adminChangePassword [PUT]
func (ECUApi *EgoClientUserApi) AdminChangePassword(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()
	var req request.AdminChangePasswordReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 请添加自己的业务逻辑
	err = ECUService.AdminChangePassword(ctx, req)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData("返回数据", c)
}

// Register 用户注
// Register 用户注册
// @Tags EgoClientUser
// @Summary 用户注册
// @Accept application/json
// @Produce application/json
// @Param data query egoclientReq.EgoClientUserSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
func (ECUApi *EgoClientUserApi) Register(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var regInfo request.UserAction
	err := c.ShouldBindJSON(&regInfo)

	body, _ := c.GetRawData()
	global.GVA_LOG.Info("ECU", zap.Any("Body", string(body)))
	global.GVA_LOG.Info("ECU", zap.Any("ecu", regInfo))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var ecu egoclient.EgoClientUser
	ecu.Username = regInfo.Username
	ecu.Password = regInfo.Password
	ecu.UserID = regInfo.UserID

	err = ECUService.CreateEgoClientUser(ctx, &ecu)
	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithMessage("注册失败", c)
		return
	}
	response.OkWithData("注册成功", c)
}

// Login 用户登录
// @Tags EgoClientUser
// @Summary 用户登录
// @Accept application/json
// @Produce application/json
// @Param data query egoclientReq.EgoClientUserSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
func (ECUApi *EgoClientUserApi) Login(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()
	var loginInfo request.UserAction
	err := c.ShouldBindJSON(&loginInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 请添加自己的业务逻辑
	var user *egoclient.EgoClientUser
	if user, err = ECUService.Login(ctx, loginInfo.UserID, loginInfo.Password); user == nil || err != nil {
		response.FailWithMessage("登录失败", c)
		return
	}
	token, claims, err := utils.LoginToken(user)

	maxAge := int(claims.RegisteredClaims.ExpiresAt.Unix() - time.Now().Unix())
	utils.SetToken(c, token, maxAge)
	if err != nil {
		response.FailWithMessage("登录失败", c)
		return
	}
	response.OkWithData(gin.H{
		"token":  token,
		"claims": claims,
		"user":   user,
	}, c)
}

// GetUserInfo 获取用户信息
// @Tags EgoClientUser
// @Summary 获取用户信息
// @Accept application/json
// @Produce application/json
// @Param data query egoclientReq.EgoClientUserSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
func (ECUApi *EgoClientUserApi) GetUserInfo(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()
	id := utils.GetUserID(c)
	if id == 0 {
		response.FailWithMessage("失败", c)
		return
	}
	// 请添加自己的业务逻辑
	user, err := ECUService.GetUserInfo(ctx, id)
	if err != nil {
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData(user, c)
}

// Logout 登出
// @Tags EgoClientUser
// @Summary 登出
// @Accept application/json
// @Produce application/json
// @Param data query egoclientReq.EgoClientUserSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /ECU/logout [POST]
func (ECUApi *EgoClientUserApi) Logout(c *gin.Context) {
	utils.ClearToken(c)

	response.Ok(c)
}
