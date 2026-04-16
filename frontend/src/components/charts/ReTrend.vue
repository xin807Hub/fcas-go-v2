<template>
  <div ref="trendCon" class="w-full h-80">
    <el-empty v-if="trendData.length === 0" :image-size="60" description="暂无数据" class="mt-4" />
  </div>
</template>

<script setup>
import { ref, nextTick, onMounted, onUnmounted, markRaw } from 'vue'
import * as echarts from 'echarts'
import { formatTimeToStr } from '@/utils/date'
import { formatByte } from '@/utils/format'

const trendCon = ref(null)
const trendConInstance = ref(null)
const initTrendOption = ref({})
const trendData = ref([])

const props = defineProps({
  title: {
    type: String,
    default: ''
  },
  staticFlag: {
    type: String,
    default: ''
  },
  searchParams: {
    type: Object,
    default: () => ({})
  },
  trendData: {
    type: Array,
    default: () => []
  },
  moduleName: {
    type: String,
    default: ''
  },
  moduleApi: {
    type: String,
    default: ''
  }
})

const xData = ref([])
const yData = ref([[], [], []])
const titles = ref(['上行流速', '下行流速', '总流速'])
const lineColors = ['#3b82f6', '#3bf680', '#f59e0b']
const lineShadowColors = ['rgba(53,142,215,0.9)', 'rgba(89,203,91,0.9)', 'rgba(245,158,11,0.9)']

const moduleName = ref(props.moduleName)
const moduleApi = ref(props.moduleApi)
const moduleMap = {
  home: () => import('@/api/dashboard'),
  traffic: () => import('@/api/traffic')
}

const getPointValue = (item, keys) => {
  for (const key of keys) {
    const value = item?.[key]
    if (value !== undefined && value !== null) {
      return Number(value) || 0
    }
  }
  return 0
}

const buildSeries = () => yData.value.map((seriesData, index) => ({
  name: titles.value[index],
  type: 'line',
  smooth: true,
  symbol: 'circle',
  symbolSize: 5,
  lineStyle: {
    color: lineColors[index]
  },
  areaStyle: {
    color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
      { offset: 0, color: 'rgba(195,211,248,0.9)' },
      { offset: 0.9, color: 'rgba(195,211,248,0)' }
    ], false),
    shadowColor: lineShadowColors[index],
    shadowBlur: 15
  },
  data: seriesData
}))

const initTrendChart = () => {
  initTrendOption.value = {
    backgroundColor: 'transparent',
    title: {
      text: props.title,
      textStyle: {
        fontSize: 15
      },
      top: '0%',
      left: '2%'
    },
    legend: {
      top: '0%',
      right: '4%',
      icon: 'roundRect',
      itemWidth: 12,
      itemHeight: 8,
      data: titles.value
    },
    grid: {
      top: '18%',
      left: '4%',
      right: '4%',
      bottom: '8%',
      containLabel: true
    },
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'shadow'
      },
      backgroundColor: 'rgba(255,255,255,1)',
      borderColor: 'rgba(255,255,255,0)',
      formatter: (params) => {
        let res = `<div><p>${params[0]?.name || ''}</p></div>`
        params.forEach((item) => {
          res += `<p style="color:${item.color}">${item.marker}${item.seriesName}：<span style="font-weight: bold;margin-left: 20px;">${formatByte(item.value, 'bps')}</span></p>`
        })
        return res
      }
    },
    xAxis: [{
      type: 'category',
      boundaryGap: false,
      axisLine: {
        show: true
      },
      axisLabel: {
        margin: 15
      },
      axisTick: {
        show: false
      },
      data: xData.value
    }],
    yAxis: [{
      type: 'value',
      splitNumber: 5,
      splitLine: {
        show: true,
        lineStyle: {
          type: 'dashed'
        }
      },
      axisLine: {
        show: true
      },
      axisLabel: {
        margin: 20,
        formatter: (value) => formatByte(value, 'bps')
      },
      axisTick: {
        show: false
      }
    }],
    series: buildSeries()
  }

  setTimeout(() => {
    if (trendCon.value) {
      trendConInstance.value = markRaw(echarts.init(trendCon.value))
      trendConInstance.value.setOption(initTrendOption.value, true)
    }
  }, 100)
}

const fillTrendData = (data) => {
  xData.value = data.map((item) => formatTimeToStr(item.startTime))
  yData.value[0] = data.map((item) => getPointValue(item, ['upBps', 'trafficUp', 'speedUp']))
  yData.value[1] = data.map((item) => getPointValue(item, ['dnBps', 'trafficDn', 'speedDn']))
  yData.value[2] = data.map((item) => {
    const up = getPointValue(item, ['upBps', 'trafficUp', 'speedUp'])
    const down = getPointValue(item, ['dnBps', 'trafficDn', 'speedDn'])
    return up + down
  })
}

const getTrendData = async () => {
  if (moduleName.value && moduleApi.value) {
    const searchInfo = props.searchParams.form
    moduleMap[moduleName.value]().then((module) => {
      const getApi = module[moduleApi.value]

      getApi(searchInfo).then((response) => {
        if (response.code === 0 && response.data) {
          const data = response.data.series ? response.data.series : response.data
          fillTrendData(data)
        }

        initTrendChart()
      })
    }).catch((err) => {
      console.log('请求模块发生错误', err)
    })
  } else {
    fillTrendData(props.trendData ?? [])
    initTrendChart()
  }
}

onMounted(() => {
  nextTick(() => {
    getTrendData()
  })
  if (trendConInstance.value) {
    window.addEventListener('resize', () => {
      trendConInstance.value.resize()
    })
  }
})

onUnmounted(() => {
  if (trendConInstance.value) {
    trendConInstance.value.dispose()
  }
})
</script>
