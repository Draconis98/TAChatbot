<script setup>
import {onMounted, ref} from 'vue';
import axios from "axios";
import {ArrowLeftBold, Back, Search} from "@element-plus/icons-vue";

// 使用ref来创建一个响应式的src属性
const iframeSrc = ref('');
const username = ref(null);

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

function back() {
  window.history.back();
}


onMounted(() => {
  // 获取当前页面的URL
  const userID = window.sessionStorage.getItem('userID');
  const cardID = window.sessionStorage.getItem('cardID');
  username.value = window.sessionStorage.getItem('username');

  iframeSrc.value = `http://127.0.0.1:8082?userID=${userID}&cardID=${cardID}`
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
              <el-icon><ArrowLeftBold/></el-icon>
              返回首页
            </el-menu-item>
          </el-menu>
        </el-header>
        <el-main>
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

.footer {
  height: 10px;
}

.logo {
  width: 25px;
  height: 25px;
  justify-items: center;
  align-items: center;
  margin-right: 10px;
}

.iframe {
  width: 100%;
  height: 80vh;
  align-items: center;
  align-content: center;
}

.affix {
  text-align: center;
  font-size: 14px;
}
</style>