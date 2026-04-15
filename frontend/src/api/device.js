import service from '@/utils/request'

export const deviceListApi = (params) => {
  return service({
    url: '/configuration/dimdeviceinfo/list',
    method: 'get',
    params
  })
}

export const createDeviceApi = (data) => {
  return service({
    url: '/configuration/dimdeviceinfo/save',
    method: 'post',
    data
  })
}

export const updateDeviceApi = (data) => {
  return service({
    url: '/configuration/dimdeviceinfo/update',
    method: 'post',
    data
  })
}

export const deleteDeviceApi = (data) => {
  return service({
    url: '/configuration/dimdeviceinfo/delete',
    method: 'post',
    data
  })
}
