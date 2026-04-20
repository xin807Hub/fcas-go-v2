<template>
  <div>
    <ToolBar
      ref="toolbarRef"
      :search-el="searchEl"
      :delete-flag="true"
      toolbar-type="config"
      @handleSearch="searchFunc"
      @handleAdd="addFunc"
      @handleDelete="deleteAllFunc"
      @handleChange="handleSearchChange"
    />
    <CommonTable
      v-if="toolbarRef !== null"
      :key="loadKey"
      ref="tableRef"
      :table-params="toolbarRef.form"
      :table-columns="tableColumns"
      :operation="true"
      module-name="policy"
      module-api="control_policyListApi"
      @handleEdit="editFunc"
      @handleDelete="deleteFunc"
    />

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="900px">
      <el-form ref="formRef" :model="form" :rules="rules" label-width="120px">
        <el-form-item label="策略名称" prop="name">
          <el-input v-model="form.name" clearable />
        </el-form-item>

        <el-form-item label="链路" prop="link_ids">
          <el-select v-model="form.link_ids" multiple filterable clearable class="w-full">
            <el-option v-for="item in linkOptions" :key="item.value" :label="item.name" :value="item.value" />
          </el-select>
        </el-form-item>

        <el-form-item label="用户范围" prop="user_type">
          <el-select v-model="form.user_type" clearable class="w-56" @change="handleUserTypeChange">
            <el-option v-for="item in userTypeOptions" :key="item.value" :label="item.name" :value="item.value" />
          </el-select>
        </el-form-item>

        <el-form-item v-if="form.user_type === 1" label="用户群组" prop="user_crowd_group_id">
          <el-select v-model="form.user_crowd_group_id" filterable clearable class="w-full">
            <el-option v-for="item in userCrowdGroupOptions" :key="item.value" :label="item.name" :value="item.value" />
          </el-select>
        </el-form-item>

        <el-form-item v-if="form.user_type === 2" label="用户群" prop="user_crowd_id">
          <el-select v-model="form.user_crowd_id" filterable clearable class="w-full">
            <el-option v-for="item in userCrowdOptions" :key="item.value" :label="item.name" :value="item.value" />
          </el-select>
        </el-form-item>

        <el-form-item v-if="form.user_type === 3" label="用户" prop="user_id">
          <el-select v-model="form.user_id" filterable clearable class="w-full">
            <el-option v-for="item in userOptions" :key="item.value" :label="item.name" :value="item.value" />
          </el-select>
        </el-form-item>

        <el-row :gutter="16">
          <el-col :span="12">
            <el-form-item label="上行限速" prop="ul_flow_rate">
              <el-input v-model="form.ul_flow_rate" clearable placeholder="单位 Mbps" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="下行限速" prop="dl_flow_rate">
              <el-input v-model="form.dl_flow_rate" clearable placeholder="单位 Mbps" />
            </el-form-item>
          </el-col>
        </el-row>

        <el-form-item label="有效时间" prop="time_range">
          <el-date-picker
            v-model="form.time_range"
            type="datetimerange"
            value-format="YYYY-MM-DD HH:mm:ss"
            range-separator="至"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
            class="w-full"
          />
        </el-form-item>

        <el-form-item label="流控类型" prop="flow_ctrl_type">
          <el-select v-model="form.flow_ctrl_type" class="w-56" @change="handleFlowCtrlTypeChange">
            <el-option v-for="item in flowCtrlTypeOptions" :key="item.value" :label="item.name" :value="item.value" />
          </el-select>
        </el-form-item>

        <template v-if="form.flow_ctrl_type === 1">
          <el-form-item label="应用大类" prop="app_type_id">
            <el-select v-model="form.app_type_id" filterable clearable class="w-full" @change="handleAppTypeChange">
              <el-option v-for="item in appTypeOptions" :key="item.value" :label="item.name" :value="item.value" />
            </el-select>
          </el-form-item>
          <el-form-item label="应用小类" prop="app_id">
            <el-select v-model="form.app_id" filterable clearable class="w-full">
              <el-option v-for="item in appIdOptions" :key="item.value" :label="item.name" :value="item.value" />
            </el-select>
          </el-form-item>
        </template>

        <template v-else>
          <el-row :gutter="16">
            <el-col :span="12">
              <el-form-item label="目的IP" prop="dst_ip">
                <el-input v-model="dstForm.dst_ip" clearable placeholder="IP 或 开始IP-结束IP" />
              </el-form-item>
            </el-col>
            <el-col :span="10">
              <el-form-item label="目的端口" prop="dst_port">
                <el-input v-model="dstForm.dst_port" clearable placeholder="端口或开始端口-结束端口" />
              </el-form-item>
            </el-col>
            <el-col :span="2" class="flex items-center">
              <el-button type="primary" icon="Plus" @click="addIpData" />
            </el-col>
          </el-row>

          <el-form-item label="目的地址列表">
            <el-table :data="ipDataList" size="small">
              <el-table-column prop="start_ip" label="开始IP" />
              <el-table-column prop="end_ip" label="结束IP" />
              <el-table-column prop="dst_port" label="端口" />
              <el-table-column label="操作" width="80">
                <template #default="scope">
                  <el-button link type="danger" icon="Delete" @click="removeIpData(scope.$index)" />
                </template>
              </el-table-column>
            </el-table>
          </el-form-item>
        </template>

        <el-form-item label="周期类型" prop="period_type">
          <el-select v-model="form.period_type" clearable class="w-56" @change="handlePeriodTypeChange">
            <el-option v-for="item in periodTypeOptions" :key="item.value" :label="item.name" :value="item.value" />
          </el-select>
        </el-form-item>

        <template v-if="form.period_type === 1">
          <el-row :gutter="16">
            <el-col :span="10">
              <el-form-item label="开始周期">
                <el-input v-model="periodForm.start_time" placeholder="HH:mm:ss" clearable />
              </el-form-item>
            </el-col>
            <el-col :span="10">
              <el-form-item label="结束周期">
                <el-input v-model="periodForm.end_time" placeholder="HH:mm:ss" clearable />
              </el-form-item>
            </el-col>
            <el-col :span="4" class="flex items-center">
              <el-button type="primary" icon="Plus" @click="addDayPeriod" />
            </el-col>
          </el-row>
          <el-form-item label="日周期列表">
            <el-table :data="dayPeriods" size="small">
              <el-table-column prop="start_time" label="开始周期" />
              <el-table-column prop="end_time" label="结束周期" />
              <el-table-column label="操作" width="80">
                <template #default="scope">
                  <el-button link type="danger" icon="Delete" @click="removeDayPeriod(scope.$index)" />
                </template>
              </el-table-column>
            </el-table>
          </el-form-item>
        </template>

        <template v-if="form.period_type === 2">
          <el-row :gutter="16">
            <el-col :span="10">
              <el-form-item label="开始周期">
                <el-select v-model="periodForm.start_week" clearable class="w-full">
                  <el-option v-for="item in weekOptions" :key="item.value" :label="item.name" :value="item.value" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="10">
              <el-form-item label="结束周期">
                <el-select v-model="periodForm.end_week" clearable class="w-full">
                  <el-option v-for="item in weekOptions" :key="item.value" :label="item.name" :value="item.value" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="4" class="flex items-center">
              <el-button type="primary" icon="Plus" @click="addWeekPeriod" />
            </el-col>
          </el-row>
          <el-form-item label="周周期列表">
            <el-table :data="weekPeriods" size="small">
              <el-table-column prop="start_week" label="开始周期" />
              <el-table-column prop="end_week" label="结束周期" />
              <el-table-column label="操作" width="80">
                <template #default="scope">
                  <el-button link type="danger" icon="Delete" @click="removeWeekPeriod(scope.$index)" />
                </template>
              </el-table-column>
            </el-table>
          </el-form-item>
        </template>

        <el-form-item label="备注" prop="remark">
          <el-input v-model="form.remark" type="textarea" :rows="3" />
        </el-form-item>
      </el-form>

      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitForm">提交</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { computed, onBeforeMount, ref } from "vue";
import ToolBar from "@/components/commonTable/toolbar.vue";
import CommonTable from "@/components/commonTable/index.vue";
import { ElMessage } from "element-plus";
import { formatDate } from "@/utils/format";
import { control_createPolicyApi, control_deletePolicyApi } from "@/api/policy";
import { dictInfoApi } from "@/api/traffic";
import { userCrowdApi, userCrowdGroupApi, userListApi } from "@/api/userMng";
import useLinkStore from "@/pinia/modules/link";

const toolbarRef = ref(null);
const tableRef = ref(null);
const formRef = ref(null);
const dialogVisible = ref(false);
const dialogTitle = ref("新建");
const loadKey = ref(0);
const isEdit = ref(false);

const flowCtrlTypeOptions = [
  { name: "应用大小类", value: 1 },
  { name: "目的地址", value: 2 },
];
const periodTypeOptions = [
  { name: "每日", value: 1 },
  { name: "每周", value: 2 },
];
const userTypeOptions = [
  { name: "用户群组", value: 1 },
  { name: "用户群", value: 2 },
  { name: "用户", value: 3 },
];
const weekOptions = [
  { name: "周一", value: 1 },
  { name: "周二", value: 2 },
  { name: "周三", value: 3 },
  { name: "周四", value: 4 },
  { name: "周五", value: 5 },
  { name: "周六", value: 6 },
  { name: "周日", value: 7 },
];

const linkOptions = ref([]);
const userOptions = ref([]);
const userCrowdOptions = ref([]);
const userCrowdGroupOptions = ref([]);
const appTypeOptions = ref([]);
const appTypeTree = ref([]);

const searchEl = ref([
  { name: "策略名称", field: "policyName" },
  { name: "用户范围", field: "userType", options: userTypeOptions },
  { name: "用户群组", field: "userCrowdGroupId", options: userCrowdGroupOptions.value },
  { name: "用户群", field: "userCrowdId", options: userCrowdOptions.value },
  { name: "用户", field: "userId", options: userOptions.value },
  { name: "应用大类", field: "appTypeId", options: appTypeOptions.value },
  { name: "应用小类", field: "appId", options: [] },
]);

const tableColumns = ref([
  { name: "策略ID", field: "id", width: 80 },
  { name: "策略名称", field: "name", width: 180 },
  { name: "用户群组", field: "user_crowd_group_name", width: 140 },
  { name: "用户群", field: "user_crowd_name", width: 140 },
  { name: "用户", field: "user_name", width: 140 },
  { name: "应用大类", field: "app_type_name", width: 140 },
  { name: "应用小类", field: "app_name", width: 140 },
  { name: "上行限速", field: "ul_flow_rate", width: 120 },
  { name: "下行限速", field: "dl_flow_rate", width: 120 },
  { name: "生效时间", field: "start_time", width: 170 },
  { name: "失效时间", field: "end_time", width: 170 },
  { name: "上行流速", field: "up_traffic_speed", width: 120 },
  { name: "下行流速", field: "dn_traffic_speed", width: 120 },
  { name: "上行通过流速", field: "up_pass_speed", width: 130 },
  { name: "下行通过流速", field: "dn_pass_speed", width: 130 },
  { name: "上行丢弃流速", field: "up_discard_speed", width: 130 },
  { name: "下行丢弃流速", field: "dn_dis_card_speed", width: 130 },
]);

const defaultForm = () => ({
  id: 0,
  name: "",
  link_ids: [],
  user_type: null,
  user_crowd_group_id: null,
  user_crowd_id: null,
  user_id: null,
  ul_flow_rate: "",
  dl_flow_rate: "",
  time_range: [],
  flow_ctrl_type: 1,
  app_type_id: null,
  app_id: null,
  period_type: null,
  remark: "",
});

const form = ref(defaultForm());
const dstForm = ref({
  dst_ip: "",
  dst_port: "",
});
const periodForm = ref({
  start_time: "",
  end_time: "",
  start_week: null,
  end_week: null,
});
const ipDataList = ref([]);
const dayPeriods = ref([]);
const weekPeriods = ref([]);

const appIdOptions = computed(() => {
  if (!form.value.app_type_id) {
    return [];
  }
  const current = appTypeTree.value.find((item) => item.id === form.value.app_type_id);
  return (current?.children || []).map((item) => ({
    value: item.id,
    name: item.name,
  }));
});

const searchAppIdOptions = computed(() => {
  const appTypeId = toolbarRef.value?.form?.appTypeId;
  if (!appTypeId) {
    return [];
  }
  const current = appTypeTree.value.find((item) => item.id === appTypeId);
  return (current?.children || []).map((item) => ({
    value: item.id,
    name: item.name,
  }));
});

const validateRate = (rule, value, callback) => {
  const other = rule.field === "ul_flow_rate" ? form.value.dl_flow_rate : form.value.ul_flow_rate;
  if ((value === "" || value === null) && (other === "" || other === null)) {
    callback(new Error("上行限速与下行限速至少填写一项"));
    return;
  }
  if (value === "" || value === null) {
    callback();
    return;
  }
  if (!/^\d+$/.test(String(value))) {
    callback(new Error("限速必须为整数"));
    return;
  }
  callback();
};

const rules = ref({
  name: [{ required: true, message: "策略名称不能为空", trigger: "blur" }],
  link_ids: [{ required: true, message: "链路不能为空", trigger: "change" }],
  time_range: [{ required: true, message: "有效时间不能为空", trigger: "change" }],
  ul_flow_rate: [{ validator: validateRate, trigger: "blur" }],
  dl_flow_rate: [{ validator: validateRate, trigger: "blur" }],
});

function syncSearchOptions() {
  const searchMap = {
    userCrowdGroupId: userCrowdGroupOptions.value,
    userCrowdId: userCrowdOptions.value,
    userId: userOptions.value,
    appTypeId: appTypeOptions.value,
    appId: searchAppIdOptions.value,
  };
  searchEl.value = searchEl.value.map((item) => ({
    ...item,
    options: searchMap[item.field] ?? item.options,
  }));
}

function resetForm() {
  form.value = defaultForm();
  dstForm.value = { dst_ip: "", dst_port: "" };
  periodForm.value = { start_time: "", end_time: "", start_week: null, end_week: null };
  ipDataList.value = [];
  dayPeriods.value = [];
  weekPeriods.value = [];
  formRef.value?.resetFields();
}

function handleUserTypeChange(value) {
  form.value.user_type = value;
  form.value.user_crowd_group_id = null;
  form.value.user_crowd_id = null;
  form.value.user_id = null;
}

function handleFlowCtrlTypeChange(value) {
  form.value.flow_ctrl_type = value;
  if (value === 1) {
    ipDataList.value = [];
  } else {
    form.value.app_type_id = null;
    form.value.app_id = null;
  }
}

function handleAppTypeChange() {
  form.value.app_id = null;
}

function handlePeriodTypeChange(value) {
  form.value.period_type = value;
  dayPeriods.value = [];
  weekPeriods.value = [];
  periodForm.value = { start_time: "", end_time: "", start_week: null, end_week: null };
}

function handleSearchChange(val, field) {
  if (field === "appTypeId" && toolbarRef.value) {
    toolbarRef.value.form.appId = null;
    syncSearchOptions();
  }
  if (field === "userType" && toolbarRef.value) {
    toolbarRef.value.form.userCrowdGroupId = null;
    toolbarRef.value.form.userCrowdId = null;
    toolbarRef.value.form.userId = null;
  }
}

function searchFunc() {
  loadKey.value += 1;
}

function addFunc() {
  isEdit.value = false;
  dialogTitle.value = "新建";
  resetForm();
  dialogVisible.value = true;
}

async function deleteAllFunc() {
  const checkedRows = tableRef.value?.checkedRows || [];
  const ids = checkedRows.map((item) => item.id);
  if (ids.length === 0) {
    ElMessage.warning("请先选择要删除的策略");
    return;
  }
  const res = await control_deletePolicyApi({ ids });
  if (res.code === 0) {
    ElMessage.success("删除成功");
    loadKey.value += 1;
  }
}

function parseLinkIds(linkIDs) {
  const values = String(linkIDs || "")
    .split(",")
    .map((item) => Number(item))
    .filter((item) => Number.isFinite(item) && item > 0);
  return values
    .map((value) => {
      const byID = linkOptions.value.find((item) => item.value === value);
      if (byID) {
        return byID.value;
      }
      const byVlan = linkOptions.value.find((item) => item.lineVlan === value);
      return byVlan?.value;
    })
    .filter((value) => value !== undefined);
}

function parsePeriods(row) {
  dayPeriods.value = [];
  weekPeriods.value = [];
  if (!row.policy_period) {
    return;
  }
  const items = String(row.policy_period)
    .split(",")
    .map((item) => item.trim())
    .filter(Boolean);
  if (row.period_type === 1) {
    dayPeriods.value = items
      .map((item) => item.split("-"))
      .filter((item) => item.length === 2)
      .map((item) => ({ start_time: item[0], end_time: item[1] }));
  }
  if (row.period_type === 2) {
    weekPeriods.value = items
      .map((item) => item.split("-"))
      .filter((item) => item.length === 2)
      .map((item) => ({ start_week: Number(item[0]), end_week: Number(item[1]) }));
  }
}

function editFunc(str, row) {
  if (row.end_time && new Date(row.end_time).getTime() <= Date.now()) {
    ElMessage.warning("已过期策略不允许编辑");
    return;
  }

  isEdit.value = true;
  dialogTitle.value = "编辑";
  resetForm();

  form.value = {
    id: row.id,
    name: row.name || "",
    link_ids: parseLinkIds(row.link_ids),
    user_type: row.user_type || null,
    user_crowd_group_id: row.user_crowd_group_id || null,
    user_crowd_id: row.user_crowd_id || null,
    user_id: row.user_id || null,
    ul_flow_rate: row.ul_flow_rate ?? "",
    dl_flow_rate: row.dl_flow_rate ?? "",
    time_range: [row.start_time, row.end_time],
    flow_ctrl_type: row.flow_ctrl_type || 1,
    app_type_id: row.app_type_id || null,
    app_id: row.app_id || null,
    period_type: row.period_type || null,
    remark: row.remark || "",
  };

  ipDataList.value = Array.isArray(row.ip_data) ? row.ip_data.map((item) => ({ ...item })) : [];
  parsePeriods(row);
  dialogVisible.value = true;
}

async function deleteFunc(str, row) {
  const res = await control_deletePolicyApi({ ids: [row.id] });
  if (res.code === 0) {
    ElMessage.success("删除成功");
    loadKey.value += 1;
  }
}

function validateIPv4(value) {
  return /^(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)$/.test(value);
}

function validatePort(port) {
  return /^([0-9]|[1-9]\d{1,3}|[1-5]\d{4}|6[0-4]\d{3}|65[0-4]\d{2}|655[0-2]\d|6553[0-5])$/.test(port);
}

function addIpData() {
  const rawIP = String(dstForm.value.dst_ip || "").trim();
  const rawPort = String(dstForm.value.dst_port || "").trim();
  if (!rawIP && !rawPort) {
    ElMessage.warning("IP地址和端口至少填写一项");
    return;
  }

  let startIP = "";
  let endIP = "";
  if (rawIP) {
    if (rawIP.includes("-")) {
      const parts = rawIP.split("-");
      if (parts.length !== 2 || !validateIPv4(parts[0]) || !validateIPv4(parts[1])) {
        ElMessage.error("请输入正确的 IPv4 地址范围");
        return;
      }
      startIP = parts[0];
      endIP = parts[1];
    } else {
      if (!validateIPv4(rawIP)) {
        ElMessage.error("请输入正确的 IPv4 地址");
        return;
      }
      startIP = rawIP;
      endIP = rawIP;
    }
  }

  if (rawPort) {
    if (rawPort.includes("-")) {
      const parts = rawPort.split("-");
      if (parts.length !== 2 || !validatePort(parts[0]) || !validatePort(parts[1]) || Number(parts[0]) >= Number(parts[1])) {
        ElMessage.error("请输入正确的端口范围");
        return;
      }
    } else if (!validatePort(rawPort)) {
      ElMessage.error("请输入正确的端口");
      return;
    }
  }

  ipDataList.value.push({
    start_ip: startIP,
    end_ip: endIP,
    dst_port: rawPort,
  });
  dstForm.value = { dst_ip: "", dst_port: "" };
}

function removeIpData(index) {
  ipDataList.value.splice(index, 1);
}

function addDayPeriod() {
  const start = String(periodForm.value.start_time || "").trim();
  const end = String(periodForm.value.end_time || "").trim();
  const regex = /^([0-1]\d|2[0-3]):[0-5]\d:[0-5]\d$/;
  if (!regex.test(start) || !regex.test(end)) {
    ElMessage.error("周期时间格式必须为 HH:mm:ss");
    return;
  }
  if (start > end) {
    ElMessage.error("开始周期不能晚于结束周期");
    return;
  }
  dayPeriods.value.push({ start_time: start, end_time: end });
  periodForm.value.start_time = "";
  periodForm.value.end_time = "";
}

function removeDayPeriod(index) {
  dayPeriods.value.splice(index, 1);
}

function addWeekPeriod() {
  const start = Number(periodForm.value.start_week);
  const end = Number(periodForm.value.end_week);
  if (!start || !end) {
    ElMessage.error("请选择完整的周周期");
    return;
  }
  if (start > end) {
    ElMessage.error("开始周期不能晚于结束周期");
    return;
  }
  weekPeriods.value.push({ start_week: start, end_week: end });
  periodForm.value.start_week = null;
  periodForm.value.end_week = null;
}

function removeWeekPeriod(index) {
  weekPeriods.value.splice(index, 1);
}

function buildPolicyPeriod() {
  if (form.value.period_type === 1) {
    return dayPeriods.value.map((item) => `${item.start_time}-${item.end_time}`).join(",");
  }
  if (form.value.period_type === 2) {
    return weekPeriods.value.map((item) => `${item.start_week}-${item.end_week}`).join(",");
  }
  return "";
}

async function submitForm() {
  await formRef.value?.validate();

  const payload = {
    id: form.value.id || undefined,
    name: form.value.name,
    user_type: form.value.user_type,
    user_crowd_group_id: form.value.user_type === 1 ? form.value.user_crowd_group_id : null,
    user_crowd_id: form.value.user_type === 2 ? form.value.user_crowd_id : null,
    user_id: form.value.user_type === 3 ? form.value.user_id : null,
    ul_flow_rate: form.value.ul_flow_rate === "" ? null : Number(form.value.ul_flow_rate),
    dl_flow_rate: form.value.dl_flow_rate === "" ? null : Number(form.value.dl_flow_rate),
    start_time: formatDate(form.value.time_range[0]),
    end_time: formatDate(form.value.time_range[1]),
    remark: form.value.remark,
    flow_ctrl_type: form.value.flow_ctrl_type,
    app_type_id: form.value.flow_ctrl_type === 1 ? form.value.app_type_id : null,
    app_id: form.value.flow_ctrl_type === 1 ? form.value.app_id : null,
    ip_data: form.value.flow_ctrl_type === 2 ? ipDataList.value : [],
    period_type: form.value.period_type,
    policy_period: buildPolicyPeriod(),
    link_ids: form.value.link_ids.join(","),
  };

  const res = await control_createPolicyApi(payload);
  if (res.code === 0) {
    ElMessage.success("操作成功");
    dialogVisible.value = false;
    loadKey.value += 1;
  }
}

async function loadLinks() {
  const linkStore = useLinkStore();
  const links = await linkStore.getLink();
  linkOptions.value = links.map((item) => ({
    value: item.id,
    name: item.lineName,
    lineVlan: item.lineVlan,
  }));
}

async function loadUsers() {
  const [userRes, crowdRes, groupRes] = await Promise.all([
    userListApi({ page: 1, limit: 999 }),
    userCrowdApi({ page: 1, limit: 999, key: "" }),
    userCrowdGroupApi({ page: 1, limit: 999, key: "" }),
  ]);

  if (userRes.code === 0) {
    userOptions.value = userRes.data.list.map((item) => ({
      value: item.id,
      name: item.userName,
    }));
  }
  if (crowdRes.code === 0) {
    userCrowdOptions.value = crowdRes.data.list.map((item) => ({
      value: item.id,
      name: item.crowdName,
    }));
  }
  if (groupRes.code === 0) {
    userCrowdGroupOptions.value = groupRes.data.list.map((item) => ({
      value: item.id,
      name: item.group_name,
    }));
  }
}

async function loadAppTypes() {
  const [typeRes, treeRes] = await Promise.all([dictInfoApi("appType"), dictInfoApi("appTypeIdTree")]);
  if (typeRes.code === 0) {
    appTypeOptions.value = typeRes.data.map((item) => ({
      value: item.id,
      name: item.name,
    }));
  }
  if (treeRes.code === 0) {
    appTypeTree.value = treeRes.data || [];
  }
}

onBeforeMount(async () => {
  await Promise.all([loadLinks(), loadUsers(), loadAppTypes()]);
  syncSearchOptions();
});
</script>
