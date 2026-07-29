<script setup lang="ts">
import { computed, ref } from 'vue';
import { useRouter } from 'vue-router';
import { storeToRefs } from 'pinia';

import {
  Building2,
  Heart,
  Loader2,
  LogIn,
  LogOut,
  Menu,
  Search,
  ShoppingBag,
  Store,
  UserRound,
  X,
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

function closeMobileMenu() {
  mobileMenuOpen.value = false;
}

function navigate(path: string) {
  closeMobileMenu();
  router.push(path);
}

function openCart() {
  closeMobileMenu();

  if (!isAuthenticated.value) {
    router.push({
      name: 'login',
      query: {
        redirect: '/store',
      },
    });

    return;
  }

  storeUi.openCart();
}
</script>

<template>
  <header class="sticky top-0 z-40 border-b border-border/70 bg-background/90 backdrop-blur-xl">
    <div
      class="mx-auto flex h-16 w-full max-w-[1500px] items-center gap-2 px-4 sm:gap-4 sm:px-6 lg:px-8"
    >
      <!-- BRAND -->
      <RouterLink to="/" class="flex min-w-0 shrink-0 items-center gap-2">
        <div
          class="nexora-gradient nexora-glow flex h-9 w-9 shrink-0 items-center justify-center rounded-xl bg-primary font-bold text-primary-foreground shadow-sm shadow-primary/20"
        >
          N
        </div>

        <span
          class="nexora-gradient-text hidden truncate text-lg font-semibold tracking-tight min-[390px]:block sm:text-xl"
        >
          {{ brand.name }}
        </span>
      </RouterLink>

      <!-- DESKTOP NAVIGATION -->
      <nav class="hidden items-center gap-1 lg:flex">
        <RouterLink
          to="/store/catalog"
          class="rounded-full px-3 py-2 text-sm font-medium text-muted-foreground transition-colors duration-200 hover:bg-muted/50 hover:text-foreground"
        >
          Explore
        </RouterLink>

        <RouterLink
          to="/store/businesses"
          class="rounded-full px-3 py-2 text-sm font-medium text-muted-foreground transition-colors duration-200 hover:bg-muted/50 hover:text-foreground"
        >
          Businesses
        </RouterLink>
      </nav>

      <!-- DESKTOP SEARCH -->
      <div class="mx-auto hidden max-w-xl flex-1 md:block">
        <div class="relative">
          <Search
            class="pointer-events-none absolute left-4 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
          />

          <Input
            class="h-11 rounded-full bg-muted/40 pl-11 pr-4"
            placeholder="Search products, categories or businesses..."
            @focus="router.push('/store/catalog')"
          />
        </div>
      </div>

      <!-- ACTIONS -->
      <div class="ml-auto flex shrink-0 items-center gap-0.5 sm:gap-1">
        <!-- FAVORITES: desktop only until feature exists -->
        <Button
          variant="ghost"
          size="icon"
          class="hidden rounded-full lg:inline-flex"
          aria-label="Saved products"
        >
          <Heart class="h-5 w-5" />
        </Button>

        <!-- SIGN IN DESKTOP -->
        <Button
          v-if="!isAuthenticated"
          variant="ghost"
          class="hidden rounded-full px-4 sm:inline-flex"
          @click="router.push('/login')"
        >
          Sign in
        </Button>

        <!-- ACCOUNT DESKTOP/TABLET -->
        <DropdownMenu v-if="isAuthenticated">
          <DropdownMenuTrigger as-child>
            <Button
              variant="ghost"
              size="icon"
              class="hidden rounded-full sm:inline-flex"
              aria-label="Open account menu"
            >
              <UserRound class="h-5 w-5" />
            </Button>
          </DropdownMenuTrigger>

          <DropdownMenuContent align="end" class="w-64">
            <DropdownMenuLabel>
              <div class="min-w-0">
                <p class="truncate text-sm font-medium">
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
              <Store class="mr-2 h-4 w-4" />
              Browse Store
            </DropdownMenuItem>

            <DropdownMenuItem @click="router.push('/store/account')">
              <UserRound class="mr-2 h-4 w-4" />
              My account
            </DropdownMenuItem>

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

        <!-- CART -->
        <Button
          v-if="isAuthenticated"
          variant="ghost"
          size="icon"
          class="relative rounded-full"
          aria-label="Open cart"
          @click="openCart"
        >
          <ShoppingBag class="h-5 w-5" />

          <span
            v-if="cartItemCount > 0"
            class="absolute right-0.5 top-0.5 flex h-4 min-w-4 items-center justify-center rounded-full bg-primary px-1 text-[10px] font-semibold text-primary-foreground"
          >
            {{ cartItemCount > 99 ? '99+' : cartItemCount }}
          </span>
        </Button>

        <!-- DESKTOP THEME -->
        <div class="hidden sm:block">
          <ThemeToggle />
        </div>

        <!-- MOBILE MENU -->
        <Button
          variant="ghost"
          size="icon"
          class="rounded-full lg:hidden"
          :aria-label="mobileMenuOpen ? 'Close Store menu' : 'Open Store menu'"
          @click="mobileMenuOpen = !mobileMenuOpen"
        >
          <X v-if="mobileMenuOpen" class="h-5 w-5" />
          <Menu v-else class="h-5 w-5" />
        </Button>
      </div>
    </div>

    <!-- MOBILE/TABLET MENU -->
    <div
      v-if="mobileMenuOpen"
      class="max-h-[calc(100vh-4rem)] overflow-y-auto border-t border-border bg-background/95 px-4 py-4 backdrop-blur-xl lg:hidden"
    >
      <div v-if="isAuthenticated" class="mb-4 rounded-2xl border border-border bg-muted/25 p-4">
        <div class="flex min-w-0 items-center gap-3">
          <div
            class="flex h-10 w-10 shrink-0 items-center justify-center rounded-full bg-primary/10 text-sm font-semibold text-primary"
          >
            {{ user?.email?.charAt(0).toUpperCase() ?? 'N' }}
          </div>

          <div class="min-w-0">
            <p class="truncate text-sm font-semibold">
              {{ user?.email }}
            </p>

            <p class="mt-1 text-xs text-muted-foreground">
              {{ isBusinessUser ? 'Business account' : 'Personal account' }}
            </p>
          </div>
        </div>
      </div>

      <nav class="space-y-1">
        <button
          type="button"
          class="flex w-full items-center gap-3 rounded-xl px-3 py-3 text-left text-sm font-medium transition hover:bg-muted"
          @click="navigate('/store')"
        >
          <Store class="h-4 w-4 text-primary" />
          Store home
        </button>

        <button
          type="button"
          class="flex w-full items-center gap-3 rounded-xl px-3 py-3 text-left text-sm font-medium transition hover:bg-muted"
          @click="navigate('/store/catalog')"
        >
          <Search class="h-4 w-4 text-primary" />
          Explore catalog
        </button>

        <button
          type="button"
          class="flex w-full items-center gap-3 rounded-xl px-3 py-3 text-left text-sm font-medium transition hover:bg-muted"
          @click="navigate('/store/businesses')"
        >
          <Building2 class="h-4 w-4 text-primary" />
          Browse businesses
        </button>

        <template v-if="isAuthenticated">
          <button
            type="button"
            class="flex w-full items-center gap-3 rounded-xl px-3 py-3 text-left text-sm font-medium transition hover:bg-muted"
            @click="navigate('/store/purchases')"
          >
            <ShoppingBag class="h-4 w-4 text-primary" />
            My purchases
          </button>

          <button
            type="button"
            class="flex w-full items-center gap-3 rounded-xl px-3 py-3 text-left text-sm font-medium transition hover:bg-muted"
            @click="navigate('/store/account')"
          >
            <UserRound class="h-4 w-4 text-primary" />
            My account
          </button>

          <button
            v-if="isBusinessUser"
            type="button"
            class="flex w-full items-center gap-3 rounded-xl px-3 py-3 text-left text-sm font-medium transition hover:bg-muted"
            @click="navigate('/erp/dashboard')"
          >
            <Building2 class="h-4 w-4 text-primary" />
            Open ERP workspace
          </button>
        </template>
      </nav>

      <div class="my-4 border-t border-border" />

      <div class="flex items-center justify-between rounded-xl px-3 py-2">
        <div>
          <p class="text-sm font-medium">Appearance</p>
          <p class="mt-1 text-xs text-muted-foreground">Change the Store theme</p>
        </div>

        <ThemeToggle />
      </div>

      <template v-if="!isAuthenticated">
        <Button variant="outline" class="mt-4 w-full rounded-full" @click="navigate('/login')">
          <LogIn class="mr-2 h-4 w-4" />
          Sign in
        </Button>

        <Button class="mt-2 w-full rounded-full" @click="navigate('/register')">
          Create account
        </Button>
      </template>

      <Button
        v-else
        variant="ghost"
        class="mt-4 w-full rounded-full text-destructive hover:text-destructive"
        :disabled="isLoggingOut"
        @click="logout"
      >
        <Loader2 v-if="isLoggingOut" class="mr-2 h-4 w-4 animate-spin" />

        <LogOut v-else class="mr-2 h-4 w-4" />

        {{ isLoggingOut ? 'Signing out...' : 'Sign out' }}
      </Button>
    </div>

    <!-- MOBILE SEARCH -->
    <div class="border-t border-border/50 px-4 py-3 md:hidden">
      <div class="relative">
        <Search
          class="pointer-events-none absolute left-4 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
        />

        <Input
          class="h-10 rounded-full bg-muted/40 pl-11"
          placeholder="Search products or businesses..."
          @focus="router.push('/store/catalog')"
        />
      </div>
    </div>
  </header>
</template>
