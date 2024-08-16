<template>
  <main-nav :title=communityName >
  </main-nav>
  <div class="community-posts">
    <n-list v-if="communityPosts.length > 0">
      <n-list-item v-for="post in communityPosts" :key="post.id">
        <post-item
            :post="post"
            :is-owner="currentUserId !== null && post.user_id === currentUserId"
            :add-friend-action="true"
            :add-follow-action="true"
            @send-whisper="handleSendWhisper"
            @handle-follow-action="handleFollowAction"
            @handle-friend-action="handleFriendAction"
        />
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
      <InfiniteLoading class="load-more" :slots="{ complete: '没有更多泡泡了', error: '加载出错' }" @infinite="nextPage()">
      <template #spinner>
        <div class="load-more-wrap">
          <n-spin :size="14" v-if="!noMore" />
          <span class="load-more-spinner">{{ noMore ? '没有更多泡泡了' : '加载更多' }}</span>
        </div>
      </template>
    </InfiniteLoading>
    </n-space>
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
import {PersonAddOutline, SearchOutline} from "@vicons/ionicons5";
import MainNav from "@/components/main-nav.vue";

const route = useRoute();
const store = useStore();

const communityName = ref('');
const communityPosts = ref<Item.PostProps[]>([]);
const currentUserId = computed(() => store.state.user?.id || null);
const communityId = computed(() => Number(route.params.id));

const page = ref(1);
const pageSize = ref(20);
const totalPosts = ref(0);
const noMore = ref(false);
const loading = ref(false);

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
  } catch (error) {
    console.error('Error fetching posts:', error);
  } finally {
    loading.value = false;
    noMore.value = true;
  }
};

const nextPage = () => {
  if (!noMore.value) {
    loadPosts();
  }
};
const handleSendWhisper = (user: Item.UserInfo) => {
  // Implement whisper functionality
};

const handleFollowAction = (post: Item.PostProps) => {
  // Implement follow/unfollow functionality
};

const handleFriendAction = (post: Item.PostProps) => {
  // Implement friend action functionality
};
</script>



<style scoped>
.community-posts {
  border: 1px solid rgba(0, 0, 0, 0.1);
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
.n-list-item{
  padding: 0;
}
</style>