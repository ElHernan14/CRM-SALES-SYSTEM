<script setup lang="ts">
import { computed, ref } from 'vue';
import { useRouter } from 'vue-router';

import {
  ArrowLeft,
  BadgeCheck,
  Building2,
  CheckCircle2,
  ImageIcon,
  Loader2,
  LockKeyhole,
  PackageCheck,
  RefreshCw,
  ShieldCheck,
  ShoppingBag,
} from 'lucide-vue-next';

import { toast } from 'vue-sonner';

import { Button } from '@/components/ui/button';

import EmptyState from '@/shared/components/erp/EmptyState.vue';

import { getProductImageUrl } from '@/shared/utils/assets';

import { useStoreCarts } from '../composables/useStoreCarts';
import { useCheckoutAllStoreCarts } from '../composables/useCheckoutAllStoreCarts';

import { saveCheckoutResult } from '../utils/checkout-session';

const router = useRouter();

const { data, isLoading, isFetching, isError, refetch } = useStoreCarts();

const checkoutMutation = useCheckoutAllStoreCarts();

const confirmingCheckout = ref(false);

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

const invoiceIds = computed(() => {
  return carts.value.map((cart) => cart.invoice_id);
});

const isRefreshing = computed(() => {
  return isFetching.value && !isLoading.value;
});

const isCheckingOut = computed(() => {
  return confirmingCheckout.value || checkoutMutation.isPending.value;
});

function formatCurrency(value: number) {
  return new Intl.NumberFormat('es-AR', {
    style: 'currency',
    currency: 'ARS',
    minimumFractionDigits: 2,
  }).format(value);
}

function continueShopping() {
  router.push({
    name: 'store-catalog',
  });
}

function openProduct(productId: number) {
  router.push({
    name: 'store-product',
    params: {
      productId,
    },
  });
}

function openSellerCatalog(sellerCompanyId: number) {
  router.push({
    name: 'store-catalog',
    query: {
      company_id: String(sellerCompanyId),
    },
  });
}

async function completeCheckout() {
  if (!hasItems.value) {
    toast.error('Your cart is empty');
    return;
  }

  if (invoiceIds.value.length === 0) {
    toast.error('No purchase orders available');
    return;
  }

  try {
    confirmingCheckout.value = true;

    const response = await checkoutMutation.mutateAsync({
      invoice_ids: invoiceIds.value,
    });

    saveCheckoutResult(response);

    toast.success('Purchase completed', {
      description: `${response.count} ${
        response.count === 1 ? 'order was' : 'orders were'
      } submitted successfully.`,
    });

    const primaryInvoiceId = response.orders[0]?.invoice_id;

    if (!primaryInvoiceId) {
      toast.error('Checkout completed, but no orders were returned');

      return;
    }

    await router.replace({
      name: 'store-order-success',

      params: {
        invoiceId: primaryInvoiceId,
      },

      state: {
        checkoutResult: response,
      },
    });
  } catch (error: any) {
    const message =
      error?.response?.data?.errorMessage ??
      error?.response?.data?.message ??
      'Unable to complete checkout';

    toast.error(message);
  } finally {
    confirmingCheckout.value = false;
  }
}
</script>

<template>
  <div class="pb-20">
    <!-- HERO -->
    <section
      class="border-b border-border bg-gradient-to-br from-background via-background to-primary/5"
    >
      <div class="mx-auto w-full max-w-[1300px] px-4 py-10 sm:px-6 sm:py-12 lg:px-8">
        <Button variant="ghost" class="-ml-3 rounded-full" @click="continueShopping">
          <ArrowLeft class="mr-2 h-4 w-4" />
          Continue shopping
        </Button>

        <div class="mt-7 max-w-3xl sm:mt-8">
          <p class="text-sm font-semibold text-primary">Nexora checkout</p>

          <h1
            class="mt-4 text-4xl font-semibold leading-[1.05] tracking-[-0.045em] text-foreground sm:text-5xl"
          >
            Review your orders before

            <span class="text-muted-foreground"> submitting the purchase. </span>
          </h1>

          <p class="mt-4 max-w-2xl text-base leading-7 text-muted-foreground">
            Products remain grouped by seller, while Nexora validates and submits every order in one
            transactional operation.
          </p>

          <div
            class="mt-6 grid gap-3 text-sm text-muted-foreground sm:flex sm:flex-wrap sm:gap-x-5 sm:gap-y-3"
          >
            <span class="flex items-center gap-2">
              <BadgeCheck class="h-4 w-4 shrink-0 text-primary" />
              Verified sellers
            </span>

            <span class="flex items-center gap-2">
              <PackageCheck class="h-4 w-4 shrink-0 text-primary" />
              Inventory validation
            </span>

            <span class="flex items-center gap-2">
              <LockKeyhole class="h-4 w-4 shrink-0 text-primary" />
              Atomic submission
            </span>
          </div>
        </div>
      </div>
    </section>

    <!-- LOADING -->
    <div
      v-if="isLoading"
      class="mx-auto grid w-full max-w-[1300px] gap-6 px-4 py-8 sm:px-6 sm:py-10 lg:grid-cols-[minmax(0,1fr)_380px] lg:gap-8 lg:px-8"
    >
      <div class="space-y-4">
        <div
          v-for="index in 2"
          :key="index"
          class="h-56 animate-pulse rounded-[1.5rem] bg-muted sm:rounded-[1.75rem]"
        />
      </div>

      <div class="h-80 animate-pulse rounded-[1.5rem] bg-muted sm:h-96 sm:rounded-[1.75rem]" />
    </div>

    <!-- ERROR -->
    <div v-else-if="isError" class="mx-auto max-w-3xl px-4 py-16 sm:px-6 sm:py-20">
      <EmptyState
        title="Unable to prepare checkout"
        description="There was a problem retrieving your purchase carts."
        :icon="ShoppingBag"
      />

      <div class="mt-5 flex justify-center">
        <Button variant="outline" class="w-full rounded-full sm:w-auto" @click="refetch()">
          Try again
        </Button>
      </div>
    </div>

    <!-- EMPTY -->
    <div v-else-if="!hasItems" class="mx-auto max-w-3xl px-4 py-16 sm:px-6 sm:py-20">
      <EmptyState
        title="Your cart is empty"
        description="Add products from the Store before continuing to checkout."
        :icon="ShoppingBag"
      />

      <div class="mt-5 flex justify-center">
        <Button class="w-full rounded-full sm:w-auto" @click="continueShopping">
          Explore catalog
        </Button>
      </div>
    </div>

    <!-- CHECKOUT -->
    <main
      v-else
      class="mx-auto grid w-full max-w-[1300px] items-start gap-8 px-4 py-8 sm:px-6 sm:py-10 lg:grid-cols-[minmax(0,1fr)_380px] lg:px-8"
    >
      <!-- ORDERS -->
      <div class="min-w-0 space-y-5">
        <div
          class="flex flex-col gap-4 min-[420px]:flex-row min-[420px]:items-center min-[420px]:justify-between"
        >
          <div>
            <h2 class="text-xl font-semibold tracking-tight">Seller orders</h2>

            <p class="mt-1 text-sm text-muted-foreground">
              {{ summary.seller_count }}

              {{ summary.seller_count === 1 ? 'seller' : 'sellers' }}

              · {{ summary.item_count }}

              {{ summary.item_count === 1 ? 'item' : 'items' }}
            </p>
          </div>

          <Button
            variant="outline"
            size="sm"
            class="w-full rounded-full min-[420px]:w-auto"
            :disabled="isRefreshing"
            @click="refetch()"
          >
            <Loader2 v-if="isRefreshing" class="mr-2 h-4 w-4 animate-spin" />

            <RefreshCw v-else class="mr-2 h-4 w-4" />

            Refresh
          </Button>
        </div>

        <section
          v-for="cart in carts"
          :key="cart.invoice_id"
          class="min-w-0 overflow-hidden rounded-[1.5rem] border border-border bg-card shadow-sm sm:rounded-[1.75rem]"
        >
          <!-- SELLER HEADER -->
          <div
            class="flex flex-col gap-4 border-b border-border bg-muted/20 px-4 py-4 min-[420px]:flex-row min-[420px]:items-center min-[420px]:justify-between sm:px-5"
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
                <p class="truncate text-sm font-semibold transition group-hover:text-primary">
                  {{ cart.seller_company }}
                </p>

                <div class="mt-1 flex min-w-0 items-center gap-2">
                  <p class="truncate text-xs text-muted-foreground">Independent seller order</p>

                  <BadgeCheck class="h-3.5 w-3.5 shrink-0 text-primary" />
                </div>
              </div>
            </button>

            <span
              class="w-fit rounded-full border border-border bg-background px-3 py-1 text-xs font-semibold"
            >
              {{ cart.items.length }}

              {{ cart.items.length === 1 ? 'line' : 'lines' }}
            </span>
          </div>

          <!-- ITEMS -->
          <div class="divide-y divide-border">
            <article v-for="item in cart.items" :key="item.id" class="min-w-0 p-4 sm:p-5">
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
                    <ImageIcon class="h-6 w-6 text-muted-foreground/40" />
                  </div>
                </button>

                <div class="min-w-0 flex-1">
                  <div
                    class="flex flex-col gap-3 min-[420px]:flex-row min-[420px]:items-start min-[420px]:justify-between"
                  >
                    <button
                      type="button"
                      class="min-w-0 text-left"
                      @click="openProduct(item.product_id)"
                    >
                      <p class="line-clamp-2 text-sm font-semibold transition hover:text-primary">
                        {{ item.product_name }}
                      </p>

                      <p class="mt-1 text-xs leading-5 text-muted-foreground">
                        {{ item.quantity }}

                        {{ item.quantity === 1 ? 'unit' : 'units' }}

                        · {{ formatCurrency(item.price) }} each
                      </p>
                    </button>

                    <div class="min-w-0 shrink-0 text-left min-[420px]:text-right">
                      <p class="text-[10px] uppercase tracking-wide text-muted-foreground">
                        Subtotal
                      </p>

                      <p class="mt-1 truncate text-sm font-semibold">
                        {{ formatCurrency(item.subtotal) }}
                      </p>
                    </div>
                  </div>
                </div>
              </div>
            </article>
          </div>

          <!-- SELLER TOTAL -->
          <div class="space-y-2 border-t border-border bg-muted/10 px-4 py-4 text-sm sm:px-5">
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

            <div class="flex min-w-0 items-end justify-between gap-4 border-t border-border pt-3">
              <div>
                <p class="font-semibold">Order total</p>

                <p class="mt-1 hidden text-xs text-muted-foreground min-[420px]:block">
                  {{ cart.seller_company }}
                </p>
              </div>

              <p class="min-w-0 truncate text-right text-lg font-semibold tracking-tight">
                {{ formatCurrency(cart.total_amount) }}
              </p>
            </div>
          </div>
        </section>
      </div>

      <!-- SUMMARY -->
      <aside class="min-w-0 lg:sticky lg:top-24">
        <div
          class="overflow-hidden rounded-[1.5rem] border border-border bg-card shadow-xl sm:rounded-[1.75rem]"
        >
          <div
            class="border-b border-border bg-gradient-to-r from-primary/10 via-primary/5 to-transparent p-4 sm:p-5"
          >
            <div class="flex min-w-0 items-center gap-3">
              <div
                class="flex h-11 w-11 shrink-0 items-center justify-center rounded-xl bg-primary text-primary-foreground"
              >
                <PackageCheck class="h-5 w-5" />
              </div>

              <div class="min-w-0">
                <p class="text-sm font-semibold">Purchase summary</p>

                <p class="mt-1 truncate text-xs text-muted-foreground">
                  {{ summary.seller_count }}

                  {{ summary.seller_count === 1 ? 'seller order' : 'seller orders' }}
                </p>
              </div>
            </div>
          </div>

          <div class="space-y-5 p-4 sm:p-5">
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

              <div class="rounded-2xl border border-primary/20 bg-primary/5 p-4">
                <p class="text-xs font-medium uppercase tracking-[0.12em] text-muted-foreground">
                  Purchase total
                </p>

                <p
                  class="mt-2 break-words text-2xl font-semibold tracking-[-0.04em] min-[390px]:text-3xl"
                >
                  {{ formatCurrency(summary.total_amount) }}
                </p>

                <p class="mt-2 text-xs leading-5 text-muted-foreground">
                  Includes all seller orders and calculated taxes.
                </p>
              </div>
            </div>

            <div class="rounded-xl border border-primary/20 bg-primary/5 p-4">
              <div class="flex items-start gap-3">
                <ShieldCheck class="mt-0.5 h-4 w-4 shrink-0 text-primary" />

                <div class="min-w-0">
                  <p class="text-sm font-medium">Transactional order submission</p>

                  <p class="mt-1 text-xs leading-5 text-muted-foreground">
                    Nexora validates every order together. If one seller order fails validation,
                    none of the orders are submitted.
                  </p>
                </div>
              </div>
            </div>

            <Button
              size="lg"
              class="w-full rounded-full"
              :disabled="isCheckingOut"
              @click="completeCheckout"
            >
              <span class="mr-2 flex h-4 w-4 items-center justify-center">
                <Loader2 v-if="isCheckingOut" class="h-4 w-4 animate-spin" />

                <LockKeyhole v-else class="h-4 w-4" />
              </span>

              {{
                isCheckingOut
                  ? 'Submitting orders...'
                  : `Submit ${summary.seller_count} ${
                      summary.seller_count === 1 ? 'order' : 'orders'
                    }`
              }}
            </Button>

            <div
              class="flex items-start justify-center gap-2 text-xs leading-5 text-muted-foreground"
            >
              <CheckCircle2 class="mt-0.5 h-3.5 w-3.5 shrink-0 text-primary" />

              <span> Orders will move from cart to purchase history. </span>
            </div>
          </div>
        </div>

        <div class="mt-4 rounded-2xl border border-border bg-muted/15 p-4">
          <div class="flex items-start gap-3">
            <ShieldCheck class="mt-0.5 h-5 w-5 shrink-0 text-primary" />

            <div class="min-w-0">
              <p class="text-sm font-semibold">What happens next?</p>

              <p class="mt-1 text-xs leading-5 text-muted-foreground">
                Each seller receives an independent pending order. You can track status and payments
                from My Purchases.
              </p>
            </div>
          </div>
        </div>
      </aside>
    </main>
  </div>
</template>
