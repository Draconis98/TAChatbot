<script setup>
import {onMounted, ref} from 'vue'
import {Search, Star} from "@element-plus/icons-vue";
import axios from "axios";

const backendURL = ref('https://callme.agileserve.org.cn:30941')
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
  axios.get(backendURL + '/auth')
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
        axios.get(backendURL + '/get/username?userID=' + userID)
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
  axios.get(backendURL + '/new/card?userID=' + userID)
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
  axios.get(backendURL + '/show/latest')
      .then((response) => {
        cards.value = response.data.card_list;
        console.log(cards.value);
      })
      .catch((error) => {
        alert('获取卡片失败\n' + error);
      });
}

function getCardsSortedByHottest() {
  axios.get(backendURL + '/show/hottest')
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
  window.location.href = '/card/details?cardID=' + cardID;
}

onMounted(() => {
  axios.get('https://callme.agileserve.org.cn:30941/show/')
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
  <div class="common-layout">
    <el-container class="container">
      <el-header class="header">
        <el-icon size="25" class="logo">
          <House/>
        </el-icon>
        <el-input
            v-model="input"
            class="w-50 m-2"
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
        <!--        <el-button class="login-text" v-else @click="user" text>{{ username }}</el-button>-->
      </el-header>
      <el-container>
        <el-header>
          <el-menu class="el-menu-demo" mode="horizontal" center>
            <el-menu-item index="1" @click="latest">按最新排序</el-menu-item>
            <el-menu-item index="2" @click="hottest" disabled>按最热排序</el-menu-item>
            <el-menu-item index="3" @click="createNewQuestion">创建新问题</el-menu-item>
          </el-menu>
        </el-header>
        <el-container>
          <!--          <el-aside width=20%>-->

          <!--          </el-aside>-->
          <el-main>
            <ul v-infinite-scroll="load" class="infinite-list" style="overflow: auto">
              <li v-for="card in cards" :key="card.id" class="infinite-list-item">
                <el-card class="box-card" shadow="hover" @click="card_details(card.id)">
                  <template #header>
                    <div class="card-header">
                      <span>#{{ card.id }}</span>
                      <el-button-group>
                        <!-- TODO: 如果用户已经收藏了这个问题，那么显示实心的星星，否则显示空心的星星 -->
                        热度{{ card.favoritesCount }}
                      </el-button-group>
                    </div>
                  </template>
                  <div class="card-body">
                    <div v-for="i in card.content.length">
                      {{ card.content[i] }}
                    </div>
                  </div>
                  <template #footer>
                    <div class="card-footer">
                      <span>{{ card.create_at }}</span>
                      <el-tag effect="light" type="info">
                        tag1
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

.footer {
  height: 10px;
}

.el-menu-demo {
  display: flex;
  justify-content: center;
  align-items: center;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 5px;
}

.card-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 5px;
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
  -webkit-line-clamp: 3;
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

.box-card {
  width: 100%;
}

.infinite-list {
  height: 600px;
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

.affix {
  text-align: center;
}
</style>