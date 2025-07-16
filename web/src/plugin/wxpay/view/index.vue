<template>
  <div>
    <!--  此代码为vue组件示例 -->
    <el-button @click="getCode">
      获取支付二维码
    </el-button>
    <qrcode-vue
      v-if="value"
      :value="value"
      :size="size"
    />
  </div>
</template>

<script setup>

import QrcodeVue from 'qrcode.vue'
import {ref} from 'vue'

import {getPayCode,getOrderById} from "@/plugin/wxpay/api/api";

var value = ref("")
var codeId = ref("")
var size = ref(200)

let timer = null
const getOrder = async () => {
  const res = await getOrderById({orderID:codeId.value})
  if(res.data.TradeState === "SUCCESS"){
    // 表明支付成功 做后面的业务即可
    clearInterval(timer)
    value.value = ""
    codeId.value = ""
  }
}
const getCode = async () => {
  const res = await getPayCode()
  value.value = res.data.codeUrl
  codeId.value = res.data.codeId
  timer = setInterval(getOrder, 1000)
}

// template 部分

</script>
