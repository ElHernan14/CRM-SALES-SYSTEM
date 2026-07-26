<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import { Building2, Minus, Plus, ShoppingCart, Loader2 } from 'lucide-vue-next';

import { Button } from '@/components/ui/button';

import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
} from '@/components/ui/sheet';

import { getProductImageUrl } from '@/shared/utils/assets';
import type { StoreProduct } from '../types/store-product.types';

const props = defineProps<{
  open: boolean;
  product: StoreProduct | null;
  adding?: boolean;
}>();

const emit = defineEmits<{
  'update:open': [value: boolean];
  add: [
    payload: {
      product: StoreProduct;
      quantity: number;
    },
  ];
}>();

const quantity = ref(1);

watch(
  () => props.product,
  () => {
    quantity.value = 1;
  }
);

const canDecrease = computed(() => quantity.value > 1);

const canIncrease = computed(() => {
  if (!props.product) return false;
  return quantity.value < props.product.available_stock;
});

const subtotal = computed(() => {
  if (!props.product) return 0;
  return props.product.price * quantity.value;
});

function formatCurrency(value: number) {
  return `$${value.toFixed(2)}`;
}

function handleAdd() {
  if (!props.product) return;

  emit('add', {
    product: props.product,
    quantity: quantity.value,
  });
}
</script>

<template>
  <Sheet :open="open" @update:open="emit('update:open', $event)">
    <SheetContent class="w-full overflow-y-auto p-0 sm:max-w-2xl">
      <div v-if="product" class="flex min-h-full flex-col">
        <div class="relative h-72 overflow-hidden border-b border-border bg-muted/30">
          <img
            v-if="getProductImageUrl(product.image_path)"
            :src="getProductImageUrl(product.image_path)!"
            :alt="product.name"
            class="h-full w-full object-cover"
          />

          <div v-else class="flex h-full items-center justify-center">
            <ShoppingCart class="h-12 w-12 text-muted-foreground/50" />
          </div>

          <span
            class="absolute left-5 top-5 rounded-full bg-primary/90 px-3 py-1 text-xs font-semibold text-primary-foreground shadow-sm"
          >
            {{ product.category }}
          </span>
        </div>

        <div class="flex flex-1 flex-col p-6">
          <SheetHeader class="text-left">
            <SheetTitle class="text-2xl">
              {{ product.name }}
            </SheetTitle>

            <SheetDescription class="flex items-center gap-2">
              <Building2 class="h-4 w-4" />
              {{ product.company_name }}
            </SheetDescription>
          </SheetHeader>

          <p class="mt-6 text-sm leading-7 text-muted-foreground">
            {{ product.description || 'No product description available.' }}
          </p>

          <div class="mt-6 grid gap-3 sm:grid-cols-2">
            <div class="rounded-xl border border-border bg-muted/30 p-4">
              <p class="text-xs text-muted-foreground">Unit price</p>

              <p class="mt-2 text-lg font-semibold">
                {{ formatCurrency(product.price) }}
              </p>
            </div>

            <div class="rounded-xl border border-border bg-muted/30 p-4">
              <p class="text-xs text-muted-foreground">Available</p>

              <p class="mt-2 text-lg font-semibold">
                {{ product.available_stock }}
              </p>
            </div>

            <div class="rounded-xl border border-border bg-muted/30 p-4">
              <p class="text-xs text-muted-foreground">Resource kind</p>

              <p class="mt-2 text-lg font-semibold capitalize">
                {{ product.kind }}
              </p>
            </div>

            <div class="rounded-xl border border-border bg-muted/30 p-4">
              <p class="text-xs text-muted-foreground">Product type</p>

              <p class="mt-2 text-lg font-semibold">
                {{ product.type }}
              </p>
            </div>
          </div>

          <div class="mt-8 rounded-2xl border border-border bg-card p-5">
            <div class="flex items-center justify-between gap-4">
              <div>
                <p class="text-sm font-medium text-foreground">Purchase quantity</p>

                <p class="mt-1 text-xs text-muted-foreground">
                  Choose up to {{ product.available_stock }} units.
                </p>
              </div>

              <div class="flex items-center rounded-full border border-border bg-muted/40 p-1">
                <Button
                  variant="ghost"
                  size="icon"
                  class="h-8 w-8 rounded-full"
                  :disabled="!canDecrease || adding"
                  @click="quantity--"
                >
                  <Minus class="h-4 w-4" />
                </Button>

                <span class="min-w-10 text-center text-sm font-semibold">
                  {{ quantity }}
                </span>

                <Button
                  variant="ghost"
                  size="icon"
                  class="h-8 w-8 rounded-full"
                  :disabled="!canIncrease || adding"
                  @click="quantity++"
                >
                  <Plus class="h-4 w-4" />
                </Button>
              </div>
            </div>

            <div
              class="mt-5 flex flex-col gap-4 border-t border-border pt-5 min-[390px]:flex-row min-[390px]:items-center min-[390px]:justify-between"
            >
              <div>
                <p class="text-xs text-muted-foreground">Estimated subtotal</p>

                <p class="mt-1 text-xl font-semibold">
                  {{ formatCurrency(subtotal) }}
                </p>
              </div>

              <Button
                :disabled="product.available_stock <= 0 || adding"
                class="w-full min-[390px]:w-auto"
                @click="handleAdd"
              >
                <Loader2 v-if="adding" class="mr-2 h-4 w-4 animate-spin" />

                <ShoppingCart v-else class="mr-2 h-4 w-4" />

                {{ adding ? 'Adding...' : 'Add to purchase' }}
              </Button>
            </div>
          </div>
        </div>
      </div>
    </SheetContent>
  </Sheet>
</template>
