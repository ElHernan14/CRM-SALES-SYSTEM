<script setup lang="ts">
import { computed, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { storeToRefs } from 'pinia';

import {
  Building2,
  ChevronDown,
  Grid2X2,
  Loader2,
  LogIn,
  LogOut,
  Menu,
  Search,
  ShoppingBag,
  Store,
  UserPlus,
  UserRound,
  X,
} from 'lucide-vue-next';

import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';

import ThemeToggle from '@/shared/components/ThemeToggle.vue';

import { useAuthStore } from '@/modules/auth/stores/auth.store';
import { useLogout } from '@/modules/auth/composables/useLogout';

import { brand } from '@/shared/config/brand';

import { useStoreUiStore } from '../../stores/store-ui.store';
import { useStoreCarts } from '../../composables/useStoreCarts';

const router = useRouter();
const route = useRoute();

const auth = useAuthStore();
const storeUi = useStoreUiStore();

const { user } = storeToRefs(auth);

const { logout, isLoggingOut } = useLogout();

const mobileMenuOpen = ref(false);

const isAuthenticated = computed(() => {
  return Boolean(user.value);
});

const isBusinessUser = computed(() => {
  return Boolean(user.value?.company_id);
});

const { data: cartsData } = useStoreCarts();

const cartItemCount = computed(() => {
  return cartsData.value?.summary.item_count ?? 0;
});

const isCatalogActive = computed(() => {
  return route.name === 'store-catalog' || route.name === 'store-product';
});

function openCatalog() {
  mobileMenuOpen.value = false;

  router.push({
    name: 'store-catalog',
  });
}

function openCategories() {
  mobileMenuOpen.value = false;

  router.push({
    name: 'store-catalog',
    query: {
      focus: 'categories',
    },
  });
}

function openPurchases() {
  mobileMenuOpen.value = false;

  if (!isAuthenticated.value) {
    router.push('/login');
    return;
  }

  router.push('/store/purchases');
}

function openAccount() {
  mobileMenuOpen.value = false;

  router.push(isAuthenticated.value ? '/store/account' : '/login');
}

function openErp() {
  mobileMenuOpen.value = false;

  router.push('/erp/dashboard');
}
</script>

<template>
  <header class="sticky top-0 z-40 border-b border-border/70 bg-background/85 backdrop-blur-xl">
    <div class="mx-auto flex h-16 w-full max-w-[1500px] items-center gap-4 px-4 sm:px-6 lg:px-8">
      <!-- BRAND -->
      <RouterLink to="/store" class="flex shrink-0 items-center gap-3">
        <div
          class="flex h-10 w-10 items-center justify-center rounded-xl bg-primary font-bold text-primary-foreground shadow-sm"
        >
          N
        </div>

        <div class="hidden sm:block">
          <p class="text-lg font-semibold leading-none tracking-[-0.025em]">
            {{ brand.name }}
          </p>

          <p class="mt-1 text-[11px] font-medium text-muted-foreground">Store</p>
        </div>
      </RouterLink>

      <!-- DESKTOP NAVIGATION -->
      <nav
        class="hidden items-center gap-1 rounded-full border border-border/60 bg-muted/25 p-1 lg:flex"
      >
        <RouterLink
          to="/store"
          class="rounded-full px-4 py-2 text-sm font-medium text-muted-foreground transition hover:text-foreground"
          active-class="bg-background text-foreground shadow-sm"
          exact-active-class="bg-background text-foreground shadow-sm"
        >
          Discover
        </RouterLink>

        <button
          type="button"
          class="rounded-full px-4 py-2 text-sm font-medium text-muted-foreground transition hover:text-foreground"
          :class="isCatalogActive ? 'bg-background text-foreground shadow-sm' : ''"
          @click="openCatalog"
        >
          Catalog
        </button>

        <RouterLink
          to="/store/businesses"
          class="rounded-full px-4 py-2 text-sm font-medium text-muted-foreground transition hover:text-foreground"
          active-class="bg-background text-foreground shadow-sm"
        >
          Businesses
        </RouterLink>
      </nav>

      <!-- SEARCH -->
      <div class="mx-auto hidden max-w-xl flex-1 md:block">
        <button type="button" class="group relative block w-full text-left" @click="openCatalog">
          <Search
            class="pointer-events-none absolute left-4 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground transition group-hover:text-foreground"
          />

          <div
            class="flex h-11 items-center rounded-full border border-transparent bg-muted/40 pl-11 pr-4 text-sm text-muted-foreground transition group-hover:border-border group-hover:bg-muted/60"
          >
            Search products, categories or businesses
          </div>
        </button>
      </div>

      <!-- ACTIONS -->
      <div class="ml-auto flex items-center gap-1">
        <ThemeToggle />

        <template v-if="!isAuthenticated">
          <Button
            variant="ghost"
            class="hidden rounded-full px-4 sm:inline-flex"
            @click="router.push('/login')"
          >
            Sign in
          </Button>

          <Button class="hidden rounded-full px-4 sm:inline-flex" @click="router.push('/register')">
            Create account
          </Button>
        </template>

        <DropdownMenu v-else>
          <DropdownMenuTrigger as-child>
            <Button variant="ghost" class="hidden max-w-56 gap-2 rounded-full px-2.5 sm:flex">
              <div
                class="flex h-8 w-8 shrink-0 items-center justify-center rounded-full bg-primary/10 text-sm font-semibold text-primary"
              >
                {{ user?.email?.charAt(0).toUpperCase() ?? 'N' }}
              </div>

              <div class="hidden min-w-0 text-left xl:block">
                <p class="max-w-32 truncate text-xs font-semibold">
                  {{ user?.email }}
                </p>

                <p class="mt-0.5 text-[10px] text-muted-foreground">
                  {{ isBusinessUser ? 'Business member' : 'Personal account' }}
                </p>
              </div>

              <ChevronDown class="h-3.5 w-3.5 shrink-0 text-muted-foreground" />
            </Button>
          </DropdownMenuTrigger>

          <DropdownMenuContent align="end" class="w-72">
            <DropdownMenuLabel>
              <div>
                <p class="truncate text-sm font-semibold">
                  {{ user?.email }}
                </p>

                <p class="mt-1 text-xs font-normal text-muted-foreground">
                  {{
                    isBusinessUser ? 'Connected to a Nexora workspace' : 'Personal commerce account'
                  }}
                </p>
              </div>
            </DropdownMenuLabel>

            <DropdownMenuSeparator />

            <DropdownMenuItem @click="router.push('/store/account')">
              <UserRound class="mr-2 h-4 w-4" />
              My account
            </DropdownMenuItem>

            <DropdownMenuItem @click="router.push('/store/purchases')">
              <ShoppingBag class="mr-2 h-4 w-4" />
              My purchases
            </DropdownMenuItem>

            <DropdownMenuItem @click="openCatalog">
              <Search class="mr-2 h-4 w-4" />
              Browse catalog
            </DropdownMenuItem>

            <template v-if="isBusinessUser">
              <DropdownMenuSeparator />

              <DropdownMenuItem @click="openErp">
                <Building2 class="mr-2 h-4 w-4" />
                Open ERP workspace
              </DropdownMenuItem>
            </template>

            <DropdownMenuSeparator />

            <DropdownMenuItem
              class="text-destructive focus:text-destructive"
              :disabled="isLoggingOut"
              @select.prevent="logout"
            >
              <Loader2 v-if="isLoggingOut" class="mr-2 h-4 w-4 animate-spin" />

              <LogOut v-else class="mr-2 h-4 w-4" />

              {{ isLoggingOut ? 'Signing out...' : 'Sign out' }}
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>

        <Button
          v-if="isAuthenticated"
          variant="ghost"
          size="icon"
          class="relative rounded-full"
          aria-label="Open shopping cart"
          @click="storeUi.openCart()"
        >
          <ShoppingBag class="h-5 w-5" />

          <span
            v-if="cartItemCount > 0"
            class="absolute right-0 top-0 flex h-4 min-w-4 items-center justify-center rounded-full bg-primary px-1 text-[10px] font-semibold text-primary-foreground"
          >
            {{ cartItemCount > 99 ? '99+' : cartItemCount }}
          </span>
        </Button>

        <Button
          variant="ghost"
          size="icon"
          class="rounded-full lg:hidden"
          :aria-label="mobileMenuOpen ? 'Close navigation' : 'Open navigation'"
          @click="mobileMenuOpen = !mobileMenuOpen"
        >
          <X v-if="mobileMenuOpen" class="h-5 w-5" />

          <Menu v-else class="h-5 w-5" />
        </Button>
      </div>
    </div>

    <!-- MOBILE SEARCH -->
    <div class="border-t border-border/50 px-4 py-3 md:hidden">
      <button type="button" class="relative block w-full text-left" @click="openCatalog">
        <Search
          class="pointer-events-none absolute left-4 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
        />

        <div
          class="flex h-10 items-center rounded-full bg-muted/40 pl-11 pr-4 text-sm text-muted-foreground"
        >
          Search Nexora Store
        </div>
      </button>
    </div>

    <!-- MOBILE MENU -->
    <div v-if="mobileMenuOpen" class="border-t border-border bg-background px-4 py-4 lg:hidden">
      <nav class="space-y-1">
        <button
          type="button"
          class="flex w-full items-center gap-3 rounded-xl px-3 py-3 text-left text-sm font-medium transition hover:bg-muted"
          @click="
            mobileMenuOpen = false;
            router.push('/store');
          "
        >
          <Store class="h-4 w-4 text-primary" />
          Discover
        </button>

        <button
          type="button"
          class="flex w-full items-center gap-3 rounded-xl px-3 py-3 text-left text-sm font-medium transition hover:bg-muted"
          @click="openCatalog"
        >
          <Search class="h-4 w-4 text-primary" />
          Browse catalog
        </button>

        <button
          type="button"
          class="flex w-full items-center gap-3 rounded-xl px-3 py-3 text-left text-sm font-medium transition hover:bg-muted"
          @click="openCategories"
        >
          <Grid2X2 class="h-4 w-4 text-primary" />
          Categories
        </button>

        <button
          v-if="isAuthenticated"
          type="button"
          class="flex w-full items-center gap-3 rounded-xl px-3 py-3 text-left text-sm font-medium transition hover:bg-muted"
          @click="openPurchases"
        >
          <ShoppingBag class="h-4 w-4 text-primary" />
          My purchases
        </button>

        <button
          type="button"
          class="flex w-full items-center gap-3 rounded-xl px-3 py-3 text-left text-sm font-medium transition hover:bg-muted"
          @click="openAccount"
        >
          <UserRound class="h-4 w-4 text-primary" />

          {{ isAuthenticated ? 'My account' : 'Sign in' }}
        </button>

        <button
          v-if="isBusinessUser"
          type="button"
          class="flex w-full items-center gap-3 rounded-xl px-3 py-3 text-left text-sm font-medium transition hover:bg-muted"
          @click="openErp"
        >
          <Building2 class="h-4 w-4 text-primary" />
          ERP workspace
        </button>

        <template v-if="!isAuthenticated">
          <div class="my-3 border-t border-border" />

          <Button
            variant="outline"
            class="w-full rounded-full"
            @click="
              mobileMenuOpen = false;
              router.push('/login');
            "
          >
            <LogIn class="mr-2 h-4 w-4" />
            Sign in
          </Button>

          <Button
            class="mt-2 w-full rounded-full"
            @click="
              mobileMenuOpen = false;
              router.push('/register');
            "
          >
            <UserPlus class="mr-2 h-4 w-4" />
            Create account
          </Button>
        </template>
      </nav>
    </div>
  </header>
</template>
