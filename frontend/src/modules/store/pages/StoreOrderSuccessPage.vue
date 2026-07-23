<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';

import { useRoute, useRouter } from 'vue-router';

import {
  ArrowRight,
  Building2,
  Check,
  CheckCircle2,
  PackageCheck,
  ReceiptText,
  ShieldCheck,
  ShoppingBag,
  Sparkles,
} from 'lucide-vue-next';

import { Button } from '@/components/ui/button';

import { getCheckoutResult } from '../utils/checkout-session';

import type { CheckoutAllResponse } from '../types/store-cart.types';
import { useStoreUiStore } from '../stores/store-ui.store';

const route = useRoute();
const router = useRouter();

const storeUi = useStoreUiStore();
const checkoutResult = ref<CheckoutAllResponse | null>(null);

onMounted(() => {
  storeUi.closeCart();

  const historyResult = window.history.state?.checkoutResult as CheckoutAllResponse | undefined;

  checkoutResult.value = historyResult ?? getCheckoutResult();
});

const orders = computed(() => {
  return checkoutResult.value?.orders ?? [];
});

const orderCount = computed(() => {
  return checkoutResult.value?.count ?? orders.value.length;
});

const totalAmount = computed(() => {
  return checkoutResult.value?.total_amount ?? 0;
});

const primaryInvoiceId = computed(() => {
  const routeId = Number(route.params.invoiceId);

  if (Number.isFinite(routeId) && routeId > 0) {
    return routeId;
  }

  return orders.value[0]?.invoice_id ?? null;
});

function formatCurrency(value: number) {
  return new Intl.NumberFormat('es-AR', {
    style: 'currency',
    currency: 'ARS',
    minimumFractionDigits: 2,
  }).format(value);
}

function getStatusLabel(status: string) {
  const labels: Record<string, string> = {
    pending: 'Awaiting payment',
    paid: 'Paid',
    draft: 'Draft',
    cancelled: 'Cancelled',
  };

  return labels[status] ?? status;
}

function openPurchases() {
  router.push('/store/purchases');
}

function continueShopping() {
  router.push('/store/catalog');
}
</script>

<template>
  <main class="relative min-h-[calc(100vh-4rem)] overflow-hidden bg-background pb-20">
    <!-- DECORATION -->
    <div
      class="pointer-events-none absolute -left-48 -top-48 h-[520px] w-[520px] rounded-full bg-primary/10 blur-3xl"
    />

    <div
      class="pointer-events-none absolute -right-56 top-32 h-[520px] w-[520px] rounded-full bg-blue-500/10 blur-3xl"
    />

    <div class="relative mx-auto max-w-[1100px] px-6 py-14 lg:px-8 lg:py-20">
      <!-- SUCCESS HERO -->
      <section class="text-center">
        <div
          class="mx-auto flex h-20 w-20 items-center justify-center rounded-[1.75rem] border border-emerald-500/20 bg-emerald-500/10 shadow-lg shadow-emerald-500/5"
        >
          <CheckCircle2 class="h-10 w-10 text-emerald-600 dark:text-emerald-400" />
        </div>

        <div
          class="mt-7 inline-flex items-center gap-2 rounded-full border border-emerald-500/20 bg-emerald-500/10 px-4 py-2 text-xs font-semibold text-emerald-700 dark:text-emerald-400"
        >
          <Check class="h-4 w-4" />
          Checkout completed successfully
        </div>

        <h1
          class="mx-auto mt-6 max-w-3xl text-4xl font-semibold tracking-[-0.045em] text-foreground sm:text-5xl lg:text-6xl"
        >
          Your purchase is
          <span class="text-muted-foreground"> now being processed. </span>
        </h1>

        <p class="mx-auto mt-5 max-w-2xl text-base leading-7 text-muted-foreground">
          Nexora created an independent order for every seller while processing your checkout as a
          single protected transaction.
        </p>
      </section>

      <!-- SUMMARY -->
      <section
        class="mx-auto mt-10 max-w-4xl overflow-hidden rounded-[2rem] border border-border bg-card shadow-2xl"
      >
        <div
          class="border-b border-border bg-gradient-to-r from-primary/10 via-primary/5 to-transparent p-6"
        >
          <div class="flex flex-col gap-5 sm:flex-row sm:items-center sm:justify-between">
            <div class="flex items-center gap-4">
              <div
                class="flex h-12 w-12 items-center justify-center rounded-2xl bg-primary text-primary-foreground shadow-sm"
              >
                <PackageCheck class="h-6 w-6" />
              </div>

              <div>
                <p class="text-sm font-semibold text-foreground">Purchase summary</p>

                <p class="mt-1 text-xs text-muted-foreground">
                  {{ orderCount }}
                  {{ orderCount === 1 ? 'seller order created' : 'seller orders created' }}
                </p>
              </div>
            </div>

            <div class="sm:text-right">
              <p class="text-xs text-muted-foreground">Checkout total</p>

              <p class="mt-1 text-3xl font-semibold tracking-tight text-foreground">
                {{ formatCurrency(totalAmount) }}
              </p>
            </div>
          </div>
        </div>

        <!-- ORDERS AVAILABLE -->
        <div v-if="orders.length > 0" class="divide-y divide-border">
          <article
            v-for="order in orders"
            :key="order.invoice_id"
            class="flex flex-col gap-4 p-5 sm:flex-row sm:items-center sm:justify-between"
          >
            <div class="flex min-w-0 items-center gap-4">
              <div
                class="flex h-11 w-11 shrink-0 items-center justify-center rounded-xl border border-border bg-muted/30"
              >
                <Building2 class="h-5 w-5 text-muted-foreground" />
              </div>

              <div class="min-w-0">
                <p class="truncate text-sm font-semibold text-foreground">
                  {{ order.seller_company }}
                </p>

                <p class="mt-1 text-xs text-muted-foreground">Seller order confirmed</p>
              </div>
            </div>

            <div class="flex items-center justify-between gap-4 sm:justify-end">
              <span
                class="rounded-full border border-amber-500/20 bg-amber-500/10 px-3 py-1 text-xs font-semibold text-amber-600 dark:text-amber-400"
              >
                {{ getStatusLabel(order.status) }}
              </span>

              <div class="text-right">
                <p class="text-[10px] uppercase tracking-wide text-muted-foreground">
                  Order reference
                </p>

                <p class="mt-1 text-sm font-semibold text-foreground">#NX-{{ order.invoice_id }}</p>
              </div>
            </div>
          </article>
        </div>

        <!-- REFRESH FALLBACK -->
        <div v-else class="p-8 text-center">
          <div
            class="mx-auto flex h-12 w-12 items-center justify-center rounded-xl border border-border bg-muted/30"
          >
            <ReceiptText class="h-5 w-5 text-muted-foreground" />
          </div>

          <p class="mt-4 text-sm font-semibold text-foreground">Purchase completed</p>

          <p class="mt-2 text-sm text-muted-foreground">
            Open your purchase history to review the newly created orders.
          </p>

          <p v-if="primaryInvoiceId" class="mt-3 text-xs text-muted-foreground">
            Primary order reference: #{{ primaryInvoiceId }}
          </p>
        </div>
      </section>

      <!-- PROCESS INFORMATION -->
      <section class="mx-auto mt-6 grid max-w-4xl gap-4 md:grid-cols-3">
        <article class="rounded-2xl border border-border bg-card p-5 shadow-sm">
          <div class="flex h-10 w-10 items-center justify-center rounded-xl bg-primary/10">
            <ReceiptText class="h-5 w-5 text-primary" />
          </div>

          <h2 class="mt-4 text-sm font-semibold text-foreground">Separate orders</h2>

          <p class="mt-2 text-xs leading-5 text-muted-foreground">
            Every seller receives an independent invoice and purchase record.
          </p>
        </article>

        <article class="rounded-2xl border border-border bg-card p-5 shadow-sm">
          <div class="flex h-10 w-10 items-center justify-center rounded-xl bg-blue-500/10">
            <ShieldCheck class="h-5 w-5 text-blue-600 dark:text-blue-400" />
          </div>

          <h2 class="mt-4 text-sm font-semibold text-foreground">Inventory secured</h2>

          <p class="mt-2 text-xs leading-5 text-muted-foreground">
            Product availability was validated before confirming the checkout.
          </p>
        </article>

        <article class="rounded-2xl border border-border bg-card p-5 shadow-sm">
          <div class="flex h-10 w-10 items-center justify-center rounded-xl bg-emerald-500/10">
            <Sparkles class="h-5 w-5 text-emerald-600 dark:text-emerald-400" />
          </div>

          <h2 class="mt-4 text-sm font-semibold text-foreground">Ready for payment</h2>

          <p class="mt-2 text-xs leading-5 text-muted-foreground">
            Pending orders can now be reviewed and paid from your purchases.
          </p>
        </article>
      </section>

      <!-- ACTIONS -->
      <section class="mx-auto mt-8 flex max-w-4xl flex-col gap-3 sm:flex-row sm:justify-center">
        <Button type="button" size="lg" class="rounded-full px-8" @click="openPurchases">
          <ReceiptText class="mr-2 h-4 w-4" />
          View my purchases
          <ArrowRight class="ml-2 h-4 w-4" />
        </Button>

        <Button
          type="button"
          size="lg"
          variant="outline"
          class="rounded-full px-8"
          @click="continueShopping"
        >
          <ShoppingBag class="mr-2 h-4 w-4" />
          Continue shopping
        </Button>
      </section>
    </div>
  </main>
</template>
