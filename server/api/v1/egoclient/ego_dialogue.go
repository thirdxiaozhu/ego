package egoclient

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	egoclientReq "github.com/flipped-aurora/gin-vue-admin/server/model/egoclient/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EgoDialogueApi struct{}

// CreateEgoDialogue 创建Ego对话
// @Tags EgoDialogue
// @Summary 创建Ego对话
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body egoclient.EgoDialogue true "创建Ego对话"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /ED/createEgoDialogue [post]
func (EDApi *EgoDialogueApi) CreateEgoDialogue(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var ED egoclient.EgoDialogue
	err := c.ShouldBindJSON(&ED)
	if err != nil {
		fmt.Println(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = EDService.CreateEgoDialogue(ctx, utils.GetUserID(c), &ED)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithData(ED, c)
}

// DeleteEgoDialogue 删除Ego对话
// @Tags EgoDialogue
// @Summary 删除Ego对话
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body egoclient.EgoDialogue true "删除Ego对话"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /ED/deleteEgoDialogue [delete]
func (EDApi *EgoDialogueApi) DeleteEgoDialogue(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := EDService.DeleteEgoDialogue(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteEgoDialogueByIds 批量删除Ego对话
// @Tags EgoDialogue
// @Summary 批量删除Ego对话
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /ED/deleteEgoDialogueByIds [delete]
func (EDApi *EgoDialogueApi) DeleteEgoDialogueByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := EDService.DeleteEgoDialogueByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateEgoDialogue 更新Ego对话
// @Tags EgoDialogue
// @Summary 更新Ego对话
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body egoclient.EgoDialogue true "更新Ego对话"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /ED/updateEgoDialogue [put]
func (EDApi *EgoDialogueApi) UpdateEgoDialogue(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var ED egoclient.EgoDialogue
	err := c.ShouldBindJSON(&ED)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = EDService.UpdateEgoDialogue(ctx, ED)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindEgoDialogue 用id查询Ego对话
// @Tags EgoDialogue
// @Summary 用id查询Ego对话
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询Ego对话"
// @Success 200 {object} response.Response{data=egoclient.EgoDialogue,msg=string} "查询成功"
// @Router /ED/findEgoDialogue [get]
func (EDApi *EgoDialogueApi) FindEgoDialogue(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reED, err := EDService.GetEgoDialogue(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reED, c)
}

// GetEgoDialogueList 分页获取Ego对话列表
// @Tags EgoDialogue
// @Summary 分页获取Ego对话列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query egoclientReq.EgoDialogueSearch true "分页获取Ego对话列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /ED/getEgoDialogueList [get]
func (EDApi *EgoDialogueApi) GetEgoDialogueList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo egoclientReq.EgoDialogueSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := EDService.GetEgoDialogueInfoList(ctx, pageInfo)
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

// GetEgoDialoguePublic 不需要鉴权的Ego对话接口
// @Tags EgoDialogue
// @Summary 不需要鉴权的Ego对话接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ED/getEgoDialoguePublic [get]
func (EDApi *EgoDialogueApi) GetEgoDialoguePublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	EDService.GetEgoDialoguePublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的Ego对话接口信息",
	}, "获取成功", c)
}

// PostEgoDialogueUserMsg 创建Ego对话的用户消息
// @Tags EgoDialogue
// @Summary 创建Ego对话
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body egoclient.EgoDialogue true "创建Ego对话"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /ED/createEgoDialogue [post]
func (EDApi *EgoDialogueApi) PostEgoDialogueUserMsg(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var ED egoclientReq.EgoDialoguePostRequest
	err := c.ShouldBindJSON(&ED)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = EDService.PostEgoDialogueUserMsg(ctx, &ED)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}
