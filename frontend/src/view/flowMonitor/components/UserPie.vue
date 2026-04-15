<template>
  <div class=" dark:bg-slate-900" ref="chartEl"/>
</template>

<script setup>
import {computed, watch, onMounted, ref, onBeforeUnmount} from 'vue'
import {formatByte} from "@/utils/format";
import * as echarts from "echarts";


const props = defineProps({
  title: {
    type: String,
    required: true
  },
  // 数据格式：[{name: 'xxx', value: 123}, {name: 'yyy', value: 456}]
  data: {
    type: Array,
    required: true
  },
})

const chartEl = ref(null)
let chartInstance = null

const totalByte = computed(() => {
  return props.data?.reduce((acc, cur) => acc + cur.value, 0)
})

const generateOptions = (chartData) => {
  return {
    title: {
      text: props.title,
      left: 10,
      top: 10,
    },
    tooltip: {
      trigger: 'item',
      formatter: (params) => {
        let res = '<div><p style="color:' + params.color + '">' + params.marker + params.name + '</p></div>'
        res += '<p>总量：<span style="font-weight: bold;margin-left: 20px;">' + formatByte(params.value) + '</span></p>'
        res += '<p>占比：<span style="font-weight: bold;margin-left: 20px;">' + (params.value / (totalByte.value) * 100).toFixed(2) + '%</span></p>'
        return res
      }
    },
    legend: {
      show: false,
      orient: 'horizontal',
    },
    series: [
      {
        name: props.title,
        type: 'pie',
        radius: ['43%', '72%'],
        center: ['50%', '52%'], // 调整饼图中心位置，向上移动一些
        avoidLabelOverlap: true,
        itemStyle: {
          borderRadius: 6,
          borderColor: '#fff',
          borderWidth: 1
        },
        label: {
          show: true,
          formatter: '{b}: {d}%',
          color: '#909090',
        },
        emphasis: {
          label: {
            show: true,
            fontSize: 18,
            fontWeight: 'bold'
          },
          itemStyle: {
            shadowBlur: 10,
            shadowOffsetX: 0,
            shadowColor: 'rgba(0, 0, 0, 0.5)'
          }
        },
        labelLine: {
          show: true,
          length: 10,
        },
        data: chartData
      }
    ],
    color: [
      '#409EFF',
      '#67C23A',
      '#E6A23C',
      '#F56C6C',
      '#909399',
      '#36CBCB',
      '#FFA2D3'
    ]
  }
}


const updateChart = async (chartData) => {
  if (!chartEl.value) return
  if (chartInstance) chartInstance.dispose()
  chartInstance = echarts.init(chartEl.value)
  chartInstance.setOption(generateOptions(chartData))
}

watch(() => props.data, (newData) => {
  updateChart(newData)
}, {deep: true})

// 修复resize事件处理函数，使用命名函数便于移除
const handleResize = () => {
  if (chartInstance) chartInstance.resize();
}

onMounted(() => {
  console.log('mounted', props.data)
  updateChart(props.data)
  window.addEventListener('resize', handleResize)
})

onBeforeUnmount(() => {
  window.removeEventListener('resize', handleResize)
  chartInstance?.dispose()
})

</script>
