<script setup>
import { formatSize } from "@/utils/format";
import UserTrend from "@/view/flowMonitor/components/UserTrend.vue";
import { ref, nextTick } from "vue";
import dayjs from "dayjs";


const props = defineProps({
  startTime: {
    type: String,
    default: '',
    required: false,
  },
  endTime: {
    type: String,
    default: '',
    required: true,
  },
})

const emit = defineEmits(['getNewRow', 'getTableData'])

const panelVisible = ref(false)
const panelRef = ref(null)
const panelTitle = ref("")



const tableData = ref([])
const fetchTableData = async (params) => {
  const asyncFunc = await new Promise((resolve) => {
    emit('getTableData', resolve)
  })
  const resp = await asyncFunc(params)
  if (resp.code === 0) {
    tableData.value = resp.data
  }
}


const handleOpenPanel = async (row) => {
  row.loading = true

  await fetchTableData(row)

  panelTitle.value = row.srcIp
  panelVisible.value = true


  nextTick(() => {
    panelRef.value.scrollIntoView({ behavior: "smooth" })
    row.loading = false
  })
}

defineExpose({
  handleOpenPanel,
})
</script>

<template>
  <div ref="panelRef" v-if="panelVisible" class="bg-white dark:bg-slate-900 rounded-md shadow-md p-2">
    <div class="w-full h-full space-y-2">
      <!-- Title -->
      <p class="text-2xl font-black ">[ {{ panelTitle }} ] - 流速详情</p>

      <!--    Level3-Chart    -->
      <user-trend class="w-full h-72" title="" :data="tableData" x-field="startTime"
        :series-config="{
          speedUp: { label: '上行流速' },
          speedDn: { label: '下行流速' },
        }" />
      <!--  Level3-Table  -->
      <div class="">
        <el-table :data="tableData" max-height="500">
          <el-table-column label="时间" prop="startTime" sortable
            :formatter="(row) => dayjs(row.startTime).format('YYYY-MM-DD HH:mm:ss')" />
          <el-table-column label="上行流速" prop="speedUp" sortable
            :formatter="(row, column, cellValue) => formatSize(cellValue, { units: ['bps', 'Kbps', 'Mbps', 'Gbps'] })" />
          <el-table-column label="下行流速" prop="speedDn" sortable
            :formatter="(row, column, cellValue) => formatSize(cellValue, { units: ['bps', 'Kbps', 'Mbps', 'Gbps'] })" />
          <el-table-column label="总流速" prop="totalSpeed" sortable
            :formatter="(row, column, cellValue) => formatSize(cellValue, { units: ['bps', 'Kbps', 'Mbps', 'Gbps'] })" />
        </el-table>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss"></style>