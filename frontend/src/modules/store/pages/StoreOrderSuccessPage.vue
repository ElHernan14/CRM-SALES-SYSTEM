<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';

import { useRoute, useRouter } from 'vue-router';

import {
  ArrowRight,
  BadgeCheck,
  Building2,
  Check,
  CheckCircle2,
  PackageCheck,
  ReceiptText,
  ShieldCheck,
  ShoppingBag,
  WalletCards,
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

function getStatusClass(status: string) {
  const classes: Record<string, string> = {
    pending: 'border-amber-500/20 bg-amber-500/10 text-amber-600 dark:text-amber-400',

    paid: 'border-emerald-500/20 bg-emerald-500/10 text-emerald-600 dark:text-emerald-400',

    draft: 'border-border bg-muted text-muted-foreground',

    cancelled: 'border-destructive/20 bg-destructive/10 text-destructive',
  };

  return classes[status] ?? 'border-border bg-muted text-muted-foreground';
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
    <div
      class="pointer-events-none absolute -left-48 -top-48 h-[520px] w-[520px] rounded-full bg-primary/10 blur-3xl"
    />

    <div
      class="pointer-events-none absolute -right-56 top-32 h-[520px] w-[520px] rounded-full bg-blue-500/10 blur-3xl"
    />

    <div
      class="relative mx-auto w-full max-w-[1100px] px-4 py-12 sm:px-6 sm:py-14 lg:px-8 lg:py-20"
    >
      <!-- SUCCESS HERO -->
      <section class="text-center">
        <div
          class="mx-auto flex h-16 w-16 items-center justify-center rounded-[1.5rem] border border-emerald-500/20 bg-emerald-500/10 shadow-lg shadow-emerald-500/5 sm:h-20 sm:w-20 sm:rounded-[1.75rem]"
        >
          <CheckCircle2 class="h-8 w-8 text-emerald-600 dark:text-emerald-400 sm:h-10 sm:w-10" />
        </div>

        <p class="mt-6 text-sm font-semibold text-emerald-600 dark:text-emerald-400 sm:mt-7">
          Purchase orders submitted
        </p>

        <h1
          class="mx-auto mt-4 max-w-3xl text-4xl font-semibold leading-[1.04] tracking-[-0.05em] text-foreground sm:text-5xl lg:text-6xl"
        >
          Your orders are ready for

          <span class="text-muted-foreground"> the next step. </span>
        </h1>

        <p class="mx-auto mt-5 max-w-2xl text-base leading-7 text-muted-foreground">
          Nexora created one independent purchase order for every seller and moved them into your
          personal purchase history.
        </p>

        <div
          class="mx-auto mt-7 grid max-w-2xl gap-3 text-left text-sm text-muted-foreground sm:flex sm:flex-wrap sm:justify-center sm:gap-x-6 sm:gap-y-3"
        >
          <span class="flex items-center gap-2">
            <BadgeCheck class="h-4 w-4 shrink-0 text-primary" />
            Seller orders confirmed
          </span>

          <span class="flex items-center gap-2">
            <PackageCheck class="h-4 w-4 shrink-0 text-primary" />
            Inventory validated
          </span>

          <span class="flex items-center gap-2">
            <WalletCards class="h-4 w-4 shrink-0 text-primary" />
            Ready for payment
          </span>
        </div>
      </section>

      <!-- SUMMARY -->
      <section
        class="mx-auto mt-10 max-w-4xl overflow-hidden rounded-[1.5rem] border border-border bg-card shadow-2xl sm:rounded-[2rem]"
      >
        <div
          class="border-b border-border bg-gradient-to-r from-primary/10 via-primary/5 to-transparent p-5 sm:p-6"
        >
          <div class="flex flex-col gap-5 sm:flex-row sm:items-end sm:justify-between">
            <div class="flex min-w-0 items-center gap-3 sm:gap-4">
              <div
                class="flex h-11 w-11 shrink-0 items-center justify-center rounded-xl bg-primary text-primary-foreground shadow-sm sm:h-12 sm:w-12 sm:rounded-2xl"
              >
                <ReceiptText class="h-5 w-5 sm:h-6 sm:w-6" />
              </div>

              <div class="min-w-0">
                <p class="text-sm font-semibold text-foreground">Submitted purchase</p>

                <p class="mt-1 text-xs leading-5 text-muted-foreground">
                  {{ orderCount }}

                  {{ orderCount === 1 ? 'independent seller order' : 'independent seller orders' }}
                </p>
              </div>
            </div>

            <div class="min-w-0 sm:text-right">
              <p class="text-xs font-medium uppercase tracking-[0.12em] text-muted-foreground">
                Purchase total
              </p>

              <p
                class="mt-2 break-words text-2xl font-semibold tracking-[-0.04em] text-foreground sm:text-3xl"
              >
                {{ formatCurrency(totalAmount) }}
              </p>
            </div>
          </div>
        </div>

        <div v-if="orders.length > 0" class="divide-y divide-border">
          <article
            v-for="order in orders"
            :key="order.invoice_id"
            class="flex min-w-0 flex-col gap-5 p-4 transition hover:bg-muted/15 sm:flex-row sm:items-center sm:justify-between sm:p-5"
          >
            <div class="flex min-w-0 items-center gap-3 sm:gap-4">
              <div
                class="flex h-11 w-11 shrink-0 items-center justify-center rounded-xl bg-primary/10 sm:h-12 sm:w-12"
              >
                <Building2 class="h-5 w-5 text-primary" />
              </div>

              <div class="min-w-0">
                <div class="flex min-w-0 items-center gap-2">
                  <p class="truncate text-base font-semibold text-foreground">
                    {{ order.seller_company }}
                  </p>

                  <BadgeCheck class="h-4 w-4 shrink-0 text-primary" />
                </div>

                <p class="mt-1 text-xs text-muted-foreground">Independent seller order</p>
              </div>
            </div>

            <div
              class="flex items-end justify-between gap-4 border-t border-border pt-4 sm:justify-end sm:border-0 sm:pt-0"
            >
              <span
                class="rounded-full border px-3 py-1 text-xs font-semibold"
                :class="getStatusClass(order.status)"
              >
                {{ getStatusLabel(order.status) }}
              </span>

              <div class="text-right">
                <p class="text-[10px] font-medium uppercase tracking-wide text-muted-foreground">
                  Order reference
                </p>

                <p class="mt-1 text-base font-semibold tracking-tight text-foreground">
                  #NX-{{ order.invoice_id }}
                </p>
              </div>
            </div>
          </article>
        </div>

        <div v-else class="p-6 text-center sm:p-8">
          <div
            class="mx-auto flex h-12 w-12 items-center justify-center rounded-xl border border-border bg-muted/30"
          >
            <ReceiptText class="h-5 w-5 text-muted-foreground" />
          </div>

          <p class="mt-4 text-sm font-semibold text-foreground">Purchase successfully submitted</p>

          <p class="mt-2 text-sm leading-6 text-muted-foreground">
            Open your purchase history to review the newly created seller orders.
          </p>

          <p v-if="primaryInvoiceId" class="mt-4 text-sm font-semibold text-foreground">
            Primary reference: #NX-{{ primaryInvoiceId }}
          </p>
        </div>

        <div class="border-t border-border bg-muted/15 p-4 sm:p-5">
          <div class="flex items-start gap-3 sm:gap-4">
            <div
              class="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl bg-amber-500/10 sm:h-11 sm:w-11"
            >
              <WalletCards class="h-5 w-5 text-amber-600 dark:text-amber-400" />
            </div>

            <div class="min-w-0">
              <p class="text-sm font-semibold">What happens next?</p>

              <p class="mt-1 text-sm leading-6 text-muted-foreground">
                Orders awaiting payment can be reviewed and paid from My Purchases. Their status
                will update as payments are recorded.
              </p>
            </div>
          </div>
        </div>
      </section>

      <!-- FLOW -->
      <section
        class="mx-auto mt-6 max-w-4xl rounded-[1.5rem] border border-border bg-card p-5 shadow-sm sm:rounded-[1.75rem] sm:p-6"
      >
        <p class="text-xs font-semibold uppercase tracking-[0.16em] text-primary">
          Nexora order flow
        </p>

        <div class="mt-5 grid gap-5 md:grid-cols-3 md:gap-4">
          <article class="flex items-start gap-3">
            <div
              class="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl bg-primary/10"
            >
              <Check class="h-5 w-5 text-primary" />
            </div>

            <div>
              <p class="text-sm font-semibold">1. Orders created</p>

              <p class="mt-1 text-xs leading-5 text-muted-foreground">
                One purchase record was generated for each seller.
              </p>
            </div>
          </article>

          <article class="flex items-start gap-3">
            <div
              class="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl bg-blue-500/10"
            >
              <ShieldCheck class="h-5 w-5 text-blue-600 dark:text-blue-400" />
            </div>

            <div>
              <p class="text-sm font-semibold">2. Inventory confirmed</p>

              <p class="mt-1 text-xs leading-5 text-muted-foreground">
                Availability was validated before submission.
              </p>
            </div>
          </article>

          <article class="flex items-start gap-3">
            <div
              class="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl bg-emerald-500/10"
            >
              <WalletCards class="h-5 w-5 text-emerald-600 dark:text-emerald-400" />
            </div>

            <div>
              <p class="text-sm font-semibold">3. Payment tracking</p>

              <p class="mt-1 text-xs leading-5 text-muted-foreground">
                Pay and follow each seller order independently.
              </p>
            </div>
          </article>
        </div>
      </section>

      <!-- ACTIONS -->
      <section class="mx-auto mt-8 flex max-w-4xl flex-col gap-3 sm:flex-row sm:justify-center">
        <Button
          type="button"
          size="lg"
          class="group w-full rounded-full px-8 sm:w-auto"
          @click="openPurchases"
        >
          <ReceiptText class="mr-2 h-4 w-4" />

          View my purchases

          <ArrowRight
            class="ml-2 h-4 w-4 transition-transform duration-200 group-hover:translate-x-0.5"
          />
        </Button>

        <Button
          type="button"
          size="lg"
          variant="outline"
          class="w-full rounded-full px-8 sm:w-auto"
          @click="continueShopping"
        >
          <ShoppingBag class="mr-2 h-4 w-4" />
          Continue shopping
        </Button>
      </section>
    </div>
  </main>
</template>
