import { defineStore } from 'pinia'
import {ref} from "vue";

export const useUserStore = defineStore('user', () => {
  const userInfo = ref({
    username:"123"
  })

  return {
    userInfo
  }
})
