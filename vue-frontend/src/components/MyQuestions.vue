<script setup>
import {onMounted, ref} from 'vue'
import {ArrowLeftBold, Search} from "@element-plus/icons-vue";
import axios from "axios";
import {
  backendURL,
  username,
  authenticate,
  back,
  check,
  createNewQuestion,
  logout,
  myquestion,
  myfavorite,
  login,
} from "@/composables/commonLogic.js";

// variables

const input = ref('')
const isStar = ref(false)
const count = ref(0)
const checked = ref(false)
let cards = ref([])

const load = () => {
  count.value += 3;
}

// function

function getCardsSortedBy(method) {
  axios.get(backendURL.value + '/show/' + method)
      .then((response) => {
        cards.value = [];
        for (let i = 0; i < response.data.card_list.length; i++) {
          if (response.data.card_list[i].content.length > 0 &&
              response.data.card_list[i].userID === window.localStorage.getItem('userID')) {
            cards.value.push(response.data.card_list[i]);
          }
        }
      })
      .catch((error) => {
        alert('获取卡片失败\n' + error);
      });
}

function card_details(cardID) {
  window.localStorage.setItem('cardID', cardID);
  window.location.href = '/card/details';
}

// Mounted

login();

onMounted(() => {
  if (window.localStorage.getItem('userID') != null) {
    check();
    getCardsSortedBy("");
  }
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
              <el-dropdown-item @click="myquestion">我的问题</el-dropdown-item>
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
            <el-menu-item index="2" @click="getCardsSortedBy('latest')">按最新排序</el-menu-item>
            <el-menu-item index="3" @click="getCardsSortedBy('hottest')">按最热排序</el-menu-item>
            <el-menu-item index="4" @click="createNewQuestion">创建新问题</el-menu-item>
          </el-menu>
        </el-header>
        <el-container>
          <el-main>
            <ul v-infinite-scroll="load" class="infinite-list" style="overflow: auto">
              <li v-for="card in cards" :key="card.id" class="infinite-list-item">
                <!-- 鼠标放在上面显示手指 -->
                <el-card class="box-card" shadow="hover" @click="card_details(card.id)">
                  <template #header>
                    <div class="card-header">
                      <span>{{ card.create_at }}</span>
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
                      <!--                      <span>{{ card.create_at }}</span>-->
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
          本平台仅供学习使用，请勿做其他用途；生成式回答，内容仅供参考。模型由Baichuan2-7B经过指令微调得到，使用Apache 2.0协议。
        </div>
      </el-footer>
    </el-container>
  </div>
</template>


<style scoped>
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
  line-clamp: 3;
}

.card-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 5px;
}

.box-card {
  width: 100%;
  cursor: pointer;
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
  margin: 10px 5px;
  //color: var(--el-color-primary);
}

.infinite-list .infinite-list-item + .list-item {
  margin-top: 10px;
}
</style>
