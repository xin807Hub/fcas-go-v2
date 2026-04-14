import service from '@/utils/request'

// 用户
export const userListApi = (params) => {
  return service({
    url: '/configuration/dimuserinfo/list',
    method: 'get',
    params
  })
}

export const createUserApi = (data) => {
  return service({
    url: '/configuration/dimuserinfo/save',
    method: 'post',
    data
  })
}

export const updateUserApi = (data) => {
  return service({
    url: '/configuration/dimuserinfo/update',
    method: 'post',
    data
  })
}

export const deleteUserApi = (data) => {
  return service({
    url: '/configuration/dimuserinfo/delete',
    method: 'post',
    data
  })
}

// 用户群
export const userCrowdApi = (params) => {
  return service({
    url: '/configuration/dimusercrowd/list',
    method: 'get',
    params
  })
}

export const createUserCrowdApi = (data) => {
  return service({
    url: '/configuration/dimusercrowd/save',
    method: 'post',
    data
  })
}

export const updateUserCrowdApi = (data) => {
  return service({
    url: '/configuration/dimusercrowd/update',
    method: 'post',
    data
  })
}

export const deleteUserCrowdApi = (data) => {
  return service({
    url: '/configuration/dimusercrowd/delete',
    method: 'post',
    data
  })
}

// 用户群组
export const userCrowdGroupApi = (params) => {
  return service({
    url: '/configuration/dimusercrowdgroup/list',
    method: 'get',
    params
  })
}

export const createUserCrowdGroupApi = (data) => {
  return service({
    url: '/configuration/dimusercrowdgroup/save',
    method: 'post',
    data
  })
}

export const updateUserCrowdGroupApi = (data) => {
  return service({
    url: '/configuration/dimusercrowdgroup/update',
    method: 'post',
    data
  })
}

export const deleteUserCrowdGroupApi = (data) => {
  return service({
    url: '/configuration/dimusercrowdgroup/delete',
    method: 'post',
    data
  })
}

// 应用分类与自定义
export const appClassifyDataApi = (params) => {
  return service({
    url: '/object/appClassify/list',
    method: 'get',
    params
  })
}
export const appClassifyImportApi = (data) => {
  return service({
    url: '/object/appClassify/import',
    method: 'post',
    data
  })
}
