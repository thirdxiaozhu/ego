import service from '@/utils/request'
// @Tags EgoDialogue
// @Summary 创建Ego对话
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EgoDialogue true "创建Ego对话"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ED/createEgoDialogue [post]
export const createEgoDialogue = (data) => {
  return service({
    url: '/ED/createEgoDialogue',
    method: 'post',
    data
  })
}

// @Tags EgoDialogue
// @Summary 删除Ego对话
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EgoDialogue true "删除Ego对话"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ED/deleteEgoDialogue [delete]
export const deleteEgoDialogue = (params) => {
  return service({
    url: '/ED/deleteEgoDialogue',
    method: 'delete',
    params
  })
}

// @Tags EgoDialogue
// @Summary 批量删除Ego对话
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Ego对话"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ED/deleteEgoDialogue [delete]
export const deleteEgoDialogueByIds = (params) => {
  return service({
    url: '/ED/deleteEgoDialogueByIds',
    method: 'delete',
    params
  })
}

// @Tags EgoDialogue
// @Summary 更新Ego对话
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EgoDialogue true "更新Ego对话"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ED/updateEgoDialogue [put]
export const updateEgoDialogue = (data) => {
  return service({
    url: '/ED/updateEgoDialogue',
    method: 'put',
    data
  })
}

// @Tags EgoDialogue
// @Summary 用id查询Ego对话
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.EgoDialogue true "用id查询Ego对话"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ED/findEgoDialogue [get]
export const findEgoDialogue = (params) => {
  return service({
    url: '/ED/findEgoDialogue',
    method: 'get',
    params
  })
}

// @Tags EgoDialogue
// @Summary 分页获取Ego对话列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Ego对话列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ED/getEgoDialogueList [get]
export const getEgoDialogueList = (params) => {
  return service({
    url: '/ED/getEgoDialogueList',
    method: 'get',
    params
  })
}

// @Tags EgoDialogue
// @Summary 不需要鉴权的Ego对话接口
// @Accept application/json
// @Produce application/json
// @Param data query egoclientReq.EgoDialogueSearch true "分页获取Ego对话列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ED/getEgoDialoguePublic [get]
export const getEgoDialoguePublic = () => {
  return service({
    url: '/ED/getEgoDialoguePublic',
    method: 'get',
  })
}
