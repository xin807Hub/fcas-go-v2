<template>
  <div>
    <el-card class="w-full h-auto">
      <div class="flex flex-wrap justify-between">
        <div class="flex">
          <!--    时间粒度    -->
          <el-dropdown
            v-if="particleInfo.show"
            class="mr-2"
            @command="handleParticle"
          >
            <el-button size="small" type="primary">
              <el-icon class="el-icon--left">
                <Coin />
              </el-icon>
              <span class="relative ml-1">时间粒度: ({{ particleText }})</span>
              <el-icon class="el-icon--right">
                <ArrowDown />
              </el-icon>
            </el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item
                  v-for="item of particleInfo.options"
                  :key="item.value"
                  :command="item.value"
                >
                  {{ item.label }}
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
          <!--    时间段    -->
          <el-dropdown
            v-if="intervalInfo.show"
            class="mr-2"
            @command="handleTimeInterval"
          >
            <el-button size="small" type="primary">
              <el-icon class="el-icon--left">
                <Connection />
              </el-icon>
              <span class="relative ml-1">时间段: ({{ intervalText }})</span>
              <el-icon class="el-icon--right">
                <ArrowDown />
              </el-icon>
            </el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item
                  v-for="item of intervalInfo.options"
                  :key="item.value"
                  :command="item.value"
                >
                  {{ item.label }}
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
          <!--    自定义的拓展条件    -->
          <el-dropdown
            v-if="extendInfo.show"
            class="mr-2"
            @command="handleExtend"
          >
            <el-button size="small" type="primary">
              <el-icon class="el-icon--left">
                <Coin />
              </el-icon>
              <span class="relative ml-1">{{ extendInfo.title }}: ({{ extendText }})</span>
              <el-icon class="el-icon--right">
                <ArrowDown />
              </el-icon>
            </el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item
                  v-for="item of extendInfo.options"
                  :key="item.value"
                  :command="item.value"
                >
                  {{ item.label }}
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
          <!--    日期 & 时间选择器    -->
          <el-date-picker
              v-if="props.datetimePicker"
              v-model="timeRange"
              class="mr-2 w-96"
              data-format="YYYY/MM/DD ddd"
              end-placeholde="请选择结束时间"
              format="YYYY-MM-DD HH:mm:00"
              size="small"
              start-placeholde="请选择开始时间"
              time-format="HH:mm:00"
              type="datetimerange"
              @change="handleChangeTime"
          />
          <!--    搜索条件    -->
          <span v-for="item in props.searchEl" class="mr-2">
            <!--            <span class="font-bold mr-2">{{ item.name }}</span>-->
            <el-select
              v-if="item.options && !item.cascader"
              v-model="form[item.field]"
              :placeholder="`请选择${item.name}`"
              :multiple="item.multiple"
              clearable
              filterable
              collapse-tags
              collapse-tags-tooltip
              size="small"
              :class="item.width ?? 'w-60'"
              @change="handleChange($event, item.field)"
            >
              <el-option
                v-for="opt in item.options"
                :key="opt.value"
                :label="opt.name"
                :value="opt.value"
              />
            </el-select>
            <el-tree-select
              v-else-if="item.treeOptions"
              v-model="form[item.field]"
              :placeholder="`请选择${item.name}`"
              :data="item.treeOptions"
              :render-after-expand="false"
              show-checkbox
              :multiple="item.multiple"
              clearable
              filterable
              collapse-tags
              collapse-tags-tooltip
              size="small"
              :class="item.width ?? 'w-60'"
              :lazy="item.lazy"
              :load="loadNode"
              :props="defaultProps"
              @node-expand="handleNodeExpand"
              @check-change="(node, checked, indeterminate) => handleCheckChange(node, checked, indeterminate, item.field)"
              @clear="handleClear(item.field)"
            />
            <el-input
              v-else
              v-model="form[item.field]"
              :placeholder="`请输入${item.name}`"
              autocomplete="off"
              size="small"
              :class="item.width ?? 'w-96'"
              clearable
            />
          </span>
          <!--   查询   -->
          <el-button type="primary" icon="Search" class="button search-button" size="small" @click="handleSearch">
            查询
          </el-button>
        </div>
        <!--        <div class="mt-0.5">-->
        <!--          <el-button type="info" icon="Refresh" class="button refresh-button" size="small" @click="handleReset">-->
        <!--            重置-->
        <!--          </el-button>-->
        <!--        </div>-->
        <div v-if="props.toolbarType === 'config'" class="mt-0.5">
          <!--    新建、编辑、刷新、启用、禁用事件    -->
          <el-button type="primary" icon="Plus" class="button" size="small" @click="handleAdd">
            新建
          </el-button>
          <el-button v-if="props.editFlag" type="warning" icon="Edit" class="button" size="small" @click="handleEdit">
            编辑
          </el-button>
          <el-button v-if="props.deleteFlag" type="danger" icon="Delete" class="button" size="small" @click="handleDelete">
            删除
          </el-button>
          <el-button v-if="props.refreshFlag" type="success" icon="Refresh" class="button" size="small" @click="handleRefresh">
            刷新
          </el-button>
        </div>
        <div v-if="props.uploadFlag || props.exportFlag">
          <!--    导入、导出    -->
          <el-button v-if="props.uploadFlag" type="success" icon="UploadFilled" class="button" size="small" @click="handleUpload">
            文件上传
          </el-button>
          <el-button v-if="props.exportFlag" type="primary" icon="Promotion" class="button" size="small" @click="handleExport">
            导出
          </el-button>
        </div>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import {computed, onMounted, ref} from 'vue'
import {ElMessage} from 'element-plus'
import {ArrowDown} from "@element-plus/icons-vue";
// 引入工具类
import {formatTimeToStr} from '@/utils/date'

// 获取父组件传值
const props = defineProps({
  toolbarType: { // 工具栏类型
    type: String,
    default: "excute"
  },
  particleInfo: { // 颗粒度
    type: Object,
    default: {}
  },
  intervalInfo: { // 时间段
    type: Object,
    default: {}
  },
  datetimePicker: {
    type: Boolean,
    default: false
  }, // 是否有日期时间选择框
  editFlag: {
    type: Boolean,
    default: false
  }, // 是否能编辑
  deleteFlag: {
    type: Boolean,
    default: false
  }, // 是否能删除
  refreshFlag: {
    type: Boolean,
    default: false
  }, // 是否能刷新
  exportFlag: {
    type: Boolean,
    default: false
  }, // 是否能导出
  uploadFlag: {
    type: Boolean,
    default: false
  }, // 是否能上传
  searchEl: { // 搜索条件
    type: Object,
    default: {}
  }
})
const particleInfo = props.particleInfo ? props.particleInfo : {
  show: false,
  initValue: "",
  initText: "",
  options: []
}
const intervalInfo = props.intervalInfo ? props.intervalInfo : {
  show: false,
  initValue: "",
  initText: "",
  options: []
}
const extendInfo = props.extendInfo ? props.extendInfo : {
  show: false,
  initValue: "",
  initText: "",
  options: []
}
const particleText = ref(particleInfo.initText)
const intervalText = ref(intervalInfo.initText)
const extendText = ref(extendInfo.initText)

// 定义查询表单
const form = ref({
  particle: particleInfo.initValue,
  interval: intervalInfo.initValue,
  extend: extendInfo.initValue,
})
// 定义时间范围变量 -- 默认时间范围是1小时 -- 且以10min向下取整
const now = new Date()
const minutes = now.getMinutes()
const roundMinutes = Math.floor(minutes / 10) * 10
const end = new Date(now)
end.setMinutes(roundMinutes, 0, 0)
const start = new Date(end.getTime() - 60 * 60 * 1000)
const timeRange = ref([formatTimeToStr(start, 'yyyy-MM-dd hh:mm:00'), formatTimeToStr(end, 'yyyy-MM-dd hh:mm:00')])
if (props.datetimePicker) {
  form.value['startTime'] = formatTimeToStr(timeRange.value[0])
  form.value['endTime'] = formatTimeToStr(timeRange.value[1])
}

// 向父组件发射事件
const emits = defineEmits([
  'handleParticle',
  'handleTimeInterval',
  'handleExtend',
  'handleSearch',
  'handleExport',
  'handleUpload',
  'handleReset',
  'handleAdd',
  'handleDelete',
  'handleEdit',
  'handleRefresh',
  'handleChange'
])
// 颗粒度变化
const handleParticle = (value) => {
  particleInfo.options.forEach(item => {
    if (item.value === value) {
      particleText.value = item.label
      form.value.particle = value
    }
  })
  emits('handleParticle', '颗粒度切换,查询参数....', form.value)
}
// 时间段变化
const handleTimeInterval = (value) => {
  const now = new Date()
  const later = new Date()
  let timeInterval = 60 * 60
  if (value === 2) {
    timeInterval = 60 * 60 * 24
  } else if (value === 3) {
    timeInterval = 60 * 60 * 24 * 7
  } else if (value === 4) {
    timeInterval = 60 * 60 * 24 * 30
  }
  later.setTime(later.getTime() - timeInterval * 1000)
  timeRange.value = [later, now]
  form.value['startTime'] = formatTimeToStr(timeRange.value[0])
  form.value['endTime'] = formatTimeToStr(timeRange.value[1])

  intervalInfo.options.forEach(item => {
    if (item.value === value) {
      intervalText.value = item.label
      form.value.interval = value
    }
  })
  emits('handleTimeInterval', '时间段切换,查询参数....', form.value)
}
// 额外条件变化
const handleExtend = (value) => {
  extendInfo.options.forEach(item => {
    if (item.value === value) {
      extendText.value = item.label
      form.value.extend = value
    }
  })
  emits('handleExtend', '其它条件切换,查询参数....', form.value)
}
// 时间范围变化
const handleChangeTime = (value) => {
  const startTimestamp = new Date(value[0]).getTime()
  const endTimestamp = new Date(value[1]).getTime()
  if ((endTimestamp - startTimestamp) < 0) {
    ElMessage({
      message: '开始时间不能晚于结束时间',
      type: 'warning',
    })
    timeRange.value = [formatTimeToStr(later), formatTimeToStr(now)]
  }
  // form.value['timeRange'] = timeRange.value
  form.value['startTime'] = formatTimeToStr(timeRange.value[0])
  form.value['endTime'] = formatTimeToStr(timeRange.value[1])
}
// 查询
const handleSearch = () => {
  form.value.appIdList = appIdList.value
  emits('handleSearch', '搜索参数...', form.value)
}
// 导出
const handleExport = () => {
  emits('handleExport', '导出参数...', form.value)
}
// 重置
const handleReset = () => {
  props.searchEl.forEach((item) => {
    form.value[item.field] = ((item.options || item.treeOptions) && item.multiple) ? [] : null
  })
  emits('handleReset', '重置参数...', form.value)
}
// 新增
const handleAdd = () => {
  emits('handleAdd', '新增...')
}
// 编辑
const handleEdit = () => {
  emits('handleEdit', '编辑...')
}
// 删除
const handleDelete = () => {
  emits('handleDelete', '删除...')
}
// 导入
const handleUpload = () => {
  emits('handleUpload', '导入...')
}
// 刷新
const handleRefresh = () => {
  emits('handleRefresh', '导出...')
}
// 下拉框选择事件
const handleChange = (val, field) => {
  if (field === "app_type_id" && form.value.app_id !== null) {
    form.value.app_id = null
  }
  if (field === "appTypeId" && form.value.appId !== null) {
    form.value.appId = null
  }
  emits('handleChange', '选择器选择...', val, field)
}
// 树形选择器懒加载
const loadNode = (node, resolve) => {
  if (node.level === 0) {
    resolve(node.data)
  } else {
    resolve(node.data.children || [])
  }
}

const defaultProps = computed(() => ({
  value: 'id',
  label: 'name',
  children: 'children',
  isLeaf: (data) => !data.children || data.children.length === 0
}))

const appIdList = ref([])
const handleNodeExpand = (node) => {
  if (node.name === "全部") {
    appIdList.value = []
  }
}
const handleCheckChange = (node, checked, indeterminate, field) => {
  if (field === "appIdList" && node.value > 0) {
    if (node.name !== "全部") {
      const appIds = node.children.reduce((arr, curr) => {
        arr.push(curr.value)
        return arr
      }, [])
      if (checked === true) {
        appIdList.value = [...appIdList.value, ...appIds]
      } else {
        appIdList.value = appIdList.value.filter(item => !appIds.includes(item))
      }
    } else {
      appIdList.value = []
      node.children.forEach(item => {
        const temp = item.children.reduce((arr, curr) => {
          arr.push(curr.value)
          return arr
        }, [])
        appIdList.value = [...appIdList.value, ...temp]
      })
    }
    // console.log(appIdList.value.length)
  }
}

const handleClear = (field) => {
  if (field === "appIdList") {
    appIdList.value = []
  } else {
    form.value[field] = null
  }
}


onMounted(() => {
  handleReset()
})

defineExpose({
  form,
  timeRange
})
</script>

<style lang="scss" scoped>
</style>
