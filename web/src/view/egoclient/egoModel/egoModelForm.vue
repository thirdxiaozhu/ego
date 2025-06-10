
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="模型供应商:" prop="modelProvider">
    <el-select v-model="formData.modelProvider" placeholder="请选择模型供应商" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in model-providerOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="服务类型:" prop="modelType">
    <el-select v-model="formData.modelType" placeholder="请选择服务类型" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in model-typeOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="模型名称:" prop="modelName">
    <el-input v-model="formData.modelName" :clearable="true" placeholder="请输入模型名称" />
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createEgoModel,
  updateEgoModel,
  findEgoModel
} from '@/api/egoclient/egoModel'

defineOptions({
    name: 'EgoModelForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const model-providerOptions = ref([])
const model-typeOptions = ref([])
const formData = ref({
            modelProvider: '',
            modelType: '',
            modelName: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findEgoModel({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    model-providerOptions.value = await getDictFunc('model-provider')
    model-typeOptions.value = await getDictFunc('model-type')
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createEgoModel(formData.value)
               break
             case 'update':
               res = await updateEgoModel(formData.value)
               break
             default:
               res = await createEgoModel(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
