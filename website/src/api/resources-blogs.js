import request from '@/utils/request'

export function getBlogs(query) {
  return request({
    url: '/resources/blogs',
    method: 'get',
    params: query
  })
}

export function getSpecBlog(id) {
  return request({
    url: '/resources/blog/' + parseInt(id),
    method: 'get'
    // params: { id }
  })
}

export function createDocs(id, data) {
  return request({
    url: '/resources/blog',
    method: 'post',
    data
  })
}

export function updateDocs(id, data) {
  return request({
    url: '/resources/blog/' + parseInt(id),
    method: 'patch',
    data
  })
}

export function deleteDocs(id, data) {
  return request({
    url: '/resources/blog/' + parseInt(id),
    method: 'delete',
    data
  })
}
