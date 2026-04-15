<!--区域选择器-->
<script setup>
import {ref, computed, onBeforeMount} from 'vue'

const props = defineProps({
  selected: {
    type: String,
    default: null,
  },
})

const emit = defineEmits(['update:selected'])

const selected = computed({
  get: () => props.selected,
  set: (value) => emit('update:selected', value)
})

const options = ref([])

onBeforeMount(async () => {
  // 获取运营商
  options.value = [{label: "上海", value: "上海"}]
})
</script>

<template>
  <div>
    <el-select
      v-model="selected"
      placeholder="请选择区域"
      size="small"
      clearable
      filterable
      style="width: 120px"
    >
      <el-option
        v-for="item in options"
        :key="item.value"
        :label="item.label"
        :value="item.value"
      />
    </el-select>
  </div>
</template>

<style scoped lang="scss">

</style>