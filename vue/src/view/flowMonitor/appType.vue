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
      module-api="appTypeRankTableDataApi"
      @handleLinkClick="linkClickFunc"
    />
    <!--    二级页面    -->
    <el-dialog
      v-model="childrenComponentVisible"
      title=""
      dragable
      align-center
      style="width: 75%;height: auto"
    >
      <template #header>
        <div class="flex justify-between">
          <p>运营商 <b class="text-blue-600 text-lg">{{ childrenComponentName }}</b> 详情</p>
          <el-button-group class="mr-8">
            <el-button type="primary" size="small" :icon="ArrowLeft" @click="handlePrev">
              上一条
            </el-button>
            <el-button type="primary" size="small" :icon="ArrowRight" @click="handleNext">
              下一条
            </el-button>
          </el-button-group>
        </div>
      </template>
      <!--    趋势图    -->
      <ReTrend
        v-if="toolbarParam!==null"
        :key="loadKey2"
        title="流量趋势图"
        static-flag=""
        :search-params="toolbarParam"
        module-name="traffic"
        module-api="appTypeRankDataApi"
      />
      <!--    二级列表    -->
      <CommonTable
        v-if="toolbarParam!==null"
        ref="tableParam"
        :key="loadKey2+1"
        :level="2"
        :table-params="toolbarParam.form"
        :table-columns="tableColumns2"
        :operation="false"
        :auto-height="true"
        :page-size="10"
        module-name="traffic"
        module-api="appTypeRankTableDataApi"
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
// 状态管理
import useLinkStore from "@/pinia/modules/link.js";
import {appTypeRankDataApi, dictInfoApi} from "@/api/traffic";
import {useUserTreeStore} from "@/pinia/modules/useUserTree";

const LinkStore = useLinkStore()
const userTreeStore = useUserTreeStore()
// 定义工具栏配置
const datetimePicker = ref(true)
const linkOptions = ref([])
const ispOptions = ref([])
const userOptions = ref([])
const appTypeOptions = ref([{
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
  {name: "用户/群/组", field: "copyIdList", treeOptions: [], multiple: true, width: 'w-60'},
  {name: "应用大类", field: "appTypeIdList", treeOptions: [], multiple: true, width: 'w-48'},
])
const levelParam = ref({
  rankLevel: "level1"
})
// 定义其他参数
const loadKey = ref(0)
const loadKey2 = ref(0)
const clickIndex = ref(0)
const tableRecords = ref([])
// 定义饼图数据
const pieData = ref({})
const gatherData = ref({})
// 列表查询逻辑
const tableParam = ref(null)
const tableColumns = ref([
  {name: "业务大类", field: "name", link: true},
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
  { name: "时间", field: "startTime" },
  { name: "上行流速", field: "upBps", sortable: true },
  { name: "下行流速", field: "dnBps", sortable: true },
  { name: "总流速", field: "totalBps", sortable: true },
])
// 链接点击跳转，展示子组件逻辑
const childrenComponentVisible = ref(false)
const childrenComponentName = ref("")

const getPieData = async () => {
  pieData.value = []
  gatherData.value = []
  const res = await appTypeRankDataApi(toolbarParam.value.form)
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
  }).filter(item => item!==undefined)

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
    url: '/api/traffic/appType/export',
    data: toolbarParam.value.form,
    responseType: 'blob',
  })
      .then((response) => {
        const filename = `appType_rank_${+new Date()}.xlsx`
        saveAs(new Blob([response.data], { type: response.headers['contnet-type'] }), filename)
      })
      .catch((error) => {
        console.error("下载文件失败：", error)
      })
}
const linkClickFunc = (str, index, row) => {
  if (tableParam.value.data) {
    tableRecords.value = tableParam.value.data
  }
  // console.log(str, index, row)
  clickIndex.value = index
  childrenComponentVisible.value = true
  childrenComponentName.value = row.name
  toolbarParam.value.form.appTypeIdList[0] = parseInt(row.appType)
  toolbarParam.value.form.rankLevel = "level2"
  toolbarParam.value.form.limit = 10

  loadKey2.value = 0
}
const handlePrev = () => {
  clickIndex.value -= 1
  if (clickIndex.value < 0) {
    clickIndex.value = 0
  }
  const currentRow = tableRecords.value[clickIndex.value]
  loadKey2.value += 1
  childrenComponentName.value = currentRow.name
  toolbarParam.value.form.appTypeIdList[0] = parseInt(currentRow.appType)
  // console.log(toolbarParam.value.form.appTypeIdList)
}
const handleNext = () => {
  clickIndex.value += 1
   if (clickIndex.value > tableRecords.value.length) {
     clickIndex.value = tableRecords.value.length - 1
  }
  const currentRow = tableRecords.value[clickIndex.value]
  loadKey2.value += 1
  childrenComponentName.value = currentRow.name
  toolbarParam.value.form.appTypeIdList[0] = parseInt(currentRow.appType)
  // console.log(toolbarParam.value.form.appTypeIdList)
}
const renameIdToValue = (data) => {
  return data.map(item => ({
    ...item,
    value: item.id,
    name: item.name ? item.name : item.label,
    children: item.children ? renameIdToValue(item.children) : []
  }))
}

onBeforeMount( async () => {
  // 获取链路
  linkOptions.value = await LinkStore.getLink()
  linkOptions.value = linkOptions.value.map(item => {
    item.name = item.lineName
    item.value = item.lineVlan
    return item
  })

  // 获取用户群组树
  userOptions.value[0] = await userTreeStore.getUserTree()

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
  // 获取应用大类
  if (localStorage.getItem('appTypeOptions')) {
    appTypeOptions.value[0].children = JSON.parse(localStorage.getItem('appTypeOptions'))
  } else {
    const res3 = await dictInfoApi('appType')
    if (res3.code === 0) {
      appTypeOptions.value[0].children = res3.data
      localStorage.setItem('appTypeOptions', JSON.stringify(res3.data))
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
    } else if (item.name === "应用大类") {
      item.treeOptions = renameIdToValue(appTypeOptions.value)
    }
  })
})

onMounted( () => {
  toolbarParam.value.form = {...toolbarParam.value.form, ...levelParam.value}

  getPieData()
})
</script>

<style scoped lang="scss">

</style>
