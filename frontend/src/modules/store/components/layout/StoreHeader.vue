<script setup lang="ts">
import { computed } from 'vue';
import { useRouter } from 'vue-router';
import { storeToRefs } from 'pinia';

import { Heart, Menu, Search, ShoppingBag, UserRound } from 'lucide-vue-next';

import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

import { useAuthStore } from '@/modules/auth/stores/auth.store';

const router = useRouter();

const auth = useAuthStore();
const { user } = storeToRefs(auth);

const isAuthenticated = computed(() => {
  return Boolean(user.value);
});

function goToAccount() {
  if (!isAuthenticated.value) {
    router.push('/login');
    return;
  }

  router.push('/store/purchases');
}
</script>

<template>
  <header class="sticky top-0 z-40 border-b border-border/70 bg-background/90 backdrop-blur-xl">
    <div class="mx-auto flex h-16 max-w-[1500px] items-center gap-4 px-4 sm:px-6 lg:px-8">
      <RouterLink to="/store" class="shrink-0 text-xl font-semibold tracking-tight text-foreground">
        SYNER
      </RouterLink>

      <nav class="hidden items-center gap-6 lg:flex">
        <RouterLink
          to="/store/catalog"
          class="text-sm font-medium text-muted-foreground transition hover:text-foreground"
        >
          Explore
        </RouterLink>

        <button
          type="button"
          class="text-sm font-medium text-muted-foreground transition hover:text-foreground"
        >
          Categories
        </button>

        <button
          type="button"
          class="text-sm font-medium text-muted-foreground transition hover:text-foreground"
        >
          New arrivals
        </button>
      </nav>

      <div class="mx-auto hidden max-w-xl flex-1 md:block">
        <div class="relative">
          <Search
            class="pointer-events-none absolute left-4 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
          />

          <Input
            class="h-11 rounded-full bg-muted/40 pl-11 pr-4"
            placeholder="Search products, categories or stores..."
            @focus="router.push('/store/catalog')"
          />
        </div>
      </div>

      <div class="ml-auto flex items-center gap-1">
        <Button variant="ghost" size="icon" class="hidden rounded-full sm:inline-flex">
          <Heart class="h-5 w-5" />
        </Button>

        <Button variant="ghost" size="icon" class="rounded-full" @click="goToAccount">
          <UserRound class="h-5 w-5" />
        </Button>

        <Button variant="ghost" size="icon" class="relative rounded-full">
          <ShoppingBag class="h-5 w-5" />

          <span
            class="absolute right-0.5 top-0.5 flex h-4 min-w-4 items-center justify-center rounded-full bg-primary px-1 text-[10px] font-semibold text-primary-foreground"
          >
            0
          </span>
        </Button>

        <Button variant="ghost" size="icon" class="rounded-full lg:hidden">
          <Menu class="h-5 w-5" />
        </Button>
      </div>
    </div>

    <div class="border-t border-border/50 px-4 py-3 md:hidden">
      <div class="relative">
        <Search
          class="pointer-events-none absolute left-4 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
        />

        <Input
          class="rounded-full bg-muted/40 pl-11"
          placeholder="Search products..."
          @focus="router.push('/store/catalog')"
        />
      </div>
    </div>
  </header>
</template>
