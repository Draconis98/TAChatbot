<script setup>
import {onMounted, ref} from 'vue'
import {Back, Search, Star} from "@element-plus/icons-vue";
import axios from "axios";

const backendURL = ref('https://callme.agileserve.org.cn:30941')
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
  let card_id = window.location.href.split('cardID=')[1]
  axios.get(backendURL.value + '/get/card?cardID=' + card_id)
      .then((response) => {
        card_data = response.data.card;
        console.log(card_data);
        qa_id = card_data.questions;
        console.log(qa_id)
        // 遍历qa_id
        for (let i = 0; i < qa_id.length; i++) {
          console.log(qa_id[i])
          axios.get(backendURL.value + '/get/question?questionID=' + qa_id[i])
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
  <div class="layout">
    <el-container class="container">
      <el-header class="navbar">
        <el-icon size="25" class="logo"><House/></el-icon>
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
<!--        <el-button :icon="Back" @click="back" circle></el-button>-->
<!--      </el-header>-->
      <el-container>
        <el-header>
          <el-menu class="menu" mode="horizontal" center :ellipsis="false">
            <el-menu-item index="1" @click="back" :icon="Back"></el-menu-item>
          </el-menu>
        </el-header>
        <el-container>
          <el-main>
<!--            <el-divider></el-divider>-->
            <el-card class="box-card" shadow="never">
              <template #header>
                <div class="card-header">
                  <span>#{{ card_data.id }}</span>
                  <el-button v-if="isStar.value===false" :icon="Star" circle disabled @click="toggleStar"></el-button>
                  <el-button v-else :icon="Star" type="warning" circle disabled @click="toggleStar"></el-button>
                </div>
              </template>
              <!-- 对qa数组内容按顺序输出 -->

              <el-card v-for="item in qa_data" :key="item" shadow="hover">
                {{ item }}
              </el-card>

            </el-card>
          </el-main>
        </el-container>
      </el-container>
      <el-footer class="footer">
        <div class="affix">
          <el-affix position="bottom" :offset="10">
            本平台仅供学习使用，请勿做其他用途；生成式回答，内容仅供参考。
          </el-affix>
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
  height: 95vh;
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

.container {
  height: 90vh;
  width: 75vw;
  justify-content: space-between;
}

.footer {
  height: 10px;
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