<template>
  <div ref="barChartRef" style="width: 100%;height: 320px;" />
</template>

<script setup>
import { ref, nextTick, onMounted, onUnmounted, markRaw } from 'vue'
import * as echarts from 'echarts'

// 图表绘制逻辑
const barChartRef = ref(null)
const barChartRefInstance = ref(null)
const initBarOption = ref({})
const toolTimer = ref(null)

const props = defineProps({
  title: String,
  legendData: Object,
  xData: Object,
  data: Object,
  config: {
    stack: {
      type: String,
      default: null
    },
    barWidth: {
      type: Number,
      default: 15
    }
  }
})
const title = ref(props.title)
const legendData = ref(props.legendData)
const xData = ref(props.xData)
const data = ref(props.data)
const stack = props.config ? props.config.stack : ""
const barWidth = props.config ? props.config.barWidth : 15

const color = [
  ["#2b89f8", "#529df6"],
  ["#07ee4e", "#57f186"],
  ["#ff5b00", "#ef864d"],
  ["#f6ea7e", "#f5ec9d"]
]
const initBarChart = () => {
  const seriesData = []
  let interval = 0
  if (data.value) {
    data.value.forEach((item, index) => {
      const series = {
        name: legendData.value[index],
        type: 'bar',
        stack: stack,
        smooth: true,
        barWidth: barWidth,
        itemStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            { offset: 0, color: color[index][0] },
            { offset: 1, color: color[index][1] }
          ])
        },
        data: item ?? []
      }
      seriesData[index] = series
    })
    if (data.value.length > 8) {
      interval = 1
    }
  }
  initBarOption.value = {
    title: [
      {
        text: title.value,
        textStyle: {
          fontSize: 13
        }
      }
    ],
    grid: {
      top: '10%',
      bottom: '13%',
      left: '7%',
      right: '2%',
    },
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'shadow'
      },
      formatter: function(params) {
        const item = params[0]
        return `<p style="color: ${item.color}">${item.name}:${item.value}</p>`
      }
    },
    legend: {
      show: true,
      data: legendData.value,
      icon: 'rect',
      orient: "horizontal",
      right: 'center',
      bottom: '0%',
      itemWidth: 12,
      itemHeight: 12,
      itemGap: 15,
      textStyle: {
        fontSize: 12
      }
    },
    xAxis: {
      type: 'category',
      data: xData.value,
      axisLine: {
        show: true,
        lineStyle: {
        }
      },
      axisTick: {
        show: false
      },
      axisLabel: {
        interval: interval,
        rotate: '0'
      }
    },
    yAxis: {
      type: 'value',
      axisLine: {
        show: false,
        lineStyle: {
        }
      },
      splitLine: {
        show: true,
        lineStyle: {
          type: 'dashed',
        }
      },
      axisTick: {
        show: false
      },
      axisLabel: {
        fontSize: 12
      }
      // boundaryGap: ['20%', '20%'],
    },
    // dataZoom: {
    //   start: 0,
    //   end: 100,
    //   type: 'inside'
    // },
    series: seriesData
  }
  if (barChartRef.value) {
    barChartRefInstance.value = markRaw(echarts.init(barChartRef.value))
    barChartRefInstance.value.setOption(initBarOption.value, true)
  }
}

// tooltip自动轮播 若使用请做销毁处理
// if (initBarOption.value && typeof initBarOption.value === 'object') {
//   let index = 0
//   toolTimer.value = setInterval(function() {
//     barChartRefInstance.value.dispatchAction({
//       type: 'showTip',
//       seriesIndex: 0,
//       dataIndex: index
//     })
//     index++
//     if (index >= 10) {
//       index = 0
//     }
//   }, 3000)
// }

// 挂载后动作 -- 初始化图表
onMounted(() => {
  nextTick(() => {
    initBarChart()
  })
  // 图表自适应
  window.addEventListener('resize', () => {
    if (barChartRefInstance.value) {
      barChartRefInstance.value.resize()
    }
  })
})

// 销毁前动作 -- 销毁图表
onUnmounted(() => {
  if (barChartRefInstance.value) {
    barChartRefInstance.value.dispose()
  }
  if (toolTimer.value) {
    clearInterval(toolTimer.value)
  }
})
</script>

<style scoped lang="scss"></style>
