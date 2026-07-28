<script setup lang="ts">
import { computed } from 'vue';
import { useRouter } from 'vue-router';
import { storeToRefs } from 'pinia';

import {
  ArrowRight,
  Building2,
  LogOut,
  Menu,
  ShoppingBag,
  UserRound,
  Loader2,
} from 'lucide-vue-next';

import { useLogout } from '@/modules/auth/composables/useLogout';

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

const router = useRouter();

const { logout, isLoggingOut } = useLogout();

const auth = useAuthStore();
const { user } = storeToRefs(auth);

const isAuthenticated = computed(() => {
  return auth.isAuthenticated;
});

const isBusinessUser = computed(() => {
  return Boolean(user.value?.company_id);
});

function openWorkspace() {
  if (!isAuthenticated.value) {
    router.push('/login');
    return;
  }

  router.push(isBusinessUser.value ? '/erp/dashboard' : '/store/purchases');
}
</script>

<template>
  <header class="sticky top-0 z-40 border-b border-border/60 bg-background/85 backdrop-blur-xl">
    <div class="mx-auto flex h-16 w-full max-w-[1500px] items-center px-4 sm:px-6 lg:px-8">
      <RouterLink to="/" class="flex items-center gap-2">
        <div
          class="nexora-gradient nexora-glow flex h-9 w-9 items-center justify-center rounded-xl bg-primary text-sm font-bold text-primary-foreground shadow-sm"
        >
          N
        </div>

        <span class="nexora-gradient-text text-xl font-semibold tracking-[-0.025em]">
          {{ brand.name }}
        </span>
      </RouterLink>

      <nav class="ml-10 hidden items-center gap-7 lg:flex">
        <RouterLink
          to="/store"
          class="text-sm font-medium text-muted-foreground transition hover:text-foreground"
        >
          Store
        </RouterLink>

        <RouterLink
          to="/store/catalog"
          class="text-sm font-medium text-muted-foreground transition hover:text-foreground"
        >
          Explore
        </RouterLink>

        <a
          href="#business"
          class="text-sm font-medium text-muted-foreground transition hover:text-foreground"
        >
          For businesses
        </a>

        <a
          href="#platform"
          class="text-sm font-medium text-muted-foreground transition hover:text-foreground"
        >
          Platform
        </a>
      </nav>

      <div class="ml-auto hidden items-center gap-2 sm:flex">
        <ThemeToggle />
        <template v-if="!isAuthenticated">
          <Button variant="ghost" @click="router.push('/login')"> Sign in </Button>

          <Button class="rounded-full px-5" @click="router.push('/register')">
            Create account
            <ArrowRight class="ml-2 h-4 w-4" />
          </Button>
        </template>

        <DropdownMenu v-if="isAuthenticated">
          <DropdownMenuTrigger as-child>
            <Button variant="outline" class="rounded-full">
              <Building2 v-if="isBusinessUser" class="mr-2 h-4 w-4" />

              <UserRound v-else class="mr-2 h-4 w-4" />

              {{ isBusinessUser ? 'Business account' : 'My account' }}
            </Button>
          </DropdownMenuTrigger>

          <DropdownMenuContent align="end" class="w-60">
            <DropdownMenuLabel> Nexora account </DropdownMenuLabel>

            <DropdownMenuSeparator />

            <DropdownMenuItem @click="openWorkspace">
              <Building2 v-if="isBusinessUser" class="mr-2 h-4 w-4" />

              <ShoppingBag v-else class="mr-2 h-4 w-4" />

              {{ isBusinessUser ? 'Open ERP workspace' : 'My purchases' }}
            </DropdownMenuItem>

            <DropdownMenuItem @click="router.push('/store')">
              <ShoppingBag class="mr-2 h-4 w-4" />
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

        <Button variant="ghost" size="icon" class="rounded-full" @click="router.push('/store')">
          <ShoppingBag class="h-5 w-5" />
        </Button>
      </div>

      <Button variant="ghost" size="icon" class="ml-auto rounded-full sm:hidden">
        <Menu class="h-5 w-5" />
      </Button>
    </div>
  </header>
</template>
