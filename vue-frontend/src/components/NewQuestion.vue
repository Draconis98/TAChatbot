<script setup>
import {onMounted, ref} from 'vue';
import axios from "axios";

// 使用ref来创建一个响应式的src属性
const iframeSrc = ref('');

onMounted(() => {
  // 获取当前页面的URL
  const IDs = window.location.href.split('?')[1]

  iframeSrc.value = `http://localhost:8082?${IDs}`
});
</script>

<template>
  <div>
    <iframe
        :src="iframeSrc"
        frameborder="0"
        class="iframe"
    />
  </div>
  <div class="footer-content">
    <p>© 2021-2022</p>
  </div>
</template>

<style scoped>
.iframe {
  width: 75vw;
  height: 90vh;
  margin: 0 0;
}
</style>