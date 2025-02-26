<script setup lang="ts">
const ping = async () => {
  const res = await $fetch('/api/action/ping')
  console.log(res)
}
</script>

<template>
  <el-button type="primary" @click="ping">Ping</el-button>
</template>

<style lang="scss" scoped></style>
