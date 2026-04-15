<template>
  <div
      v-if="resolvedData.label"
      class="tag-badge inline-flex h-6 items-center rounded-md px-2
           text-xs text-center font-medium transition-colors duration-200"
      :class="[resolvedStyle.bgClass, resolvedStyle.textClass, resolvedStyle.hoverBgClass]"
  >
    {{ resolvedData.label }}
  </div>
  <!-- 如果只是空值，渲染一个占位符或空 -->
  <span v-else>-</span>
</template>

<script setup>
import { computed } from 'vue';

/**
 * @file EnumTagCell.vue
 * @description 字典/枚举标签渲染器。
 *              接收一个 options 数组，根据当前 value 查找对应的 label 并显示为标签。
 *
 * @usage
 * {
 *   prop: 'status',
 *   label: '状态',
 *   type: 'enum-tag',
 *
    // 核心配置：传递 options 数组
    options: [
      { label: '启用', value: 1, type: 'success' },
      { label: '禁用', value: 0, type: 'danger' },
      { label: '待审核', value: 2 } // 默认 type 为 'info'
    ]
 * }
 */

const props = defineProps({
  value: { type: [String, Number, Boolean], default: null },
  column: { type: Object, default: () => ({}) },
});

// --- 样式映射表 (复用原逻辑) ---
const tagStyles = {
  primary: { bgClass: 'bg-blue-100 dark:bg-blue-800', textClass: 'text-blue-800 dark:text-blue-100', hoverBgClass: 'hover:bg-blue-200 dark:hover:bg-blue-700' },
  success: { bgClass: 'bg-green-100 dark:bg-green-800', textClass: 'text-green-800 dark:text-green-100', hoverBgClass: 'hover:bg-green-200 dark:hover:bg-green-700' },
  warning: { bgClass: 'bg-orange-100 dark:bg-orange-800', textClass: 'text-orange-800 dark:text-orange-100', hoverBgClass: 'hover:bg-orange-200 dark:hover:bg-orange-700' },
  danger:  { bgClass: 'bg-red-100 dark:bg-red-800', textClass: 'text-red-800 dark:text-red-100', hoverBgClass: 'hover:bg-red-200 dark:hover:bg-red-700' },
  info:    { bgClass: 'bg-slate-100 dark:bg-slate-700', textClass: 'text-slate-800 dark:text-slate-200', hoverBgClass: 'hover:bg-slate-200 dark:hover:bg-slate-600' },
};

// --- 核心逻辑: 在数组中查找匹配项 ---
const resolvedData = computed(() => {
  // 1. 获取配置的 options 数组，通常在 column 中
  const options = props.column.options || [];

  // 2. 在数组中查找 value 匹配的项
  // 注意：这里使用 '==' 弱类型比较，兼容 '1' 和 1 的情况。如果需要严格匹配请用 '==='
  const found = options.find(item => item.value == props.value);

  if (found) {
    return {
      label: found.label,
      // 如果 options 里定义了 type (颜色)，则使用它，否则默认 'info'
      type: found.type || 'info'
    };
  }

  // 3. 如果没找到匹配项，优雅降级：直接显示原始值，颜色设为灰色(info)
  return {
    label: String(props.value ?? ''), // 处理 null/undefined
    type: 'info'
  };
});

// --- 样式计算 ---
const resolvedStyle = computed(() => {
  return tagStyles[resolvedData.value.type] || tagStyles.info;
});
</script>
