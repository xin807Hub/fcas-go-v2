<template>
  <div>
    <!--    工具栏    -->
    <ToolBar
      ref="toolbarParam"
      :search-el="searchEl"
      :datetime-picker="datetimePicker"
      toolbar-type="config"
      :export-flag="false"
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
      module-name="user"
      module-api="userListApi"
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
            v-if="config.type !== 'radio' && config.type !== 'multiple'"
            v-model="form[config.field]"
            :placeholder="`请输入${config.name}`"
            autocomplete="off"
            class="w-112"
            :type="config.type === 'textarea' ? 'textarea' : 'input'"
          />
          <el-radio-group
            v-else-if="config.type === 'radio'"
            v-model="form[config.field]"
          >
            <el-radio label="是" :value="1">
              是
            </el-radio>
            <el-radio label="否" :value="0">
              否
            </el-radio>
          </el-radio-group>
          <div
            v-for="(item, index) in ipConfs"
            v-else-if="config.type === 'multiple'"
          >
            <el-input
              v-model="item.value"
              placeholder="请输入IP地址/IP地址段"
              autocomplete="off"
              class="w-96 mr-4"
            />
            <el-button
              v-if="index === 0" type="primary" icon="Plus" class="w-8 h-8" round size="small" title="添加"
              @click="handleAddIpConf"
            />
            <el-button
              v-else type="danger" icon="Delete" class="w-8 h-8" round size="small" title="删除"
              @click="handleDeleteIpConf(index)"
            />
          </div>
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
import {nextTick, onMounted, ref} from "vue"
import {ElMessage} from 'element-plus'
// 引入验证类
// 引入公共组件
import ToolBar from "@/components/commonTable/toolbar.vue"
import CommonTable from "@/components/commonTable/index.vue"
// 引入接口
import {createUserApi, deleteUserApi, updateUserApi} from "@/api/userMng"
import {useUserTreeStore} from "@/pinia/modules/useUserTree";

const userTreeStore = useUserTreeStore()

// 定义工具栏配置
const toolbarParam = ref(null)
const datetimePicker = ref(false)
const searchEl = ref([
  {name: "用户名称或者IP地址/IP地址段", field: "key"}
])
// 列表查询逻辑
const loadKey = ref(0)
const tableParam = ref(null)
const tableColumns = ref([
  {name: "用户名称", field: "userName"},
  {name: "IP地址/IP地址段", field: "ipAddress"},
  {name: "IP地址个数", field: "ipAddressNum"},
  {name: "备注", field: "userRemark"},
])
// 配置逻辑
const formRef = ref(null)
const form = ref({
  id: null
})
const configVisible = ref(false)
const configTitle = ref("新增")
const configItems = ref([
  {name: "用户名称", field: "userName", type: "string"},
  {name: "是否为监控用户", field: "userType", type: "radio"},
  {name: "IP地址段", field: "ipAddress", type: "multiple"},
  {name: "用户备注", field: "userRemark", type: "textarea"},
])
const rules = ref({
  userName: [{required: true, message: '请输入用户名称', trigger: 'blur'}],
})
// IP地址段配置逻辑
const ipConfCount = ref(0)
const ipConfs = ref([])
const handleAddIpConf = (event, data) => {
  if (!event) {
    ipConfCount.value = 0
    ipConfs.value = []
  }
  if (!data) {
    ipConfCount.value += 1
  } else {
    ipConfCount.value = data.length
  }
  for (let i = 0; i < ipConfCount.value; i++) {
    if (!ipConfs.value[i]) {
      ipConfs.value[i] = {
        name: "IP地址段" + (i + 1),
        field: "ipAddress" + (i + 1),
        value: data ? data[i] : ""
      }
    }
  }
}
const handleDeleteIpConf = (index) => {
  ipConfs.value.splice(index)
  ipConfCount.value -= 1
}

const initForm = () => {
  const temp_form = {}
  configItems.value.forEach(item => {
    temp_form[item.field] = null
    temp_form.userType = 0
    temp_form.ipAddress = []
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

  ipConfCount.value = 0
  handleAddIpConf()
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

  handleAddIpConf('', row.ipAddress)
}

const deleteFunc = async (str, row) => {
  const res = await deleteUserApi([row.id])
  if (res.code === 0) {
    ElMessage.success(`删除成功！`)
    loadKey.value += 1
    await userTreeStore.init()
  }
}

const handleSubmit = async () => {
  formRef.value.validate(async valid => {
    if (valid) {
      form.value.userType = parseInt(form.value.userType)
      form.value.ipAddress = ipConfs.value.map(item => item.value)
      let res = {}
      if (configTitle.value === "新建") {
        res = await createUserApi(form.value)
      } else {
        res = await updateUserApi(form.value)
      }
      if (res.code === 0) {
        ElMessage.success(`操作成功！`)
        configVisible.value = false
        loadKey.value += 1
        await userTreeStore.init()
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
