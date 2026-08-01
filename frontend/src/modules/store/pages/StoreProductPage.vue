<script setup lang="ts">
import { computed, ref } from 'vue';

import { useRoute, useRouter } from 'vue-router';
import { storeToRefs } from 'pinia';

import {
  ArrowLeft,
  ArrowRight,
  BadgeCheck,
  Building2,
  Check,
  CircleOff,
  ImageIcon,
  Layers3,
  Loader2,
  Minus,
  PackageCheck,
  Plus,
  ShieldCheck,
  ShoppingBag,
  Store,
} from 'lucide-vue-next';

import { Button } from '@/components/ui/button';

import EmptyState from '@/shared/components/erp/EmptyState.vue';

import { getProductImageUrl } from '@/shared/utils/assets';

import { useStoreProduct } from '../composables/useStoreProduct';
import { useAddToStoreCart } from '../composables/useAddToStoreCart';
import { useAuthStore } from '@/modules/auth/stores/auth.store';

const route = useRoute();
const router = useRouter();
const auth = useAuthStore();
const { user } = storeToRefs(auth);

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

const isService = computed(() => {
  return product.value?.kind === 'service';
});

const isOwnCompanyProduct = computed(() => {
  return Boolean(user.value?.company_id && product.value?.company_id === user.value.company_id);
});

const isProductAvailable = computed(() => {
  return product.value?.is_available ?? true;
});

const unavailableReason = computed(() => {
  return product.value?.unavailable_reason ?? '';
});

const unavailableTitle = computed(() => {
  switch (unavailableReason.value) {
    case 'out_of_stock':
      return 'This product is currently out of stock';
    case 'seller_unavailable':
      return 'This seller is currently unavailable';
    case 'product_unavailable':
      return 'This product is no longer available';
    default:
      return 'This product is not available for purchase';
  }
});

const unavailableDescription = computed(() => {
  switch (unavailableReason.value) {
    case 'out_of_stock':
      return 'You can still review the product details, but new purchases are paused until stock is replenished.';
    case 'seller_unavailable':
      return 'The seller is not accepting new orders through the Store right now.';
    case 'product_unavailable':
      return 'This listing was removed from the active catalog. Previous orders remain available in your purchase history.';
    default:
      return 'You can review this historical product page, but it cannot be added to a cart.';
  }
});

const canDecrease = computed(() => {
  return quantity.value > 1 && !adding.value;
});

const canIncrease = computed(() => {
  if (!isProductAvailable.value) {
    return false;
  }

  if (adding.value) {
    return false;
  }

  if (isService.value) {
    return quantity.value < 999;
  }

  return quantity.value < availableStock.value;
});

const totalPrice = computed(() => {
  if (!product.value) return 0;

  return product.value.price * quantity.value;
});

const availabilityLabel = computed(() => {
  if (!isProductAvailable.value) {
    if (unavailableReason.value === 'out_of_stock') {
      return 'Out of stock';
    }

    return 'No longer available';
  }

  if (isService.value) {
    return 'Available to request';
  }

  if (availableStock.value <= 0) {
    return 'Currently unavailable';
  }

  if (availableStock.value <= 10) {
    return 'Limited availability';
  }

  return 'Available now';
});

const availabilityDescription = computed(() => {
  if (!isProductAvailable.value) {
    return unavailableDescription.value;
  }

  if (isService.value) {
    return 'This service can be included in your Nexora purchase.';
  }

  if (availableStock.value <= 0) {
    return 'There are no units currently available for purchase.';
  }

  return `${availableStock.value} units ready to purchase`;
});

const availabilityClass = computed(() => {
  if (!isProductAvailable.value) {
    return 'text-destructive';
  }

  if (isService.value) {
    return 'text-primary';
  }

  if (availableStock.value <= 0) {
    return 'text-destructive';
  }

  if (availableStock.value <= 10) {
    return 'text-amber-600 dark:text-amber-400';
  }

  return 'text-emerald-600 dark:text-emerald-400';
});

const canAddToCart = computed(() => {
  if (!isProductAvailable.value) {
    return false;
  }

  if (isOwnCompanyProduct.value) {
    return false;
  }

  if (adding.value) {
    return false;
  }

  if (isService.value) {
    return true;
  }

  return availableStock.value > 0;
});

const addButtonLabel = computed(() => {
  if (adding.value) {
    return 'Adding to cart...';
  }

  if (isOwnCompanyProduct.value) {
    return 'Your product';
  }

  if (!isProductAvailable.value) {
    return 'Unavailable for purchase';
  }

  return isService.value ? 'Add service to cart' : 'Add to cart';
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
  if (!product.value || !canAddToCart.value) return;

  await addToCart(product.value, quantity.value, true);
}

function openSellerCatalog() {
  if (!product.value) return;

  router.push({
    name: 'store-catalog',
    query: {
      company_id: String(product.value.company_id),
    },
  });
}

function openCategory() {
  if (!product.value) return;

  router.push({
    name: 'store-catalog',
    query: {
      category_id: String(product.value.category_id),
    },
  });
}
</script>

<template>
  <div class="pb-20">
    <!-- TOP NAVIGATION -->
    <div
      class="mx-auto flex w-full max-w-[1500px] items-center justify-between gap-4 px-5 py-6 sm:px-6 lg:px-8"
    >
      <Button type="button" variant="ghost" class="-ml-3 rounded-full" @click="router.back()">
        <ArrowLeft class="mr-2 h-4 w-4" />
        Back to catalog
      </Button>

      <div
        v-if="isRefreshing"
        class="hidden items-center gap-2 rounded-full border border-border bg-card px-3 py-1.5 text-xs text-muted-foreground sm:flex"
      >
        <Loader2 class="h-3.5 w-3.5 animate-spin" />
        Updating product
      </div>
    </div>

    <!-- LOADING -->
    <div
      v-if="isLoading"
      class="mx-auto grid w-full max-w-[1500px] gap-10 px-6 py-8 lg:grid-cols-2 lg:px-8"
    >
      <div class="aspect-square animate-pulse rounded-[2rem] bg-muted" />

      <div class="space-y-5 py-4">
        <div class="h-5 w-32 animate-pulse rounded bg-muted" />

        <div class="h-14 w-4/5 animate-pulse rounded bg-muted" />

        <div class="h-5 w-full animate-pulse rounded bg-muted" />

        <div class="h-5 w-3/4 animate-pulse rounded bg-muted" />

        <div class="h-24 animate-pulse rounded-2xl bg-muted" />

        <div class="h-56 animate-pulse rounded-[1.75rem] bg-muted" />
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
      class="mx-auto grid w-full max-w-[1500px] items-start gap-10 px-6 pb-16 lg:grid-cols-[1.05fr_0.95fr] lg:gap-16 lg:px-8"
    >
      <!-- PRODUCT VISUAL -->
      <section class="lg:sticky lg:top-24">
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

          <div
            class="absolute inset-x-0 bottom-0 h-40 bg-gradient-to-t from-background/70 to-transparent"
          />

          <button
            type="button"
            class="absolute left-5 top-5 max-w-[70%] truncate rounded-full border border-white/10 bg-background/85 px-3 py-1.5 text-xs font-semibold shadow-sm backdrop-blur transition hover:bg-background"
            @click="openCategory"
          >
            {{ product.category }}
          </button>

          <span
            class="absolute right-5 top-5 rounded-full px-3 py-1.5 text-xs font-semibold shadow-sm"
            :class="
              isProductAvailable
                ? 'bg-primary text-primary-foreground'
                : 'bg-destructive text-destructive-foreground'
            "
          >
            {{ isProductAvailable ? (isService ? 'Service' : 'Product') : 'Unavailable' }}
          </span>

          <div class="absolute bottom-5 left-5 right-5 flex items-end justify-between gap-4">
            <div
              class="inline-flex min-w-0 items-center gap-3 rounded-2xl border border-white/10 bg-background/85 px-4 py-3 shadow-md backdrop-blur"
            >
              <div
                class="flex h-9 w-9 shrink-0 items-center justify-center rounded-xl bg-primary/10"
              >
                <Building2 class="h-4 w-4 text-primary" />
              </div>

              <div class="min-w-0">
                <p class="truncate text-xs text-muted-foreground">Sold by</p>

                <p class="truncate text-sm font-semibold">
                  {{ product.company_name }}
                </p>
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- PRODUCT INFORMATION -->
      <section class="flex flex-col py-2 lg:py-5">
        <!-- SELLER -->
        <button
          type="button"
          class="group flex w-full items-center justify-between gap-4 rounded-2xl border border-border bg-card p-4 text-left shadow-sm transition hover:border-primary/25 hover:shadow-md"
          @click="openSellerCatalog"
        >
          <div class="flex min-w-0 items-center gap-3">
            <div
              class="flex h-11 w-11 shrink-0 items-center justify-center rounded-xl bg-primary/10"
            >
              <Store class="h-5 w-5 text-primary" />
            </div>

            <div class="min-w-0">
              <div class="flex items-center gap-2">
                <p class="truncate text-sm font-semibold">
                  {{ product.company_name }}
                </p>

                <BadgeCheck class="h-4 w-4 shrink-0 text-primary" />
              </div>

              <p class="mt-1 text-xs text-muted-foreground">Business connected to Nexora</p>
            </div>
          </div>

          <ArrowRight
            class="h-4 w-4 shrink-0 text-muted-foreground transition group-hover:translate-x-0.5 group-hover:text-primary"
          />
        </button>

        <!-- IDENTITY -->
        <div class="mt-8">
          <div class="flex flex-wrap items-center gap-2">
            <button
              type="button"
              class="rounded-full border border-primary/20 bg-primary/10 px-3 py-1 text-xs font-semibold text-primary"
              @click="openCategory"
            >
              {{ product.category }}
            </button>

            <span
              class="inline-flex items-center gap-1.5 rounded-full border border-border bg-muted/40 px-3 py-1 text-xs font-medium text-muted-foreground"
            >
              <Layers3 class="h-3.5 w-3.5" />
              {{ product.type }}
            </span>
          </div>

          <h1 class="mt-5 text-4xl font-semibold tracking-[-0.05em] text-foreground sm:text-5xl">
            {{ product.name }}
          </h1>

          <p class="mt-5 max-w-xl text-base leading-7 text-muted-foreground">
            {{
              product.description ||
              'Available from a company operating throughout the Nexora commerce network.'
            }}
          </p>

          <div
            v-if="!isProductAvailable"
            class="mt-6 rounded-[1.25rem] border border-destructive/25 bg-destructive/10 p-4"
          >
            <div class="flex gap-3">
              <div
                class="mt-0.5 flex h-9 w-9 shrink-0 items-center justify-center rounded-full bg-destructive/10"
              >
                <CircleOff class="h-4 w-4 text-destructive" />
              </div>

              <div>
                <p class="text-sm font-semibold text-foreground">
                  {{ unavailableTitle }}
                </p>

                <p class="mt-1 text-sm leading-6 text-muted-foreground">
                  {{ unavailableDescription }}
                </p>
              </div>
            </div>
          </div>
        </div>

        <!-- PRICE -->
        <div class="mt-8">
          <p class="text-xs font-medium uppercase tracking-[0.12em] text-muted-foreground">
            {{ isProductAvailable ? 'Current price' : 'Last listed price' }}
          </p>

          <p class="mt-2 text-4xl font-semibold tracking-[-0.04em] text-foreground">
            {{ formatCurrency(product.price) }}
          </p>

          <p class="mt-2 text-sm text-muted-foreground">
            <template v-if="isProductAvailable">
              Price per
              {{ isService ? 'service unit' : 'unit' }}.
            </template>

            <template v-else>
              Historical reference from the last active listing.
            </template>
          </p>
        </div>

        <!-- COMMERCE CONFIDENCE -->
        <div class="mt-8 overflow-hidden rounded-[1.5rem] border border-border bg-muted/15">
          <div class="grid divide-y divide-border sm:grid-cols-3 sm:divide-x sm:divide-y-0">
            <div class="flex items-start gap-3 p-4">
              <PackageCheck class="mt-0.5 h-5 w-5 shrink-0 text-primary" />

              <div>
                <p class="text-sm font-semibold">
                  {{ availabilityLabel }}
                </p>

                <p class="mt-1 text-xs leading-5 text-muted-foreground">
                  {{ availabilityDescription }}
                </p>
              </div>
            </div>

            <div class="flex items-start gap-3 p-4">
              <ShieldCheck class="mt-0.5 h-5 w-5 shrink-0 text-primary" />

              <div>
                <p class="text-sm font-semibold">Verified business</p>

                <p class="mt-1 text-xs leading-5 text-muted-foreground">
                  Seller identity connected to Nexora ERP.
                </p>
              </div>
            </div>

            <div class="flex items-start gap-3 p-4">
              <ShoppingBag class="mt-0.5 h-5 w-5 shrink-0 text-primary" />

              <div>
                <p class="text-sm font-semibold">Unified checkout</p>

                <p class="mt-1 text-xs leading-5 text-muted-foreground">
                  Orders remain organized by seller.
                </p>
              </div>
            </div>
          </div>
        </div>

        <!-- BUY BOX -->
        <div class="mt-8 overflow-hidden rounded-[1.75rem] border border-border bg-card shadow-lg">
          <div class="border-b border-border bg-muted/15 px-5 py-4">
            <div class="flex items-start justify-between gap-4">
              <div>
                <p class="text-sm font-semibold">Build your order</p>

                <p class="mt-1 text-xs leading-5 text-muted-foreground">
                  <template v-if="isOwnCompanyProduct">
                    Preview how this product appears to buyers. Self-purchase is disabled.
                  </template>

                  <template v-else-if="!isProductAvailable">
                    This product remains visible for reference, but it is not accepting new orders.
                  </template>

                  <template v-else>
                    Select the quantity before adding this
                    {{ isService ? 'service' : 'product' }}
                    to your cart.
                  </template>
                </p>
              </div>

              <div class="text-right">
                <p class="text-xs text-muted-foreground">Availability</p>

                <p class="mt-1 text-sm font-semibold" :class="availabilityClass">
                  {{ availabilityLabel }}
                </p>
              </div>
            </div>
          </div>

          <div class="p-5">
            <div class="flex flex-col gap-6 sm:flex-row sm:items-end sm:justify-between">
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
                <p class="text-xs font-medium uppercase tracking-wide text-muted-foreground">
                  Order subtotal
                </p>

                <p class="mt-1 text-3xl font-semibold tracking-[-0.035em]">
                  {{ formatCurrency(totalPrice) }}
                </p>
              </div>
            </div>

            <Button
              type="button"
              size="lg"
              class="mt-6 w-full rounded-full"
              :disabled="!canAddToCart"
              @click="handleAddToCart"
            >
              <span class="mr-2 flex h-4 w-4 items-center justify-center">
                <Loader2 v-if="adding" class="h-4 w-4 animate-spin" />

                <ShoppingBag v-else class="h-4 w-4" />
              </span>

              {{ addButtonLabel }}
            </Button>

            <div class="mt-4 flex items-center justify-center gap-2 text-xs text-muted-foreground">
              <Check class="h-3.5 w-3.5 text-primary" />

              <span v-if="isOwnCompanyProduct">
                Published by your business. Purchasing is unavailable.
              </span>

              <span v-else-if="!isProductAvailable">
                Historical product page. New cart actions are unavailable.
              </span>

              <span v-else>
                Your order will be connected to
                {{ product.company_name }}.
              </span>
            </div>
          </div>
        </div>
      </section>
    </main>
  </div>
</template>
