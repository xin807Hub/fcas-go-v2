<template>
  <div class=" dark:bg-slate-900" ref="chartEl" />
</template>

<script setup>
import { ref, watch, onMounted, onBeforeUnmount } from 'vue'
import * as echarts from 'echarts'
import { formatSize } from '@/utils/format'
import { generateColorStyles } from '@/utils/charts'

const props = defineProps({
  title: {
    type: String,
    default: ''
  },
  data: {
    type: Array,
    required: true
  }
})

const chartEl = ref(null)
let chartInstance = null

const generateOptions = (chartData) => {
  const seriesConfig = {
    trafficUp: { label: '上行流量' },
    trafficDn: { label: '下行流量' },
  }

  const colorSets = generateColorStyles(Object.keys(seriesConfig).length)

  const series = Object.entries(seriesConfig).map(([key, cfg], index) => {
    return {
      id: key,
      name: cfg.label,
      data: chartData?.map(d => d[key] ?? 0),
      type: 'bar',
      stack: 'total',
      itemStyle: {
        color: colorSets[index].lineColor
      },
      areaStyle: colorSets[index].areaStyle,
      label: {
        show: false
      },
    }
  })

  return {
    title: {
      text: props.title,
      top: 8,
      textStyle: {
        fontSize: 14,
      },
    },
    tooltip: {
      trigger: 'axis',
      axisPointer: { type: 'shadow' },
      formatter: (params) => {
        const up = params.find(p => p.seriesName === '上行流量')
        const dn = params.find(p => p.seriesName === '下行流量')
        const total = (up?.value ?? 0) + (dn?.value ?? 0)
        return `
          ${up?.name}<br/>
          ${up.marker} 上行: ${formatSize(up.value, 'B')}<br/>
          ${dn.marker} 下行: ${formatSize(dn.value, 'B')}<br/>
          总流量: ${formatSize(total, 'B')}
        `
      }
    },
    grid: {
      left: '1%',
      right: '3%',
      top: '8%',
      bottom: '3%',
      containLabel: true
    },
    xAxis: {
      type: 'value',
      axisLabel: {
        formatter: (val) => formatSize(val, 'B')
      }
    },
    yAxis: {
      type: 'category',
      data: chartData.map(item => item.srcIp),
      axisTick: {
        show: false
      }
    },
    legend: {
      top: '2%',
      data: Object.values(seriesConfig).map(v => v.label)
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
