<template>
    <div>
        <main-nav :title="t('sidebar.topic')"/>

        <n-list class="main-content-wrap tags-wrap" bordered>
            <n-tabs type="line" animated @update:value="changeTab">
                <n-tab-pane name="hot" :tab="t('topic.hot')" />
                <n-tab-pane name="new" :tab="t('topic.new')" />
                <n-tab-pane v-if="store.state.userLogined"
                name="follow" :tab="t('topic.follow')" />
                <template v-if="store.state.userLogined" #suffix>
                    <n-tag v-model:checked="tagsChecked" checkable>
                        {{tagsEditText}}
                    </n-tag>
                </template>
            </n-tabs>
            <n-spin :show="loading">
                <n-space>
                    <tag-item
                        v-for="tag in tags"
                        :key="tag.id"
                        :tag="tag"
                        :showAction="store.state.userLogined && tagsChecked"
                        :checkFollowing="inFollowTab"
                    >
                    </tag-item>
                </n-space>
            </n-spin>
        </n-list>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch} from 'vue';
import { getTags } from '@/api/post';
import { useStore } from 'vuex';
import { useI18n } from 'vue-i18n';
import type { Item } from '@/types/Item';

const { t } = useI18n();
const store = useStore();
const tags = ref<Item.TagProps[]>([]);
const tagType = ref<"hot" | "new" | "follow">('hot');
const loading = ref(false);
const tagsChecked = ref(false)
const inFollowTab = ref(false)

watch(tagsChecked, () => {
    if (!tagsChecked.value) {
        window.$message.success("保存成功");
        store.commit("refreshTopicFollow")
    }
});
const tagsEditText = computed({  
    get: () => {  
        let text = t('topic.edit');
        if (tagsChecked.value) {
            text = t('topic.save');
        }
        return text;
    },
    set: (newVal) => {
        // do nothing
    },
});
const loadTags = () => {
    loading.value = true;
    getTags({
        type: tagType.value,
        num: 50,
    })
        .then((res) => {
            tags.value = res.topics;
            loading.value = false;
        })
        .catch((err) => {
            console.log(err);
            loading.value = false;
        });
};
const changeTab = (tab: "hot" | "new" | "follow") => {
    tagType.value = tab;
    if (tab == "follow") {
        inFollowTab.value = true
    } else {
        inFollowTab.value = false
    }
    loadTags();
};
onMounted(() => {
    loadTags();
});
</script>

<style lang="less" scoped>
.tags-wrap {
    padding: 20px;
}
.dark {
    .tags-wrap {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>