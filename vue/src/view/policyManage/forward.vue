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
        module-api="White_policyListApi"
        @handleEdit="editFunc"
        @handleDelete="deleteFunc"
      />
      <!--  配置二级   -->
      <el-dialog
        v-model="configVisible"
        :title="configTitle"
        class="w-182 h-auto"
      >
        <!--          二级页面-->
        <el-form
          ref="formRef"
          class="config-form"
          :model="form"
          :rules="rules"
        >
          <!--            策略名称-->
          <el-form-item
            :label="configTable[0].name"
            label-width="120px"
            :prop="configTable[0].field"
          >
            <el-input
              v-model="form[configTable[0].field]"
              placeholder="请输入策略名称"
              autocomplete="off"
              class="w-112"
            />
          </el-form-item>
          <!--            用户类型-->
          <el-form-item
            :label="configTable[1].name"
            label-width="120px"
            :prop="configTable[1].field"
          >
            <el-select
              v-model="form[configTable[1].field]"
              :placeholder="`请选择${configTable[1].name}`"
              class="w-112"
              @change="handleChange($event, configTable[1].field)"
            >
              <el-option
                v-for="item in configTable[1].options"
                :label="item.name"
                :value="item.value"
              />
            </el-select>
          </el-form-item>

          <!--            用户群组-->
          <el-form-item
            v-if="secondUserCrowdGroupVisible"
            :label="configTable[2].name"
            label-width="120px"
            :prop="configTable[2].field"
          >
            <el-select
              v-model="form[configTable[2].field]"
              :placeholder="`请选择${configTable[2].name}`"
              class="w-112"
              @change="handleChange($event, configTable[2].field)"
            >
              <el-option
                v-for="item in configTable[2].options"
                :label="item.name"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          <!--            用户群-->
          <el-form-item
            v-if="secondUserCrowdVisible"
            :label="configTable[3].name"
            label-width="120px"
            :prop="configTable[3].field"
          >
            <el-select
              v-model="form[configTable[3].field]"
              :placeholder="`请选择${configTable[3].name}`"
              filterable
              class="w-112"
              @change="handleChange($event, configTable[3].field)"
            >
              <el-option
                v-for="item in configTable[3].options"
                :label="item.name"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          <!--            用户-->
          <el-form-item
            v-if="secondUserVisible"
            :label="configTable[4].name"
            label-width="120px"
            :prop="configTable[4].field"
          >
            <el-select
              v-model="form[configTable[4].field]"
              :placeholder="`请选择${configTable[4].name}`"
              filterable
              class="w-112"
              @change="handleChange($event, configTable[4].field)"
            >
              <el-option
                v-for="item in configTable[4].options"
                :label="item.name"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          <!--            上行tos-->
          <el-row :gutter="10">
            <el-col :span="12">
          <el-form-item
            :label="configTable[5].name"
            label-width="120px"
            :prop="configTable[5].field"
          >
            <el-input
              v-model="form[configTable[5].field]"
              :placeholder="`请输入${configTable[5].name}`"
              autocomplete="off"
              class="w-112"
            />
          </el-form-item>
            </el-col>
          <!--            下行tos-->
            <el-col :span="12">
          <el-form-item
            :label="configTable[6].name"
            label-width="120px"
            :prop="configTable[6].field"
          >
            <el-input
              v-model="form[configTable[6].field]"
              :placeholder="`请输入${configTable[6].name}`"
              autocomplete="off"
              class="w-112"
            />
          </el-form-item>
            </el-col>
          </el-row>
          <!--            应用大类-->
          <el-row :gutter="40">
            <el-col :span="12">
          <el-form-item
            :label="configTable[7].name"
            label-width="120px"
            :prop="configTable[7].field"
          >
            <el-select
              v-model="form[configTable[7].field]"
              filterable
              :placeholder="`请选择${configTable[7].name}`"
              class="w-112"
              @change="handleChange($event, configTable[7].field)"
            >
              <el-option
                v-for="item in configTable[7].options"
                :label="item.name"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
            </el-col>
          <!--            应用小类-->
            <el-col :span="12">
          <el-form-item
            v-if="secondSelectVisible"
            label-width="120px"
            :label="configTable[8].name"
            :prop="configTable[8].field"
          >
            <el-select
              v-model="form[configTable[8].field]"
              filterable
              :placeholder="`请选择${configTable[8].name}`"
              class="w-112"
              @change="handleChange($event, configTable[8].field)"
            >
              <el-option
                v-for="(item,index) in configTable[8].options"
                :key="index"
                :label="item.name"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
            </el-col>
          </el-row>
          <!--            有效时间-->
          <el-form-item
            :label="configTable[9].name"
            label-width="120px"
            :prop="configTable[9].field"
          >
            <el-date-picker
              v-model="form['timeRange']"
              type="datetimerange"
              range-separator="至"
              start-placeholder="生效时间"
              end-placeholder="失效时间"
              format="YYYY-MM-DD HH:mm:ss"
              data-format="YYYY/MM/DD ddd"
              time-format="HH:mm:ss"
              size="small"
            />
          </el-form-item>
          <!--            备注-->
          <el-form-item
            :label="configTable[12].name"
            label-width="120px"
            :prop="configTable[12].field"
          >
            <el-input
              v-model="form[configTable[12].field]"
              :placeholder="`请输入${configTable[12].name}`"
              autocomplete="off"
              class="w-112"
              :type="configTable[12].type = 'textarea'"
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
    </div>
  </div>
</template>

<script setup>
import {onBeforeMount, ref} from "vue"
import ToolBar from "@/components/commonTable/toolbar.vue";
import CommonTable from "@/components/commonTable/index.vue";
import {ElMessage} from "element-plus";
import {formatDate} from '@/utils/format'
import {dictInfoApi} from "@/api/traffic";
import useLinkStore from "@/pinia/modules/link";
import {White_createPolicyApi, White_deletePolicyApi } from "@/api/policy"
import {useUserTreeStore} from "@/pinia/modules/useUserTree";
import {formatTimeToStr} from "@/utils/date";

const LinkStore = useLinkStore()
const userTreeStore = useUserTreeStore()


// 定义工具栏配置
const timeRange = ref(null)
const formRef = ref(null)
const form = ref({
  id: null
})
const configVisible = ref(false)
const configTitle = ref("新增")
const toolbarParam = ref(null)
const deleteFlag = ref(true)
const linkOptions = ref([])
const appTypeIdOptions = ref([])//应用大类
const appTypeIdTreeOptions = ref([])//应用小类
const userOptions = ref([])//用户群组树
const searchEl = ref([
  {name: "请输入策略名称", field: "policyName"},
  {
    name: "请选择用户群组/用户", field: "user_type", options: [
      {name: "用户群组", value: 1},
      {name: "用户群", value: 2},
      {name: "用户", value: 3},
    ]
  },
  {name: "用户群组", field: "uerCrowdGroupId", options: []},
  {name: "应用大类", field: "appTypeId", options: []},
  {name: "应用小类", field: "appId", options: []},
])

//应用小类下拉框显示值
const secondSelectVisible = ref(false)
const secondUserCrowdGroupVisible = ref(false)
const secondUserCrowdVisible = ref(false)
const secondUserVisible = ref(false)


// 列表查询逻辑
const loadKey = ref(0)
const tableParam = ref(null)
const tableColumns = ref([
  {name: "策略ID", field: "id", sortable: true},
  {name: "策略名称", field: "name", sortable: true},
  {name: "用户群组", field: "user_crowd_group_name", sortable: true},
  {name: "用户群", field: "user_crowd_name", sortable: true},
  {name: "用户", field: "user_name", sortable: true},
  {name: "上行tos", field: "ul_tos", sortable: true},
  {name: "下行tos", field: "dl_tos", sortable: true},
  {name: "应用大类", field: "app_type_name", sortable: true},
  {name: "应用小类", field: "app_name", sortable: true},
  {name: "生效时间", field: "start_time", sortable: true},
  {name: "失效时间", field: "end_time", sortable: true},
])


const configTable = ref([
  {name: "策略名称", field: "name", type: "string"},
  {
    name: "用户类型", field: "user_type", type: "number", options: [
      {name: "用户群组", value: 1},
      {name: "用户群", value: 2},
      {name: "用户", value: 3},
    ]
  },
  {name: "用户群组", field: "user_crowd_group_id", type: "number", options: []},
  {name: "用户群", field: "user_crowd_id", type: "number", options: []},
  {name: "用户", field: "user_id", type: "number", options: []},
  {name: "上行tos", field: "ul_tos", type: "number"},
  {name: "下行tos", field: "dl_tos", type: "number"},
  {name: "应用大类", field: "app_type_id", type: "number", options: []},
  {name: "应用小类", field: "app_id", type: "number", options: []},
  {name: "有效时间", field: "timeRange", option: []},
  {name: "生效时间", field: "start_time", type: "string", option: []},
  {name: "失效时间", field: "end_time", type: "string", option: []},
  {name: "备注", field: "remark", type: "string"},
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

const validateUser = (rule, value, callback) => {
  const appTypeId = form.value.app_type_id
  if (appTypeId === null && value === null) {
    callback(new Error('应用大小类与用户至少要选择一个'))
  } else {
    callback();
  }
}

const validateApp = (rule, value, callback) => {
  const userType = form.value.user_type
  if (userType === null && value === null) {
    callback(new Error('应用大小类与用户至少要选择一个'))
  } else {
    callback();
  }
}
const validateAppId = (rule, value, callback) => {
  const appTypeId = form.value.app_type_id
  if (appTypeId !== null && value === null) {
    callback(new Error('请选择应用小类'))
  } else {
    callback();
  }
}

const validateTos = (rule, value, callback) => {
  const ulTos = form.value.ul_tos
  const dlTos = form.value.dl_tos
  if (ulTos === null && dlTos === null) {
    callback(new Error('下行tos与上行tos必须有一个不为空'))
  }else {
    callback();
  }
}

const validateInt = (rule, value, callback) => {
  const regex = /^$|^(?:[1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/
  if(value !== null) {
    if (!regex.test(value)) {
      callback(new Error('请输入正确的数值,数值范围应在0-255!'))
    } else {
      callback();
    }
  }else{
    callback();
  }
}


const rules = ref({
  name: [{required: true, message: '名称不能为空', trigger: 'blur'}],
  user_type:[{message:"应用大小类与用户至少要选择一个",validator:validateUser,trigger: 'blur'}],
  app_type_id:[{message:"应用大小类与用户至少要选择一个",validator:validateApp,trigger: 'blur'}],
  ul_tos:[{ message: '下行tos与上行tos必须有一个不为空', validator:validateTos,trigger: 'blur'},
    { message: '请输入正确的数值,数值范围应在0-255!', validator:validateInt,trigger: 'blur'}],
  dl_tos:[{ message: '下行tos与上行tos必须有一个不为空',validator:validateTos, trigger: 'blur'},
    { message: '请输入正确的数值,数值范围应在0-255!', validator:validateInt,trigger: 'blur'}],
  timeRange:[{ required: true,message: '有效时间不能为空', trigger: 'blur'}],
  start_time: [
    {required: true, message: '开始时间不能为空', trigger: 'change'},
    {message: "开始时间大于结束时间", validator: validateStartTime, trigger: 'change'}
  ],
  end_time: [
    {required: true, message: '结束时间不能为空', trigger: 'change'},
    {message: "开始时间大于结束时间", validator: validateEndTime, trigger: 'change'}
  ],
  app_id: [{ message: '应用小类ID不能为空', validator: validateAppId,trigger: 'blur'}],
})



//工具栏
const initForm = () => {
  const temp_form = {}
  configTable.value.forEach(item => {
    temp_form[item.field] = null
  })
  form.value = {...form.value, ...temp_form}
  secondUserCrowdGroupVisible.value = false;
  secondUserCrowdVisible.value = false;
  secondUserVisible.value = false;
  timeRange.value = null;
  formRef.value.clearValidate()
}

const addFunc = () => {
  configVisible.value = true
  configTitle.value = "新建"
  initForm()
}

const deleteAllFunc = async () => {
  const checkedRows = tableParam.value.checkedRows
  let ids = []
  checkedRows.map(item => ids.push(item.id))
  const res = await White_deletePolicyApi({ids: ids})
  if (res.code === 0) {
    ElMessage.success(`删除成功！`)
    loadKey.value += 1
  }
}

function collectOptions(items, depth, tmpOption = []) {
  if (depth === 0) {
    items.forEach(item => {
      tmpOption.push({
        value: item.id,
        name: item.label
      });
    });
    return tmpOption;
  }
  items.forEach(item => {
    if (item.children && Array.isArray(item.children) && depth >= 1) {
      collectOptions(item.children, depth - 1, tmpOption);
    }
  });
  return tmpOption;
}

//递归
const handleKey = (str, val, field) => {
  // console.log(val, field)
  if (field === "user_type") {
    searchEl.value.forEach(item => {
      const tmpOption = collectOptions(userOptions.value, Number(val))
      if (item.name === "用户群组") {
        if (val === 1) {
          item.field = "uerCrowdGroupId"
        } else if (val === 2) {
          item.field = "uerCrowdId"
          tmpOption.forEach(item => {
            item.value = item.value.split('-').pop()
          })
        } else if (val === 3) {
          item.field = "uerId"
          tmpOption.forEach(item => {
            item.value = item.value.split('-').pop()
          })
        }
        item.options = tmpOption
      }
    })
  }
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


onBeforeMount(async () => {
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
  userOptions.value[0] = await userTreeStore.getUserTree()


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
  })

  configTable.value.forEach(item => {
    if (item.name === "应用大类") {
      item.options = appTypeIdOptions.value
    } else if (item.field === "user_crowd_group_id") {
      item.options = collectOptions(userOptions.value, 1)
    } else if (item.field === "user_crowd_id") {
      item.options = collectOptions(userOptions.value, 2)
    } else if (item.field === "user_id") {
      item.options = collectOptions(userOptions.value, 3)
    }
  })
})

const searchFunc = () => { // 搜索
  loadKey.value += 1
}


const editFunc = (str, row) => {
  if(formRef.value ){formRef.value.clearValidate()}
  // console.log(row)
  handleChange(row.user_type, 'user_type')
  handleChange(row.app_type_id, 'app_type_id')
  configVisible.value = true
  configTitle.value = "编辑"
  configTable.value.forEach(item => {
    form.value.id = row.id
    form.value[item.field] = row[item.field]
    if(item.field === "user_type" && row[item.field] === 0){
      form.value[item.field] = null
    }
    if(item.field === "user_type" && row[item.field] === 0){
      form.value[item.field] = null
    }
    if(item.field === "user_type" && row[item.field] === 0){
      form.value[item.field] = null
    }
    if(item.field === "app_type_id" && row[item.field] === 0){
      form.value[item.field] = null
    }
    if(item.field === "app_id" && row[item.field] === 0){
      form.value[item.field] = null
    }
    if (item.field === "user_id") {
      form.value[item.field] = row.user_name
    }
    if (item.field === "user_crowd_id") {
      form.value[item.field] = row.user_crowd_name
    }
    // form.value.user_id += ""
    // form.value.user_crowd_id += ""
    form.value.user_crowd_group_id += ""
  })
  form.value['timeRange']=[formatTimeToStr(row.start_time, 'yyyy-MM-dd 00:00:00'), formatTimeToStr(row.end_time, 'yyyy-MM-dd hh:mm:00')]

}

const clearselect = () => {
  initForm()
}

const deleteFunc = async (str, row) => {
  // console.log(row)
  const res = await White_deletePolicyApi({ids: [row.id]})
  if (res.code === 0) {
    ElMessage.success(`删除成功！`)
    loadKey.value += 1
  }
}

//二级页面控制下拉框
const handleChange = (value, field) => {
  if (field === "user_type") {
    if (value === 1) {
      secondUserCrowdGroupVisible.value = true;
      secondUserCrowdVisible.value = false;
      secondUserVisible.value = false;
    }
    if (value === 2) {
      secondUserCrowdVisible.value = true;
      secondUserCrowdGroupVisible.value = false;
      secondUserVisible.value = false;
    }
    if (value === 3) {
      secondUserVisible.value = true;
      secondUserCrowdVisible.value = false;
      secondUserCrowdGroupVisible.value = false;
    }
  }
  if (field === "app_type_id") {
    configTable.value.forEach(item => {
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
    // console.log(form.value)
    if (valid) {
      let res = {}
      let temp = {}
      let temp2 = {}
      if(form.value.ul_tos!==null) {
        form.value.ul_tos = parseInt(form.value.ul_tos)
      }
      if(form.value.dl_tos!==null) {
        form.value.dl_tos = parseInt(form.value.dl_tos)
      }
      form.value.user_crowd_group_id = parseInt(form.value.user_crowd_group_id)
      if(form.value['user_id']!==null) {
        temp.value = form.value['user_id']
        temp.value = temp.value.split('-').pop()
        form.value['user_id'] = parseInt(temp.value)
      }
      if(form.value['user_crowd_id']!==null) {
        temp2.value = form.value['user_crowd_id']
        temp2.value = temp2.value.split('-').pop()
        form.value['user_crowd_id'] = parseInt(temp.value)
      }
      form.value['start_time'] = formatDate(form.value['timeRange'][0])
      form.value['end_time'] = formatDate(form.value['timeRange'][1])
      res = await White_createPolicyApi(form.value)
      if (res.code === 0) {
        ElMessage.success(`操作成功！`)
        configVisible.value = false
        loadKey.value += 1
      }
    } else {
      console.log('error submit')
    }
  })
}
</script>

<style scoped lang="scss">
.config-form {
  width: 500px;
}
</style>
