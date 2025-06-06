import service from '@/utils/request'

export const login = (data) => {
  return service({
    url: '/ECU/login',
    method: 'post',
    data: data
  })
}


export const register = (data) => {
  return service({
    url: '/ECU/register',
    method: 'post',
    data: data
  })
}


export const getUserInfo = () =>{
    return service({
        url: '/ECU/getUserInfo',
        method: 'get'
    })
}


export const logout = () =>{
  return service({
    url: '/ECU/logout',
    method: 'post'
  })
}
