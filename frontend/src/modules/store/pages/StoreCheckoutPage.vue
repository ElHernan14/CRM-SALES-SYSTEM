<script setup lang="ts">
import { computed, ref } from 'vue';
import { useRouter } from 'vue-router';

import {
  ArrowLeft,
  Building2,
  CheckCircle2,
  ImageIcon,
  Loader2,
  LockKeyhole,
  PackageCheck,
  ShieldCheck,
  ShoppingBag,
  Sparkles,
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
    <section
      class="border-b border-border bg-gradient-to-br from-background via-background to-primary/5"
    >
      <div class="mx-auto max-w-[1300px] px-6 py-12 lg:px-8">
        <Button variant="ghost" class="-ml-3 rounded-full" @click="router.push('/store/catalog')">
          <ArrowLeft class="mr-2 h-4 w-4" />
          Continue shopping
        </Button>

        <div class="mt-8 max-w-3xl">
          <div
            class="inline-flex items-center gap-2 rounded-full border border-primary/20 bg-primary/10 px-4 py-2 text-xs font-semibold text-primary"
          >
            <LockKeyhole class="h-4 w-4" />
            Secure checkout
          </div>

          <h1 class="mt-5 text-4xl font-semibold tracking-[-0.04em] text-foreground sm:text-5xl">
            Review and complete
            <span class="text-muted-foreground"> your purchase. </span>
          </h1>

          <p class="mt-4 max-w-2xl text-base leading-7 text-muted-foreground">
            Nexora organizes your products into one order per seller while completing the entire
            checkout as a single protected operation.
          </p>
        </div>
      </div>
    </section>

    <!-- LOADING -->
    <div
      v-if="isLoading"
      class="mx-auto grid max-w-[1300px] gap-8 px-6 py-10 lg:grid-cols-[minmax(0,1fr)_380px] lg:px-8"
    >
      <div class="space-y-4">
        <div
          v-for="index in 2"
          :key="index"
          class="h-56 animate-pulse rounded-[1.75rem] bg-muted"
        />
      </div>

      <div class="h-96 animate-pulse rounded-[1.75rem] bg-muted" />
    </div>

    <!-- ERROR -->
    <div v-else-if="isError" class="mx-auto max-w-3xl px-6 py-20">
      <EmptyState
        title="Unable to prepare checkout"
        description="There was a problem retrieving your purchase carts."
        :icon="ShoppingBag"
      />

      <div class="mt-5 flex justify-center">
        <Button variant="outline" @click="refetch()"> Try again </Button>
      </div>
    </div>

    <!-- EMPTY -->
    <div v-else-if="!hasItems" class="mx-auto max-w-3xl px-6 py-20">
      <EmptyState
        title="Your cart is empty"
        description="Add products from the Store before continuing to checkout."
        :icon="ShoppingBag"
      />

      <div class="mt-5 flex justify-center">
        <Button class="rounded-full" @click="router.push('/store/catalog')">
          Explore catalog
        </Button>
      </div>
    </div>

    <!-- CHECKOUT -->
    <main
      v-else
      class="mx-auto grid max-w-[1300px] items-start gap-8 px-6 py-10 lg:grid-cols-[minmax(0,1fr)_380px] lg:px-8"
    >
      <!-- ORDERS -->
      <div class="space-y-5">
        <div class="flex items-center justify-between">
          <div>
            <h2 class="text-xl font-semibold tracking-tight">Seller orders</h2>

            <p class="mt-1 text-sm text-muted-foreground">
              {{ summary.seller_count }}
              {{ summary.seller_count === 1 ? 'seller' : 'sellers' }}
              ·
              {{ summary.item_count }}
              {{ summary.item_count === 1 ? 'item' : 'items' }}
            </p>
          </div>

          <Button variant="ghost" size="sm" :disabled="isRefreshing" @click="refetch()">
            <Loader2 v-if="isRefreshing" class="mr-2 h-4 w-4 animate-spin" />

            Refresh
          </Button>
        </div>

        <section
          v-for="cart in carts"
          :key="cart.invoice_id"
          class="overflow-hidden rounded-[1.75rem] border border-border bg-card shadow-sm"
        >
          <div
            class="flex items-center justify-between gap-4 border-b border-border bg-muted/20 px-5 py-4"
          >
            <div class="flex min-w-0 items-center gap-3">
              <div
                class="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl bg-primary/10"
              >
                <Building2 class="h-5 w-5 text-primary" />
              </div>

              <div class="min-w-0">
                <p class="truncate text-sm font-semibold">
                  {{ cart.seller_company }}
                </p>

                <p class="mt-1 text-xs text-muted-foreground">Independent seller order</p>
              </div>
            </div>

            <span
              class="rounded-full border border-border bg-background px-3 py-1 text-xs font-semibold"
            >
              {{ cart.items.length }}
              {{ cart.items.length === 1 ? 'line' : 'lines' }}
            </span>
          </div>

          <div class="divide-y divide-border">
            <article v-for="item in cart.items" :key="item.id" class="flex gap-4 p-5">
              <div
                class="h-20 w-20 shrink-0 overflow-hidden rounded-xl border border-border bg-muted/30"
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
              </div>

              <div class="min-w-0 flex-1">
                <div class="flex items-start justify-between gap-4">
                  <div>
                    <p class="line-clamp-2 text-sm font-semibold">
                      {{ item.product_name }}
                    </p>

                    <p class="mt-1 text-xs text-muted-foreground">
                      {{ item.quantity }} ×
                      {{ formatCurrency(item.price) }}
                    </p>
                  </div>

                  <p class="shrink-0 text-sm font-semibold">
                    {{ formatCurrency(item.subtotal) }}
                  </p>
                </div>
              </div>
            </article>
          </div>

          <div class="space-y-2 border-t border-border bg-muted/10 px-5 py-4 text-sm">
            <div class="flex justify-between text-muted-foreground">
              <span>Subtotal</span>
              <span>
                {{ formatCurrency(cart.subtotal) }}
              </span>
            </div>

            <div class="flex justify-between text-muted-foreground">
              <span>Taxes</span>
              <span>
                {{ formatCurrency(cart.taxes) }}
              </span>
            </div>

            <div class="flex justify-between border-t border-border pt-3 font-semibold">
              <span>Seller total</span>
              <span>
                {{ formatCurrency(cart.total_amount) }}
              </span>
            </div>
          </div>
        </section>
      </div>

      <!-- SUMMARY -->
      <aside class="lg:sticky lg:top-24">
        <div class="overflow-hidden rounded-[1.75rem] border border-border bg-card shadow-xl">
          <div
            class="border-b border-border bg-gradient-to-r from-primary/10 via-primary/5 to-transparent p-5"
          >
            <div class="flex items-center gap-3">
              <div
                class="flex h-11 w-11 items-center justify-center rounded-xl bg-primary text-primary-foreground"
              >
                <PackageCheck class="h-5 w-5" />
              </div>

              <div>
                <p class="text-sm font-semibold">Order summary</p>

                <p class="mt-1 text-xs text-muted-foreground">Multi-seller checkout</p>
              </div>
            </div>
          </div>

          <div class="space-y-5 p-5">
            <div class="space-y-3">
              <div class="flex justify-between text-sm">
                <span class="text-muted-foreground"> Sellers </span>

                <span class="font-medium">
                  {{ summary.seller_count }}
                </span>
              </div>

              <div class="flex justify-between text-sm">
                <span class="text-muted-foreground"> Items </span>

                <span class="font-medium">
                  {{ summary.item_count }}
                </span>
              </div>

              <div class="flex justify-between text-sm">
                <span class="text-muted-foreground"> Subtotal </span>

                <span>
                  {{ formatCurrency(summary.subtotal) }}
                </span>
              </div>

              <div class="flex justify-between text-sm">
                <span class="text-muted-foreground"> Taxes </span>

                <span>
                  {{ formatCurrency(summary.taxes) }}
                </span>
              </div>

              <div class="flex items-end justify-between border-t border-border pt-4">
                <div>
                  <p class="font-semibold">Total</p>

                  <p class="mt-1 text-xs text-muted-foreground">All seller orders included</p>
                </div>

                <p class="text-2xl font-semibold tracking-tight">
                  {{ formatCurrency(summary.total_amount) }}
                </p>
              </div>
            </div>

            <div class="rounded-xl border border-primary/20 bg-primary/5 p-4">
              <div class="flex items-start gap-3">
                <ShieldCheck class="mt-0.5 h-4 w-4 shrink-0 text-primary" />

                <div>
                  <p class="text-sm font-medium">Transactional checkout</p>

                  <p class="mt-1 text-xs leading-5 text-muted-foreground">
                    Every seller order is confirmed together. If one validation fails, no order is
                    submitted.
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
              <Loader2 v-if="isCheckingOut" class="mr-2 h-4 w-4 animate-spin" />

              <LockKeyhole v-else class="mr-2 h-4 w-4" />

              {{ isCheckingOut ? 'Completing purchase...' : 'Complete purchase' }}
            </Button>

            <div class="flex items-center justify-center gap-2 text-xs text-muted-foreground">
              <CheckCircle2 class="h-3.5 w-3.5 text-primary" />
              Secure multi-tenant processing
            </div>
          </div>
        </div>

        <div class="mt-4 grid gap-3 sm:grid-cols-2 lg:grid-cols-1">
          <div class="rounded-2xl border border-border bg-card p-4">
            <ShieldCheck class="h-5 w-5 text-primary" />

            <p class="mt-3 text-sm font-semibold">Protected purchase</p>

            <p class="mt-1 text-xs leading-5 text-muted-foreground">
              Inventory and ownership are validated before submission.
            </p>
          </div>

          <div class="rounded-2xl border border-border bg-card p-4">
            <Sparkles class="h-5 w-5 text-primary" />

            <p class="mt-3 text-sm font-semibold">One checkout</p>

            <p class="mt-1 text-xs leading-5 text-muted-foreground">
              Nexora creates the correct independent order for every seller.
            </p>
          </div>
        </div>
      </aside>
    </main>
  </div>
</template>
