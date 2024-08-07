<template>
    <n-modal v-model:show="store.state.authModalShow" class="auth-card" preset="card" size="small"
        :mask-closable="false" :bordered="false" :style="{
            width: '360px',
        }">
        <div class="auth-wrap">
            <n-card :bordered="false">
                <a :href="loginUrl">Please go to vrpanda website to login</a>
            </n-card>
        </div>
    </n-modal>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useStore } from 'vuex';
import { userInfo, cookieLogin } from '@/api/auth';
import { VueCookies } from 'vue-cookies';
import { inject } from 'vue'

const store = useStore();
const loading = ref(false); 
const loginUrl = import.meta.env.VITE_APP_LOGIN_URL;


const handleLogin = async (cookie: string) => {
    loading.value = true;
    await cookieLogin({
        Token: cookie,
    }).then((res) => {
        const token = res?.token || '';
        // 写入用户信息
        console.log("I GOT HERE");

        localStorage.setItem('PAOPAO_TOKEN', token);
        return userInfo(token);
    }).then((res) => {
        window.$message.success('登录成功');
        loading.value = false;
        store.commit('updateUserinfo', res);
        console.log(res);

        store.commit('triggerAuth', false);
        store.commit('refresh');
    }).catch((err) => {
            console.log(err);
            loading.value = false;
        });
};


const $cookies = inject<VueCookies>('$cookies');

const autoLoginOrRegister = async () => {
    const cookie = $cookies?.get('userinfo');
    if (!cookie) {
        throw new Error('No cookie found');
    }
    await handleLogin(cookie)
        .catch((err) => {
            console.log("I GOT ERROR HERE", err);
            // store.commit('userLogout'); 
        });

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
                console.log(err);
                store.commit('userLogout');
            });
    } else {
        autoLoginOrRegister()
            .catch((err) => {
                console.log(err);
            });
    }
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