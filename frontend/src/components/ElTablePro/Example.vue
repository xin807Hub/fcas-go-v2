<template>
  <ElTablePro
    v-model:current-page="currentPage"
    v-model:page-size="pageSize"
    :columns="columns"
    :data="tableData"
    :total="total"
    :loading="loading"
    show-detail
    :actions-config="{
      width: 140,
    }"
    border
    highlight-current-row
    @page-change="handlePageChange"
    @sort-change="handleSortChange"
    @selection-change="handleSelectionChange"
  >
    <!-- 1. Toolbar 插槽 -->
    <template #toolbar>
      <div class="flex flex-wrap gap-4 justify-between items-center">
        <h2 class="text-lg font-semibold text-slate-700">
          服务监控列表
        </h2>
        <div class="flex items-center gap-2">
          <el-input
            v-model="searchQuery"
            placeholder="模拟搜索..."
            clearable
            class="w-48"
          />
          <el-button
            type="primary"
            :icon="Search"
            @click=""
          >
            搜索
          </el-button>
          <el-button
            :icon="Refresh"
            circle
            @click=""
          />
        </div>
      </div>
    </template>

    <!-- 2. Expand 插槽 -->
    <template #expand="{ row }">
      <div class="p-4">
        <CompoundCell
          :row="row"
          :column="{
            items: [
              { prop: 'sourceIp', label: '源 IP', type: 'ip' },
              { prop: 'region', label: '区域', format: (val) => `📍 ${val}` },
              { prop: 'owner', label: '负责人', type: 'tag', config: { '张三': { type: 'primary' }, '李四': { type: 'success' }, '王五': { type: 'warning' } } },
              { prop: 'fullRequest', label: '完整请求', type: 'json', previewText: '查看完整请求体' }
            ],
            config: {
              labelWidth: '100px', // 自定义展开区域内 CompoundCell 的标签宽度
            }
          }"
        />
      </div>
    </template>

    <!-- 3. 自定义列插槽 -->
    <template #serverName="{ row }">
      <div
        class="flex items-center space-x-1 group cursor-pointer"
        @click="viewDetails(row)"
      >
        <el-icon
          :size="16"
          class="text-blue-500"
        >
          <Platform />
        </el-icon>
        <span class="font-bold group-hover:text-blue-600 group-hover:underline">{{ row.serverName }}</span>
      </div>
    </template>

    <!-- 4. 操作列插槽 -->
    <template #actions="{ row }">
      <div class="inline-flex">
        <el-button
          type="primary"
          link
          size="small"
          @click="handleEdit(row)"
        >
          编辑
        </el-button>
        <el-popconfirm
          title="确定删除此行吗？"
          @confirm="handleDelete(row)"
        >
          <template #reference>
            <el-button
              type="danger"
              link
              size="small"
            >
              删除
            </el-button>
          </template>
        </el-popconfirm>
      </div>
    </template>
  </ElTablePro>
</template>

<script setup>
import { ref, watch } from 'vue';
import { ElMessage, ElMessageBox, ElPopconfirm } from 'element-plus';
import { Platform, Search, Refresh } from '@element-plus/icons-vue';
import ElTablePro from './ElTablePro.vue'; // 确保路径正确
import CompoundCell from './renderers/compoundCell.vue'; // 为展开行导入

// --- 响应式状态 ---
const total = ref(3);
const loading = ref(false);
const searchQuery = ref('');
const selectedRows = ref([]);

const currentPage = ref(1);
const pageSize = ref(10);

watch(currentPage, (newVal) => {
  console.log('currentPage', currentPage.value)
}
)

watch(pageSize, (newVal) => {
  console.log('pageSize', pageSize.value)
}
)

const wait = (ms) => new Promise((resolve) => {
  window.setTimeout(resolve, ms);
});

const tableData = ref([
  {
    "id": 1,
    "serverName": "DB-Master-1",
    "status": "up",
    "memoryUsage": 0.486,
    "packetLoss": 79.8,
    "uploadRate": 183.6835,
    "downloadSize": 346882529206.6157,
    "lastSeen": '2025-06-07 12:00:00',
    "nodeType": "edge",
    "serviceType": "database",
    "requestCount": 134322349251,
    "region": "华北",
    "sourceIp": "133.74.144.187",
    "logFile": "https://example.com/logs/app-1.log",
    "owner": "李四",
    "shortRequest": "{\"page\":1,\"limit\":10}",
    "fullRequest": {
      "method": "POST",
      "headers": {
        "Content-Type": "application/json",
        "X-Request-ID": "req-1"
      },
      "body": {
        "serverName": "DB-Master-1",
        "action": "restart"
      }
    },
    "protocol": 6,
    "port": "18080",
    "featureEnabled": 1,
    "toggleLocked": false,
    "toggleShouldFail": false,
    "toggleNote": "关闭时会二次确认，异步保存后保留状态",
  },
  {
    "id": 5,
    "serverName": "ETL-Worker-5",
    "status": "down",
    "memoryUsage": 96.7,
    "packetLoss": 0.56,
    "uploadRate": 175775860.7806006,
    "downloadSize": 103568046748.29298,
    "lastSeen": 1750011156636.395,
    "nodeType": "edge",
    "serviceType": "database",
    "requestCount": 508813.3154338964,
    "sourceIp": "157.14.151.61",
    "logFile": "https://example.com/logs/app-5.log",
    "owner": "王五",
    "shortRequest": "{\"page\":1,\"limit\":10}",
    "fullRequest": {
      "method": "POST",
      "headers": {
        "Content-Type": "application/json",
        "X-Request-ID": "req-5"
      },
      "body": {
        "serverName": "ETL-Worker-5",
        "action": "restart"
      }
    },
    "protocol": 17,
    "port": 3306,
    "featureEnabled": 0,
    "toggleLocked": true,
    "toggleShouldFail": false,
    "toggleNote": "该行被 disabled 规则锁定，不允许切换",
  },
  {
    "id": 9,
    "serverName": "Cache-Replica-9",
    "status": "pending",
    "memoryUsage": 34.2,
    "packetLoss": 1.35,
    "uploadRate": 128576.4,
    "downloadSize": 2458291201,
    "lastSeen": '2025-06-07 13:45:10',
    "nodeType": "core",
    "serviceType": "cache",
    "requestCount": 36210500,
    "region": "华东",
    "sourceIp": "10.23.9.88",
    "logFile": "https://example.com/logs/cache-9.log",
    "owner": "张三",
    "shortRequest": "{\"sync\":false,\"cache\":\"warmup\"}",
    "fullRequest": {
      "method": "PUT",
      "headers": {
        "Content-Type": "application/json",
        "X-Request-ID": "req-9"
      },
      "body": {
        "serverName": "Cache-Replica-9",
        "action": "flush"
      }
    },
    "protocol": 6,
    "port": 6379,
    "featureEnabled": 0,
    "toggleLocked": false,
    "toggleShouldFail": true,
    "toggleNote": "模拟接口失败，切换后会自动回滚",
  }
]);


// --- 列定义 (覆盖所有特性) ---
const columns = ref([
  // 原生类型
  { type: 'expand', width: 55, fixed: 'left' },
  { type: 'selection', width: 55, fixed: 'left' },
  { type: 'index', label: 'index', width: 80, fixed: 'left' },
  { type: 'id', prop: 'id', label: 'ID', width: 50, fixed: 'left' },
    { type: 'datetime', prop: 'lastSeen', label: 'datetime', width: 140, config: { format: 'YYYY/MM/DD HH:mm' } },

  // 自定义插槽列 (type: default)
  { prop: 'serverName', label: 'default', width: 140, fixed: 'left' },
  { prop: 'owner', label: 'default-format', width: 150, formatter: (value) => ` ${value}` },

  // 多级表头
  {
    label: 'multi-header',
    children: [
      {
        type: 'status',
        prop: 'status',
        label: 'status',
        width: 100,
        config: {
          'up': { label: '运行中', type: 'success' },
          'restarting': { label: '重启中', type: 'warning', isPing: true }
        }
      },
      {
        type: 'progress',
        prop: 'memoryUsage',
        label: 'progress',
        width: 100,
        config: { valueType: 'decimal', thresholds: { warning: 35, danger: 85 } }
      },
      {
        type: 'progress',
        prop: 'memoryUsage',
        label: 'progress-circle',
        width: 140,
        align: 'center',
        config: {
          type: 'circle',
        }
      },
    ]
  },
  // 聚合单元格 (Compound)
  {
    type: 'compound',
    prop: 'networkMetrics',
    label: 'comound',
    width: 180,
    // 为 compound 列配置排序
    sortBy: 'uploadRate',
    items: [
      { type: 'signal', prop: 'packetLoss', label: 'signal', align: 'right', config: { inverse: true } },
      { type: 'size', prop: 'uploadRate', label: 'size', align: 'center' },
    ],
  },
  { type: 'size', prop: 'uploadRate', label: 'size', width: 100, },
  // 各种独立渲染器
  { type: 'ip', prop: 'sourceIp', label: 'ip', width: 150, align: 'center' },
  { type: 'protocol', prop: 'protocol', label: 'protocol', width: 120 },
  { type: 'port', prop: 'port', label: 'port', width: 120 },
  {
    type: 'tag',
    prop: 'nodeType',
    label: 'tag',
    width: 100,
    config: {
      'core': { label: '核心', type: 'danger' },
      'edge': { label: '边缘', type: 'success' },
    },
  },
  { type: 'icon', prop: 'serviceType', label: 'icon', width: 100 },
  {
    type: 'switch',
    prop: 'featureEnabled',
    label: 'switch',
    width: 110,
    align: 'center',
    config: {
      activeValue: 1,
      inactiveValue: 0,
      inlinePrompt: true,
      activeText: '开',
      inactiveText: '关',
      disabled: ({ row }) => row.toggleLocked,
      beforeChange: async ({ row, nextValue }) => {
        if (nextValue !== 0) return true;

        try {
          await ElMessageBox.confirm(
            `确认关闭 ${row.serverName} 的功能开关吗？`,
            '关闭确认',
            {
              type: 'warning',
              confirmButtonText: '确认关闭',
              cancelButtonText: '取消',
            }
          );
          return true;
        } catch (_error) {
          return false;
        }
      },
      onChange: async ({ row, nextValue }) => {
        await wait(700);

        if (row.toggleShouldFail) {
          ElMessage.error(`${row.serverName} 模拟保存失败，已自动回滚`);
          throw new Error('mock failure');
        }

        row.toggleNote = nextValue === 1
          ? '最近一次切换：已开启'
          : '最近一次切换：已关闭';
        ElMessage.success(`${row.serverName} 开关已${nextValue === 1 ? '开启' : '关闭'}`);
        return true;
      },
    },
  },
  { prop: 'toggleNote', label: 'switch-note', minWidth: 220 },
  {
    type: 'unit',
    prop: 'requestCount',
    label: 'unit',
    align: 'right',
    width: 80,
    config: {
      unitBase: 1000,
      units: ['', 'K', 'M', 'B'],
      decimals: 1,
    },
  },
  { type: 'url', prop: 'logFile', label: 'url', width: 150 },
  { type: 'json', prop: 'shortRequest', label: 'json', width: 150, previewText: '查看参数' },

]);

// --- 事件处理 ---
const handlePageChange = ()=>{
  // 处理向后端获取数据逻辑...
  console.log('page-change', currentPage.value, pageSize.value)
}

const handleSortChange = ({ prop, order }) => {
  ElMessage.info(`后端排序触发：列=${prop}, 顺序=${order}`);
  // 如果是 avgLatency 列，则可以调用 fetchData 重新获取数据
  if (prop === 'avgLatency') {
    // fetchData({ sortBy: prop, sortOrder: order });
  }
};

const handleSelectionChange = (val) => {
  selectedRows.value = val;
  console.log('选中的行:', selectedRows.value);
};

const viewDetails = (row) => ElMessage.info(`查看详情: ${row.serverName}`);
const handleEdit = (row) => ElMessage.success(`编辑: ${row.serverName}`);
const handleDelete = (row) => ElMessage.error(`删除成功: ${row.serverName}`);

</script>
<style scoped>
::v-deep .el-table .el-table__header-wrapper th {
  @apply p-2 !important;
}

::v-deep .el-table .el-table__row td {
  @apply p-2 !important;
}
</style>
