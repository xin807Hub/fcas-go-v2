<template>
  <div>
    <el-card class="w-full">
      <div class="inline-flex w-full justify-between">
        <div class="inline-flex gap-2">
          <el-input
            v-model="searchForm.key" placeholder="请输入分流器名称或分流器IP" class="w-96" size="small" clearable
            @clear="fetchTableData"
          />
          <el-button type="primary" class="button" size="small" icon="Search" @click="fetchTableData">
            查询
          </el-button>
        </div>
        <el-button type="primary" class="button" size="small" icon="Plus" @click="addFunc">
          新建
        </el-button>
      </div>
    </el-card>

    <el-card class="h-full mt-2">
      <el-table v-loading="loading" size="small" height="720" :data="tableData.list" highlight-current-row>
        <el-table-column type="index" label="序号" width="60" align="center" />
        <el-table-column prop="bypassName" label="分流器名称" width="380" show-overflow-tooltip />
        <el-table-column prop="olpId" label="Bypass编号" width="150" align="center" show-overflow-tooltip>
          <template #default="{ row }">
            <el-tag type="success" round>
              {{ row.olpId }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="bypassIp" label="分流器IP" width="160" show-overflow-tooltip>
          <template #default="{ row }">
            <el-tag effect="dark">
              {{ row.bypassIp }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="bypassPort" label="分流器端口" width="160" align="center" show-overflow-tooltip>
          <template #default="{ row }">
            <el-tag round>
              {{ row.bypassPort }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="Bypass状态" align="center" width="220">
          <!-- 0:bypass；1-串行 -->
          <template #default="{ row }">
            <el-switch
              v-if="btnAuth.bypassStatus"
              v-model="row.status" :loading="row.loading" :active-value="1" :inactive-value="0" width="80"
              inline-prompt active-text="串行" inactive-text="bypass"
              :before-change="() => handleConfirmChange(row)"
            />
            <div v-else>
              <el-tag v-if="row.status===1" type="success">
                串行
              </el-tag>
              <el-tag v-else type="info">
                bypass
              </el-tag>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="remark" label="备注" show-overflow-tooltip />
        <el-table-column prop="option" label="操作" align="center" fixed="right" width="100">
          <template #default="{ row }">
            <el-button link type="success" icon="Edit" title="编辑" size="large" @click="editFunc(row)" />
            <el-button link type="danger" icon="Delete" title="删除" size="large" @click="deleteFunc(row)" />
          </template>
        </el-table-column>
      </el-table>

      <el-pagination
        v-model:current-page="pagination.page" v-model:page-size="pagination.pageSize" class="mt-2"
        background size="small" :total="tableData.total" :page-sizes="[10, 50, 100]"
        layout="sizes, prev, pager, next, total" @current-change="fetchTableData" @size-change="fetchTableData"
      />
    </el-card>

    <ConfirmDialog ref="confirmDialogRef" />

    <!--  配置   -->
    <el-dialog v-model="configVisible" :title="configTitle" class="w-176 h-auto">
      <el-form ref="formRef" :model="form" :rules="rules">
        <el-form-item v-for="config in configItems" :label="config.name" label-width="120px" :prop="config.field">
          <el-input
            v-model="form[config.field]" :placeholder="`请输入${config.name}`" autocomplete="off" class="w-112"
            :type="config.type === 'textarea' ? 'textarea' : 'input'"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="configVisible = !configVisible">取消</el-button>
          <el-button type="primary" :loading="loading" @click="handleSubmit">提交</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {ref, onMounted} from "vue"
import {ElMessage} from 'element-plus'
import {isInteger, validateIP, validatePort} from '@/utils/validate'
import {createBypassApi, updateBypassApi, deleteBypassApi, bypassListApi, setStatusApi} from "@/api/bypass"
import ConfirmDialog from "@/view/deviceUser/ConfirmDialog.vue"
import {useBtnAuth} from "@/utils/btnAuth";

const btnAuth = useBtnAuth()

// 列表查询逻辑
const searchForm = ref({
  key: null,
})
const tableData = ref({
  list: [],
  total: 0,
})
const pagination = ref({
  page: 1,
  limit: 20,
})

const loading = ref(false)
const fetchTableData = async () => {
  const params = {
    ...searchForm.value,
    ...pagination.value,
  }

  console.log('params', params)

  try {
    loading.value = true
    const resp = await bypassListApi(params)
    console.log('resp', resp)
    // 添加用于展开列的id字段
    tableData.value.list = resp?.data?.list || []
    tableData.value.total = resp?.data?.totalCount || 0

    console.log('tableData', tableData.value)
  } finally {
    loading.value = false
  }
}

// 配置逻辑
const formRef = ref(null)
const form = ref({
  id: null
})
const configVisible = ref(false)
const configTitle = ref("新增")
const configItems = ref([
  {name: "Bypass编号", field: "olpId", type: "number"},
  {name: "分流器IP", field: "bypassIp", type: "string"},
  {name: "分流器端口", field: "bypassPort", type: "string"},
  {name: "分流器名称", field: "bypassName", type: "string"},
  {name: "备注", field: "remark", type: "textarea"},
])
const rules = ref({
  olpId: [
    {required: true, message: '请输入Bypass编号', trigger: 'blur'},
    {message: 'Bypass编号不合法', validator: isInteger, trigger: 'blur'}
  ],
  bypassIp: [
    {required: true, message: '请输入分流器IP', trigger: 'blur'},
    {message: 'IP地址不合法', validator: validateIP, trigger: 'blur'}
  ],
  bypassPort: [
    {required: true, message: '请输入分流器端口', trigger: 'blur'},
    {message: '端口不合法', validator: validatePort, trigger: 'blur'}
  ],
})

const initForm = () => {
  const temp_form = {}
  configItems.value.forEach(item => {
    temp_form[item.field] = null
  })
  form.value = {...form.value, ...temp_form}
}


const addFunc = () => {
  configVisible.value = true
  configTitle.value = "新建"
  initForm()
}

const editFunc = (row) => {
  configVisible.value = true
  configTitle.value = "编辑"
  configItems.value.forEach(item => {
    form.value.id = row.id
    form.value[item.field] = row[item.field]
  })
}

const deleteFunc = async (row) => {
  const res = await deleteBypassApi([row.id])
  if (res.code === 0) {
    ElMessage.success(`删除成功！`)
    await fetchTableData()
  }
}

const handleSubmit = async () => {
  loading.value = true
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
        ElMessage.success(`操作成功！`)
        configVisible.value = false
        await fetchTableData()
      }
    }
  })
  loading.value = false
}

// ------------------------------ 状态切换逻辑 ----------------------------------

const confirmDialogRef = ref(null)

const handleConfirmChange = async (row) => {
  try {

    const bypassPassword = await confirmDialogRef.value.show(row.bypassName)

    handleStatusChange(row, bypassPassword)

    return true
  } catch (e) {
    return false;
  }
};

const handleStatusChange = async (row, bypassPassword) => {
  row.loading = true
  try {
    const params = {
      bypass_info: {...row},
      bypass_password: bypassPassword,
    }
    const resp = await setStatusApi(params)
    if (resp.code === 0) {
      ElMessage.success('切换成功')
    } else {
      // 若请求失败，则恢复原来的状态
      row.status = Number(!row.status)
    }
  } finally {
    row.loading = false
  }
}

onMounted(() => {
  fetchTableData()
})
</script>

<style scoped lang="scss"></style>
