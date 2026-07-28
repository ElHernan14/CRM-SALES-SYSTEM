<script setup lang="ts">
import { computed, ref } from 'vue';
import { useRouter } from 'vue-router';
import { storeToRefs } from 'pinia';

import {
  ArrowRight,
  Building2,
  ChevronDown,
  Loader2,
  LogIn,
  LogOut,
  Menu,
  ShoppingBag,
  Store,
  UserPlus,
  UserRound,
  X,
} from 'lucide-vue-next';

import { Button } from '@/components/ui/button';

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';

import ThemeToggle from '@/shared/components/ThemeToggle.vue';

import { brand } from '@/shared/config/brand';

import { useAuthStore } from '@/modules/auth/stores/auth.store';
import { useLogout } from '@/modules/auth/composables/useLogout';

const router = useRouter();

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

const primaryWorkspaceLabel = computed(() => {
  if (!isAuthenticated.value) {
    return 'Get started';
  }

  return isBusinessUser.value ? 'Open workspace' : 'Explore Store';
});

function scrollToSection(sectionId: string) {
  mobileMenuOpen.value = false;

  const section = document.getElementById(sectionId);

  section?.scrollIntoView({
    behavior: 'smooth',
    block: 'start',
  });
}

function openPrimaryExperience() {
  mobileMenuOpen.value = false;

  if (!isAuthenticated.value) {
    router.push('/register');
    return;
  }

  router.push(isBusinessUser.value ? '/erp/dashboard' : '/store');
}

function openStore() {
  mobileMenuOpen.value = false;

  router.push('/store');
}

function openAccount() {
  mobileMenuOpen.value = false;

  if (!isAuthenticated.value) {
    router.push('/login');
    return;
  }

  router.push('/store/account');
}
</script>

<template>
  <header
    class="fixed inset-x-0 top-0 z-50 border-b border-border/60 bg-background/75 backdrop-blur-xl"
  >
    <div class="mx-auto flex h-16 w-full max-w-[1500px] items-center gap-4 px-4 sm:px-6 lg:px-8">
      <!-- BRAND -->
      <RouterLink to="/" class="flex shrink-0 items-center gap-3">
        <div
          class="nexora-gradient nexora-glow flex h-10 w-10 items-center justify-center rounded-xl bg-primary font-bold text-primary-foreground shadow-sm shadow-primary/20"
        >
          N
        </div>

        <div class="hidden min-[390px]:block">
          <p class="nexora-gradient-text text-lg font-semibold leading-none tracking-[-0.03em]">
            {{ brand.name }}
          </p>

          <p class="mt-1 hidden text-[11px] font-medium text-muted-foreground sm:block">
            Commerce & operations
          </p>
        </div>
      </RouterLink>

      <!-- DESKTOP NAVIGATION -->
      <nav
        class="mx-auto hidden items-center gap-1 rounded-full border border-border/60 bg-muted/25 p-1 lg:flex"
      >
        <button
          type="button"
          class="rounded-full px-4 py-2 text-sm font-medium text-muted-foreground transition hover:bg-background hover:text-foreground hover:shadow-sm"
          @click="scrollToSection('nexora-product')"
        >
          Product
        </button>

        <button
          type="button"
          class="rounded-full px-4 py-2 text-sm font-medium text-muted-foreground transition hover:bg-background hover:text-foreground hover:shadow-sm"
          @click="scrollToSection('nexora-solutions')"
        >
          Solutions
        </button>

        <button
          type="button"
          class="rounded-full px-4 py-2 text-sm font-medium text-muted-foreground transition hover:bg-background hover:text-foreground hover:shadow-sm"
          @click="scrollToSection('nexora-commerce')"
        >
          Commerce
        </button>

        <button
          type="button"
          class="rounded-full px-4 py-2 text-sm font-medium text-muted-foreground transition hover:bg-background hover:text-foreground hover:shadow-sm"
          @click="scrollToSection('nexora-architecture')"
        >
          Platform
        </button>
      </nav>

      <!-- ACTIONS -->
      <div class="ml-auto flex items-center gap-1.5">
        <ThemeToggle />

        <template v-if="!isAuthenticated">
          <Button
            variant="ghost"
            class="hidden rounded-full px-4 sm:inline-flex"
            @click="router.push('/login')"
          >
            Sign in
          </Button>

          <Button class="hidden rounded-full px-5 sm:inline-flex" @click="router.push('/register')">
            Get started

            <ArrowRight class="ml-2 h-4 w-4" />
          </Button>
        </template>

        <DropdownMenu v-else>
          <DropdownMenuTrigger as-child>
            <Button
              variant="ghost"
              class="flex h-10 max-w-56 gap-2 rounded-full px-1.5 sm:px-2.5"
              aria-label="Open account menu"
            >
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
                  {{ isBusinessUser ? 'Business workspace' : 'Personal account' }}
                </p>
              </div>

              <ChevronDown class="hidden h-3.5 w-3.5 text-muted-foreground xl:block" />
            </Button>
          </DropdownMenuTrigger>

          <DropdownMenuContent align="end" class="w-72">
            <DropdownMenuLabel>
              <p class="truncate text-sm font-semibold">
                {{ user?.email }}
              </p>

              <p class="mt-1 text-xs font-normal text-muted-foreground">
                {{ isBusinessUser ? 'Connected to a Nexora business' : 'Nexora personal account' }}
              </p>
            </DropdownMenuLabel>

            <DropdownMenuSeparator />

            <DropdownMenuItem v-if="isBusinessUser" @click="router.push('/erp/dashboard')">
              <Building2 class="mr-2 h-4 w-4" />
              Open business workspace
            </DropdownMenuItem>

            <DropdownMenuItem @click="router.push('/store')">
              <Store class="mr-2 h-4 w-4" />
              Explore Store
            </DropdownMenuItem>

            <DropdownMenuItem @click="router.push('/store/purchases')">
              <ShoppingBag class="mr-2 h-4 w-4" />
              My purchases
            </DropdownMenuItem>

            <DropdownMenuItem @click="openAccount">
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

        <Button
          v-if="isAuthenticated"
          class="hidden rounded-full px-5 sm:inline-flex"
          @click="openPrimaryExperience"
        >
          {{ primaryWorkspaceLabel }}

          <ArrowRight class="ml-2 h-4 w-4" />
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

    <!-- MOBILE MENU -->
    <div
      v-if="mobileMenuOpen"
      class="max-h-[calc(100vh-4rem)] overflow-y-auto border-t border-border bg-background/95 px-4 py-4 backdrop-blur-xl lg:hidden"
    >
      <nav class="space-y-1">
        <button
          type="button"
          class="w-full rounded-xl px-3 py-3 text-left text-sm font-medium transition hover:bg-muted"
          @click="scrollToSection('nexora-product')"
        >
          Product
        </button>

        <button
          type="button"
          class="w-full rounded-xl px-3 py-3 text-left text-sm font-medium transition hover:bg-muted"
          @click="scrollToSection('nexora-solutions')"
        >
          Solutions
        </button>

        <button
          type="button"
          class="w-full rounded-xl px-3 py-3 text-left text-sm font-medium transition hover:bg-muted"
          @click="scrollToSection('nexora-commerce')"
        >
          Commerce
        </button>

        <button
          type="button"
          class="w-full rounded-xl px-3 py-3 text-left text-sm font-medium transition hover:bg-muted"
          @click="scrollToSection('nexora-architecture')"
        >
          Platform
        </button>

        <button
          type="button"
          class="flex w-full items-center gap-3 rounded-xl px-3 py-3 text-left text-sm font-medium transition hover:bg-muted"
          @click="openStore"
        >
          <Store class="h-4 w-4 text-primary" />
          Explore Store
        </button>

        <div class="my-3 border-t border-border" />

        <template v-if="!isAuthenticated">
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

        <template v-else>
          <div class="mb-3 rounded-2xl border border-border bg-muted/25 p-4">
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
                  {{ isBusinessUser ? 'Business workspace' : 'Personal account' }}
                </p>
              </div>
            </div>
          </div>

          <button
            v-if="isBusinessUser"
            type="button"
            class="flex w-full items-center gap-3 rounded-xl px-3 py-3 text-left text-sm font-medium transition hover:bg-muted"
            @click="
              mobileMenuOpen = false;
              router.push('/erp/dashboard');
            "
          >
            <Building2 class="h-4 w-4 text-primary" />
            Open business workspace
          </button>

          <button
            type="button"
            class="flex w-full items-center gap-3 rounded-xl px-3 py-3 text-left text-sm font-medium transition hover:bg-muted"
            @click="
              mobileMenuOpen = false;
              router.push('/store/purchases');
            "
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
            My account
          </button>

          <Button class="mt-3 w-full rounded-full" @click="openPrimaryExperience">
            {{ primaryWorkspaceLabel }}

            <ArrowRight class="ml-2 h-4 w-4" />
          </Button>

          <Button
            variant="ghost"
            class="mt-2 w-full rounded-full text-destructive hover:text-destructive"
            :disabled="isLoggingOut"
            @click="logout"
          >
            <Loader2 v-if="isLoggingOut" class="mr-2 h-4 w-4 animate-spin" />

            <LogOut v-else class="mr-2 h-4 w-4" />

            {{ isLoggingOut ? 'Signing out...' : 'Sign out' }}
          </Button>
        </template>
      </nav>
    </div>
  </header>
</template>
