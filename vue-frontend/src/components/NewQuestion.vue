<script setup>
import {onMounted, ref} from 'vue';
import axios from "axios";
import {ArrowLeftBold, Search} from "@element-plus/icons-vue";
import {
  backendURL,
  username,
  authenticate,
  back,
  logout,
  myquestion,
  myfavorite,
} from "@/composables/commonLogic.js";

// variables

const iframeSrc = ref('');
const input = ref('');
const display = ref(true);

// function

function isdisplay(display) {
  let cardID = window.localStorage.getItem('cardID');
  axios.get(backendURL.value + '/get/isdisplay?cardID=' + cardID + '&display=' + display)
      .then((response) => {
        console.log(response.data);
      })
      .catch((error) => {
        console.log('获取认证URL失败：', error);
      });
  console.log(display);
}

// Mounted

onMounted(() => {
  // 获取当前页面的URL
  const userID = window.localStorage.getItem('userID');
  const cardID = window.localStorage.getItem('cardID');
  username.value = window.localStorage.getItem('username');

  iframeSrc.value = `https://callme.agileserve.org.cn:30942/?userID=${userID}&cardID=${cardID}`;
});
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
        <el-main>
          <div>
            <el-tooltip content="选择是否在首页显示，默认展示" placement="right">
              <el-switch
                  v-model="display"
                  active-text="显示"
                  inactive-text="隐藏"
                  @change="isdisplay(display)"
              />
            </el-tooltip>
          </div>
          <div>
            <iframe
                :src="iframeSrc"
                frameborder="0"
                class="iframe"
            />
          </div>
        </el-main>
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
.el-dropdown-link {
  cursor: pointer;
  color: #329eff;
  text-decoration: underline;
  display: flex;
  align-items: center;
  padding: 0 10px;
}

.iframe {
  width: 100%;
  height: 76vh;
  align-items: center;
  align-content: center;
}
</style>