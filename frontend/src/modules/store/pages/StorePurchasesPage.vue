<script setup lang="ts">
import { computed, ref, watch } from 'vue';

import {
  ArrowDownUp,
  Building2,
  CalendarDays,
  CheckCircle2,
  Clock3,
  CreditCard,
  Eye,
  Loader2,
  ReceiptText,
  RefreshCw,
  Search,
  ShoppingBag,
  WalletCards,
} from 'lucide-vue-next';

import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

import { Select, SelectContent, SelectItem, SelectTrigger } from '@/components/ui/select';

import EmptyState from '@/shared/components/erp/EmptyState.vue';
import DataPagination from '@/shared/components/erp/DataPagination.vue';

import StorePurchaseDetailsDrawer from '../components/purchases/StorePurchaseDetailsDrawer.vue';

import { useStorePurchases } from '../composables/useStorePurchases';

import type { StorePurchase } from '../types/store-purchase.types';

import type { InvoiceStatus } from '@/modules/invoices/types/invoice.types';

const page = ref(1);
const limit = ref(8);

const search = ref('');
const debouncedSearch = ref('');

const statusInvoice = ref('all');

const sortColumn = ref<'created_at' | 'total_amount' | 'paid_amount' | 'status_invoice'>(
  'created_at'
);

const order = ref<'asc' | 'desc'>('desc');

let searchTimer: ReturnType<typeof setTimeout> | undefined;

watch(search, (value) => {
  if (searchTimer) {
    clearTimeout(searchTimer);
  }

  searchTimer = setTimeout(() => {
    debouncedSearch.value = value.trim();
  }, 300);
});

watch([debouncedSearch, statusInvoice, sortColumn, order], () => {
  page.value = 1;
});

const params = computed(() => ({
  page: page.value,
  limit: limit.value,

  search: debouncedSearch.value.length >= 2 ? debouncedSearch.value : undefined,

  status_invoice:
    statusInvoice.value !== 'all' ? (statusInvoice.value as InvoiceStatus) : undefined,

  sort_column: sortColumn.value,
  order: order.value,
}));

const { data, isLoading, isFetching, isError, refetch } = useStorePurchases(params);

const purchases = computed(() => {
  return data.value?.items ?? [];
});

const totalPages = computed(() => {
  return data.value?.meta.total_pages ?? 1;
});

const isRefreshing = computed(() => {
  return isFetching.value && !isLoading.value;
});

/*
 * Métricas correspondientes a la página cargada.
 * Más adelante pueden venir de un endpoint summary.
 */
const pageMetrics = computed(() => {
  return {
    awaitingPayment: purchases.value.filter((purchase) => purchase.status_invoice === 'pending')
      .length,

    paid: purchases.value.filter((purchase) => purchase.status_invoice === 'paid').length,

    outstanding: purchases.value.reduce((total, purchase) => total + purchase.remaining_amount, 0),

    spent: purchases.value.reduce((total, purchase) => total + purchase.paid_amount, 0),
  };
});

const selectedPurchase = ref<StorePurchase | null>(null);

const purchaseDetailsOpen = ref(false);

const purchaseInitialAction = ref<'overview' | 'payment'>('overview');

function openPurchase(purchase: StorePurchase) {
  selectedPurchase.value = purchase;
  purchaseInitialAction.value = 'overview';
  purchaseDetailsOpen.value = true;
}

function openPurchasePayment(purchase: StorePurchase) {
  selectedPurchase.value = purchase;
  purchaseInitialAction.value = 'payment';
  purchaseDetailsOpen.value = true;
}

function canPay(purchase: StorePurchase) {
  return purchase.status_invoice === 'pending' && purchase.remaining_amount > 0;
}

function clearFilters() {
  search.value = '';
  debouncedSearch.value = '';

  statusInvoice.value = 'all';

  sortColumn.value = 'created_at';
  order.value = 'desc';

  page.value = 1;
}

function formatCurrency(value: number) {
  return new Intl.NumberFormat('es-AR', {
    style: 'currency',
    currency: 'ARS',
    minimumFractionDigits: 2,
  }).format(value);
}

function formatDate(value: string) {
  return new Intl.DateTimeFormat('es-AR', {
    day: '2-digit',
    month: 'short',
    year: 'numeric',
  }).format(new Date(value));
}

function getStatusLabel(status: InvoiceStatus) {
  const labels: Record<InvoiceStatus, string> = {
    draft: 'Draft',
    pending: 'Awaiting payment',
    paid: 'Paid',
    cancelled: 'Cancelled',
  };

  return labels[status];
}

function getStatusClass(status: InvoiceStatus) {
  const classes: Record<InvoiceStatus, string> = {
    draft: 'border-border bg-muted text-muted-foreground',

    pending: 'border-amber-500/20 bg-amber-500/10 text-amber-600 dark:text-amber-400',

    paid: 'border-emerald-500/20 bg-emerald-500/10 text-emerald-600 dark:text-emerald-400',

    cancelled: 'border-destructive/20 bg-destructive/10 text-destructive',
  };

  return classes[status];
}
</script>

<template>
  <div class="pb-20">
    <!-- HERO -->
    <section
      class="relative overflow-hidden border-b border-border bg-gradient-to-br from-background via-background to-primary/5"
    >
      <div
        class="pointer-events-none absolute -right-40 -top-56 h-[520px] w-[520px] rounded-full bg-primary/10 blur-3xl"
      />

      <div class="relative mx-auto max-w-[1400px] px-6 py-14 lg:px-8">
        <div class="flex flex-col justify-between gap-8 lg:flex-row lg:items-end">
          <div class="max-w-3xl">
            <p class="text-sm font-semibold text-primary">Nexora account</p>

            <h1 class="mt-3 text-4xl font-semibold tracking-[-0.04em] sm:text-5xl">My purchases</h1>

            <p class="mt-4 max-w-2xl text-base leading-7 text-muted-foreground">
              Review seller orders, outstanding balances, payments and completed purchases.
            </p>
          </div>

          <Button
            variant="outline"
            class="self-start rounded-full lg:self-auto"
            :disabled="isRefreshing"
            @click="refetch()"
          >
            <RefreshCw
              class="mr-2 h-4 w-4"
              :class="{
                'animate-spin': isRefreshing,
              }"
            />

            Refresh purchases
          </Button>
        </div>
      </div>
    </section>

    <main class="mx-auto max-w-[1400px] space-y-8 px-6 py-8 lg:px-8">
      <!-- METRICS -->
      <section class="grid gap-4 md:grid-cols-2 xl:grid-cols-4">
        <article class="rounded-2xl border border-border bg-card p-5 shadow-sm">
          <div class="flex items-start justify-between">
            <div>
              <p class="text-sm text-muted-foreground">Awaiting payment</p>

              <p class="mt-2 text-3xl font-semibold">
                {{ pageMetrics.awaitingPayment }}
              </p>
            </div>

            <div class="rounded-xl bg-amber-500/10 p-3">
              <Clock3 class="h-5 w-5 text-amber-600 dark:text-amber-400" />
            </div>
          </div>

          <p class="mt-4 text-xs text-muted-foreground">On this results page</p>
        </article>

        <article class="rounded-2xl border border-border bg-card p-5 shadow-sm">
          <div class="flex items-start justify-between">
            <div>
              <p class="text-sm text-muted-foreground">Paid orders</p>

              <p class="mt-2 text-3xl font-semibold">
                {{ pageMetrics.paid }}
              </p>
            </div>

            <div class="rounded-xl bg-emerald-500/10 p-3">
              <CheckCircle2 class="h-5 w-5 text-emerald-600 dark:text-emerald-400" />
            </div>
          </div>

          <p class="mt-4 text-xs text-muted-foreground">Completed on this page</p>
        </article>

        <article class="rounded-2xl border border-border bg-card p-5 shadow-sm">
          <div class="flex items-start justify-between gap-4">
            <div class="min-w-0">
              <p class="text-sm text-muted-foreground">Outstanding</p>

              <p class="mt-2 truncate text-2xl font-semibold">
                {{ formatCurrency(pageMetrics.outstanding) }}
              </p>
            </div>

            <div class="rounded-xl bg-primary/10 p-3">
              <WalletCards class="h-5 w-5 text-primary" />
            </div>
          </div>

          <p class="mt-4 text-xs text-muted-foreground">Remaining on visible orders</p>
        </article>

        <article class="rounded-2xl border border-border bg-card p-5 shadow-sm">
          <div class="flex items-start justify-between gap-4">
            <div class="min-w-0">
              <p class="text-sm text-muted-foreground">Amount paid</p>

              <p class="mt-2 truncate text-2xl font-semibold">
                {{ formatCurrency(pageMetrics.spent) }}
              </p>
            </div>

            <div class="rounded-xl bg-blue-500/10 p-3">
              <CreditCard class="h-5 w-5 text-blue-600 dark:text-blue-400" />
            </div>
          </div>

          <p class="mt-4 text-xs text-muted-foreground">Across visible purchases</p>
        </article>
      </section>

      <!-- FILTERS -->
      <section class="rounded-[1.75rem] border border-border bg-card p-5 shadow-sm">
        <div class="grid gap-3 lg:grid-cols-[minmax(0,1fr)_220px_220px_170px_auto]">
          <div class="relative">
            <Search
              class="pointer-events-none absolute left-4 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
            />

            <Input
              v-model="search"
              class="h-11 rounded-full bg-muted/30 pl-11"
              placeholder="Search seller or product..."
            />

            <Loader2
              v-if="isRefreshing"
              class="absolute right-4 top-1/2 h-4 w-4 -translate-y-1/2 animate-spin text-muted-foreground"
            />
          </div>

          <Select v-model="statusInvoice">
            <SelectTrigger class="h-11 rounded-full">
              <span :class="statusInvoice === 'all' ? 'text-muted-foreground' : 'text-foreground'">
                {{
                  statusInvoice === 'all'
                    ? 'All statuses'
                    : getStatusLabel(statusInvoice as InvoiceStatus)
                }}
              </span>
            </SelectTrigger>

            <SelectContent>
              <SelectItem value="all"> All statuses </SelectItem>

              <SelectItem value="pending"> Awaiting payment </SelectItem>

              <SelectItem value="paid"> Paid </SelectItem>

              <SelectItem value="cancelled"> Cancelled </SelectItem>
            </SelectContent>
          </Select>

          <Select v-model="sortColumn">
            <SelectTrigger class="h-11 rounded-full">
              <span>
                {{
                  sortColumn === 'created_at'
                    ? 'Purchase date'
                    : sortColumn === 'total_amount'
                      ? 'Order total'
                      : sortColumn === 'paid_amount'
                        ? 'Amount paid'
                        : 'Status'
                }}
              </span>
            </SelectTrigger>

            <SelectContent>
              <SelectItem value="created_at"> Purchase date </SelectItem>

              <SelectItem value="total_amount"> Order total </SelectItem>

              <SelectItem value="paid_amount"> Amount paid </SelectItem>

              <SelectItem value="status_invoice"> Status </SelectItem>
            </SelectContent>
          </Select>

          <Select v-model="order">
            <SelectTrigger class="h-11 rounded-full">
              <span>
                {{ order === 'desc' ? 'Descending' : 'Ascending' }}
              </span>
            </SelectTrigger>

            <SelectContent>
              <SelectItem value="desc"> Descending </SelectItem>

              <SelectItem value="asc"> Ascending </SelectItem>
            </SelectContent>
          </Select>

          <Button variant="ghost" class="h-11 rounded-full" @click="clearFilters">
            Clear filters
          </Button>
        </div>

        <div class="mt-4 flex items-center justify-between border-t border-border pt-4">
          <p class="text-sm text-muted-foreground">
            {{ data?.meta.total ?? 0 }}
            purchase records
          </p>

          <div class="flex items-center gap-2 text-xs text-muted-foreground">
            <ArrowDownUp class="h-3.5 w-3.5" />
            Personal order history
          </div>
        </div>
      </section>

      <!-- LOADING -->
      <div v-if="isLoading" class="space-y-4">
        <div
          v-for="index in 4"
          :key="index"
          class="h-48 animate-pulse rounded-[1.75rem] bg-muted"
        />
      </div>

      <EmptyState
        v-else-if="isError"
        title="Unable to load purchases"
        description="There was a problem retrieving your order history."
        :icon="ShoppingBag"
      />

      <EmptyState
        v-else-if="purchases.length === 0"
        title="No purchases found"
        description="Completed Store checkouts will appear here."
        :icon="ReceiptText"
      />

      <!-- PURCHASE CARDS -->
      <template v-else>
        <div class="space-y-4">
          <article
            v-for="purchase in purchases"
            :key="purchase.id"
            class="group overflow-hidden rounded-[1.75rem] border border-border bg-card shadow-sm transition hover:border-primary/25 hover:shadow-lg"
          >
            <div
              class="grid gap-5 p-5 lg:grid-cols-[minmax(0,1fr)_180px_180px_auto] lg:items-center"
            >
              <div class="flex min-w-0 gap-4">
                <div
                  class="flex h-12 w-12 shrink-0 items-center justify-center rounded-2xl bg-primary/10"
                >
                  <Building2 class="h-5 w-5 text-primary" />
                </div>

                <div class="min-w-0">
                  <div class="flex flex-wrap items-center gap-2">
                    <h2 class="truncate text-base font-semibold">
                      {{ purchase.seller_company }}
                    </h2>

                    <span
                      class="inline-flex rounded-full border px-2.5 py-1 text-xs font-semibold"
                      :class="getStatusClass(purchase.status_invoice)"
                    >
                      {{ getStatusLabel(purchase.status_invoice) }}
                    </span>
                  </div>

                  <div
                    class="mt-2 flex flex-wrap items-center gap-x-4 gap-y-2 text-xs text-muted-foreground"
                  >
                    <span class="flex items-center gap-1.5">
                      <CalendarDays class="h-3.5 w-3.5" />

                      {{ formatDate(purchase.created_at) }}
                    </span>

                    <span>
                      {{ purchase.item_count }}
                      {{ purchase.item_count === 1 ? 'item' : 'items' }}
                    </span>

                    <span> Order #NX-{{ purchase.id }} </span>
                  </div>
                </div>
              </div>

              <div>
                <p class="text-xs text-muted-foreground">Order total</p>

                <p class="mt-1 text-lg font-semibold">
                  {{ formatCurrency(purchase.total_amount) }}
                </p>
              </div>

              <div>
                <p class="text-xs text-muted-foreground">Remaining</p>

                <p
                  class="mt-1 text-lg font-semibold"
                  :class="
                    purchase.remaining_amount > 0
                      ? 'text-amber-600 dark:text-amber-400'
                      : 'text-emerald-600 dark:text-emerald-400'
                  "
                >
                  {{ formatCurrency(purchase.remaining_amount) }}
                </p>
              </div>

              <div class="flex flex-wrap gap-2 lg:justify-end">
                <Button variant="outline" class="rounded-full" @click="openPurchase(purchase)">
                  <Eye class="mr-2 h-4 w-4" />
                  View order
                </Button>

                <Button
                  v-if="canPay(purchase)"
                  class="rounded-full"
                  @click="openPurchasePayment(purchase)"
                >
                  <WalletCards class="mr-2 h-4 w-4" />
                  Pay now
                </Button>
              </div>
            </div>

            <div class="h-1 bg-muted">
              <div
                class="h-full bg-primary transition-all"
                :style="{
                  width: `${
                    purchase.total_amount > 0
                      ? Math.min((purchase.paid_amount / purchase.total_amount) * 100, 100)
                      : 0
                  }%`,
                }"
              />
            </div>
          </article>
        </div>

        <DataPagination
          class="mt-8"
          :page="page"
          :total-pages="totalPages"
          :total="data?.meta.total"
          @previous="page--"
          @next="page++"
        />
      </template>
    </main>

    <StorePurchaseDetailsDrawer
      v-model:open="purchaseDetailsOpen"
      :purchase="selectedPurchase"
      :initial-action="purchaseInitialAction"
    />
  </div>
</template>
