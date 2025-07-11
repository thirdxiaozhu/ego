import service from '@/utils/request'
// @Tags EgoNoramlAgent
// @Summary 创建EGO普通智能体
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EgoNoramlAgent true "创建EGO普通智能体"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ENA/createEgoNoramlAgent [post]
export const createEgoNoramlAgent = (data) => {
  return service({
    url: '/ENA/createEgoNoramlAgent',
    method: 'post',
    data
  })
}

// @Tags EgoNoramlAgent
// @Summary 删除EGO普通智能体
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EgoNoramlAgent true "删除EGO普通智能体"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ENA/deleteEgoNoramlAgent [delete]
export const deleteEgoNoramlAgent = (params) => {
  return service({
    url: '/ENA/deleteEgoNoramlAgent',
    method: 'delete',
    params
  })
}

// @Tags EgoNoramlAgent
// @Summary 批量删除EGO普通智能体
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EGO普通智能体"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ENA/deleteEgoNoramlAgent [delete]
export const deleteEgoNoramlAgentByIds = (params) => {
  return service({
    url: '/ENA/deleteEgoNoramlAgentByIds',
    method: 'delete',
    params
  })
}

// @Tags EgoNoramlAgent
// @Summary 更新EGO普通智能体
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EgoNoramlAgent true "更新EGO普通智能体"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ENA/updateEgoNoramlAgent [put]
export const updateEgoNoramlAgent = (data) => {
  return service({
    url: '/ENA/updateEgoNoramlAgent',
    method: 'put',
    data
  })
}

// @Tags EgoNoramlAgent
// @Summary 用id查询EGO普通智能体
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.EgoNoramlAgent true "用id查询EGO普通智能体"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ENA/findEgoNoramlAgent [get]
export const findEgoNoramlAgent = (params) => {
  return service({
    url: '/ENA/findEgoNoramlAgent',
    method: 'get',
    params
  })
}

// @Tags EgoNoramlAgent
// @Summary 分页获取EGO普通智能体列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EGO普通智能体列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ENA/getEgoNoramlAgentList [get]
export const getEgoNoramlAgentList = (params) => {
  return service({
    url: '/ENA/getEgoNoramlAgentList',
    method: 'get',
    params
  })
}

// @Tags EgoNoramlAgent
// @Summary 不需要鉴权的EGO普通智能体接口
// @Accept application/json
// @Produce application/json
// @Param data query egoclientReq.EgoNoramlAgentSearch true "分页获取EGO普通智能体列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ENA/getEgoNoramlAgentPublic [get]
export const getEgoNoramlAgentPublic = () => {
  return service({
    url: '/ENA/getEgoNoramlAgentPublic',
    method: 'get',
  })
}
