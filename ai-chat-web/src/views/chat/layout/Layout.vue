<script setup lang='ts'>
import { computed } from 'vue'
import { NLayout, NLayoutContent } from 'naive-ui'
import { useRouter } from 'vue-router'
// import Sider from './sider/index.vue'
// import SiderRight from './SiderRight/index.vue'
// import Permission from './Permission.vue'
import { useBasicLayout } from '@/hooks/useBasicLayout'
import { useAppStore, useChatStore } from '@/store'

const router = useRouter()
const appStore = useAppStore()
const chatStore = useChatStore()
// const authStore = useAuthStore()

router.replace({ name: 'Chat', params: { uuid: chatStore.active } })

const { isMobile } = useBasicLayout()

const collapsed = computed(() => appStore.siderCollapsed)
/*
onBeforeMount(() => {
  const access_token = getCookieValue('sso_0voice_access_token')
  if (!access_token)
    window.location.href = 'https://user.0voice.com?sys=ai'
})

*/
// const needPermission = computed(() => !!authStore.session?.auth && !authStore.token)
// const needPermission = !localStorage.access_token
const getMobileClass = computed(() => {
  if (isMobile.value)
    return ['rounded-none', 'shadow-none']
  return ['border', 'rounded-md', 'shadow-md', 'dark:border-neutral-800']
})

const getContainerClass = computed(() => {
  return [
    'h-full',
    // { 'pl-[260px]': !isMobile.value && !collapsed.value },
    // { 'right-[0]': !isMobile.value && !collapsed.value },
  ]
})
</script>

<template>
  <div class="h-full dark:bg-[#24272e] transition-all" :class="[isMobile ? 'p-0' : 'p-4']">
    <div class="h-full overflow-hidden" :class="getMobileClass">
      <NLayout class="z-40 transition" :class="getContainerClass" has-sider>
        <!-- <Sider /> -->
        <NLayoutContent class="h-full">
          <RouterView v-slot="{ Component, route }">
            <component :is="Component" :key="route.fullPath" />
          </RouterView>
        </NLayoutContent>
        <!-- <SiderRight /> -->
      </NLayout>
    </div>
    <!-- <Permission :visible="needPermission" /> -->
  </div>
</template>
