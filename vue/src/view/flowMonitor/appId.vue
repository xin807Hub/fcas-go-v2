<template>
  <div>
    <!--    工具栏    -->
    <ToolBar
      ref="toolbarParam"
      :search-el="searchEl"
      :datetime-picker="datetimePicker"
      :export-flag="true"
      @handleSearch="searchFunc"
      @handleRefresh="refreshFunc"
      @handleExport="exportFunc"
    />
    <!--    饼图    -->
    <div class="grid grid-cols-1 md:grid-cols-1 lg:grid-cols-2 py-0 gap-4 md:gap-3 mt-2">
      <gva-card title="上下行总量占比" custom-class="col-span-1 h-80" :without-padding="false">
        <RePie
          v-if="Object.values(pieData).length > 0"
          :key="loadKey"
          title=""
          static-flag="totalByte"
          :data="pieData"
        />
      </gva-card>
      <gva-card title="下行总量占比" custom-class="col-span-1 h-80" :without-padding="false">
        <RePie
          v-if="Object.values(pieData).length > 0"
          :key="loadKey"
          title=""
          static-flag="dnByte"
          :data="pieData"
        />
      </gva-card>
    </div>
    <!--    列表    -->
    <CommonTable
      v-if="toolbarParam!==null"
      :key="loadKey"
      ref="tableParam"
      :table-params="toolbarParam.form"
      :table-columns="tableColumns"
      :gather-data="gatherData"
      :operation="false"
      :auto-height="true"
      module-name="traffic"
      module-api="appIdRankTrendDataApi"
      @handleLinkClick="linkClickFunc"
    />
    <!--    二级页面    -->
    <el-dialog
      v-model="childrenComponentVisible"
      title=""
      dragable
      align-center
      class="w-3/4 h-198"
    >
      <template #header>
        <div class="flex justify-between">
          <p>业务小类 <b class="text-blue-600 text-lg">{{ childrenComponentName }}</b> 详情</p>
          <el-button-group class="mr-8">
            <el-button type="primary" size="small" :icon="ArrowLeft" @click="handlePrev('level2')">
              上一条
            </el-button>
            <el-button type="primary" size="small" :icon="ArrowRight" @click="handleNext('level2')">
              下一条
            </el-button>
          </el-button-group>
        </div>
      </template>
      <div
        :key="loadKey2"
        class="w-full h-176 overflow-y-scroll"
      >
        <!--    趋势图    -->
        <ReTrend
          v-if="trendData2!==null"
          title="流量趋势图"
          static-flag=""
          :trend-data="trendData2"
        />
        <!--    二级列表    -->
        <CommonTable
            v-if="toolbarParam!==null"
            ref="tableParam"
            :table-params="toolbarParam.form"
            :table-columns="tableColumns2"
            :operation="false"
            :auto-height="true"
            module-name="traffic"
            module-api="appIdRankTrendDataApi"
            @handleLinkClick="linkClickFunc"
        />

        <!--  柱状图切换TopN  -->
        <div class="flex justify-end items-center mr-8 mt-2" v-if="barDstIpData.length > 0 || barHostData.length > 0">
          <el-input-number v-model="topN" :min="1" :max="20" size="small" class="mr-4"/>
          <el-button type="primary" icon="Search" class="button search-button" size="small" @click="getLevel2Data">
            查询
          </el-button>
        </div>

        <!--   目的IP和域名排行   -->
        <ReDoubleBar
            v-if="barDstIpData.length > 0 || barHostData.length > 0"
            :left-data="barDstIpData"
            :right-data="barHostData"
            :legend-data="['目的IP', '域名']"
        />
        <CommonTable
            v-if="toolbarParam!==null"
            ref="tableParamDstIp"
            :level="2"
            :table-params="{...toolbarParam.form, ...{rankType:'dstIp'}}"
            :table-columns="[
            { name: '目的IP', field: 'name', link: true },
            { name: '归属信息', field: 'location' },
            { name: '上行平均', field: 'avgUpBps' },
            { name: '下行平均', field: 'avgDnBps' },
            { name: '上行总量', field: 'upByte' },
            { name: '下行总量', field: 'dnByte' },
          ]"
            :operation="false"
            :auto-height="true"
            :page-size="10"
            module-name="traffic"
            module-api="appIdRankTableDataApi"
            @handleLinkClick="linkClickFuncDstIp"
        />
        <CommonTable
            v-if="toolbarParam!==null"
            ref="tableParamHost"
            :level="2"
            :table-params="{...toolbarParam.form, ...{rankType:'host'}}"
            :table-columns="[
            { name: '域名', field: 'name', link: true },
            { name: '上行平均', field: 'avgUpBps' },
            { name: '下行平均', field: 'avgDnBps' },
            { name: '上行总量', field: 'upByte' },
            { name: '下行总量', field: 'dnByte' },
          ]"
            :operation="false"
            :auto-height="true"
            :page-size="10"
            module-name="traffic"
            module-api="appIdRankTableDataApi"
            @handleLinkClick="linkClickFuncHost"
        />
        <el-backtop :visibility-height="300" :right="100" :bottom="100" />
      </div>
    </el-dialog>
    <!--    三级页面    -->
    <el-dialog
      v-model="childrenComponentVisible2"
      title=""
      dragable
      align-center
      style="width: 75%;height: auto"
    >
      <template #header>
        <div class="flex justify-between">
          <p><b class="text-blue-600 text-lg">{{ childrenComponentName2 }}</b> 详情</p>
          <el-button-group class="mr-8">
            <el-button type="primary" size="small" :icon="ArrowLeft" @click="handlePrev('level3')">
              上一条
            </el-button>
            <el-button type="primary" size="small" :icon="ArrowRight" @click="handleNext('level3')">
              下一条
            </el-button>
          </el-button-group>
        </div>
      </template>
      <!--    趋势图    -->
      <ReTrend
        v-if="toolbarParam!==null"
        :key="loadKey3"
        title="流量趋势图"
        static-flag=""
        :search-params="toolbarParam"
        module-name="traffic"
        module-api="appIdRankDataApi"
      />
      <!--    三级列表    -->
      <CommonTable
        v-if="toolbarParam!==null"
        :key="loadKey3+1"
        ref="tableParam"
        :level="3"
        :table-params="toolbarParam.form"
        :table-columns="tableColumns2"
        :operation="false"
        :auto-height="true"
        :page-size="10"
        module-name="traffic"
        module-api="appIdRankTrendDataApi"
      />
    </el-dialog>
  </div>
</template>

<script setup>
import {onBeforeMount, onMounted, ref} from "vue"
import axios from "axios"
import {saveAs} from "file-saver"
import {ArrowLeft, ArrowRight} from '@element-plus/icons-vue'
// 引入公共组件
import ToolBar from "@/components/commonTable/toolbar.vue"
import CommonTable from "@/components/commonTable/index.vue"
import {GvaCard} from "@/view/dashboard/components"
import RePie from "@/components/charts/RePie.vue";
import ReTrend from "@/components/charts/ReTrend.vue";
import ReDoubleBar from "@/components/charts/ReDoubleBar.vue";
// 引入接口
import {appIdRankDataApi, dictInfoApi} from "@/api/traffic"
// 状态管理
import useLinkStore from "@/pinia/modules/link.js";
import {useUserTreeStore} from "@/pinia/modules/useUserTree";

const LinkStore = useLinkStore()
const userTreeStore = useUserTreeStore()

// 定义工具栏配置
const datetimePicker = ref(true)
const linkOptions = ref([])
const ispOptions = ref([])
const userOptions = ref([])
const appTypeIdOptions = ref([{
  id: 0,
  name: "全部",
  children: []
}])
const toolbarParam = ref(null)
const provinceOptions = [
  {name: "上海", value: "上海"},
]
const searchEl = ref([
  {name: "运营商", field: "ispNameList", options: [], multiple: true, width: 'w-40'},
  {name: "链路", field: "linkIdList", options: [], multiple: true, width: 'w-60'},
  {name: "区域", field: "dstProvince", options: provinceOptions, multiple: false, width: 'w-32'},
  {name: "用户/群/组", field: "copyIdList", treeOptions: [], multiple: true, width: 'w-60', lazy: false},
  {name: "应用大小类", field: "appIdList", treeOptions: [], multiple: true, width: 'w-60', lazy: true},
])
const levelParam = ref({
  rankLevel: "level1"
})
// 定义其他参数
const loadKey = ref(0)
const loadKey2 = ref(0)
const clickIndex = ref(0)
const tableRecords = ref([])
const loadKey3 = ref(0)
const clickIndex2 = ref(0)
const tableRecordsDstIp = ref([])
const tableRecordsHost = ref([])
const level2Flag = ref("")
// 定义饼图数据
const pieData = ref({})
const gatherData = ref({})
// 列表查询逻辑
const tableParam = ref(null)
const tableParamDstIp = ref(null)
const tableParamHost = ref(null)
const tableColumns = ref([
  {name: "业务小类", field: "name", link: true},
  {name: "上行峰值", field: "maxUpBps", sortable: true},
  {name: "下行峰值", field: "maxDnBps", sortable: true},
  {name: "上行平均", field: "avgUpBps", sortable: true},
  {name: "下行平均", field: "avgDnBps", sortable: true},
  {name: "上行总量", field: "upByte", sortable: true},
  {name: "下行总量", field: "dnByte", sortable: true},
  {name: "上下行总量", field: "totalByte", sortable: true},
  {name: "总量占比", field: "totalRatio"},
])
const tableColumns2 = ref([
  {name: "时间", field: "startTime"},
  {name: "上行流速", field: "upBps", sortable: true},
  {name: "下行流速", field: "dnBps", sortable: true},
  {name: "总流速", field: "totalBps", sortable: true},
])
// 链接点击跳转，展示子组件逻辑
const childrenComponentVisible = ref(false)
const childrenComponentName = ref("")
const childrenComponentVisible2 = ref(false)
const childrenComponentName2 = ref("")

const topN = ref(20)

const getPieData = async () => {
  pieData.value = []
  gatherData.value = []
  const res = await appIdRankDataApi(toolbarParam.value.form)
  if (res.code === 0) {
    pieData.value = res.data
    gatherData.value = res.data.gatherData ?? {}
  }
}
// 查询
const searchFunc = () => {
  toolbarParam.value.form.isp = null
  toolbarParam.value.form.rankLevel = "level1"
  toolbarParam.value.form.limit = 20
  toolbarParam.value.form.userIdList = toolbarParam.value.form.copyIdList.map(item => {
    if (item.split('-').length === 3) {
      const lastPart = item.split('-').pop()
      return parseInt(lastPart, 10)
    } else {
      return undefined
    }
  }).filter(item => item !== undefined)

  getPieData()
  loadKey.value += 1
}
// 刷新
const refreshFunc = () => {
  loadKey.value += 1
}
// 导出
const exportFunc = async () => {
  axios({
    method: "post",
    url: '/api/traffic/appId/export',
    data: toolbarParam.value.form,
    responseType: 'blob',
  })
      .then((response) => {
        const filename = `appId_rank_${+new Date()}.xlsx`
        saveAs(new Blob([response.data], {type: response.headers['contnet-type']}), filename)
      })
      .catch((error) => {
        console.error("下载文件失败：", error)
      })
}

// 获取二级页面数据
const barDstIpData = ref([])
const barHostData = ref([])
const trendData2 = ref(null)
const getLevel2Data = async () => {
  toolbarParam.value.form.topN = topN.value
  barDstIpData.value = []
  barHostData.value = []
  const res = await appIdRankDataApi(toolbarParam.value.form)
  if (res.code === 0) {
    barDstIpData.value = res.data.dstIpRankBar ?? []
    barHostData.value = res.data.hostRankBar ?? []
    trendData2.value = res.data.trendData ?? []
  }
}
const linkClickFunc = async (str, index, row) => {
  clickIndex.value = index
  childrenComponentVisible.value = true
  childrenComponentName.value = row.name
  toolbarParam.value.form.appIdList[0] = parseInt(row.appId)
  toolbarParam.value.form.dstIpParam = null
  toolbarParam.value.form.hostParam = null
  toolbarParam.value.form.rankLevel = "level2"
  toolbarParam.value.form.limit = 10

  if (tableParam.value.data) {
    tableRecords.value = tableParam.value.data
  }

  await getLevel2Data()
  loadKey2.value = 0
}
const handlePrev = async (str) => {
  if (str === 'level2') {
    clickIndex.value -= 1
    if (clickIndex.value < 0) {
      clickIndex.value = 0
    }
    const currentRow = tableRecords.value[clickIndex.value]
    childrenComponentName.value = currentRow.name
    toolbarParam.value.form.appIdList[0] = parseInt(currentRow.appId)

    await getLevel2Data()
    loadKey2.value += 1
  } else {
    clickIndex2.value -= 1
    if (clickIndex2.value < 0) {
      clickIndex2.value = 0
    }
    if (level2Flag.value === "DstIp") {
      const currentRow = tableRecordsDstIp.value[clickIndex2.value]
      childrenComponentName2.value = currentRow.name
      toolbarParam.value.form.dstIpParam = currentRow.name
    } else {
      const currentRow = tableRecordsHost.value[clickIndex2.value]
      childrenComponentName2.value = currentRow.name
      toolbarParam.value.form.hostParam = currentRow.name
    }

    loadKey3.value += 1
  }
}
const handleNext = async (str) => {
  if (str === 'level2') {
    clickIndex.value += 1
    if (clickIndex.value > tableRecords.value.length) {
      clickIndex.value = tableRecords.value.length - 1
    }
    const currentRow = tableRecords.value[clickIndex.value]
    childrenComponentName.value = currentRow.name
    toolbarParam.value.form.appIdList[0] = parseInt(currentRow.appId)

    await getLevel2Data()
    loadKey2.value += 1
  } else {
    clickIndex2.value += 1
    if (level2Flag.value === "DstIp") {
      if (clickIndex2.value > tableRecordsDstIp.value.length) {
        clickIndex2.value = tableRecordsDstIp.value.length - 1
      }
      const currentRow = tableRecordsDstIp.value[clickIndex2.value]
      childrenComponentName2.value = currentRow.name
      toolbarParam.value.form.dstIpParam = currentRow.name
    } else {
      if (clickIndex2.value > tableRecordsHost.value.length) {
        clickIndex2.value = tableRecordsHost.value.length - 1
      }
      const currentRow = tableRecordsHost.value[clickIndex2.value]
      childrenComponentName2.value = currentRow.name
      toolbarParam.value.form.hostParam = currentRow.name
    }
    loadKey3.value += 1
  }
}
const linkClickFuncDstIp = (str, index, row) => {
  clickIndex2.value = index
  childrenComponentVisible2.value = true
  childrenComponentName2.value = row.name
  toolbarParam.value.form.dstIpParam = row.name
  toolbarParam.value.form.rankLevel = "level3"
  toolbarParam.value.form.limit = 10

  loadKey3.value += 1
  level2Flag.value = "DstIp"
  if (tableParamDstIp.value.data) {
    tableRecordsDstIp.value = tableParamDstIp.value.data
  }
}
const linkClickFuncHost = (str, index, row) => {
  clickIndex2.value = index
  childrenComponentVisible2.value = true
  childrenComponentName2.value = row.name
  toolbarParam.value.form.hostParam = row.name
  toolbarParam.value.form.rankLevel = "level3"
  toolbarParam.value.form.limit = 10

  loadKey3.value += 1
  level2Flag.value = "Host"
  if (tableParamHost.value.data) {
    tableRecordsHost.value = tableParamHost.value.data
  }
}
const renameIdToValue = (data) => {
  return data.map(item => ({
    ...item,
    value: item.id,
    name: item.name ? item.name : item.label,
    children: item.children ? renameIdToValue(item.children) : []
  }))
}

const handleTopSearch = () => {
  toolbarParam.value.form.topN = topN.value
  loadKey2.value += 1
}

onBeforeMount(async () => {
  // 获取链路
  linkOptions.value = await LinkStore.getLink()
  linkOptions.value = linkOptions.value.map(item => {
    item.name = item.lineName
    item.value = item.lineVlan
    return item
  })
  // 获取用户群组树
  userOptions.value[0] =  await userTreeStore.getUserTree()
  // 获取运营商
  if (localStorage.getItem('ispOptions') !== '[]') {
    ispOptions.value = JSON.parse(localStorage.getItem('ispOptions'))
  } else {
    const res2 = await dictInfoApi('ispSelect')
    if (res2.code === 0) {
      ispOptions.value = res2.data
      localStorage.setItem('ispOptions', JSON.stringify(res2.data))
    }
  }

  // 获取应用大小类
  if (localStorage.getItem('appTypeIdTree')) {
    appTypeIdOptions.value[0].children = JSON.parse(localStorage.getItem('appTypeIdTree'))
  } else {
    const res3 = await dictInfoApi('appTypeIdTree')
    if (res3.code === 0) {
      appTypeIdOptions.value[0].children = res3.data
      localStorage.setItem('appTypeIdTree', JSON.stringify(res3.data))
    }
  }

  searchEl.value.forEach(item => {
    if (item.name === "运营商") {
      item.options = ispOptions.value.map(item => {
        item.value = item.name
        return item
      })
    } else if (item.name === "链路") {
      item.options = linkOptions.value
    } else if (item.name === "用户/群/组") {
      item.treeOptions = renameIdToValue(userOptions.value)
    } else if (item.name === "应用大小类") {
      item.treeOptions = renameIdToValue(appTypeIdOptions.value)
    }
  })
})

onMounted(() => {
  toolbarParam.value.form = {...toolbarParam.value.form, ...levelParam.value}

  getPieData()
})
</script>

<style scoped lang="scss">

</style>
