<script setup>
import { formatSize } from "@/utils/format";
import UserBar from "@/view/flowMonitor/components/UserBar.vue";
import UserTrend from "@/view/flowMonitor/components/UserTrend.vue";
import { ref, nextTick } from "vue";
import UserLevel3View from "@/view/flowMonitor/components/UserLevel3View.vue";

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

const emit = defineEmits(['getLevel2TableData', 'getLevel2TrendChartData'])

const panelRef = ref(null)
const panelVisible = ref(false)
const panelTitle = ref("")

const trendChartData = ref([])
const fetchTrendChartData = async (params) => {
  const asyncFunc = await new Promise((resolve) => {
    emit('getLevel2TrendChartData', resolve)
  })

  const resp = await asyncFunc(params)
  if (resp.code === 0) {
    trendChartData.value = resp.data
  }
}

const barChartData = ref([])
const tableData = ref([])
const fetchTableData = async (params) => {
  const asyncFunc = await new Promise((resolve) => {
    emit('getLevel2TableData', resolve)
  })
  const resp = await asyncFunc(params)
  if (resp.code === 0) {
    tableData.value = resp.data
    barChartData.value = resp.data?.slice()?.sort((a, b) => a.trafficTotal - b.trafficTotal)
  }
}

const handleOpenPanel = async (row) => {
  row.loading = true

  await fetchTrendChartData(row)
  await fetchTableData(row)

  panelTitle.value = row.name
  panelVisible.value = true

  nextTick(() => {
    panelRef.value.scrollIntoView({ behavior: "smooth" })
    row.loading = false
  })
}

const level3ViewRef = ref(null)

const getLevel3TableData = async (params) => {
  emit("getLevel3TableData", params)
}


defineExpose({
  handleOpenPanel,
})
</script>

<template>
  <div ref="panelRef" v-if="panelVisible" class="bg-white dark:bg-slate-900  rounded-md shadow-md p-2">
    <div class="w-full h-full flex flex-col gap-3">
      <!-- Title -->
      <p class="text-2xl font-black ">[ {{ panelTitle }} ] - 流量趋势</p>

      <!--   Level2-Chart   -->
      <user-trend class="w-full h-72" title="" :data="trendChartData" x-field="startTime" :series-config="{
        speedUp: { label: '上行流量' },
        speedDn: { label: '下行流量' },
      }" />
      <user-bar class="w-full h-[510px]" title="IP排行" :data="barChartData" />

      <!--    Level2-Table    -->
      <div class="">
        <el-table :data="tableData" max-height="500">
          <el-table-column label="源IP" prop="srcIp">
            <template #default="{ row }">
              <el-button :loading="row.loading" size="small" icon="Link" type="primary" link
                @click="level3ViewRef?.handleOpenPanel(row)">
                {{ row.srcIp }}
              </el-button>
            </template>
          </el-table-column>
          <el-table-column label="上行平均" prop="avgSpeedUp" sortable
            :formatter="(row, column, cellValue) => formatSize(cellValue, { units: ['bps', 'Kbps', 'Mbps', 'Gbps'] })" />
          <el-table-column label="下行平均" prop="avgSpeedDn" sortable
            :formatter="(row, column, cellValue) => formatSize(cellValue, { units: ['bps', 'Kbps', 'Mbps', 'Gbps'] })" />
          <el-table-column label="上行总量" prop="trafficUp" sortable
            :formatter="(row, column, cellValue) => formatSize(cellValue, { units: ['bps', 'Kbps', 'Mbps', 'Gbps'] })" />
          <el-table-column label="下行总量" prop="trafficDn" sortable
            :formatter="(row, column, cellValue) => formatSize(cellValue, { units: ['bps', 'Kbps', 'Mbps', 'Gbps'] })" />
          <el-table-column label="上下行总量" prop="trafficTotal" sortable
            :formatter="(row, column, cellValue) => formatSize(cellValue)" />

          <el-table-column label="总量占比" prop="totalProportion" sortable>
            <template #default="{ row }">
              {{ (row.totalProportion * 100).toFixed(4) }}%
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>
  </div>


  <user-level3-view ref="level3ViewRef" :start-time="props.startTime" :end-time="props.endTime"
    @getTableData="getLevel3TableData" />
</template>

<style scoped lang="scss"></style>