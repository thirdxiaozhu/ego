
<template>
    <div class="app-container">
      <el-container>
        <el-header>
          <HeaderBar />
</el-header>
        <el-container>
          <el-aside width="200px"> <Sidebar/></el-aside>
          <el-main>

            <el-row>
              <el-col :span="8">

                <el-select
                  v-model="selectedValue"
                  clearable
                  placeholder="请选择"
                >
                  <el-option
                    v-for="item in modelList"
                    :key="item.modelName"
                    :label="item.modelName"
                    :value="item.modelName"
                  />
                </el-select>
              </el-col>
            </el-row>
          </el-main>
        </el-container>
      </el-container>
    </div>
</template>

<script setup>
import cookie from "js-cookie";
import {ref,onMounted} from "vue";
import {useRouter} from "vue-router";
import {getUserInfo, logout} from "@/api/user";
const router = useRouter()

import HeaderBar from '@/components/HeaderBar.vue';
import Sidebar from '@/components/Sidebar.vue';
import {getEgoModelAll, getEgoModelList} from "@/api/egoModel.js";

// 模拟数据
const stats = [
  { title: '用户总数', value: '1,284', change: '+12%', isPositive: true },
  { title: '今日访问', value: '328', change: '+8%', isPositive: true },
  { title: '待处理任务', value: '16', change: '+3', isPositive: false },
  { title: '系统状态', value: '正常', change: '运行中', isPositive: true }
];

const activities = [
  { icon: 'user-plus', user: '张三', action: '创建了新用户', target: '李四', time: '10分钟前' },
  { icon: 'document', user: '王五', action: '更新了文档', target: '系统使用手册', time: '30分钟前' },
  { icon: 'warning', user: '系统', action: '发出警告:', target: '服务器负载过高', time: '1小时前' }
];

const announcements = [
  {
    title: '系统维护通知',
    content: '本系统将于2023年6月15日凌晨2:00-4:00进行系统维护，期间可能无法正常访问。',
    time: '2023-06-10'
  },
  {
    title: '新功能上线',
    content: '数据分析模块已更新，支持更多维度的数据可视化和导出功能。',
    time: '2023-06-05'
  }
];

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

const modelList = ref([])
const selectedValue = ref('')

onMounted(async () => {
  if (!cookie.get('x-token')) {
    router.push("/login")
  }else{
    getUserInfoFunc()
  }

  const ret = await getEgoModelAll()
  if(ret.code === 0) {
    modelList.value = ret.data
    console.log(ret.data)
  }
})
</script>


<style scoped  lang="scss">
.app-container {
  height: 100vh;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.main-content {
  display: flex;
  flex: 1;
  overflow: hidden;
}

.content-wrapper {
  flex: 1;
  padding: 20px;
  overflow-y: auto;
  background-color: #f5f7fa;
}

.content-header {
  margin-bottom: 20px;
}

.content-header h1 {
  margin: 0;
  font-size: 24px;
  font-weight: 500;
  color: #303133;
}

.dashboard-card {
  margin-bottom: 20px;
}

.stats-cards {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 20px;
}

.stat-card {
  padding: 20px;
  border-left: 4px solid;
  transition: all 0.3s;
}

.stat-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.stat-title {
  font-size: 14px;
  color: #606266;
  margin-bottom: 8px;
}

.stat-value {
  font-size: 24px;
  font-weight: 500;
  color: #303133;
  margin-bottom: 8px;
}

.stat-change {
  font-size: 12px;
  display: flex;
  align-items: center;
}

.positive {
  color: #67c23a;
}

.negative {
  color: #f56c6c;
}

.two-column-layout {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 20px;
}

.content-card {
  height: 100%;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.activity-list,
.announcement-list {
  padding: 10px 0;
}

.activity-item {
  display: flex;
  padding: 15px 0;
  border-bottom: 1px solid #f2f3f5;
}

.activity-item:last-child {
  border-bottom: none;
}

.activity-icon {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background-color: #f0f2f5;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 15px;
}

.activity-icon i {
  color: #409eff;
}

.activity-content {
  flex: 1;
}

.activity-content p {
  margin: 0;
}

.activity-user,
.activity-target {
  font-weight: 500;
}

.activity-time {
  font-size: 12px;
  color: #909399;
  margin-top: 5px;
}

.announcement-item {
  padding: 15px 0;
  border-bottom: 1px solid #f2f3f5;
}

.announcement-item:last-child {
  border-bottom: none;
}

.announcement-title {
  font-size: 16px;
  font-weight: 500;
  margin-bottom: 8px;
}

.announcement-content {
  font-size: 14px;
  color: #606266;
  margin-bottom: 8px;
}

.announcement-time {
  font-size: 12px;
  color: #909399;
}

/* 响应式布局 */
@media (max-width: 1024px) {
  .stats-cards {
    grid-template-columns: repeat(2, 1fr);
  }

  .two-column-layout {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 768px) {
  .stats-cards {
    grid-template-columns: 1fr;
  }
}

</style>
