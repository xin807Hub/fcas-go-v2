<template>
  <div>
    <el-card class="mt-2">
      <!--   表格主体   -->
      <el-table
        class="w-full"
        :loading="loading"
        fit
        stripe
        size="small"
        highlight-current-row
        header-row-class-name="row-class-name"
        show-overflow-tooltip
        table-layout="auto"
        :data="data"
        :height="getHeight"
        :expand-row-keys="expandKeys"
        row-key="name"
        @row-click="handleRowClick"
        @selection-change="handleSelectionChange"
        @expand-change="handleExpandChange"
      >
        <el-table-column
          v-if="expand!==''"
          type="expand"
        >
          <template #default="props">
            <el-table
              class="w-5/6 ml-14"
              :loading="loading"
              fit
              stripe
              size="small"
              highlight-current-row
              header-row-class-name="row-class-name"
              show-overflow-tooltip
              table-layout="auto"
              :data="expandData"
            >
              <el-table-column
                v-for="(item, index) in expandColumns"
                :key="index"
                show-overflow-tooltip
                :prop="item.field"
                :label="item.name"
                :sortable="item.sortable"
                :width="item.width"
                :formatter="colFormatter"
              />
            </el-table>
          </template>
        </el-table-column>
        <el-table-column
          type="index"
          label="序号"
          width="80"
        />
        <el-table-column
          v-for="(item, index) in columns"
          :key="index"
          show-overflow-tooltip
          :prop="item.field"
          :label="item.name"
          :sortable="item.sortable"
          :width="item.width"
          :formatter="colFormatter"
        >
          <template
            v-if="item.link===true"
            #default="scope"
          >
            <el-link
              :icon="Link"
              type="primary"
              @click="handleLinkClick(scope.$index, scope.row)"
            >
              {{ scope.row[item.field] }}
            </el-link>
          </template>
          <template
            v-if="item.field.indexOf('Time') > - 1"
            #default="scope"
          >
            {{ formatTimeToStr(scope.row[item.field]) }}
          </template>
          <template
            v-if="item.field.indexOf('time') > - 1"
            #default="scope"
          >
            {{
              (scope.row['dl_tos'] || scope.row['ul_tos']) ? formatTimeToStr(scope.row[item.field]) : ((scope.row['ul_flow_rate'] || scope.row['dl_flow_rate']) ? formatTimeToStr(scope.row[item.field]) : scope.row[item.field])
            }}
          </template>
          <template
            v-else-if="item.field==='crowds'"
            #default="scope"
          >
            <el-tag
              v-for="crowd in scope.row[item.field]"
              type="success"
              size="small"
              class="mr-1"
            >
              {{ crowd.crowdName }}
            </el-tag>
          </template>
          <template
            v-else-if="item.field==='users'"
            #default="scope"
          >
            <el-tag
              v-for="user in scope.row[item.field]"
              type="success"
              size="small"
              class="mr-1"
            >
              {{ user.userName }}
            </el-tag>
          </template>
          <template
            v-else-if="item.field==='ipAddress'"
            #default="scope"
          >
            <div
              v-for="(addr, index) in scope.row[item.field]"
              :key="index"
            >
              <el-tag type="success">
                {{ addr }}
              </el-tag>
            </div>
          </template>
          <template
            v-else-if="item.field.indexOf('Ip') > -1 && item.field !== 'srcIp'"
            #default="scope"
          >
            <el-tag
              type="primary"
              size="small"
              effect="dark"
              class="mr-1"
            >
              {{ scope.row[item.field] }}
            </el-tag>
          </template>
          <template
            v-else-if="item.field==='linkState'"
            #default="scope"
          >
            <el-tag
              v-if="scope.row[item.field] === 0"
              type="danger"
              size="small"
            >
              断开
            </el-tag>
            <el-tag
              v-else-if="scope.row[item.field] === 1"
              type="success"
              size="small"
            >
              正常
            </el-tag>
          </template>
          <template
            v-else-if="item.field==='totalRatio'"
            #default="scope"
          >
            {{
              gather.totalByte ? ((scope.row['totalByte'] / gather.totalByte) * 100).toFixed(4) + "%" : ((scope.row['trafficTotal'] / gather.trafficTotal) * 100).toFixed(4) + "%"
            }}
          </template>
          <template
            v-else-if="item.field==='totalProportion'"
            #default="scope"
          >
            {{ (scope.row[item.field] * 100).toFixed(4) + "%" }}
          </template>
          <template
            v-else-if="item.field==='diskUseMap'"
            #default="{row}"
          >
            <div
              v-for="(value, key) in row.diskUseMap"
              :key="key"
              class="space-x-2"
            >
              <el-tag type="success">
                {{ key }}: {{ (value[1] / value[0] * 100).toFixed(2) }}%
              </el-tag>
              <el-tag type="primary">
                {{ formatSize(value[1], {units: ['KB', 'MB', 'GB', 'TB'], base: 1024}) }} / {{ formatSize(value[0], {units: ['KB', 'MB', 'GB', 'TB'], base: 1024}) }}
              </el-tag>
            </div>
          </template>
          <template
            v-else-if="item.field==='processMap'"
            #default="{row}"
          >
            <div
              v-for="(value, key) in row.processMap"
              :key="key"
            >
              <el-tag
                v-if="value!=0"
                type="success"
              >
                {{ key }}
              </el-tag>
              <el-tag
                v-else
                type="danger"
              >
                {{ key }}
              </el-tag>
            </div>
          </template>
        </el-table-column>
        <el-table-column
          v-if="props.operation"
          prop="option"
          label="操作"
          fixed="right"
          width="100"
        >
          <template #default="scope">
            <el-button
              link
              type="success"
              icon="Edit"
              title="编辑"
              size="large"
              @click="handleEdit(scope.row)"
            />
            <el-button
              link
              type="danger"
              icon="Delete"
              title="删除"
              size="large"
              @click="handleDelete(scope.row)"
            />
          </template>
        </el-table-column>
      </el-table>
      <div class="flex justify-between">
        <!--   分页   -->
        <Pagination
          v-if="total > 0"
          ref="paginationParam"
          :total="total"
          :page-size="pageSize"
          :refresh="getTableData ? getTableData : refresh"
        />
        <!--   汇总信息   -->
        <span
          v-if="Object.values(gather).length > 0"
          class="mt-4"
        >
          <el-text>
            <b>汇总：</b>
            上行平均：<span class="text-blue-500 font-bold">{{
              formatByte((gather.avgUpBps ? gather.avgUpBps : gather.avgSpeedUp), "bps")
            }}</span>；
            下行平均：<span class="text-blue-500 font-bold">{{
              formatByte((gather.avgDnBps ? gather.avgDnBps : gather.avgSpeedDn), "bps")
            }}</span>；
            上行总量：<span
              class="text-blue-500 font-bold"
            >{{ formatByte(gather.upByte ? gather.upByte : gather.trafficUp) }}</span>；
            下行总量：<span
              class="text-blue-500 font-bold"
            >{{ formatByte(gather.dnByte ? gather.dnByte : gather.trafficDn) }}</span>；
            上下行总量：<span class="text-blue-500 font-bold">{{
              formatByte(gather.totalByte ? gather.totalByte : gather.trafficTotal)
            }}</span>
          </el-text>
        </span>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import {computed, nextTick, onMounted, ref} from 'vue'
import {Link} from '@element-plus/icons-vue'
// 公共请求方法
import {formatTimeToStr} from '@/utils/date'
import {formatByte, formatIsOversea, formatSize} from '@/utils/format'
// 引入分页组件
import Pagination from "./pagination.vue"
// 获取父组件传值
const props = defineProps({
  tableParams: {
    type: Object,
    default: {}
  },
  tableColumns: {
    type: Array,
    default: []
  },
  tableData: {
    type: Array,
    default: []
  },
  operation: {
    type: Boolean,
    default: true
  },
  refresh: {
    type: Function,
    default: null
  },
  autoHeight: {
    type: Boolean,
    default: false
  },
  autoSearch: {
    type: Boolean,
    default: true
  },
  pageSize: {
    type: Number,
    default: 20
  },
  level: {
    type: Number,
    default: 1
  },
  expand: {
    type: String,
    default: ""
  },
  gatherData: {
    type: Object,
    default: {}
  },
  moduleName: {
    type: String,
    default: ""
  },
  moduleApi: {
    type: String,
    default: ""
  },
})
const gather = ref({})
// 动态引入 api 接口
const moduleName = ref(props.moduleName)
const moduleApi = ref(props.moduleApi)
const moduleMap = {
  device: () => import('@/api/device'),
  link: () => import('@/api/link'),
  user: () => import('@/api/userMng'),
  bypass: () => import('@/api/bypass'),
  traffic: () => import('@/api/traffic'),
  policy: () => import('@/api/policy'),
}
// 声明 列表 需要的常量
const loading = ref(false)
const searchParams = ref(props.tableParams) // 查询条件
const columns = ref(props.tableColumns) // 初始化表格结构体
const data = ref([]) //初始化表格数据
const expandColumns = ref(props.tableColumns) // 展开表格字段
const expandData = ref([]) // 展开表格
const selectionRow = ref({}) // 选中行
const checkedRows = ref([]) // 选中行
// 定义分页需要的常量
const paginationParam = ref(null)
const total = ref(0)

// 获取数据
const getTableData = (type = '') => {
  loading.value = true
  if (moduleName.value && moduleApi.value) {
    // 基本检索条件
    const searchInfo = {
      page: paginationParam.value ? paginationParam.value.page : 1,
      limit: paginationParam.value ? paginationParam.value.pageSize : 20,
    }
    moduleMap[moduleName.value]().then(module => {

      let getApi = module[moduleApi.value]
      if (type === '' && moduleApi.value === "userActionDataApi") {
        searchParams.value.dstIp = null
        searchParams.value.appId = null
      } else if (type === "detail") {
        getApi = module["userActionDetailApi"]
      }

      console.log('searchParams', searchParams.value)

      getApi({...searchParams.value, ...searchInfo}).then(response => {
        if (response.code === 0 && response.data) {
          if (type === '') {
            if (response.data.list === null) {
              data.value = []
            } else {
              if (response.data.list) {
                data.value = response.data.list
              } else if (response.data.rows) {
                data.value = response.data.rows
              } else if (response.data) {
                data.value = response.data
              } else {
                data.value = []
              }
            }
            total.value = (response.data.list ? response.data.totalCount : 0) ?? 0
            gather.value = (response.data.gather ? response.data.gather : props.gatherData)
          } else if (type === 'detail') {
            expandData.value = (response.data.list ? response.data.list : response.data) ?? []
          }
        }
      })
    }).catch(err => {
      console.log("请求模块发生错误", err)
    })
  } else {
    data.value = props.tableData ?? []
    total.value = props.tableData.length ?? 0
  }
  loading.value = false
}
// 定义单击行事件
const handleRowClick = (val) => {
  selectionRow.value = val
}
// 定义 选中行 的方法
const handleSelectionChange = (val) => {
  selectionRow.value = val[0] ?? {}
  checkedRows.value = val ?? []
}
// 定义行展开
const expandKeys = ref([])
const handleExpandChange = async (row, expandedRows) => {
  if (moduleApi.value === "userActionDataApi") {
    if (searchParams.value.dataType === "dst_ip") {
      searchParams.value.dstIp = row.name
      expandColumns.value = [
        {name: "应用名称", field: "name"},
        {name: "域名", field: "host"},
        {name: "上行流量", field: "upByte", sortable: true},
        {name: "下行流量", field: "dnByte", sortable: true},
        {name: "总流量", field: "totalByte", sortable: true},
      ]
    } else if (searchParams.value.dataType === "app_id") {
      searchParams.value.appId = row.appId
      expandColumns.value = [
        {name: "目的IP", field: "name"},
        {name: "域名", field: "host"},
        {name: "上行流量", field: "upByte", sortable: true},
        {name: "下行流量", field: "dnByte", sortable: true},
        {name: "总流量", field: "totalByte", sortable: true},
      ]
    }
    await getTableData('detail')
  }

  if (expandedRows.length > 1) {
    expandKeys.value = [row.name]
  } else {
    expandKeys.value = expandedRows.map(row => row.name)
  }
}
// 数据格式化方法
const colFormatter = (row, column) => {
  if (column.property.indexOf("time") && column.property === "update_at") {
    return formatTimeToStr(row[column.property])
  } else if (column.property.indexOf("Bps") > -1 || column.property.indexOf("Speed") > -1 || column.property.indexOf("speed") > -1) {
    return formatByte(row[column.property], "bps")
  } else if (column.property.indexOf("Byte") > -1 || column.property.indexOf("traffic") > -1) {
    return formatByte(row[column.property], "B")
  } else if (column.property === 'isOversea') {
    return formatIsOversea(row[column.property])
  } else if (column.property.indexOf("flow_rate") > -1) {
    return formatByte(row[column.property]*1000000, "bps")
  } else {
    return row[column.property] !== "" ? row[column.property] : "-"
  }
}
// 编辑
const handleEdit = (row) => {
  emits('handleEdit', 'Edittttttttttt ================= ', row)
}
// 删除
const handleDelete = (row) => {
  emits('handleDelete', 'Deleteeeeeeeee ================= ', row)
}
// 点击链接跳转
const handleLinkClick = (index, row) => {
  emits('handleLinkClick', 'LinkClick ================= ', index, row)
}
// 定义表格高度
const getHeight = computed(() => {
  return window.innerHeight - (props.autoHeight ? 562 : 250)
})
// 向父组件发射事件
const emits = defineEmits(['handleSendSingle', 'handleEdit', 'handleDelete', 'handleLinkClick'])

onMounted(() => {
  nextTick(() => {
    if (props.autoSearch) {
      getTableData()
    }
  })
})

defineExpose({
  selectionRow,
  checkedRows,
  paginationParam,
  ...([1, 2].includes(props.level) && {data})
})

</script>

<style lang="scss" scoped>

</style>
