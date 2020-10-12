import request from '@/utils/request'

export function getCategory(query) {
  return request({
    url: '/resources/categorys',
    method: 'get',
    params: query
  })
}

export function getSpecCategory(id) {
  return request({
    url: '/resources/category/'+parseInt(id),
    method: 'get',
    //params: { id }
  })
}


export function createCategory(id,data) {
  return request({
    url: '/resources/category',
    method: 'post',
    data
  })
}

export function updateCategory(id,data) {
  return request({
    url: '/resources/category/'+parseInt(id),
    method: 'patch',
    data
  })
}

export function deleteCategory(id,data) {
  return request({
    url: '/resources/category/'+parseInt(id),
    method: 'patch',
    data
  })
}
