<template>
  <div class="flex items-center justify-center h-screen">
    <div class="flex flex-col items-center">
      <img class="w-20 h-20" src="/public/logo.png" alt="logo" />
      <h1 class="text-3xl mb-4">欢迎登录</h1>
      <el-form
          ref="ruleFormRef"
          :model="form"
          :rules="rules"
      >
        <el-form-item prop="userID">
          <el-input v-model="form.userID" placeholder="请输入用户名"></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input type="password" v-model="form.password" placeholder="请输入密码"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button class="w-full" type="primary" @click="loginFunc">登录</el-button>
          <div class="flex w-full justify-end">
            <span class="cursor-pointer text-blue-500" @click="toRegister">前去注册</span>
          </div>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import { ref,onMounted } from 'vue'
import { login,getUserInfo,logout } from '@/api/user'
import {useRouter} from "vue-router";
import cookie from 'js-cookie'


const router = useRouter()

const form = ref({
  userID: '',
  password: ''
})

const rules = {
  userID: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 5, max: 20, message: '长度在 5 到 20 个字符', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, max: 12, message: '长度在 6 到 12 个字符', trigger: 'blur' }
  ]
}

const ruleFormRef = ref(null)

const getUserInfoFunc = async () => {
  const res = await getUserInfo()
  if(res.code === 0){
    router.push("/")
  }else{
    console.log("<UNK>")
  }
}

const loginFunc = async () =>{
  ruleFormRef.value.validate(async valid => {
    if(!valid) return
    const res = await login(form.value)
    if(res.code === 0){
      getUserInfoFunc()
      console.log("OK")
    }
  })
}

const toRegister = () => {
  router.push('register')
}

onMounted(()=>{
  if (cookie.get('x-token')) {
    getUserInfoFunc()
  }
})

</script>

<style scoped lang="scss">

</style>
