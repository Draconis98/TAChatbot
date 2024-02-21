<script setup>
import axios from "axios";
import {onMounted} from "vue";

onMounted(() => {
      let code = window.location.href.split('code=')[1];
      if (code) {
        axios.get('http://10.30.19.40:8080/callback/gitlab?code=' + code)
            .then((response) => {
              console.log(response.data);
              let username = response.data.username;
              let email = response.data.email;
              axios.get('http://10.30.19.40:8080/check/user?username=' + username + '&email=' + email)
                  .then((response) => {
                    console.log(response.data);
                    if (response.data.error === null) {
                      window.sessionStorage.setItem('userID', response.data.userID);
                      // window.location.href = 'http://10.30.19.40:8081/?userID=' + response.data.userID;
                      window.location.href = '10.30.19.40:8081';
                    } else {
                      console.log(response.data.error);
                    }

                  })
                  .catch((error) => {
                    console.log('获取用户信息失败：', error);
                  });
            })
            .catch((error) => {
              console.log('获取认证URL失败：', error);
            });
      } else {
        console.log('没有code');
      }
    }
)
</script>

<template>
  <div class="loading">
    <div class="loading-content">
      <div class="loading-icon">
        <img src="@/assets/loading.gif" alt="loading" width="20" height="20"/>
        认证成功，等待后端数据同步...
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>
