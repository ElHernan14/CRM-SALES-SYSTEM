<script setup lang="ts">
import {
  ArrowUpRight,
  Building2,
  ImageIcon,
  Loader2,
  PackageCheck,
  ShoppingBag,
} from 'lucide-vue-next';

import { Button } from '@/components/ui/button';

import { getProductImageUrl } from '@/shared/utils/assets';

import type { StoreProduct } from '../../types/store-product.types';

defineProps<{
  product: StoreProduct;
  adding?: boolean;
  actionsDisabled?: boolean;
}>();

const emit = defineEmits<{
  view: [product: StoreProduct];
  add: [product: StoreProduct];
}>();

function formatCurrency(value: number) {
  return new Intl.NumberFormat('es-AR', {
    style: 'currency',
    currency: 'ARS',
    maximumFractionDigits: 0,
  }).format(value);
}

function getAvailabilityLabel(stock: number) {
  if (stock <= 0) {
    return 'Unavailable';
  }

  if (stock <= 10) {
    return 'Only a few left';
  }

  return 'Available';
}

function getAvailabilityClass(stock: number) {
  if (stock <= 0) {
    return 'text-destructive';
  }

  if (stock <= 10) {
    return 'text-amber-600 dark:text-amber-400';
  }

  return 'text-emerald-600 dark:text-emerald-400';
}

function getAvailabilityDotClass(stock: number) {
  if (stock <= 0) {
    return 'bg-destructive';
  }

  if (stock <= 10) {
    return 'bg-amber-500';
  }

  return 'bg-emerald-500';
}
</script>

<template>
  <article
    class="group relative flex h-full flex-col overflow-hidden rounded-[1.75rem] border border-border bg-card shadow-sm transition duration-300 hover:-translate-y-1 hover:border-primary/25 hover:shadow-xl"
  >
    <!-- PRODUCT IMAGE -->
    <button
      type="button"
      class="relative block h-64 overflow-hidden text-left"
      @click="emit('view', product)"
    >
      <img
        v-if="getProductImageUrl(product.image_path)"
        :src="getProductImageUrl(product.image_path)!"
        :alt="product.name"
        class="h-full w-full object-cover transition duration-700 group-hover:scale-105"
      />

      <div
        v-else
        class="flex h-full w-full items-center justify-center bg-gradient-to-br from-muted via-muted/50 to-background"
      >
        <ImageIcon class="h-12 w-12 text-muted-foreground/40" />
      </div>

      <div
        class="absolute inset-0 bg-gradient-to-t from-background/75 via-transparent to-transparent"
      />

      <!-- CATEGORY -->
      <span
        class="absolute left-4 top-4 max-w-[75%] truncate rounded-full border border-white/10 bg-background/85 px-3 py-1.5 text-xs font-semibold text-foreground shadow-sm backdrop-blur"
      >
        {{ product.category }}
      </span>

      <!-- VIEW INDICATOR -->
      <div
        class="absolute right-4 top-4 flex h-9 w-9 items-center justify-center rounded-full border border-white/10 bg-background/85 text-muted-foreground opacity-0 shadow-sm backdrop-blur transition group-hover:opacity-100"
      >
        <ArrowUpRight class="h-4 w-4" />
      </div>

      <!-- COMPANY OVER IMAGE -->
      <div class="absolute bottom-4 left-4 right-4">
        <p class="flex items-center gap-2 truncate text-xs font-medium text-muted-foreground">
          <Building2 class="h-3.5 w-3.5 shrink-0" />

          <span class="truncate">
            {{ product.company_name }}
          </span>
        </p>
      </div>
    </button>

    <!-- PRODUCT CONTENT -->
    <div class="flex flex-1 flex-col p-5">
      <div class="flex-1">
        <div class="flex items-start justify-between gap-4">
          <div class="min-w-0 flex-1">
            <button type="button" class="block text-left" @click="emit('view', product)">
              <h2
                class="line-clamp-2 min-h-14 text-xl font-semibold tracking-[-0.025em] text-foreground transition group-hover:text-primary"
              >
                {{ product.name }}
              </h2>
            </button>

            <p class="mt-1 truncate text-xs font-medium text-muted-foreground">
              {{ product.type }}
            </p>
          </div>
        </div>

        <p class="mt-4 line-clamp-2 min-h-10 text-sm leading-5 text-muted-foreground">
          {{
            product.description ||
            'Available from a business connected to the Nexora commerce network.'
          }}
        </p>
      </div>

      <!-- PRICE -->
      <div class="mt-6">
        <p class="text-xs font-medium text-muted-foreground">Current price</p>

        <p class="mt-1 text-2xl font-semibold tracking-[-0.035em] text-foreground">
          {{ formatCurrency(product.price) }}
        </p>
      </div>

      <!-- AVAILABILITY -->
      <div class="mt-5 flex items-center justify-between gap-4 border-t border-border pt-4">
        <div class="flex min-w-0 items-center gap-3">
          <div class="flex h-9 w-9 shrink-0 items-center justify-center rounded-xl bg-muted/50">
            <PackageCheck class="h-4 w-4" :class="getAvailabilityClass(product.available_stock)" />
          </div>

          <div class="min-w-0">
            <div class="flex items-center gap-2">
              <span
                class="h-2 w-2 shrink-0 rounded-full"
                :class="getAvailabilityDotClass(product.available_stock)"
              />

              <p
                class="truncate text-sm font-semibold"
                :class="getAvailabilityClass(product.available_stock)"
              >
                {{ getAvailabilityLabel(product.available_stock) }}
              </p>
            </div>

            <p class="mt-0.5 text-xs text-muted-foreground">
              {{
                product.available_stock > 0
                  ? `${product.available_stock} units in stock`
                  : 'Currently unavailable for purchase'
              }}
            </p>
          </div>
        </div>

        <Button
          type="button"
          class="shrink-0 rounded-full px-5"
          :disabled="product.available_stock <= 0 || actionsDisabled"
          @click="emit('add', product)"
        >
          <Loader2 v-if="adding" class="mr-2 h-4 w-4 animate-spin" />

          <ShoppingBag v-else class="mr-2 h-4 w-4" />

          {{ adding ? 'Adding...' : 'Add' }}
        </Button>
      </div>
    </div>
  </article>
</template>
