<!--链路选择器-->
<script setup>
import {ref, computed, onBeforeMount} from 'vue'
import {useLinkStoreV2} from "@/pinia/modules/useLink";

const linkStore = useLinkStoreV2()

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
  options.value = await linkStore.getOptions()
})
</script>

<template>
  <div>
    <el-select
      v-model="selected"
      placeholder="请选择链路"
      size="small"
      multiple
      collapse-tags
      collapse-tags-tooltip
      :max-collapse-tags="1"
      clearable
      filterable
      style="width: 220px"
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