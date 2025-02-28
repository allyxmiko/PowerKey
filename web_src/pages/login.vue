<script setup lang="ts">
import { ref } from 'vue'

const password = ref('')
const hasChecked = ref(false)

const onSubmit = () => {
  if (password.value === '') {
    ElMessage.error('密码不能为空!')
    return
  }
}
</script>

<template>
  <div class="flex items-center justify-center h-screen bg-blue-50">
    <div class="flex flex-col items-center mb-30">
      <div class="text-6xl font-bold mb-25 text-shadow powerkey">
        <span v-for="v in 'PowerKey'" class="text-blue-500">{{ v }}</span>
      </div>
      <div class="p-20 w-120 form-card">
        <div>
          <el-input
            placeholder="请输入登录密码..."
            show-password
            type="password"
            v-model="password"
            class="w-full u-font e-height"
          ></el-input>
        </div>
        <div class="mt-6">
          <el-checkbox v-model="hasChecked" style="color: #a3a6ad"
            >记住密码</el-checkbox
          >
        </div>
        <div class="mt-6">
          <el-button
            @click="onSubmit"
            type="primary"
            round
            class="w-full u-font bg-blue-300 text-white e-height"
            >登 录</el-button
          >
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.text-shadow {
  text-shadow: 2px 2px 4px rgba(20, 141, 215, 0.3);
}

.e-height {
  height: 45px;
}

.form-card {
  border-radius: 8px;
  box-shadow: rgba(142, 197, 255, 0.2) 0 0 10px;
  backdrop-filter: blur(10px);
  background-color: rgba(255, 255, 255, 0.5);
}

.powerkey {
  user-select: none;
}

.u-font {
  font-weight: bold;
  letter-spacing: 0.5px;
}
</style>
