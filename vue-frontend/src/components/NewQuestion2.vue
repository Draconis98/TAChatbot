<script setup>
import axios from "axios";
import {onMounted, ref} from "vue";

const textarea = ref('');
const cards = ref([]);

function submit() {
  console.log('提交问题');
  let username = window.location.href.split('username=')[1];
  let question = textarea.value;

  if (textarea.value.trim() !== '') {
    cards.value.push(question);
    textarea.value = '';
  }

  axios.get(`http://127.0.0.1:8080/new-question?username=${username}&question=${question}`)
      .then((response) => {
        console.log(response.data);
      })
      .catch((error) => {
        alert('提交问题失败，请重新尝试\n' + error);
        console.log('提交问题失败：', error);
      });
}
</script>

<template>
  <div class="common-layout">
    <el-container class="container1">
      <el-header class="header">
        Header
      </el-header>
      <el-divider/>
      <el-container class="container2">
        <el-aside width=5% class="aside">
        </el-aside>
        <el-main class="main">
            <el-card v-for="card in cards" :key="card">
              <p>{{ card }}</p>
            </el-card>
        </el-main>
        <el-aside width=5% class="aside">
        </el-aside>
      </el-container>
      <el-footer class="footer">
        <el-input
            v-model="textarea"
            autosize
            type="textarea"
            placeholder="请输入问题"
            clearable
        />
        <el-button type="primary" class="submit" v-if="textarea === ''" disabled>
          提问
        </el-button>
        <el-button type="primary" class="submit" @click="submit" v-else>
          提问
        </el-button>
      </el-footer>
      <el-text>本平台仅供学习用途，请勿做其他用途。答案由大语言模型生成，内容仅供参考。</el-text>
    </el-container>
  </div>
</template>

<style scoped>
.container1 {
  height: 90vh;
  width: 90vw;
  justify-content: space-between;
}

.container2 {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 30px;
}

.aside {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.main {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.footer {
  display: flex;
  justify-content: space-between;
  justify-items: center;
  align-items: center;
}

.submit {
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 5px;
}
</style>