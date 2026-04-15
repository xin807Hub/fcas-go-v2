<template>
  <div class="space-y-2">
    <el-card>
      <div class="inline-flex w-full gap-2 justify-self-center">
        <date-picker v-model:start-time="searchForm.startTime" v-model:end-time="searchForm.endTime" />
        <isp-select v-model:selected="searchForm.ispNameList" />
        <link-select v-model:selected="searchForm.linkIdList" />
        <app-type-select v-model:selected="searchForm.appIdList" />
        <user-tree-select v-model:selected="searchForm.crowdIdList" tree-type="crowdTree" />
        <div class="inline-flex w-full items-center justify-between">
          <el-button icon="Search" type="primary" size="small" class="button" :loading="level1TableState.loading || pieLoading" @click="handleSearch">查询</el-button>
          <el-button icon="Promotion" type="primary" size="small" class="button justify-self-end" @click="handleExport">导出</el-button>
        </div>
      </div>
    </el-card>

    <div class="space-y-3 dark:bg-slate-900">
      <div v-loading="pieLoading" class="grid grid-cols-1 gap-3 xl:grid-cols-2">
        <div class="bg-white rounded-xl dark:border-slate-700">
          <h3 class="chart-card-title">下行总量占比</h3>
          <PieChart :data="level1PieChartData.trafficDown" name-key="name" value-key="value" :show-pie-percent="true" legend="none" :height="280" :formatter="formatTraffic" />
        </div>
        <div class="bg-white rounded-xl dark:border-slate-700">
          <h3 class="chart-card-title">上下行总量占比</h3>
          <PieChart :data="level1PieChartData.trafficTotal" name-key="name" value-key="value" :show-pie-percent="true" legend="none" :height="280" :formatter="formatTraffic" />
        </div>
      </div>

      <div class="bg-white rounded-xl p-2">
        <ElTablePro :data="level1TableState.rows" :columns="tableColumns" :loading="level1TableState.loading" :show-pagination="false" :height="445" border>
          <template #name="{ row }">
            <el-button :loading="row.panelLoading" size="small" icon="Link" type="primary" link @click="handleOpenLevel2(row)">
              <el-text class="w-full" type="primary" truncated>{{ row.name }}</el-text>
            </el-button>
          </template>
          <template #totalRatio="{ row }">{{ formatTotalRatio(row.trafficTotal) }}</template>
        </ElTablePro>

        <div class="mt-2 ml-2 inline-flex flex-wrap items-center gap-5 text-gray-500">
          <p class="text-lg font-black">汇总</p>
          <p v-for="(value, key) in level1TableState.gather" :key="key">{{ key }}：<span class="font-bold text-blue-500">{{ value }}</span></p>
        </div>
      </div>
    </div>

    <div v-if="level2PanelVisible" ref="level2PanelRef" class="rounded-md bg-white p-3 shadow-md space-y-3 dark:bg-slate-900">
      <p class="text-2xl font-black">[ {{ selectedLevel2UserName }} ] - 用户详情</p>
      <div class="rounded-md border border-slate-200 p-2 dark:border-slate-700">
        <h3 class="detail-card-title">流量趋势</h3>
        <TrendChart v-loading="level2TrendLoading" :data="level2TrendChartData" x-axis-key="startTime" :series="level2TrendSeries" :start-time="searchForm.startTime" :end-time="searchForm.endTime" y-axis-name="速率" :height="320" legend="top" />
      </div>
      <div class="rounded-md border border-slate-200 p-2 dark:border-slate-700">
        <h3 class="detail-card-title">IP排行</h3>
        <PieChart v-loading="level2TableLoading" :data="level2BarChartData" chart-type="bar" bar-direction="horizontal" name-key="srcIp" value-key="trafficTotal" :name-max-length="18" :height="520" :show-y-axis="true" :show-bar-value="false" :formatter="formatTraffic" />
      </div>
      <div class="rounded-md border border-slate-200 p-2 dark:border-slate-700">
        <h3 class="detail-card-title">IP明细</h3>
        <ElTablePro :data="level2TableData" :columns="level2TableColumns" :loading="level2TableLoading" :show-pagination="false" :height="420" border>
          <template #srcIp="{ row }">
            <el-button :loading="row.level3Loading" size="small" icon="Link" type="primary" link @click="handleOpenLevel3(row)">
              <el-text class="w-full" type="primary" truncated>{{ row.srcIp }}</el-text>
            </el-button>
          </template>
          <template #totalProportion="{ row }">{{ ((row.totalProportion || 0) * 100).toFixed(4) }}%</template>
        </ElTablePro>
      </div>
    </div>

    <div v-if="level3PanelVisible" ref="level3PanelRef" class="rounded-md bg-white p-3 shadow-md space-y-3 dark:bg-slate-900">
      <p class="text-2xl font-black">[ {{ selectedLevel3Ip }} ] - 流速详情</p>
      <div class="rounded-md border border-slate-200 p-2 dark:border-slate-700">
        <h3 class="detail-card-title">流速趋势</h3>
        <TrendChart v-loading="level3Loading" :data="level3TableData" x-axis-key="startTime" :series="level3TrendSeries" :start-time="searchForm.startTime" :end-time="searchForm.endTime" y-axis-name="速率" :height="320" legend="top" />
      </div>
      <div class="rounded-md border border-slate-200 p-2 dark:border-slate-700">
        <h3 class="detail-card-title">流速明细</h3>
        <ElTablePro :data="level3TableData" :columns="level3TableColumns" :loading="level3Loading" :show-pagination="false" :height="420" border />
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, nextTick, onMounted, ref } from "vue";
import axios from "axios";
import { saveAs } from "file-saver";
import dayjs from "dayjs";
import ElTablePro from "@/components/ElTablePro/ElTablePro.vue";
import PieChart from "@/components/charts/PieChart.vue";
import TrendChart from "@/components/charts/TrendChart.vue";
import IspSelect from "@/components/searchItem/IspSelect.vue";
import UserTreeSelect from "@/components/searchItem/UserTreeSelect.vue";
import AppTypeSelect from "@/components/searchItem/AppTypeSelect.vue";
import DatePicker from "@/components/searchItem/DatePicker.vue";
import LinkSelect from "@/components/searchItem/LinkSelect.vue";
import { formatSize } from "@/utils/format";
import { userCrowdRankLevel1PieApi, userCrowdRankLevel1TableApi, userCrowdRankLevel2TableApi, userCrowdRankLevel2TrendApi, userCrowdRankLevel3TableApi } from "@/api/traffic";

const searchForm = ref({ startTime: dayjs().subtract(1, "hour").format("YYYY-MM-DD HH:mm:ss"), endTime: dayjs().format("YYYY-MM-DD HH:mm:ss"), ispNameList: [], linkIdList: [], crowdIdList: [], appIdList: [] });
const pieLoading = ref(false);
const level2PanelVisible = ref(false);
const level2PanelRef = ref(null);
const selectedLevel2UserName = ref("");
const selectedLevel2UserId = ref(null);
const level2TrendLoading = ref(false);
const level2TableLoading = ref(false);
const level2TrendChartData = ref([]);
const level2TableData = ref([]);
const level3PanelVisible = ref(false);
const level3PanelRef = ref(null);
const selectedLevel3Ip = ref("");
const level3Loading = ref(false);
const level3TableData = ref([]);
const level1PieChartData = ref({ trafficTotal: [], trafficDown: [] });
const level1TableState = ref({ loading: false, trafficTotal: 0, rows: [], gather: {} });

const tableColumns = computed(() => [
  { type: "index", label: "排名", width: 80 },
  { prop: "name", label: "用户名称", minWidth: 220 },
  { type: "bps", prop: "maxSpeedUp", label: "上行峰值", minWidth: 120, sortable: true },
  { type: "bps", prop: "maxSpeedDn", label: "下行峰值", minWidth: 120, sortable: true },
  { type: "bps", prop: "avgSpeedUp", label: "上行平均", minWidth: 120, sortable: true },
  { type: "bps", prop: "avgSpeedDn", label: "下行平均", minWidth: 120, sortable: true },
  { type: "size", prop: "trafficUp", label: "上行总量", minWidth: 120, sortable: true },
  { type: "size", prop: "trafficDn", label: "下行总量", minWidth: 120, sortable: true },
  { type: "size", prop: "trafficTotal", label: "上下行总量", minWidth: 130, sortable: true },
  { prop: "totalRatio", label: "总量占比", minWidth: 110 },
]);
const level2TableColumns = computed(() => [
  { type: "index", label: "排名", width: 80 },
  { prop: "srcIp", label: "源IP", minWidth: 180 },
  { type: "bps", prop: "avgSpeedUp", label: "上行平均", minWidth: 120, sortable: true },
  { type: "bps", prop: "avgSpeedDn", label: "下行平均", minWidth: 120, sortable: true },
  { type: "size", prop: "trafficUp", label: "上行总量", minWidth: 120, sortable: true },
  { type: "size", prop: "trafficDn", label: "下行总量", minWidth: 120, sortable: true },
  { type: "size", prop: "trafficTotal", label: "上下行总量", minWidth: 130, sortable: true },
  { prop: "totalProportion", label: "总量占比", minWidth: 110 },
]);
const level2TrendSeries = computed(() => [
  { name: "speedUp", label: "上行流量", formatter: formatRate },
  { name: "speedDn", label: "下行流量", formatter: formatRate },
]);
const level2BarChartData = computed(() => level2TableData.value.slice().sort((a, b) => a.trafficTotal - b.trafficTotal));
const level3TrendSeries = computed(() => [
  { name: "speedUp", label: "上行流速", formatter: formatRate },
  { name: "speedDn", label: "下行流速", formatter: formatRate },
  { name: "totalSpeed", label: "总流速", formatter: formatRate },
]);
const level3TableColumns = computed(() => [
  { type: "index", label: "序号", width: 80 },
  { prop: "startTime", label: "时间", minWidth: 180, sortable: true, formatter: (row) => dayjs(row.startTime).format("YYYY-MM-DD HH:mm:ss") },
  { type: "bps", prop: "speedUp", label: "上行流速", minWidth: 120, sortable: true },
  { type: "bps", prop: "speedDn", label: "下行流速", minWidth: 120, sortable: true },
  { type: "bps", prop: "totalSpeed", label: "总流速", minWidth: 120, sortable: true },
]);

function formatTraffic(value) { return formatSize(value); }
function formatRate(value) { return formatSize(value, { units: ["bps", "Kbps", "Mbps", "Gbps", "Tbps"] }).replace(" (", "").replace(")", ""); }
function formatTotalRatio(trafficTotal) { return `${(((trafficTotal || 0) / (level1TableState.value.trafficTotal || 1)) * 100).toFixed(4)}%`; }

const fetchLevel1PieChartData = async () => {
  pieLoading.value = true;
  try {
    const res = await userCrowdRankLevel1PieApi({ ...searchForm.value });
    if (res.code === 0) {
      level1PieChartData.value.trafficTotal = res.data?.map((item) => ({ name: item.name, value: item.trafficTotal })) || [];
      level1PieChartData.value.trafficDown = res.data?.map((item) => ({ name: item.name, value: item.trafficDn })) || [];
    }
  } finally { pieLoading.value = false; }
};

const fetchLevel1TableData = async () => {
  level1TableState.value.loading = true;
  try {
    const res = await userCrowdRankLevel1TableApi({ ...searchForm.value });
    if (res.code === 0) {
      level1TableState.value.trafficTotal = res.data?.gather?.trafficTotal || 0;
      level1TableState.value.rows = res.data?.rows?.map((item) => ({ ...item, panelLoading: false })) || [];
      const formatConfig = { avgSpeedUp: { label: "上行平均", units: ["bps", "Kbps", "Mbps", "Gbps"] }, avgSpeedDn: { label: "下行平均", units: ["bps", "Kbps", "Mbps", "Gbps"] }, trafficUp: { label: "上行总量" }, trafficDn: { label: "下行总量" }, trafficTotal: { label: "上下行总量" } };
      const gather = res.data?.gather || {};
      level1TableState.value.gather = Object.entries(formatConfig).reduce((acc, [key, value]) => { acc[value.label] = formatSize(gather[key], value); return acc; }, {});
    }
  } finally { level1TableState.value.loading = false; }
};

const loadLevel2TrendData = async (row) => {
  level2TrendLoading.value = true;
  try {
    const res = await userCrowdRankLevel2TrendApi({ startTime: searchForm.value.startTime, endTime: searchForm.value.endTime, crowdIdList: [row.id], topN: 20 });
    level2TrendChartData.value = res.code === 0 ? res.data || [] : [];
  } finally { level2TrendLoading.value = false; }
};

const loadLevel2TableData = async (row) => {
  level2TableLoading.value = true;
  try {
    const res = await userCrowdRankLevel2TableApi({ startTime: searchForm.value.startTime, endTime: searchForm.value.endTime, crowdIdList: [row.id], topN: 20 });
    level2TableData.value = res.code === 0 ? (res.data || []).map((item) => ({ ...item, level3Loading: false })) : [];
  } finally { level2TableLoading.value = false; }
};

const handleOpenLevel2 = async (row) => {
  row.panelLoading = true;
  try {
    selectedLevel2UserName.value = row.name;
    selectedLevel2UserId.value = row.id;
    await Promise.all([loadLevel2TrendData(row), loadLevel2TableData(row)]);
    level2PanelVisible.value = true;
    nextTick(() => level2PanelRef.value?.scrollIntoView({ behavior: "smooth" }));
  } finally { row.panelLoading = false; }
};

const handleOpenLevel3 = async (row) => {
  row.level3Loading = true;
  level3Loading.value = true;
  try {
    selectedLevel3Ip.value = row.srcIp;
    const res = await userCrowdRankLevel3TableApi({ startTime: searchForm.value.startTime, endTime: searchForm.value.endTime, srcIp: row.srcIp });
    level3TableData.value = res.code === 0 ? res.data || [] : [];
    level3PanelVisible.value = true;
    nextTick(() => level3PanelRef.value?.scrollIntoView({ behavior: "smooth" }));
  } finally { level3Loading.value = false; row.level3Loading = false; }
};

const handleSearch = async () => {
  await Promise.all([fetchLevel1PieChartData(), fetchLevel1TableData()]);
  if (level2PanelVisible.value && selectedLevel2UserId.value !== null) {
    const currentRow = level1TableState.value.rows.find((item) => item.id === selectedLevel2UserId.value);
    if (currentRow) await Promise.all([loadLevel2TrendData(currentRow), loadLevel2TableData(currentRow)]);
  }
  if (level3PanelVisible.value && selectedLevel3Ip.value) {
    const res = await userCrowdRankLevel3TableApi({ startTime: searchForm.value.startTime, endTime: searchForm.value.endTime, srcIp: selectedLevel3Ip.value });
    level3TableData.value = res.code === 0 ? res.data || [] : [];
  }
};

const handleExport = async () => {
  axios({ method: "post", url: "/api/traffic/userCrowdRank/export", data: searchForm.value, responseType: "blob" })
    .then((response) => {
      const filename = `user_rank_${+new Date()}.xlsx`;
      saveAs(new Blob([response.data], { type: response.headers["contnet-type"] }), filename);
    })
    .catch((error) => console.error("下载文件失败：", error));
};

onMounted(async () => { await handleSearch(); });
</script>

<style scoped lang="scss">
.chart-card-title { display: block; width: 100%; text-align: center; font-size: 18px; font-weight: 700; letter-spacing: 0.5px; color: rgb(51 65 85); }
.dark .chart-card-title { color: rgb(226 232 240); }
.detail-card-title { margin-bottom: 12px; text-align: left; font-size: 18px; font-weight: 700; line-height: 28px; color: rgb(51 65 85); }
.dark .detail-card-title { color: rgb(226 232 240); }
</style>
