<script setup lang="ts">
import { computed } from 'vue';

import { ArrowRight, BadgeCheck, Building2, PackageSearch, Store } from 'lucide-vue-next';

import { Button } from '@/components/ui/button';

import { Sheet, SheetContent } from '@/components/ui/sheet';

import { getCompanyCoverUrl, getCompanyLogoUrl } from '@/shared/utils/assets';

import type { StoreBusiness } from '../../types/store-business.types';

const props = defineProps<{
  open: boolean;
  business: StoreBusiness | null;
}>();

const emit = defineEmits<{
  'update:open': [value: boolean];
  catalog: [business: StoreBusiness];
}>();

const companyCover = computed(() => {
  return getCompanyCoverUrl(props.business?.cover_image);
});

const companyLogo = computed(() => {
  return getCompanyLogoUrl(props.business?.logo);
});
</script>

<template>
  <Sheet :open="open" @update:open="emit('update:open', $event)">
    <SheetContent class="w-full overflow-y-auto p-0 sm:max-w-2xl">
      <template v-if="business">
        <!-- BUSINESS HERO -->
        <section class="relative min-h-72 overflow-hidden border-b border-border bg-muted/30">
          <img
            v-if="companyCover"
            :src="companyCover"
            :alt="`${business.name} cover`"
            class="absolute inset-0 h-full w-full object-cover"
          />

          <div
            v-else
            class="absolute inset-0 bg-gradient-to-br from-primary/20 via-muted/50 to-background"
          />

          <div
            class="absolute inset-0 bg-gradient-to-t from-background via-background/35 to-transparent"
          />

          <div class="relative flex min-h-72 flex-col justify-end p-6">
            <div
              class="flex h-20 w-20 items-center justify-center overflow-hidden rounded-[1.5rem] border-4 border-background bg-background shadow-xl"
            >
              <img
                v-if="companyLogo"
                :src="companyLogo"
                :alt="`${business.name} logo`"
                class="h-full w-full object-contain p-2"
              />

              <Building2 v-else class="h-7 w-7 text-primary" />
            </div>

            <div class="mt-5">
              <div class="flex flex-wrap items-center gap-2">
                <span
                  class="rounded-full border border-primary/20 bg-primary/10 px-3 py-1 text-xs font-semibold text-primary"
                >
                  {{ business.category }}
                </span>

                <span
                  class="inline-flex items-center gap-1.5 rounded-full border border-border bg-background/70 px-3 py-1 text-xs font-medium text-muted-foreground backdrop-blur"
                >
                  <BadgeCheck class="h-3.5 w-3.5 text-primary" />
                  Connected business
                </span>
              </div>

              <h2 class="mt-4 text-3xl font-semibold tracking-[-0.04em]">
                {{ business.name }}
              </h2>
            </div>
          </div>
        </section>

        <div class="space-y-6 p-6">
          <!-- DESCRIPTION -->
          <section>
            <p class="text-xs font-semibold uppercase tracking-[0.16em] text-primary">
              Business profile
            </p>

            <h3 class="mt-2 text-xl font-semibold">About this seller</h3>

            <p class="mt-3 text-sm leading-7 text-muted-foreground">
              {{
                business.description ||
                'This company publishes products and services through the Nexora commerce network.'
              }}
            </p>
          </section>

          <!-- COMMERCE PRESENCE -->
          <section class="rounded-[1.5rem] border border-border bg-muted/15 p-5">
            <div class="flex items-start gap-4">
              <div
                class="flex h-11 w-11 shrink-0 items-center justify-center rounded-xl bg-primary/10"
              >
                <Store class="h-5 w-5 text-primary" />
              </div>

              <div class="min-w-0">
                <p class="text-sm font-semibold">Nexora commerce presence</p>

                <p class="mt-2 text-sm leading-6 text-muted-foreground">
                  Inventory published by this company is connected to its active business workspace.
                </p>
              </div>
            </div>

            <div class="mt-5 flex items-center justify-between border-t border-border pt-4">
              <div>
                <p class="text-xs text-muted-foreground">Available catalog</p>

                <p class="mt-1 text-2xl font-semibold">
                  {{ business.total_products }}
                </p>
              </div>

              <PackageSearch class="h-6 w-6 text-primary" />
            </div>
          </section>

          <Button
            size="lg"
            class="w-full rounded-full"
            :disabled="business.total_products <= 0"
            @click="emit('catalog', business)"
          >
            Explore business catalog

            <ArrowRight class="ml-2 h-4 w-4" />
          </Button>
        </div>
      </template>
    </SheetContent>
  </Sheet>
</template>
