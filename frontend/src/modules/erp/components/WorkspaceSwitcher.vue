<script setup lang="ts">
import { computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';

import { Building2, Check, ChevronsUpDown, Loader2, LogOut, ShoppingBag } from 'lucide-vue-next';

import { Button } from '@/components/ui/button';

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';

import { useCompanyMe } from '@/modules/company/composables/useCompanyMe';
import { useLogout } from '@/modules/auth/composables/useLogout';

import { getCompanyLogoUrl } from '@/shared/utils/assets';

const route = useRoute();
const router = useRouter();

const { data: company, isLoading: companyLoading } = useCompanyMe();

const { logout, isLoggingOut } = useLogout();

const companyLogo = computed(() => {
  return getCompanyLogoUrl(company.value?.logo);
});

const companyName = computed(() => {
  if (companyLoading.value) {
    return 'Loading workspace...';
  }

  return company.value?.name ?? 'Business workspace';
});

const companyCategory = computed(() => {
  return company.value?.category ?? 'Nexora ERP';
});

const isErpActive = computed(() => {
  return route.path.startsWith('/erp');
});

function openErp() {
  if (isErpActive.value) return;

  router.push({
    name: 'erp-dashboard',
  });
}

function openStore() {
  router.push({
    name: 'store-home',
  });
}
</script>

<template>
  <DropdownMenu>
    <DropdownMenuTrigger as-child>
      <Button
        type="button"
        variant="ghost"
        class="h-10 min-w-0 max-w-[112px] justify-between gap-1.5 rounded-xl border border-transparent px-1.5 hover:border-border hover:bg-muted/60 min-[390px]:max-w-[150px] sm:max-w-[260px] sm:gap-2 sm:px-3"
        aria-label="Switch Nexora experience"
      >
        <div class="flex min-w-0 items-center gap-2 sm:gap-2.5">
          <div
            class="flex h-8 w-8 shrink-0 items-center justify-center overflow-hidden rounded-lg border border-border bg-muted/60"
          >
            <img
              v-if="companyLogo"
              :src="companyLogo"
              :alt="`${companyName} logo`"
              class="h-full w-full object-contain p-1"
            />

            <Building2 v-else class="h-4 w-4 text-primary" />
          </div>

          <div class="hidden min-w-0 text-left sm:block">
            <p class="truncate text-sm font-semibold leading-none text-foreground">
              {{ companyName }}
            </p>

            <p class="mt-1 truncate text-[11px] text-muted-foreground">
              {{ companyCategory }}
            </p>
          </div>
        </div>

        <ChevronsUpDown class="h-4 w-4 shrink-0 text-muted-foreground" />
      </Button>
    </DropdownMenuTrigger>

    <DropdownMenuContent align="start" :side-offset="8" class="w-[calc(100vw-2rem)] max-w-72">
      <DropdownMenuLabel>
        <div>
          <p class="text-sm font-semibold">Switch experience</p>

          <p class="mt-1 text-xs font-normal leading-5 text-muted-foreground">
            Move between your business workspace and Store.
          </p>
        </div>
      </DropdownMenuLabel>

      <DropdownMenuSeparator />

      <DropdownMenuItem class="gap-3 py-3" @select.prevent="openErp">
        <div
          class="flex h-9 w-9 shrink-0 items-center justify-center overflow-hidden rounded-xl bg-primary/10"
        >
          <img
            v-if="companyLogo"
            :src="companyLogo"
            :alt="`${companyName} logo`"
            class="h-full w-full object-contain p-1.5"
          />

          <Building2 v-else class="h-4 w-4 text-primary" />
        </div>

        <div class="min-w-0 flex-1">
          <p class="truncate text-sm font-medium">
            {{ companyName }}
          </p>

          <p class="mt-1 truncate text-xs text-muted-foreground">
            {{ companyCategory }}
          </p>
        </div>

        <Check v-if="isErpActive" class="h-4 w-4 shrink-0 text-primary" />
      </DropdownMenuItem>

      <DropdownMenuItem class="gap-3 py-3" @select.prevent="openStore">
        <div class="flex h-9 w-9 shrink-0 items-center justify-center rounded-xl bg-blue-500/10">
          <ShoppingBag class="h-4 w-4 text-blue-600 dark:text-blue-400" />
        </div>

        <div class="min-w-0 flex-1">
          <p class="text-sm font-medium">Nexora Store</p>

          <p class="mt-1 text-xs text-muted-foreground">Browse the consumer catalog</p>
        </div>
      </DropdownMenuItem>

      <DropdownMenuSeparator />

      <DropdownMenuItem
        class="gap-3 py-3 text-destructive focus:text-destructive"
        :disabled="isLoggingOut"
        @select.prevent="logout"
      >
        <div class="flex h-9 w-9 shrink-0 items-center justify-center rounded-xl bg-destructive/10">
          <Loader2 v-if="isLoggingOut" class="h-4 w-4 animate-spin" />

          <LogOut v-else class="h-4 w-4" />
        </div>

        <div class="min-w-0">
          <p class="text-sm font-medium">
            {{ isLoggingOut ? 'Signing out...' : 'Sign out' }}
          </p>

          <p class="mt-1 truncate text-xs text-muted-foreground">End the current Nexora session</p>
        </div>
      </DropdownMenuItem>
    </DropdownMenuContent>
  </DropdownMenu>
</template>
