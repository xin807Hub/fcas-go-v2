<template>
  <div id="statisticCon" style="width: 100%;height: 300px;" />
</template>

<script setup>
import { ref, nextTick, onMounted, onUnmounted, markRaw } from 'vue'
// 引入echarts
import * as echarts from 'echarts'
// 引入 api 接口
import { getDatasetDataApi } from '@/api/dataset'
// 公共请求方法
import { getDataFormValue } from '@/utils/datasetForm'

const statisticCon = ref(null)
const statisticConInstance = ref(null)
const initStatisticOption = ref({})
const statisticData = ref([])

const props = defineProps({
  datasetId: String,
  flag: String
})
const dataset_id = props.datasetId

const xAxisData = ref(['疑似被诈骗者', '模型数量', '涉诈线索', '疑似涉诈分子'])
if (props.flag === 'app') {
  xAxisData.value = ['点击率', '涉诈APP数量']
}
const barData = ref([])
const getStatisticData = async() => {
  const params = {}
  params.dataset_id = dataset_id
  const dataForm = getDataFormValue(params)
  const res = await getDatasetDataApi(dataForm)
  statisticData.value = res.data.tableRow
  initStatisticChart(statisticData.value)
  // console.log(statisticData.value, ' ========  Statistic ======== ', props.flag)
}

const initStatisticChart = (data) => {
  if (data) {
    barData.value = Object.values(data[0])
  }
  const colors = []
  for (let i = 0; i < 4; i++) {
    colors.push({
      type: 'linear',
      x: 0,
      y: 0,
      x2: 1,
      y2: 0,
      colorStops: [
        { offset: 0, color: '#6fcff5' },
        { offset: 0.5, color: '#76cbee' },
        { offset: 0.5, color: '#90d0ea' },
        { offset: 1, color: '#2abaf3' }
      ]
    })
  }
  initStatisticOption.value = {
    backgroundColor: 'transparent',
    tooltip: {
      trigger: 'axis',
      formatter: '{b} : {c}',
      axisPointer: {
        type: 'shadow'
      }
    },
    grid: {
      top: '20%',
      bottom: '20%',
      left: '2%',
      right: '2%'
    },
    xAxis: {
      data: xAxisData.value,
      axisLine: {
        show: false
      },
      axisTick: {
        show: false
      },
      axisLabel: {
        margin: 25,
        color: '#000',
        fontSize: 12
      }
    },
    yAxis: {
      show: false
    },
    series: [
      {
        type: 'bar',
        barWidth: 60,
        itemStyle: {
          color: function(params) {
            return colors[params.dataIndex % 7]
          }
        },
        label: {
          show: true,
          color: '#000',
          fontSize: 12,
          position: 'insideTop',
          offset: [0, -35]
        },
        data: barData
      },
      {
        z: 3,
        type: 'pictorialBar',
        data: barData,
        symbol: 'diamond',
        symbolOffset: [0, '50%'],
        symbolSize: [60, 60 * 0.5],
        itemStyle: {
          color: function(params) {
            return colors[params.dataIndex % 7]
          }
        }
      },
      {
        z: 4,
        type: 'pictorialBar',
        data: barData,
        symbol: 'diamond',
        symbolPosition: 'end',
        symbolOffset: [0, '-50%'],
        symbolSize: [60, 60 * 0.5],
        itemStyle: {
          borderWidth: 0,
          color: '#50c6f5'
        }
      }
    ]
  }
  statisticCon.value = document.getElementById('statisticCon')
  statisticConInstance.value = markRaw(echarts.init(statisticCon.value))
  statisticConInstance.value.setOption(initStatisticOption.value, true)
}

onMounted(() => {
  nextTick(() => {
    getStatisticData()
  })
  window.addEventListener('resize', () => {
    statisticConInstance.value.resize()
  })
})

onUnmounted(() => {
  statisticConInstance.value.dispose()
})
</script>
