<!--链路选择器-->
<script setup>
import { ref, computed, onBeforeMount } from 'vue'
import { useAppTypeStore } from "@/pinia/modules/useAppType";

const appTypeStore = useAppTypeStore()

const props = defineProps({
  selected: {
    type: Array,
    default: () => []
  },
  includeChildren: {
    type: Boolean,
    default: true,
  }
})

const emit = defineEmits(['update:selected'])

const selected = computed({
  get: () => props.selected,
  set: (value) => emit('update:selected', value)
})

const options = ref([])

const cascaderProps = {
  value: 'id',
  label: 'name',
  multiple: true,
}

const handleChange = (val) => {
  selected.value = val.map(item => item[item.length - 1])
}

onBeforeMount(async () => {
  if (props.includeChildren) {
    options.value = await appTypeStore.getOptions()
  } else {
    options.value = await appTypeStore.getOptionExcluseChildren()
  }

  // console.log('appTypeOptions', options.value)
})
</script>

<template>
  <div>
    <el-cascader
      v-model="selected"
      :options="options"
      :props="cascaderProps"
      :show-all-levels="false"
      filterable
      clearable
      collapse-tags
      collapse-tags-tooltip
      :max-collapse-tags="1"
      placeholder="请选择应用类型"
      size="small"
      style="width: 260px;"
      @change="handleChange"
    />
  </div>
</template>

<style scoped lang="scss"></style>