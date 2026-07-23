<script setup lang="ts">
import { Building2, Check, ChevronsUpDown, Loader2, LogOut, ShoppingBag } from 'lucide-vue-next';

import { useRoute, useRouter } from 'vue-router';

import { computed } from 'vue';

import { useCompanyMe } from '@/modules/company/composables/useCompanyMe';

import { getCompanyLogoUrl } from '@/shared/utils/assets';

import { Button } from '@/components/ui/button';

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';

import { useLogout } from '@/modules/auth/composables/useLogout';

const route = useRoute();
const router = useRouter();

const { data: company, isLoading: companyLoading } = useCompanyMe();

const companyLogo = computed(() => {
  return getCompanyLogoUrl(company.value?.logo);
});

const { logout, isLoggingOut } = useLogout();

function openErp() {
  if (route.path.startsWith('/erp')) return;

  router.push('/erp/dashboard');
}

function openStore() {
  router.push('/store');
}
</script>

<template>
  <DropdownMenu>
    <DropdownMenuTrigger as-child>
      <div class="flex min-w-0 items-center gap-2.5">
        <div
          class="flex h-8 w-8 shrink-0 items-center justify-center overflow-hidden rounded-lg border border-border bg-muted/60"
        >
          <img
            v-if="companyLogo"
            :src="companyLogo"
            :alt="company?.name ?? 'Company logo'"
            class="h-full w-full object-contain p-1"
          />

          <Building2 v-else class="h-4 w-4 text-primary" />
        </div>

        <div class="hidden min-w-0 text-left sm:block">
          <p class="truncate text-sm font-semibold leading-none text-foreground">
            {{ companyLoading ? 'Loading workspace...' : (company?.name ?? 'Business workspace') }}
          </p>

          <p class="mt-1 truncate text-[11px] text-muted-foreground">
            {{ company?.category ?? 'Nexora ERP' }}
          </p>
        </div>
      </div>
    </DropdownMenuTrigger>

    <DropdownMenuContent align="start" class="w-72">
      <DropdownMenuLabel> Switch experience </DropdownMenuLabel>

      <DropdownMenuSeparator />

      <DropdownMenuItem class="gap-3 py-3" @click="openErp">
        <div
          class="flex h-9 w-9 shrink-0 items-center justify-center overflow-hidden rounded-xl bg-primary/10"
        >
          <img
            v-if="companyLogo"
            :src="companyLogo"
            :alt="company?.name ?? 'Company logo'"
            class="h-full w-full object-contain p-1.5"
          />

          <Building2 v-else class="h-4 w-4 text-primary" />
        </div>

        <div class="min-w-0 flex-1">
          <p class="truncate text-sm font-medium">
            {{ company?.name ?? 'Business workspace' }}
          </p>

          <p class="mt-1 truncate text-xs text-muted-foreground">
            {{ company?.category ?? 'Products, sales and operations' }}
          </p>
        </div>

        <Check v-if="route.path.startsWith('/erp')" class="h-4 w-4 shrink-0 text-primary" />
      </DropdownMenuItem>

      <DropdownMenuItem class="gap-3 py-3" @click="openStore">
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
        @click="logout"
      >
        <div class="flex h-9 w-9 shrink-0 items-center justify-center rounded-xl bg-destructive/10">
          <Loader2 v-if="isLoggingOut" class="h-4 w-4 animate-spin" />

          <LogOut v-else class="h-4 w-4" />
        </div>

        <div>
          <p class="text-sm font-medium">
            {{ isLoggingOut ? 'Signing out...' : 'Sign out' }}
          </p>

          <p class="mt-1 text-xs text-muted-foreground">End the current Nexora session</p>
        </div>
      </DropdownMenuItem>
    </DropdownMenuContent>
  </DropdownMenu>
</template>
