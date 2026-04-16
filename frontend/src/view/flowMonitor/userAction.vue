<template>
  <div class="flex flex-col gap-2">
    <el-card>
      <div class="inline-flex w-full gap-2 justify-self-center">
        <date-picker
          v-model:start-time="searchForm.startTime"
          v-model:end-time="searchForm.endTime"
          max-time-range="1h"
          :shortcuts="[]"
        />

        <el-select
          v-model="searchForm.dataType"
          clearable
          size="small"
          placeholder="请选择目的IP或应用小类"
          style="width: 200px"
        >
          <el-option
            v-for="item in dataTypeOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>

        <el-input
          v-model="searchForm.srcIp"
          placeholder="请输入源IP"
          clearable
          size="small"
          style="width: 200px"
        />

        <el-button
          icon="Search"
          type="primary"
          size="small"
          class="button"
          @click="onQuery"
        >
          查询
        </el-button>
      </div>
    </el-card>

    <div class="rounded-md bg-white p-2 shadow-md">
      <el-table
        v-loading="loading"
        size="small"
        height="720"
        :data="tableData.list"
        highlight-current-row
      >
        <el-table-column
          type="index"
          label="序号"
          width="60"
          align="center"
        />
        <el-table-column
          prop="name"
          :label="dataTypeOptions.find((item) => item.value === searchForm.dataType)?.label || '名称'"
          width="380"
          show-overflow-tooltip
        />
        <el-table-column
          prop="upByte"
          label="上行流量"
          width="380"
          show-overflow-tooltip
          :formatter="(row, column, cellValue) => formatSize(cellValue)"
        />
        <el-table-column
          prop="dnByte"
          label="下行流量"
          width="380"
          show-overflow-tooltip
          :formatter="(row, column, cellValue) => formatSize(cellValue)"
        />
        <el-table-column
          prop="totalByte"
          label="总流量"
          width="380"
          show-overflow-tooltip
          :formatter="(row, column, cellValue) => formatSize(cellValue)"
        />
      </el-table>

      <el-pagination
        v-model:current-page="pagination.page"
        v-model:page-size="pagination.limit"
        class="mt-2"
        background
        size="small"
        :total="tableData.total"
        :page-sizes="[10, 50, 100]"
        layout="sizes, prev, pager, next, total"
        @current-change="onPageChange"
        @size-change="onPageChange"
      />
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref } from "vue";
import dayjs from "dayjs";
import { ElMessage } from "element-plus";
import { userActionPageDataApi } from "@/api/traffic";
import DatePicker from "@/components/searchItem/DatePicker.vue";
import { formatSize } from "@/utils/format";

const dataTypeOptions = [
  { label: "目的IP", value: "dst_ip" },
  { label: "应用小类", value: "app_id" },
];

const searchForm = ref({
  startTime: dayjs().subtract(1, "hour").format("YYYY-MM-DD HH:mm:ss"),
  endTime: dayjs().format("YYYY-MM-DD HH:mm:ss"),
  dataType: "dst_ip",
  srcIp: "",
});

const tableData = ref({
  list: [],
  total: 0,
});

const pagination = ref({
  page: 1,
  limit: 20,
});

const loading = ref(false);

const resetTableData = () => {
  tableData.value.list = [];
  tableData.value.total = 0;
};

const validateQuery = () => {
  if (!searchForm.value.dataType) {
    ElMessage.warning("请选择目的IP或应用小类");
    return false;
  }
  if (searchForm.value.srcIp.trim() === "") {
    ElMessage.warning("请输入源IP");
    return false;
  }
  return true;
};

const fetchTableData = async () => {
  const params = {
    ...searchForm.value,
    ...pagination.value,
  };

  try {
    loading.value = true;
    const resp = await userActionPageDataApi(params);
    tableData.value.list = resp?.data?.list || [];
    tableData.value.total = resp?.data?.totalCount || 0;
  } finally {
    loading.value = false;
  }
};

const onQuery = async () => {
  pagination.value.page = 1;
  if (!validateQuery()) {
    resetTableData();
    return;
  }
  await fetchTableData();
};

const onPageChange = async () => {
  if (!validateQuery()) {
    resetTableData();
    return;
  }
  await fetchTableData();
};

onMounted(() => {
  resetTableData();
});
</script>

<style scoped lang="scss"></style>
