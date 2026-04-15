<template>
  <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-6 gap-4 md:gap-2 gva-container2">
    <ToolBar ref="toolbarParam" :search-el="searchEl" :datetime-picker="datetimePicker"
      class="col-span-1 md:col-span-6 lg:col-span-6" @handleSearch="searchFunc" />
    <gva-card title="总流量趋势图" custom-class="col-span-1 md:col-span-2 md:row-start-2 lg:col-span-6 col-start-1 row-span-2"
      class="h-80">
      <ReTrend v-if="toolbarParam !== null" :key="loadKey" v-loading="loading" title="单位：Gbps" static-flag=""
        :search-params="toolbarParam" module-name="home" module-api="flowDataApi" />
    </gva-card>
    <gva-card title="链路流量分析" custom-class="col-span-1 md:col-span-6 row-span-2">
      <!-- <gva-table v-if="loading" :key="loadKey" :data="tableData" /> -->

      <!--表格-->
      <el-table v-loading="loading" :data="tableData" height="380">
        <el-table-column label="链路名称" prop="linkName" sortable align="center" show-overflow-tooltip />
        <el-table-column label="上行峰值" prop="maxUpBps" sortable
          :formatter="(row, column, cellValue) => formatSize(cellValue, { units: ['bps', 'Kbps', 'Mbps', 'Gbps'] })" />
        <el-table-column label="下行峰值" prop="maxDnBps" sortable
          :formatter="(row, column, cellValue) => formatSize(cellValue, { units: ['bps', 'Kbps', 'Mbps', 'Gbps'] })" />
        <el-table-column label="上行平均" prop="avgUpBps" sortable
          :formatter="(row, column, cellValue) => formatSize(cellValue, { units: ['bps', 'Kbps', 'Mbps', 'Gbps'] })" />
        <el-table-column label="下行平均" prop="avgDnBps" sortable
          :formatter="(row, column, cellValue) => formatSize(cellValue, { units: ['bps', 'Kbps', 'Mbps', 'Gbps'] })" />
      </el-table>
    </gva-card>
  </div>
</template>

<script setup>
import { onBeforeMount, onMounted, ref } from "vue"
import { GvaCard } from "./components"
import ReTrend from "@/components/charts/ReTrend.vue"
import ToolBar from "@/components/commonTable/toolbar.vue"
import useLinkStore from "@/pinia/modules/link.js";
import { flowDataApi } from "@/api/dashboard"
import { formatSize } from "@/utils/format";

// 公共请求方法
const LinkStore = useLinkStore()
// 定义工具栏配置 和 执行方法
const linkOptions = ref([])
const toolbarParam = ref(null)
const datetimePicker = ref(true)
const searchEl = ref([
  { name: '链路', field: 'linkIdList', options: [], multiple: true, width: 'w-120' },
])
// 定义数据
const loadKey = ref(0)
const loading = ref(false)
const tableData = ref([])
// 定义方法
const searchFunc = (str, form) => { // 搜索
  getChartData()
}
const getChartData = async () => {
  loading.value = true
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
    loading.value = false
  }
}

onBeforeMount(async () => {
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

onMounted(async () => {
  await getChartData()
})
</script>

<style lang="scss" scoped></style>
