<!--运营商选择器-->
<script setup>
import {ref, computed, onBeforeMount} from 'vue'
import {useIspStore} from "@/pinia/modules/useIsp";

const ispStore = useIspStore()

const props = defineProps({
  selected: {
    type: Array,
    default: () => []
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
  options.value = await ispStore.getOptions()
})
</script>

<template>
  <div>
    <el-select
      v-model="selected"
      placeholder="请选择运营商"
      size="small"
      multiple
      collapse-tags
      collapse-tags-tooltip
      :max-collapse-tags="1"
      clearable
      filterable
      style="width: 160px"
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