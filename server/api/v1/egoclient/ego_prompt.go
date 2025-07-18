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

type EgoPromptApi struct{}

// CreateEgoPrompt 创建Ego提示词记忆
// @Tags EgoPrompt
// @Summary 创建Ego提示词记忆
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body egoclient.EgoPrompt true "创建Ego提示词记忆"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /EP/createEgoPrompt [post]
func (EPApi *EgoPromptApi) CreateEgoPrompt(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var EP egoclient.EgoPrompt
	err := c.ShouldBindJSON(&EP)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	EP.OwnerID = utils.GetUserID(c)
	err = EPService.CreateEgoPrompt(ctx, &EP)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteEgoPrompt 删除Ego提示词记忆
// @Tags EgoPrompt
// @Summary 删除Ego提示词记忆
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body egoclient.EgoPrompt true "删除Ego提示词记忆"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /EP/deleteEgoPrompt [delete]
func (EPApi *EgoPromptApi) DeleteEgoPrompt(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := EPService.DeleteEgoPrompt(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteEgoPromptByIds 批量删除Ego提示词记忆
// @Tags EgoPrompt
// @Summary 批量删除Ego提示词记忆
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /EP/deleteEgoPromptByIds [delete]
func (EPApi *EgoPromptApi) DeleteEgoPromptByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := EPService.DeleteEgoPromptByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateEgoPrompt 更新Ego提示词记忆
// @Tags EgoPrompt
// @Summary 更新Ego提示词记忆
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body egoclient.EgoPrompt true "更新Ego提示词记忆"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /EP/updateEgoPrompt [put]
func (EPApi *EgoPromptApi) UpdateEgoPrompt(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var EP egoclient.EgoPrompt
	err := c.ShouldBindJSON(&EP)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = EPService.UpdateEgoPrompt(ctx, EP)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindEgoPrompt 用id查询Ego提示词记忆
// @Tags EgoPrompt
// @Summary 用id查询Ego提示词记忆
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询Ego提示词记忆"
// @Success 200 {object} response.Response{data=egoclient.EgoPrompt,msg=string} "查询成功"
// @Router /EP/findEgoPrompt [get]
func (EPApi *EgoPromptApi) FindEgoPrompt(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reEP, err := EPService.GetEgoPrompt(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reEP, c)
}

// FindEgoPromptByOwner 用owner查询Ego提示词记忆
// @Tags EgoPrompt
// @Summary 用owner查询Ego提示词记忆
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用owner查询Ego提示词记忆"
// @Success 200 {object} response.Response{data=egoclient.EgoPrompt,msg=string} "查询成功"
// @Router /EP/findEgoPrompt [get]
func (EPApi *EgoPromptApi) FindEgoPromptByOwner(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	userID := utils.GetUserID(c)
	reEP, err := EPService.GetEgoPromptByOwner(ctx, userID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reEP, c)
}

// GetEgoPromptList 分页获取Ego提示词记忆列表
// @Tags EgoPrompt
// @Summary 分页获取Ego提示词记忆列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query egoclientReq.EgoPromptSearch true "分页获取Ego提示词记忆列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /EP/getEgoPromptList [get]
func (EPApi *EgoPromptApi) GetEgoPromptList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo egoclientReq.EgoPromptSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := EPService.GetEgoPromptInfoList(ctx, pageInfo)
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

// GetEgoPromptPublic 不需要鉴权的Ego提示词记忆接口
// @Tags EgoPrompt
// @Summary 不需要鉴权的Ego提示词记忆接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /EP/getEgoPromptPublic [get]
func (EPApi *EgoPromptApi) GetEgoPromptPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	EPService.GetEgoPromptPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的Ego提示词记忆接口信息",
	}, "获取成功", c)
}
