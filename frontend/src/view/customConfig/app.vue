<template>
  <div>
    <!--    工具栏    -->
    <ToolBar
      ref="toolbarParam"
      :search-el="searchEl"
      :datetime-picker="datetimePicker"
      :upload-flag="true"
      toolbar-type="excute"
      @handleSearch="searchFunc"
      @handleRefresh="refreshFunc"
      @handleUpload="uploadFunc"
    />
    <!--    数据列表    -->
    <CommonTable
      v-if="toolbarParam!==null"
      :key="loadKey"
      ref="tableParam"
      :table-params="toolbarParam.form"
      :table-columns="tableColumns"
      :operation="false"
      :page-size="50"
      module-name="user"
      module-api="appClassifyDataApi"
    />
    <!--  文件上传  -->
    <el-dialog
      v-model="uploadVisible"
      title="文件上传"
      class="w-176 h-auto"
    >
      <el-upload
        v-model:file-list="fileList"
        class="upload-demo"
        drag
        :multiple="false"
        accept=".list"
        :limit="1"
        :headers="{'b-token': userStore.token}"
        :action="`/api/object/appClassify/import`"
        :on-success="uploadOnSuccess"
      >
        <el-icon class="el-icon-upload" size="56">
          <upload-filled class="w-40 h-40" />
        </el-icon>
        <div class="el-upload__text">
          拖动文件到这里 或者 <em>点击上传</em>
        </div>
        <template #tip>
          <div class="el-upload__tip">
            支持上传文件类型为.list, 文件大小不超过20M
          </div>
        </template>
      </el-upload>
    </el-dialog>
  </div>
</template>

<script setup>
import {ref, onMounted, nextTick} from "vue"
import { ElMessage } from 'element-plus'
import { UploadFilled } from '@element-plus/icons-vue'
// 状态管理
import { useUserStore } from '@/pinia'
const userStore = useUserStore()
// 引入公共组件
import ToolBar from "@/components/commonTable/toolbar.vue"
import CommonTable from "@/components/commonTable/index.vue"
// 定义工具栏配置
const toolbarParam = ref(null)
const datetimePicker = ref(false)
const searchEl = ref([
  { name: "应用大类编号", field: "appTypeId", width: 'w-44' },
  { name: "应用小类编号", field: "appId", width: 'w-44' },
])
// 列表查询逻辑
const loadKey = ref(0)
const tableParam = ref(null)
const tableColumns = ref([
  { name: "应用大类编号", field: "appTypeId" },
  { name: "应用大类名称", field: "appTypeName" },
  { name: "应用小类编号", field: "appId" },
  { name: "应用小类名称", field: "appName" },
  // { name: "应用小类业务细分编号", field: "upByte" },
  // { name: "应用小类业务细分名称", field: "upByte" },
])

// 查询
const searchFunc = () => {
  loadKey.value += 1
}
// 刷新
const refreshFunc = () => {
  loadKey.value += 1
}
// 文件上传
const uploadVisible = ref(false)
const fileList = ref([])
const uploadFunc = () => {
  uploadVisible.value = true
}
const uploadOnSuccess = (res, file) => {
  if (res.code === 0) {
    ElMessage.success("文件上传成功!!")
    uploadVisible.value = false
    fileList.value = []
    loadKey.value += 1
  } else {
    ElMessage.error(res.msg)
  }
}

onMounted(() => {
  nextTick(() => {
  })
})
</script>

<style scoped lang="scss">

</style>
