<script setup>
import {computed, onMounted, ref, toRaw} from 'vue'
import {
  ArrowLeftBold,
  CircleCloseFilled,
  Edit,
  RemoveFilled,
  Search,
  Star,
  SuccessFilled
} from "@element-plus/icons-vue";
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
  createNewQuestion,
} from "@/composables/commonLogic.js";
import MarkdownIt from 'markdown-it';
import hljs from 'highlight.js';
import 'highlight.js/styles/github.css';

// variables

const input = ref('')
const isStar = ref(false)
const dialogFormVisible = ref(false)
const comment = ref('')
const questionID = ref('')
const question = ref('')
const answer = ref('')
const userID = ref('')

let card_data = ref({})
let qa_id = ref([])
let qa_data = ref([])
let role = ref(null)

// function

function toggleStar() {
  if (isStar.value === false) {
    isStar.value = true;
    console.log('收藏');
  } else {
    isStar.value = false;
    console.log('取消收藏');
  }
}

function dialog(item) {
  dialogFormVisible.value = !dialogFormVisible.value;
  questionID.value = item.id;
  question.value = item.question;
  answer.value = item.answer;
}

function cancel() {
  dialogFormVisible.value = false;
  comment.value = '';
}

function submit() {
  dialogFormVisible.value = false;
  console.log(comment.value);
  axios.get(backendURL.value + '/new/comment?questionID='
      + questionID.value + '&comment=' + encodeURIComponent(comment.value)
      + '&userID=' + window.localStorage.getItem('userID'))
      .then((response) => {
        console.log(response.data);
        comment.value = '';
        axios.get(backendURL.value + '/get/email?cardID=' +
            window.localStorage.getItem('cardID') +
            '&questionID=' + questionID.value)
            .then((response) => {
              console.log(response.data);
            }).catch((error) => {
          console.log('邮件发送失败\n', error);
        });
        window.location.reload();
      }).catch((error) => {
    console.log('添加评论失败\n', error);
  });
}

function parseMarkdown(text) {
  const md = new MarkdownIt({
    highlight: function (str, lang) {
      if (lang && hljs.getLanguage(lang)) {
        try {
          return '<pre class="hljs"><code>' +
              hljs.highlight(lang, str, true).value +
              '</code></pre>';
        } catch (__) {
        }
      }
      return '<pre class="hljs"><code>' + md.utils.escapeHtml(str) + '</code></pre>';
    }
  });
  return md.render(text);
}

// Mounted

onMounted(() => {
  if (window.localStorage.getItem('userID') !== null) {
    check();
    username.value = window.localStorage.getItem('username');
  }
});

onMounted(() => {
  axios.get(backendURL.value + '/get/card?cardID=' + window.localStorage.getItem('cardID'))
      .then((response) => {
        card_data = response.data.card;

        if (window.localStorage.getItem('userID') !== card_data.userID && sessionStorage.getItem("isPageReloaded") !== "true") {
          console.log('不是本人的卡片');
          sessionStorage.setItem("isPageReloaded", "true");
          axios.get(backendURL.value + '/get/click?cardID=' + window.localStorage.getItem('cardID'))
              .then((response) => {
                console.log('点击量+1');
              }).catch((error) => {
            console.log('点击量+1失败\n', error);
          });
        }

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

// 决定是否显示评论功能
onMounted(() => {
  axios.get(backendURL.value + '/get/userRole?userID=' + window.localStorage.getItem('userID'))
      .then((response) => {
        if (response.data.role === 'TA' || response.data.role === 'teacher') {
          window.localStorage.setItem('role', response.data.role);
          role.value = window.localStorage.getItem('role');
        }
      }).catch((error) => {
    console.log('获取用户角色失败\n', error);
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
            <el-menu-item index="1" @click="back">
              <el-icon>
                <ArrowLeftBold/>
              </el-icon>
              返回首页
            </el-menu-item>
            <el-menu-item index="2" @click="createNewQuestion">创建新问题</el-menu-item>
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
                  <el-timeline-item v-for="item in qa_data" :key="item" :timestamp="item.create_at" placement="top">
                    <el-card shadow="hover">
                      <div class="question">
                        <div v-html="parseMarkdown(item.question)"></div>
                      </div>
                      <el-divider>模型回答，仅供参考</el-divider>
                      <el-container>
                        <div v-html="parseMarkdown(item.answer)"></div>
                        <el-tooltip v-if="item.likes === 1" content="原提问者对这个答案表示满意" placement="bottom">
                          <el-icon class="like" color="green">
                            <SuccessFilled/>
                          </el-icon>
                        </el-tooltip>
                        <el-tooltip v-else-if="item.likes === 2" content="原提问者对这个答案表示不满"
                                    placement="bottom">
                          <el-icon class="like" color="blue">
                            <CircleCloseFilled/>
                          </el-icon>
                        </el-tooltip>
                        <el-tooltip v-else content="原提问者对这个答案没有表示" placement="bottom">
                          <el-icon class="like" color="gray">
                            <RemoveFilled/>
                          </el-icon>
                        </el-tooltip>
                        <el-footer height="0.5vh" v-if="item.comments.length === 0 && role === 'TA'"
                                   class="foot-button">
                          <el-button
                              plain @click="dialog(item)"
                              style="float: right"
                              :icon="Edit"
                          >
                            添加评论/答案修正
                          </el-button>
                        </el-footer>
                      </el-container>
                      <el-dialog v-model="dialogFormVisible"
                                 title="添加评论/答案修正"
                                 width="80vw"
                                 :modal="false"
                                 :close-on-click-modal="false"
                                 :close-on-press-escape="false"
                                 :show-close="false"
                      >
                        <el-form :model="form">
                          <el-form-item label="原问题">
                            <pre class="pre">{{ question }}</pre>
                          </el-form-item>
                          <el-form-item label="原答案">
                            <pre class="pre">{{ answer }}</pre>
                          </el-form-item>
                          <el-form-item>
                            <el-input
                                v-model="comment"
                                autosize
                                type="textarea"
                                placeholder="请输入评论/答案修正"
                            />
                          </el-form-item>
                        </el-form>
                        <template #footer>
                          <div class="dialog-footer">
                            <el-button v-model="comment" @click="cancel">取消</el-button>
                            <el-button v-model="comment" type="primary" @click="submit" v-if="comment === ''" disabled>
                              提交
                            </el-button>
                            <el-button v-model="comment" type="primary" @click="submit" v-else>提交</el-button>
                          </div>
                        </template>
                      </el-dialog>
                      <el-divider v-if="item.comments.length !== 0">评论/纠正</el-divider>
                      <el-container v-if="item.comments.length !== 0">
                        <pre class="pre">{{ item.comments_content[item.comments.length - 1] }}</pre>
                        <el-footer height="0.5vh" class="foot-button">
                          <el-button
                              v-if="role === 'TA' || 'teacher'"
                              plain @click="dialog(item)"
                              style="float: right"
                              :icon="Edit"
                          >
                            修改
                          </el-button>
                        </el-footer>
                      </el-container>
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
.pre {
  display: flex;
  white-space: pre-wrap;
  white-space: -moz-pre-wrap;
  white-space: -o-pre-wrap;
  word-wrap: break-word;
  font-size: 14px;
}

.question {
  font-weight: bold;
}

b {
  font-weight: bold;
}

.like {
  display: flex;
  align-items: center;
  justify-content: center;
  align-content: center;
  justify-items: center;
}

.foot-button {
  display: flex;
  align-content: center;
  justify-content: center;
  align-items: center;
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
