<template>
  <div>
    <p>用户名: {{ userInfo.username }}</p>
    <p>邮箱: {{ userInfo.email }}</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      userInfo: {}
    }
  },
  mounted() {
    // 假设您的Gin后端已经设置好并返回了用户信息
    // 这里是获取信息的逻辑，可能需要适当修改以适应实际情况
    this.fetchUserInfo();
  },
  methods: {
    fetchUserInfo() {
      // 使用fetch API或Vue的HTTP库(如axios)获取信息
      // 以下代码假设您的后端返回的信息位于response.data中
      fetch("http://localhost:8080/callback/github")
          .then(response => response.json())
          .then(data => {
            this.userInfo = data;
          })
          .catch(error => console.error("Error:", error));
    }
  }
}
</script>
