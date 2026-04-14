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
      module-name="bypass"
      module-api="bypassListApi"
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
import {isInteger, validateIP, validatePort} from '@/utils/validate'
// 引入公共组件
import ToolBar from "@/components/commonTable/toolbar.vue"
import CommonTable from "@/components/commonTable/index.vue"
// 引入接口
import { createBypassApi, updateBypassApi, deleteBypassApi } from "@/api/bypass"
// 定义工具栏配置
const toolbarParam = ref(null)
const datetimePicker = ref(false)
const searchEl = ref([
  { name: "分流器名称或者分流器IP", field: "key" }
])
// 列表查询逻辑
const loadKey = ref(0)
const tableParam = ref(null)
const tableColumns = ref([
  { name: "分流器名称", field: "bypassName" },
  { name: "Bypass编号", field: "olpId" },
  { name: "分流器IP", field: "bypassIp" },
  { name: "分流器端口", field: "bypassPort" },
  { name: "Bypass状态", field: "status" },
  { name: "备注", field: "remark" },
])
// 配置逻辑
const formRef = ref(null)
const form = ref({
  id: null
})
const configVisible = ref(false)
const configTitle = ref("新增")
const configItems = ref([
  { name: "Bypass编号", field: "olpId", type: "number" },
  { name: "分流器IP", field: "bypassIp", type: "string" },
  { name: "分流器端口", field: "bypassPort", type: "string" },
  { name: "分流器名称", field: "bypassName", type: "string" },
  { name: "备注", field: "remark", type: "textarea" },
])
const rules = ref({
  olpId: [
    { required: true, message: '请输入Bypass编号', trigger: 'blur' },
    { message: 'Bypass编号不合法', validator: isInteger, trigger: 'blur' }
  ],
  bypassIp: [
    { required: true, message: '请输入分流器IP', trigger: 'blur' },
    { message: 'IP地址不合法', validator: validateIP, trigger: 'blur' }
  ],
  bypassPort: [
    { required: true, message: '请输入分流器端口', trigger: 'blur' },
    { message: '端口不合法', validator: validatePort, trigger: 'blur' }
  ],
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
  const res = await deleteBypassApi([row.id])
  if (res.code === 0) {
    ElMessage.success( `删除成功！`)
    loadKey.value += 1
  }
}

const handleSubmit = async () => {
  formRef.value.validate(async valid => {
    if (valid) {
      form.value.olpId = parseInt(form.value.olpId)
      form.value.bypassPort = parseInt(form.value.bypassPort)
      let res = {}
      if (configTitle.value === "新建") {
        res = await createBypassApi(form.value)
      } else {
        res = await updateBypassApi(form.value)
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
