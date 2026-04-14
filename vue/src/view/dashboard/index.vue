<template>
  <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-6 py-2 gap-4 md:gap-2 gva-container2">
    <ToolBar
      ref="toolbarParam"
      :search-el="searchEl"
      :datetime-picker="datetimePicker"
      class="col-span-1 md:col-span-6 lg:col-span-6"
      @handleSearch="searchFunc"
    />
    <gva-card title="总流量趋势图" custom-class="col-span-1 md:col-span-2 md:row-start-2 lg:col-span-6 col-start-1 row-span-2" class="h-80">
      <ReTrend
        v-if="toolbarParam!==null"
        :key="loadKey"
        title="单位：Gbps"
        static-flag=""
        :search-params="toolbarParam"
        module-name="home"
        module-api="flowDataApi"
      />
    </gva-card>
    <gva-card title="链路流量分析" custom-class="col-span-1 md:col-span-6 row-span-2">
      <gva-table
        v-if="loadChildComponent"
        :key="loadKey"
        :data="tableData"
      />
    </gva-card>
  </div>
</template>

<script setup>
import {onBeforeMount, onMounted, ref} from "vue"
// 引入组件
import {GvaCard, GvaTable} from "./components"
import ReTrend from "@/components/charts/ReTrend.vue"
import ToolBar from "@/components/commonTable/toolbar.vue"
// 状态管理
import useLinkStore from "@/pinia/modules/link.js";
// 引入接口
import {flowDataApi} from "@/api/dashboard"
// 公共请求方法
const LinkStore = useLinkStore()
// 定义工具栏配置 和 执行方法
const linkOptions = ref([])
const toolbarParam = ref(null)
const datetimePicker = ref(true)
const searchEl = ref([
  {name: '链路', field: 'linkIdList', options: [], multiple: true, width: 'w-120'},
])
// 定义数据
const loadKey = ref(0)
const loadChildComponent = ref(false)
const tableData = ref([])
// 定义方法
const searchFunc = (str, form) => { // 搜索
  // console.log(str, '=======================', form)
  getChartData()
}
const getChartData = async () => {
  const param = {
    startTime: toolbarParam.value.form.startTime,
    endTime: toolbarParam.value.form.endTime,
    linkIdList: toolbarParam.value.form.linkIdList ?? []
  }
  const res = await flowDataApi(param)
  if (res.code === 0) {
    // 表格处理
    tableData.value = res.data.tables

    loadKey.value += 1
    loadChildComponent.value = true
  }
}

onBeforeMount( async () => {
  linkOptions.value = await LinkStore.getLink()
  linkOptions.value = linkOptions.value.map(item => {
    item.name = item.lineName
    item.value = item.lineVlan
    return item
  })

  searchEl.value.forEach(item => {
    if (item.name === "链路") {
      item.options = linkOptions.value
    }
  })
})

onMounted( async () => {
  await getChartData()
})
</script>

<style lang="scss" scoped>
</style>
