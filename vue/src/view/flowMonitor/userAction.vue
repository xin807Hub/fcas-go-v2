<template>
  <div>
    <!--    工具栏    -->
    <ToolBar
      ref="toolbarParam"
      :search-el="searchEl"
      :datetime-picker="datetimePicker"
      toolbar-type="excute"
      @handleSearch="searchFunc"
      @handleRefresh="refreshFunc"
      @handleExport="exportFunc"
    />
    <!--    数据列表    -->
    <CommonTable
      v-if="toolbarParam!==null"
      :key="loadKey"
      ref="tableParam"
      :table-params="toolbarParam.form"
      :table-columns="tableColumns"
      :operation="false"
      :expand="expand"
      :auto-search="autoSearch"
      :page-size="50"
      module-name="traffic"
      module-api="userActionDataApi"
    />
  </div>
</template>

<script setup>
import {ref, onMounted, nextTick} from "vue"
import axios from "axios"
import { saveAs } from "file-saver"
// 引入公共组件
import ToolBar from "@/components/commonTable/toolbar.vue"
import CommonTable from "@/components/commonTable/index.vue"
// 定义工具栏配置
const toolbarParam = ref(null)
const datetimePicker = ref(true)
const dataTypeOptions = [
  { name: "目的IP", value: "dst_ip" },
  { name: "应用小类", value: "app_id" }
]
const searchEl = ref([
  { name: "源IP", field: "srcIp", width: 'w-72' },
  { name: "数据类型", field: "dataType", options: dataTypeOptions, multiple: false, width: 'w-56' },
])
// 列表查询逻辑
const loadKey = ref(0)
const expand = ref("dstIp")
const autoSearch = ref(false)
const tableParam = ref(null)
const tableColumns = ref([
  { name: "目的IP", field: "name" },
  { name: "上行流量", field: "upByte", sortable: true },
  { name: "下行流量", field: "dnByte", sortable: true },
  { name: "总流量", field: "totalByte", sortable: true },
])

// 查询
const searchFunc = () => {
  autoSearch.value = true
  loadKey.value += 1
  tableColumns.value[0].name = "目的IP"
  if (toolbarParam.value.form.dataType === 'app_id') {
    tableColumns.value[0].name = "应用名称"
  }
}
// 刷新
const refreshFunc = () => {
  loadKey.value += 1
}
// 导出
const exportFunc = async () => {
  axios({
    method: "post",
    url: '/api/traffic/userAction/export',
    data: toolbarParam.value.form,
    responseType: 'blob',
  })
      .then((response) => {
        const filename = `userAction_rank_${+new Date()}.xlsx`
        saveAs(new Blob([response.data], { type: response.headers['contnet-type'] }), filename)
      })
      .catch((error) => {
        console.error("下载文件失败：", error)
      })
}

onMounted(() => {
  nextTick(() => {
  })
})
</script>

<style scoped lang="scss">

</style>
