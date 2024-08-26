<!-- src/views/Admin.vue -->
<template>
  <div>
    <main-nav :title="t('sidebar.admin')" theme />
    <!-- <n-card class="Reported">
      <n-list>
        <n-list-item>
          <n-icon size="20" color="#f00" class="icon" />
          <span class="text">Reported Posts</span>
        </n-list-item>
        <n-list-item>
          <n-icon size="20" color="#f00" class="icon" />
          <span class="text">Reported Comments</span>
        </n-list-item>
        <n-list-item>
          <n-icon size="20" color="#f00" class="icon" />
          <span class="text">Reported Users</span>
        </n-list-item>
      </n-list>
    </n-card> -->
    <n-card title="Reported Posts" class="reported-posts">
      <n-list v-if="reportedPosts.length">
        <n-list-item v-for="post in reportedPosts" :key="post.id">
          <n-thing>
            <template #header><b>Post {{ post.post_id }}</b>, reported by {{ post.reporter_id }}</template>
            <template #description> Reason: {{ post.reason }} <br> created at: {{ post.created_at }} <br> updated at: {{ post.updated_at }} </template>
          </n-thing>
        </n-list-item>
      </n-list>
      <n-empty v-else description="No reported posts found" />
    </n-card>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useI18n } from 'vue-i18n';
import MainNav from "@/components/main-nav.vue";
import { deleteReport, getReportList } from '@/api/post';
import type { Item } from '@/types/Item';
const reportedPosts = ref<Item.ReportedPostProps[]>([]);

const { t } = useI18n();

onMounted(async () => {
  try {
    const res = await getReportList();
    reportedPosts.value = res.reports || [];
  } catch (error) {
    console.error('Failed to fetch reported posts:', error);
  }
});
</script>

<style scoped>

</style>