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
        style="border: 0; margin: 0; padding: 0; overflow: hidden; display: block;"
    />
  </div>
  <div class="affix">
    <el-affix position="bottom" :offset="10">
      本平台仅供学习使用，请勿做其他用途；生成式回答，内容仅供参考。
    </el-affix>
  </div>
</template>

<style scoped>
.iframe {
  width: 75vw;
  height: 90vh;
  margin: 0 0;
}

.affix {
  text-align: center;
}
</style>