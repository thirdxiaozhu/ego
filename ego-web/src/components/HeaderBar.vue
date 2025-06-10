<template>
  <header class="header-bar">
    <div class="header-left">
      <button class="sidebar-toggle" @click="toggleSidebar">
        <i class="el-icon-s-fold"></i>
      </button>
      <div class="logo">
        <span>管理系统</span>
      </div>
    </div>

    <div class="header-right">
      <div class="search-box">
        <el-input
          placeholder="搜索..."
          prefix-icon="el-icon-search"
          size="small"
        >
        </el-input>
      </div>

      <div class="notification">
        <el-badge value="3">
          <button class="header-icon">
            <i class="el-icon-bell"></i>
          </button>
        </el-badge>
      </div>

      <el-button @click="logoutFunc">退出登录</el-button>
      <div class="user-menu">
        <el-dropdown trigger="click">
          <div class="user-info">
            <img src="https://picsum.photos/id/1005/40/40" alt="用户头像" class="user-avatar" />
            <span class="user-name">管理员</span>
            <i class="el-icon-caret-bottom"></i>
          </div>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item>
                <i class="el-icon-user"></i>
                <span>个人资料</span>
              </el-dropdown-item>
              <el-dropdown-item>
                <i class="el-icon-setting"></i>
                <span>设置</span>
              </el-dropdown-item>
              <el-dropdown-item divided>
                <i class="el-icon-switch-button"></i>
                <span>退出登录</span>
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </div>
  </header>
</template>

<script setup>
import { ref, defineEmits } from 'vue';
import {logout} from "@/api/user.js";
import cookie from "js-cookie";
import {useRouter} from "vue-router";
const router = useRouter()

const emits = defineEmits(['toggle-sidebar']);

const toggleSidebar = () => {
  emits('toggle-sidebar');
};

const logoutFunc = async () => {
  await logout()
  cookie.remove('x-token')
  router.push("/login")
}
</script>

<style scoped>
.header-bar {
  height: 60px;
  background-color: #fff;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.1);
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
  z-index: 100;
  position: relative;
}

.header-left {
  display: flex;
  align-items: center;
}

.sidebar-toggle {
  background: none;
  border: none;
  cursor: pointer;
  padding: 8px;
  margin-right: 15px;
  transition: color 0.3s;
}

.sidebar-toggle:hover {
  color: #409eff;
}

.logo {
  font-size: 18px;
  font-weight: 500;
  color: #303133;
}

.header-right {
  display: flex;
  align-items: center;
}

.search-box {
  width: 200px;
  margin-right: 20px;
}

.notification {
  margin-right: 20px;
  position: relative;
}

.header-icon {
  background: none;
  border: none;
  cursor: pointer;
  padding: 8px;
  transition: color 0.3s;
}

.header-icon:hover {
  color: #409eff;
}

.user-menu {
  cursor: pointer;
  display: flex;
  align-items: center;
}

.user-info {
  display: flex;
  align-items: center;
}

.user-avatar {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  margin-right: 10px;
  object-fit: cover;
}

.user-name {
  margin-right: 5px;
  font-size: 14px;
  color: #303133;
}

.el-dropdown-menu__item {
  display: flex;
  align-items: center;
}

.el-dropdown-menu__item i {
  margin-right: 8px;
  width: 16px;
}
</style>
