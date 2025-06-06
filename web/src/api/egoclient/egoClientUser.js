import service from '@/utils/request'
// @Tags EgoClientUser
// @Summary 创建EGO用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EgoClientUser true "创建EGO用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ECU/createEgoClientUser [post]
export const createEgoClientUser = (data) => {
  return service({
    url: '/ECU/createEgoClientUser',
    method: 'post',
    data
  })
}

// @Tags EgoClientUser
// @Summary 删除EGO用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EgoClientUser true "删除EGO用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ECU/deleteEgoClientUser [delete]
export const deleteEgoClientUser = (params) => {
  return service({
    url: '/ECU/deleteEgoClientUser',
    method: 'delete',
    params
  })
}

// @Tags EgoClientUser
// @Summary 批量删除EGO用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EGO用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ECU/deleteEgoClientUser [delete]
export const deleteEgoClientUserByIds = (params) => {
  return service({
    url: '/ECU/deleteEgoClientUserByIds',
    method: 'delete',
    params
  })
}

// @Tags EgoClientUser
// @Summary 更新EGO用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EgoClientUser true "更新EGO用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ECU/updateEgoClientUser [put]
export const updateEgoClientUser = (data) => {
  return service({
    url: '/ECU/updateEgoClientUser',
    method: 'put',
    data
  })
}

// @Tags EgoClientUser
// @Summary 用id查询EGO用户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.EgoClientUser true "用id查询EGO用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ECU/findEgoClientUser [get]
export const findEgoClientUser = (params) => {
  return service({
    url: '/ECU/findEgoClientUser',
    method: 'get',
    params
  })
}

// @Tags EgoClientUser
// @Summary 分页获取EGO用户列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EGO用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ECU/getEgoClientUserList [get]
export const getEgoClientUserList = (params) => {
  return service({
    url: '/ECU/getEgoClientUserList',
    method: 'get',
    params
  })
}

// @Tags EgoClientUser
// @Summary 不需要鉴权的EGO用户接口
// @Accept application/json
// @Produce application/json
// @Param data query egoclientReq.EgoClientUserSearch true "分页获取EGO用户列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ECU/getEgoClientUserPublic [get]
export const getEgoClientUserPublic = () => {
  return service({
    url: '/ECU/getEgoClientUserPublic',
    method: 'get',
  })
}
// AdminChangePassword 管理员修改密码
// @Tags EgoClientUser
// @Summary 管理员修改密码
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /ECU/adminChangePassword [PUT]
export const adminChangePassword = (data) => {
  console.log(data)
  return service({
    url: '/ECU/adminChangePassword',
    method: 'PUT',
    data
  })
}
// Register 用户注册
// @Tags EgoClientUser
// @Summary 用户注册
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /ECU/register [POST]
export const register = () => {
  return service({
    url: '/ECU/register',
    method: 'POST'
  })
}
// Login 用户登录
// @Tags EgoClientUser
// @Summary 用户登录
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /ECU/login [POST]
export const login = () => {
  return service({
    url: '/ECU/login',
    method: 'POST'
  })
}
// GetUserInfo 获取用户信息
// @Tags EgoClientUser
// @Summary 获取用户信息
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /ECU/getUserInfo [GET]
export const getUserInfo = () => {
  return service({
    url: '/ECU/getUserInfo',
    method: 'GET'
  })
}
