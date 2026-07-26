<script setup lang="ts">
import { Compass, Home, ShoppingBag, UserRound, Building2 } from 'lucide-vue-next';
import { useStoreUiStore } from '../../stores/store-ui.store';
import { computed } from 'vue';

import { useAuthStore } from '@/modules/auth/stores/auth.store';
import { router } from '@/app/router';

const auth = useAuthStore();

const isAuthenticated = computed(() => {
  return auth.isAuthenticated;
});

const storeUi = useStoreUiStore();

function openAccount() {
  router.push(isAuthenticated.value ? '/store/account' : '/login');
}
</script>

<template>
  <nav
    class="fixed inset-x-0 bottom-0 z-40 border-t border-border bg-background/95 px-3 py-2 backdrop-blur-xl md:hidden"
  >
    <div class="grid gap-1" :class="isAuthenticated ? 'grid-cols-4' : 'grid-cols-3'">
      <RouterLink
        to="/store"
        class="flex flex-col items-center gap-1 rounded-xl px-2 py-2 text-xs text-muted-foreground transition hover:bg-muted hover:text-foreground"
      >
        <Home class="h-5 w-5" />
        Home
      </RouterLink>

      <RouterLink
        to="/store/catalog"
        class="flex flex-col items-center gap-1 rounded-xl px-2 py-2 text-xs text-muted-foreground transition hover:bg-muted hover:text-foreground"
      >
        <Compass class="h-5 w-5" />
        Explore
      </RouterLink>

      <button
        v-if="isAuthenticated"
        class="flex flex-col items-center gap-1 rounded-xl px-2 py-2 text-xs text-muted-foreground transition hover:bg-muted hover:text-foreground"
        @click="storeUi.openCart()"
      >
        <ShoppingBag class="h-5 w-5" />
        Cart
      </button>

      <button
        type="button"
        class="flex flex-col items-center gap-1 rounded-xl px-2 py-2 text-xs text-muted-foreground transition hover:bg-muted hover:text-foreground"
        @click="router.push('/store/businesses')"
      >
        <Building2 class="h-4 w-4 text-primary" />
        Businesses
      </button>

      <button
        class="flex flex-col items-center gap-1 rounded-xl px-2 py-2 text-xs text-muted-foreground transition hover:bg-muted hover:text-foreground"
        @click="openAccount"
      >
        <UserRound class="h-5 w-5" />

        {{ isAuthenticated ? 'Account' : 'Sign in' }}
      </button>
    </div>
  </nav>
</template>
