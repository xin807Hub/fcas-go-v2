<script setup>
import { onMounted, ref, watch } from "vue";
import { ElMessage } from 'element-plus'
import dayjs from "dayjs";
import { useVModels } from "@vueuse/core";

// 使用该组件时需传入 v-model:startTime 和 v-model:endTime 两个 props
const props = defineProps({
  startTime: {
    type: [Date, dayjs, String],
    default: null,
  },
  endTime: {
    type: [Date, dayjs, String],
    default: null,
  },
  // 最大时间范围
  maxTimeRange: {
    type: String, // e.g., '5m', '2h', '7d'
    default: null,
  },
  // 快速选择时间范围
  shortcuts: {
    type: Array,
    default: () => [
      {
        text: "最近一天",
        value: () => {
          const end = dayjs();
          return [end.subtract(1, 'day').toDate(), end.toDate()];
        },
      },
      {
        text: "最近一周",
        value: () => {
          const end = dayjs();
          return [end.subtract(7, 'day').toDate(), end.toDate()];
        },
      },
      {
        text: "近一个月",
        value: () => {
          const end = dayjs();
          return [end.subtract(30, 'day').toDate(), end.toDate()];
        },
      },
    ],
  }
})

const { startTime, endTime } = useVModels(props)

// 解析 maxTimeRange 字符串为 {unit, amount}
const unitTextMap = { m: '分钟', h: '小时', d: '天' };
const parseMaxTimeRange = (timeRange) => {
  if (!timeRange) return null;
  const unit = timeRange.slice(-1);
  const amount = parseInt(timeRange);
  const unitToSeconds = {
    'm': 60,    // 分钟转换为秒
    'h': 3600,  // 小时转换为秒
    'd': 86400, // 天转换为秒
    's': 1,     // 秒不变
  };

  const seconds = unitToSeconds[unit];

  if (seconds === undefined) return null; // 如果时间单位不在映射对象中，则返回 null

  return {
    unit: 's',
    amount: amount * seconds,
  };
}
const maxRange = parseMaxTimeRange(props.maxTimeRange)

const selectedDate = ref([]);

watch(selectedDate, (newRange) => {

  if (!newRange || !newRange[0] || !newRange[1] || newRange.length !== 2) {
    selectedDate.value = [null, null]
    return
  }

  let [start, end] = newRange.map(d => dayjs(d));

  if (maxRange && end.diff(start, maxRange.unit) > maxRange.amount) {
    ElMessage.warning(`选择的时间范围不能超过 ${parseInt(props.maxTimeRange)} ${unitTextMap[props.maxTimeRange.slice(-1)]}`);

    end = start.add(maxRange.amount, maxRange.unit);
    selectedDate.value = [start.toDate(), end.toDate()];
  }

  startTime.value = start.format('YYYY-MM-DD HH:mm:ss')
  endTime.value = end.format('YYYY-MM-DD HH:mm:ss')
});


onMounted(() => {
  if (!props.startTime && !props.endTime) {
    // 默认一小时
    selectedDate.value = [dayjs().subtract(1, 'hour'), dayjs()]
    return
  }
  selectedDate.value = [dayjs(props.startTime), dayjs(props.endTime)]
})


</script>

<template>
  <div>
    <el-date-picker v-model="selectedDate" type="datetimerange" :shortcuts="props.shortcuts" range-separator="-"
      size="small" :clearable="false" />
  </div>
</template>

<style scoped>
/* 日期范围选择器宽度 */
:deep(.el-date-editor) {
  width: 320px;
  ;
}

:deep(.el-date-editor .el-range-separator) {
  flex: none;
  ;
}

:deep(.el-date-editor .el-range__icon) {
  //display: none;
}

:deep(.el-date-editor .el-range__close-icon) {}

:deep(.el-range-input) {
  width: 140px;
}
</style>