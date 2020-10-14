import request from '@/utils/request'

export function getNodes(query) {
  return request({
    url: '/cloud/resources/nodes',
    method: 'get',
    params: query
  })
}
export function getNamespaces(query) {
  return request({
    url: '/cloud/resources/namespaces',
    method: 'get',
    params: query
  })
}

export function getService(query) {
  return request({
    url: '/cloud/resources/service',
    method: 'get',
    params: query
  })
}

export function getConfigmap(query) {
  return request({
    url: '/cloud/resources/configmap',
    method: 'get',
    params: query
  })
}

export function getServiceAccount(query) {
  return request({
    url: '/cloud/resources/serviceaccount',
    method: 'get',
    params: query
  })
}

export function getPod(query) {
  return request({
    url: '/cloud/resources/pod',
    method: 'get',
    params: query
  })
}

export function getJob(query) {
  return request({
    url: '/cloud/resources/job',
    method: 'get',
    params: query
  })
}
export function getBatchJob(query) {
  return request({
    url: '/cloud/resources/batchjob',
    method: 'get',
    params: query
  })
}
export function getDeployment(query) {
  return request({
    url: '/cloud/resources/deployment',
    method: 'get',
    params: query
  })
}
export function getDaemonset(query) {
  return request({
    url: '/cloud/resources/daemonset',
    method: 'get',
    params: query
  })
}

export function getReplicaset(query) {
  return request({
    url: '/cloud/resources/replicaset',
    method: 'get',
    params: query
  })
}


export function getStatefulSet(query) {
  return request({
    url: '/cloud/resources/statefulset',
    method: 'get',
    params: query
  })
}

export function getPv(query) {
  return request({
    url: '/cloud/resources/pv',
    method: 'get',
    params: query
  })
}

export function getPvc(query) {
  return request({
    url: '/cloud/resources/pvc',
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
