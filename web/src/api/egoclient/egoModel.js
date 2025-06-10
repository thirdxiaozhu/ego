import service from '@/utils/request'
// @Tags EgoModel
// @Summary 创建模型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EgoModel true "创建模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /eModel/createEgoModel [post]
export const createEgoModel = (data) => {
  return service({
    url: '/eModel/createEgoModel',
    method: 'post',
    data
  })
}

// @Tags EgoModel
// @Summary 删除模型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EgoModel true "删除模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /eModel/deleteEgoModel [delete]
export const deleteEgoModel = (params) => {
  return service({
    url: '/eModel/deleteEgoModel',
    method: 'delete',
    params
  })
}

// @Tags EgoModel
// @Summary 批量删除模型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /eModel/deleteEgoModel [delete]
export const deleteEgoModelByIds = (params) => {
  return service({
    url: '/eModel/deleteEgoModelByIds',
    method: 'delete',
    params
  })
}

// @Tags EgoModel
// @Summary 更新模型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EgoModel true "更新模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /eModel/updateEgoModel [put]
export const updateEgoModel = (data) => {
  return service({
    url: '/eModel/updateEgoModel',
    method: 'put',
    data
  })
}

// @Tags EgoModel
// @Summary 用id查询模型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.EgoModel true "用id查询模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /eModel/findEgoModel [get]
export const findEgoModel = (params) => {
  return service({
    url: '/eModel/findEgoModel',
    method: 'get',
    params
  })
}

// @Tags EgoModel
// @Summary 分页获取模型列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取模型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eModel/getEgoModelList [get]
export const getEgoModelList = (params) => {
  return service({
    url: '/eModel/getEgoModelList',
    method: 'get',
    params
  })
}

// @Tags EgoModel
// @Summary 不需要鉴权的模型接口
// @Accept application/json
// @Produce application/json
// @Param data query egoclientReq.EgoModelSearch true "分页获取模型列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /eModel/getEgoModelPublic [get]
export const getEgoModelPublic = () => {
  return service({
    url: '/eModel/getEgoModelPublic',
    method: 'get',
  })
}
