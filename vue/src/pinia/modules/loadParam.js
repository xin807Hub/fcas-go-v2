import { defineStore } from 'pinia'
import { ref } from 'vue'
export const useLoadParamStore = defineStore('param', () => {
  const param = ref({})

  const setParam = (value) => {
    param.value = { ...param.value, ...value }
  }

  return {
    param,
    setParam
  }
})
