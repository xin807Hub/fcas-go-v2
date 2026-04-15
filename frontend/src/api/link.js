import service from '@/utils/request'

// 链路列表
export const linkListApi = (params) => {
  return service({
    url: '/configuration/dimlineinfo/list',
    method: 'get',
    params
  })
}

export const createLinkApi = (data) => {
  return service({
    url: '/configuration/dimlineinfo/save',
    method: 'post',
    data
  })
}

export const updateLinkApi = (data) => {
  return service({
    url: '/configuration/dimlineinfo/update',
    method: 'post',
    data
  })
}

export const deleteLinkApi = (data) => {
  return service({
    url: '/configuration/dimlineinfo/delete',
    method: 'post',
    data
  })
}
