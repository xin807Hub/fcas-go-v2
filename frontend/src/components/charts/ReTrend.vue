<template>
  <div ref="trendCon" class="w-full h-80">
    <el-empty v-if="trendData.length === 0" :image-size="60" description="暂无数据" class="mt-4" />
  </div>
</template>

<script setup>
import { ref, nextTick, onMounted, onUnmounted, markRaw } from 'vue'
// 引入echarts
import * as echarts from 'echarts'
// 引入工具类
import { formatTimeToStr } from '@/utils/date'
import { formatByte } from "@/utils/format";

const trendCon = ref(null)
const trendConInstance = ref(null)
const initTrendOption = ref({})
const trendData = ref([])

const props = defineProps({
  title: {
    type: String,
    default: ""
  },
  staticFlag: {
    type: String,
    default: ""
  },
  searchParams: {
    type: Object,
    default: {}
  },
  trendData: {
    type: Array,
    default: []
  },
  moduleName: {
    type: String,
    default: ""
  },
  moduleApi: {
    type: String,
    default: ""
  }
})
const xData = ref([])
const yData = ref([[],[]])
const titles = ref([])

// 动态引入 api 接口
const moduleName = ref(props.moduleName)
const moduleApi = ref(props.moduleApi)
const moduleMap = {
  home: () => import('@/api/dashboard'),
  traffic: () => import('@/api/traffic')
}

const initTrendChart = () => {
  initTrendOption.value = {
    backgroundColor: 'transparent',
    title: {
      text: props.title,
      textStyle: {
        fontSize: 15,
        // color: '#fff',
      },
      top: '0%',
      left: '2%',
    },
    grid: {
      top: '14%',
      left: '4%',
      right: '4%',
      bottom: '8%',
      containLabel: true,
    },
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'shadow'
      },
      backgroundColor: 'rgba(255,255,255,1)',
      borderColor: 'rgba(255,255,255,0)',
      formatter: (params) => {
        let res = '<div><p>'+ params[0].name + '</p></div>'
        res += '<p style="color:' + params[0].color + '">' + params[0].marker + params[0].seriesName + '：<span style="font-weight: bold;margin-left: 20px;">' + formatByte(params[0].value, "bps") + '</span></p>'
        res += '<p style="color:' + params[1].color + '">' + params[1].marker + params[1].seriesName + '：<span style="font-weight: bold;margin-left: 20px;">' + formatByte(params[1].value, "bps") + '</span></p>'
        return res
      }
    },
    xAxis: [{
      type: 'category',
      boundaryGap: false,
      axisLine: { //坐标轴轴线相关设置。数学上的x轴
        show: true,
        lineStyle: {
          // color: '#233e64'
        },
      },
      axisLabel: { //坐标轴刻度标签的相关设置
        margin: 15,
        // color: '#6a9cd5',
      },
      axisTick: {
        show: false,
      },
      data: xData.value,
    }],
    yAxis: [{
      type: 'value',
      // min: 0,
      // max:140,
      splitNumber: 5,
      splitLine: {
        show: true,
        lineStyle: {
          type: "dashed",
          // color: '#233e64'
        }
      },
      axisLine: {
        show: true,
      },
      axisLabel: {
        margin: 20,
        // color: '#6a9cd5',
        formatter: (value) => {
          return formatByte(value, "bps")
        }
      },
      axisTick: {
        show: false,
      },
    }],
    series: [{
      name: titles.value[0],
      type: 'line',
      smooth: true, //是否平滑曲线显示
      symbol: 'circle',  // 默认是空心圆（中间是白色的），改成实心圆
      symbolSize: 5,
      lineStyle: {
        color: "#3b82f6"   // 线条颜色
      },
      areaStyle: { //区域填充样式
        //线性渐变，前4个参数分别是x0,y0,x2,y2(范围0~1);相当于图形包围盒中的百分比。如果最后一个参数是‘true’，则该四个值是绝对像素位置。
        color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
          { offset: 0,  color: 'rgba(195,211,248,0.9)'},
          { offset: 0.9,  color: 'rgba(195,211,248,0)'}
        ], false),
        shadowColor: 'rgba(53,142,215, 0.9)', //阴影颜色
        shadowBlur: 15 //shadowBlur设图形阴影的模糊大小。配合shadowColor,shadowOffsetX/Y, 设置图形的阴影效果。
      },
      data: yData.value[0]
    },{
      name: titles.value[1],
      type: 'line',
      smooth: true, //是否平滑曲线显示
      symbol: 'circle',  // 默认是空心圆（中间是白色的），改成实心圆
      symbolSize: 5,
      lineStyle: {
        color: "#3bf680"   // 线条颜色
      },
      areaStyle: { //区域填充样式
        //线性渐变，前4个参数分别是x0,y0,x2,y2(范围0~1);相当于图形包围盒中的百分比。如果最后一个参数是‘true’，则该四个值是绝对像素位置。
        color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
          { offset: 0,  color: 'rgba(195,211,248,0.9)'},
          { offset: 0.9,  color: 'rgba(195,211,248,0)'}
        ], false),
        shadowColor: 'rgba(89,203,91,0.9)', //阴影颜色
        shadowBlur: 15 //shadowBlur设图形阴影的模糊大小。配合shadowColor,shadowOffsetX/Y, 设置图形的阴影效果。
      },
      data: yData.value[1]
    }]
  }
  setTimeout(() => {
    if (trendCon.value) {
      trendConInstance.value = markRaw(echarts.init(trendCon.value))
      trendConInstance.value.setOption(initTrendOption.value, true)
    }
  }, 100)
}
const getTrendData = async() => {
  if (moduleName.value && moduleApi.value) {
    const searchInfo = props.searchParams.form
    moduleMap[moduleName.value]().then(module => {
      const getApi = module[moduleApi.value]

      getApi(searchInfo).then(response => {
        if (response.code === 0 && response.data) {
          let data = []
          if (response.data.series) {
            data = response.data.series
          } else {
            data = response.data
          }
          xData.value = data.map(item => formatTimeToStr(item.startTime))
          yData.value[0] = data.map(item => item.upBps || item.trafficUp || item.speedUp)
          yData.value[1] = data.map(item => item.dnBps || item.trafficDn || item.speedDn)
          titles.value = ["上行", "下行"]
        }

        initTrendChart()
      })
    }).catch(err => {
      console.log("请求模块发生错误", err)
    })
  } else {
    let data = props.trendData ?? []
    xData.value = data.map(item => formatTimeToStr(item.startTime))
    yData.value[0] = data.map(item => item.upBps)
    yData.value[1] = data.map(item => item.dnBps)

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
