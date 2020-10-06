import request from '@/utils/request'

export function getResourceArticles() {
  return request({
    url: '/plat-resources/resources/articles',
    method: 'get'
  })
}

export function addRole(data) {
  return request({
    url: '/auth-rbac/role',
    method: 'post',
    data
  })
}

export function updateRole(id, data) {
  return request({
    url: `/auth-rbac/role/${id}`,
    method: 'put',
    data
  })
}

export function deleteRole(id) {
  return request({
    url: `/auth-rbac/role/${id}`,
    method: 'delete'
  })
}

export function getPermissions() {
  return request({
    url: '/auth-rbac/permissions',
    method: 'get'
  })
}

export function addPermissions(data) {
  return request({
    url: '/auth-rbac/permission',
    method: 'post',
    data
  })
}

export function updatePermissions(id, data) {
  return request({
    url: `/auth-rbac/permission/${id}`,
    method: 'put',
    data
  })
}

export function deletePermissions(id) {
  return request({
    url: `/auth-rbac/permission/${id}`,
    method: 'delete'
  })
}
