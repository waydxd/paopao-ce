<template>
    <n-modal
        v-model:show="store.state.authModalShow"
        class="auth-card"
        preset="card"
        size="small"
        :mask-closable="false"
        :bordered="false"
        :style="{
            width: '360px',
        }"
    >
        <div class="auth-wrap">
            <n-card :bordered="false">
            <a href="http://localhost:9527">Please go to vrpanda website to login</a>
            </n-card>
        </div>
    </n-modal>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useStore } from 'vuex';
import { userLogin, userRegister, userInfo } from '@/api/auth';
import { VueCookies } from 'vue-cookies';
import { inject } from 'vue'
import * as jose from 'jose'

const store = useStore();
const loading = ref(false); // Potentially redundant
const password = 'password';



const handleLogin = async (username: string) => {
    try {
        loading.value = true;
        const res = await userLogin({
            username: username,
            password: password,
        });
        const token = res?.token || '';
        // 写入用户信息
        localStorage.setItem('PAOPAO_TOKEN', token);
        const userInfoResponse = await userInfo(token);
        window.$message.success('登录成功');
        loading.value = false;
        store.commit('updateUserinfo', userInfoResponse);
        store.commit('triggerAuth', false);
        store.commit('refresh');

    } catch (err) {
        console.log(err);
        loading.value = false;
        throw err;
    }
};

const handleRegister = async (username: string) => {
    try {
        loading.value = true;
        await userRegister({
            username: username,
            password: password,
        });
        const res = await userLogin({
            username: username,
            password: password,
        });
        const token = res?.token || '';
        // 写入用户信息
        localStorage.setItem('PAOPAO_TOKEN', token);

        const userInfoResponse = await userInfo(token);
        window.$message.success('注册成功');
        loading.value = false;
        store.commit('updateUserinfo', userInfoResponse);
        store.commit('triggerAuth', false);
    } catch (err) {
        loading.value = false;
        throw err;
    }
};
const $cookies = inject<VueCookies>('$cookies'); 

const autoLoginOrRegister = async () => {
    const cookie = $cookies?.get('userinfo');
    if (!cookie) {
        return;
    }
    const jwk = await jose.importJWK({
          "kty": "oct",
          "use": "sig",
          "alg": "HS256",
          "kid": "595c072f-ad49-4a8b-8df9-85154963bf51",
          "k": "LhVU5U8jT-H7QIzp74eEtiACa_fF0Op_h7F2ZmmkmjVUQmJjcmz1fyiXCVmJLEs4AOY8nxxdLDbWiuZFDV1kig"
        });
    const token = localStorage.getItem('PAOPAO_TOKEN') || '';
    if (token) {
        // token found, just log it in
        const userInfoResponse = await userInfo(token);
        store.commit('triggerAuth', true);
        // return;
    }
    // Decode the JWT to get the username
    const { payload, protectedHeader } = await jose.compactVerify(cookie, jwk);
    
    // Try to get user info
    const username = JSON.parse(new TextDecoder().decode(payload)).data.username;
    try{    
    // If we get here, the user exists and is logged in
        await handleLogin(username).catch((err) => {
            console.log("user not found, try registering");
            handleRegister(username);
        });

    } catch (err) { // Potentially redundant 
        console.log(err);
      // Store the new token
        //   setCookie('PAOPAO_TOKEN', loginResponse.token, 7);
        //   const userInfoResponse = await userInfo(loginResponse.token);
        //   store.commit('updateUserinfo', userInfoResponse);
        //   store.commit('triggerAuth', false);
      store.commit('triggerAuth', true);
    }
};

onMounted(() => {
    const token = localStorage.getItem('PAOPAO_TOKEN') || '';
    if (token) {
        userInfo(token)
            .then((res) => {
                store.commit('updateUserinfo', res);
                store.commit('triggerAuth', false);
            })
            .catch((err) => {
                store.commit('userLogout');
            });
    } else {
        store.commit('userLogout');
    }
    autoLoginOrRegister();
});
</script>

<style lang="less" scoped>
.auth-wrap {
    margin-top: -30px;
}
.dark {
    .auth-wrap {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>