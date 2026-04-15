<template>
  <div
    ref="chartEl"
    class=" dark:bg-slate-900"
  />
</template>

<script setup>
import { ref, watch, onMounted, onBeforeUnmount } from 'vue'
import * as echarts from 'echarts'
import { formatSize } from '@/utils/format'

const props = defineProps({
  title: {
    type: String,
    default: ''
  },
  data: {
    type: Array,
    required: true
  },
  isReverse: {
    // 是否反转y轴和x轴
    type: Boolean,
    default: false,
  },
  color: {
    // 柱状图颜色
    type: String,
    default: '#5470C6'
  },
  seriesName: {
    // 系列名称
    type: String,
    default: ''
  }
})

const chartEl = ref(null)
let chartInstance = null

const generateOptions = (chartData) => {
  const sortedChartData = chartData.slice().sort((a, b) => b.value - a.value);

  const series = [
    {
      name: props.seriesName,
      data: sortedChartData,
      type: 'bar',
      itemStyle: {
        color: props.color,
        barBorderRadius: props.isReverse ? [10, 0, 0, 10] : [0, 10, 10, 0],  // 四个角的圆角半径
        shadowBlur: 12,
        shadowColor: '#333',
        shadowOffsetX: 0,
        shadowOffsetY: 5,
      },
      barMinWidth: 8,  // 增加最小宽度
      label: {
        show: true,
        position: props.isReverse ? 'insideRight' : 'insideLeft',
        padding: [0, 12],  // 减少内边距
        fontSize: 12,
        borderWidth: 1,
        formatter: (params) => {
          const name = params.name.length > 24 ? params.name.slice(0, 24) + '...' : params.name;
          return `{name|${name}}: {value|${formatSize(params.value, 'B')}}`;
        },
        rich: {
          name:{
            color: '#1F2D3D',
          },
          value: {
            fontWeight: 'bold',
            color: '#1F2D3D',
            fontSize: 13,
          }
        },
      }
    },
  ]

  return {
    title: {
      text: props.title,
      top: 2,
      textStyle: {
        fontSize: 14,
      },
    },
    tooltip: {
      trigger: 'axis',
      axisPointer: { type: 'shadow' },
      formatter: (params) => {
        const param = params[0];
        return `
           ${param?.name}: ${formatSize(param.value, 'B')}
        `;
      }
    },
    grid: {
      left: '5%',
      right: '5%',
      top: '5%',
      bottom: '3%',
      containLabel: true
    },
    xAxis: {
      type: 'value',
      inverse: props.isReverse ? true : false, // 是否反转x轴
      axisLabel: {
        formatter: (val) => formatSize(val, 'B')
      }
    },
    yAxis: {
      type: 'category',
      data: sortedChartData.map(item => item.name),
      position: props.isReverse ? 'right' : 'left', // 是否反转y轴
      axisLabel: {
        show: false  // 隐藏y轴标签
      },
      axisTick: {
        show: false  // 隐藏x轴刻度线
      }
    },
    legend: {
      data: [props.seriesName],  // 使用seriesName作为legend数据
      top: '0%',
      left: 'center'
    },
    series: series
  }
}



const updateChart = (chartData) => {
  if (!chartEl.value) return
  if (chartInstance) chartInstance.dispose()
  chartInstance = echarts.init(chartEl.value)
  chartInstance.setOption(generateOptions(chartData))
}

watch(() => props.data, (newVal) => {
  if (newVal) updateChart(newVal)
}, { deep: true })

const handleResize = () => {
  if (chartInstance) chartInstance.resize();
}

onMounted(() => {
  updateChart(props.data)
  window.addEventListener('resize', handleResize)
})

onBeforeUnmount(() => {
  window.removeEventListener('resize', handleResize)
  chartInstance?.dispose()
})
</script>

<style scoped></style>
