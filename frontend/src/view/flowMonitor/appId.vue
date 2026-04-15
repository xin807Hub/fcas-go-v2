<template>
  <div class="space-y-2">
    <el-card>
      <div class="inline-flex w-full gap-2 justify-self-center">
        <date-picker
          v-model:start-time="searchForm.startTime"
          v-model:end-time="searchForm.endTime"
        />
        <area-select v-model:selected="searchForm.dstProvince" />
        <isp-select v-model:selected="searchForm.ispNameList" />
        <link-select v-model:selected="searchForm.linkIdList" />
        <app-type-select v-model:selected="searchForm.appIdList" />
        <user-tree-select v-model:selected="searchForm.userIdList" />
        <div class="inline-flex w-full items-center justify-between">
          <el-button
            icon="Search"
            type="primary"
            size="small"
            class="button"
            :loading="level1TableState.loading || pieLoading"
            @click="handleSearch"
          >
            查询
          </el-button>
          <el-button
            icon="Promotion"
            type="primary"
            size="small"
            class="button justify-self-end"
            @click="handleExport"
          >
            导出
          </el-button>
        </div>
      </div>
    </el-card>

    <div class="space-y-3 dark:bg-slate-900">
      <div
        v-loading="pieLoading"
        class="grid grid-cols-1 gap-3 xl:grid-cols-2"
      >
        <div class="rounded-xl bg-white dark:border-slate-700">
          <h3 class="chart-card-title">下行总量占比</h3>
          <PieChart
            :data="level1PieChartData.trafficDown"
            name-key="name"
            value-key="value"
            :show-pie-percent="true"
            legend="none"
            :height="280"
            :formatter="formatTraffic"
          />
        </div>

        <div class="rounded-xl bg-white dark:border-slate-700">
          <h3 class="chart-card-title">上下行总量占比</h3>
          <PieChart
            :data="level1PieChartData.trafficTotal"
            name-key="name"
            value-key="value"
            :show-pie-percent="true"
            legend="none"
            :height="280"
            :formatter="formatTraffic"
          />
        </div>
      </div>

      <div class="rounded-xl bg-white p-2">
        <ElTablePro
          :data="level1TableState.rows"
          :columns="level1Columns"
          :loading="level1TableState.loading"
          :show-pagination="false"
          :height="425"
          border
        >
          <template #name="{ row }">
            <el-button
              :loading="row.panelLoading"
              size="small"
              icon="Link"
              type="primary"
              link
              @click="handleOpenLevel2(row)"
            >
              <el-text class="w-full" type="primary" truncated>
                {{ row.name }}
              </el-text>
            </el-button>
          </template>

          <template #totalRatio="{ row }">
            {{ formatLevel1Ratio(row.totalByte) }}
          </template>
        </ElTablePro>

        <div class="flex items-end justify-between">
          <div
            class="ml-2 inline-flex flex-wrap items-center gap-5 text-gray-500"
          >
            <p class="text-lg font-black">汇总</p>
            <p v-for="(value, key) in level1TableState.gather" :key="key">
              {{ key }}：
              <span class="font-bold text-blue-500">{{ value }}</span>
            </p>
          </div>

          <el-pagination
            v-model:current-page="pagination.page"
            v-model:page-size="pagination.limit"
            background
            size="small"
            :total="level1TableState.total"
            :page-sizes="[20, 50, 100]"
            layout="prev, pager, next, total"
            @current-change="fetchLevel1TableData"
            @size-change="fetchLevel1TableData"
          />
        </div>
      </div>
    </div>

    <div
      v-if="level2PanelVisible"
      ref="level2PanelRef"
      class="space-y-3 rounded-md bg-white p-3 shadow-md dark:bg-slate-900"
    >
      <p class="text-2xl font-black">[ {{ selectedLevel2Name }} ] - 业务详情</p>

      <div class="rounded-md border border-slate-200 p-2 dark:border-slate-700">
        <h3 class="detail-card-title">流速趋势</h3>
        <TrendChart
          v-loading="level2TrendLoading"
          :data="level2TrendChartData"
          x-axis-key="startTime"
          :series="level2TrendSeries"
          :start-time="searchForm.startTime"
          :end-time="searchForm.endTime"
          y-axis-name="速率"
          :height="320"
          legend="top"
        />
      </div>

      <div class="rounded-md border border-slate-200 p-2 dark:border-slate-700">
        <h3 class="detail-card-title">流速明细</h3>
        <ElTablePro
          :data="level2TrendChartData"
          :columns="level2TrendColumns"
          :loading="level2TrendLoading"
          :show-pagination="false"
          :height="420"
          border
        />
      </div>

      <div class="rounded-md border border-slate-200 p-2 dark:border-slate-700">
        <h3 class="detail-card-title">目的IP排行</h3>
        <PieChart
          v-loading="level2DstIpLoading"
          :data="level2BarChartData.dstIp"
          chart-type="bar"
          bar-direction="horizontal"
          name-key="name"
          value-key="value"
          :name-max-length="22"
          :height="520"
          :show-y-axis="true"
          :show-bar-value="false"
          :formatter="formatTraffic"
        />
      </div>

      <div class="rounded-md border border-slate-200 p-2 dark:border-slate-700">
        <h3 class="detail-card-title">目的IP明细</h3>
        <ElTablePro
          :data="level2DstIpTableState.rows"
          :columns="level2DstIpColumns"
          :loading="level2DstIpLoading"
          :show-pagination="false"
          :height="420"
          border
        >
          <template #name="{ row }">
            <el-button
              :loading="row.level3Loading"
              size="small"
              icon="Link"
              type="primary"
              link
              @click="handleOpenLevel3(false, row)"
            >
              <el-text class="w-full" type="primary" truncated>
                {{ row.name }}
              </el-text>
            </el-button>
          </template>
        </ElTablePro>

        <div class="mt-2 flex justify-end">
          <el-pagination
            v-model:current-page="level2DstIpTableState.page"
            v-model:page-size="level2DstIpTableState.limit"
            background
            size="small"
            :total="level2DstIpTableState.total"
            :page-sizes="[20, 50, 100]"
            layout="prev, pager, next, total"
            @current-change="handleLevel2TablePageChange('dstIp')"
            @size-change="handleLevel2TablePageChange('dstIp')"
          />
        </div>
      </div>

      <div class="rounded-md border border-slate-200 p-2 dark:border-slate-700">
        <h3 class="detail-card-title">域名排行</h3>
        <PieChart
          v-loading="level2HostLoading"
          :data="level2BarChartData.host"
          chart-type="bar"
          bar-direction="horizontal"
          name-key="name"
          value-key="value"
          :name-max-length="26"
          :height="520"
          :show-y-axis="true"
          :show-bar-value="false"
          :formatter="formatTraffic"
        />
      </div>

      <div class="rounded-md border border-slate-200 p-2 dark:border-slate-700">
        <h3 class="detail-card-title">域名明细</h3>
        <ElTablePro
          :data="level2HostTableState.rows"
          :columns="level2HostColumns"
          :loading="level2HostLoading"
          :show-pagination="false"
          :height="420"
          border
        >
          <template #name="{ row }">
            <el-button
              :loading="row.level3Loading"
              size="small"
              icon="Link"
              type="primary"
              link
              @click="handleOpenLevel3(true, row)"
            >
              <el-text class="w-full" type="primary" truncated>
                {{ row.name }}
              </el-text>
            </el-button>
          </template>
        </ElTablePro>

        <div class="mt-2 flex justify-end">
          <el-pagination
            v-model:current-page="level2HostTableState.page"
            v-model:page-size="level2HostTableState.limit"
            background
            size="small"
            :total="level2HostTableState.total"
            :page-sizes="[20, 50, 100]"
            layout="prev, pager, next, total"
            @current-change="handleLevel2TablePageChange('host')"
            @size-change="handleLevel2TablePageChange('host')"
          />
        </div>
      </div>
    </div>

    <div
      v-if="level3PanelVisible"
      ref="level3PanelRef"
      class="space-y-3 rounded-md bg-white p-3 shadow-md dark:bg-slate-900"
    >
      <p class="text-2xl font-black">{{ level3PanelTitle }}</p>

      <div class="rounded-md border border-slate-200 p-2 dark:border-slate-700">
        <h3 class="detail-card-title">流量趋势</h3>
        <TrendChart
          v-loading="level3TrendLoading"
          :data="level3TrendChartData"
          x-axis-key="startTime"
          :series="level3TrendSeries"
          :start-time="searchForm.startTime"
          :end-time="searchForm.endTime"
          y-axis-name="速率"
          :height="320"
          legend="top"
        />
      </div>

      <div class="rounded-md border border-slate-200 p-2 dark:border-slate-700">
        <h3 class="detail-card-title">流速明细</h3>
        <ElTablePro
          :data="level3TableState.rows"
          :columns="level3Columns"
          :loading="level3TableLoading"
          :show-pagination="false"
          :height="420"
          border
        />

        <div class="mt-2 flex justify-end">
          <el-pagination
            v-model:current-page="level3TableState.page"
            v-model:page-size="level3TableState.limit"
            background
            size="small"
            :total="level3TableState.total"
            :page-sizes="[20, 50, 100]"
            layout="prev, pager, next, total"
            @current-change="fetchLevel3TableData"
            @size-change="fetchLevel3TableData"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, nextTick, onMounted, ref } from "vue";
import axios from "axios";
import { saveAs } from "file-saver";
import dayjs from "dayjs";
import { ElMessage } from "element-plus";
import ElTablePro from "@/components/ElTablePro/ElTablePro.vue";
import PieChart from "@/components/charts/PieChart.vue";
import TrendChart from "@/components/charts/TrendChart.vue";
import LinkSelect from "@/components/searchItem/LinkSelect.vue";
import UserTreeSelect from "@/components/searchItem/UserTreeSelect.vue";
import DatePicker from "@/components/searchItem/DatePicker.vue";
import AreaSelect from "@/components/searchItem/AreaSelect.vue";
import IspSelect from "@/components/searchItem/IspSelect.vue";
import AppTypeSelect from "@/components/searchItem/AppTypeSelect.vue";
import { formatSize } from "@/utils/format";
import {
  appIdRankDataApi,
  appIdRankTableDataApi,
  appIdRankTrendDataApi,
} from "@/api/traffic";

const searchForm = ref({
  rankLevel: "level1",
  startTime: dayjs().subtract(1, "hour").format("YYYY-MM-DD HH:mm:ss"),
  endTime: dayjs().format("YYYY-MM-DD HH:mm:ss"),
  dstProvince: null,
  ispNameList: [],
  linkIdList: [],
  userIdList: [],
  appIdList: [],
});

const pieLoading = ref(false);
const level1TableState = ref({
  loading: false,
  rows: [],
  total: 0,
  trafficTotal: 0,
  gather: {},
});
const level1PieChartData = ref({
  trafficDown: [],
  trafficTotal: [],
});
const pagination = ref({
  page: 1,
  limit: 20,
});

const level2PanelVisible = ref(false);
const level2PanelRef = ref(null);
const selectedLevel2Row = ref(null);
const selectedLevel2Name = ref("");
const level2TrendLoading = ref(false);
const level2TrendChartData = ref([]);
const level2DstIpLoading = ref(false);
const level2HostLoading = ref(false);
const level2BarChartData = ref({
  dstIp: [],
  host: [],
});
const level2DstIpTableState = ref({
  rows: [],
  total: 0,
  page: 1,
  limit: 20,
});
const level2HostTableState = ref({
  rows: [],
  total: 0,
  page: 1,
  limit: 20,
});

const level3PanelVisible = ref(false);
const level3PanelRef = ref(null);
const level3PanelTitle = ref("");
const selectedLevel3Target = ref({
  isHost: false,
  name: "",
});
const level3TrendLoading = ref(false);
const level3TableLoading = ref(false);
const level3TrendChartData = ref([]);
const level3TableState = ref({
  rows: [],
  total: 0,
  page: 1,
  limit: 20,
});

const level1Columns = computed(() => [
  { type: "index", label: "排名", width: 80 },
  { prop: "name", label: "业务小类", minWidth: 180 },
  {
    type: "bps",
    prop: "maxUpBps",
    label: "上行峰值",
    minWidth: 120,
    sortable: true,
  },
  {
    type: "bps",
    prop: "maxDnBps",
    label: "下行峰值",
    minWidth: 120,
    sortable: true,
  },
  {
    type: "bps",
    prop: "avgUpBps",
    label: "上行平均",
    minWidth: 120,
    sortable: true,
  },
  {
    type: "bps",
    prop: "avgDnBps",
    label: "下行平均",
    minWidth: 120,
    sortable: true,
  },
  {
    type: "size",
    prop: "upByte",
    label: "上行总量",
    minWidth: 120,
    sortable: true,
  },
  {
    type: "size",
    prop: "dnByte",
    label: "下行总量",
    minWidth: 120,
    sortable: true,
  },
  {
    type: "size",
    prop: "totalByte",
    label: "上下行总量",
    minWidth: 130,
    sortable: true,
  },
  { prop: "totalRatio", label: "总量占比", minWidth: 110 },
]);

const level2TrendSeries = computed(() => [
  { name: "upBps", label: "上行流速", formatter: formatRate },
  { name: "dnBps", label: "下行流速", formatter: formatRate },
]);

const level2TrendColumns = computed(() => [
  {
    prop: "startTime",
    label: "时间",
    minWidth: 180,
    sortable: true,
    formatter: (row) => dayjs(row.startTime).format("YYYY-MM-DD HH:mm:ss"),
  },
  {
    type: "bps",
    prop: "upBps",
    label: "上行流速",
    minWidth: 120,
    sortable: true,
  },
  {
    type: "bps",
    prop: "dnBps",
    label: "下行流速",
    minWidth: 120,
    sortable: true,
  },
  {
    type: "bps",
    prop: "totalBps",
    label: "总流速",
    minWidth: 120,
    sortable: true,
  },
]);

const level2DstIpColumns = computed(() => [
  { type: "index", label: "排名", width: 80 },
  { prop: "name", label: "目的IP", minWidth: 180 },
  { prop: "location", label: "归属信息", minWidth: 180 },
  {
    type: "bps",
    prop: "avgUpBps",
    label: "上行平均",
    minWidth: 120,
    sortable: true,
  },
  {
    type: "bps",
    prop: "avgDnBps",
    label: "下行平均",
    minWidth: 120,
    sortable: true,
  },
  {
    type: "size",
    prop: "upByte",
    label: "上行总量",
    minWidth: 120,
    sortable: true,
  },
  {
    type: "size",
    prop: "dnByte",
    label: "下行总量",
    minWidth: 120,
    sortable: true,
  },
]);

const level2HostColumns = computed(() => [
  { type: "index", label: "排名", width: 80 },
  { prop: "name", label: "域名", minWidth: 220 },
  {
    type: "bps",
    prop: "avgUpBps",
    label: "上行平均",
    minWidth: 120,
    sortable: true,
  },
  {
    type: "bps",
    prop: "avgDnBps",
    label: "下行平均",
    minWidth: 120,
    sortable: true,
  },
  {
    type: "size",
    prop: "upByte",
    label: "上行总量",
    minWidth: 120,
    sortable: true,
  },
  {
    type: "size",
    prop: "dnByte",
    label: "下行总量",
    minWidth: 120,
    sortable: true,
  },
]);

const level3TrendSeries = computed(() => [
  { name: "upBps", label: "上行流速", formatter: formatRate },
  { name: "dnBps", label: "下行流速", formatter: formatRate },
]);

const level3Columns = computed(() => [
  {
    prop: "startTime",
    label: "时间",
    minWidth: 180,
    sortable: true,
    formatter: (row) => dayjs(row.startTime).format("YYYY-MM-DD HH:mm:ss"),
  },
  {
    type: "bps",
    prop: "upBps",
    label: "上行流速",
    minWidth: 120,
    sortable: true,
  },
  {
    type: "bps",
    prop: "dnBps",
    label: "下行流速",
    minWidth: 120,
    sortable: true,
  },
  {
    type: "bps",
    prop: "totalBps",
    label: "总流速",
    minWidth: 120,
    sortable: true,
  },
]);

function formatTraffic(value) {
  return formatSize(value);
}

function formatRate(value) {
  return formatSize(value, {
    units: ["bps", "Kbps", "Mbps", "Gbps", "Tbps"],
  })
    .replace(" (", "")
    .replace(")", "");
}

function formatLevel1Ratio(totalByte) {
  return `${(
    ((totalByte || 0) / (level1TableState.value.trafficTotal || 1)) *
    100
  ).toFixed(4)}%`;
}

const fetchLevel1TableData = async () => {
  level1TableState.value.loading = true;
  try {
    const resp = await appIdRankTrendDataApi({
      ...searchForm.value,
      ...pagination.value,
    });
    if (resp.code === 0) {
      level1TableState.value.rows =
        resp.data?.list?.map((item) => ({
          ...item,
          panelLoading: false,
        })) || [];
      level1TableState.value.total = resp.data?.totalCount || 0;
    }
  } finally {
    level1TableState.value.loading = false;
  }
};

const fetchLevel1PieChartData = async () => {
  pieLoading.value = true;
  try {
    const res = await appIdRankDataApi({ ...searchForm.value });
    if (res.code === 0) {
      level1PieChartData.value.trafficTotal =
        res.data?.pieData?.map((item) => ({
          name: item.name,
          value: item.totalByte,
        })) || [];
      level1PieChartData.value.trafficDown =
        res.data?.pieData?.map((item) => ({
          name: item.name,
          value: item.dnByte,
        })) || [];

      level1TableState.value.trafficTotal = res.data?.gatherData?.totalByte || 1;

      const formatConfig = {
        avgUpBps: { label: "上行平均", units: ["bps", "Kbps", "Mbps", "Gbps"] },
        avgDnBps: { label: "下行平均", units: ["bps", "Kbps", "Mbps", "Gbps"] },
        upByte: { label: "上行总量" },
        dnByte: { label: "下行总量" },
        totalByte: { label: "上下行总量" },
      };
      const gather = res.data?.gatherData || {};
      level1TableState.value.gather = Object.entries(formatConfig).reduce(
        (acc, [key, config]) => {
          acc[config.label] = formatSize(gather[key], config);
          return acc;
        },
        {},
      );
    }
  } finally {
    pieLoading.value = false;
  }
};

const fetchLevel2Data = async () => {
  level2TrendLoading.value = true;
  try {
    const resp = await appIdRankDataApi({
      rankLevel: "level2",
      startTime: searchForm.value.startTime,
      endTime: searchForm.value.endTime,
      appIdList: [selectedLevel2Row.value.appId],
    });
    if (resp.code === 0) {
      level2TrendChartData.value =
        resp.data?.trendData?.slice().sort((a, b) => new Date(a.startTime) - new Date(b.startTime)) || [];
      level2BarChartData.value = {
        dstIp:
          resp.data?.dstIpRankBar?.map((item) => ({
            name: item.name,
            value: item.value ?? item.totalByte ?? 0,
          })) || [],
        host:
          resp.data?.hostRankBar?.map((item) => ({
            name: item.name,
            value: item.value ?? item.totalByte ?? 0,
          })) || [],
      };
    } else {
      level2TrendChartData.value = [];
      level2BarChartData.value = { dstIp: [], host: [] };
    }
  } finally {
    level2TrendLoading.value = false;
  }
};

const fetchLevel2TableData = async (rankType) => {
  const targetState =
    rankType === "dstIp" ? level2DstIpTableState.value : level2HostTableState.value;
  const loadingRef = rankType === "dstIp" ? level2DstIpLoading : level2HostLoading;

  loadingRef.value = true;
  try {
    const resp = await appIdRankTableDataApi({
      rankLevel: "level2",
      startTime: searchForm.value.startTime,
      endTime: searchForm.value.endTime,
      appIdList: [selectedLevel2Row.value.appId],
      rankType,
      page: targetState.page,
      limit: targetState.limit,
    });
    if (resp.code === 0) {
      targetState.rows =
        resp.data?.list?.map((item) => ({
          ...item,
          level3Loading: false,
        })) || [];
      targetState.total = resp.data?.totalCount || 0;
    }
  } finally {
    loadingRef.value = false;
  }
};

const handleLevel2TablePageChange = async (rankType) => {
  await fetchLevel2TableData(rankType);
};

const fetchLevel3TrendData = async () => {
  level3TrendLoading.value = true;
  try {
    const params = {
      rankLevel: "level3",
      startTime: searchForm.value.startTime,
      endTime: searchForm.value.endTime,
      appIdList: [selectedLevel2Row.value.appId],
    };
    if (selectedLevel3Target.value.isHost) {
      params.hostParam = selectedLevel3Target.value.name;
    } else {
      params.dstIpParam = selectedLevel3Target.value.name;
    }
    const resp = await appIdRankDataApi(params);
    level3TrendChartData.value = resp.code === 0 ? resp.data || [] : [];
  } finally {
    level3TrendLoading.value = false;
  }
};

const fetchLevel3TableData = async () => {
  level3TableLoading.value = true;
  try {
    const params = {
      rankLevel: "level3",
      startTime: searchForm.value.startTime,
      endTime: searchForm.value.endTime,
      appIdList: [selectedLevel2Row.value.appId],
      page: level3TableState.value.page,
      limit: level3TableState.value.limit,
    };
    if (selectedLevel3Target.value.isHost) {
      params.hostParam = selectedLevel3Target.value.name;
    } else {
      params.dstIpParam = selectedLevel3Target.value.name;
    }

    const resp = await appIdRankTrendDataApi(params);
    if (resp.code === 0) {
      level3TableState.value.rows =
        resp.data?.list?.slice().sort((a, b) => new Date(a.startTime) - new Date(b.startTime)) || [];
      level3TableState.value.total = resp.data?.totalCount || 0;
    }
  } finally {
    level3TableLoading.value = false;
  }
};

const handleOpenLevel2 = async (row) => {
  const startTime = dayjs(searchForm.value.startTime);
  const endTime = dayjs(searchForm.value.endTime);
  if (endTime.diff(startTime, "day") > 2) {
    ElMessage.warning("时间范围超过2天，无法下钻");
    return;
  }

  row.panelLoading = true;
  try {
    selectedLevel2Row.value = { ...row };
    selectedLevel2Name.value = row.name;
    level2DstIpTableState.value.page = 1;
    level2HostTableState.value.page = 1;
    await Promise.all([
      fetchLevel2Data(),
      fetchLevel2TableData("dstIp"),
      fetchLevel2TableData("host"),
    ]);
    level2PanelVisible.value = true;
    nextTick(() => {
      level2PanelRef.value?.scrollIntoView({ behavior: "smooth" });
    });
  } finally {
    row.panelLoading = false;
  }
};

const handleOpenLevel3 = async (isHost, row) => {
  row.level3Loading = true;
  try {
    selectedLevel3Target.value = {
      isHost,
      name: row.name,
    };
    level3PanelTitle.value = `${row.name} - ${isHost ? "域名" : "目的IP"}流量趋势`;
    level3TableState.value.page = 1;
    await Promise.all([fetchLevel3TrendData(), fetchLevel3TableData()]);
    level3PanelVisible.value = true;
    nextTick(() => {
      level3PanelRef.value?.scrollIntoView({ behavior: "smooth" });
    });
  } finally {
    row.level3Loading = false;
  }
};

const handleSearch = async () => {
  await Promise.all([fetchLevel1TableData(), fetchLevel1PieChartData()]);

  if (level2PanelVisible.value && selectedLevel2Row.value?.appId) {
    await Promise.all([
      fetchLevel2Data(),
      fetchLevel2TableData("dstIp"),
      fetchLevel2TableData("host"),
    ]);
  }

  if (level3PanelVisible.value && selectedLevel3Target.value.name) {
    await Promise.all([fetchLevel3TrendData(), fetchLevel3TableData()]);
  }
};

const handleExport = async () => {
  axios({
    method: "post",
    url: "/api/traffic/appId/export",
    data: searchForm.value,
    responseType: "blob",
  })
    .then((response) => {
      const filename = `app_id_rank_${+new Date()}.xlsx`;
      saveAs(
        new Blob([response.data], { type: response.headers["contnet-type"] }),
        filename,
      );
    })
    .catch((error) => {
      console.error("下载文件失败：", error);
    });
};

onMounted(async () => {
  await handleSearch();
});
</script>

<style scoped lang="scss">
.chart-card-title {
  display: block;
  width: 100%;
  text-align: center;
  font-size: 18px;
  font-weight: 700;
  letter-spacing: 0.5px;
  color: rgb(51 65 85);
}

.dark .chart-card-title {
  color: rgb(226 232 240);
}

.detail-card-title {
  margin-bottom: 12px;
  text-align: left;
  font-size: 18px;
  font-weight: 700;
  line-height: 28px;
  color: rgb(51 65 85);
}

.dark .detail-card-title {
  color: rgb(226 232 240);
}
</style>
