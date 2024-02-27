<script setup>
import {onMounted, ref} from 'vue'
import {ArrowLeftBold, CircleCloseFilled, Search, Star, SuccessFilled} from "@element-plus/icons-vue";
import axios from "axios";
import {
  backendURL,
  username,
  authenticate,
  back,
  check,
  logout,
  myquestion,
  myfavorite,
} from "@/composables/commonLogic.js";

const input = ref('')
const isStar = ref(false)
const card_id = ref('')
let card_data = ref({})
let qa_id = ref([])
let qa_data = ref([])
let count = ref(0)

// Mounted

onMounted(() => {
  if (window.localStorage.getItem('userID') !== null) {
    check();
    username.value = window.localStorage.getItem('username');
  }
});

onMounted(() => {
  card_id.value = window.localStorage.getItem('cardID');
  axios.get(backendURL.value + '/get/card?cardID=' + card_id.value)
      .then((response) => {
        card_data = response.data.card;
        qa_id = card_data.questions;
        for (let i = 0; i < qa_id.length; i++) {
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
                  <span></span>
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
</style>
