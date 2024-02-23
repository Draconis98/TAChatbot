<script setup>
import {onMounted, ref} from 'vue'
import {Search, Star} from "@element-plus/icons-vue";
import axios from "axios";

const backendURL = ref('10.30.19.40:8080')
const input = ref('')
const isStar = ref(false)
const count = ref(0)
const checked = ref(false)
const username = ref(null)
let cards = ref([])

const load = () => {
  console.log('加载更多');
  count.value += 3;
}

function onChange(status) {
  checked.value = status;
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

function toggleStar() {
  if (isStar.value === false) {
    isStar.value = true;
    console.log('收藏');
  } else {
    isStar.value = false;
    console.log('取消收藏');
  }
}

function latest() {
  getCardsSortedByLatest();
  console.log('按最新排序');
}

function hottest() {
  getCardsSortedByHottest();
  console.log('按最热排序');
}

function createNewQuestion() {
  console.log('创建新问题');
  if (!username.value) {
    alert('请先登录');
    authenticate();
    return;
  }
  console.log('创建一个新的卡片来记录问题');
  // let userID = window.location.href.split('userID=')[1];
  let userID = window.sessionStorage.getItem('userID');
  console.log(userID)
  axios.get(backendURL.value + '/new/card?userID=' + userID)
      .then((response) => {
        console.log(response.data);
        window.sessionStorage.setItem('cardID', response.data.cardID);
        window.location.href = '/new/question';
      })
      .catch((error) => {
        alert('服务器出现错误，无法创建新问题，请重新尝试');
        console.log('创建新卡片失败：', error);
      });
}

function getCardsSortedByLatest() {
  axios.get(backendURL.value + '/show/latest')
      .then((response) => {
        cards.value = response.data.card_list;
        console.log(cards.value);
      })
      .catch((error) => {
        alert('获取卡片失败\n' + error);
      });
}

function getCardsSortedByHottest() {
  axios.get(backendURL.value + '/show/hottest')
      .then((response) => {
        cards.value = response.data.card_list;
        console.log(cards.value);
      })
      .catch((error) => {
        alert('获取卡片失败\n' + error);
      });
}

function card_details(cardID) {
  console.log('查看卡片详');
  console.log(cardID);
  window.sessionStorage.setItem('cardID', cardID);
  window.location.href = '/card/details';
}

onMounted(() => {
  axios.get(backendURL.value + '/show/')
      .then((response) => {
        cards.value = response.data.card_list;
        console.log(cards.value);
      })
      .catch((error) => {
        alert('获取卡片失败\n' + error);
      });
})
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
            <el-menu-item index="1" @click="latest">按最新排序</el-menu-item>
            <el-menu-item index="2" @click="hottest" disabled>按最热排序</el-menu-item>
            <el-menu-item index="3" @click="createNewQuestion">创建新问题</el-menu-item>
          </el-menu>
        </el-header>
        <el-container>
          <el-main>
            <ul v-infinite-scroll="load" class="infinite-list" style="overflow: auto">
              <li v-for="card in cards" :key="card.id" class="infinite-list-item">
                <el-card class="box-card" shadow="hover" @click="card_details(card.id)">
                  <template #header>
                    <div class="card-header">
                      <span>#{{ card.id.slice(10) }}</span>
                      <el-button-group>
                        <!-- TODO: 如果用户已经收藏了这个问题，那么显示实心的星星，否则显示空心的星星 -->
                        热度{{ card.favoritesCount }}
                      </el-button-group>
                    </div>
                  </template>
                  <div class="card-body">
                    <div v-for="item in card.content">
                      {{ item }}
                    </div>
                  </div>
                  <template #footer>
                    <div class="card-footer">
                      <span>{{ card.create_at }}</span>
                      <el-tag effect="light" type="info">
<!--                        tag1-->
                      </el-tag>
                    </div>
                  </template>
                </el-card>
              </li>
            </ul>
          </el-main>
        </el-container>
      </el-container>
      <el-footer class="footer">
        <div class="claim">
          本平台仅供学习使用，请勿做其他用途；生成式回答，内容仅供参考。
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
  height: 3px;
}

.card-body {
  display: -webkit-box;
  justify-content: space-between;
  align-items: center;
  height: 70px;
  width: 100%;
  overflow: hidden;
  text-overflow: ellipsis;
  //white-space: nowrap;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 4;
}

.card-footer {
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

.box-card {
  width: 100%;
}

.infinite-list {
  height: 80vh;
  padding: 0;
  margin: 0;
  list-style: none;
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