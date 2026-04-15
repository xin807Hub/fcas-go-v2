<template>
  <el-pagination
    class="mt-2"
    background
    size="small"
    layout="sizes, prev, pager, next, total"
    :current-page="page"
    :page-size="pageSize"
    :page-sizes="[10, 20, 50, 100]"
    :total="props.total"
    @size-change="handleSizeChange"
    @current-change="handleCurrentChange"
  />
</template>

<script setup>
import { ref, watch } from 'vue'
// import { useRouterStore } from '@/pinia/modules/router'
// const routerStore = useRouterStore()

// 接收父组件传参
const props = defineProps({
  total: {
    type: Number,
    default: 0
  }, // 总数
  pageSize: {
    type: Number,
    default: 20
  },
  refresh: {
    type: Function,
    default: null
  }
})

// 声明 列表 需要的常量
const page = ref(1) // 当前页码
const currentPage = ref(1) // 当前页码
const pageSize = ref(20 ?? props.pageSize) // 页数

// 定义 切换页码 的方法
const handleSizeChange = (val) => {
  pageSize.value = val
  // props.refresh()
}

// 定义 翻页 的方法
const handleCurrentChange = (val) => {
  page.value = val
}

// // 向父组件发射事件
// const emits = defineEmits(['handleSizeChange', 'handleCurrentChange'])

watch(pageSize, () => {
  props.refresh()
})

watch(page, () => {
  props.refresh()
})

defineExpose({
  page,
  pageSize
})
</script>
<style lang="scss" scoped>
</style>
