import request from '@/utils/request'

export function getSysBasicInfos(query) {
  return request({
    url: '/sys/monitor/basicinfo',
    method: 'get',
    params: query
  })
}