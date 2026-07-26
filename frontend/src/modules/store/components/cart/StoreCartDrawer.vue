<script setup lang="ts">
import { computed, ref, watch } from 'vue';

import { useRouter } from 'vue-router';
import { useQueryClient } from '@tanstack/vue-query';
import { toast } from 'vue-sonner';

import {
  ArrowRight,
  BadgeCheck,
  Building2,
  ImageIcon,
  Loader2,
  Minus,
  PackageOpen,
  Plus,
  RefreshCw,
  ShoppingBag,
  Trash2,
} from 'lucide-vue-next';

import { Button } from '@/components/ui/button';

import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
} from '@/components/ui/sheet';

import { getProductImageUrl } from '@/shared/utils/assets';

import { useStoreUiStore } from '../../stores/store-ui.store';
import { useAuthStore } from '@/modules/auth/stores/auth.store';

import { useStoreCarts } from '../../composables/useStoreCarts';
import { useUpdateStoreCartItem } from '../../composables/useUpdateStoreCartItem';
import { useDeleteStoreCartItem } from '../../composables/useDeleteStoreCartItem';

import type { StoreCartItem, StoreSellerCart } from '../../types/store-cart.types';

const router = useRouter();
const queryClient = useQueryClient();

const storeUi = useStoreUiStore();
const auth = useAuthStore();

watch(
  () => auth.isAuthenticated,
  (authenticated) => {
    if (!authenticated) {
      storeUi.closeCart();
    }
  }
);

const { data, isLoading, isFetching, isError, refetch } = useStoreCarts();

const updateMutation = useUpdateStoreCartItem();
const deleteMutation = useDeleteStoreCartItem();

const processingItemId = ref<number | null>(null);
const isSynchronizing = ref(false);

const carts = computed(() => {
  return data.value?.carts ?? [];
});

const summary = computed(() => {
  return (
    data.value?.summary ?? {
      seller_count: 0,
      item_count: 0,
      subtotal: 0,
      taxes: 0,
      total_amount: 0,
    }
  );
});

const hasItems = computed(() => {
  return summary.value.item_count > 0;
});

const isRefreshing = computed(() => {
  return isFetching.value && !isLoading.value;
});

const isChangingCart = computed(() => {
  return isSynchronizing.value || updateMutation.isPending.value || deleteMutation.isPending.value;
});

const isCartBusy = computed(() => {
  return isRefreshing.value || isChangingCart.value;
});

function formatCurrency(value: number) {
  return new Intl.NumberFormat('es-AR', {
    style: 'currency',
    currency: 'ARS',
    minimumFractionDigits: 2,
  }).format(value);
}

async function synchronizeStoreCart() {
  try {
    isSynchronizing.value = true;

    await Promise.all([
      refetch(),

      queryClient.refetchQueries({
        queryKey: ['b2c-store-products'],
        type: 'active',
      }),

      queryClient.invalidateQueries({
        queryKey: ['b2c-store-product'],
      }),
    ]);
  } finally {
    isSynchronizing.value = false;
  }
}

async function updateQuantity(cart: StoreSellerCart, item: StoreCartItem, quantity: number) {
  if (quantity < 1) return;

  try {
    processingItemId.value = item.id;

    await updateMutation.mutateAsync({
      invoiceId: cart.invoice_id,
      sellerCompanyId: cart.seller_company_id,
      itemId: item.id,
      quantity,
    });

    await synchronizeStoreCart();
  } catch (error: any) {
    const message =
      error?.response?.data?.errorMessage ??
      error?.response?.data?.message ??
      'Failed to update cart quantity';

    toast.error(message);
  } finally {
    processingItemId.value = null;
  }
}

async function deleteItem(cart: StoreSellerCart, item: StoreCartItem) {
  try {
    processingItemId.value = item.id;

    await deleteMutation.mutateAsync({
      invoiceId: cart.invoice_id,
      sellerCompanyId: cart.seller_company_id,
      itemId: item.id,
    });

    await synchronizeStoreCart();

    toast.success('Product removed from cart');
  } catch (error: any) {
    const message =
      error?.response?.data?.errorMessage ??
      error?.response?.data?.message ??
      'Failed to remove product';

    toast.error(message);
  } finally {
    processingItemId.value = null;
  }
}

function continueShopping() {
  storeUi.closeCart();

  router.push({
    name: 'store-catalog',
  });
}

function goToCheckout() {
  storeUi.closeCart();

  router.push({
    name: 'store-checkout',
  });
}

function openProduct(productId: number) {
  storeUi.closeCart();

  router.push({
    name: 'store-product',
    params: {
      productId,
    },
  });
}

function openSellerCatalog(sellerCompanyId: number) {
  storeUi.closeCart();

  router.push({
    name: 'store-catalog',
    query: {
      company_id: String(sellerCompanyId),
    },
  });
}
</script>

<template>
  <Sheet :open="storeUi.cartOpen" @update:open="$event ? storeUi.openCart() : storeUi.closeCart()">
    <SheetContent class="flex h-full w-full flex-col overflow-hidden p-0 sm:max-w-2xl">
      <!-- HEADER -->
      <div
        class="shrink-0 border-b border-border bg-gradient-to-r from-primary/5 via-muted/15 to-transparent px-4 py-5 sm:px-6"
      >
        <SheetHeader class="pr-8 text-left">
          <div class="flex min-w-0 items-start justify-between gap-3">
            <div class="flex min-w-0 items-center gap-3">
              <div
                class="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl bg-primary text-primary-foreground shadow-sm sm:h-11 sm:w-11"
              >
                <ShoppingBag class="h-5 w-5" />
              </div>

              <div class="min-w-0">
                <SheetTitle class="truncate"> Nexora cart </SheetTitle>

                <SheetDescription class="mt-1 line-clamp-2">
                  One checkout, organized into independent seller orders.
                </SheetDescription>
              </div>
            </div>

            <Button
              v-if="hasItems"
              type="button"
              variant="ghost"
              size="icon"
              class="mr-5 shrink-0 rounded-full sm:mr-3"
              :disabled="isCartBusy"
              aria-label="Synchronize cart"
              @click="synchronizeStoreCart"
            >
              <RefreshCw
                class="h-4 w-4"
                :class="{
                  'animate-spin': isCartBusy,
                }"
              />
            </Button>
          </div>
        </SheetHeader>
      </div>

      <!-- LOADING -->
      <div v-if="isLoading" class="flex min-h-0 flex-1 items-center justify-center p-6 sm:p-8">
        <div class="text-center">
          <Loader2 class="mx-auto h-7 w-7 animate-spin text-primary" />

          <p class="mt-3 text-sm text-muted-foreground">Loading your cart...</p>
        </div>
      </div>

      <!-- ERROR -->
      <div v-else-if="isError" class="flex min-h-0 flex-1 items-center justify-center p-6 sm:p-8">
        <div class="w-full max-w-sm text-center">
          <p class="text-sm font-semibold">Unable to load your cart</p>

          <p class="mt-2 text-sm leading-6 text-muted-foreground">
            Try refreshing your purchase information.
          </p>

          <Button variant="outline" class="mt-5 w-full rounded-full sm:w-auto" @click="refetch()">
            Try again
          </Button>
        </div>
      </div>

      <!-- EMPTY -->
      <div v-else-if="!hasItems" class="flex min-h-0 flex-1 items-center justify-center p-6 sm:p-8">
        <div class="w-full max-w-sm text-center">
          <div
            class="mx-auto flex h-16 w-16 items-center justify-center rounded-2xl border border-border bg-muted/30"
          >
            <PackageOpen class="h-7 w-7 text-muted-foreground" />
          </div>

          <h3 class="mt-5 text-lg font-semibold">Your cart is empty</h3>

          <p class="mt-2 text-sm leading-6 text-muted-foreground">
            Browse products from connected businesses and start building your next purchase.
          </p>

          <Button class="mt-6 w-full rounded-full sm:w-auto" @click="continueShopping">
            Explore products
          </Button>
        </div>
      </div>

      <!-- CART -->
      <template v-else>
        <!-- SCROLLABLE CART CONTENT -->
        <div
          class="relative min-h-0 flex-1 overflow-y-auto overscroll-contain px-4 py-4 sm:px-6 sm:py-5"
        >
          <!-- SYNC STATUS -->
          <div
            v-if="isRefreshing || isSynchronizing"
            class="sticky top-0 z-20 mb-4 overflow-hidden rounded-xl border border-primary/20 bg-background/95 shadow-md backdrop-blur"
          >
            <div class="flex items-start gap-3 p-3 sm:p-4">
              <Loader2 class="mt-0.5 h-4 w-4 shrink-0 animate-spin text-primary" />

              <div class="min-w-0">
                <p class="text-sm font-medium">Updating your cart</p>

                <p class="mt-0.5 text-xs leading-5 text-muted-foreground">
                  Synchronizing quantities, totals and current availability.
                </p>
              </div>
            </div>
          </div>

          <div class="space-y-4 sm:space-y-5">
            <section
              v-for="cart in carts"
              :key="cart.invoice_id"
              class="min-w-0 overflow-hidden rounded-2xl border border-border bg-card shadow-sm"
            >
              <!-- SELLER -->
              <div
                class="flex flex-col gap-4 border-b border-border bg-muted/15 px-4 py-4 min-[420px]:flex-row min-[420px]:items-center min-[420px]:justify-between"
              >
                <button
                  type="button"
                  class="group flex min-w-0 items-center gap-3 text-left"
                  @click="openSellerCatalog(cart.seller_company_id)"
                >
                  <div
                    class="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl bg-primary/10"
                  >
                    <Building2 class="h-5 w-5 text-primary" />
                  </div>

                  <div class="min-w-0">
                    <div class="flex min-w-0 items-center gap-2">
                      <p class="truncate text-sm font-semibold">
                        {{ cart.seller_company }}
                      </p>

                      <BadgeCheck class="h-4 w-4 shrink-0 text-primary" />
                    </div>

                    <p
                      class="mt-1 truncate text-xs text-muted-foreground transition group-hover:text-foreground"
                    >
                      Independent order · View seller
                    </p>
                  </div>
                </button>

                <div
                  class="flex shrink-0 items-end justify-between gap-4 border-t border-border/70 pt-3 min-[420px]:block min-[420px]:border-0 min-[420px]:pt-0 min-[420px]:text-right"
                >
                  <p class="text-[10px] font-medium uppercase tracking-wide text-muted-foreground">
                    Seller total
                  </p>

                  <p class="truncate text-base font-semibold tracking-tight">
                    {{ formatCurrency(cart.total_amount) }}
                  </p>
                </div>
              </div>

              <!-- ITEMS -->
              <div class="divide-y divide-border">
                <article v-for="item in cart.items" :key="item.id" class="min-w-0 p-4">
                  <div class="flex min-w-0 gap-3 sm:gap-4">
                    <button
                      type="button"
                      class="h-16 w-16 shrink-0 overflow-hidden rounded-xl border border-border bg-muted/30 sm:h-20 sm:w-20"
                      @click="openProduct(item.product_id)"
                    >
                      <img
                        v-if="getProductImageUrl(item.product_image_path)"
                        :src="getProductImageUrl(item.product_image_path)!"
                        :alt="item.product_name"
                        class="h-full w-full object-cover"
                      />

                      <div v-else class="flex h-full w-full items-center justify-center">
                        <ImageIcon class="h-6 w-6 text-muted-foreground/50" />
                      </div>
                    </button>

                    <div class="min-w-0 flex-1">
                      <div class="flex min-w-0 items-start justify-between gap-2 sm:gap-3">
                        <button
                          type="button"
                          class="min-w-0 flex-1 text-left"
                          @click="openProduct(item.product_id)"
                        >
                          <p
                            class="line-clamp-2 text-sm font-semibold transition hover:text-primary"
                          >
                            {{ item.product_name }}
                          </p>

                          <p class="mt-1 truncate text-xs text-muted-foreground">
                            {{ formatCurrency(item.price) }} per unit
                          </p>
                        </button>

                        <Button
                          variant="ghost"
                          size="icon"
                          class="h-8 w-8 shrink-0 rounded-full text-muted-foreground hover:bg-destructive/10 hover:text-destructive"
                          :disabled="isChangingCart"
                          aria-label="Remove product"
                          @click="deleteItem(cart, item)"
                        >
                          <Loader2
                            v-if="processingItemId === item.id && deleteMutation.isPending.value"
                            class="h-4 w-4 animate-spin"
                          />

                          <Trash2 v-else class="h-4 w-4" />
                        </Button>
                      </div>

                      <div
                        class="mt-4 flex flex-col gap-3 min-[390px]:flex-row min-[390px]:items-end min-[390px]:justify-between"
                      >
                        <div
                          class="inline-flex w-fit items-center rounded-full border border-border bg-muted/30 p-1"
                        >
                          <Button
                            variant="ghost"
                            size="icon"
                            class="h-7 w-7 rounded-full"
                            :disabled="item.quantity <= 1 || isChangingCart"
                            @click="updateQuantity(cart, item, item.quantity - 1)"
                          >
                            <Loader2
                              v-if="processingItemId === item.id && updateMutation.isPending.value"
                              class="h-3.5 w-3.5 animate-spin"
                            />

                            <Minus v-else class="h-3.5 w-3.5" />
                          </Button>

                          <span class="min-w-9 text-center text-sm font-semibold">
                            {{ item.quantity }}
                          </span>

                          <Button
                            variant="ghost"
                            size="icon"
                            class="h-7 w-7 rounded-full"
                            :disabled="isChangingCart"
                            @click="updateQuantity(cart, item, item.quantity + 1)"
                          >
                            <Loader2
                              v-if="processingItemId === item.id && updateMutation.isPending.value"
                              class="h-3.5 w-3.5 animate-spin"
                            />

                            <Plus v-else class="h-3.5 w-3.5" />
                          </Button>
                        </div>

                        <div class="min-w-0 text-left min-[390px]:text-right">
                          <p
                            class="text-[10px] font-medium uppercase tracking-wide text-muted-foreground"
                          >
                            Line subtotal
                          </p>

                          <p class="mt-1 truncate text-base font-semibold">
                            {{ formatCurrency(item.subtotal) }}
                          </p>
                        </div>
                      </div>
                    </div>
                  </div>
                </article>
              </div>

              <!-- SELLER SUMMARY -->
              <div class="border-t border-border bg-muted/10 px-4 py-4">
                <div class="space-y-2 text-sm">
                  <div class="flex min-w-0 justify-between gap-4 text-muted-foreground">
                    <span>Subtotal</span>

                    <span class="min-w-0 truncate text-right">
                      {{ formatCurrency(cart.subtotal) }}
                    </span>
                  </div>

                  <div class="flex min-w-0 justify-between gap-4 text-muted-foreground">
                    <span>Taxes</span>

                    <span class="min-w-0 truncate text-right">
                      {{ formatCurrency(cart.taxes) }}
                    </span>
                  </div>

                  <div
                    class="flex min-w-0 items-end justify-between gap-4 border-t border-border pt-3"
                  >
                    <span class="font-semibold"> Seller order </span>

                    <span class="min-w-0 truncate text-right text-lg font-semibold tracking-tight">
                      {{ formatCurrency(cart.total_amount) }}
                    </span>
                  </div>
                </div>
              </div>
            </section>
          </div>
        </div>

        <!-- GLOBAL SUMMARY -->
        <div
          class="shrink-0 border-t border-border bg-background px-4 pb-[calc(1.25rem+env(safe-area-inset-bottom))] pt-4 shadow-[0_-12px_30px_-20px_rgba(0,0,0,0.35)] sm:px-6 sm:py-5"
        >
          <div class="rounded-[1.5rem] border border-border bg-muted/15 p-4 sm:p-5">
            <div class="mb-5 flex min-w-0 items-start justify-between gap-4">
              <div class="min-w-0">
                <p class="text-sm font-semibold">Complete purchase</p>

                <p class="mt-1 text-xs leading-5 text-muted-foreground">
                  Nexora will create one independent order per seller.
                </p>
              </div>

              <ShoppingBag class="h-5 w-5 shrink-0 text-primary" />
            </div>

            <div class="space-y-3">
              <div class="flex items-center justify-between gap-4 text-sm">
                <span class="text-muted-foreground"> Sellers </span>

                <span class="font-medium">
                  {{ summary.seller_count }}
                </span>
              </div>

              <div class="flex items-center justify-between gap-4 text-sm">
                <span class="text-muted-foreground"> Items </span>

                <span class="font-medium">
                  {{ summary.item_count }}
                </span>
              </div>

              <div class="flex min-w-0 items-center justify-between gap-4 text-sm">
                <span class="text-muted-foreground"> Subtotal </span>

                <span class="min-w-0 truncate text-right">
                  {{ formatCurrency(summary.subtotal) }}
                </span>
              </div>

              <div class="flex min-w-0 items-center justify-between gap-4 text-sm">
                <span class="text-muted-foreground"> Taxes </span>

                <span class="min-w-0 truncate text-right">
                  {{ formatCurrency(summary.taxes) }}
                </span>
              </div>

              <div class="flex min-w-0 items-end justify-between gap-4 border-t border-border pt-4">
                <div class="min-w-0">
                  <p class="text-sm font-semibold">Cart total</p>

                  <p class="mt-1 hidden text-xs text-muted-foreground sm:block">
                    All seller orders included
                  </p>
                </div>

                <p
                  class="min-w-0 truncate text-right text-xl font-semibold tracking-tight sm:text-2xl"
                >
                  {{ formatCurrency(summary.total_amount) }}
                </p>
              </div>
            </div>
          </div>

          <div class="mt-3 grid gap-2 sm:mt-4 sm:grid-cols-[auto_1fr]">
            <Button
              variant="outline"
              class="w-full rounded-full"
              :disabled="isCartBusy"
              @click="continueShopping"
            >
              Continue shopping
            </Button>

            <Button
              size="lg"
              class="w-full rounded-full"
              :disabled="isCartBusy"
              @click="goToCheckout"
            >
              Continue to checkout

              <ArrowRight class="ml-2 h-4 w-4" />
            </Button>
          </div>
        </div>
      </template>
    </SheetContent>
  </Sheet>
</template>
