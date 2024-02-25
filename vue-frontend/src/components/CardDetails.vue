<script setup>
import {onMounted, ref} from 'vue'
import {ArrowLeftBold, CircleCloseFilled, Search, Star, SuccessFilled} from "@element-plus/icons-vue";
import axios from "axios";

const backendURL = ref('https://callme.agileserve.org.cn:30941')
const input = ref('')
const username = ref(null)
const isStar = ref(false)
const card_id = ref('')
let card_data = ref({})
let qa_id = ref([])
let qa_data = ref([])
let count = ref(0)

const load = () => {
  console.log('加载更多');
  count.value += 3;
}

function user() {
  console.log('用户信息');
}

function authenticate() {
  axios.get(backendURL.value + '/auth')
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


function myquestion() {
  // TODO: 跳转到我的问题页面
}

function myfavorite() {
  // TODO: 跳转到我的收藏页面
}

function logout() {
  console.log('退出登录');
  window.sessionStorage.removeItem('userID');
  window.sessionStorage.removeItem('username');
  window.sessionStorage.removeItem('cardID');
  window.location.href = '/';
}

onMounted(() => {
      console.log('mounted');
      // userID = window.location.href.split('userID=')[1];
      let userID = window.sessionStorage.getItem('userID');
      console.log(userID)
      if (userID != null) {
        axios.get(backendURL.value + '/get/username?userID=' + userID)
            .then((response) => {
              console.log(response.data);
              window.sessionStorage.setItem('username', response.data.username);
              username.value = window.sessionStorage.getItem('username');
            }).catch((error) => {
          if (error.response.status === 401) {
            alert('用户不存在，请重新登录');
            authenticate();
          } else if (error.response.status === 500) {
            alert('获取用户名失败, 请重新登录');
            authenticate();
          }
        });
      }

      // console.log(username.value);
    }
)

onMounted(() => {
  // let card_id = window.location.href.split('cardID=')[1]
  card_id.value = window.sessionStorage.getItem('cardID');
  console.log("card id: ", card_id.value)
  axios.get(backendURL.value + '/get/card?cardID=' + card_id.value)
      .then((response) => {
        card_data = response.data.card;
        console.log(card_data);
        qa_id = card_data.questions;
        console.log(qa_id);
        for (let i = 0; i < qa_id.length; i++) {
          console.log(qa_id[i])
          axios.get(backendURL.value + '/get/question?questionID=' + qa_id[i])
              .then((response) => {
                qa_data.value[i] = response.data;
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
  <div class="layout">
    <el-container class="container">
      <el-header class="navbar">
        <el-icon size="25" class="logo">
          <House/>
        </el-icon>
        <el-input
            v-model="input"
            class="searchbar"
            size="large"
            placeholder="Search..."
            :prefix-icon="Search"
            clearable
            disabled
        />
        <el-button class="login-text" v-if="!username" @click="authenticate" text>Login</el-button>
        <el-dropdown>
          <span class="el-dropdown-link">
            {{ username }}
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item @click="myquestion" disabled>我的问题</el-dropdown-item>
              <el-dropdown-item @click="myfavorite" disabled>我的收藏</el-dropdown-item>
              <el-dropdown-item @click="logout" divided>退出登录</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </el-header>
      <el-container>
        <el-header>
          <el-menu class="menu" mode="horizontal" center :ellipsis="false">
            <el-menu-item index="1" @click="back">
              <el-icon>
                <ArrowLeftBold/>
              </el-icon>
              返回首页
            </el-menu-item>
          </el-menu>
        </el-header>
        <el-container>
          <el-main>
            <el-card class="box-card" shadow="never">
              <template #header>
                <div class="card-header">
                  <span>#{{ card_id.slice(10) }}</span>
                  <el-button v-if="isStar.value===false" :icon="Star" circle disabled @click="toggleStar"></el-button>
                  <el-button v-else :icon="Star" type="warning" circle disabled @click="toggleStar"></el-button>
                </div>
              </template>
              <el-scrollbar height="70vh">
                <el-timeline>
                  <el-timeline-item v-for="item in qa_data" :key="item" :timestamp="item.create_at"
                                    placement="top">
                    <el-card shadow="hover">
                      <div class="question">{{ item.question }}</div>
                      <el-divider></el-divider>
                      <div>{{ item.answer }}
                      <el-tooltip v-if="item.likes === 1" content="原提问者对这个答案表示满意" placement="right">
                        <el-icon class="like" color="green">
                          <SuccessFilled/>
                        </el-icon>
                      </el-tooltip>
                      <el-tooltip v-else-if="item.likes === 2" content="原提问者对这个答案表示不满" placement="right">
                        <el-icon class="like" color="blue">
                          <CircleCloseFilled/>
                        </el-icon>
                      </el-tooltip>
                      </div>
                    </el-card>
                  </el-timeline-item>
                </el-timeline>
              </el-scrollbar>
            </el-card>
          </el-main>
        </el-container>
      </el-container>
      <el-footer class="footer">
        <div class="claim">
          本平台仅供学习使用，请勿做其他用途；生成式回答，内容仅供参考。模型由Baichuan2-7B经过指令微调得到，使用Apache 2.0协议。
        </div>
      </el-footer>
    </el-container>
  </div>
</template>


<style scoped>
.layout {
  display: -webkit-flex;
  justify-content: space-between;
  align-items: center;
  padding: 0;
  margin: 0;
  height: 100vh;
  width: 100vw;
}

.container {
  height: 100vh;
  width: 100vw;
  justify-content: space-between;
  padding: 0 0;
  margin: 0;
}

.navbar {
  display: -webkit-inline-flex;
  height: 5vh;
  align-items: center;
}

.searchbar {
  height: 40px;
}

.login-text {
  cursor: pointer;
  color: #329eff;
  text-decoration: underline;
  padding: 0 10px;
}

.el-dropdown-link {
  cursor: pointer;
  color: #329eff;
  text-decoration: underline;
  display: flex;
  align-items: center;
  padding: 0 10px;
}

.footer {
  height: 5vh;
}

.menu {
  display: -webkit-flex;
  justify-content: start;
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

.question {
  font-weight: bold;
}

.like {
  align-items: center;
  justify-content: center;
  align-content: center;
  justify-items: center;
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

.claim {
  text-align: center;
  font-size: 14px;
}
</style>