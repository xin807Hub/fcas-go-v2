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
        @handleChange="changeFunc"
    />
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
        module-api="alarmLogListApi"
    />


  </div>
</template>

<script setup>

import ToolBar from "@/components/commonTable/toolbar.vue";
import CommonTable from "@/components/commonTable/index.vue";
import {onBeforeMount, ref, onMounted } from "vue";
import {dictInfoApi} from "@/api/traffic";

const loadKey = ref(0)
// 状态管理
import useLinkStore from "@/pinia/modules/link.js";
const LinkStore = useLinkStore()
// 定义工具栏配置
const datetimePicker = ref(true)
const linkOptions = ref([])
const toolbarParam = ref(null)
const appTypeOptions = ref([])
const appTypeIdTreeOptions = ref([])
// 列表查询逻辑
const tableParam = ref(null)
const tableColumns = ref([
  {name: "开始时间", field: "start_time"},
  {name: "策略名称", field: "policy_name", sortable: true},
  {name: "应用大类", field: "app_type_name", sortable: true},
  {name: "应用小类", field: "app_name", sortable: true},
  {name: "链路", field: "link_name", sortable: true},
  {name: "告警类型", field: "alarm_type", sortable: true},
  {name: "流量上浮/下降比例", field: "flowPercent", sortable: true},
  {name: "实际流速", field: "realFlowSpeed", sortable: true},
])
const searchEl = ref([
  {name: "应用大类", field: "appTypeId", options: [], width: 'w-40'},
  {name: "应用小类", field: "appId", options: [], width: 'w-40'},
  {name: "链路", field: "linkIds", options: [], multiple: true, width: 'w-112'},
])

const searchFunc = () => { // 搜索
  toolbarParam.value.form.startTime = new Date(toolbarParam.value.form.startTime)
  toolbarParam.value.form.endTime = new Date(toolbarParam.value.form.endTime)
  loadKey.value += 1
}

const changeFunc = (str, val, field) => {
  if (localStorage.getItem("appTypeIdTreeOptions")) {
    const res = JSON.parse(localStorage.getItem("appTypeIdTreeOptions"))
    if (field === 'appTypeId') {
      searchEl.value.forEach(item => {
        if (item.name === "应用小类") {
          item.options = val ? res.filter(item => item.id === val)[0].children.map(item => {
            return {
              value: item.id,
              name: item.name
            }
          }) : []
        }
      })
    }
  }
}

onBeforeMount(async () => {
  // 获取链路
  linkOptions.value = await LinkStore.getLink()
  linkOptions.value = linkOptions.value.map(item => {
    item.name = item.lineName
    item.value = item.lineVlan
    return item
  })

// 获取应用大类
  if (localStorage.getItem('appTypeOptions')) {
    appTypeOptions.value = JSON.parse(localStorage.getItem('appTypeOptions'))
  } else {
    const res3 = await dictInfoApi('appType')
    if (res3.code === 0) {
      appTypeOptions.value = res3.data
      localStorage.setItem('appTypeOptions', JSON.stringify(res3.data))
    }
  }

  // 获取应用小类
  if (localStorage.getItem('appTypeIdTreeOptions')) {
    appTypeIdTreeOptions.value = JSON.parse(localStorage.getItem('appTypeIdTreeOptions'))
  } else {
    const res4 = await dictInfoApi('appTypeIdTree')
    if (res4.code === 0) {
      appTypeIdTreeOptions.value = res4.data
      localStorage.setItem('appTypeIdTreeOptions', JSON.stringify(res4.data))
    }
  }

  appTypeOptions.value = appTypeOptions.value.map(item => {
    return {
      value: item.id,
      name: item.name
    }
  })

  searchEl.value.forEach(item => {
    if (item.name === "应用大类") {
      item.options = (appTypeOptions.value)
    }else if (item.name==="链路"){
      item.options = linkOptions.value
    }
  })
})

const refreshFunc = () => {
  loadKey.value += 1
}

onMounted(() => {
  toolbarParam.value.form.startTime = new Date(toolbarParam.value.form.startTime)
  toolbarParam.value.form.endTime = new Date(toolbarParam.value.form.endTime)
})

</script>

<style scoped lang="scss">

</style>
