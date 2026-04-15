import service from '@/utils/request'

// 获取用户群组树
export const userCrowdGroupTreeApi = (params) => {
  return service({
    url: '/configuration/dimusercrowdgroup/getGroupTree/'+params,
    method: 'get',
  })
}

// 获取运营商、应用大类、应用大小类
export const dictInfoApi = (params) => {
  return service({
    url: '/configuration/dimDict/infoList/'+params,
    method: 'get',
  })
}

// 运营商分析
export const ispRankDataApi = (data) => {
  return service({
    url: '/traffic/isp/rankData',
    method: 'post',
    data
  })
}
export const ispRankTableDataApi = (data) => {
  return service({
    url: '/traffic/isp/rankTableData',
    method: 'post',
    data
  })
}

// 业务大类分析
export const appTypeRankDataApi = (data) => {
  return service({
    url: '/traffic/appType/rankData',
    method: 'post',
    data
  })
}
export const appTypeRankTableDataApi = (data) => {
  return service({
    url: '/traffic/appType/rankTableData',
    method: 'post',
    data
  })
}

// 业务小类分析
export const appIdRankDataApi = (data) => {
  return service({
    url: '/traffic/appId/rankData',
    method: 'post',
    data
  })
}
export const appIdRankTableDataApi = (data) => {
  return service({
    url: '/traffic/appId/rankTableData',
    method: 'post',
    data
  })
}
export const appIdRankTrendDataApi = (data) => {
  return service({
    url: '/traffic/appId/trendTableData',
    method: 'post',
    data
  })
}

// 用户排名分析
export const userRankLevel1PieApi = (data) => {
  return service({
    url: '/traffic/userRank/level1Pie',
    method: 'post',
    data
  })
}
export const userRankLevel1TableApi = (data) => {
  return service({
    url: '/traffic/userRank/level1Table',
    method: 'post',
    data
  })
}
export const userRankLevel2TrendApi = (data) => {
  return service({
    url: '/traffic/userRank/level2Trend',
    method: 'post',
    data
  })
}
export const userRankLevel2TableApi = (data) => {
  return service({
    url: '/traffic/userRank/level2Table',
    method: 'post',
    data
  })
}
export const userRankLevel3TableApi = (data) => {
  return service({
    url: '/traffic/userRank/level3Table',
    method: 'post',
    data
  })
}

// 用户群排名分析
export const userCrowdRankLevel1PieApi = (data) => {
  return service({
    url: '/traffic/userCrowdRank/level1Pie',
    method: 'post',
    data
  })
}
export const userCrowdRankLevel1TableApi = (data) => {
  return service({
    url: '/traffic/userCrowdRank/level1Table',
    method: 'post',
    data
  })
}
export const userCrowdRankLevel2TableApi = (data) => {
  return service({
    url: '/traffic/userCrowdRank/level2Table',
    method: 'post',
    data
  })
}
export const userCrowdRankLevel2TrendApi = (data) => {
  return service({
    url: '/traffic/userCrowdRank/level2Trend',
    method: 'post',
    data
  })
}
export const userCrowdRankLevel3TableApi = (data) => {
  return service({
    url: '/traffic/userCrowdRank/level3Table',
    method: 'post',
    data
  })
}

// 用户群组排名分析
export const userCrowdGroupRankLevel1PieApi = (data) => {
  return service({
    url: '/traffic/userCrowdGroupRank/level1Pie',
    method: 'post',
    data
  })
}
export const userCrowdGroupRankLevel1TableApi = (data) => {
  return service({
    url: '/traffic/userCrowdGroupRank/level1Table',
    method: 'post',
    data
  })
}
export const userCrowdGroupRankLevel2TableApi = (data) => {
  return service({
    url: '/traffic/userCrowdGroupRank/level2Table',
    method: 'post',
    data
  })
}
export const userCrowdGroupRankLevel2TrendApi = (data) => {
  return service({
    url: '/traffic/userCrowdGroupRank/level2Trend',
    method: 'post',
    data
  })
}
export const userCrowdGroupRankLevel3TableApi = (data) => {
  return service({
    url: '/traffic/userCrowdGroupRank/level3Table',
    method: 'post',
    data
  })
}
export const userCrowdGroupRankLevel3TrendApi = (data) => {
  return service({
    url: '/traffic/userCrowdGroupRank/level3Table',
    method: 'post',
    data
  })
}

// 用户行为分析
export const userActionPageDataApi = (data) => {
  return service({
    url: '/traffic/userAction/pageData',
    method: 'post',
    data
  })
}
export const userActionDetailApi = (data) => {
  return service({
    url: '/traffic/userAction/detail',
    method: 'post',
    data
  })
}

// 流量告警日志
export const alarmLogListApi = (data)=> {
  return service({
    url:'/traffic/alarmLog/list',
    method:'post',
    data
  })
}
