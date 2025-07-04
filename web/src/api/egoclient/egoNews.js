import service from '@/utils/request'
// @Tags EgoNews
// @Summary 创建Ego新闻推送
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EgoNews true "创建Ego新闻推送"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /EN/createEgoNews [post]
export const createEgoNews = (data) => {
  return service({
    url: '/EN/createEgoNews',
    method: 'post',
    data
  })
}

// @Tags EgoNews
// @Summary 删除Ego新闻推送
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EgoNews true "删除Ego新闻推送"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /EN/deleteEgoNews [delete]
export const deleteEgoNews = (params) => {
  return service({
    url: '/EN/deleteEgoNews',
    method: 'delete',
    params
  })
}

// @Tags EgoNews
// @Summary 批量删除Ego新闻推送
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Ego新闻推送"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /EN/deleteEgoNews [delete]
export const deleteEgoNewsByIds = (params) => {
  return service({
    url: '/EN/deleteEgoNewsByIds',
    method: 'delete',
    params
  })
}

// @Tags EgoNews
// @Summary 更新Ego新闻推送
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EgoNews true "更新Ego新闻推送"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /EN/updateEgoNews [put]
export const updateEgoNews = (data) => {
  return service({
    url: '/EN/updateEgoNews',
    method: 'put',
    data
  })
}

// @Tags EgoNews
// @Summary 用id查询Ego新闻推送
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.EgoNews true "用id查询Ego新闻推送"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /EN/findEgoNews [get]
export const findEgoNews = (params) => {
  return service({
    url: '/EN/findEgoNews',
    method: 'get',
    params
  })
}

// @Tags EgoNews
// @Summary 分页获取Ego新闻推送列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Ego新闻推送列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /EN/getEgoNewsList [get]
export const getEgoNewsList = (params) => {
  return service({
    url: '/EN/getEgoNewsList',
    method: 'get',
    params
  })
}

// @Tags EgoNews
// @Summary 不需要鉴权的Ego新闻推送接口
// @Accept application/json
// @Produce application/json
// @Param data query egoclientReq.EgoNewsSearch true "分页获取Ego新闻推送列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /EN/getEgoNewsPublic [get]
export const getEgoNewsPublic = () => {
  return service({
    url: '/EN/getEgoNewsPublic',
    method: 'get',
  })
}
