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
        <n-list-item v-for="report in reportedPosts" :key="report.id">
          <n-thing>
            <template #header>
              <n-tag class="top-tag" 
              :type="getStatusType(report.status)"
              size="small" round> 
              <b>
              {{ report.status.charAt(0).toUpperCase() + report.status.slice(1) }}
              </b>
            </n-tag>
              <b> Post {{  report.post_id }}</b>, reported by {{ report.reporter_id }}
            </template>
            <template #description> Reason: {{ report.reason }} 
              <br> created at: {{  formatDate(report.created_at.toLocaleString()) }} 
              <br> {{ lastStatus(report.status) }}: {{ formatDate(report.updated_at.toLocaleString()) }} 
            </template>
          </n-thing>
          <div class="button-group">
            <n-button color="#87CBAC" strong secondary round type="info" v-if="report.status !== 'resolved'" @click="handleViewReportedPost(report.post_id, report.id)"> <eye-outline class="icon"/> View</n-button>
            <n-button color="#7A7978" strong secondary round type="info" @click="handleDeleteReport(report.id)"> <trash-bin-outline class="icon"/> Close Report</n-button>
            <n-button color="#6593C1" strong secondary round type="info" v-if="report.status !== 'resolved'" @click="handleDeletePost(report.post_id, report.id)"> <trash-outline class="icon"/> Delete Post</n-button>
          </div>
        </n-list-item>
      </n-list>
      <n-empty v-else description="No reported posts found" />
    </n-card>
    <n-modal v-model:show="showReportedPost">
            <n-card title="Post Content" closable @close="showReportedPost = false">
              <post-item v-if="postDetail" :post="postDetail" :is-owner="false"
              :add-friend-action="false" :add-follow-action="false" :add-whisper-action="false"/>
            </n-card>
    </n-modal>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useI18n } from 'vue-i18n';
import MainNav from "@/components/main-nav.vue";
import { deleteReport, getReportList, getPost, deletePost, changeReportStatus } from '@/api/post';
import type { Item } from '@/types/Item';
import type { NetParams } from '@/types/NetParams';
import type { NetReq } from '@/types/NetReq';
import { NButton, NTag } from 'naive-ui';
import { TrashOutline, EyeOutline, TrashBinOutline } from "@vicons/ionicons5";
import PostItem from '@/components/post-item.vue';
import { format } from 'date-fns';


const { t } = useI18n();

const reportedPosts = ref<Item.ReportedPostProps[]>([]);
const showReportedPost = ref<boolean>(false);
const postDetail = ref<NetReq.PostGetPost>();

const formatDate = (dateString: string) => {
  return format(new Date(dateString), 'h:mm a, MMM d, yyyy');
};
const getStatusType = (status: string) => {
  switch (status) {
    case 'resolved': return 'success';
    case 'reviewed': return 'info';
    default: return 'error';
  }
};
const lastStatus = (status: string) => {
  switch (status) {
    case 'resolved': return 'deleted at';
    case 'reviewed': return 'last checked at';
    default: return 'updated at';
  }
};
const handleDeletePost = (id: number, reportId: number) => {
  const data: NetParams.PostDeletePost = {
    id: id,
  };
  deletePost(data).then(() => {
    window.$message.success('Post deleted successfully');
    changeReportStatus({
      report_id: reportId,
      status: 'resolved'
    }).then(() => {
      const reportIndex = reportedPosts.value.findIndex(report => report.id === reportId);
      if (reportIndex !== -1) {
        reportedPosts.value[reportIndex].status = 'resolved';
      }
    });;
  }).catch((err) => {
    window.$message.error('Failed to delete post:', err);
  });
};

const handleViewReportedPost = (id: number, reportId: number) => {
  const data: NetParams.PostGetPost = {
    id: id,
  };
  getPost(data).then((res) => {
    postDetail.value = res;
    showReportedPost.value = true;
    changeReportStatus({
      report_id: reportId,
      status: 'reviewed'
    }).then(()=>{
      const reportIndex = reportedPosts.value.findIndex(report => report.id === reportId);
      if (reportIndex !== -1) {
        reportedPosts.value[reportIndex].status = 'reviewed';
      }
    });
  }).catch((err) => {
    window.$message.error('Failed to fetch post detail:', err);
  });
};
const handleDeleteReport = (id: number) => {
  const data: NetParams.PostDeleteReport = {
    id: id,
  };
  deleteReport(data).then(()=>{
    window.$message.success('Report deleted successfully');
  }).catch((err)=>{
    window.$message.error('Failed to delete report:', err);
  });
};

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
.n-modal {
  width:50%;
}
.button-group {
  display: flex;
  gap: 10px;
  justify-content: flex-end;
}
.icon {
  width: 16px;
  height: 16px;
  margin-right: 5.5px;
}

</style>