<script setup lang="ts">
import { computed } from 'vue';
import { useRouter } from 'vue-router';
import { storeToRefs } from 'pinia';

import {
  Building2,
  Heart,
  Loader2,
  LogOut,
  Menu,
  Search,
  ShoppingBag,
  UserRound,
} from 'lucide-vue-next';

import { useLogout } from '@/modules/auth/composables/useLogout';
import { useStoreCarts } from '../../composables/useStoreCarts';

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';

import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import ThemeToggle from '@/shared/components/ThemeToggle.vue';

import { useAuthStore } from '@/modules/auth/stores/auth.store';
import { brand } from '@/shared/config/brand';
import { useStoreUiStore } from '../../stores/store-ui.store';

const router = useRouter();

const storeUi = useStoreUiStore();
const auth = useAuthStore();
const { user } = storeToRefs(auth);

const { logout, isLoggingOut } = useLogout();

const isAuthenticated = computed(() => {
  return Boolean(user.value);
});

const { data: cartsData } = useStoreCarts();

const cartItemCount = computed(() => {
  return cartsData.value?.summary.item_count ?? 0;
});

const isBusinessUser = computed(() => {
  return Boolean(user.value?.company_id);
});

function openAuthenticatedArea() {
  if (!isAuthenticated.value) {
    router.push('/login');
    return;
  }

  router.push(isBusinessUser.value ? '/erp/dashboard' : '/store/purchases');
}
</script>

<template>
  <header class="sticky top-0 z-40 border-b border-border/70 bg-background/90 backdrop-blur-xl">
    <div class="mx-auto flex h-16 max-w-[1500px] items-center gap-4 px-4 sm:px-6 lg:px-8">
      <RouterLink to="/" class="flex shrink-0 items-center gap-2">
        <div
          class="flex h-9 w-9 items-center justify-center rounded-xl bg-primary font-bold text-primary-foreground"
        >
          N
        </div>

        <span class="text-xl font-semibold tracking-tight">
          {{ brand.name }}
        </span>
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

        <Button
          v-if="!isAuthenticated"
          variant="ghost"
          class="hidden rounded-full px-4 sm:inline-flex"
          @click="router.push('/login')"
        >
          Sign in
        </Button>

        <DropdownMenu v-if="isAuthenticated">
          <DropdownMenuTrigger as-child>
            <Button variant="ghost" size="icon" class="rounded-full">
              <UserRound class="h-5 w-5" />
            </Button>
          </DropdownMenuTrigger>

          <DropdownMenuContent align="end" class="w-64">
            <DropdownMenuLabel>
              <div>
                <p class="text-sm font-medium">
                  {{ user?.email }}
                </p>

                <p class="mt-1 text-xs font-normal text-muted-foreground">
                  {{ isBusinessUser ? 'Business account' : 'Personal account' }}
                </p>
              </div>
            </DropdownMenuLabel>

            <DropdownMenuSeparator />

            <DropdownMenuItem v-if="isBusinessUser" @click="router.push('/erp/dashboard')">
              <Building2 class="mr-2 h-4 w-4" />
              Open ERP workspace
            </DropdownMenuItem>

            <DropdownMenuItem @click="router.push('/store/purchases')">
              <ShoppingBag class="mr-2 h-4 w-4" />
              My purchases
            </DropdownMenuItem>

            <DropdownMenuItem @click="router.push('/store')">
              <Search class="mr-2 h-4 w-4" />
              Browse Store
            </DropdownMenuItem>

            <DropdownMenuSeparator />

            <DropdownMenuItem
              class="text-destructive focus:text-destructive"
              :disabled="isLoggingOut"
              @click="logout"
            >
              <Loader2 v-if="isLoggingOut" class="mr-2 h-4 w-4 animate-spin" />

              <LogOut v-else class="mr-2 h-4 w-4" />

              {{ isLoggingOut ? 'Signing out...' : 'Sign out' }}
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>

        <Button
          variant="ghost"
          size="icon"
          class="relative rounded-full"
          @click="storeUi.openCart()"
        >
          <ShoppingBag class="h-5 w-5" />

          <span
            v-if="cartItemCount > 0"
            class="absolute right-0.5 top-0.5 flex h-4 min-w-4 items-center justify-center rounded-full bg-primary px-1 text-[10px] font-semibold text-primary-foreground"
          >
            {{ cartItemCount > 99 ? '99+' : cartItemCount }}
          </span>
        </Button>

        <Button variant="ghost" size="icon" class="rounded-full lg:hidden">
          <Menu class="h-5 w-5" />
        </Button>
        <ThemeToggle />
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
