<template>
  <div v-if="!height" ref="barChartRef" class="w-full h-80">
    <el-empty v-if="barData.length === 0" :image-size="60" description="暂无数据" class="mt-4" />
  </div>
  <div v-else ref="barChartRef" class="w-full h-52">
    <el-empty v-if="barData.length === 0" :image-size="60" description="暂无数据" class="mt-4" />
  </div>
</template>

<script setup>
import {ref, nextTick, onMounted, onUnmounted, markRaw} from 'vue'
import * as echarts from 'echarts'
// 引入工具类
import { formatByte } from "@/utils/format";
// 图表绘制逻辑
const barChartRef = ref(null)
const barChartRefInstance = ref(null)
const initBarOption = ref({})
const barData = ref([])
const toolTimer = ref(null)
const props = defineProps({
  title: {
    type: String,
    default: ""
  },
  staticFlag: {
    type: String,
    default: ""
  },
  height: {
    type: Boolean,
    default: false
  },
  interval: {
    type: Number,
    default: 1
  },
  line: {
    type: Boolean,
    default: true
  },
  multipleObj: {
    type: Object,
    default: {
      isMultiple: false,
      multipleName: []
    }
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
const xData = ref([])
const yData = ref([])
const yData2 = ref([])

// 动态引入 api 接口
const moduleName = ref(props.moduleName)
const moduleApi = ref(props.moduleApi)
const moduleMap = {
  traffic: () => import('@/api/traffic')
}
const getBarData = async () => {
  if (moduleName.value && moduleApi.value) {
    const searchInfo = props.searchParams.form
    moduleMap[moduleName.value]().then(module => {
      const getApi = module[moduleApi.value]

      getApi(searchInfo).then(response => {
        if (response.code === 0 && response.data) {
          const result = response.data
          resolveData(result)
          initBarChart()
        }
      })
    }).catch(err => {
      console.log("请求模块发生错误", err)
    })
  }
}

// 数据处理流程
const resolveData = (data) => {
  xData.value = data.map(item => item.srcIp)
  yData.value = data.map(item => item.trafficTotal)
}
// 初始化柱状图
const initBarChart = () => {
  initBarOption.value = {
    backgroundColor: 'transparent',
    title: {
      text: props.title,
      textStyle: {
        fontSize: 15,
        // color: '#fff',
      },
      top: '-1%',
      left: '2%',
    },
    grid: {
      top: !props.height ? "12%" : "8%",
      left: !props.height ? "12%" : (props.interval > 0 ? "8%" : "4%"),
      right: "4%",
      bottom: !props.height ? "22%" : (props.interval > 0 ? "12%" : "8%"),
    },
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'shadow'
      },
      backgroundColor: 'rgba(255,255,255,1)',
      borderColor: 'rgba(255,255,255,0)',
      formatter: function(params) {
        const item = params[0]
        const titles = props.multipleObj.multipleName
        if (props.multipleObj.multipleName.length > 0) {
          const title =  `<b class="text-blue-600">${item.name}</b>` +
              `<p style="color: ${item.color}">${item.marker}${titles[0]}：` +
              `<b class="text-slate-600">${formatByte(item.value, "bps")}</b>` +
              `</p>`
          if (!props.multipleObj.isMultiple > 0) {
            return title
          } else {
            const item2 = params[1]
            return title +
                `<p style="color: ${item2.color}">${item2.marker}${titles[1]}：`+
                `<b class="text-slate-600">${formatByte(item2.value, "bps")}</b>`+
                `</p>`
          }
        } else {
          return `<p style="color: ${item.color}">${item.marker}${item.name}：</p>` +
              `<b class="text-slate-600">${formatByte(item.value, "bps")}</b>`
        }
      }
    },
    legend: {
      show: false,
      data: [],
      top: "0%",
      textStyle: {
        // color: "#ffffff"
      }
    },
    xAxis: {
      data: xData.value,
      axisLine: {
        show: true, //隐藏X轴轴线
        lineStyle: {
          // color: '#01FCE3'
        }
      },
      axisTick: {
        show: true //隐藏X轴刻度
      },
      axisLabel: {
        show: true,
        // color: "#ebf8ac", //X轴文字颜色
        interval: props.interval,
        rotate: 30
      },
    },
    yAxis: [{
      type: "value",
      name: "",
      nameTextStyle: {
        // color: "#ebf8ac"
      },
      splitNumber: 5,
      splitLine: {
        show: false,
        lineStyle: {
          type: "dashed",
          // color: '#233e64'
        }
      },
      axisTick: {
        show: false
      },
      axisLine: {
        show: true,
        lineStyle: {
          // color: '#FFFFFF'
        }
      },
      axisLabel: {
        show: true,
        margin: 20,
        // color: "#ebf8ac"
        formatter: (value) => {
          return formatByte(value, "bps")
        }
      },
    }, {
      show: false,
      type: "value",
      name: "同比",
      nameTextStyle: {
        // color: "#ebf8ac"
      },
      position: "right",
      splitLine: {
        show: false
      },
      axisTick: {
        show: false
      },
      axisLine: {
        show: false
      },
      axisLabel: {
        show: true,
        formatter: "{value} %", //右侧Y轴文字显示
        // color: "#ebf8ac"
      }
    }, {
      type: "value",
      gridIndex: 0,
      // min: 50,
      // max: 100,
      splitNumber: 8,
      splitLine: {
        show: false
      },
      axisLine: {
        show: false
      },
      axisTick: {
        show: false
      },
      axisLabel: {
        show: false
      },
      splitArea: {
        show: true,
        areaStyle: {
          color: ["rgba(250,250,250,0.0)", "rgba(250,250,250,0.05)"]
        }
      }
    }],
    series: [{
      name: "",
      type: "line",
      yAxisIndex: 1, //使用的 y 轴的 index，在单个图表实例中存在多个 y轴的时候有用
      smooth: true, //平滑曲线显示
      showAllSymbol: false, //显示所有图形。
      symbol: "circle", //标记的图形为实心圆
      symbolSize: 5, //标记的大小
      itemStyle: { //折线拐点标志的样式
        color: "#3b82f6"
      },
      lineStyle: {
        color: "#3b82f6"
      },
      areaStyle: {
        color: "rgba(108,187,253,0.15)"
      },
      data: yData.value
    }, {
      name: "",
      type: "bar",
      barWidth: 30,
      itemStyle: {
        color: function (params) {
          const colorList = ['#3b82f6', '#3bf680', '#8036f6'];
          if (params.dataIndex < 3) {
            return colorList[params.dataIndex]
          } else {
            return "#71f1f1"
          }
        },
        // color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
        //   {offset: 0, color: "#d2e3fc"},
        //   {offset: 1, color: "#3b82f6"}
        // ], false)
      },
      data: yData.value
    }]
  }

  // 如果是多种数据集合
  if (props.multipleObj.isMultiple) {
    const extendSeries = {
      name: "",
      type: "bar",
      barWidth: 15,
      itemStyle: {
        color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
          {offset: 0, color: "#faf5dd"},
          {offset: 1, color: "#ecd049"}
        ], false)
      },
      data: yData2.value
    }
    initBarOption.value.series.push(extendSeries)
  }

  if (!props.line) {
    initBarOption.value.series.splice(0, 1)
  }

  setTimeout(() => {
    if (barChartRef.value) {
      barChartRefInstance.value = markRaw(echarts.init(barChartRef.value))
      barChartRefInstance.value.setOption(initBarOption.value, true)
    }
  }, 100)
}

// 挂载后动作 -- 初始化图表
onMounted(() => {
  nextTick(() => {
    getBarData()
  })
  // 图表自适应
  if (barChartRefInstance.value) {
    window.addEventListener('resize', () => {
      barChartRefInstance.value.resize()
    })
  }
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
