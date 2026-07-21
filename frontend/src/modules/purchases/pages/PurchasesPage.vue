<script setup lang="ts">
import { computed, ref, watch } from 'vue';

import {
  ArrowDownUp,
  Building2,
  CalendarDays,
  CreditCard,
  Eye,
  MoreHorizontal,
  ReceiptText,
  RefreshCw,
  ShoppingBag,
  WalletCards,
  Search,
} from 'lucide-vue-next';

import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';

import PageContainer from '@/shared/components/erp/PageContainer.vue';
import PageHeader from '@/shared/components/erp/PageHeader.vue';
import SectionCard from '@/shared/components/erp/SectionCard.vue';
import DataTable from '@/shared/components/erp/DataTable.vue';
import DataPagination from '@/shared/components/erp/DataPagination.vue';
import EmptyState from '@/shared/components/erp/EmptyState.vue';

import PurchaseDetailsDrawer from '../components/PurchaseDetailsDrawer.vue';

import { usePurchases } from '../composables/usePurchases';

import type { PurchaseListItem } from '../types/purchase.types';

import type { InvoiceStatus } from '@/modules/invoices/types/invoice.types';

const page = ref(1);
const limit = ref(10);

const statusInvoice = ref('all');

const supplierSearch = ref('');

const sortColumn = ref<'created_at' | 'total_amount' | 'paid_amount' | 'status_invoice'>(
  'created_at'
);

const order = ref<'asc' | 'desc'>('desc');

watch([supplierSearch, statusInvoice, sortColumn, order], () => {
  page.value = 1;
});

const purchaseParams = computed(() => ({
  page: page.value,
  limit: limit.value,

  status_invoice:
    statusInvoice.value !== 'all' ? (statusInvoice.value as InvoiceStatus) : undefined,

  seller_company: supplierSearch.value.trim().length >= 2 ? supplierSearch.value.trim() : undefined,

  sort_column: sortColumn.value,
  order: order.value,
}));

const selectedPurchase = ref<PurchaseListItem | null>(null);

const purchaseDetailsOpen = ref(false);

const { data, isLoading, isFetching, isError, refetch } = usePurchases(purchaseParams);

const purchases = computed(() => data.value?.items ?? []);

const totalPages = computed(() => data.value?.meta.total_pages ?? 1);

const isRefreshing = computed(() => isFetching.value && !isLoading.value);

const columns = [
  {
    key: 'seller_company',
    label: 'Supplier',
  },
  {
    key: 'status_invoice',
    label: 'Status',
  },
  {
    key: 'created_at',
    label: 'Purchase date',
  },
  {
    key: 'total_amount',
    label: 'Total',
  },
  {
    key: 'paid_amount',
    label: 'Paid',
  },
  {
    key: 'remaining_amount',
    label: 'Remaining',
  },
  {
    key: 'actions',
    label: '',
  },
];

const rows = computed(() =>
  purchases.value.map((purchase) => ({
    ...purchase,

    remaining_amount: Math.max(purchase.total_amount - purchase.paid_amount, 0),

    actions: purchase.id,
  }))
);

/*
 * Estos KPIs representan únicamente los registros
 * de la página actualmente cargada.
 */
const pageMetrics = computed(() => {
  const items = purchases.value;

  return {
    pending: items.filter((purchase) => purchase.status_invoice === 'pending').length,

    paid: items.filter((purchase) => purchase.status_invoice === 'paid').length,

    outstanding: items.reduce(
      (total, purchase) => total + Math.max(purchase.total_amount - purchase.paid_amount, 0),
      0
    ),

    spent: items.reduce((total, purchase) => total + purchase.paid_amount, 0),
  };
});

function formatCurrency(value: unknown) {
  return new Intl.NumberFormat('es-AR', {
    style: 'currency',
    currency: 'ARS',
    minimumFractionDigits: 2,
  }).format(Number(value));
}

function formatDate(value: unknown) {
  return new Intl.DateTimeFormat('es-AR', {
    day: '2-digit',
    month: 'short',
    year: 'numeric',
  }).format(new Date(String(value)));
}

function getStatusLabel(value: unknown) {
  const status = String(value);

  const labels: Record<string, string> = {
    draft: 'Draft',
    pending: 'Awaiting payment',
    paid: 'Paid',
    cancelled: 'Cancelled',
  };

  return labels[status] ?? status;
}

function getStatusClass(value: unknown) {
  const status = String(value);

  const classes: Record<string, string> = {
    draft: 'border-border bg-muted text-muted-foreground',

    pending: 'border-amber-500/20 bg-amber-500/10 text-amber-600 dark:text-amber-400',

    paid: 'border-emerald-500/20 bg-emerald-500/10 text-emerald-600 dark:text-emerald-400',

    cancelled: 'border-destructive/20 bg-destructive/10 text-destructive',
  };

  return classes[status] ?? 'border-border bg-muted text-muted-foreground';
}

const purchaseInitialAction = ref<'overview' | 'checkout' | 'payment'>('overview');

function canPayPurchase(purchase: PurchaseListItem) {
  return purchase.status_invoice === 'pending' && purchase.paid_amount < purchase.total_amount;
}

function openPurchaseDetails(purchase: PurchaseListItem) {
  selectedPurchase.value = purchase;
  purchaseInitialAction.value = 'overview';
  purchaseDetailsOpen.value = true;
}

function openPurchaseItems(purchase: PurchaseListItem) {
  console.log('Open purchase items:', purchase.id);
}

function openPurchasePayments(purchase: PurchaseListItem) {
  console.log('Open purchase payments:', purchase.id);
}

function openPurchasePayment(purchase: PurchaseListItem) {
  console.log('Pay purchase:', purchase.id);
}

function handleSort(columnKey: string) {
  const supportedColumns = ['created_at', 'total_amount', 'paid_amount', 'status_invoice'] as const;

  if (!supportedColumns.includes(columnKey as (typeof supportedColumns)[number])) {
    return;
  }

  if (sortColumn.value === columnKey) {
    order.value = order.value === 'asc' ? 'desc' : 'asc';

    return;
  }

  sortColumn.value = columnKey as typeof sortColumn.value;

  order.value = 'desc';
}

function clearFilters() {
  statusInvoice.value = 'all';
  sortColumn.value = 'created_at';
  supplierSearch.value = '';
  order.value = 'desc';
  page.value = 1;
}

function openPurchaseForCheckout(purchase: PurchaseListItem) {
  selectedPurchase.value = purchase;
  purchaseInitialAction.value = 'checkout';
  purchaseDetailsOpen.value = true;
}

function openPurchaseForPayment(purchase: PurchaseListItem) {
  selectedPurchase.value = purchase;
  purchaseInitialAction.value = 'payment';
  purchaseDetailsOpen.value = true;
}
</script>

<template>
  <PageContainer>
    <PageHeader
      title="Purchases"
      description="Review supplier orders, track outstanding balances and manage buyer-side payments."
    />

    <!-- KPIs -->
    <section class="grid gap-4 md:grid-cols-2 xl:grid-cols-4">
      <article class="rounded-2xl border border-border bg-card p-5 shadow-sm">
        <div class="flex items-start justify-between">
          <div>
            <p class="text-sm text-muted-foreground">Awaiting payment</p>

            <p class="mt-2 text-3xl font-semibold tracking-tight text-foreground">
              {{ pageMetrics.pending }}
            </p>
          </div>

          <div class="rounded-xl bg-amber-500/10 p-3">
            <ReceiptText class="h-5 w-5 text-amber-600 dark:text-amber-400" />
          </div>
        </div>

        <p class="mt-4 text-xs text-muted-foreground">On this results page</p>
      </article>

      <article class="rounded-2xl border border-border bg-card p-5 shadow-sm">
        <div class="flex items-start justify-between">
          <div>
            <p class="text-sm text-muted-foreground">Completed purchases</p>

            <p class="mt-2 text-3xl font-semibold tracking-tight text-foreground">
              {{ pageMetrics.paid }}
            </p>
          </div>

          <div class="rounded-xl bg-emerald-500/10 p-3">
            <ShoppingBag class="h-5 w-5 text-emerald-600 dark:text-emerald-400" />
          </div>
        </div>

        <p class="mt-4 text-xs text-muted-foreground">Fully paid on this page</p>
      </article>

      <article class="rounded-2xl border border-border bg-card p-5 shadow-sm">
        <div class="flex items-start justify-between gap-4">
          <div class="min-w-0">
            <p class="text-sm text-muted-foreground">Outstanding balance</p>

            <p class="mt-2 truncate text-2xl font-semibold tracking-tight text-foreground">
              {{ formatCurrency(pageMetrics.outstanding) }}
            </p>
          </div>

          <div class="rounded-xl bg-primary/10 p-3">
            <WalletCards class="h-5 w-5 text-primary" />
          </div>
        </div>

        <p class="mt-4 text-xs text-muted-foreground">Remaining on visible purchases</p>
      </article>

      <article class="rounded-2xl border border-border bg-card p-5 shadow-sm">
        <div class="flex items-start justify-between gap-4">
          <div class="min-w-0">
            <p class="text-sm text-muted-foreground">Amount paid</p>

            <p class="mt-2 truncate text-2xl font-semibold tracking-tight text-foreground">
              {{ formatCurrency(pageMetrics.spent) }}
            </p>
          </div>

          <div class="rounded-xl bg-blue-500/10 p-3">
            <CreditCard class="h-5 w-5 text-blue-600 dark:text-blue-400" />
          </div>
        </div>

        <p class="mt-4 text-xs text-muted-foreground">Paid across visible purchases</p>
      </article>
    </section>

    <SectionCard
      title="Purchase records"
      description="Invoices where your company participates as the buyer."
    >
      <template #actions>
        <Button variant="outline" size="sm" :disabled="isRefreshing" @click="refetch()">
          <RefreshCw
            class="mr-2 h-4 w-4"
            :class="{
              'animate-spin': isRefreshing,
            }"
          />

          Refresh
        </Button>
      </template>

      <!-- Filters -->
      <div class="mb-5 space-y-4">
        <div
          class="grid gap-3 md:grid-cols-2 xl:grid-cols-[minmax(260px,1fr)_220px_220px_180px_auto]"
        >
          <div class="relative">
            <Search
              class="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
            />

            <Input v-model="supplierSearch" class="pl-9" placeholder="Search supplier..." />
          </div>

          <Select v-model="statusInvoice">
            <SelectTrigger>
              <span :class="statusInvoice === 'all' ? 'text-muted-foreground' : 'text-foreground'">
                {{ statusInvoice === 'all' ? 'All statuses' : getStatusLabel(statusInvoice) }}
              </span>
            </SelectTrigger>

            <SelectContent>
              <SelectItem value="all"> All statuses </SelectItem>

              <SelectItem value="draft"> Draft </SelectItem>

              <SelectItem value="pending"> Awaiting payment </SelectItem>

              <SelectItem value="paid"> Paid </SelectItem>

              <SelectItem value="cancelled"> Cancelled </SelectItem>
            </SelectContent>
          </Select>

          <Select v-model="sortColumn">
            <SelectTrigger>
              <SelectValue placeholder="Sort by" />
            </SelectTrigger>

            <SelectContent>
              <SelectItem value="created_at"> Purchase date </SelectItem>

              <SelectItem value="total_amount"> Total amount </SelectItem>

              <SelectItem value="paid_amount"> Paid amount </SelectItem>

              <SelectItem value="status_invoice"> Status </SelectItem>
            </SelectContent>
          </Select>

          <Select v-model="order">
            <SelectTrigger>
              <SelectValue placeholder="Order" />
            </SelectTrigger>

            <SelectContent>
              <SelectItem value="desc"> Descending </SelectItem>

              <SelectItem value="asc"> Ascending </SelectItem>
            </SelectContent>
          </Select>

          <Button variant="ghost" class="justify-self-start" @click="clearFilters">
            Clear filters
          </Button>
        </div>

        <div class="flex items-center justify-between">
          <p class="text-sm text-muted-foreground">
            {{ data?.meta.total ?? 0 }}
            purchase records
          </p>

          <div class="flex items-center gap-2 text-xs text-muted-foreground">
            <ArrowDownUp class="h-3.5 w-3.5" />
            Buyer-side records
          </div>
        </div>
      </div>

      <!-- Loading -->
      <div
        v-if="isLoading"
        class="rounded-xl border border-border bg-muted/30 p-8 text-center text-sm text-muted-foreground"
      >
        Loading purchases...
      </div>

      <!-- Error -->
      <EmptyState
        v-else-if="isError"
        title="Unable to load purchases"
        description="There was a problem connecting to the purchases API."
        :icon="ShoppingBag"
      />

      <!-- Empty -->
      <EmptyState
        v-else-if="rows.length === 0"
        title="No purchases found"
        description="Completed marketplace checkouts will appear here."
        :icon="ShoppingBag"
      />

      <!-- Table -->
      <template v-else>
        <DataTable
          :columns="columns"
          :rows="rows"
          :sortable-columns="['status_invoice', 'created_at', 'total_amount', 'paid_amount']"
          :sort-column="sortColumn"
          :sort-order="order"
          @sort="handleSort"
        >
          <template #cell-seller_company="{ row }">
            <div class="flex min-w-0 items-center gap-3">
              <div
                class="flex h-9 w-9 shrink-0 items-center justify-center rounded-xl border border-border bg-muted/40"
              >
                <Building2 class="h-4 w-4 text-muted-foreground" />
              </div>

              <div class="min-w-0">
                <p class="truncate text-sm font-medium text-foreground">
                  {{ row.seller_company }}
                </p>

                <p class="text-xs text-muted-foreground">B2B supplier</p>
              </div>
            </div>
          </template>

          <template #cell-status_invoice="{ value }">
            <span
              class="inline-flex rounded-full border px-2.5 py-1 text-xs font-semibold"
              :class="getStatusClass(value)"
            >
              {{ getStatusLabel(value) }}
            </span>
          </template>

          <template #cell-created_at="{ value }">
            <div class="flex items-center gap-2 text-sm text-muted-foreground">
              <CalendarDays class="h-4 w-4" />
              {{ formatDate(value) }}
            </div>
          </template>

          <template #cell-total_amount="{ value }">
            <span class="font-semibold text-foreground">
              {{ formatCurrency(value) }}
            </span>
          </template>

          <template #cell-paid_amount="{ value }">
            <span class="text-sm text-emerald-600 dark:text-emerald-400">
              {{ formatCurrency(value) }}
            </span>
          </template>

          <template #cell-remaining_amount="{ value }">
            <span
              class="text-sm font-medium"
              :class="
                Number(value) > 0 ? 'text-amber-600 dark:text-amber-400' : 'text-muted-foreground'
              "
            >
              {{ formatCurrency(value) }}
            </span>
          </template>

          <template #cell-actions="{ row }">
            <DropdownMenu>
              <DropdownMenuTrigger as-child>
                <Button variant="ghost" size="icon" class="h-8 w-8">
                  <MoreHorizontal class="h-4 w-4" />
                </Button>
              </DropdownMenuTrigger>

              <DropdownMenuContent align="end" class="w-52">
                <DropdownMenuItem @click="openPurchaseDetails(row as PurchaseListItem)">
                  <Eye class="mr-2 h-4 w-4" />
                  View purchase
                </DropdownMenuItem>

                <template v-if="row.status_invoice === 'draft'">
                  <DropdownMenuSeparator />

                  <DropdownMenuItem @click="openPurchaseForCheckout(row as PurchaseListItem)">
                    <ShoppingBag class="mr-2 h-4 w-4" />
                    Continue checkout
                  </DropdownMenuItem>
                </template>

                <template v-else-if="canPayPurchase(row as PurchaseListItem)">
                  <DropdownMenuSeparator />

                  <DropdownMenuItem @click="openPurchaseForPayment(row as PurchaseListItem)">
                    <WalletCards class="mr-2 h-4 w-4" />
                    Pay invoice
                  </DropdownMenuItem>
                </template>
              </DropdownMenuContent>
            </DropdownMenu>
          </template>
        </DataTable>

        <DataPagination
          class="mt-5"
          :page="page"
          :total-pages="totalPages"
          :total="data?.meta.total"
          @previous="page--"
          @next="page++"
        />
      </template>
    </SectionCard>

    <PurchaseDetailsDrawer
      v-model:open="purchaseDetailsOpen"
      :purchase="selectedPurchase"
      :initial-action="purchaseInitialAction"
    />
  </PageContainer>
</template>
