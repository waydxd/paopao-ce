<template xmlns:max-width="http://www.w3.org/1999/xhtml">
  <div>
    <main-nav :title="t('sidebar.community')">
      <div class="nav-buttons">
        <!-- <button class="nav-button">
          <search-outline class="nav-icon" />
        </button> -->
        <button @click="toggleDrawer" v-if="userInfo.is_admin" class="nav-button">
          <person-add-outline class="nav-icon" />
        </button>
      </div>
    </main-nav>
    <n-modal
    :mask-closable="false"
                    preset="dialog"
                    :title="t('community.create_community')"
                    v-model:show="isDrawerOpen" closable>
      <div class="modal-content">
      <n-form :model="formModel" :rules="rules" ref="formRef">
        <n-form-item :label="t('community.name')" path="name">
          <n-input v-model:value="formModel.name" />
        </n-form-item>
        <n-form-item :label="t('community.description')" path="description">
          <n-input v-model:value="formModel.description" type="textarea" />
        </n-form-item>
        <n-form-item :label="t('community.avatar_url')" path="avatar_url">
          <n-input v-model:value="formModel.avatar_url" />
        </n-form-item>
        <n-form-item :label="t('community.banner_url')" path="banner_url">
          <n-input v-model:value="formModel.banner_url" />
        </n-form-item>
        <n-button type="primary" @click="submitForm">{{ t('community.submit') }}</n-button>
      </n-form>
    </div>
    </n-modal>
    <div class="forum">
      <n-list class="main-content-wrap" bordered>
        <ul class="grid">
          <li v-for="forum in forums" :key="forum.id">
            <router-link :to="{ name: 'CommunityPosts', params: { id: forum.id } }" class="card-link">
              <div class="card">
                <img alt="" :src="forum.avatar_url" class="community_icon"/>
                <div class="text">
                  <h3 class="card__title">{{ forum.name }}
                  </h3>
                  <p class="card__content">{{ forum.description }}</p>
                  <div class="card__date">
                    {{ forum.created_at }}
                  </div>
                </div>
                <div class="card__arrow">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" height="15" width="15">
                    <path fill="#fff"
                      d="M13.4697 17.9697C13.1768 18.2626 13.1768 18.7374 13.4697 19.0303C13.7626 19.3232 14.2374 19.3232 14.5303 19.0303L20.3232 13.2374C21.0066 12.554 21.0066 11.446 20.3232 10.7626L14.5303 4.96967C14.2374 4.67678 13.7626 4.67678 13.4697 4.96967C13.1768 5.26256 13.1768 5.73744 13.4697 6.03033L18.6893 11.25H4C3.58579 11.25 3.25 11.5858 3.25 12C3.25 12.4142 3.58579 12.75 4 12.75H18.6893L13.4697 17.9697Z">
                    </path>
                  </svg>
                </div>
              </div>
            </router-link>
          </li>
        </ul>
      </n-list>
    </div>
  </div>
</template>


<script lang="ts" setup>
import { PersonAddOutline, SearchOutline } from '@vicons/ionicons5'
import { ref, onMounted, reactive } from 'vue';
import { getCommunityList } from "@/api/user";
import { NetReq } from "@/types/NetReq";
import MainNav from "@/components/main-nav.vue";
import { NForm, NFormItem, NInput, NButton } from 'naive-ui';
import { useStore } from 'vuex';
import { useI18n } from 'vue-i18n';
import { createCommunity } from "@/api/user";
import type { FormInst } from 'naive-ui';
import type { Item } from "@/types/Item";

const {t} = useI18n();
const userStore = useStore();

const userInfo = userStore.state.userInfo;
const isDrawerOpen = ref<boolean>(false);
const forums = reactive<Item.Community[]>([]);

const formRef = ref<FormInst | null>(null);
const formModel = reactive<NetReq.createCommunityReq>({
  name: '',
  description: '',
  avatar_url: '',
  banner_url: '',
});

const rules = {
  name: {
    required: true,
    message: t('community.name_required'),
    trigger: 'blur',
  },
  description: {
    required: false,
    message: t('community.description_required'),
    trigger: 'blur',
  },
  avatar_url: {
    required: false,
    message: t('community.avatar_url_required'),
    trigger: 'blur',
  },
  banner_url: {
    required: false,
    message: t('community.banner_url_required'),
    trigger: 'blur',
  },
};

function submitForm() {
  formRef.value?.validate(async (errors: any) => {
    if (!errors) {
      try {
        const response = await createCommunity(formModel);
        isDrawerOpen.value = false;
        forums.push(response.community);
      } catch (error) {
        window.$message.error('Error creating community: '+ error);
      }
    } else {
      window.$message.error('Form validation failed');
    }
  });
}

onMounted(async () => {
  try {
    const response = await getCommunityList({
      offset: 0,
      limit: 20,
    });
    forums.splice(0, forums.length, ...response.communities);
  } catch (error) {
    console.error('Error fetching community list:', error);
  }
})

function toggleDrawer() {
  isDrawerOpen.value = !isDrawerOpen.value;
}
</script>

<style scoped>
.nav-icon {
  width: 1.3rem;
  vertical-align: -0.125em;
}
.card .text{
  position: relative;
  margin-left: 10px;
}
.community_icon {
  width: 6rem;
  height: 6rem;
  position: relative;
  align-self: center;
  border-radius: 10px;
}

.nav-buttons {
  display: flex;
  margin-left: auto;
}

.nav-button {
  background: transparent;
  border: 0;
  color: #333;
  padding: 5px;
  margin-right: 10px;
  cursor: pointer;
}


.grid {
  display: grid;
  grid-template-columns: repeat(1, 1fr);
  gap: 20px;
}

.forum {
  font-family: Arial, sans-serif;
}

h1 {
  color: #333;
}
h3{
  margin-top: 0;
  margin-bottom: 0;
  margin-left: 0;
}
ul {
  list-style-type: none;
  padding: 0;
}

li {
  margin: 10px 0;
  display: grid;
  place-items: center;
}

.card-link {
  text-decoration: none;
  color: inherit;
  display: flex;
  justify-content: center;
  width: 100%;
}

/* From Uiverse.io by satyamchaudharydev */
/* this card is inspired form this - https://georgefrancis.dev/ */

.card {
  --border-radius: 0.75rem;
  --primary-color: #ffab00;
  --secondary-color: #3c3852;
  width: 75%;
  font-family: "Arial";
  padding: 1rem;
  cursor: pointer;
  border-radius: var(--border-radius);
  background: #f1f1f3;
  box-shadow: 0px 8px 16px 0px rgb(0 0 0 / 3%);
  position: relative;
  display: flex;
}

.card>*+* {
  margin-top: 1.1em;
}

.card .card__content {
  color: var(--secondary-color);
  font-size: 0.86rem;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  line-clamp: 2;
  -webkit-box-orient: vertical;
  max-width: 200px;
  overflow: hidden;
  text-overflow: ellipsis;
}

.card .card__title {
  padding: 0;
  font-size: 1.3rem;
  font-weight: bold;
}

.card .card__date {
  color: #6e6b80;
  font-size: 0.8rem;
}

.card .card__arrow {
  position: absolute;
  background: var(--primary-color);
  padding: 0.4rem;
  border-top-left-radius: var(--border-radius);
  border-bottom-right-radius: var(--border-radius);
  bottom: 0;
  right: 0;
  transition: 0.2s;
  display: flex;
  justify-content: center;
  align-items: center;
}

.card svg {
  transition: 0.2s;
}

/* hover */
.card:hover .card__title {
  color: var(--primary-color);
  text-decoration: underline;
}

.card:hover .card__arrow {
  background: #111;
}

.card:hover .card__arrow svg {
  transform: translateX(3px);
}
</style>