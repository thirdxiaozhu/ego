<template>
  <aside class="sidebar" :class="{ collapsed: isCollapsed }">
    <el-menu
      default-active="1"
      class="el-menu-vertical-demo"
      background-color="#1e293b"
      text-color="#fff"
      active-text-color="#409eff"
      :collapse="isCollapsed"
    >
      <el-menu-item index="1">
        <i class="el-icon-s-home"></i>
        <template #title>首页</template>
      </el-menu-item>

      <el-sub-menu index="2">
        <template #title>
          <i class="el-icon-user"></i>
          <span>用户管理</span>
        </template>
        <el-menu-item index="2-1">用户列表</el-menu-item>
        <el-menu-item index="2-2">角色管理</el-menu-item>
        <el-menu-item index="2-3">权限设置</el-menu-item>
      </el-sub-menu>

      <el-sub-menu index="3">
        <template #title>
          <i class="el-icon-document"></i>
          <span>内容管理</span>
        </template>
        <el-menu-item index="3-1">文章列表</el-menu-item>
        <el-menu-item index="3-2">分类管理</el-menu-item>
        <el-menu-item index="3-3">标签管理</el-menu-item>
      </el-sub-menu>

      <el-sub-menu index="4">
        <template #title>
          <i class="el-icon-s-order"></i>
          <span>订单管理</span>
        </template>
        <el-menu-item index="4-1">订单列表</el-menu-item>
        <el-menu-item index="4-2">退款处理</el-menu-item>
        <el-menu-item index="4-3">数据分析</el-menu-item>
      </el-sub-menu>

      <el-sub-menu index="5">
        <template #title>
          <i class="el-icon-s-marketing"></i>
          <span>营销活动</span>
        </template>
        <el-menu-item index="5-1">优惠券</el-menu-item>
        <el-menu-item index="5-2">限时折扣</el-menu-item>
        <el-menu-item index="5-3">满减活动</el-menu-item>
      </el-sub-menu>

      <el-sub-menu index="6">
        <template #title>
          <i class="el-icon-s-tools"></i>
          <span>系统设置</span>
        </template>
        <el-menu-item index="6-1">基础设置</el-menu-item>
        <el-menu-item index="6-2">通知设置</el-menu-item>
        <el-menu-item index="6-3">日志管理</el-menu-item>
      </el-sub-menu>
    </el-menu>
  </aside>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue';

const isCollapsed = ref(false);

// 监听窗口大小变化，在小屏幕上自动折叠侧边栏
const handleResize = () => {
  if (window.innerWidth < 768) {
    isCollapsed.value = true;
  } else {
    isCollapsed.value = false;
  }
};

onMounted(() => {
  handleResize();
  window.addEventListener('resize', handleResize);
});

// 接收来自父组件的折叠/展开侧边栏的事件
const props = defineProps({
  collapsed: {
    type: Boolean,
    default: false
  }
});

watch(() => props.collapsed, (newVal) => {
  isCollapsed.value = newVal;
});
</script>

<style scoped>
.sidebar {
  background-color: #1e293b;
  width: 200px;
  transition: width 0.3s;
  height: 100%;
  overflow-y: auto;
}

.sidebar.collapsed {
  width: 64px;
}

.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
}

.el-menu-item,
.el-sub-menu__title {
  height: 50px;
  line-height: 50px;
}

.el-sub-menu .el-menu-item {
  height: 45px;
  line-height: 45px;
}

.el-menu-item-group__title {
  padding: 0;
}

.el-menu--collapse .el-sub-menu__title span {
  height: 0;
  width: 0;
  overflow: hidden;
  visibility: hidden;
  display: inline-block;
}

.el-menu--collapse .el-sub-menu .el-sub-menu__title i {
  margin-right: 0;
}

.el-menu--collapse .el-sub-menu .el-sub-menu__title .el-sub-menu__icon-arrow {
  display: none;
}

.el-menu-item:hover,
.el-sub-menu__title:hover {
  background-color: rgba(255, 255, 255, 0.1);
}

.el-menu-item.is-active {
  background-color: rgba(64, 158, 255, 0.1);
}
</style>
