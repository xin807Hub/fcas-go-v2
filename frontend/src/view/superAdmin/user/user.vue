<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="addUser">
          新增用户
        </el-button>
      </div>
      <el-table :data="tableData" row-key="ID">
        <el-table-column align="left" label="ID" min-width="50" prop="ID" />
        <el-table-column align="left" label="用户名" min-width="150" prop="userName" />
        <el-table-column align="left" label="昵称" min-width="150" prop="nickName" />
        <el-table-column align="left" label="手机号" min-width="180" prop="phone" />
        <el-table-column align="left" label="邮箱" min-width="180" prop="email" />
        <!--        <el-table-column align="left" label="Bypass切换密码" min-width="180" prop="bypassPassword">-->
        <!--          <template #default="{row}">-->
        <!--            <el-input type="password" :value="row.bypassPassword" disabled />-->
        <!--          </template>-->
        <!--        </el-table-column>-->
        <el-table-column align="left" label="用户角色" min-width="">
          <template #default="{row}">
            {{ authOptions.find(item => item.authorityId === row.authority.authorityId)?.authorityName }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="启用" min-width="150">
          <template #default="scope">
            <el-switch
              v-model="scope.row.enable" inline-prompt :active-value="1" :inactive-value="2"
              @change="() => { switchEnable(scope.row) }"
            />
          </template>
        </el-table-column>

        <el-table-column label="操作" min-width="250" align="center" fixed="right">
          <template #default="scope">
            <el-button type="primary" link icon="edit" @click="openEdit(scope.row)">
              编辑
            </el-button>
            <!-- <el-button type="primary" link icon="magic-stick" @click="resetPasswordFunc(scope.row)">重置密码</el-button> -->
            <el-button type="primary" link icon="lock" @click="openChangePassword(scope.row)">
              修改密码
            </el-button>
            <el-button type="primary" link icon="delete" @click="deleteUserFunc(scope.row)">
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total"
          layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <el-drawer
      v-model="addUserDialog" size="60%" :show-close="false" :close-on-press-escape="false"
      :close-on-click-modal="false"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">用户</span>
          <div>
            <el-button @click="closeAddUserDialog">
              取 消
            </el-button>
            <el-button type="primary" @click="enterAddUserDialog">
              确 定
            </el-button>
          </div>
        </div>
      </template>

      <el-form ref="userForm" :rules="rules" :model="userInfo" label-width="140px">
        <el-form-item v-if="dialogFlag === 'add'" label="用户名" prop="userName">
          <el-input v-model="userInfo.userName" />
        </el-form-item>
        <el-form-item v-if="dialogFlag === 'add'" label="密码" prop="password">
          <el-input v-model="userInfo.password" />
        </el-form-item>
        <el-form-item label="昵称" prop="nickName">
          <el-input v-model="userInfo.nickName" />
        </el-form-item>
        <el-form-item label="手机号" prop="phone">
          <el-input v-model="userInfo.phone" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="userInfo.email" />
        </el-form-item>
        <el-form-item
          v-if="userInfo.authorityId === 888"
          label="Bypass切换密码"
          prop="bypassPassword"
          :rules="[
            {required: true, message: '请输入bypassPassword密码', trigger: 'blur'}
          ]"
        >
          <el-input v-model="userInfo.bypassPassword" />
        </el-form-item>
        <el-form-item
          v-if="btnAuth.authChange"
          label="用户角色" prop="authorityId"
        >
          <el-cascader
            v-model="userInfo.authorityId" style="width:100%" :options="authOptions"
            :show-all-levels="false"
            :props="{ checkStrictly: true, label: 'authorityName', value: 'authorityId', disabled: 'disabled', emitPath: false }"
            :clearable="false"
          />
        </el-form-item>
        <el-form-item label="启用" prop="disabled">
          <el-switch v-model="userInfo.enable" inline-prompt :active-value="1" :inactive-value="2" />
        </el-form-item>
      </el-form>
    </el-drawer>


    <!-- 新增密码修改对话框组件 -->
    <el-dialog v-model="changePasswordDialog" title="修改密码" width="30%">
      <el-form ref="passwordFormEl" :model="passwordInfo" :rules="passwordRules" label-width="80px">
        <el-form-item label="原密码" prop="password">
          <el-input v-model="passwordInfo.password" type="password" />
        </el-form-item>
        <el-form-item label="新密码" prop="newPassword">
          <el-input v-model="passwordInfo.newPassword" type="password" />
        </el-form-item>
        <el-form-item label="确认密码" prop="confirmPassword">
          <el-input v-model="passwordInfo.confirmPassword" type="password" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="closeChangePasswordDialog">取 消</el-button>
          <el-button type="primary" @click="submitChangePassword">确 定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>

import {
  getUserList,
  register,
  deleteUser, changePassword
} from '@/api/user'

import {getAuthorityList} from '@/api/authority'
import {setUserInfo} from '@/api/user.js'

import {nextTick, ref, onMounted} from 'vue'
import {ElMessage, ElMessageBox} from 'element-plus'
import {useBtnAuth} from "@/utils/btnAuth";

const btnAuth = useBtnAuth()

defineOptions({
  name: 'User',
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async () => {
  const table = await getUserList({page: page.value, pageSize: pageSize.value})
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

const authOptions = ref([])
const fetchAuthOptions = async () => {
  const res = await getAuthorityList({page: 1, pageSize: 999})
  authOptions.value = res.data.list
}


const deleteUserFunc = async (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await deleteUser({id: row.ID})
    if (res.code === 0) {
      ElMessage.success('删除成功')
      await getTableData()
    }
  })
}

// 弹窗相关
const userInfo = ref({
  username: '',
  password: '',
  nickName: '',
  authorityId: '',
  enable: 1,
})

const rules = ref({
  userName: [
    {required: true, message: '请输入用户名', trigger: 'blur'},
    {min: 5, message: '最低5位字符', trigger: 'blur'}
  ],
  password: [
    {required: true, message: '请输入用户密码', trigger: 'blur'},
    {min: 6, message: '最低6位字符', trigger: 'blur'}
  ],
  nickName: [
    {required: true, message: '请输入用户昵称', trigger: 'blur'}
  ],
  phone: [
    {pattern: /^1([38][0-9]|4[014-9]|[59][0-35-9]|6[2567]|7[0-8])\d{8}$/, message: '请输入合法手机号', trigger: 'blur'},
  ],
  email: [
    {
      pattern: /^([0-9A-Za-z\-_.]+)@([0-9a-z]+\.[a-z]{2,3}(\.[a-z]{2})?)$/g,
      message: '请输入正确的邮箱',
      trigger: 'blur'
    },
  ],
  authorityId: [
    {required: true, message: '请选择用户角色', trigger: 'blur'}
  ]
})
const userForm = ref(null)
const enterAddUserDialog = async () => {
  userInfo.value.authorityIds = [userInfo.value.authorityId]

  userForm.value.validate(async valid => {
    if (valid) {
      const req = {
        ...userInfo.value
      }
      if (dialogFlag.value === 'add') {
        const res = await register(req)
        if (res.code === 0) {
          ElMessage({type: 'success', message: '创建成功'})
          await getTableData()
          closeAddUserDialog()
        }
      }
      if (dialogFlag.value === 'edit') {
        const res = await setUserInfo(req)
        if (res.code === 0) {
          ElMessage({type: 'success', message: '编辑成功'})
          await getTableData()
          closeAddUserDialog()
        }
      }
    }
  })
}

const addUserDialog = ref(false)
const closeAddUserDialog = () => {
  userForm.value.resetFields()
  addUserDialog.value = false
}

const dialogFlag = ref('add')

const addUser = () => {
  dialogFlag.value = 'add'
  addUserDialog.value = true
}


const openEdit = (row) => {
  dialogFlag.value = 'edit'
  userInfo.value = JSON.parse(JSON.stringify(row))
  addUserDialog.value = true
}

const switchEnable = async (row) => {
  userInfo.value = JSON.parse(JSON.stringify(row))
  await nextTick()
  const req = {
    ...userInfo.value
  }
  const res = await setUserInfo(req)
  if (res.code === 0) {
    ElMessage({type: 'success', message: `${req.enable === 2 ? '禁用' : '启用'}成功`})
    await getTableData()
  }
}


// 密码修改对话框相关
const changePasswordDialog = ref(false)
const passwordFormEl = ref(null)
const passwordInfo = ref({
  ID: '',
  password: '',
  newPassword: '',
  confirmPassword: '',
})

const validateConfirmPassword = (rule, value, callback) => {
  if (value !== passwordInfo.value.newPassword) {
    callback(new Error('两次输入的密码不一致'))
  } else {
    callback()
  }
}

const passwordRules = ref({
  password: [
    {required: true, message: '请输入原密码', trigger: 'blur'},
  ],
  newPassword: [
    {required: true, message: '请输入新密码', trigger: 'blur'},
    {min: 6, message: '最低6位字符', trigger: 'blur'}
  ],
  confirmPassword: [
    {required: true, message: '请确认密码', trigger: 'blur'},
    {validator: validateConfirmPassword, trigger: 'blur'}
  ],
})


const openChangePassword = (row) => {
  passwordInfo.value.ID = row.ID
  passwordInfo.value.password = ''
  passwordInfo.value.newPassword = ''
  passwordInfo.value.confirmPassword = ''
  changePasswordDialog.value = true
}

const closeChangePasswordDialog = () => {
  passwordFormEl.value.resetFields()
  changePasswordDialog.value = false
}

const submitChangePassword = async () => {
  passwordFormEl.value.validate(async valid => {
    if (valid) {
      const res = await changePassword({
        password: passwordInfo.value.password,
        newPassword: passwordInfo.value.newPassword,
      })
      if (res.code === 0) {
        ElMessage({type: 'success', message: '密码修改成功'})
        closeChangePasswordDialog()
      } else {
        ElMessage({type: 'error', message: res.msg})
      }
    }
  })
}

onMounted(async () => {
  await getTableData()
  await fetchAuthOptions()
})

</script>

<style lang="scss">
.header-img-box {
  @apply w-52 h-52 border border-solid border-gray-300 rounded-xl flex justify-center items-center cursor-pointer;
}
</style>
