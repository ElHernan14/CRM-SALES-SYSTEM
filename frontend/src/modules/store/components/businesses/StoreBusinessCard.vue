<script setup lang="ts">
import { ArrowRight, BadgeCheck, Building2, PackageSearch } from 'lucide-vue-next';

import { Button } from '@/components/ui/button';

import { getCompanyCoverUrl, getCompanyLogoUrl } from '@/shared/utils/assets';

import type { StoreBusiness } from '../../types/store-business.types';

defineProps<{
  business: StoreBusiness;
}>();

const emit = defineEmits<{
  view: [business: StoreBusiness];
  catalog: [business: StoreBusiness];
}>();
</script>

<template>
  <article
    class="nexora-interactive nexora-shine nexora-surface group flex h-full min-w-0 flex-col overflow-hidden rounded-[1.75rem] border border-border bg-card shadow-sm hover:border-primary/25 hover:shadow-xl"
  >
    <!-- BRAND VISUAL -->
    <button
      type="button"
      class="relative block h-44 w-full shrink-0 overflow-hidden text-left sm:h-52"
      @click="emit('view', business)"
    >
      <img
        v-if="getCompanyCoverUrl(business.cover_image)"
        :src="getCompanyCoverUrl(business.cover_image)!"
        :alt="`${business.name} cover`"
        class="h-full w-full object-cover transition duration-700 group-hover:scale-105"
        loading="lazy"
        decoding="async"
      />

      <div
        v-else
        class="absolute inset-0 bg-gradient-to-br from-primary/20 via-muted/50 to-background"
      />

      <div
        class="absolute inset-0 bg-gradient-to-t from-background/95 via-background/20 to-transparent"
      />

      <span
        class="absolute left-4 top-4 max-w-[calc(100%-2rem)] truncate rounded-full border border-white/10 bg-background/85 px-3 py-1.5 text-xs font-semibold shadow-sm backdrop-blur"
      >
        {{ business.category }}
      </span>

      <!-- BUSINESS IDENTITY -->
      <div class="absolute bottom-4 left-4 right-4 flex min-w-0 items-end gap-3">
        <div
          class="flex h-14 w-14 shrink-0 items-center justify-center overflow-hidden rounded-2xl border-4 border-background bg-background shadow-lg sm:h-16 sm:w-16"
        >
          <img
            v-if="getCompanyLogoUrl(business.logo)"
            :src="getCompanyLogoUrl(business.logo)!"
            :alt="`${business.name} logo`"
            class="h-full w-full object-contain p-1.5"
            loading="lazy"
            decoding="async"
          />

          <Building2 v-else class="h-6 w-6 text-primary" />
        </div>

        <div class="min-w-0 flex-1 pb-1">
          <div class="flex min-w-0 items-center gap-2">
            <h2 class="min-w-0 truncate text-base font-semibold tracking-tight sm:text-lg">
              {{ business.name }}
            </h2>

            <BadgeCheck class="h-4 w-4 shrink-0 text-primary" />
          </div>

          <p class="mt-1 truncate text-xs text-muted-foreground">Business connected to Nexora</p>
        </div>
      </div>
    </button>

    <!-- INFORMATION -->
    <div class="flex min-w-0 flex-1 flex-col p-5">
      <p class="line-clamp-3 text-sm leading-6 text-muted-foreground sm:min-h-[4.5rem]">
        {{ business.description || 'Explore products and services from this business.' }}
      </p>

      <div
        class="mt-5 flex min-w-0 items-center gap-3 rounded-2xl border border-border bg-muted/20 p-4"
      >
        <div class="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl bg-primary/10">
          <PackageSearch class="h-5 w-5 text-primary" />
        </div>

        <div class="min-w-0">
          <p class="truncate text-lg font-semibold">
            {{ business.total_products }}
          </p>

          <p class="truncate text-xs text-muted-foreground">
            {{ business.total_products === 1 ? 'available resource' : 'available resources' }}
          </p>
        </div>
      </div>

      <div class="mt-auto grid gap-2 border-t border-border pt-5 xl:grid-cols-2">
        <Button
          type="button"
          variant="outline"
          class="group w-full rounded-full"
          @click="emit('view', business)"
        >
          Business details
        </Button>

        <Button
          type="button"
          class="w-full rounded-full"
          :disabled="business.total_products <= 0"
          @click="emit('catalog', business)"
        >
          View catalog

          <ArrowRight class="ml-2 h-4 w-4 transition-transform duration-200 group-hover:translate-x-0.5" />
        </Button>
      </div>
    </div>
  </article>
</template>
