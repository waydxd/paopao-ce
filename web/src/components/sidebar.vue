<template>
    <div class="sidebar-wrap">
        <div class="logo-wrap">
            <a href="https://vrpanda.org" target="_blank">
                <n-image class="logo-img" width="36" :src="LOGO" :preview-disabled="true" />
            </a>
        </div>
        <!-- menu -->
        <n-menu class="menu" :accordion="true" :icon-size="24" :options="menuOptions" :render-label="renderMenuLabel"
            :render-icon="renderMenuIcon" :value="selectedPath" @update:value="goRouter"
            :theme-overrides="menuThemeOverrides" />

        <div class="user-wrap" v-if="store.state.userInfo.id > 0">
            <n-avatar class="user-avatar" round :size="34" :src="store.state.userInfo.avatar" />

            <div class="user-info">
                <div class="nickname">
                    <span class="nickname-txt">
                        {{ store.state.userInfo.nickname }}
                    </span>
                    <n-button class="logout" quaternary circle size="tiny" @click="handleLogout">
                        <template #icon>
                            <n-icon>
                                <log-out-outline />
                            </n-icon>
                        </template>
                    </n-button>
                </div>
                <div class="username">@{{ store.state.userInfo.username }}</div>
            </div>

            <div class="user-mini-wrap">
                <n-button class="logout" quaternary circle @click="handleLogout">
                    <template #icon>
                        <n-icon :size="24">
                            <log-out-outline />
                        </n-icon>
                    </template>
                </n-button>
            </div>
        </div>
        <div class="user-wrap" v-else>
            <div v-if="!store.state.profile.allowUserRegister" class="login-only-wrap">
                <n-button strong secondary round type="primary" @click="triggerAuth('signin')">
                    {{ t('sidebar.login') }}
                </n-button>
            </div>
            <div v-if="store.state.profile.allowUserRegister" class="login-wrap">
                <n-button strong secondary round type="primary" @click="triggerAuth('signin')">
                    {{ t('sidebar.login') }}
                </n-button>
                <n-button strong secondary round type="info" @click="triggerAuth('signup')">
                    {{ t('sidebar.register') }}
                </n-button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { h, ref, watch, computed, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useStore } from 'vuex';
import { NIcon, NBadge, useMessage } from 'naive-ui';
import {
    HomeOutline,
    BookmarksOutline,
    MegaphoneOutline,
    ChatbubblesOutline,
    LeafOutline,
    PeopleOutline,
    WalletOutline,
    SettingsOutline,
    LogOutOutline,
    PeopleCircleOutline,
    BuildOutline
} from '@vicons/ionicons5';
import { Hash } from '@vicons/tabler';
import { getUnreadMsgCount } from '@/api/user';
import LOGO from '@/assets/img/logo.png';
import { VueCookies } from 'vue-cookies';
import { inject } from 'vue'
import { useI18n } from 'vue-i18n';

const { t } = useI18n();

const menuThemeOverrides = {
    itemTextColor: '#333333',           // Dark gray for normal text
    itemTextColorHover: '#ff8c00',      // Dark orange for hover text
    itemTextColorActive: '#ffffff',     // White for active text
    itemTextColorActiveHover: '#ffffff', // White for active hover text
    itemColorHover: '#fff0e0',          // Light orange for hover background
    itemColorActive: '#ffa500',         // Orange for active item background
    itemColorActiveHover: '#ff9500',    // Slightly darker orange for active hover
    itemColorActivePressed: '#ff8000',  // Even darker orange for active pressed
    itemColorActiveHoverPressed: '#ff7500', // Darkest orange for active hover and pressed
    itemColorActiveHoverFocus: '#ffb700',   // Brighter orange for active hover focus

    itemIconColor: '#333333',           // Dark gray for normal icon
    itemIconColorHover: '#ff8c00',      // Dark orange for hover icon
    itemIconColorActive: '#ffffff',     // White for active icon
    itemIconColorActiveHover: '#ffffff' // White for active hover icon
};

const store = useStore();
const route = useRoute();
const router = useRouter();
const hasUnreadMsg = ref(false);
const selectedPath = ref<any>(route.name || '');
const msgLoop = ref();

const enableAnnoucement = (import.meta.env.VITE_ENABLE_ANOUNCEMENT.toLowerCase() === 'true');

watch(route, () => {
    selectedPath.value = route.name;
});
watch(store.state, () => {
    hasUnreadMsg.value = store.state.unreadMsgCount > 0;
    if (store.state.userInfo.id > 0) {
        if (!msgLoop.value) {
            getUnreadMsgCount()
                .then((res) => {
                    hasUnreadMsg.value = res.count > 0;
                    store.commit("updateUnreadMsgCount", res.count)
                })
                .catch((err) => {
                    console.log(err);
                });

            msgLoop.value = setInterval(() => {
                getUnreadMsgCount()
                    .then((res) => {
                        hasUnreadMsg.value = res.count > 0;
                        store.commit("updateUnreadMsgCount", res.count)
                    })
                    .catch((err) => {
                        console.log(err);
                    });
            }, store.state.profile.defaultMsgLoopInterval);
        }
    } else {
        if (msgLoop.value) {
            clearInterval(msgLoop.value);
        }
    }
});
onMounted(() => {
    window.onresize = () => {
        store.commit('triggerCollapsedLeft', document.body.clientWidth <= 821);
        store.commit('triggerCollapsedRight', document.body.clientWidth <= 821);
    };
});
const menuOptions = computed(() => {
    const options = [
        {
            label: t('sidebar.home'),
            key: 'home',
            icon: () => h(HomeOutline),
            href: '/',
        },
        {
            label: t('sidebar.topic'),
            key: 'topic',
            icon: () => h(Hash),
            href: '/topic',
        },
    ];
    if (enableAnnoucement) {
        options.push({
            label: t('sidebar.anouncement'),
            key: 'announcement',
            icon: () => h(MegaphoneOutline),
            href: '/announcement',
        });
    }
    options.push({
        label: t('sidebar.profile'),
        key: 'profile',
        icon: () => h(LeafOutline),
        href: '/profile',
    });
    options.push({
        label: t('sidebar.message'),
        key: 'messages',
        icon: () => h(ChatbubblesOutline),
        href: '/messages',
    });
    options.push({
      label: t('sidebar.community'),
      key: 'community',
      icon: () => h(PeopleCircleOutline),
      href: '/community',
    });
    options.push({
        label: t('sidebar.archive'),
        key: 'collection',
        icon: () => h(BookmarksOutline),
        href: '/collection',
    });
    if (store.state.profile.useFriendship) {
        options.push({
            label: t('sidebar.friend'),
            key: 'contacts',
            icon: () => h(PeopleOutline),
            href: '/contacts',
        });
    }
    if (store.state.profile.enableWallet) {
        options.push({
            label: t('sidebar.wallet'),
            key: 'wallet',
            icon: () => h(WalletOutline),
            href: '/wallet',
        });
    }
    options.push({
        label: t('sidebar.setting'),
        key: 'setting',
        icon: () => h(SettingsOutline),
        href: '/setting',
    });
    if (store.state.userInfo.is_admin) {
        options.push({
            label: t('sidebar.admin'),
            key: 'admin',
            icon: () => h(BuildOutline),
            href: '/admin',
        });
    }
    return store.state.userInfo.id > 0
        ? options
        : [
            {
                label: t('sidebar.home'),
                key: 'home',
                icon: () => h(HomeOutline),
                href: '/',
            },
            {
                label: t('sidebar.topic'),
                key: 'topic',
                icon: () => h(Hash),
                href: '/topic',
            },
        ];
});

const renderMenuLabel = (option: AnyObject) => {
    if ('href' in option) {
        return h('div', {}, option.label);
    }
    return option.label;
};
const renderMenuIcon = (option: AnyObject) => {
    if (option.key === 'messages') {
        return h(
            NBadge,
            {
                dot: true,
                show: hasUnreadMsg.value,
                processing: true,
            },
            {
                default: () =>
                    h(
                        NIcon,
                        {
                            color:
                                option.key === selectedPath.value
                                    ? 'var(--n-item-icon-color-active)'
                                    : 'var(--n-item-icon-color)',
                        },
                        { default: option.icon }
                    ),
            }
        );
    }
    return h(NIcon, null, { default: option.icon });
};

const goRouter = (name: string, item: any = {}) => {
    selectedPath.value = name;
    router.push({
        name, query: {
            t: (new Date().getTime())
        }
    });
};

const triggerAuth = (key: string) => {
    store.commit('triggerAuth', true);
    store.commit('triggerAuthKey', key);
};
const $cookies = inject<VueCookies>('$cookies');

const handleLogout = () => {
    $cookies?.remove('userinfo'); // prevent user refresh page to login again
    store.commit('userLogout');
    store.commit('refresh')

    window.location.href = 'http://localhost:9527';

};
window.$store = store;
window.$message = useMessage();
</script>

<style lang="less">
:root {
    // 如果要变更中间栏的大小，修改此处即可
    --content-main: 620px;
    --n-item-color-active: #ff9f56;
    --n-item-text-color: #ff9f56;
    --n-item-text-color-hover: #ff9f56;
    --n-item-text-color-active: #ff9f56;
    --n-item-color-hover: #ff9f56;

}

.sidebar-wrap::-webkit-scrollbar {
    width: 0;
    /* 隐藏滚动条的宽度 */
    height: 0;
    /* 隐藏滚动条的高度 */
}

.sidebar-wrap {
    z-index: 99;
    width: 200px;
    height: 100vh;
    position: fixed;
    right: calc(50% + var(--content-main) / 2 + 10px);
    padding: 12px 0;
    box-sizing: border-box;
    max-height: calc(100vh);
    /* 调整高度 */
    overflow: auto;

    .n-menu .n-menu-item-content::before {
        border-radius: 21px;
    }


    .logo-wrap {
        display: flex;
        justify-content: flex-start;
        margin-bottom: 12px;

        .logo-img {
            margin-left: 24px;

            &:hover {
                cursor: pointer;
            }
        }
    }

    .user-wrap {
        display: flex;
        align-items: center;
        position: absolute;
        bottom: 12px;
        left: 12px;
        right: 12px;

        .user-mini-wrap {
            display: none;
        }

        .user-avatar {
            margin-right: 8px;
        }

        .user-info {
            display: flex;
            flex-direction: column;

            .nickname {
                font-size: 16px;
                font-weight: bold;
                line-height: 16px;
                height: 16px;
                margin-bottom: 2px;
                display: flex;
                align-items: center;

                .nickname-txt {
                    max-width: 90px;
                    text-overflow: ellipsis;
                    overflow: hidden;
                    white-space: nowrap;
                }

                .logout {
                    margin-left: 6px;
                }
            }

            .username {
                font-size: 14px;
                line-height: 16px;
                height: 16px;
                width: 120px;
                text-overflow: ellipsis;
                overflow: hidden;
                white-space: nowrap;
                opacity: 0.75;
            }
        }

        .login-only-wrap {
            display: flex;
            justify-content: center;
            width: 100%;

            button {
                margin: 0 4px;
                width: 80%
            }
        }

        .login-wrap {
            display: flex;
            justify-content: center;
            width: 100%;

            button {
                margin: 0 4px;
            }
        }
    }
}

.auth-card {
    .n-card-header {
        z-index: 999;
    }
}

@media screen and (max-width: 821px) {
    .sidebar-wrap {
        width: 200px;
        right: calc(100% - 200px);
    }

    .logo-wrap {
        .logo-img {
            margin-left: 12px !important;
        }
    }

    .user-wrap {

        .user-avatar,
        .user-info,
        .login-only-wrap,
        .login-wrap {
            margin-bottom: 32px;
        }

        //     .user-mini-wrap {
        //         display: block !important;
        //     }
    }
}
</style>