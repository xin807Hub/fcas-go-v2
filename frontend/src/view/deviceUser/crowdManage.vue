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
      module-name="user"
      module-api="userCrowdApi"
      @handleEdit="editFunc"
      @handleDelete="deleteFunc"
    />
    <!--  配置   -->
    <el-dialog
      v-model="configVisible"
      :title="configTitle"
      class="w-198 h-auto"
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
            v-if="config.type !== 'transfer'"
            v-model="form[config.field]"
            :placeholder="`请输入${config.name}`"
            autocomplete="off"
            class="w-112"
            :type="config.type === 'textarea' ? 'textarea' : 'input'"
          />
          <el-transfer
            v-else
            v-model="bindData"
            filterable
            :titles="['用户', '关联用户']"
            :format="{
              noChecked: '${total}',
              hasChecked: '${checked}/${total}'
            }"
            :data="data"
            @change="handleChange"
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
import {isInteger} from '@/utils/validate'
// 引入公共组件
import ToolBar from "@/components/commonTable/toolbar.vue"
import CommonTable from "@/components/commonTable/index.vue"
// 引入接口
import { createUserCrowdApi, updateUserCrowdApi, deleteUserCrowdApi, userListApi } from "@/api/userMng"
import {useUserTreeStore} from "@/pinia/modules/useUserTree";

const userTreeStore = useUserTreeStore()

// 定义工具栏配置
const toolbarParam = ref(null)
const datetimePicker = ref(false)
const searchEl = ref([
  { name: "用户群名称", field: "key" }
])
// 列表查询逻辑
const loadKey = ref(0)
const tableParam = ref(null)
const tableColumns = ref([
  { name: "用户群名称", field: "crowdName" },
  { name: "关联用户", field: "users" },
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
  { name: "用户群名称", field: "crowdName", type: "string" },
  { name: "备注", field: "remark", type: "textarea" },
  { name: "关联用户", field: "userIds", type: "transfer" },
])
const rules = ref({
  crowdName: [{ required: true, message: '请输入链路名称', trigger: 'blur' }],
})
// 穿梭框数据
const bindData = ref([])
const data = ref([])
const handleChange = (value) => {
  form.value.userIds = value
}
const getUserData = async () => {
  // if (localStorage.getItem('userData')) {
  //   data.value = JSON.parse(localStorage.getItem('userData'))
  // } else {
    const res = await userListApi({
      page: 1,
      limit: Number.MAX_SAFE_INTEGER,
      key: ""
    })
    if (res.code === 0) {
      const temp_data = res.data.list.map(item => {
        return {
          key: item.id,
          label: item.userName
        }
      })
      // localStorage.setItem('userData',JSON.stringify(temp_data))
      data.value = temp_data
    }
  // }
}

const initForm = () => {
  const temp_form = {}
  configItems.value.forEach(item => {
    temp_form[item.field] = null
    temp_form.userIds = []
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

  bindData.value = []
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

  bindData.value = row.users.map(item => item.id)
}

const deleteFunc = async (str, row) => {
  const res = await deleteUserCrowdApi([row.id])
  if (res.code === 0) {
    ElMessage.success( `删除成功！`)
    loadKey.value += 1
    await userTreeStore.syncTree()
  }
}

const handleSubmit = async () => {
  formRef.value.validate(async valid => {
    if (valid) {
      let res = {}
      if (configTitle.value === "新建") {
        res = await createUserCrowdApi(form.value)
      } else {
        res = await updateUserCrowdApi(form.value)
      }
      if (res.code === 0) {
        ElMessage.success( `操作成功！`)
        configVisible.value = false
        loadKey.value += 1
        await userTreeStore.syncTree()
      }
    }
  })
}

onMounted(() => {
  nextTick(() => {
    getUserData()
  })
})
</script>

<style scoped lang="scss">

</style>
