import {onMounted, ref} from "vue";
import axios from "axios";

// variables that are used in multiple components

export const backendURL = ref("http://127.0.0.1:8080");
export const username = ref(null)


// functions that are used in multiple components

export function authenticate() {
    axios.get(backendURL.value + '/auth')
        .then(response => {
            let url = new URL(response.data.url);
            window.location.href = url.href;
        })
        .catch(error => {
            alert("获取认证URL失败\n" + error)
        });
}

export function logout() {
    window.localStorage.removeItem('username');
    window.localStorage.removeItem('userID');
    window.localStorage.removeItem('cardID');
    window.location.href = '/';
}

export function back() {
    window.location.href = '/';
}

export function myquestion() {
    // TODO: 跳转到我的问题页面
    window.location.href = '/questions';
}

export function myfavorite() {
    // TODO: 跳转到我的收藏页面
}

export function login() {
    onMounted(() => {
        if (window.localStorage.getItem('userID') === null) {
            authenticate()
        }
        console.log("认证结束")
        axios.get(backendURL.value + '/get/username?userID=' + window.localStorage.getItem('userID'))
            .then(response => {
                window.localStorage.setItem('username', response.data.username);
                window.localStorage.setItem('role', response.data.role);
                username.value = window.localStorage.getItem('username');
            })
            .catch(error => {
                if (error.response.status === 401) {
                    alert('用户不存在，请重新登录\n' + error)
                } else {
                    alert("获取用户名失败\n" + error)
                }
                authenticate()
            });
    });
}

export function check() {
    axios.get(backendURL.value + '/get/username?userID=' + window.localStorage.getItem('userID'))
        .then((response) => {
            if (response.data.username !== window.localStorage.getItem('username')) {
                alert('用户信息不匹配，请重新登录');
                authenticate();
            }
        })
        .catch((error) => {
            console.log('获取用户信息失败\n' + error);
        });
}

export function createNewQuestion() {
    let userID = window.localStorage.getItem('userID');
    axios.get(backendURL.value + '/new/card?userID=' + userID)
        .then((response) => {
            console.log(response.data);
            window.localStorage.setItem('cardID', response.data.cardID);
            window.location.href = '/new/question';
        })
        .catch((error) => {
            alert('服务器出现错误，无法创建新问题，请重新尝试');
            console.log('创建新卡片失败：', error);
        });
}