<template>
  <div>
    <ToolBar
        ref="toolbarParam"
        :search-el="searchEl"
        :datetime-picker="datetimePicker"
        :delete-flag="deleteFlag"
        :refresh-flag="refreshFlag"
        toolbar-type="config"
        @handleSearch="searchFunc"
        @handleAdd="addFunc"
        @handleDelete="deleteAllFunc"
        @handleChange="changeFunc"
    />
    <!--    列表    -->
    <!--    设备列表    -->
    <CommonTable
        v-if="toolbarParam!==null"
        :key="loadKey"
        ref="tableParam"
        :table-params="toolbarParam.form"
        :table-columns="tableColumns"
        :operation="true"
        module-name="policy"
        module-api="alarmConfigPageApi"
        @handleEdit="editFunc"
        @handleDelete="deleteFunc"
    />
    <!--  配置   -->
    <el-dialog
        v-model="configVisible"
        :title="configTitle"
        class="w-176 h-auto"
    >
      <el-form
          ref="formRef"
          :model="form"
          :rules="rules"
      >
        <el-form-item
            v-for="(config,index) in configItems"
            :key="index"
            :label="config.name"
            label-width="120px"
            :prop="config.field"
        >
          <el-input
              v-if="!config.options"
              v-model="form[config.field]"
              :placeholder="`请输入${config.name}`"
              autocomplete="off"
              class="w-112"
          />
          <el-select
              v-else
              :multiple="config.multiple"
              filterable
              v-model="form[config.field]"
              :placeholder="`请输入${config.name}`"
              class="w-112"
              @change="changeConfigFunc($event, config.field)"
          >
            <el-option
                v-for="item in config.options"
                :key="item.value"
                :label="item.name"
                :value="item.value"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="configVisible = !configVisible">取消</el-button>
          <el-button type="primary" @click="handleSubmit">提交</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>

import ToolBar from "@/components/commonTable/toolbar.vue";
import CommonTable from "@/components/commonTable/index.vue";
import {nextTick, onBeforeMount, onMounted, ref} from "vue";
import {ElMessage} from "element-plus";
import {validateEMail} from "@/utils/validate";
import {dictInfoApi} from "@/api/traffic";
// 状态管理
import useLinkStore from "@/pinia/modules/link.js";
import {deleteAlarmConfigApi, saveOrUpdateAlarmConfigApi} from "@/api/policy";

const loadKey = ref(0)

const LinkStore = useLinkStore()
// 定义工具栏配置
const toolbarParam = ref(null)
const datetimePicker = ref(false)
const deleteFlag = ref(true)
const refreshFlag = ref(false)
const linkOptions = ref([])
const appTypeOptions = ref([])
const appTypeIdTreeOptions = ref([])
// 列表查询逻辑
const tableParam = ref(null)
const tableColumns = ref([
  {name: "名称", field: "name", sortable: true},
  {name: "开始时间", field: "start_time", sortable: true},
  {name: "结束时间", field: "end_time", sortable: true},
  {name: "应用大类", field: "app_type_name", sortable: true},
  {name: "应用小类", field: "app_name", sortable: true},
  {name: "流量上浮比例%", field: "increase_ratio", sortable: true},
  {name: "流量下降比例%", field: "decrease_ratio", sortable: true},
  {name: "邮箱", field: "email", sortable: true},
])

const timeOptions = ref([])
const generateTimeOptions = () => {
  const hours = 24;
  const minutesInterval = 10;
  for (let hour = 0; hour < hours; hour++) {
    for (let minute = 0; minute < 60; minute += minutesInterval) {
      const name = formatTime(hour, minute);
      timeOptions.value.push({name, value: name});
    }
  }
}
const formatTime = (hour, minute) => {
  const formattedHour = String(hour).padStart(2, '0');
  const formattedMinute = String(minute).padStart(2, '0');
  return `${formattedHour}:${formattedMinute}:00`;
}

generateTimeOptions()

// 配置逻辑
const formRef = ref(null)
const form = ref({
  id: null
})
const configVisible = ref(false)
const configTitle = ref("新增")
const configItems = ref([
  {name: "名称", field: "name", type: "string"},
  {name: "开始时间", field: "start_time", type: "string", options: timeOptions.value,multiple: false},
  {name: "结束时间", field: "end_time", type: "string", options: timeOptions.value,multiple: false},
  {name: "应用大类", field: "app_type_id", type: "string", options: appTypeOptions,multiple: false},
  {name: "应用小类", field: "app_id", type: "string", options: [],multiple: false},
  {
    name: "链路",
    field: "link_ids",
    type: "string",
    options: linkOptions,
    multiple: true
  },
  {name: "流量上浮比例%", field: "increase_ratio", type: "number"},
  {name: "流量下降比例%", field: "decrease_ratio", type: "number"},
  {name: "邮箱", field: "email", type: "string"},
])

//验证开始时间是否大于结束时间
const validateStartTime = (rule, value, callback) => {
  const endTime = form.value.end_time
  if (new Date('1970-01-01 ' + value).getTime() >= new Date('1970-01-01 ' + endTime).getTime()) {
    callback(new Error('开始时间大于结束时间'))
  } else {
    callback();
  }
}
//验证结束时间是否大于开始时间
const validateEndTime = (rule, value, callback) => {
  const startTime = form.value.start_time
  if (new Date('1970-01-01 ' + value).getTime() <= new Date('1970-01-01 ' + startTime).getTime()) {
    callback(new Error('开始时间大于结束时间'))
  } else {
    callback();
  }
}

const validateRatio = (rule, value, callback) => {
  const regex = /^(?!\.?$)\d+(\.\d+)?$/
  if (value != null && value!=='') {
    if (!regex.test(value)) {
      callback(new Error('请输入正确的数值'))
    } else {
      callback()
    }
  }else {
    callback()
  }
}

const rules = ref({
  name: [{required: true, message: '名称不能为空', trigger: 'blur'}],
  start_time: [
    {required: true, message: '开始时间不能为空', trigger: 'change'},
    {message: "开始时间大于结束时间", validator: validateStartTime, trigger: 'change'}
  ],
  end_time: [
    {required: true, message: '结束时间不能为空', trigger: 'change'},
    {message: "开始时间大于结束时间", validator: validateEndTime, trigger: 'change'}
  ],
  app_type_id: [{required: true, message: '应用大类ID不能为空', trigger: 'blur'}],
  app_id: [{required: true, message: '应用小类ID不能为空', trigger: 'blur'}],
  link_ids: [{required: true, message: '链路不能为空', trigger: 'blur'}],
  increase_ratio: [
    {required: true, message: '流量上浮比例不能为空', trigger: 'blur'},
    {message: "请输入正确的数值", validator: validateRatio, trigger: 'blur'}
  ],
  decrease_ratio: [
    {validator: validateRatio, trigger: 'blur'}
  ],
  email: [{message: '邮箱不合法', validator: validateEMail, trigger: 'blur'}],
})


const searchEl = ref([
  {name: "名称", field: "name", width: 'w-40'},
  {name: "应用大类", field: "app_type_id", options: [], width: 'w-60'},
  {name: "应用小类", field: "app_id", options: [], width: 'w-48'},
])

const searchFunc = () => { // 搜索
  loadKey.value += 1
}

const isEdit = ref(false)
const addFunc = () => {
  isEdit.value = false
  configVisible.value = true
  configTitle.value = "新建"
  initForm()
}
const editFunc = (str, row) => {
  isEdit.value = true
  if (formRef.value) {
    formRef.value.clearValidate()
  }
  configVisible.value = true
  configTitle.value = "编辑"
  configItems.value.forEach(item => {
    form.value.id = row.id
    form.value[item.field] = row[item.field]
    if (item.name === "应用小类") {
      item.options = row.app_type_id ? appTypeIdTreeOptions.value.filter(item => item.id === row.app_type_id)[0].children.map(item => {
        return {
          value: item.id,
          name: item.name
        }
      }) : []
    }
  })
  form.value.link_ids = row.link_ids.split(",").map(item => parseInt(item))
}
const deleteFunc = async (str, row) => {
  const res = await deleteAlarmConfigApi({ids: [row.id]})
  if (res.code === 0) {
    ElMessage.success(`删除成功！`)
    loadKey.value += 1
  }
}

const deleteAllFunc = async () => {
  const checkedRows = tableParam.value.checkedRows
  let ids = []
  checkedRows.map(item => ids.push(item.id))
  const res = await deleteAlarmConfigApi({ids: ids})
  if (res.code === 0) {
    ElMessage.success(`删除成功！`)
    loadKey.value += 1
  }
}
const changeConfigFunc = (val, field) => {
  if (field === 'app_type_id') {
    form.value.app_id = null
    configItems.value.forEach(item => {
      if (item.name === "应用小类") {
        item.options = val ? appTypeIdTreeOptions.value.filter(item => item.id === val)[0].children.map(item => {
          return {
            value: item.id,
            name: item.name
          }
        }) : []
      }
    })
  }
}
const changeFunc = (str, val, field) => {
  if (localStorage.getItem("appTypeIdTreeOptions")) {
    const res = JSON.parse(localStorage.getItem("appTypeIdTreeOptions"))
    if (field === 'app_type_id') {
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


const handleSubmit = async () => {
  formRef.value.validate(async valid => {
    if (valid) {
      let res = {}
      form.value.increase_ratio = parseFloat(form.value.increase_ratio)
      form.value.decrease_ratio = parseFloat(form.value.decrease_ratio)
        form.value.link_ids = form.value.link_ids.toString()
      res = await saveOrUpdateAlarmConfigApi(form.value)
      if (res.code === 0) {
        ElMessage.success(`操作成功！`)
        configVisible.value = false
        loadKey.value += 1
      }
    }
  })
}

const initForm = () => {
  const temp_form = {}
  configItems.value.forEach(item => {
    temp_form[item.field] = null
    temp_form.id = 0
  })
  form.value = {...form.value, ...temp_form}
  formRef.value.clearValidate()
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
    }
  })
})

onMounted(() => {
  nextTick(() => {
  })
})
</script>

<style scoped lang="scss">

</style>
