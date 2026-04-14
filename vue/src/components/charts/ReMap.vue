<template>
  <div id="geoMapCon" style="width: 100%;height: 550px;" />
</template>

<script setup>
import { ref, nextTick, onMounted, onUnmounted, markRaw } from 'vue'
import { ElMessage } from 'element-plus'
// 引入echarts
import * as echarts from 'echarts'
// 获取 行政区域编号
import { regionCode } from '@/api/appCommon'
// 引入 api 接口
import { getDatasetDataApi } from '@/api/dataset'
// 公共请求方法
import { getDataFormValue } from '@/utils/datasetForm'

const geoMapCon = ref(null)
const geoMapConInstance = ref(null)
const initGeoMapOption = ref({})
const geoMapData = ref([])
const geoCoordMap = ref({})
const customerBatteryCityData = ref([])
const mapCode = ref({})
const mapJson = ref({})

const getRegionCode = async() => {
  const res = await regionCode({})
  if (res.code === 0) {
    mapCode.value = res.data.Code
    // console.log(mapCode.value)
    fetch(`/maps/${mapCode.value}.json`)
      .then(response => response.json())
      .then(map => {
        mapJson.value = ReMap
        ReMap.features.forEach(item => {
          const name = item.properties.name
          const properties = item.properties.center
          geoCoordMap.value[name] = properties
          customerBatteryCityData.value.push({ name: name, value: 0 })
        })
        getGeoMapData()
      })
  } else {
    ElMessage({
      type: 'error',
      message: '获取数据失败'
    })
  }
}

const getGeoMapData = async() => {
  const params = {}
  params.dataset_id = '7b80ffeb-996f-434d-950e-0a8f88f6d568'
  const dataForm = getDataFormValue(params)
  const res = await getDatasetDataApi(dataForm)
  if (res.code === 0) {
    geoMapData.value = res.data.tableRow
  } else {
    geoMapData.value = customerBatteryCityData.value
  }
  console.log(geoMapData.value, '=========================  GeoMap')
  initGeoMapChart(geoMapData.value)
}

const initGeoMapChart = (data) => {
  data = data ?? []
  echarts.registerMap('sichuan', mapJson.value)
  data.forEach((item, index) => {
    customerBatteryCityData.value.forEach((_item, _index) => {
      if (item.city === _item.name) {
        customerBatteryCityData.value[_index]['value'] = item.bz_count + item.zz_count
      }
    })
  })
  initGeoMapOption.value = {
    backgroundColor: 'transparent',
    tooltip: {
      // borderWidth: 0,
      trigger: 'item',
      show: true,
      enterable: true,
      fontSize: 12,
      color: '#000',
      backgroundColor: 'rgba(255, 255, 255, 1)',
      formatter: (params) => {
        let res = '<div><p style="color:' + params.color + '">' + params.name + '</p></div>'
        res += '<p>' + params.marker + ' <span style="font-weight: bold;margin-left: 20px;">' + params.value + '</span></p>'
        return res
      }
      // formatter: '{b}<br />{c}'
    },
    geo: [
      {
        map: 'sichuan',
        aspectScale: 0.9,
        roam: false, // 是否允许缩放
        zoom: 1, // 默认显示级别
        scaleLimit: {
          min: 1,
          max: 2
        },
        layoutSize: '95%',
        layoutCenter: ['50%', '50%'],
        itemStyle: {
          areaColor: {
            type: 'linear-gradient',
            x: 0,
            y: 400,
            x2: 0,
            y2: 0,
            colorStops: [
              {
                offset: 0,
                color: 'rgba(37,108,190,0.3)' // 0% 处的颜色
              },
              {
                offset: 1,
                color: 'rgba(15,169,195,0.3)' // 50% 处的颜色
              }
            ],
            global: true // 缺省为 false
          },
          borderColor: '#4ecee6',
          borderWidth: 1,
          emphasis: {
            areaColor: {
              type: 'linear-gradient',
              x: 0,
              y: 300,
              x2: 0,
              y2: 0,
              colorStops: [
                {
                  offset: 0,
                  color: 'rgba(37,108,190,1)' // 0% 处的颜色
                },
                {
                  offset: 1,
                  color: 'rgba(15,169,195,1)' // 50% 处的颜色
                }
              ],
              global: true // 缺省为 false
            }
          }
        },
        emphasis: {
          itemStyle: {
            areaColor: '#0160AD'
          },
          label: {
            show: 0,
            color: '#000'
          }
        },
        zlevel: 3
      },
      {
        map: 'sichuan',
        aspectScale: 0.9,
        roam: false, // 是否允许缩放
        zoom: 1, // 默认显示级别
        scaleLimit: {
          min: 1,
          max: 2
        },
        layoutSize: '95%',
        layoutCenter: ['50%', '50%'],
        itemStyle: {
          borderColor: 'rgba(192,245,249,.6)',
          borderWidth: 2,
          shadowColor: '#2C99F6',
          shadowOffsetY: 0,
          shadowBlur: 120,
          areaColor: 'rgba(29,85,139,.2)'
        },
        zlevel: 2,
        silent: true
      },
      {
        map: 'sichuan',
        aspectScale: 0.9,
        roam: false, // 是否允许缩放
        zoom: 1, // 默认显示级别
        scaleLimit: {
          min: 1,
          max: 2
        },
        layoutSize: '95%',
        layoutCenter: ['50%', '51.5%'],
        itemStyle: {
          // areaColor: '#005DDC',
          areaColor: 'rgba(0,27,95,0.4)',
          borderColor: '#004db5',
          borderWidth: 1
        },
        zlevel: 1,
        silent: true
      }
    ],
    series: [
      {
        geoIndex: 0,
        // coordinateSystem: 'geo',
        showLegendSymbol: true,
        type: 'map',
        roam: true,
        zoom: 1,
        scaleLimit: {
          min: 1,
          max: 2
        },
        label: {
          show: false,
          color: '#000',
          emphasis: {
            show: false,
            color: '#000'
          }
        },
        itemStyle: {
          borderColor: '#2ab8ff',
          borderWidth: 1.5,
          areaColor: '#12235c',
          emphasis: {
            areaColor: '#2AB8FF',
            borderWidth: 0,
            color: 'red'
          }
        },
        map: 'sichuan', // 使用
        data: customerBatteryCityData
        // data: difficultData //热力图数据   不同区域 不同的底色
      },
      {
        type: 'lines',
        zlevel: 5,
        effect: {
          show: false,
          // period: 4, //箭头指向速度，值越小速度越快
          // trailLength: 0.02, //特效尾迹长度[0,1]值越大，尾迹越长重
          // symbol: 'arrow', //箭头图标
          // symbol: imgDatUrl,
          symbolSize: 5 // 图标大小
        },
        lineStyle: {
          width: 17, // 尾迹线条宽度
          color: {
            type: 'linear',
            x: 0,
            y: 0,
            x2: 1,
            y2: 0,
            colorStops: [
              {
                offset: 0,
                color: 'rgb(199,145,41)' // 0% 处的颜色
              },
              {
                offset: 0.5,
                color: 'rgb(199,145,41)' // 0% 处的颜色
              },
              {
                offset: 0.5,
                color: 'rgb(223,176,32)' // 0% 处的颜色
              },
              {
                offset: 1,
                color: 'rgb(223,176,32)' // 0% 处的颜色
              },
              {
                offset: 1,
                color: 'rgb(199,145,41)' // 100% 处的颜色
              }
            ],
            global: false // 缺省为 false
          },
          opacity: 1, // 尾迹线条透明度
          curveness: 0 // 尾迹线条曲直度
        },
        label: {
          show: 0,
          position: 'end',
          formatter: '245'
        },
        silent: true,
        data: lineData()
      },
      {
        type: 'scatter',
        coordinateSystem: 'geo',
        geoIndex: 0,
        zlevel: 5,
        label: {
          show: false,
          position: 'bottom',
          padding: [4, 8],
          backgroundColor: '#003F5E',
          borderRadius: 5,
          borderColor: '#67F0EF',
          borderWidth: 1,
          color: '#67F0EF'
        },
        symbol: 'diamond',
        symbolSize: [17, 8],
        color: '#ffd133',
        opacity: 1,
        silent: true,
        data: scatterData()
      },
      {
        type: 'scatter',
        coordinateSystem: 'geo',
        geoIndex: 0,
        zlevel: 4,
        label: {
          formatter: '{b}',
          position: 'bottom',
          color: '#000',
          fontSize: 12,
          distance: 10,
          show: true
        },
        symbol: 'diamond',
        symbolSize: [17, 8],
        itemStyle: {
          // color: '#F7AF21',
          color: {
            type: 'linear',
            x: 0,
            y: 0,
            x2: 1,
            y2: 0,
            colorStops: [
              {
                offset: 0,
                color: 'rgb(199,145,41)' // 0% 处的颜色
              },
              {
                offset: 0.5,
                color: 'rgb(199,145,41)' // 0% 处的颜色
              },
              {
                offset: 0.5,
                color: 'rgb(223,176,32)' // 0% 处的颜色
              },
              {
                offset: 1,
                color: 'rgb(223,176,32)' // 0% 处的颜色
              },
              {
                offset: 1,
                color: 'rgb(199,145,41)' // 100% 处的颜色
              }
            ],
            global: false // 缺省为 false
          },
          opacity: 1
        },
        silent: true,
        data: scatterData2()
      }
    ]
  }
  geoMapCon.value = document.getElementById('geoMapCon')
  geoMapConInstance.value = markRaw(echarts.init(geoMapCon.value))
  geoMapConInstance.value.setOption(initGeoMapOption.value, true)
}
// 动态计算柱形图的高度（定一个max）
const lineMaxHeight = () => {
  const maxValue = Math.max(...customerBatteryCityData.value.map(item => item.value))
  return 0.9 / maxValue
}
// 柱状体的主干
const lineData = () => {
  return customerBatteryCityData.value.map(item => {
    return {
      coords: [
        geoCoordMap.value[item.name],
        [
          geoCoordMap.value[item.name][0],
          geoCoordMap.value[item.name][1] + item.value * lineMaxHeight()
        ]
      ]
    }
  })
}
// 柱状体的顶部
const scatterData = () => {
  return customerBatteryCityData.value.map(item => {
    return [
      geoCoordMap.value[item.name][0],
      geoCoordMap.value[item.name][1] + item.value * lineMaxHeight()
    ]
  })
}
// 柱状体的底部
const scatterData2 = () => {
  return customerBatteryCityData.value.map(item => {
    return {
      name: item.name,
      value: geoCoordMap.value[item.name]
    }
  })
}

onMounted(() => {
  nextTick(() => {
    getRegionCode()
  })
  window.addEventListener('resize', () => {
    geoMapConInstance.value.resize()
  })
})

onUnmounted(() => {
  geoMapConInstance.value.dispose()
})
</script>
