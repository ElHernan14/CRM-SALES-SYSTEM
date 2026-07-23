<script setup lang="ts">
import { computed, ref, toRef } from 'vue';

import { useRoute, useRouter } from 'vue-router';

import {
  ArrowLeft,
  Building2,
  Check,
  Heart,
  ImageIcon,
  Loader2,
  Minus,
  PackageCheck,
  Plus,
  ShieldCheck,
  ShoppingBag,
  Sparkles,
} from 'lucide-vue-next';

import { toast } from 'vue-sonner';

import { Button } from '@/components/ui/button';

import EmptyState from '@/shared/components/erp/EmptyState.vue';

import { getProductImageUrl } from '@/shared/utils/assets';

import { useStoreProduct } from '../composables/useStoreProduct';
import { useAddToStoreCart } from '../composables/useAddToStoreCart';

const route = useRoute();
const router = useRouter();

const productId = computed<number | null>(() => {
  const id = Number(route.params.productId);

  return Number.isFinite(id) && id > 0 ? id : null;
});

const { data: product, isLoading, isFetching, isError, refetch } = useStoreProduct(productId);

const quantity = ref(1);

const { addToCart, addingProductId } = useAddToStoreCart();

const adding = computed(() => {
  return product.value !== undefined && addingProductId.value === product.value.id;
});

const isRefreshing = computed(() => {
  return isFetching.value && !isLoading.value;
});

const availableStock = computed(() => {
  return product.value?.available_stock ?? 0;
});

const canDecrease = computed(() => {
  return quantity.value > 1 && !adding.value;
});

const canIncrease = computed(() => {
  return !adding.value && quantity.value < availableStock.value;
});

const totalPrice = computed(() => {
  if (!product.value) return 0;

  return product.value.price * quantity.value;
});

function formatCurrency(value: number) {
  return new Intl.NumberFormat('es-AR', {
    style: 'currency',
    currency: 'ARS',
    minimumFractionDigits: 2,
  }).format(value);
}

function decreaseQuantity() {
  if (!canDecrease.value) return;

  quantity.value--;
}

function increaseQuantity() {
  if (!canIncrease.value) return;

  quantity.value++;
}

async function handleAddToCart() {
  if (!product.value) return;

  await addToCart(product.value, quantity.value, true);
}

function saveProduct() {
  if (!product.value) return;

  toast.success('Saved for later', {
    description: product.value.name,
  });
}
</script>

<template>
  <div class="pb-20">
    <div class="mx-auto max-w-[1500px] px-5 py-6 sm:px-6 lg:px-8">
      <Button type="button" variant="ghost" class="-ml-3 rounded-full" @click="router.back()">
        <ArrowLeft class="mr-2 h-4 w-4" />
        Back to catalog
      </Button>
    </div>

    <!-- LOADING -->
    <div
      v-if="isLoading"
      class="mx-auto grid max-w-[1500px] gap-10 px-6 py-8 lg:grid-cols-2 lg:px-8"
    >
      <div class="aspect-square animate-pulse rounded-[2rem] bg-muted" />

      <div class="space-y-5 py-4">
        <div class="h-5 w-32 animate-pulse rounded bg-muted" />
        <div class="h-14 w-4/5 animate-pulse rounded bg-muted" />
        <div class="h-5 w-full animate-pulse rounded bg-muted" />
        <div class="h-5 w-3/4 animate-pulse rounded bg-muted" />
        <div class="h-28 animate-pulse rounded-2xl bg-muted" />
        <div class="h-14 animate-pulse rounded-full bg-muted" />
      </div>
    </div>

    <!-- ERROR -->
    <div v-else-if="isError" class="mx-auto max-w-3xl px-6 py-20">
      <EmptyState
        title="Product unavailable"
        description="This product could not be found or is no longer available for purchase."
        :icon="ShoppingBag"
      />

      <div class="mt-5 flex justify-center gap-3">
        <Button variant="outline" @click="refetch()"> Try again </Button>

        <Button @click="router.push('/store/catalog')"> Browse catalog </Button>
      </div>
    </div>

    <!-- PRODUCT -->
    <main
      v-else-if="product"
      class="mx-auto grid max-w-[1500px] gap-10 px-6 pb-16 lg:grid-cols-[1.05fr_0.95fr] lg:gap-16 lg:px-8"
    >
      <!-- IMAGE -->
      <section>
        <div
          class="group relative aspect-square overflow-hidden rounded-[2rem] border border-border bg-muted/30 shadow-sm"
        >
          <img
            v-if="getProductImageUrl(product.image_path)"
            :src="getProductImageUrl(product.image_path)!"
            :alt="product.name"
            class="h-full w-full object-cover transition duration-700 group-hover:scale-[1.02]"
          />

          <div
            v-else
            class="flex h-full w-full items-center justify-center bg-gradient-to-br from-muted via-muted/50 to-background"
          >
            <ImageIcon class="h-16 w-16 text-muted-foreground/40" />
          </div>

          <div class="absolute left-5 top-5 flex flex-wrap gap-2">
            <span
              class="rounded-full border border-white/15 bg-background/85 px-3 py-1.5 text-xs font-semibold shadow-sm backdrop-blur"
            >
              {{ product.category }}
            </span>

            <span
              class="rounded-full bg-primary px-3 py-1.5 text-xs font-semibold text-primary-foreground shadow-sm"
            >
              {{ product.type }}
            </span>
          </div>

          <Button
            type="button"
            variant="secondary"
            size="icon"
            class="absolute right-5 top-5 rounded-full bg-background/85 shadow-md backdrop-blur"
            @click="saveProduct"
          >
            <Heart class="h-5 w-5" />
          </Button>
        </div>
      </section>

      <!-- INFORMATION -->
      <section class="flex flex-col py-2 lg:py-8">
        <div>
          <RouterLink
            :to="{
              name: 'store-catalog',
              query: {
                company_id: product.company_id,
              },
            }"
            class="inline-flex items-center gap-2 text-sm font-medium text-muted-foreground transition hover:text-foreground"
          >
            <Building2 class="h-4 w-4" />
            {{ product.company_name }}
          </RouterLink>

          <h1 class="mt-5 text-4xl font-semibold tracking-[-0.045em] text-foreground sm:text-5xl">
            {{ product.name }}
          </h1>

          <p class="mt-5 max-w-xl text-base leading-7 text-muted-foreground">
            {{ product.description || 'Discover this product from a verified Nexora store.' }}
          </p>
        </div>

        <div class="mt-8">
          <p class="text-sm text-muted-foreground">Current price</p>

          <p class="mt-2 text-4xl font-semibold tracking-[-0.035em] text-foreground">
            {{ formatCurrency(product.price) }}
          </p>
        </div>

        <!-- TRUST -->
        <div class="mt-8 grid gap-3 sm:grid-cols-3">
          <div class="rounded-2xl border border-border bg-card p-4">
            <PackageCheck class="h-5 w-5 text-primary" />

            <p class="mt-3 text-sm font-semibold">Real availability</p>

            <p class="mt-1 text-xs text-muted-foreground">{{ availableStock }} units ready</p>
          </div>

          <div class="rounded-2xl border border-border bg-card p-4">
            <ShieldCheck class="h-5 w-5 text-primary" />

            <p class="mt-3 text-sm font-semibold">Verified seller</p>

            <p class="mt-1 text-xs text-muted-foreground">Connected business</p>
          </div>

          <div class="rounded-2xl border border-border bg-card p-4">
            <Sparkles class="h-5 w-5 text-primary" />

            <p class="mt-3 text-sm font-semibold">Secure checkout</p>

            <p class="mt-1 text-xs text-muted-foreground">Protected purchase flow</p>
          </div>
        </div>

        <!-- BUY BOX -->
        <div class="mt-8 rounded-[1.75rem] border border-border bg-card p-5 shadow-lg">
          <div class="flex flex-col gap-5 sm:flex-row sm:items-end sm:justify-between">
            <div>
              <p class="text-sm font-medium text-foreground">Quantity</p>

              <div
                class="mt-3 inline-flex items-center rounded-full border border-border bg-muted/30 p-1"
              >
                <Button
                  type="button"
                  variant="ghost"
                  size="icon"
                  class="h-9 w-9 rounded-full"
                  :disabled="!canDecrease"
                  @click="decreaseQuantity"
                >
                  <Minus class="h-4 w-4" />
                </Button>

                <span class="min-w-12 text-center text-sm font-semibold">
                  {{ quantity }}
                </span>

                <Button
                  type="button"
                  variant="ghost"
                  size="icon"
                  class="h-9 w-9 rounded-full"
                  :disabled="!canIncrease"
                  @click="increaseQuantity"
                >
                  <Plus class="h-4 w-4" />
                </Button>
              </div>
            </div>

            <div class="text-left sm:text-right">
              <p class="text-xs text-muted-foreground">Order subtotal</p>

              <p class="mt-1 text-2xl font-semibold">
                {{ formatCurrency(totalPrice) }}
              </p>
            </div>
          </div>

          <Button
            type="button"
            size="lg"
            class="mt-5 w-full rounded-full"
            :disabled="adding || availableStock <= 0"
            @click="handleAddToCart"
          >
            <Loader2 v-if="adding" class="mr-2 h-4 w-4 animate-spin" />

            <ShoppingBag v-else class="mr-2 h-4 w-4" />

            {{ adding ? 'Adding to cart...' : 'Add to cart' }}
          </Button>

          <div class="mt-4 flex items-center justify-center gap-2 text-xs text-muted-foreground">
            <Check class="h-3.5 w-3.5 text-primary" />
            Available from {{ product.company_name }}
          </div>
        </div>
      </section>
    </main>
  </div>
</template>
