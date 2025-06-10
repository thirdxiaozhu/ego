package egoclient

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/egoclient"
	egoclientReq "github.com/flipped-aurora/gin-vue-admin/server/model/egoclient/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EgoModelApi struct{}

// CreateEgoModel 创建模型
// @Tags EgoModel
// @Summary 创建模型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body egoclient.EgoModel true "创建模型"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /eModel/createEgoModel [post]
func (eModelApi *EgoModelApi) CreateEgoModel(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var eModel egoclient.EgoModel
	err := c.ShouldBindJSON(&eModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = eModelService.CreateEgoModel(ctx, &eModel)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteEgoModel 删除模型
// @Tags EgoModel
// @Summary 删除模型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body egoclient.EgoModel true "删除模型"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /eModel/deleteEgoModel [delete]
func (eModelApi *EgoModelApi) DeleteEgoModel(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := eModelService.DeleteEgoModel(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteEgoModelByIds 批量删除模型
// @Tags EgoModel
// @Summary 批量删除模型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /eModel/deleteEgoModelByIds [delete]
func (eModelApi *EgoModelApi) DeleteEgoModelByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := eModelService.DeleteEgoModelByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateEgoModel 更新模型
// @Tags EgoModel
// @Summary 更新模型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body egoclient.EgoModel true "更新模型"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /eModel/updateEgoModel [put]
func (eModelApi *EgoModelApi) UpdateEgoModel(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var eModel egoclient.EgoModel
	err := c.ShouldBindJSON(&eModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = eModelService.UpdateEgoModel(ctx, eModel)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindEgoModel 用id查询模型
// @Tags EgoModel
// @Summary 用id查询模型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询模型"
// @Success 200 {object} response.Response{data=egoclient.EgoModel,msg=string} "查询成功"
// @Router /eModel/findEgoModel [get]
func (eModelApi *EgoModelApi) FindEgoModel(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reeModel, err := eModelService.GetEgoModel(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reeModel, c)
}

// GetEgoModelList 分页获取模型列表
// @Tags EgoModel
// @Summary 分页获取模型列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query egoclientReq.EgoModelSearch true "分页获取模型列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /eModel/getEgoModelList [get]
func (eModelApi *EgoModelApi) GetEgoModelList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo egoclientReq.EgoModelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := eModelService.GetEgoModelInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithData(list, c)
}

// GetEgoModelAll 分页获取模型列表
// @Tags EgoModel
// @Summary 分页获取模型列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query egoclientReq.EgoModelSearch true "分页获取模型列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /eModel/getEgoModelAll [get]
func (eModelApi *EgoModelApi) GetEgoModelAll(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	list, total, err := eModelService.GetEgoModelInfoList(ctx, pageInfo)
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

// GetEgoModelPublic 不需要鉴权的模型接口
// @Tags EgoModel
// @Summary 不需要鉴权的模型接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /eModel/getEgoModelPublic [get]
func (eModelApi *EgoModelApi) GetEgoModelPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	eModelService.GetEgoModelPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的模型接口信息",
	}, "获取成功", c)
}
