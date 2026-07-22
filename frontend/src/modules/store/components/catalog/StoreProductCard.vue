<script setup lang="ts">
import { Heart, ImageIcon, ShoppingBag, Sparkles } from 'lucide-vue-next';

import { Button } from '@/components/ui/button';

import { getProductImageUrl } from '@/shared/utils/assets';

import type { StoreProduct } from '../../types/store-product.types';

defineProps<{
  product: StoreProduct;
  adding?: boolean;
}>();

const emit = defineEmits<{
  view: [product: StoreProduct];
  add: [product: StoreProduct];
  favorite: [product: StoreProduct];
}>();

function formatCurrency(value: number) {
  return new Intl.NumberFormat('es-AR', {
    style: 'currency',
    currency: 'ARS',
    minimumFractionDigits: 2,
  }).format(value);
}
</script>

<template>
  <article
    class="group relative flex h-full flex-col overflow-hidden rounded-[1.75rem] border border-border bg-card shadow-sm transition duration-300 hover:-translate-y-1 hover:border-primary/25 hover:shadow-xl"
  >
    <button
      type="button"
      class="relative block h-64 overflow-hidden text-left"
      @click="emit('view', product)"
    >
      <img
        v-if="getProductImageUrl(product.image_path)"
        :src="getProductImageUrl(product.image_path)!"
        :alt="product.name"
        class="h-full w-full object-cover transition duration-500 group-hover:scale-105"
      />

      <div
        v-else
        class="flex h-full w-full items-center justify-center bg-gradient-to-br from-muted via-muted/50 to-background"
      >
        <ImageIcon class="h-12 w-12 text-muted-foreground/40" />
      </div>

      <div
        class="absolute inset-0 bg-gradient-to-t from-background/80 via-transparent to-transparent"
      />

      <div class="absolute left-4 top-4 flex max-w-[75%] flex-wrap gap-2">
        <span
          class="rounded-full border border-white/15 bg-background/85 px-3 py-1 text-xs font-semibold text-foreground shadow-sm backdrop-blur"
        >
          {{ product.category }}
        </span>

        <span
          class="rounded-full bg-primary/90 px-3 py-1 text-xs font-semibold text-primary-foreground shadow-sm"
        >
          {{ product.type }}
        </span>
      </div>

      <Button
        type="button"
        variant="secondary"
        size="icon"
        class="absolute right-4 top-4 rounded-full bg-background/85 shadow-sm backdrop-blur"
        @click.stop="emit('favorite', product)"
      >
        <Heart class="h-4 w-4" />
      </Button>

      <div class="absolute bottom-4 left-4 right-4 flex items-end justify-between gap-4">
        <div class="min-w-0">
          <p class="truncate text-xs font-medium text-muted-foreground">
            {{ product.company_name }}
          </p>

          <h2 class="mt-1 line-clamp-2 text-lg font-semibold tracking-tight text-foreground">
            {{ product.name }}
          </h2>
        </div>

        <div
          class="shrink-0 rounded-2xl border border-border bg-background/90 px-3 py-2 text-right shadow-md backdrop-blur"
        >
          <p class="text-[10px] font-medium uppercase tracking-wider text-muted-foreground">
            Price
          </p>

          <p class="mt-0.5 text-lg font-semibold tracking-tight text-foreground">
            {{ formatCurrency(product.price) }}
          </p>
        </div>
      </div>
    </button>

    <div class="flex flex-1 flex-col p-5">
      <p class="line-clamp-2 min-h-10 text-sm leading-5 text-muted-foreground">
        {{ product.description || 'Discover this product from a verified store.' }}
      </p>

      <div class="mt-5 flex items-center justify-between border-t border-border pt-4">
        <div class="flex items-center gap-2">
          <span
            class="h-2.5 w-2.5 rounded-full"
            :class="
              product.available_stock > 10
                ? 'bg-emerald-500'
                : product.available_stock > 0
                  ? 'bg-amber-500'
                  : 'bg-destructive'
            "
          />

          <div>
            <p class="text-xs font-medium text-foreground">
              {{ product.available_stock > 0 ? 'Available' : 'Unavailable' }}
            </p>

            <p class="text-[11px] text-muted-foreground">
              {{ product.available_stock }}
              in stock
            </p>
          </div>
        </div>

        <Button
          type="button"
          class="rounded-full px-5"
          :disabled="product.available_stock <= 0 || adding"
          @click="emit('add', product)"
        >
          <ShoppingBag v-if="!adding" class="mr-2 h-4 w-4" />

          <Sparkles v-else class="mr-2 h-4 w-4 animate-pulse" />

          {{ adding ? 'Adding...' : 'Add to cart' }}
        </Button>
      </div>
    </div>
  </article>
</template>
