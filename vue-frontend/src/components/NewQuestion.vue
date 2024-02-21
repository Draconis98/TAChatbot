<script setup>
import {onMounted, ref} from 'vue';
import axios from "axios";
import {Back} from "@element-plus/icons-vue";

// 使用ref来创建一个响应式的src属性
const iframeSrc = ref('');

onMounted(() => {
  // 获取当前页面的URL
  const userID = window.sessionStorage.getItem('userID');
  const cardID = window.sessionStorage.getItem('cardID');

  iframeSrc.value = `http://10.30.19.40:8082?userID=${userID}&cardID=${cardID}`
});
</script>

<template>
  <el-container>
    <el-main class="container-main">
      <div>
        <iframe
            :src="iframeSrc"
            frameborder="0"
            class="iframe"
            style="border: 0; margin: 0; padding: 0; overflow: hidden; display: block;"
        />
      </div>
    </el-main>
    <el-footer class="container-footer">
      <div class="affix">
        <el-affix position="bottom" :offset="10">
          本平台仅供学习使用，请勿做其他用途；生成式回答，内容仅供参考。
        </el-affix>
      </div>
    </el-footer>
  </el-container>
</template>

<style scoped>
.container-main {
  height: 90vh;
  width: 75vw;
  justify-content: space-between;
}

.container-footer {
  height: 5vh;
  width: 75vw;
  justify-content: space-between;
}


.iframe {
  width: 75vw;
  height: 90vh;
  margin: 0 0;
}

.affix {
  text-align: center;
}
</style>