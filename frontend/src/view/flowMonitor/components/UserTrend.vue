<template>
  <div class=" dark:bg-slate-900" ref="chartEl"/>
</template>

<script setup>
import {ref, watch, onMounted, onBeforeUnmount} from 'vue'
import dayjs from "dayjs";
import * as echarts from 'echarts'
import {formatSize} from "@/utils/format";
import {generateColorStyles} from "@/utils/charts";

const props = defineProps({
  title: {
    type: String,
    required: true
  },
  data: {
    type: Array,
    required: true
  },
  seriesConfig: {
    type: Object,
    required: true
  },
  xField: {
    type: String,
    required: true
  }
})

const chartEl = ref(null)
let chartInstance = null

const generateOptions = (chartData) => {
  const xData = chartData?.map(item => dayjs(item[props.xField]).format("YYYY-MM-DD HH:mm:ss"))

  const seriesConfig = props.seriesConfig

  // 生成颜色配置
  const colorSets = generateColorStyles(Object.keys(seriesConfig).length)

  const series = Object.entries(seriesConfig).map(([key, cfg], index) => {
    return {
      id: key,
      name: cfg.label,
      data: chartData?.map(d => d[key] ?? 0),
      type: 'line',
      smooth: true,
      symbol: 'circle',
      symbolSize: 6,
      lineStyle: {
        color: colorSets[index].lineColor
      },
      areaStyle: colorSets[index].areaStyle,
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
      formatter: (params) => {
        let time = params[0]?.axisValue;
        let lines = params.map(p => {
          const key = Object.entries(seriesConfig).find(([, v]) => v.label === p.seriesName)?.[0];
          const unit = seriesConfig[key]?.unit ?? 'B';
          return `${p.marker}${p.seriesName}: ${formatSize(p.data, unit)}`
        })
        return [time, ...lines].join('<br/>')
      }
    },
    legend: {
      top: '2%',
      data: Object.values(seriesConfig).map(v => v.label),
    },
    grid: {
      left: '6%',
      right: '5%',
      top: '16%',
      bottom: '12%',
    },
    xAxis: {
      type: 'category',
      data: xData,
      boundaryGap: false,
      axisTick: {
        show: false,
      },
    },
    yAxis: {
      type: 'value',
      axisLabel: {
        formatter: (val) => formatSize(val, 'B')
      }
    },
    series,
  }
}

const updateChart = async (chartData) => {
  if (!chartEl.value) return
  if (chartInstance) chartInstance.dispose()
  chartInstance = echarts.init(chartEl.value)
  chartInstance.setOption(generateOptions(chartData))
}

watch(() => props.data, (newData) => {
  console.log('watch data', newData)
  updateChart(newData)
}, {deep: true})

// 修复resize事件处理函数，使用命名函数便于移除
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

<style scoped>

</style>