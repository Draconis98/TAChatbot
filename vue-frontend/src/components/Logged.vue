<script setup>
import { ref, onMounted } from 'vue';

// 使用ref创建一个响应式的username变量
const username = ref('');

onMounted(() => {
  // 当组件挂载完成后，从URL中获取username参数并更新username变量
  const params = new URLSearchParams(window.location.search);
  const user = params.get('username');
  if (user) {
    username.value = user;
  } else {
    console.log('没有username');
  }
});
</script>

<template>
  <div class="logged">
    <img src="@/assets/logo.svg" alt="Logo" width="20" height="20" class="logo"/>
    <input type="text" placeholder="search...">
    <div class="user-info">
      <span>{{ username }}</span>
    </div>
  </div>
</template>

<style scoped>

</style>
