<template>
  <div v-if="!height" ref="pieCon" class="w-full" style="height: 250px">
    <el-empty v-if="pieData.length === 0" :image-size="60" description="暂无数据" class="mt-4" />
  </div>
  <div v-else ref="pieCon" class="w-full" style="height: 200px">
    <el-empty v-if="pieData.length === 0" :image-size="60" description="暂无数据" class="mt-4" />
  </div>
</template>

<script setup>
import {ref, nextTick, onMounted, onUnmounted, markRaw} from 'vue'
// 引入echarts
import * as echarts from 'echarts'
// 引入工具类
import { formatByte } from "@/utils/format";
// 接收父组件传参
const props = defineProps({
  title: {
    type: String,
    default: ""
  },
  staticFlag: {
    type: String,
    default: ""
  },
  data: {
    type: Object,
    default: {}
  },
  height: {
    type: Boolean,
    default: false
  },
  searchParams: {
    type: Object,
    default: {}
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
// console.log(props.moduleName + "/" + props.moduleApi + "/" + props.staticFlag)

const pieCon = ref(null)
const pieConInstance = ref(null)
const initPieOption = ref({})
const pieData = ref([])
const seriesData = ref([])
const legendData = ref([])
const totalByte = ref(1)
// 动态引入 api 接口
const moduleName = ref(props.moduleName)
const moduleApi = ref(props.moduleApi)
const moduleMap = {
  traffic: () => import('@/api/traffic')
}
const initPieChart = () => {
  const colorList = [
    '#3b82f6',
    '#409B5C',
    '#ff5b00',
    '#ffe000',
    '#ffa800',
    '#00ffff',
    '#ff3000',
    '#f33ccc',
    '#ffdf12',
    '#4ec284',
    '#509c30',
    '#062efa',
  ]
  seriesData.value.forEach((item, index) => {
    legendData.value[index] = item.name + ":" + item.count
  })
  initPieOption.value = {
    backgroundColor: 'transparent',
    title: {
      text: props.title,
      textStyle: {
        fontSize: 15,
      },
      top: '-1%',
      left: '2%',
    },
    grid: {
      top: "10%",
      left: "8%",
      right: "4%",
      bottom: "4%"
    },
    legend: {
      show: props.height,
      icon: 'circle',
      orient: "vertical",
      top: "center",
      right: '3%',
      itemWidth: 10,
      itemHeight: 10,
      itemGap: 12,
      textStyle: {
        // color: '#000',
        fontSize: 12
      }
    },
    tooltip: {
      trigger: 'item',
      formatter: (params) => {
        let res = '<div><p style="color:' + params.color + '">' + params.marker + params.name + '</p></div>'
        res += '<p>总量：<span style="font-weight: bold;margin-left: 20px;">' + formatByte(params.value) + '</span></p>'
        res += '<p>占比：<span style="font-weight: bold;margin-left: 20px;">' + (params.value/(totalByte.value)*100).toFixed(2) + '%</span></p>'
        return res
      }
    },
    series: [{
      type: 'pie',
      center: ['50%', '50%'],
      radius: ['35%', '55%'],
      clockwise: true,
      avoidLabelOverlap: true,
      hoverOffset: 15,
      itemStyle: {
        color: function (params) {
          return colorList[params.dataIndex]
        }
      },
      label: {
        show: true,
        position: 'outside',
        color: '#909090',
        formatter: '{a|{b}：{d}%}\n{hr|}',
        rich: {
          hr: {
            backgroundColor: 't',
            borderRadius: 3,
            width: 3,
            height: 3,
            padding: [3, 3, 0, 0],
          },
          a: {
            padding: [-25, 15, -50, 15]
          }
        }
      },
      labelLine: {
        length: 20,
        length2: 30,
        lineStyle: {
          width: 1
        }
      },
      data: seriesData.value,
    }]
  }
  setTimeout(() => {
    if (pieCon.value) {
      pieConInstance.value = markRaw(echarts.init(pieCon.value))
      pieConInstance.value.setOption(initPieOption.value, true)
    }
  }, 200)
}

const getPieData = async () => {
  if (moduleName.value && moduleApi.value) {
    const searchInfo = props.searchParams.form
    moduleMap[moduleName.value]().then(module => {
      const getApi = module[moduleApi.value]

      getApi(searchInfo).then(response => {
        if (response.code === 0 && response.data) {
          seriesData.value = (response.data.pieData ? response.data.pieData : response.data).filter(item => item.name).map(item => {
            return {
              name: item.name,
              value: item[props.staticFlag]
            }
          })

          initPieChart()
        }
      })
    }).catch(err => {
      console.log("请求模块发生错误", err)
    })
  } else {
    seriesData.value = (props.data.pieData ? props.data.pieData : props.data).filter(item => item.name).map(item => {
      return {
        name: item.name,
        value: item[props.staticFlag]
      }
    })
    if (props.data.gatherData) {
      totalByte.value = props.data.gatherData.totalByte ?? 1
    } else {
      totalByte.value = props.data.reduce((arr, curr) => {
        return arr + curr.trafficTotal
      }, 0)
    }

    initPieChart()
  }
}

// 自动播放
let index = 0
const timer = ref(null)
const autoPlay = () => {
  timer.value = setInterval(() => {
    if (pieConInstance.value) {
      pieConInstance.value.dispatchAction({
        type: 'hideTip',
        seriesIndex: 0,
        dataIndex: index
      });
      // 显示提示框
      pieConInstance.value.dispatchAction({
        type: 'showTip',
        seriesIndex: 0,
        dataIndex: index
      });
      // 取消高亮指定的数据图形
      pieConInstance.value.dispatchAction({
        type: 'downplay',
        seriesIndex: 0,
        dataIndex: index === 0 ? 5 : index - 1
      });
      pieConInstance.value.dispatchAction({
        type: 'highlight',
        seriesIndex: 0,
        dataIndex: index
      });
      // 在加到5的时候，重新开始循环
      index++;
      if (index > 5) {
        index = 0;
      }
    }
  }, 3000)
}

onMounted(() => {
  nextTick(() => {
    getPieData()

    if (pieConInstance.value) {
      window.addEventListener('resize', () => {
        pieConInstance.value.resize()
      })
    }

    setTimeout(function() {
      // autoPlay()
    }, 1500);
  })
})

onUnmounted(() => {
  if (timer.value) {
    clearInterval(timer.value)
  }
  if (pieConInstance.value) {
    pieConInstance.value.dispose()
  }
})
</script>
