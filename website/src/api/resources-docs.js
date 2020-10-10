import request from '@/utils/request'

export function getDocs(query) {
  return request({
    url: '/resources/docs',
    method: 'get',
    params: query
  })
}

export function getSpecDocs(id) {
  return request({
    url: '/resources/docs/'+parseInt(id),
    method: 'get',
    //params: { id }
  })
}


export function createDocs(id,data) {
  return request({
    url: '/resources/docs',
    method: 'post',
    data
  })
}

export function updateDocs(id,data) {
  return request({
    url: '/resources/docs/'+parseInt(id),
    method: 'patch',
    data
  })
}

export function deleteDocs(id,data) {
  return request({
    url: '/resources/docs/'+parseInt(id),
    method: 'patch',
    data
  })
}
