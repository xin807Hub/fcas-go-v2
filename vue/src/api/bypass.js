import service from '@/utils/request'

// 链路列表
export const bypassListApi = (params) => {
  return service({
    url: '/configuration/dimbypass/list',
    method: 'get',
    params
  })
}

export const createBypassApi = (data) => {
  return service({
    url: '/configuration/dimbypass/save',
    method: 'post',
    data
  })
}

export const updateBypassApi = (data) => {
  return service({
    url: '/configuration/dimbypass/update',
    method: 'post',
    data
  })
}

export const deleteBypassApi = (data) => {
  return service({
    url: '/configuration/dimbypass/delete',
    method: 'post',
    data
  })
}
