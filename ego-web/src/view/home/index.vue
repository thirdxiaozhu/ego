
<template>
  <div>
    您已登录
    {{info}}
    <el-button @click="logoutFunc">退出登录</el-button>
  </div>
</template>

<script setup>
import cookie from "js-cookie";
import {ref,onMounted} from "vue";
import {useRouter} from "vue-router";
import {getUserInfo, logout} from "@/api/user";
const router = useRouter()

const info = ref({})

const getUserInfoFunc = async () => {
  const res = await getUserInfo()
  if (res.code === 0) {
    info.value = res.data
  }
}

const logoutFunc = async () => {
  await logout()
  cookie.remove('x-token')
  router.push("/login")
}

onMounted(()=>{
  if (!cookie.get('x-token')) {
    router.push("/login")
  }else{
    getUserInfoFunc()
  }
})
</script>


<style scoped  lang="scss">


</style>
