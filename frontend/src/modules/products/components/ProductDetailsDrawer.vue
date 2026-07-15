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
    <SheetContent class="w-full sm:max-w-xl">
      <SheetHeader>
        <SheetTitle>
          {{ product?.name ?? 'Product details' }}
        </SheetTitle>

        <SheetDescription> Product catalog information and inventory snapshot. </SheetDescription>
      </SheetHeader>

      <div
        v-if="product"
        class="relative h-56 overflow-hidden rounded-2xl border border-border bg-muted/30"
      >
        <img
          v-if="getProductImageUrl(product.image_path)"
          :src="getProductImageUrl(product.image_path)!"
          :alt="product.name"
          class="h-full w-full object-cover"
        />

        <div
          v-else
          class="flex h-full w-full items-center justify-center bg-gradient-to-br from-muted to-background"
        >
          <ImageIcon class="h-10 w-10 text-muted-foreground/50" />
        </div>

        <div
          class="absolute inset-0 bg-gradient-to-t from-background/70 via-transparent to-transparent"
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

        <div class="grid gap-4 sm:grid-cols-2">
          <div class="rounded-xl border border-border bg-card p-4">
            <p class="text-xs text-muted-foreground">Type</p>

            <p class="mt-2 text-sm font-medium text-foreground">
              {{ product.type }}
            </p>
          </div>

          <div class="rounded-xl border border-border bg-card p-4">
            <p class="text-xs text-muted-foreground">Status</p>

            <p class="mt-2 text-sm font-medium text-foreground">
              {{ product.status === 1 ? 'Active' : 'Inactive' }}
            </p>
          </div>

          <div class="rounded-xl border border-border bg-card p-4">
            <p class="text-xs text-muted-foreground">Price</p>

            <p class="mt-2 text-sm font-medium text-foreground">${{ product.price.toFixed(2) }}</p>
          </div>

          <div class="rounded-xl border border-border bg-card p-4">
            <p class="text-xs text-muted-foreground">Stock</p>

            <p class="mt-2 text-sm font-medium text-foreground">
              {{ product.stock }}
            </p>
          </div>
        </div>

        <div class="rounded-xl border border-border bg-muted/40 p-4">
          <p class="text-xs text-muted-foreground">Tenant</p>

          <p class="mt-2 text-sm text-foreground">Company ID: {{ product.company_id }}</p>
        </div>
      </div>
    </SheetContent>
  </Sheet>
</template>
