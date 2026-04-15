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
      module-name="link"
      module-api="linkListApi"
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
          label="链路名称"
          prop="lineName"
          label-width="120px"
        >
          <el-input
            v-model="form.lineName"
            placeholder=""
            class="w-112"
          />
        </el-form-item>
        <el-form-item
          label="链路vlan"
          prop="lineVlan"
          label-width="120px"
        >
          <el-input
            v-model.number="form.lineVlan"
            placeholder=""
            class="w-112"
          />
        </el-form-item>
        <el-form-item
          label="链路备注"
          prop="remark"
          label-width="120px"
        >
          <el-input
            v-model="form.remark"
            placeholder=""
            type="textarea"
            class="w-112"
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
import {ElMessage} from 'element-plus'
// 引入验证类
import {isInteger} from '@/utils/validate'
// 引入公共组件
import ToolBar from "@/components/commonTable/toolbar.vue"
import CommonTable from "@/components/commonTable/index.vue"
// 引入接口
import {createLinkApi, updateLinkApi, deleteLinkApi} from "@/api/link"
import {useLinkStoreV2} from "@/pinia/modules/useLink";

const linkStore = useLinkStoreV2()

// 定义工具栏配置
const toolbarParam = ref(null)
const datetimePicker = ref(false)
const searchEl = ref([
  {name: "链路名称或者链路vlan", field: "key"}
])
// 列表查询逻辑
const loadKey = ref(0)
const tableParam = ref(null)
const tableColumns = ref([
  {name: "链路名称", field: "lineName"},
  {name: "链路vlan", field: "lineVlan"},
  {name: "链路备注", field: "remark"},
])

// 配置逻辑
const configVisible = ref(false)
const configTitle = ref("新增")
const formRef = ref(null)
const defaultForm = () => ({
  id: null,
  lineName: '',
  lineVlan: null,
  remark: ''
})
const form = ref(defaultForm())
const rules = ref({
  lineName: [{required: true, message: '请输入链路名称', trigger: 'blur'}],
  lineVlan: [
    {required: true, message: '请输入链路vlan', trigger: 'blur'},
    {validator: isInteger, message: '链路vlan不合法', trigger: 'blur'},
    {type: "number", min: 0, max: 4095, message: '链路vlan不合法, 范围[0-4095]', trigger: 'blur'},
  ],
})

const searchFunc = () => { // 搜索
  loadKey.value += 1
}

const addFunc = () => {
  configVisible.value = true
  configTitle.value = "新建"
  form.value = defaultForm()
}

const refreshFunc = () => {
  loadKey.value += 1
}

const editFunc = (str, row) => {
  configVisible.value = true
  configTitle.value = "编辑"
  form.value = {...row}
}

const deleteFunc = async (str, row) => {
  const res = await deleteLinkApi([row.id])
  if (res.code === 0) {
    ElMessage.success(`删除成功！`)
    loadKey.value += 1
  }
}

const handleSubmit = async () => {
  formRef.value.validate(async valid => {
    if (valid) {
      form.value.lineVlan = parseInt(form.value.lineVlan)
      let res = {}
      if (configTitle.value === "新建") {
        res = await createLinkApi(form.value)
      } else {
        res = await updateLinkApi(form.value)
      }
      if (res.code === 0) {
        ElMessage.success(`操作成功！`)
        await linkStore.sync()
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
