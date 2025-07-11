package egoclient

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	egoclientReq "github.com/flipped-aurora/gin-vue-admin/server/model/egoclient/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EgoNoramlAgentApi struct{}

// CreateEgoNoramlAgent 创建EGO普通智能体
// @Tags EgoNoramlAgent
// @Summary 创建EGO普通智能体,创建者ID为当前登录用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body egoclient.EgoNoramlAgent true "创建EGO普通智能体"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /ENA/createEgoNoramlAgent [post]
func (ENAApi *EgoNoramlAgentApi) CreateEgoNoramlAgent(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var ENA egoclient.EgoNoramlAgent
	err := c.ShouldBindJSON(&ENA)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	ENA.OwnerID = utils.GetUserID(c)
	err = ENAService.CreateEgoNoramlAgent(ctx, &ENA)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteEgoNoramlAgent 删除EGO普通智能体
// @Tags EgoNoramlAgent
// @Summary 删除EGO普通智能体
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body egoclient.EgoNoramlAgent true "删除EGO普通智能体"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /ENA/deleteEgoNoramlAgent [delete]
func (ENAApi *EgoNoramlAgentApi) DeleteEgoNoramlAgent(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := ENAService.DeleteEgoNoramlAgent(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteEgoNoramlAgentByIds 批量删除EGO普通智能体
// @Tags EgoNoramlAgent
// @Summary 批量删除EGO普通智能体
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /ENA/deleteEgoNoramlAgentByIds [delete]
func (ENAApi *EgoNoramlAgentApi) DeleteEgoNoramlAgentByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := ENAService.DeleteEgoNoramlAgentByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateEgoNoramlAgent 更新EGO普通智能体
// @Tags EgoNoramlAgent
// @Summary 更新EGO普通智能体
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body egoclient.EgoNoramlAgent true "更新EGO普通智能体"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /ENA/updateEgoNoramlAgent [put]
func (ENAApi *EgoNoramlAgentApi) UpdateEgoNoramlAgent(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var ENA egoclient.EgoNoramlAgent
	err := c.ShouldBindJSON(&ENA)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = ENAService.UpdateEgoNoramlAgent(ctx, ENA)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindEgoNoramlAgent 用id查询EGO普通智能体
// @Tags EgoNoramlAgent
// @Summary 用id查询EGO普通智能体
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询EGO普通智能体"
// @Success 200 {object} response.Response{data=egoclient.EgoNoramlAgent,msg=string} "查询成功"
// @Router /ENA/findEgoNoramlAgent [get]
func (ENAApi *EgoNoramlAgentApi) FindEgoNoramlAgent(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reENA, err := ENAService.GetEgoNoramlAgent(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reENA, c)
}

// GetEgoNoramlAgentList 分页获取EGO普通智能体列表
// @Tags EgoNoramlAgent
// @Summary 分页获取EGO普通智能体列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query egoclientReq.EgoNoramlAgentSearch true "分页获取EGO普通智能体列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /ENA/getEgoNoramlAgentList [get]
func (ENAApi *EgoNoramlAgentApi) GetEgoNoramlAgentList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo egoclientReq.EgoNoramlAgentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := ENAService.GetEgoNoramlAgentInfoList(ctx, pageInfo)
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

// GetEgoNoramlAgentListUser 分页获取EGO普通智能体列表 (User)
// @Tags EgoNoramlAgent
// @Summary 分页获取EGO普通智能体列表 (User)
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query egoclientReq.EgoNoramlAgentSearch true "分页获取EGO普通智能体列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /ENA/getEgoNoramlAgentList [get]
func (ENAApi *EgoNoramlAgentApi) GetEgoNoramlAgentListUser(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var toSearch egoclientReq.EgoNoramlAgentSearchUser
	err := c.ShouldBindQuery(&toSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	toSearch.OwnerID = utils.GetUserID(c)
	list, total, err := ENAService.GetEgoNoramlAgentInfoListUser(ctx, toSearch)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     toSearch.Page,
		PageSize: toSearch.PageSize,
	}, "获取成功", c)
}

// GetEgoNoramlAgentPublic 不需要鉴权的EGO普通智能体接口
// @Tags EgoNoramlAgent
// @Summary 不需要鉴权的EGO普通智能体接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ENA/getEgoNoramlAgentPublic [get]
func (ENAApi *EgoNoramlAgentApi) GetEgoNoramlAgentPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	ENAService.GetEgoNoramlAgentPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的EGO普通智能体接口信息",
	}, "获取成功", c)
}
