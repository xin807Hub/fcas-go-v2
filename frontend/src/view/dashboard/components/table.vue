<template>
  <div>
    <el-table
      :data="tableData"
      fit
      stripe
      border
      show-overflow-tooltip
      highlight-current-row
      header-row-class-name="row-class-name"
      table-layout="auto"
      class="w-full"
    >
      <el-table-column prop="linkName" label="链路名称" width="600" align="center" />
      <el-table-column prop="maxUpBps" label="上行峰值" align="center" />
      <el-table-column prop="maxDnBps" label="下行峰值" align="center" />
      <el-table-column prop="avgUpBps" label="上行平均" align="center" />
      <el-table-column prop="avgDnBps" label="下行平均" align="center" />
    </el-table>
  </div>
</template>

<script setup>
import {ref} from "vue"
// 引入字节格式化方法
import {getDataByteFormat} from "@/utils/methods.js"
// 接收父组件传参
const props = defineProps({
  data: Array
})
const tableData = ref([])
tableData.value = props.data.map(item => {
  const _item = {
    linkName: item.linkName,
    maxUpBps: getDataByteFormat(item.maxUpBps, 'bps'),
    maxDnBps: getDataByteFormat(item.maxDnBps, 'bps'),
    avgUpBps: getDataByteFormat(item.avgUpBps, 'bps'),
    avgDnBps: getDataByteFormat(item.avgDnBps, 'bps')
  }
  return _item
})
</script>

<style scoped lang="scss">

</style>
