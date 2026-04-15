<!--运营商选择器-->
<script setup>
import {ref, watch, computed, onBeforeMount} from 'vue'
import {useUserTreeStore} from "@/pinia/modules/useUserTree";

const userTreeStore = useUserTreeStore()

const props = defineProps({
  selected: {
    type: Array,
    default: () => []
  },
  treeType: {
    type: String,
    default: 'userTree'
  },
})

const emit = defineEmits(['update:selected'])

const selected = computed({
  get: () => props.selected,
  set: (value) => emit('update:selected', value)
})

const checkAll = ref(false)
const indeterminate = ref(false)
const value = ref([])
const options = ref([])

const cascaderProps = {
  multiple: true,
}

const getAllValuePaths = computed(() => {
  const result = []
  const queue = options.value.map(
    (node) => ({ node, path: [node.value] })
  )

  while (queue.length > 0) {
    const { node, path } = queue.shift()
    if (node.children?.length) {
      node.children.forEach((child) => {
        queue.push({ node: child, path: [...path, child.value] })
      })
    } else {
      result.push(path)
    }
  }
  
  // console.log('getAllValuePaths', result);
  
  return result
})

const handleChange = (val) => {
  selected.value = val.map(item => item[item.length - 1])
}

const handleCheckAll = (val) => {
  indeterminate.value = false
  value.value = val ? getAllValuePaths.value : []
  handleChange(val ? getAllValuePaths.value : [])
}

watch(value, (val) => {
  if (val.length === 0) {
    checkAll.value = false
    indeterminate.value = false
  } else if (val.length === getAllValuePaths.value.length) {
    checkAll.value = true
    indeterminate.value = false
  } else {
    indeterminate.value = true
  }
})

onBeforeMount(async () => {
  let op

  // 获取运营商
  switch (props.treeType) {
    case 'crowdTree':
      op = await userTreeStore.getCrowdTree()
      break;
    case 'groupTree':
      op = await userTreeStore.getGroupTree()
      break;
    default:
      op = await userTreeStore.getUserTree()
      break;
  }

  options.value = op?.children || []

  // console.log('userTreeOptions', options.value)
})
</script>

<template>
  <div>
    <el-cascader
      v-model="value"
      :options="options"
      :props="cascaderProps"
      :show-all-levels="true"
      clearable
      filterable
      collapse-tags
      collapse-tags-tooltip
      :max-collapse-tags="1"
      placeholder="请选择用户|群|组"
      size="small"
      style="width: 260px;"
      @change="handleChange"
    >
      <template #header>
        <el-checkbox
          v-model="checkAll"
          :indeterminate="indeterminate"
          @change="handleCheckAll"
        >
          All
        </el-checkbox>
      </template>
    </el-cascader>
  </div>
</template>

<style scoped lang="scss">

</style>