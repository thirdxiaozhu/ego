<template>
  <div class="flex items-center justify-center h-screen">
    <div class="flex flex-col items-center">
      <img class="w-20 h-20" src="/public/logo.png" alt="logo" />
      <h1 class="text-3xl mb-4">欢迎注册</h1>
      <el-form
          ref="ruleFormRef"
          :model="form"
          :rules="rules"
      >
        <el-form-item prop="username">
          <el-input v-model="form.username" placeholder="请输入用户名"></el-input>
        </el-form-item>
        <el-form-item prop="nickname">
          <el-input v-model="form.nickname" placeholder="请输入用户昵称"></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input type="password" v-model="form.password" placeholder="请输入密码"></el-input>
        </el-form-item>
        <el-form-item prop="rePassword">
          <el-input type="password" v-model="form.rePassword" placeholder="请再次输入密码"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button class="w-full" type="primary" @click="registerFunc">注册</el-button>
          <div class="flex w-full justify-end">
            <span class="cursor-pointer text-blue-500" @click="toLogin">已有帐号</span>
          </div>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { register } from '@/api/user'
import {useRouter} from "vue-router";
import {ElMessage} from "element-plus";

const router = useRouter()

const form = ref({
  username: '',
  password: '',
  rePassword: ''
})

const validatePass = (rule, value, callback) => {
  if (value !== form.value.password) {
    callback(new Error('两次输入的密码不一致'))
    return
  }
    callback()
}

const rules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 5, max: 10, message: '长度在 5 到 10 个字符', trigger: 'blur' }
  ],
  nickname: [
    { required: true, message: '请输入用户昵称', trigger: 'blur' },
    { min: 3, max: 10, message: '长度在 3 到 10 个字符', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, max: 12, message: '长度在 6 到 12 个字符', trigger: 'blur' }
  ],
  rePassword: [
    {validator: validatePass, trigger: 'blur'}
  ]
}

const ruleFormRef = ref(null)

const registerFunc = async () =>{
  ruleFormRef.value.validate(async valid => {
    if(!valid) return
    const res = await register(form.value)
    if(res.code === 0){
      ElMessage.success('注册成功')
      toLogin()
    }
  })
}

const toLogin = () => {
  router.push('login')
}
</script>

<style scoped lang="scss">

</style>
