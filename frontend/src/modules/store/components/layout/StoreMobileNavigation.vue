<script setup lang="ts">
import { computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { storeToRefs } from 'pinia';

import { Building2, Home, Search, ShoppingBag, UserRound } from 'lucide-vue-next';

import { useAuthStore } from '@/modules/auth/stores/auth.store';

const route = useRoute();
const router = useRouter();

const auth = useAuthStore();
const { user } = storeToRefs(auth);

const isAuthenticated = computed(() => Boolean(user.value));

function isActive(path: string) {
  if (path === '/store') {
    return route.path === '/store';
  }

  return route.path.startsWith(path);
}

function openPurchases() {
  if (!isAuthenticated.value) {
    router.push({
      name: 'login',
      query: {
        redirect: '/store/purchases',
      },
    });

    return;
  }

  router.push('/store/purchases');
}
</script>

<template>
  <nav
    class="fixed inset-x-0 bottom-0 z-40 border-t border-border bg-background/95 pb-[env(safe-area-inset-bottom)] backdrop-blur-xl md:hidden"
    aria-label="Store mobile navigation"
  >
    <div class="grid h-16 grid-cols-4">
      <button
        type="button"
        class="flex flex-col items-center justify-center gap-1 text-[11px] font-medium transition"
        :class="isActive('/store') ? 'text-primary' : 'text-muted-foreground'"
        @click="router.push('/store')"
      >
        <Home class="h-5 w-5" />
        Home
      </button>

      <button
        type="button"
        class="flex flex-col items-center justify-center gap-1 text-[11px] font-medium transition"
        :class="isActive('/store/catalog') ? 'text-primary' : 'text-muted-foreground'"
        @click="router.push('/store/catalog')"
      >
        <Search class="h-5 w-5" />
        Catalog
      </button>

      <button
        type="button"
        class="flex flex-col items-center justify-center gap-1 text-[11px] font-medium transition"
        :class="isActive('/store/businesses') ? 'text-primary' : 'text-muted-foreground'"
        @click="router.push('/store/businesses')"
      >
        <Building2 class="h-5 w-5" />
        Businesses
      </button>

      <button
        type="button"
        class="flex flex-col items-center justify-center gap-1 text-[11px] font-medium transition"
        :class="isActive('/store/purchases') ? 'text-primary' : 'text-muted-foreground'"
        @click="openPurchases"
      >
        <ShoppingBag v-if="isAuthenticated" class="h-5 w-5" />
        <UserRound v-else class="h-5 w-5" />

        {{ isAuthenticated ? 'Purchases' : 'Sign in' }}
      </button>
    </div>
  </nav>
</template>
