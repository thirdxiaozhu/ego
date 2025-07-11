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

type EgoNewsApi struct{}

// CreateEgoNews 创建Ego新闻推送
// @Tags EgoNews
// @Summary 创建Ego新闻推送
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body egoclient.EgoNews true "创建Ego新闻推送"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /EN/createEgoNews [post]
func (ENApi *EgoNewsApi) CreateEgoNews(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	userid := utils.GetUserID(c)

	var EN egoclient.EgoNews
	err := c.ShouldBindJSON(&EN)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	EN.PublisherID = &userid
	err = ENService.CreateEgoNews(ctx, &EN)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteEgoNews 删除Ego新闻推送
// @Tags EgoNews
// @Summary 删除Ego新闻推送
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body egoclient.EgoNews true "删除Ego新闻推送"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /EN/deleteEgoNews [delete]
func (ENApi *EgoNewsApi) DeleteEgoNews(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := ENService.DeleteEgoNews(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteEgoNewsByIds 批量删除Ego新闻推送
// @Tags EgoNews
// @Summary 批量删除Ego新闻推送
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /EN/deleteEgoNewsByIds [delete]
func (ENApi *EgoNewsApi) DeleteEgoNewsByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := ENService.DeleteEgoNewsByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateEgoNews 更新Ego新闻推送
// @Tags EgoNews
// @Summary 更新Ego新闻推送
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body egoclient.EgoNews true "更新Ego新闻推送"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /EN/updateEgoNews [put]
func (ENApi *EgoNewsApi) UpdateEgoNews(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var EN egoclient.EgoNews
	err := c.ShouldBindJSON(&EN)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = ENService.UpdateEgoNews(ctx, EN)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindEgoNews 用id查询Ego新闻推送
// @Tags EgoNews
// @Summary 用id查询Ego新闻推送
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询Ego新闻推送"
// @Success 200 {object} response.Response{data=egoclient.EgoNews,msg=string} "查询成功"
// @Router /EN/findEgoNews [get]
func (ENApi *EgoNewsApi) FindEgoNews(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reEN, err := ENService.GetEgoNews(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reEN, c)
}

// GetEgoNewsList 分页获取Ego新闻推送列
// GetEgoNewsList 分页获取Ego新闻推送列表
// @Tags EgoNews
// @Summary 分页获取Ego新闻推送列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query egoclientReq.EgoNewsSearch true "分页获取Ego新闻推送列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /EN/getEgoNewsList [get]
func (ENApi *EgoNewsApi) GetEgoNewsList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo egoclientReq.EgoNewsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := ENService.GetEgoNewsInfoList(ctx, pageInfo)
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

// GetEgoNewsPublic 不需要鉴权的Ego新闻推送接口
// @Tags EgoNews
// @Summary 不需要鉴权的Ego新闻推送接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /EN/getEgoNewsPublic [get]
func (ENApi *EgoNewsApi) GetEgoNewsPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	ENService.GetEgoNewsPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的Ego新闻推送接口信息",
	}, "获取成功", c)
}
