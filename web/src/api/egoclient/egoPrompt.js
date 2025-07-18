import service from '@/utils/request'
// @Tags EgoPrompt
// @Summary 创建Ego提示词记忆
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EgoPrompt true "创建Ego提示词记忆"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /EP/createEgoPrompt [post]
export const createEgoPrompt = (data) => {
  return service({
    url: '/EP/createEgoPrompt',
    method: 'post',
    data
  })
}

// @Tags EgoPrompt
// @Summary 删除Ego提示词记忆
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EgoPrompt true "删除Ego提示词记忆"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /EP/deleteEgoPrompt [delete]
export const deleteEgoPrompt = (params) => {
  return service({
    url: '/EP/deleteEgoPrompt',
    method: 'delete',
    params
  })
}

// @Tags EgoPrompt
// @Summary 批量删除Ego提示词记忆
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Ego提示词记忆"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /EP/deleteEgoPrompt [delete]
export const deleteEgoPromptByIds = (params) => {
  return service({
    url: '/EP/deleteEgoPromptByIds',
    method: 'delete',
    params
  })
}

// @Tags EgoPrompt
// @Summary 更新Ego提示词记忆
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EgoPrompt true "更新Ego提示词记忆"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /EP/updateEgoPrompt [put]
export const updateEgoPrompt = (data) => {
  return service({
    url: '/EP/updateEgoPrompt',
    method: 'put',
    data
  })
}

// @Tags EgoPrompt
// @Summary 用id查询Ego提示词记忆
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.EgoPrompt true "用id查询Ego提示词记忆"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /EP/findEgoPrompt [get]
export const findEgoPrompt = (params) => {
  return service({
    url: '/EP/findEgoPrompt',
    method: 'get',
    params
  })
}

// @Tags EgoPrompt
// @Summary 分页获取Ego提示词记忆列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Ego提示词记忆列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /EP/getEgoPromptList [get]
export const getEgoPromptList = (params) => {
  return service({
    url: '/EP/getEgoPromptList',
    method: 'get',
    params
  })
}

// @Tags EgoPrompt
// @Summary 不需要鉴权的Ego提示词记忆接口
// @Accept application/json
// @Produce application/json
// @Param data query egoclientReq.EgoPromptSearch true "分页获取Ego提示词记忆列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /EP/getEgoPromptPublic [get]
export const getEgoPromptPublic = () => {
  return service({
    url: '/EP/getEgoPromptPublic',
    method: 'get',
  })
}
