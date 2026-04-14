<template>
  <div ref="trendCon" class="w-full h-128 mt-4">
    <el-empty v-if="trendData.length === 0" :image-size="60" description="暂无数据" class="mt-4" />
  </div>
</template>

<script setup>
import { ref, nextTick, onMounted, onUnmounted, markRaw } from 'vue'
// 引入echarts
import * as echarts from 'echarts'
import {formatByte} from "@/utils/format";

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
  leftData: {
    type: Array,
    default: []
  },
  rightData: {
    type: Array,
    default: []
  },
  legendData: {
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

// 动态引入 api 接口
const moduleName = ref(props.moduleName)
const moduleApi = ref(props.moduleApi)
const moduleMap = {
  traffic: () => import('@/api/traffic')
}

let data = {
  left: {
    data: props.leftData.map(item => {
      return {
        value: parseFloat(item.value).toFixed(2),
        label: [item.name].join(" ")
      }
    })
  },
  right: {
    data: props.rightData.map(item => {
      return {
        value: parseFloat(item.value).toFixed(2),
        label: [item.name].join(" ")
      }
    })
  }
};
let yAxisLabelLeft = data.left.data.map(e => e.label);
let yAxisLabelRight = data.right.data.map(e => e.label);
let yAxisDataLeft = data.left.data.map(e => e.value).sort((a, b) => a - b);
let yAxisDataRight = data.right.data.map(e => e.value).sort((a, b) => a - b);
let top = 60;
let bottom = 10;
const labelSetting = {
  label: "11",
  yLabel: "11"
};
const attackSourcesColor = [
  "#EB3B5A",
  "#FA8231",
  "#F7B731",
  "#3B84F6",
  "#1089E7",
  "#F57474",
  "#56D0E3",
  "#1089E7",
  "#F57474",
  "#1089E7",
  "#F57474",
  "#F57474"
];
const contains = (arr, dst) => {
  let i = arr.length;
  while ((i -= 1)) {
    if (arr[i] === dst) {
      return i;
    }
  }
  return false;
}


const initBarData = () => {
  initTrendOption.value = {
    backgroundColor: "transparent",
    title: {
      show: true,
      text: "目的IP和域名排行",
      x: "left",
      textStyle: {
        color: "#000",
        fontSize: 15
      }
    },
    tooltip: {
      show: true,
      trigger: "axis",
      axisPointer: {
        type: "shadow"
      },
      formatter: (params) => {
        let res = '<div><p>'+ params[0].name + '</p></div>'
        res += '<p style="color:' + params[0].color + '">' + params[0].marker + params[0].seriesName + '：<span style="font-weight: bold;margin-left: 20px;">' + formatByte(parseFloat(params[0].value), "bps") + '</span></p>'
        return res
      }
    },
    legend: {
      left: "center",
      itemGap: 100,
      top: 10,
      itemWidth: 0, //图例标记的图形宽度
      itemHeight: 0, //图例标记的图形高度
      y: "center",
      borderRadius: 0,
      inactiveColor: "#000",
      formatter: name => {
        if (name === props.legendData[0]) {
          return "{a|" + name + "}";
        } else {
          return "{b|" + name + "}";
        }
      },
      textStyle: {
        rich: {
          a: {
            padding: [8, 10, 5, 10],
            align: "center",
            fontSize: 13,
            backgroundColor: "#3b84f6",
            color: "#fff"
          },
          b: {
            padding: [8, 10, 5, 10],
            align: "center",
            fontSize: 13,
            backgroundColor: "#4AB57D",
            color: "#fff"
          }
        }
      },
      data: props.legendData
    },
    grid: [
      {
        right: "60%",
        width: "40%",
        containLabel: false,
        top: top,
        bottom: bottom
      },
      {
        left: "51%",
        width: 0,
        top: top,
        bottom: bottom
      },
      {
        left: "60%",
        width: "40%",
        containLabel: false,
        top: top,
        bottom: bottom
      }
    ],
    xAxis: [
      {
        type: "value",
        inverse: true,
        axisLine: {
          show: false
        },
        axisTick: {
          show: false
        },
        axisLabel: {
          show: false
        },
        splitLine: {
          show: false
        }
      },
      {
        gridIndex: 1,
        show: false
      },
      {
        gridIndex: 2,
        type: "value",
        axisLine: {
          show: false
        },
        axisTick: {
          show: false
        },
        axisLabel: {
          show: false
        },
        splitLine: {
          show: false
        }
      }
    ],
    yAxis: [
      {
        // show: false,
        offset: 10,
        //padding: [-23, -30],
        position: "right",
        axisLabel: {
          color: "#000",
          fontSize: labelSetting.yLabel,
        },
        axisLine: {
          show: false
        },
        type: "category",
        inverse: false,
        axisTick: {
          show: false
        },
        data: yAxisLabelLeft.reverse()
      },
      {
        gridIndex: 1,
        position: "",
        type: "category",
        inverse: true,
        axisLabel: {
          show: false,
          margin: 0,
          color: "#fff",
          fontSize: 10,
          position: "right",
          rich: {
            nt1: {
              color: "#fff",
              backgroundColor: attackSourcesColor[0],
              width: 15,
              height: 15,
              fontSize: 10,
              align: "center",
              borderRadius: 100
            },
            nt2: {
              color: "#fff",
              backgroundColor: attackSourcesColor[1],
              width: 15,
              height: 15,
              fontSize: 10,
              align: "center",
              borderRadius: 100
            },
            nt3: {
              color: "#fff",
              backgroundColor: attackSourcesColor[2],
              width: 15,
              height: 15,
              fontSize: 10,
              align: "center",
              borderRadius: 100
            },
            nt: {
              color: "#fff",
              backgroundColor: attackSourcesColor[3],
              width: 15,
              height: 15,
              fontSize: 10,
              align: "center",
              lineHeight: 3,
              borderRadius: 100
            }
          },
          formatter: function(value, index) {
            index = contains(yAxisDataRight, value) + 1;
            if (index - 1 < 3) {
              return ["{nt" + index + "|" + index + "}"].join("\n");
            } else {
              return ["{nt|" + index + "}"].join("\n");
            }
          }
        },
        axisTick: {
          show: false
        },
        axisLine: {
          show: false
        },
        axisPointer: {
          label: {
            show: true,
            margin: 30
          }
        },
        data: yAxisDataRight
      },
      {
        //show: false,
        offset: 10,
        gridIndex: 2,
        position: "left",
        axisLabel: {
          color: `#000`,
          fontSize: labelSetting.yLabel,
        },
        axisLine: {
          show: false
        },
        //type: "category",
        inverse: false,
        axisTick: {
          show: false
        },
        data: yAxisLabelRight.reverse()
      },
    ],
    series: [
      {
        name: props.legendData[0],
        type: "bar",
        barWidth: "15",
        // barMinHeight: "50",
        // lineHeight:"5",
        label: {
          show: true,
          fontSize: labelSetting.label,
          lineHeight: 56,
          distance: 10,
          color: "#fff",
          width: "15",
          align: "center",
          padding:[2,0,0,0],
          formatter: (param) => {
            return formatByte(parseFloat(param.value), "bps")
          }
        },
        itemStyle: {
          borderRadius: 30,
          color: new echarts.graphic.LinearGradient(0, 1, 1, 1, [
            { offset: 0, color: "#98b6e7" },
            { offset: 1, color: "#3b84f6" }
          ])
        },
        data: yAxisDataLeft
      },
      {
        type: "bar",
        barWidth: "15",
        // barMinHeight: "50",
        // lineHeight:"5",
        label: {
          show: true,
          fontSize: labelSetting.label,
          distance: 10,
          color: "#fff",
          padding:[2,0,0,0],
          formatter: (param) => {
            return formatByte(parseFloat(param.value), "bps")
          }
        },
        xAxisIndex: 2,
        yAxisIndex: 2,
        name: props.legendData[1],
        itemStyle: {
          borderRadius: 30,
          color: new echarts.graphic.LinearGradient(0, 1, 1, 1, [
            { offset: 0, color: "#4AB57D" },
            { offset: 1, color: "#97ecc3" },
          ])
        },
        data: yAxisDataRight
      }
    ]
  }
  setTimeout(() => {
    if (trendCon.value) {
      trendConInstance.value = markRaw(echarts.init(trendCon.value))
      trendConInstance.value.setOption(initTrendOption.value, true)
    }
  }, 100)
}
const getBarData = async() => {
  if (moduleName.value && moduleApi.value) {
    const searchInfo = props.searchParams.form
    moduleMap[moduleName.value]().then(module => {
      const getApi = module[moduleApi.value]

      getApi(searchInfo).then(response => {
        if (response.code === 0 && response.data) {
        }

        initBarData()
      })
    }).catch(err => {
      console.log("请求模块发生错误", err)
    })
  } else {
    initBarData()
  }
}

onMounted(() => {
  nextTick(() => {
    getBarData()
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
