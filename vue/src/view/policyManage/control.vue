<template>
  <div>
    <!--    工具栏    -->
    <ToolBar
      ref="toolbarParam"
      :search-el="searchEl"
      :delete-flag="deleteFlag"
      toolbar-type="config"
      @handleSearch="searchFunc"
      @handleAdd="addFunc"
      @handleDelete="deleteAllFunc"
      @handleChange="handleKey"
    />
    <div class="gva-table-box">
      <!--    设备列表    -->
      <CommonTable
        v-if="toolbarParam!==null"
        :key="loadKey"
        ref="tableParam"
        :table-params="toolbarParam.form"
        :table-columns="tableColumns"
        :operation="true"
        module-name="policy"
        module-api="control_policyListApi"
        @handleEdit="editFunc"
        @handleDelete="deleteFunc"
      />
      <!--  配置二级   -->
      <!--      class="w-176 h-auto"-->
      <el-dialog
        v-model="configVisible"
        :title="configTitle"
        class="w-198 h-auto"
      >
        <!--          二级页面-->
        <el-form
          ref="formRef"
          :model="form"
          :rules="rules"
        >
          <!--            策略名称-->
          <el-form-item
            :label="configtable[0].name"
            label-width="120px"
            :prop="configtable[0].field"
          >
            <el-input
              v-model="form[configtable[0].field]"
              :placeholder="`请输入${configtable[0].name}`"
              autocomplete="off"
              class="w-112"
            />
          </el-form-item>
          <!--            链路-->
          <el-form-item
            :label="configtable[1].name"
            label-width="120px"
            :prop="configtable[1].field"
          >
            <el-select
              v-model="form[configtable[1].field]"
              :placeholder="`请选择${configtable[1].name}`"
              multiple
              filterable
              class="w-112"
              @change="handlechange($event, configtable[1].field)"
            >
              <el-option
                v-for="item in configtable[1].options"
                :label="item.name"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          <!--            用户类型-->
          <el-form-item
            :label="configtable[2].name"
            label-width="120px"
            :prop="configtable[2].field"
          >
            <el-select
              v-model="form[configtable[2].field]"
              :placeholder="`请选择${configtable[2].name}`"
              filterable
              class="w-112"
              @change="handlechange($event, configtable[2].field)"
            >
              <el-option
                v-for="item in configtable[2].options"
                :label="item.name"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          <!--            上行限速-->
          <el-form-item
            :label="configtable[3].name"
            label-width="120px"
            :prop="configtable[3].field"
          >
            <el-input
              v-model="form[configtable[3].field]"
              :placeholder="`请输入${configtable[3].name}Mbps`"
              autocomplete="off"
              class="w-112"
            />
          </el-form-item>
          <!--            下行限速-->
          <el-form-item
            :label="configtable[4].name"
            label-width="120px"
            :prop="configtable[4].field"
          >
            <el-input
              v-model="form[configtable[4].field]"
              :placeholder="`请输入${configtable[4].name}Mbps`"
              autocomplete="off"
              class="w-112"
            />
          </el-form-item>
          <!--            有效时间-->
          <el-form-item
            :label="configtable[5].name"
            label-width="120px"
            style="width:510px;"
            :prop="configtable[5].field"
          >
            <el-date-picker
              v-model="form[configtable[5].field]"
              type="datetimerange"
              range-separator="至"
              start-placeholder="生效时间"
              end-placeholder="失效时间"
              size="small"
            />
          </el-form-item>
          <!--            流控类型-->
          <el-form-item
            :label="configtable[8].name"
            label-width="120px"
            :prop="configtable[8].field"
          >
            <el-select
              v-model="form[configtable[8].field]"
              :placeholder="`请选择${configtable[8].name}`"
              class="w-112"
              @change="handlechange($event, configtable[8].field)"
            >
              <el-option
                v-for="item in configtable[8].options"
                :label="item.name"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          <!--            应用大类-->
          <el-form-item
            v-if="firstSelectVisible"
            :label="configtable[12].name"
            label-width="120px"
            :prop="configtable[12].field"
          >
            <el-select
              v-model="form[configtable[12].field]"
              :placeholder="`请选择${configtable[12].name}`"
              filterable
              class="w-112"
              @change="handlechange($event, configtable[12].field)"
            >
              <el-option
                v-for="item in configtable[12].options"
                :label="item.name"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          <!--            应用小类-->
          <el-form-item
             v-if="firstSelectVisible"
            :label="configtable[13].name"
            label-width="120px"
            :prop="configtable[13].field"
          >
            <el-select
              v-model="form[configtable[13].field]"
              :placeholder="`请选择${configtable[13].name}`"
              filterable
              class="w-112"
              @change="handlechange($event, configtable[13].field)"
            >
              <el-option
                v-for="item in configtable[13].options"
                :label="item.name"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          <!--            IP-->
          <el-form-item
            v-if="firstSelectIPVisible"
            :label="configtable[9].name"
            label-width="120px"
            :prop="configtable[9].field"
          >
            <el-input
              v-model="form[configtable[9].field]"
              :placeholder="`IP地址或者开始ip-结束ip`"
              autocomplete="off"
              class="w-112"
            />
          </el-form-item>
          <el-form-item
            v-if="firstSelectIPVisible"
            :label="configtable[10].name"
            label-width="120px"
            :prop="configtable[10].field"
          >
            <el-input
              v-model="form[configtable[10].field]"
              :placeholder="`请输入${configtable[10].name}`"
              autocomplete="off"
              class="w-112"
            />
            <el-button type="primary" icon="Plus" class="button" size="small" @click="handleAddToIPTable" />
<!--            目的地址-->
          </el-form-item>
          <el-form
            v-if="firstSelectIPVisible "
            :label="configtable[11].name"
            label-width="120px"
          >
            <el-form-item>
<!--              <el-button type="primary" icon="Plus" class="button" size="small" @click="handleAddToTable" />-->
              <el-table :data="IpData">
                <el-table-column prop="start_ip" label="目的起始ip" width="150" />
                <el-table-column prop="end_ip" label="目的结束ip" width="150" />
                <el-table-column prop="dst_port" label="目的端口" width="150" />
                <el-table-column prop="option" label="操作" fixed="right" width="100">
                  <template #default="scope">
                    <el-button link type="danger" icon="Delete" title="删除" size="large" @click="handleDelFromTable(scope.row, scope.$index)" />
                  </template>
                </el-table-column>
              </el-table>
            </el-form-item>
          </el-form>
          <!--            周期类型-->
          <el-form-item
            :label="configtable[14].name"
            label-width="120px"
            :prop="configtable[14].field"
          >
            <el-select
              v-model="form[configtable[14].field]"
              :placeholder="`请选择${configtable[14].name}`"
              class="w-112"
              @change="handlechange($event, configtable[14].field)"
            >
              <el-option
                v-for="item in configtable[14].options"
                :label="item.name"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
<!--          时间选择-->
<!--          每日-->
          <el-form-item
              v-if="TimeVisible"
              label="开始周期"
              label-width="120px"
              :prop="configtable[17].field"
          >
            <el-input
                v-model="form[configtable[17].field]"
                placeholder="格式HH:mm:ss"
                autocomplete="off"
                class="w-112"
            />
          </el-form-item>
          <el-form-item
              v-if="TimeVisible"
              label="结束周期"
              label-width="120px"
              :prop="configtable[18].field"
          >
            <el-input
                v-model="form[configtable[18].field]"
                placeholder="格式HH:mm:ss"
                autocomplete="off"
                class="w-112"
            />
            <el-button type="primary" icon="Plus" class="button" size="small" @click="handleAddToTimeTable" />
            <!--            表-->
          </el-form-item>
          <el-form
              v-if="TimeVisible"
              label="周期时间"
              label-width="120px"
              :prop="TimeData"
          >
            <el-form-item>
              <el-table :data="TimeData">
                <el-table-column prop="start_time" label="开始周期" width="150" />
                <el-table-column prop="end_time" label="结束周期" width="150" />
                <el-table-column prop="option" label="操作" fixed="right" width="100">
                  <template #default="scope">
                    <el-button link type="danger" icon="Delete" title="删除" size="large" @click="handleDelFromTable(scope.row, scope.$index)" />
                  </template>
                </el-table-column>
              </el-table>
            </el-form-item>
          </el-form>
          <!--          每周-->
          <el-form-item
              v-if="WeekVisible"
              label="开始周期"
              label-width="120px"
              :prop="configtable[19].field"
          >
            <el-select
                v-model="form[configtable[19].field]"
                placeholder="请选择"
                class="w-112"
                @change="handlechange($event, configtable[19].field)"
            >
              <el-option
                  v-for="item in configtable[19].options"
                  :label="item.name"
                  :value="item.value"
              />
            </el-select>
          </el-form-item>
          <el-form-item
              v-if="WeekVisible"
              label="结束周期"
              label-width="120px"
              :prop="configtable[20].field"
          >
            <el-select
                v-model="form[configtable[20].field]"
                placeholder="请选择"
                class="w-112"
                @change="handlechange($event, configtable[20].field)"
            >
              <el-option
                  v-for="item in configtable[20].options"
                  :label="item.name"
                  :value="item.value"
              />
            </el-select>
            <el-button type="primary" icon="Plus" class="button" size="small" @click="handleAddToWeekTable" />
            <!--            表-->
          </el-form-item>
          <el-form
              v-if="WeekVisible"
              label="周期时间"
              label-width="120px"
              :prop="WeekData"
          >
            <el-form-item>
              <el-table :data="WeekData">
                <el-table-column prop="start_week" label="开始周期" width="150" />
                <el-table-column prop="end_week" label="结束周期" width="150" />
                <el-table-column prop="option" label="操作" fixed="right" width="100">
                  <template #default="scope">
                    <el-button link type="danger" icon="Delete" title="删除" size="large" @click="handleDelFromTable(scope.row, scope.$index)" />
                  </template>
                </el-table-column>
              </el-table>
            </el-form-item>
          </el-form>
          <!--            备注-->
          <el-form-item
              v-if="configtable[16].name === '备注'"
              :label="configtable[16].name"
              :prop="configtable[16].field"
              label-width="120px"
          >
            <el-input
                v-model="form[configtable[16].field]"
                :placeholder="`请输入${configtable[16].name}`"
                :type="configtable[16].type = 'textarea' "
                autocomplete="off"
                class="w-112"
            />
          </el-form-item>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="configVisible = !configVisible;clearselect">取消</el-button>
            <el-button type="primary" @click="handleSubmit">提交</el-button>
          </span>
        </template>
      </el-dialog>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import {onBeforeMount, ref} from "vue"
import ToolBar from "@/components/commonTable/toolbar.vue";
import CommonTable from "@/components/commonTable/index.vue";
import {ElMessage} from "element-plus";
import {validatePort} from "@/utils/validate";
import {dictInfoApi, userCrowdGroupTreeApi} from "@/api/traffic";
import useLinkStore from "@/pinia/modules/link";
import {control_createPolicyApi, control_deletePolicyApi} from "@/api/policy"
import {formatDate} from "@/utils/format";
import {formatTimeToStr} from "@/utils/date";

// 定义工具栏配置
const formRef = ref(null)
const form = ref({
  id: null
})
const configVisible = ref(false)
const configTitle = ref("新增")
const toolbarParam = ref(null)
const  deleteFlag = ref(true)
const linkOptions = ref([])
const LinkStore = useLinkStore()
const appTypeIdOptions = ref([])//应用大类
const appTypeIdTreeOptions = ref([])//应用小类
const userOptions = ref([])//用户群组树
const searchEl = ref([
  { name: "请输入策略名称", field: "policyName" },
  { name: "用户", field: "uerId",options:[]},
  { name: "应用大类", field: "appTypeId",options:[]},
  // { name: "应用小类", field: "appId",options:[]},
])

//应用小类下拉框显示值
const secondSelectVisible = ref(false)
const firstSelectVisible = ref(false)
const firstSelectIPVisible = ref(false)
const TimeVisible = ref(false)
const WeekVisible = ref(false)

const isEdit = ref(false)


// 列表查询逻辑
const loadKey = ref(0)
const tableParam = ref(null)
const tableColumns = ref([
  { name: "策略ID", field: "id" },
  { name: "策略名称", field: "name" },
  { name: "用户", field: "user_name" },
  { name: "上行限速", field: "ul_flow_rate" },
  { name: "下行限速", field: "dl_flow_rate" },
  { name: "生效时间", field: "start_time" },
  { name: "失效时间", field: "end_time" },
  { name: "上行流速", field: "up_traffic_speed" },
  { name: "下行流速", field: "dn_traffic_speed" },
  { name: "上行通过流速", field: "up_pass_speed" },
  { name: "下行通过流速", field: "dn_pass_speed" },
  { name: "上行丢弃流速", field: "up_discard_speed" },
])
const timeRange = ref()
const configtable = ref([
  { name: "策略名称", field: "name" , type: "string"},
  { name: "链路", field: "link_ids" , type: "number",options:[]},
  { name: "用户", field: "user_id" , type: "number",options:[]},
  { name: "上行限速", field: "ul_flow_rate" },
  { name: "下行限速", field: "dl_flow_rate" },
  { name: "有效时间", field: "time" , type: "string",option:[]},
  { name: "生效时间", field: "start_time" ,option:[]},
  { name: "失效时间", field: "end_time" ,option:[]},
  { name: "流控类型", field: "flow_ctrl_type" ,options:[
      {name:"应用大小类",value:1},
      {name:"目的地址",value:2},
    ]},
  { name: "目的IP", field: "dst_ip"},
  { name: "目的端口", field: "dst_port"},
  { name: "目的地址", field: "dst_place"},
  { name: "应用大类", field: "app_type_id" , type: "number",options:[]},
  { name: "应用小类", field: "app_id", type: "number" ,options:[]},
  { name: "周期类型", field: "period_type" , type: "number",options:[
      {name:"每日",value:1},
      {name:"每周",value:2},
    ]},
  { name: "周期时间", field: "policy_period", type: "string" },
  { name: "备注", field: "remark" , type: "string"},
  { name: "开始周期", field: "startTime" , type: "string"},
  { name: "结束周期", field: "endTime" , type: "string"},
  { name: "开始星期", field: "start_week" , type: "string",options:[
      {name:"周一",value:"周一"},
      {name:"周二",value:"周二"},
      {name:"周三",value:"周三"},
      {name:"周四",value:"周四"},
      {name:"周五",value:"周五"},
      {name:"周六",value:"周六"},
      {name:"周日",value:"周日"},
    ]},
  { name: "结束星期", field: "end_week" , type: "string",options:[
      {name:"周一",value:"周一"},
      {name:"周二",value:"周二"},
      {name:"周三",value:"周三"},
      {name:"周四",value:"周四"},
      {name:"周五",value:"周五"},
      {name:"周六",value:"周六"},
      {name:"周日",value:"周日"},
    ]},
])

//验证开始时间是否大于结束时间
const validateStartTime = (rule, value, callback) => {
  const endTime = form.value.end_time
  if (new Date('1970-01-01 ' + value).getTime() >= new Date('1970-01-01 ' + endTime).getTime()) {
    callback(new Error('开始时间大于结束时间'))
  } else {
    callback();
  }
}
//验证结束时间是否大于开始时间
const validateEndTime = (rule, value, callback) => {
  const startTime = form.value.start_time
  if (new Date('1970-01-01 ' + value).getTime() <= new Date('1970-01-01 ' + startTime).getTime()) {
    callback(new Error('开始时间大于结束时间'))
  } else {
    callback();
  }
}

const validateIp=(rule, value, callback) =>{
  const reg = /^(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$/
   if(String(value).includes('-')){
      const ips = String(value).split('-');
      if(ips.length===2 && reg.test(ips[0]) && reg.test(ips[1])){
        callback();
    }
     callback(new Error('请输入正确的ip格式'))
  }else if(reg.test(value)){
     callback();
   }
  callback(new Error('请输入正确的ip格式'))
}

const validateFlow = (rule, value, callback) => {
  const ulRate = form.value.ul_flow_rate
  const dlRate = form.value.dl_flow_rate
  if (ulRate === null && dlRate === null) {
    callback(new Error('下行限速与上行限速必须有一个不为空'))
  } else {
    callback();
  }
}



const rules = ref({
  name: [{required: true, message: '名称不能为空', trigger: 'blur'}],
  user_id:[{required: true, message: '选择用户不能为空', trigger: 'blur'}],
  time:[{required: true, message: '有效时间不能为空', trigger: 'blur'}],
  ul_flow_rate:[{ message: '下行限速与上行限速必须有一个不为空', validator:validateFlow,trigger: 'blur'}],
  dl_flow_rate:[{ message: '下行限速与上行限速必须有一个不为空',validator:validateFlow, trigger: 'blur'}],
  start_time: [
    {required: true, message: '开始时间不能为空', trigger: 'change'},
    {message: "开始时间大于结束时间", validator: validateStartTime, trigger: 'change'}
  ],
  end_time: [
    {required: true, message: '结束时间不能为空', trigger: 'change'},
    {message: "开始时间大于结束时间", validator: validateEndTime, trigger: 'change'}
  ],
  link_ids: [{required: true, message: '链路不能为空', trigger: 'blur'}],
  dst_ip: [
    { required: true, message: '请输入设备IP', trigger: 'blur' },
    { message: 'IP地址不合法', validator: validateIp, trigger: 'blur' }
  ],
  dst_port: [{ message: '端口不合法', validator: validatePort, trigger: 'blur' }],
})

//页面
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)

let IpData = ref([])
let TimeData = ref([])
const WeekData = ref([])

const handleAddToWeekTable = () => {
  WeekData.value.push({
    start_week: form.value.start_week,
    end_week: form.value.end_week
  })
}

const handleAddToTimeTable = () => {
  TimeData.value.push({
      start_time: form.value.start_time,
      end_time: form.value.end_time
    })
}
const handleAddToIPTable = () => {
  if(form.value.dst_ip.includes('-')){
    IpData.value.push({
      start_ip:form.value.dst_ip.split('-')[0],
      end_ip: form.value.dst_ip.split('-')[1],
      dst_port: form.value.dst_port
    })
  }else {
    IpData.value.push({
      start_ip: form.value.dst_ip,
      end_ip: form.value.dst_ip,
      dst_port: form.value.dst_port
    })
  }
}

const handleDelFromTable = (row, index) => {
  IpData.value.splice(index, 1)
  TimeData.value.splice(index, 1)
  WeekData.value.splice(index, 1)
}


// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
}

const handleCurrentChange = (val) => {
  page.value = val
}
//工具栏
const initForm = () => {
  const temp_form = {}
  configtable.value.forEach(item => {
    temp_form[item.field] = null
  })
  form.value = {...form.value, ...temp_form}
  secondSelectVisible.value = false;
  firstSelectVisible.value = false;
  firstSelectIPVisible.value =false;
  TimeVisible.value = false;
  WeekVisible.value = false;
  timeRange.value = null;
  formRef.value.clearValidate()
  // TimeData.value = null;
  // const  WeekData = null;
  // const IpData = null;
}

const addFunc = () => {
  isEdit.value = false
  configVisible.value = true
  configTitle.value = "新建"
  initForm()
}

const deleteAllFunc = async () => {
  const checkedRows = tableParam.value.checkedRows
  let ids = []
  checkedRows.map(item => ids.push(item.id))
  const res = await control_deletePolicyApi({ids: ids})
  if (res.code === 0) {
    ElMessage.success(`删除成功！`)
    loadKey.value += 1
  }
}

function collectOptions(items,depth,tmpOption = []){
  if(depth === 0) {
    items.forEach(item => {
      tmpOption.push({
        value: item.id,
        name: item.label
      });
    });
    return tmpOption;
  }
  items.forEach(item=>{
    if(item.children && Array.isArray(item.children) && depth >= 1){
      collectOptions(item.children,depth - 1,tmpOption);
    }
  });
  return tmpOption;
}

//递归
const handleKey = (str, val, field) => {
    searchEl.value.forEach(item => {
      if(item.field==="userId") {
        const tmpOption = collectOptions(userOptions.value, 3)
        tmpOption.forEach(item => {
          item.value = item.value.split('-').pop()
        })
        item.options = tmpOption
      }
    })


  if (field === "appTypeId") {
    searchEl.value.forEach(item => {
      if (item.name === "应用小类") {
        item.options = val ? appTypeIdTreeOptions.value.filter(item => item.id === val)[0].children.map(item => {
          return {
            value: item.id,
            name: item.name
          }
        }) : []
      }
    })
  }
}


onBeforeMount( async () => {
  // 获取链路
  linkOptions.value = await LinkStore.getLink()
  linkOptions.value = linkOptions.value.map(item => {
    item.name = item.lineName
    item.value = item.lineVlan
    return item
  })

  // 获取应用大类
  if (localStorage.getItem('appType')) {
    appTypeIdOptions.value = JSON.parse(localStorage.getItem('appType'))
  } else {
    const res3 = await dictInfoApi('appType')
    if (res3.code === 0) {
      appTypeIdOptions.value = res3.data
      localStorage.setItem('appType', JSON.stringify(res3.data))
    }
  }

  // 获取应用小类
  if (localStorage.getItem('appTypeIdTreeOptions')) {
    appTypeIdTreeOptions.value = JSON.parse(localStorage.getItem('appTypeIdTreeOptions'))
  } else {
    const res4 = await dictInfoApi('appTypeIdTree')
    if (res4.code === 0) {
      appTypeIdTreeOptions.value = res4.data
      localStorage.setItem('appTypeIdTreeOptions', JSON.stringify(res4.data))
    }
  }

  // 获取用户群组树
  const res = await userCrowdGroupTreeApi(3)
  if (res.code === 0) {
    userOptions.value[0] = res.data
    localStorage.setItem('userOptions', JSON.stringify(res.data))
  }

  appTypeIdOptions.value = appTypeIdOptions.value.map(item => {
    return {
      value: item.id,
      name: item.name
    }
  })

  searchEl.value.forEach(item => {
    if (item.name === "应用大类") {
      item.options = (appTypeIdOptions.value)
    }
    if (item.name === "用户") {
      const tmpOption = collectOptions(userOptions.value, 3)
      tmpOption.forEach(item => {
        item.value = item.value.split('-').pop()
      })
      item.options = tmpOption
    }
  })

  configtable.value.forEach(item => {
    if (item.name === "应用大类") {
      item.options = (appTypeIdOptions.value)
    }
    if (item.name === "链路") {
      item.options = (linkOptions.value)
    }
    if (item.name === "用户") {
      item.options = collectOptions(userOptions.value, 3)
    }
  })
})

const searchFunc = () => { // 搜索
  loadKey.value += 1
}

const editFunc = (str, row) => {
  isEdit.value = true
  if(formRef.value ){formRef.value.clearValidate()}
  // console.log(row)
  IpData.value = []
  WeekData.value = []
  TimeData.value = []
  handlechange(row.period_type, 'period_type')
  handlechange(row.flow_ctrl_type, 'flow_ctrl_type')
  handlechange(row.user_type, 'user_type')
  handlechange(row.app_type_id, 'app_type_id')

  configVisible.value = true
  configTitle.value = "编辑"
  if(row.dst_ip !== "") {
    let temp = {}
    let port = {}
    temp.value = String(row.dst_ip).split(',')
    port.value = String(row.dst_port).split(',')

    temp.value.forEach((date, index) => {
      let ip = {}
      if (String(date).includes('-')) {
        ip.value = String(temp.value).split('-')
        IpData.value.push({
          start_ip: ip.value[0],
          end_ip: ip.value[1],
          dst_port: port.value[index]
        })
      } else {
        IpData.value.push({
          start_ip: date,
          end_ip: date,
          dst_port: port.value[index]
        })
      }
    })
  }

  if(row.period_type === 1 && row.policy_period !== ""){
    let temp1 = {}
    temp1.value = row.policy_period
    temp1.value = String(temp1.value).split(',')
    temp1.value.forEach(date=>{
      let time1 = {}
      time1.value = String(date).split('-')
      TimeData.value.push({
        start_time: time1.value[0],
        end_time: time1.value[1]
      })
    })
  }

  if(row.period_type === 2 && row.policy_period !==""){
    let temp = {}
    temp.value = row.policy_period
    temp.value = String(temp.value).split(',')
    temp.value.forEach(date=>{
      let time = {}
      time.value = String(date).split('-')
      WeekData.value.push({
        start_week: time.value[0],
        end_week: time.value[1]
      })
    })
  }

  configtable.value.forEach(item => {
    form.value.id = row.id
    form.value[item.field] = row[item.field]
    if(item.field === "app_type_id" && row[item.field] === 0){
      form.value[item.field] = null
    }
    if(item.field === "app_id" && row[item.field] === 0){
      form.value[item.field] = null
    }
    if(item.field==="user_id"){
      form.value[item.field]= row.user_name
    }
  })
  form.value['time']=[formatTimeToStr(row.start_time, 'yyyy-MM-dd 00:00:00'), formatTimeToStr(row.end_time, 'yyyy-MM-dd hh:mm:00')]

  const link_ids = []
  row.link_ids.split(',').forEach((item,index)=>{
    link_ids[index]=linkOptions.value.filter(object=>object.lineVlan===parseInt(item))[0].lineName
  })
  form.value.link_ids = link_ids
}

const clearselect = () => { // 搜索
  initForm()
}

const deleteFunc = async (str, row) => {
  const res = await control_deletePolicyApi({ids: [row.id]})
  if (res.code === 0) {
    ElMessage.success( `删除成功！`)
    loadKey.value += 1
  }
}


//二级页面控制下拉框
const handlechange = (value, field) => {
  // console.log(value, field)
  if(value === 1 && field === "period_type"){
    TimeVisible.value = true;
    WeekVisible.value = false;
  }else if(value === 2 && field === "period_type"){
    TimeVisible.value = false;
    WeekVisible.value = true;
  }


  if(value === 1 && field === "flow_ctrl_type"){
    firstSelectVisible.value = true;
    firstSelectIPVisible.value = false;
  }else if(value === 2 && field === "flow_ctrl_type"){
    firstSelectIPVisible.value = true;
    firstSelectVisible.value = false;
  }

  if (field === "app_type_id") {
    configtable.value.forEach(item => {
      secondSelectVisible.value = true;
      if (item.name === "应用小类") {
        item.options = value ? appTypeIdTreeOptions.value.filter(item => item.id === value)[0].children.map(item => {
          return {
            value: item.id,
            name: item.name
          }
        }) : []
      }
    })
  }


}
//提交创建与修改
const handleSubmit = async () => {
  formRef.value.validate(async valid => {
    if (valid) {
      let link_ids = []
      let res = {}
      const regex = /^\d+$/
      let temp={}
      // console.log(form.value)
      if(form.value.ul_flow_rate!==null) {
        form.value.ul_flow_rate = parseInt(form.value.ul_flow_rate)
      }
      if(form.value.dl_flow_rate!==null) {
        form.value.dl_flow_rate = parseInt(form.value.dl_flow_rate)
      }
      temp.value = form.value['user_id']
      temp.value = temp.value.split('-').pop()
      form.value['user_id'] =parseInt(temp.value)
      if (isEdit.value) {
        form.value.link_ids.map(item=>{
          if (!regex.test(item)){
            link_ids.push(linkOptions.value.filter(object=>object.lineName===item)[0].lineVlan)
          }else {
            link_ids.push(item)
          }
        })
        form.value.link_ids=link_ids
        form.value.link_ids = form.value.link_ids.toString()
      } else {
        form.value.link_ids = form.value.link_ids.toString()
      }
      form.value['start_time'] = formatDate(form.value['time'][0])
      form.value['end_time'] = formatDate(form.value['time'][1])
      if(form.value.flow_ctrl_type === null){
        form.value.flow_ctrl_type = 1
      }
      if(form.value.period_type === null){
        form.value.period_type = 1
      }
      // 给端口，ip，周期赋值
      if(form.value.flow_ctrl_type === 2){
        IpData.value.forEach((date,index)=>{
          let temp = {}
          let port = {}
          temp = date.start_ip+ "-" + date.end_ip
          port = date.dst_port
          if(index === 0){
            form.value.dst_ip =  temp
            form.value.dst_port =  port

          }else {
            form.value.dst_ip = form.value.dst_ip +","+ temp
            //form.value.dst_port =  form.value.dst_port +","+ port
          }
        })
      }
      if(form.value.period_type === 1){
        TimeData.value.forEach((date,index)=>{
          let temp = {}
          temp = date.start_time + "-" + date.end_time
          if(index === 0){
            form.value.policy_period =  temp
          }else {
            form.value.policy_period = form.value.policy_period +","+ temp
          }
        })
      }else{
       WeekData.value.forEach((date,index)=>{
          let temp = {}
          temp = date.start_week + "-" + date.end_week
          if(index === 0){
            form.value.policy_period =  temp
          }else {
            form.value.policy_period = form.value.policy_period +","+ temp
          }
        })
      }

      res = await control_createPolicyApi(form.value)
      if (res.code === 0) {
        ElMessage.success( `操作成功！`)

        configVisible.value = false
        loadKey.value += 1
      }
    }
  })
}

</script>

<style scoped lang="scss">

</style>
