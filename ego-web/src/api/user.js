import service from '@/utils/request'

export const login = (data) => {
  return service({
    url: '/qmUser/login',
    method: 'post',
    data: data
  })
}


export const register = (data) => {
  return service({
    url: '/qmUser/register',
    method: 'post',
    data: data
  })
}


export const getUserInfo = () =>{
    return service({
        url: '/qmUser/getUserInfo',
        method: 'get'
    })
}


export const logout = () =>{
  return service({
    url: '/qmUser/logout',
    method: 'post'
  })
}
