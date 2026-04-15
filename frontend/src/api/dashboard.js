import service from '@/utils/request'

// 总流量趋势图
export const flowDataApi = (data) => {
  return service({
    url: '/traffic/dwsTotalTraffic/home',
    method: 'post',
    data
  })
}
