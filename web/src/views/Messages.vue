<template>
    <div>
            <main-nav :title="t('sidebar.message')" />

            <n-list class="main-content-wrap messages-wrap" bordered>
             <!-- 私信组件 -->
            <whisper :show="showWhisper" :user="whisperReceiver" @success="whisperSuccess" />
            <n-space justify="space-between">
                        <div class="title title-action">
                            <n-button text size="small" :focusable="false" @click="handleUnreadMessage">
                                <template #icon>
                                    <n-icon>
                                        <UnreadIcon />
                                    </n-icon>
                                </template>
                                {{ store.state.unreadMsgCount }} {{ t('message.notRead') }}
                            </n-button>
                            <n-divider vertical />
                            <n-button text size="small" :focusable="false" @click="handleReadAll">{{ t('message.readAll') }}</n-button>
                        </div>
                        <div class="title title-filter">
                            <n-dropdown 
                            placement="bottom-end"
                            trigger="click"
                            size="small"
                            :options="options"
                            @select="handleAction">
                                <n-button text>
                                    <template #icon>
                                        <n-icon>
                                            <OptionsIcon />
                                        </n-icon>
                                    </template>
                                    {{ messageStyle }}
                                </n-button>
                            </n-dropdown>
                        </div>
            </n-space>
            <div v-if="loading && list.length === 0" class="skeleton-wrap">
                <message-skeleton :num="pageSize" />
            </div>
            <div v-else>
                <div class="empty-wrap" v-if="list.length === 0">
                    <n-empty size="large" :description="t('message.noMoreMessages')" />
                </div>
                <div v-else>
                    <n-list-item v-for="m in list" :key="m.id">
                        <message-item :message="m" @send-whisper="onSendWhisper" @reload="reloadMessages" />
                     </n-list-item>
                </div>
            </div>
        </n-list>
        <n-space v-if="totalPage > 0" justify="center">
            <InfiniteLoading class="load-more" :slots="{ complete: t('NoMoreData'), error: t('error.loading') }" @infinite="nextPage">
                <template #spinner>
                    <div class="load-more-wrap">
                        <n-spin :size="14" v-if="!noMore" />
                        <span class="load-more-spinner">{{ noMore ? t('NoMoreData') : t('LoadMore') }}</span>
                    </div>
                </template>
            </InfiniteLoading>
        </n-space>
    </div>
</template>

<script setup lang="ts">
import { h, ref, onMounted, computed } from 'vue';
import type { Component } from 'vue'
import { NIcon, DropdownOption } from 'naive-ui'
import { useStore } from 'vuex';
import { useRoute } from 'vue-router';
import InfiniteLoading from "v3-infinite-loading";
import { getMessages, readAllMessage } from '@/api/user';
import { 
    LayersOutline as AllIcon,
    AtOutline as SystemIcon,
    PaperPlaneOutline as WhisperIcon,
    PersonAddOutline as RequestingIcon,
    ChatbubbleEllipsesOutline as UnreadIcon,
    OptionsOutline as OptionsIcon, 
 } from '@vicons/ionicons5'
import { Item } from '@/types/Item';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
const store = useStore();
const route = useRoute();
const loading = ref(false);
const noMore = ref(false);
const page = ref(+(route.query.p as string) || 1);
const pageSize = ref(20);
const totalPage = ref(0);
const list = ref<Item.MessageProps[]>([]);
const messageStyle = ref<string>(t('message.all'))
const messageStyleVal = ref<'all' | 'system' | 'whisper' | 'requesting' | 'unread'>('all')
const showWhisper = ref(false);
const whisperReceiver = ref<Item.UserInfo>({
    id: 0,
    avatar: '',
    username: '',
    nickname: '',
    is_admin: false,
    is_friend: true,
    is_following: false,
    created_on: 0,
    follows: 0,
    followings: 0,
    status: 1,
});

const reset = () => {
    noMore.value = false;
    page.value = 1;
    totalPage.value = 0;
    list.value = [];
}

const renderIcon = (icon: Component) => {
  return () => {
    return h(NIcon, null, {
      default: () => h(icon)
    })
  }
}

const options = computed(() => {
    let opts: DropdownOption[];
    switch (messageStyle.value) {
    case t('message.all'):
        opts = [
            {
                label: t('message.system'),
                key: 'system',
                icon: renderIcon(SystemIcon)
            },
            {
                label: t('message.whisper'),
                key: 'whisper',
                icon: renderIcon(WhisperIcon)
            },
            {
                label: t('message.requesting'),
                key: 'requesting',
                icon: renderIcon(RequestingIcon)
            },
            {
                label: t('message.unread'),
                key: 'unread',
                icon: renderIcon(UnreadIcon)
            }
        ]
        break;
    case t('message.system'):
        opts = [
            {
                label: t('message.all'),
                key: 'all',
                icon: renderIcon(AllIcon)
            },
            {
                label: t('message.whisper'),
                key: 'whisper',
                icon: renderIcon(WhisperIcon)
            },
            {
                label: t('message.requesting'),
                key: 'requesting',
                icon: renderIcon(RequestingIcon)
            },
            {
                label: t('message.unread'),
                key: 'unread',
                icon: renderIcon(UnreadIcon)
            }
        ]
        break;
    case t('message.whisper'):
        opts = [
            {
                label: t('message.all'),
                key: 'all',
                icon: renderIcon(AllIcon)
            },
            {
                label: t('message.system'),
                key: 'system',
                icon: renderIcon(SystemIcon)
            },
            {
                label: t('message.requesting'),
                key: 'requesting',
                icon: renderIcon(RequestingIcon)
            },
            {
                label: t('message.unread'),
                key: 'unread',
                icon: renderIcon(UnreadIcon)
            }
        ]
        break;
    case t('message.requesting'):
            opts = [
            {
                label: t('message.all'),
                key: 'all',
                icon: renderIcon(AllIcon)
            },
            {
                label: t('message.system'),
                key: 'system',
                icon: renderIcon(SystemIcon)
            },
            {
                label: t('message.whisper'),
                key: 'whisper',
                icon: renderIcon(WhisperIcon)
            },
            {
                label: t('message.unread'),
                key: 'unread',
                icon: renderIcon(UnreadIcon)
            }
        ]
        break;
    case t('message.unread'):
            opts = [
            {
                label: t('message.all'),
                key: 'all',
                icon: renderIcon(AllIcon)
            },
            {
                label: t('message.system'),
                key: 'system',
                icon: renderIcon(SystemIcon)
            },
            {
                label: t('message.whisper'),
                key: 'whisper',
                icon: renderIcon(WhisperIcon)
            },
            {
                label: t('message.requesting'),
                key: 'requesting',
                icon: renderIcon(RequestingIcon)
            }
        ]
        break;
    default:
        opts = [];
        break;
    }
    return opts;
});

const handleAction = (
    item: 'all' | 'system' | 'whisper' | 'requesting' | 'unread'
) => {
    switch (item) {
    case 'all':
        messageStyle.value = t('message.all');
        break;
    case 'system':
        messageStyle.value = t('message.system');
        break;
    case 'whisper':
        messageStyle.value = t('message.whisper');
        break;
    case 'requesting':
        messageStyle.value = t('message.requesting');
        break;
    case 'unread':
        messageStyle.value = t('message.unread');
        break;
    }
    messageStyleVal.value = item
    reset();
    loadMessages();
};

const handleUnreadMessage = () => {
    handleAction('unread')
}

const handleReadAll = () => {
    if (store.state.unreadMsgCount > 0 && list.value.length > 0) {
        readAllMessage().then((_res) => {
            if (messageStyleVal.value != "unread") {
                for (let idx in list.value) {
                    list.value[idx].is_read = 1;
                }
            } else {
                list.value = [];
            }
            store.commit("updateUnreadMsgCount", 0)
        })
        .catch((err) => {
            console.log(err);
        });
    }
}

const onSendWhisper =  (user: Item.UserInfo) => {
    whisperReceiver.value = user;
    showWhisper.value = true;
};

const whisperSuccess = () => {
    showWhisper.value = false;
};

const reloadMessages = () => {
    reset();
    loadMessages();
};

const loadMessages = () => {
    loading.value = true;
    getMessages({
        style: messageStyleVal.value,
        page: page.value,
        page_size: pageSize.value,
    }).then((res) => {
        loading.value = false;
        if (res.list.length === 0) {
            noMore.value = true
        }
        if (page.value > 1) {
            list.value = list.value.concat(res.list);
        } else {
            list.value = res.list;
            window.scrollTo(0, 0);
        }
        totalPage.value = Math.ceil(res.pager.total_rows / pageSize.value);
    })
    .catch((_err) => {
        loading.value = false;
        if (page.value > 1) {
            page.value--
        }
    });
};
const nextPage = () => {
    if (page.value < totalPage.value || totalPage.value == 0) {
        noMore.value = false;
        page.value++;
        loadMessages();
    } else {
        noMore.value = true;
    }
};
onMounted(() => {
    loadMessages();
});
</script>

<style lang="less" scoped>
.load-more {
    margin: 20px;

    .load-more-wrap {
        display: flex;
        flex-direction: row;
        justify-content: center;
        align-items: center;
        gap: 14px;

        .load-more-spinner {
            font-size: 14px;
            opacity: 0.65;
        }
    }
}
.title {
    padding-top: 4px;
    opacity: 0.9;
}
.title-action {
    display: flex;
    align-items: center;
    margin-left: 20px;
}
.title-filter {
    margin-right: 20px;
}
.dark {
    .empty-wrap {
        background-color: rgba(16, 16, 20, 0.75);
    }
    .messages-wrap, .pagination-wrap {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>