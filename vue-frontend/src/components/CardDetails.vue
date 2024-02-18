<script setup>
import {onMounted, ref} from 'vue'
import {Back, Search, Star} from "@element-plus/icons-vue";
import axios from "axios";

const input = ref('')
const username = ref('')
const isStar = ref(false)
let card_data = ref({})
let qa_id = ref([])
let qa_data = ref([])
let count = ref(0)

function user() {
  console.log('用户信息');
}

function authenticate() {
  axios.get('http://127.0.0.1:8080/auth')
      .then((response) => {
            console.log(response.data);
            // 解析url
            let url = new URL(response.data.url);
            console.log(url);
            window.location.href = url.href;
          }
      )
      .catch((error) => {
        alert('获取认证URL失败：' + error)
        console.log('获取认证URL失败：', error);
      });
}

onMounted(() => {
      console.log('mounted');
      let userID = window.location.href.split('userID=')[1];
      axios.get('http://127.0.0.1:8080/get/username?userID=' + userID)
          .then((response) => {
            console.log(response.data);
            username.value = response.data.username;
          }).catch((error) => {
        if (error.response.status === 401) {
          console.log('用户不存在');
          alert('用户不存在，请重新登录');
          authenticate();
        } else if (error.response.status === 500) {
          console.log('获取用户名失败');
          alert('获取用户名失败, 请重新登录');
          authenticate();
        }
      });
      console.log(username.value);
    }
)

onMounted(() => {
  let card_id = window.location.href.split('cardID=')[1]
  axios.get('http://127.0.0.1:8080/get/card?cardID=' + card_id)
      .then((response) => {
        card_data = response.data.card;
        console.log(card_data);
        qa_id = card_data.questions;
        console.log(qa_id)
        // 遍历qa_id
        for (let i = 0; i < qa_id.length; i++) {
          axios.get('http://127.0.0.1:8080/get/question?questionID=' + qa_id[i])
              .then((response) => {
                qa_data.value.push("Q: " + response.data.question);
                qa_data.value.push("A: " + response.data.answer);
              }).catch((error) => {
            console.log('获取问题失败\n', error);
          });
        }
      }).catch((error) => {
    console.log('获取卡片失败\n', error);
  });
})

function toggleStar() {
  if (isStar.value === false) {
    isStar.value = true;
    console.log('收藏');
  } else {
    isStar.value = false;
    console.log('取消收藏');
  }
}

function back() {
  window.history.back();
}
</script>

<template>
  <div class="common-layout">
    <el-container class="container">
      <el-header class="header">
        <el-button :icon="Back" @click="back" circle></el-button>
        <!--        <el-icon size="25" class="logo"><House/></el-icon>-->
        <el-button class="login-text" v-if="!username" @click="authenticate" text>Login</el-button>
        <el-button class="login-text" v-else @click="user" text>{{ username }}</el-button>
      </el-header>
      <el-container>
        <el-container>
          <el-main>
            <el-divider></el-divider>
            <el-card class="box-card" shadow="never">
              <template #header>
                <div class="card-header">
                  <span>#{{ card_data.id }}</span>
                  <el-button :icon="Star" @click="toggleStar" circle disabled></el-button>
                </div>
              </template>
              <el-card v-for="qa in qa_data" shadow="hover">
                {{ qa }}
              </el-card>

            </el-card>
          </el-main>
        </el-container>
      </el-container>
    </el-container>
  </div>
  <div class="affix">
    <el-affix position="bottom" :offset="10">
      本平台仅供学习使用，请勿做其他用途；生成式回答，内容仅供参考。
    </el-affix>
  </div>
</template>


<style scoped>
.container {
  height: 90vh;
  width: 75vw;
  justify-content: space-between;
}

.common-layout {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 5px;
}

.logo {
  width: 25px;
  height: 25px;
  justify-items: center;
  align-items: center;
  margin-right: 10px;
}

.login-text {
  cursor: pointer;
  color: blue;
  text-decoration: underline;
  padding: 5px 10px;
  margin-left: auto;
}

.w-50 .m-2 {
  width: 80%;
  padding: 0 10px;
}

.infinite-list .infinite-list-item {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 200px;
  background: var(--el-color-primary-light-9);
  margin: 5px;
  //color: var(--el-color-primary);
}

.infinite-list .infinite-list-item + .list-item {
  margin-top: 10px;
}

.affix {
  text-align: center;
}
</style>