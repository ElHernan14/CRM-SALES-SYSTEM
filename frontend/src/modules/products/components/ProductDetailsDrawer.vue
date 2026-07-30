<script setup lang="ts">
import { ImageIcon } from 'lucide-vue-next';
import { getProductImageUrl } from '@/shared/utils/assets';

import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
} from '@/components/ui/sheet';

import type { ProductListItem } from '../types/product.types';

defineProps<{
  open: boolean;
  product: ProductListItem | null;
}>();

const emit = defineEmits<{
  'update:open': [value: boolean];
}>();
</script>

<template>
  <Sheet :open="open" @update:open="emit('update:open', $event)">
    <SheetContent class="w-full overflow-y-auto sm:max-w-xl">
      <SheetHeader>
        <SheetTitle>
          {{ product?.name ?? 'Product details' }}
        </SheetTitle>

        <SheetDescription> Product catalog information and inventory snapshot. </SheetDescription>
      </SheetHeader>
      <div
        v-if="product"
        class="relative mt-3 h-56 overflow-hidden rounded-2xl border border-border bg-muted/30"
      >
        <img
          v-if="getProductImageUrl(product.image_path)"
          :src="getProductImageUrl(product.image_path)!"
          :alt="product.name"
          class="h-full w-full object-cover"
        />

        <div
          v-else
          class="flex h-full w-full items-center justify-center bg-linear-to-br from-muted to-background"
        >
          <ImageIcon class="h-10 w-10 text-muted-foreground/50" />
        </div>

        <div
          class="absolute inset-0 bg-linear-to-t from-background/70 via-transparent to-transparent"
        />

        <div class="absolute bottom-4 left-4 flex flex-wrap gap-2">
          <span
            class="rounded-full bg-primary/90 px-3 py-1 text-xs font-semibold text-primary-foreground shadow-sm"
          >
            {{ product.category }}
          </span>

          <span
            class="rounded-full bg-background/90 px-3 py-1 text-xs font-medium text-foreground shadow-sm"
          >
            {{ product.type }}
          </span>
        </div>
      </div>

      <div v-if="product" class="mt-6 space-y-6">
        <div class="rounded-xl border border-border bg-card p-4">
          <p class="text-xs text-muted-foreground">Description</p>

          <p class="mt-2 text-sm text-foreground">
            {{ product.description || 'No description provided.' }}
          </p>
        </div>

        <div class="grid gap-4 sm:grid-cols-3">
          <div class="rounded-xl border border-border bg-card p-4">
            <p class="text-xs text-muted-foreground">Kind</p>

            <p class="mt-2 text-sm font-medium capitalize text-foreground">
              {{ product.kind }}
            </p>
          </div>

          <div class="rounded-xl border border-border bg-card p-4">
            <p class="text-xs text-muted-foreground">Category</p>

            <p class="mt-2 text-sm font-medium text-foreground">
              {{ product.category }}
            </p>
          </div>

          <div class="rounded-xl border border-border bg-card p-4">
            <p class="text-xs text-muted-foreground">Type</p>

            <p class="mt-2 text-sm font-medium text-foreground">
              {{ product.type }}
            </p>
          </div>
        </div>

        <div class="grid gap-4 sm:grid-cols-2">
          <div class="rounded-xl border border-border bg-card p-4">
            <p class="text-xs text-muted-foreground">Price</p>

            <p class="mt-2 text-lg font-semibold text-foreground">${{ product.price.toFixed(2) }}</p>
          </div>

          <div class="rounded-xl border border-border bg-card p-4">
            <p class="text-xs text-muted-foreground">Status</p>

            <div class="mt-2">
              <span
                class="inline-flex items-center gap-2 rounded-full border px-2.5 py-1 text-xs font-semibold"
                :class="
                  product.status === 1
                    ? 'border-emerald-500/20 bg-emerald-500/10 text-emerald-600 dark:text-emerald-400'
                    : 'border-border bg-muted text-muted-foreground'
                "
              >
                <span
                  class="h-1.5 w-1.5 rounded-full"
                  :class="product.status === 1 ? 'bg-emerald-500' : 'bg-muted-foreground'"
                />

                {{ product.status === 1 ? 'Active' : 'Inactive' }}
              </span>
            </div>
          </div>
        </div>

        <div class="rounded-xl border border-border bg-card p-4">
          <div class="flex items-center justify-between gap-4">
            <div>
              <p class="text-xs text-muted-foreground">Inventory snapshot</p>

              <p class="mt-1 text-sm text-muted-foreground">
                Physical stock, units reserved by drafts, and quantity available to sell.
              </p>
            </div>
          </div>

          <div class="mt-4 grid gap-3 sm:grid-cols-3">
            <div class="rounded-lg border border-border bg-muted/30 p-3">
              <p class="text-[10px] font-medium uppercase tracking-wide text-muted-foreground">
                Stock
              </p>

              <p class="mt-1 text-xl font-semibold text-foreground">
                {{ product.stock }}
              </p>
            </div>

            <div class="rounded-lg border border-border bg-muted/30 p-3">
              <p class="text-[10px] font-medium uppercase tracking-wide text-muted-foreground">
                Reserved
              </p>

              <p class="mt-1 text-xl font-semibold text-amber-600 dark:text-amber-400">
                {{ product.reserved_stock ?? 0 }}
              </p>
            </div>

            <div class="rounded-lg border border-border bg-muted/30 p-3">
              <p class="text-[10px] font-medium uppercase tracking-wide text-muted-foreground">
                Available
              </p>

              <p
                class="mt-1 text-xl font-semibold"
                :class="
                  (product.available_stock ?? product.stock - (product.reserved_stock ?? 0)) > 0
                    ? 'text-emerald-600 dark:text-emerald-400'
                    : 'text-destructive'
                "
              >
                {{ product.available_stock ?? product.stock - (product.reserved_stock ?? 0) }}
              </p>
            </div>
          </div>
        </div>
      </div>
    </SheetContent>
  </Sheet>
</template>
