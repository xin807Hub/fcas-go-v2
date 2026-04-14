<template>
  <div>
    <!--    工具栏    -->
    <ToolBar
      ref="toolbarParam"
      :search-el="searchEl"
      :datetime-picker="datetimePicker"
      toolbar-type="config"
      @handleSearch="searchFunc"
      @handleAdd="addFunc"
      @handleRefresh="refreshFunc"
    />
    <!--    设备列表    -->
    <CommonTable
      v-if="toolbarParam!==null"
      :key="loadKey"
      ref="tableParam"
      :table-params="toolbarParam.form"
      :table-columns="tableColumns"
      :operation="true"
      module-name="device"
      module-api="deviceListApi"
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
          v-for="config in configItems"
          :label="config.name"
          label-width="120px"
          :prop="config.field"
        >
          <el-input
            v-model="form[config.field]"
            :placeholder="`请输入${config.name}`"
            autocomplete="off"
            class="w-112"
            :type="config.type === 'textarea' ? 'textarea' : 'input'"
          />
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
import {ref, onMounted, nextTick} from "vue"
import { ElMessage } from 'element-plus'
// 引入验证类
import {validateIP, validatePort} from '@/utils/validate'
// 引入公共组件
import ToolBar from "@/components/commonTable/toolbar.vue"
import CommonTable from "@/components/commonTable/index.vue"
// 引入接口
import { createDeviceApi, updateDeviceApi, deleteDeviceApi } from "@/api/device"
// 定义工具栏配置
const toolbarParam = ref(null)
const datetimePicker = ref(false)
const searchEl = ref([
  { name: "设备名称或者设备IP", field: "key" }
])
// 列表查询逻辑
const loadKey = ref(0)
const tableParam = ref(null)
const tableColumns = ref([
  { name: "设备名称", field: "deviceName" },
  { name: "设备IP", field: "deviceIp" },
  { name: "磁盘使用情况", field: "diskUseMap" },
  { name: "服务状态", field: "processMap" },
  { name: "连接状态", field: "linkState" },
])
// 配置逻辑
const formRef = ref(null)
const form = ref({
  id: null
})
const configVisible = ref(false)
const configTitle = ref("新增")
const configItems = ref([
  { name: "设备名称", field: "deviceName", type: "string" },
  { name: "设备IP", field: "deviceIp", type: "string" },
  { name: "SNMP读写团体名", field: "snmpName", type: "string" },
  { name: "UDP端口号", field: "udpPort", type: "number" },
  { name: "设备备注", field: "remark", type: "textarea" },
])
const rules = ref({
  deviceName: [{ required: true, message: '请输入设备名称', trigger: 'blur' }],
  deviceIp: [
    { required: true, message: '请输入设备IP', trigger: 'blur' },
    { message: 'IP地址不合法', validator: validateIP, trigger: 'blur' }
  ],
  udpPort: [{ message: '端口不合法', validator: validatePort, trigger: 'blur' }],
})

const initForm = () => {
  const temp_form = {}
  configItems.value.forEach(item => {
    temp_form[item.field] = null
  })
  form.value = {...form.value, ...temp_form}
}

const searchFunc = () => { // 搜索
  loadKey.value += 1
}

const addFunc = () => {
  configVisible.value = true
  configTitle.value = "新建"
  initForm()
}

const refreshFunc = () => {
  loadKey.value += 1
}

const editFunc = (str, row) => {
  configVisible.value = true
  configTitle.value = "编辑"
  configItems.value.forEach(item => {
    form.value.id = row.id
    form.value[item.field] = row[item.field]
  })
}

const deleteFunc = async (str, row) => {
  const res = await deleteDeviceApi([row.id])
  if (res.code === 0) {
    ElMessage.success( `删除成功！`)
    loadKey.value += 1
  }
}

const handleSubmit = async () => {
  formRef.value.validate(async valid => {
    if (valid) {
      form.value.udpPort = parseInt(form.value.udpPort)
      let res = {}
      if (configTitle.value === "新建") {
        res = await createDeviceApi(form.value)
      } else {
        res = await updateDeviceApi(form.value)
      }
      if (res.code === 0) {
        ElMessage.success( `操作成功！`)

        configVisible.value = false
        loadKey.value += 1
      }
    }
  })
}

onMounted(() => {
  nextTick(() => {
  })
})
</script>

<style scoped lang="scss">

</style>
