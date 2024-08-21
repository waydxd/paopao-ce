<template>
  <div>
    <main-nav :title=communityName>
      <div class="nav-buttons">
        <button class="nav-button" @click="onClickCompose">
          <create-outline class="nav-icon" />
        </button>
      </div>
    </main-nav>
    <div class="community-posts">
      <compose class="compose-box" :community-id=communityId v-if="compose" @post-success="onPostSuccess" />
      <n-list v-if="communityPosts.length > 0">
        <n-list-item v-for="post in communityPosts" :key="communityId">
          <post-item :post="post" :is-owner="currentUserId !== null && post.user_id === currentUserId"
            :add-friend-action="false" :add-follow-action="false" @send-whisper="handleSendWhisper" />
        </n-list-item>
        <!--      <InfiniteLoading class="load-more" :slots="{ complete: '没有更多泡泡了', error: '加载出错' }" @infinite="nextPage">-->
        <!--        <template #spinner>-->
        <!--          <div class="load-more-wrap">-->
        <!--            <n-spin :size="14" />-->
        <!--            <span class="load-more-spinner">加载更多</span>-->
        <!--          </div>-->
        <!--        </template>-->
        <!--      </InfiniteLoading>-->
      </n-list>
      <n-empty v-else description="No posts found" />
      <n-space v-if="!noMore" justify="center">
        <InfiniteLoading class="load-more" :slots="{ complete: t('NoMorePosts'), error: t('LoadError') }"
          @infinite="nextPage()">
          <template #spinner>
            <div class="load-more-wrap">
              <n-spin :size="14" v-if="!noMore" />
              <span class="load-more-spinner">{{ noMore ? t('NoMorePosts') : t('LoadMore') }}</span>
            </div>
          </template>
        </InfiniteLoading>
      </n-space>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRoute } from 'vue-router';
import { useStore } from 'vuex';
import PostItem from '@/components/post-item.vue';
import { getCommunityByID, getCommunityPosts } from "@/api/user";
import InfiniteLoading from "v3-infinite-loading";
import "v3-infinite-loading/lib/style.css";
import { CreateOutline } from "@vicons/ionicons5";
import MainNav from "@/components/main-nav.vue";
import { Item } from "@/types/Item";
import Compose from "@/components/compose.vue";
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
const route = useRoute();
const store = useStore();
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

const communityName = ref('');
const communityPosts = ref<Item.PostProps[]>([]);
const currentUserId = computed(() => store.state.user?.id || null);
const communityId = computed(() => Number(route.params.id));

const page = ref(1);
const pageSize = ref(20);
const totalPosts = ref(0);
const noMore = ref(false);
const loading = ref(false);

const compose = ref<boolean>(false);
const color = ref<string>("#333");

const onClickCompose = () => {
  compose.value = !compose.value;
  color.value = (color.value === "#333") ? "#ffab00" : "#333";
}

const onPostSuccess = async (res: Item.PostProps) => {
  loading.value = false;
  console.log(res);
  await loadPosts();
}
onMounted(async () => {
  const communityId = Number(route.params.id);
  try {
    const communityData = await getCommunityByID({ community_id: communityId });
    communityName.value = communityData.community.name;
    await loadPosts();
  } catch (error) {
    console.error('Error fetching community data:', error);
  }
});

const loadPosts = async () => {
  if (loading.value) return;
  loading.value = true;
  try {
    const postsData = await getCommunityPosts({
      community_id: communityId.value,
      offset: (page.value - 1) * pageSize.value,
      limit: pageSize.value
    });

    const newPosts = postsData.posts.Tweets;
    totalPosts.value = postsData.Total;

    if (page.value === 1) {
      communityPosts.value = newPosts;
    } else {
      communityPosts.value.push(...newPosts);
    }

    noMore.value = communityPosts.value.length >= totalPosts.value;
    page.value++;
    // console.log("Rerun");
  } catch (error) {
    console.error('Error fetching posts:', error);
  } finally {
    noMore.value = true;
  }
};

const nextPage = () => {
  if (!noMore.value) {
    loadPosts();
  }
};
const handleSendWhisper = (user: Item.UserInfo) => {
  whisperReceiver.value = user;
  showWhisper.value = true;
};
</script>



<style scoped>
.compose-box {
  position: sticky;
}

.nav-icon {
  width: 1.3rem;
  vertical-align: -0.125em;
}

.nav-buttons {
  display: flex;
  margin-left: auto;
}

.nav-button {
  background: transparent;
  border: 0;
  padding: 5px;
  cursor: pointer;
}

.nav-button {
  color: v-bind(color);
}

.community-posts {
  border: 0.75px solid rgba(0, 0, 0, 0.05);
}

.load-more {
  margin: 20px 0;
}

.load-more-wrap {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 10px;
}

.n-list-item {
  padding: 0;
}
</style>