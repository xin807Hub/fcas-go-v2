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
          static-flag="trafficTotal"
          :data="pieData"
        />
      </gva-card>
      <gva-card title="下行总量占比" custom-class="col-span-1 h-80" :without-padding="false">
        <RePie
          v-if="Object.values(pieData).length > 0"
          :key="loadKey"
          title=""
          static-flag="trafficDn"
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
      :operation="false"
      :auto-height="true"
      module-name="traffic"
      module-api="userRankLevel1TableApi"
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
          <p>用户 <b class="text-blue-600 text-lg">{{ childrenComponentName }}</b> 详情</p>
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
      <div :key="loadKey2" class="w-full h-176 overflow-y-scroll">
        <!--    趋势图    -->
        <ReTrend
          v-if="toolbarParam!==null"
          :key="loadKey2"
          title="流量趋势图"
          static-flag=""
          :search-params="toolbarParam"
          module-name="traffic"
          module-api="userRankLevel2TrendApi"
        />
        <div class="flex justify-end items-center mr-8">
          <el-input-number v-model="topN" :min="1" :max="20" size="small" class="mr-4" />
          <el-button type="primary" icon="Search" class="button search-button" size="small" @click="handleTopSearch">
            查询
          </el-button>
        </div>
        <!--    TOP图    -->
        <ReTop
          v-if="toolbarParam!==null"
          :key="loadKey2"
          title="IP排行"
          static-flag=""
          :search-params="toolbarParam"
          :interval="0"
          :line="false"
          module-name="traffic"
          module-api="userRankLevel2TableApi"
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
          module-api="userRankLevel2TableApi"
          @handleLinkClick="linkClickFunc2"
        />
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
        module-api="userCrowdGroupRankLevel3TrendApi"
      />
      <!--    三级列表    -->
      <CommonTable
        v-if="toolbarParam!==null"
        :key="loadKey3+1"
        ref="tableParam"
        :level="3"
        :table-params="toolbarParam.form"
        :table-columns="tableColumns3"
        :operation="false"
        :auto-height="true"
        :page-size="10"
        module-name="traffic"
        module-api="userCrowdGroupRankLevel3TableApi"
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
import ReTop from "@/components/charts/ReTop.vue";
// 引入接口
import {dictInfoApi, userRankLevel1PieApi} from "@/api/traffic";
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
const searchEl = ref([
  {name: "运营商", field: "ispNameList", options: [], multiple: true, width: 'w-40'},
  {name: "链路", field: "linkIdList", options: [], multiple: true, width: 'w-112'},
  {name: "用户/群/组", field: "copyIdList", treeOptions: [], multiple: true, width: 'w-60', lazy: false},
  {name: "应用大小类", field: "appIdList", treeOptions: [], multiple: true, width: 'w-60', lazy: true},
])
const loadKey = ref(0)
const loadKey2 = ref(0)
const clickIndex = ref(0)
const tableRecords = ref([])
const loadKey3 = ref(0)
const clickIndex2 = ref(0)
const tableRecords2 = ref([])
// 定义饼图数据
const pieData = ref({})
const gatherData = ref({})
// 列表查询逻辑
const tableParam = ref(null)
const tableColumns = ref([
  {name: "用户", field: "name", link: true},
  {name: "上行峰值", field: "maxSpeedUp", sortable: true},
  {name: "下行峰值", field: "maxSpeedDn", sortable: true},
  {name: "上行平均", field: "avgSpeedUp", sortable: true},
  {name: "下行平均", field: "avgSpeedDn", sortable: true},
  {name: "上行总量", field: "trafficUp", sortable: true},
  {name: "下行总量", field: "trafficDn", sortable: true},
  {name: "上下行总量", field: "trafficTotal", sortable: true},
  {name: "总量占比", field: "totalRatio"},
])
const tableColumns2 = ref([
  {name: "源IP", field: "srcIp", width: '180', link: true},
  {name: "上行平均", field: "avgSpeedUp"},
  {name: "下行平均", field: "avgSpeedDn"},
  {name: "上行总量", field: "trafficUp"},
  {name: "下行总量", field: "trafficDn"},
  {name: "上下行总量", field: "trafficTotal"},
  {name: "总量占比", field: "totalProportion"},
])
const tableColumns3 = ref([
  {name: "时间", field: "startTime"},
  {name: "上行流速", field: "speedUp"},
  {name: "下行流速", field: "speedDn"},
  {name: "总流速", field: "totalSpeed"},
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
  const res = await userRankLevel1PieApi(toolbarParam.value.form)
  if (res.code === 0) {
    pieData.value = res.data
    gatherData.value = res.data.gatherData ?? {}
  }
}
// 查询
const searchFunc = () => {
  // console.log(toolbarParam.value.form.copyIdList)
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
    url: '/api/traffic/userRank/export',
    data: toolbarParam.value.form,
    responseType: 'blob',
  })
      .then((response) => {
        const filename = `user_rank_${+new Date()}.xlsx`
        saveAs(new Blob([response.data], {type: response.headers['contnet-type']}), filename)
      })
      .catch((error) => {
        console.error("下载文件失败：", error)
      })
}

const linkClickFunc = (str, index, row) => {
  clickIndex.value = index
  childrenComponentVisible.value = true
  childrenComponentName.value = row.name
  toolbarParam.value.form.copyIdList[0] = parseInt(row.id)
  toolbarParam.value.form.userIdList = toolbarParam.value.form.copyIdList
  toolbarParam.value.form.topN = 20

  if (tableParam.value.data) {
    tableRecords.value = tableParam.value.data
  }

  loadKey2.value += 1
}
const handlePrev = (str) => {
  if (str === 'level2') {
    clickIndex.value -= 1
    if (clickIndex.value < 0) {
      clickIndex.value = 0
    }
    const currentRow = tableRecords.value[clickIndex.value]
    childrenComponentName.value = currentRow.name
    toolbarParam.value.form.copyIdList[0] = parseInt(currentRow.id)
    toolbarParam.value.form.userIdList = toolbarParam.value.form.copyIdList
    toolbarParam.value.form.topN = 20

    loadKey2.value += 1
  } else {
    clickIndex2.value -= 1
    if (clickIndex2.value < 0) {
      clickIndex2.value = 0
    }
    const currentRow = tableRecords2.value[clickIndex2.value]
    childrenComponentName2.value = currentRow.srcIp
    toolbarParam.value.form.srcIp = currentRow.srcIp

    loadKey3.value += 1
  }
}
const handleNext = (str) => {
  if (str === 'level2') {
    clickIndex.value += 1
    if (clickIndex.value >= tableRecords.value.length) {
      clickIndex.value = tableRecords.value.length - 1
    }
    const currentRow = tableRecords.value[clickIndex.value]
    childrenComponentName.value = currentRow.name
    toolbarParam.value.form.copyIdList[0] = parseInt(currentRow.id)
    toolbarParam.value.form.userIdList = toolbarParam.value.form.copyIdList
    toolbarParam.value.form.topN = 20

    loadKey2.value += 1
  } else {
    clickIndex2.value += 1
    if (clickIndex2.value >= tableRecords2.value.length) {
      clickIndex2.value = tableRecords2.value.length - 1
    }
    const currentRow = tableRecords2.value[clickIndex2.value]
    childrenComponentName2.value = currentRow.srcIp
    toolbarParam.value.form.srcIp = currentRow.srcIp

    loadKey3.value += 1
  }
}
const handleTopSearch = () => {
  toolbarParam.value.form.topN = topN.value
  loadKey2.value += 1
}
const linkClickFunc2 = (str, index, row) => {
  clickIndex.value = index
  childrenComponentVisible2.value = true
  childrenComponentName2.value = row.srcIp
  toolbarParam.value.form.copyIdList[0] = parseInt(row.id)
  toolbarParam.value.form.userIdList = toolbarParam.value.form.copyIdList
  toolbarParam.value.form.srcIp = row.srcIp

  loadKey3.value += 1
  if (tableParam.value.data) {
    tableRecords2.value = tableParam.value.data
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

onBeforeMount(async () => {
  // 获取链路
  linkOptions.value = await LinkStore.getLink()
  linkOptions.value = linkOptions.value.map(item => {
    item.name = item.lineName
    item.value = item.lineVlan
    return item
  })

  userOptions.value[0] = await userTreeStore.getUserTree()

  // 获取运营商
  if (localStorage.getItem('ispOptions')) {
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
  getPieData()
})
</script>

<style scoped lang="scss">

</style>
