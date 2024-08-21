import i18n from '@/i18n';
import { createRouter, createWebHashHistory } from "vue-router";

const { t } = i18n.global;
const routes = [
  {
    path: "/",
    name: "home",
    meta: {
      title: t('sidebar.home'),
      keepAlive: true,
    },
    component: () => import("@/views/Home.vue"),
  },
  {
    path: "/post",
    name: "post",
    meta: {
      title: t('router.post'),
    },
    component: () => import("@/views/Post.vue"),
  },
  {
    path: "/topic",
    name: "topic",
    meta: {
      title: t('sidebar.topic'),
    },
    component: () => import("@/views/Topic.vue"),
  },
  {
    path: "/community",
    name: "community",
    meta: {
      title: t('sidebar.community'),
    },
    component: () => import("@/views/Community.vue"),
  },
  {
    path: "/announcement",
    name: "announcement",
    meta: {
      title: t('sidebar.announcement'),
    },
    component: () => import("@/views/Announcement.vue"),
  },
  {
    path: "/profile",
    name: "profile",
    meta: {
      title: t('sidebar.profile'),
    },
    component: () => import("@/views/Profile.vue"),
  },
  {
    path: "/u",
    name: "user",
    meta: {
      title: t('router.user'),
    },
    component: () => import("@/views/User.vue"),
  },
  {
    path: "/messages",
    name: "messages",
    meta: {
      title: t('sidebar.message'),
    },
    component: () => import("@/views/Messages.vue"),
  },
  {
    path: "/collection",
    name: "collection",
    meta: {
      title: t('sidebar.archive'),
    },
    component: () => import("@/views/Collection.vue"),
  },
  {
    path: "/contacts",
    name: "contacts",
    meta: {
      title: t('sidebar.friend'),
    },
    component: () => import("@/views/Contacts.vue"),
  },
  {
    path: "/following",
    name: "following",
    meta: {
      title: t('router.following'),
    },
    component: () => import("@/views/Following.vue"),
  },
  {
    path: "/wallet",
    name: "wallet",
    meta: {
      title: t('sidebar.wallet'),
    },
    component: () => import("@/views/Wallet.vue"),
  },
  {
    path: "/setting",
    name: "setting",
    meta: {
      title: t("sidebar.setting"),
    },
    component: () => import("@/views/Setting.vue"),
  },
  {
    path: "/404",
    name: "404",
    meta: {
      title: "404",
    },
    component: () => import("@/views/404.vue"),
  },
  {
    path: "/:pathMatch(.*)",
    redirect: "/404",
  },
  {
    path: '/community/:id',
    name: 'CommunityPosts',
    meta: {
      title: t('sidebar.community'),
    },
    props: true,
    component: () => import("@/views/Community-Posts.vue"),
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

router.beforeEach((to, from, next) => {

  document.title = `${to.meta.title} | ${t('router.siteTitle')}`;
  next();
});

export default router;
