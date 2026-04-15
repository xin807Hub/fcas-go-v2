<template>
  <el-dialog :model-value="visible" title="警告" width="540px" :close-on-click-modal="false" @close="onCancel">
    <div class="p-4 mb-4 text-center">
      <p class="text-lg text-red-500 font-semibold mb-3">
        请确认是否切换分流器 <span class="text-blue-700 font-bold">{{ targetName }}</span> 的 Bypass 状态。
      </p>
      <p class="text-base text-red-600">
        该操作可能影响当前网络行为，为防止误操作，请输入Bypass密码进行确认。
      </p>
    </div>

    <div class="mb-4">
      <el-input
        v-model="inputValue" type="password" placeholder="请输入Bypass密码进行确认" clearable
        @input="checkInput"
      />
    </div>

    <template #footer>
      <el-button @click="onCancel">
        取消
      </el-button>
      <el-button type="danger" :disabled="inputValue === ''" @click="onConfirm">
        确认切换
      </el-button>
    </template>
  </el-dialog>
</template>

<script setup>
import {ref} from 'vue'
import {validateBypassPasswordApi} from "@/api/bypass";
import {ElMessage} from "element-plus";


const visible = ref(false)
const inputValue = ref('')
const targetName = ref('')

let resolver = null
let rejecter = null

const show = async (bypassName) => {
  targetName.value = bypassName
  inputValue.value = ''
  visible.value = true


  return new Promise((resolve, reject) => {
    resolver = resolve
    rejecter = reject
  })
}

const onConfirm = async () => {
  const resp = await validateBypassPasswordApi({
    bypass_password: inputValue.value
  })
  if (resp.code === 0) {
    if (resp.data.succ) {
      visible.value = false
      if (resolver) resolver(inputValue.value)
    }else{
      ElMessage.error(resp.data.msg)
    }
  }
}

const onCancel = () => {
  visible.value = false
  if (rejecter) rejecter('cancel')
}


defineExpose({
  show
})
</script>
