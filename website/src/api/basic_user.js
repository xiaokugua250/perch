import request from '@/utils/request'

export function login(data) {
  return request({
    // url: '/vue-element-admin/user/login',
    url: '/basic/user/login',
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    // url: '/vue-element-admin/user/info',
    url: '/basic/user/info',
    method: 'get',
    params: { token }
  })
}

export function logout() {
  return request({
    url: '/basic/user/logout',
    method: 'post'
  })
}

// -------------------用户管理相关API

export function authuserCreate(data) {
  return request({
    url: '/basic/auth-user/user',
    method: 'post',
    data
  })
}
export function authuserDelete(id) {
  return request({
    url: '/basic/auth-user/user/' + parseInt(id),
    method: 'delete'
  })
}
export function authuserUpdate(id, data) {
  return request({
    url: '/basic/auth-user/user/' + parseInt(id),
    method: 'patch',
    data
  })
}
export function authuserGet(query) {
  return request({
    url: '/basic/auth-user/users',
    method: 'get',
    params: query
  })
}

