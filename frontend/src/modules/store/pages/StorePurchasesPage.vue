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

import { useRoute, useRouter } from 'vue-router';

const route = useRoute();
const router = useRouter();

function queryString(value: unknown) {
  return typeof value === 'string' ? value : '';
}

const limit = ref(8);

const search = ref(queryString(route.query.search));

const debouncedSearch = ref(queryString(route.query.search).trim());

const statusInvoice = ref(queryString(route.query.status) || 'all');

const page = ref(Number(route.query.page) > 0 ? Number(route.query.page) : 1);

const querySort = queryString(route.query.sort);

const sortColumn = ref<'created_at' | 'total_amount' | 'paid_amount' | 'status_invoice'>(
  ['created_at', 'total_amount', 'paid_amount', 'status_invoice'].includes(querySort)
    ? (querySort as 'created_at' | 'total_amount' | 'paid_amount' | 'status_invoice')
    : 'created_at'
);

const order = ref<'asc' | 'desc'>(queryString(route.query.order) === 'asc' ? 'asc' : 'desc');

let searchTimer: ReturnType<typeof setTimeout> | undefined;

watch(search, (value) => {
  if (searchTimer) {
    clearTimeout(searchTimer);
  }

  searchTimer = setTimeout(() => {
    debouncedSearch.value = value.trim();
  }, 300);
});

const purchaseFilterSources = [debouncedSearch, statusInvoice, sortColumn, order] as const;

watch(
  purchaseFilterSources,
  () => {
    if (page.value !== 1) {
      page.value = 1;
    }
  },
  {
    flush: 'sync',
  }
);

watch([...purchaseFilterSources, page], () => {
  const query: Record<string, string> = {};

  if (debouncedSearch.value.length >= 2) {
    query.search = debouncedSearch.value;
  }

  if (statusInvoice.value !== 'all') {
    query.status = statusInvoice.value;
  }

  if (sortColumn.value !== 'created_at') {
    query.sort = sortColumn.value;
  }

  if (order.value !== 'desc') {
    query.order = order.value;
  }

  if (page.value > 1) {
    query.page = String(page.value);
  }

  router.replace({
    name: 'store-purchases',
    query,
  });
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

function getPaymentProgress(purchase: StorePurchase) {
  if (purchase.total_amount <= 0) {
    return 0;
  }

  return Math.min(Math.max((purchase.paid_amount / purchase.total_amount) * 100, 0), 100);
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

      <div class="relative mx-auto w-full max-w-[1400px] px-4 py-12 sm:px-6 sm:py-14 lg:px-8">
        <div class="flex flex-col justify-between gap-7 lg:flex-row lg:items-end">
          <div class="max-w-3xl">
            <p class="text-sm font-semibold text-primary">Nexora account</p>

            <h1 class="mt-3 text-4xl font-semibold leading-[1.05] tracking-[-0.04em] sm:text-5xl">
              My purchases
            </h1>

            <p class="mt-4 max-w-2xl text-base leading-7 text-muted-foreground">
              Review seller orders, outstanding balances, payments and completed purchases.
            </p>
          </div>

          <Button
            variant="outline"
            class="w-full self-start rounded-full sm:w-auto lg:self-auto"
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

    <main class="mx-auto w-full max-w-[1400px] space-y-7 px-4 py-8 sm:space-y-8 sm:px-6 lg:px-8">
      <!-- PAGE SNAPSHOT -->
      <section
        class="overflow-hidden rounded-[1.5rem] border border-border bg-card shadow-sm sm:rounded-[1.75rem]"
      >
        <div
          class="flex flex-col gap-3 border-b border-border bg-muted/15 px-4 py-4 sm:flex-row sm:items-center sm:justify-between sm:px-5"
        >
          <div>
            <p class="text-sm font-semibold">Current results snapshot</p>

            <p class="mt-1 text-xs leading-5 text-muted-foreground">
              Values calculated from the purchases visible on this page.
            </p>
          </div>

          <span
            class="self-start rounded-full border border-border bg-background px-3 py-1 text-xs font-medium text-muted-foreground sm:self-auto"
          >
            Page {{ page }} of {{ totalPages }}
          </span>
        </div>

        <div class="grid grid-cols-1 gap-px bg-border min-[390px]:grid-cols-2 lg:grid-cols-4">
          <div class="min-w-0 bg-card p-4 sm:p-5">
            <div class="flex min-w-0 items-center gap-2">
              <Clock3 class="h-4 w-4 shrink-0 text-amber-600 dark:text-amber-400" />

              <p class="truncate text-xs text-muted-foreground">Awaiting payment</p>
            </div>

            <p class="mt-3 text-2xl font-semibold">
              {{ pageMetrics.awaitingPayment }}
            </p>
          </div>

          <div class="min-w-0 bg-card p-4 sm:p-5">
            <div class="flex min-w-0 items-center gap-2">
              <CheckCircle2 class="h-4 w-4 shrink-0 text-emerald-600 dark:text-emerald-400" />

              <p class="truncate text-xs text-muted-foreground">Paid orders</p>
            </div>

            <p class="mt-3 text-2xl font-semibold">
              {{ pageMetrics.paid }}
            </p>
          </div>

          <div class="min-w-0 bg-card p-4 sm:p-5">
            <div class="flex min-w-0 items-center gap-2">
              <WalletCards class="h-4 w-4 shrink-0 text-primary" />

              <p class="truncate text-xs text-muted-foreground">Outstanding</p>
            </div>

            <p
              class="mt-3 break-words text-lg font-semibold text-amber-600 dark:text-amber-400 sm:text-xl"
            >
              {{ formatCurrency(pageMetrics.outstanding) }}
            </p>
          </div>

          <div class="min-w-0 bg-card p-4 sm:p-5">
            <div class="flex min-w-0 items-center gap-2">
              <CreditCard class="h-4 w-4 shrink-0 text-blue-600 dark:text-blue-400" />

              <p class="truncate text-xs text-muted-foreground">Amount paid</p>
            </div>

            <p
              class="mt-3 break-words text-lg font-semibold text-emerald-600 dark:text-emerald-400 sm:text-xl"
            >
              {{ formatCurrency(pageMetrics.spent) }}
            </p>
          </div>
        </div>
      </section>

      <!-- FILTERS -->
      <section
        class="rounded-[1.5rem] border border-border bg-card p-4 shadow-sm sm:rounded-[1.75rem] sm:p-5"
      >
        <div class="grid gap-3 lg:grid-cols-[minmax(0,1fr)_220px_220px_170px_auto]">
          <div class="relative">
            <Search
              class="pointer-events-none absolute left-4 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
            />

            <Input
              v-model="search"
              class="h-11 rounded-full bg-muted/30 pl-11 pr-10"
              placeholder="Search seller or product..."
            />

            <Loader2
              v-if="isRefreshing"
              class="absolute right-4 top-1/2 h-4 w-4 -translate-y-1/2 animate-spin text-muted-foreground"
            />
          </div>

          <Select v-model="statusInvoice">
            <SelectTrigger class="h-11 rounded-full">
              <span
                class="truncate"
                :class="statusInvoice === 'all' ? 'text-muted-foreground' : 'text-foreground'"
              >
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
              <span class="truncate">
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

          <Button variant="ghost" class="h-11 w-full rounded-full lg:w-auto" @click="clearFilters">
            Clear filters
          </Button>
        </div>

        <div class="mt-4 border-t border-border pt-4">
          <p class="text-sm text-muted-foreground">
            {{ data?.meta.total ?? 0 }}

            {{ data?.meta.total === 1 ? 'purchase record' : 'purchase records' }}
          </p>
        </div>
      </section>

      <!-- LOADING -->
      <div v-if="isLoading" class="space-y-4">
        <div
          v-for="index in 4"
          :key="index"
          class="h-64 animate-pulse rounded-[1.5rem] bg-muted sm:h-48 sm:rounded-[1.75rem]"
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
            class="group min-w-0 overflow-hidden rounded-[1.5rem] border border-border bg-card shadow-sm transition hover:border-primary/25 hover:shadow-lg sm:rounded-[1.75rem]"
          >
            <div
              class="grid min-w-0 gap-5 p-4 sm:p-5 lg:grid-cols-[minmax(0,1fr)_160px_160px_auto] lg:items-center lg:gap-6"
            >
              <!-- ORDER IDENTITY -->
              <div class="flex min-w-0 gap-3 sm:gap-4">
                <div
                  class="flex h-11 w-11 shrink-0 items-center justify-center rounded-xl bg-primary/10 sm:h-12 sm:w-12 sm:rounded-2xl"
                >
                  <Building2 class="h-5 w-5 text-primary" />
                </div>

                <div class="min-w-0">
                  <div class="flex min-w-0 flex-wrap items-center gap-2">
                    <h2 class="min-w-0 truncate text-base font-semibold">
                      {{ purchase.seller_company }}
                    </h2>

                    <span
                      class="inline-flex shrink-0 rounded-full border px-2.5 py-1 text-xs font-semibold"
                      :class="getStatusClass(purchase.status_invoice)"
                    >
                      {{ getStatusLabel(purchase.status_invoice) }}
                    </span>
                  </div>

                  <div class="mt-2 flex flex-wrap items-center gap-2">
                    <span class="text-sm font-semibold"> #NX-{{ purchase.id }} </span>

                    <span class="h-1 w-1 rounded-full bg-muted-foreground/50" />

                    <span class="flex items-center gap-1.5 text-xs text-muted-foreground">
                      <CalendarDays class="h-3.5 w-3.5 shrink-0" />

                      {{ formatDate(purchase.created_at) }}
                    </span>
                  </div>

                  <p class="mt-2 text-xs leading-5 text-muted-foreground">
                    {{ purchase.item_count }}

                    {{
                      purchase.item_count === 1 ? 'item from this seller' : 'items from this seller'
                    }}
                  </p>
                </div>
              </div>

              <!-- FINANCIALS MOBILE/TABLET -->
              <div class="grid grid-cols-1 gap-3 min-[390px]:grid-cols-2 lg:contents">
                <div
                  class="min-w-0 rounded-xl border border-border bg-muted/15 p-4 lg:rounded-none lg:border-0 lg:bg-transparent lg:p-0"
                >
                  <p class="text-[10px] font-medium uppercase tracking-wide text-muted-foreground">
                    Order total
                  </p>

                  <p class="mt-1 break-words text-lg font-semibold tracking-tight">
                    {{ formatCurrency(purchase.total_amount) }}
                  </p>

                  <p class="mt-1 break-words text-xs text-muted-foreground">
                    Paid: {{ formatCurrency(purchase.paid_amount) }}
                  </p>
                </div>

                <div
                  class="min-w-0 rounded-xl border border-border bg-muted/15 p-4 lg:rounded-none lg:border-0 lg:bg-transparent lg:p-0"
                >
                  <p class="text-[10px] font-medium uppercase tracking-wide text-muted-foreground">
                    Outstanding
                  </p>

                  <p
                    class="mt-1 break-words text-lg font-semibold tracking-tight"
                    :class="
                      purchase.remaining_amount > 0
                        ? 'text-amber-600 dark:text-amber-400'
                        : 'text-emerald-600 dark:text-emerald-400'
                    "
                  >
                    {{ formatCurrency(purchase.remaining_amount) }}
                  </p>

                  <p class="mt-1 text-xs leading-5 text-muted-foreground">
                    {{
                      purchase.remaining_amount > 0 ? 'Payment still required' : 'Order fully paid'
                    }}
                  </p>
                </div>
              </div>

              <!-- ACTIONS -->
              <div
                class="grid w-full gap-2 min-[390px]:grid-cols-2 lg:flex lg:w-auto lg:flex-wrap lg:justify-end"
              >
                <Button
                  variant="outline"
                  class="w-full rounded-full"
                  @click="openPurchase(purchase)"
                >
                  <Eye class="mr-2 h-4 w-4" />
                  Details
                </Button>

                <Button
                  v-if="canPay(purchase)"
                  class="w-full rounded-full"
                  @click="openPurchasePayment(purchase)"
                >
                  <WalletCards class="mr-2 h-4 w-4" />
                  Pay now
                </Button>
              </div>
            </div>

            <!-- PAYMENT PROGRESS -->
            <div class="border-t border-border bg-muted/10 px-4 py-3 sm:px-5">
              <div class="flex items-center justify-between gap-4 text-xs">
                <span class="text-muted-foreground"> Payment progress </span>

                <span class="font-semibold"> {{ Math.round(getPaymentProgress(purchase)) }}% </span>
              </div>

              <div class="mt-2 h-1.5 overflow-hidden rounded-full bg-muted">
                <div
                  class="h-full rounded-full transition-all duration-500"
                  :class="purchase.remaining_amount <= 0 ? 'bg-emerald-500' : 'bg-primary'"
                  :style="{
                    width: `${getPaymentProgress(purchase)}%`,
                  }"
                />
              </div>
            </div>
          </article>
        </div>

        <DataPagination
          class="mt-8 w-full"
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
