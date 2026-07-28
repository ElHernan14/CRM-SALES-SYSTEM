<script setup lang="ts">
import { computed, ref, toRef } from 'vue';
import { toast } from 'vue-sonner';

import {
  Building2,
  Loader2,
  PackageOpen,
  RefreshCw,
  ShoppingBag,
  ShoppingCart,
  Minus,
  Plus,
  Trash2,
} from 'lucide-vue-next';
import { useQueryClient } from '@tanstack/vue-query';

import { Button } from '@/components/ui/button';

import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
} from '@/components/ui/sheet';

import { useMarketplacePurchaseCart } from '../composables/useMarketplacePurchaseCart';

import { useCheckoutMarketplacePurchaseCart } from '../composables/useCheckoutMarketplacePurchaseCart';
import { useUpdatePurchaseCartItem } from '../composables/useUpdatePurchaseCartItem';
import { useDeletePurchaseCartItem } from '../composables/useDeletePurchaseCartItem';

import type { MarketplaceCartItem } from '../types/purchase-cart.types';

const queryClient = useQueryClient();

const props = defineProps<{
  open: boolean;
  sellerCompanyId: number | null;
}>();

const emit = defineEmits<{
  'update:open': [value: boolean];
  checkoutCompleted: [invoiceId: number];
}>();

const sellerCompanyIdRef = toRef(props, 'sellerCompanyId');

const {
  data: cart,
  isLoading,
  isFetching,
  isError,
  refetch,
} = useMarketplacePurchaseCart(sellerCompanyIdRef);

const checkoutMutation = useCheckoutMarketplacePurchaseCart();

const updateItemMutation = useUpdatePurchaseCartItem();

const deleteItemMutation = useDeletePurchaseCartItem();

const processingItemId = ref<number | null>(null);

const isSynchronizingCart = ref(false);

const isChangingCart = computed(() => {
  return (
    isSynchronizingCart.value ||
    updateItemMutation.isPending.value ||
    deleteItemMutation.isPending.value ||
    isRefreshing.value
  );
});

async function synchronizeCart() {
  try {
    isSynchronizingCart.value = true;

    await Promise.all([
      refetch(),

      queryClient.refetchQueries({
        queryKey: ['store-products'],
        type: 'active',
      }),

      queryClient.invalidateQueries({
        queryKey: ['marketplace-purchase-cart'],
      }),
    ]);
  } finally {
    isSynchronizingCart.value = false;
  }
}

const isCheckingOut = computed(() => checkoutMutation.isPending.value);

const isRefreshing = computed(() => isFetching.value && !isLoading.value);

const itemCount = computed(() => {
  return cart.value?.items.reduce((total, item) => total + item.quantity, 0) ?? 0;
});

const hasItems = computed(() => {
  return (cart.value?.items.length ?? 0) > 0;
});

function formatCurrency(value: number) {
  return new Intl.NumberFormat('es-AR', {
    style: 'currency',
    currency: 'ARS',
    minimumFractionDigits: 2,
  }).format(value);
}

async function checkout() {
  if (!cart.value) {
    toast.error('No active purchase cart');
    return;
  }

  if (!hasItems.value) {
    toast.error('Add at least one product before checkout');
    return;
  }

  try {
    const response = await checkoutMutation.mutateAsync(cart.value.invoice_id);

    if (response.source !== 'erp') {
      throw new Error('Marketplace checkout returned an invalid source');
    }

    await Promise.all([
      queryClient.refetchQueries({
        queryKey: ['store-products'],
        type: 'active',
      }),

      queryClient.invalidateQueries({
        queryKey: ['marketplace-purchase-cart'],
      }),

      queryClient.invalidateQueries({
        queryKey: ['purchases'],
      }),
    ]);

    toast.success('Purchase submitted', {
      description: 'The B2B draft was converted into a pending purchase invoice.',
      duration: 4000,
    });

    emit('update:open', false);

    emit('checkoutCompleted', response.invoice_id);
  } catch (error: any) {
    const message =
      error?.response?.data?.errorMessage ??
      error?.response?.data?.message ??
      'Failed to complete checkout';

    toast.error(message);
  }
}

async function updateQuantity(item: MarketplaceCartItem, nextQuantity: number) {
  if (!cart.value) return;
  if (nextQuantity < 1) return;

  try {
    processingItemId.value = item.id;

    await updateItemMutation.mutateAsync({
      invoiceId: cart.value.invoice_id,
      sellerCompanyId: cart.value.seller_company_id,
      itemId: item.id,
      quantity: nextQuantity,
    });

    await synchronizeCart();

    toast.success('Purchase quantity updated');
  } catch (error: any) {
    const message = error?.response?.data?.errorMessage ?? 'Failed to update quantity';

    toast.error(message);
  } finally {
    processingItemId.value = null;
  }
}

async function deleteItem(item: MarketplaceCartItem) {
  if (!cart.value) return;

  try {
    processingItemId.value = item.id;

    await deleteItemMutation.mutateAsync({
      invoiceId: cart.value.invoice_id,
      sellerCompanyId: cart.value.seller_company_id,
      itemId: item.id,
    });

    await synchronizeCart();

    toast.success('Product removed from purchase');
  } catch (error: any) {
    const message = error?.response?.data?.errorMessage ?? 'Failed to remove product';

    toast.error(message);
  } finally {
    processingItemId.value = null;
  }
}
</script>

<template>
  <Sheet :open="open" @update:open="emit('update:open', $event)">
    <SheetContent class="flex w-full flex-col overflow-hidden p-0 sm:max-w-2xl">
      <!-- HEADER -->
      <div class="border-b border-border bg-muted/20 px-6 py-5">
        <SheetHeader class="text-left">
          <div class="flex items-start justify-between gap-4">
            <div class="flex min-w-0 items-center gap-3">
              <div
                class="flex h-11 w-11 shrink-0 items-center justify-center rounded-xl border border-border bg-background shadow-sm"
              >
                <ShoppingBag class="h-5 w-5 text-muted-foreground" />
              </div>

              <div class="min-w-0">
                <SheetTitle> Purchase cart </SheetTitle>

                <SheetDescription class="mt-1">
                  Review this B2B order before checkout.
                </SheetDescription>
              </div>
            </div>

            <Button
              v-if="cart"
              type="button"
              variant="ghost"
              size="icon"
              :disabled="isChangingCart || isCheckingOut"
              @click="synchronizeCart"
            >
              <RefreshCw
                class="h-4 w-4"
                :class="{
                  'animate-spin': isRefreshing,
                }"
              />
            </Button>
          </div>
        </SheetHeader>
      </div>

      <!-- INITIAL LOADING -->
      <div v-if="isLoading" class="flex flex-1 items-center justify-center p-8">
        <div class="text-center">
          <Loader2 class="mx-auto h-7 w-7 animate-spin text-muted-foreground" />

          <p class="mt-3 text-sm text-muted-foreground">Loading purchase cart...</p>
        </div>
      </div>

      <!-- ERROR -->
      <div v-else-if="isError" class="flex flex-1 items-center justify-center p-8">
        <div class="max-w-sm text-center">
          <div
            class="mx-auto flex h-12 w-12 items-center justify-center rounded-2xl border border-destructive/20 bg-destructive/10"
          >
            <ShoppingCart class="h-5 w-5 text-destructive" />
          </div>

          <p class="mt-4 text-sm font-medium text-foreground">Unable to load the cart</p>

          <p class="mt-1 text-sm text-muted-foreground">Try refreshing the purchase information.</p>

          <Button type="button" variant="outline" class="mt-4" @click="refetch()">
            <RefreshCw class="mr-2 h-4 w-4" />
            Try again
          </Button>
        </div>
      </div>

      <!-- EMPTY / NONEXISTENT CART -->
      <div v-else-if="!cart || !hasItems" class="flex flex-1 items-center justify-center p-8">
        <div class="max-w-sm text-center">
          <div
            class="mx-auto flex h-16 w-16 items-center justify-center rounded-2xl border border-border bg-muted/30 shadow-sm"
          >
            <PackageOpen class="h-7 w-7 text-muted-foreground" />
          </div>

          <h3 class="mt-5 text-lg font-semibold text-foreground">Your purchase cart is empty</h3>

          <p class="mt-2 text-sm leading-6 text-muted-foreground">
            Browse this supplier catalog and add products to start building a B2B purchase.
          </p>

          <Button type="button" class="mt-6" @click="emit('update:open', false)">
            Continue browsing
          </Button>
        </div>
      </div>

      <!-- CART -->
      <template v-else-if="cart">
        <!-- SUPPLIER SUMMARY -->
        <div class="border-b border-border px-6 py-4">
          <div
            class="flex items-center justify-between gap-4 rounded-2xl border border-border bg-card p-4 shadow-sm"
          >
            <div class="flex min-w-0 items-center gap-3">
              <div
                class="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl bg-primary/10"
              >
                <Building2 class="h-5 w-5 text-primary" />
              </div>

              <div class="min-w-0">
                <p class="truncate text-sm font-semibold text-foreground">
                  {{ cart.seller_company }}
                </p>

                <p class="mt-1 text-xs text-muted-foreground">B2B purchase order</p>
              </div>
            </div>

            <div class="flex shrink-0 items-center gap-2">
              <span
                class="rounded-full border border-amber-500/20 bg-amber-500/10 px-2.5 py-1 text-xs font-semibold text-amber-600 dark:text-amber-400"
              >
                Draft
              </span>

              <span
                class="rounded-full border border-border bg-muted px-2.5 py-1 text-xs font-medium text-muted-foreground"
              >
                {{ itemCount }}
                {{ itemCount === 1 ? 'unit' : 'units' }}
              </span>
            </div>
          </div>
        </div>

        <!-- SCROLLABLE ITEMS -->
        <div class="relative min-h-0 flex-1 overflow-y-auto px-6 py-5">
          <div
            v-if="isRefreshing || isSynchronizingCart"
            class="absolute inset-x-6 top-5 z-10 rounded-xl border border-border bg-background/90 p-4 shadow-sm backdrop-blur"
          >
            <div class="flex items-center gap-3">
              <Loader2 class="h-4 w-4 animate-spin text-muted-foreground" />

              <p class="text-sm text-muted-foreground">Synchronizing purchase and stock...</p>
            </div>
          </div>

          <div class="space-y-3">
            <article
              v-for="item in cart.items"
              :key="item.id"
              class="rounded-2xl border border-border bg-card p-4 shadow-sm transition hover:border-primary/20"
            >
              <div class="flex items-start justify-between gap-4">
                <div class="min-w-0">
                  <p class="line-clamp-2 text-sm font-semibold text-foreground">
                    {{ item.product_name }}
                  </p>

                  <p class="mt-1 text-xs text-muted-foreground">
                    {{ formatCurrency(item.price) }} per unit
                  </p>
                </div>

                <div class="flex shrink-0 items-center gap-2">
                  <p class="text-sm font-semibold text-foreground">
                    {{ formatCurrency(item.subtotal) }}
                  </p>

                  <Button
                    type="button"
                    variant="ghost"
                    size="icon"
                    class="h-8 w-8 text-muted-foreground hover:text-destructive"
                    :disabled="isChangingCart"
                    @click="deleteItem(item)"
                  >
                    <Loader2
                      v-if="processingItemId === item.id && deleteItemMutation.isPending.value"
                      class="h-4 w-4 animate-spin"
                    />

                    <Trash2 v-else class="h-4 w-4" />
                  </Button>
                </div>
              </div>

              <div class="mt-4 grid grid-cols-2 gap-3 border-t border-border pt-4 sm:grid-cols-3">
                <div>
                  <p class="text-xs text-muted-foreground">Quantity</p>

                  <div
                    class="mt-2 inline-flex items-center rounded-full border border-border bg-muted/40 p-1"
                  >
                    <Button
                      type="button"
                      variant="ghost"
                      size="icon"
                      class="h-7 w-7 rounded-full"
                      :disabled="item.quantity <= 1 || isChangingCart"
                      @click="updateQuantity(item, item.quantity - 1)"
                    >
                      <Loader2
                        v-if="processingItemId === item.id && updateItemMutation.isPending.value"
                        class="h-3.5 w-3.5 animate-spin"
                      />

                      <Minus v-else class="h-3.5 w-3.5" />
                    </Button>

                    <span class="min-w-9 text-center text-sm font-semibold">
                      {{ item.quantity }}
                    </span>

                    <Button
                      type="button"
                      variant="ghost"
                      size="icon"
                      class="h-7 w-7 rounded-full"
                      :disabled="isChangingCart"
                      @click="updateQuantity(item, item.quantity + 1)"
                    >
                      <Loader2
                        v-if="processingItemId === item.id && updateItemMutation.isPending.value"
                        class="h-3.5 w-3.5 animate-spin"
                      />

                      <Plus v-else class="h-3.5 w-3.5" />
                    </Button>
                  </div>
                </div>

                <div>
                  <p class="text-xs text-muted-foreground">Unit price</p>

                  <p class="mt-1 text-sm font-medium text-foreground">
                    {{ formatCurrency(item.price) }}
                  </p>
                </div>

                <div class="col-span-2 sm:col-span-1">
                  <p class="text-xs text-muted-foreground">Line subtotal</p>

                  <p class="mt-1 text-sm font-semibold text-foreground">
                    {{ formatCurrency(item.subtotal) }}
                  </p>
                </div>
              </div>
            </article>
          </div>
        </div>

        <!-- FINANCIAL SUMMARY -->
        <div class="border-t border-border bg-background px-6 py-5">
          <div class="rounded-2xl border border-border bg-muted/20 p-5">
            <div class="space-y-3">
              <div class="flex items-center justify-between">
                <span class="text-sm text-muted-foreground"> Subtotal </span>

                <span class="text-sm font-medium text-foreground">
                  {{ formatCurrency(cart.subtotal) }}
                </span>
              </div>

              <div class="flex items-center justify-between">
                <span class="text-sm text-muted-foreground"> Taxes </span>

                <span class="text-sm font-medium text-foreground">
                  {{ formatCurrency(cart.taxes) }}
                </span>
              </div>

              <div class="flex items-end justify-between border-t border-border pt-4">
                <div>
                  <p class="text-sm font-medium text-foreground">Purchase total</p>

                  <p class="mt-1 text-xs text-muted-foreground">Tax included</p>
                </div>

                <p class="text-2xl font-semibold tracking-tight text-foreground">
                  {{ formatCurrency(cart.total_amount) }}
                </p>
              </div>
            </div>
          </div>

          <div class="mt-4 grid gap-2 sm:grid-cols-[auto_1fr]">
            <Button
              type="button"
              variant="outline"
              :disabled="isCheckingOut"
              @click="emit('update:open', false)"
            >
              Continue browsing
            </Button>

            <Button
              type="button"
              class="w-full"
              :disabled="isCheckingOut || isRefreshing || !hasItems"
              @click="checkout"
            >
              <Loader2 v-if="isCheckingOut" class="mr-2 h-4 w-4 animate-spin" />

              <ShoppingBag v-else class="mr-2 h-4 w-4" />

              {{ isCheckingOut ? 'Processing purchase...' : 'Proceed to checkout' }}
            </Button>
          </div>
        </div>
      </template>
    </SheetContent>
  </Sheet>
</template>
