<script setup lang="ts">
const router = useRouter()

onMounted(() => {
  router.push('/login')
})
</script>

<template>
  <el-button type="primary">Ping</el-button>
</template>

<style lang="scss" scoped></style>
