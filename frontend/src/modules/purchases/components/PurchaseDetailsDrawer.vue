<script setup lang="ts">
import { computed, toRef } from 'vue';

import {
  Building2,
  CalendarDays,
  CircleDollarSign,
  CreditCard,
  FileText,
  Loader2,
  PackageSearch,
  ReceiptText,
  WalletCards,
} from 'lucide-vue-next';

import { Button } from '@/components/ui/button';

import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
} from '@/components/ui/sheet';

import EmptyState from '@/shared/components/erp/EmptyState.vue';

import { useInvoiceDetail } from '@/modules/invoices/composables/useInvoiceDetail';
import type { PurchaseListItem } from '../types/purchase.types';

const props = defineProps<{
  open: boolean;
  purchase: PurchaseListItem | null;
}>();

const emit = defineEmits<{
  'update:open': [value: boolean];

  viewItems: [invoiceId: number];
  viewPayments: [invoiceId: number];
  pay: [invoiceId: number];
}>();

const purchaseId = computed<number | null>(() => {
  return props.purchase?.id ?? null;
});

const { data: invoice, isLoading, isFetching, isError, refetch } = useInvoiceDetail(purchaseId);

const isRefreshing = computed(() => {
  return isFetching.value && !isLoading.value;
});

const remainingAmount = computed(() => {
  if (!invoice.value) return 0;

  return Math.max(invoice.value.total_amount - invoice.value.paid_amount, 0);
});

const paymentProgress = computed(() => {
  if (!invoice.value) return 0;

  if (invoice.value.total_amount <= 0) {
    return 0;
  }

  return Math.min(Math.round((invoice.value.paid_amount / invoice.value.total_amount) * 100), 100);
});

const canPay = computed(() => {
  if (!invoice.value) return false;

  return invoice.value.status_invoice === 'pending' && remainingAmount.value > 0;
});

function formatCurrency(value: number) {
  return new Intl.NumberFormat('es-AR', {
    style: 'currency',
    currency: 'ARS',
    minimumFractionDigits: 2,
  }).format(value);
}

function formatDate(value?: string | null) {
  if (!value) return '—';

  return new Intl.DateTimeFormat('es-AR', {
    day: '2-digit',
    month: 'short',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  }).format(new Date(value));
}

function getStatusLabel(status?: string) {
  const labels: Record<string, string> = {
    draft: 'Draft',
    pending: 'Awaiting payment',
    paid: 'Paid',
    cancelled: 'Cancelled',
  };

  return status ? (labels[status] ?? status) : 'Unknown';
}

function getStatusClass(status?: string) {
  const classes: Record<string, string> = {
    draft: 'border-border bg-muted text-muted-foreground',

    pending: 'border-amber-500/20 bg-amber-500/10 text-amber-600 dark:text-amber-400',

    paid: 'border-emerald-500/20 bg-emerald-500/10 text-emerald-600 dark:text-emerald-400',

    cancelled: 'border-destructive/20 bg-destructive/10 text-destructive',
  };

  return classes[status ?? ''] ?? 'border-border bg-muted text-muted-foreground';
}

function handleViewItems() {
  if (!invoice.value) return;

  emit('viewItems', invoice.value.id);
}

function handleViewPayments() {
  if (!invoice.value) return;

  emit('viewPayments', invoice.value.id);
}

function handlePay() {
  if (!invoice.value || !canPay.value) return;

  emit('pay', invoice.value.id);
}
</script>

<template>
  <Sheet :open="open" @update:open="emit('update:open', $event)">
    <SheetContent class="w-full overflow-y-auto p-0 sm:max-w-2xl">
      <!-- HEADER -->
      <div class="border-b border-border bg-muted/20 px-6 py-5">
        <SheetHeader class="text-left">
          <div class="flex items-start justify-between gap-4">
            <div class="flex min-w-0 items-center gap-3">
              <div
                class="flex h-11 w-11 shrink-0 items-center justify-center rounded-xl border border-border bg-background shadow-sm"
              >
                <ReceiptText class="h-5 w-5 text-muted-foreground" />
              </div>

              <div class="min-w-0">
                <SheetTitle> Purchase details </SheetTitle>

                <SheetDescription class="mt-1">
                  Review the supplier order and its financial status.
                </SheetDescription>
              </div>
            </div>

            <Button
              v-if="invoice"
              type="button"
              variant="ghost"
              size="icon"
              :disabled="isRefreshing"
              @click="refetch()"
            >
              <Loader2 v-if="isRefreshing" class="h-4 w-4 animate-spin" />

              <FileText v-else class="h-4 w-4" />
            </Button>
          </div>
        </SheetHeader>
      </div>

      <!-- LOADING -->
      <div v-if="isLoading" class="flex min-h-96 items-center justify-center p-8">
        <div class="text-center">
          <Loader2 class="mx-auto h-7 w-7 animate-spin text-muted-foreground" />

          <p class="mt-3 text-sm text-muted-foreground">Loading purchase details...</p>
        </div>
      </div>

      <!-- ERROR -->
      <div v-else-if="isError" class="p-6">
        <EmptyState
          title="Unable to load purchase"
          description="There was a problem retrieving the invoice details."
          :icon="ReceiptText"
        />

        <div class="mt-4 flex justify-center">
          <Button variant="outline" @click="refetch()"> Try again </Button>
        </div>
      </div>

      <!-- CONTENT -->
      <div v-else-if="invoice && purchase" class="space-y-6 p-6">
        <!-- SUPPLIER -->
        <section
          class="relative overflow-hidden rounded-2xl border border-border bg-card p-5 shadow-sm"
        >
          <div
            class="pointer-events-none absolute inset-0 bg-gradient-to-br from-primary/5 via-transparent to-muted/40"
          />

          <div class="relative flex items-start justify-between gap-4">
            <div class="flex min-w-0 items-center gap-4">
              <div
                class="flex h-12 w-12 shrink-0 items-center justify-center rounded-2xl border border-border bg-background shadow-sm"
              >
                <Building2 class="h-5 w-5 text-primary" />
              </div>

              <div class="min-w-0">
                <p class="text-xs font-medium uppercase tracking-wide text-muted-foreground">
                  Supplier
                </p>

                <h3 class="mt-1 truncate text-lg font-semibold text-foreground">
                  {{ purchase.seller_company }}
                </h3>

                <p class="mt-1 text-xs text-muted-foreground">B2B marketplace purchase</p>
              </div>
            </div>

            <span
              class="shrink-0 rounded-full border px-3 py-1 text-xs font-semibold"
              :class="getStatusClass(invoice.status_invoice)"
            >
              {{ getStatusLabel(invoice.status_invoice) }}
            </span>
          </div>
        </section>

        <!-- SUMMARY CARDS -->
        <section class="grid gap-3 sm:grid-cols-2">
          <article class="rounded-2xl border border-border bg-muted/20 p-4">
            <div class="flex items-center gap-3">
              <div class="rounded-xl bg-primary/10 p-2.5">
                <CircleDollarSign class="h-4 w-4 text-primary" />
              </div>

              <div>
                <p class="text-xs text-muted-foreground">Purchase total</p>

                <p class="mt-1 text-lg font-semibold text-foreground">
                  {{ formatCurrency(invoice.total_amount) }}
                </p>
              </div>
            </div>
          </article>

          <article class="rounded-2xl border border-border bg-muted/20 p-4">
            <div class="flex items-center gap-3">
              <div class="rounded-xl bg-emerald-500/10 p-2.5">
                <CreditCard class="h-4 w-4 text-emerald-600 dark:text-emerald-400" />
              </div>

              <div>
                <p class="text-xs text-muted-foreground">Amount paid</p>

                <p class="mt-1 text-lg font-semibold text-foreground">
                  {{ formatCurrency(invoice.paid_amount) }}
                </p>
              </div>
            </div>
          </article>

          <article class="rounded-2xl border border-border bg-muted/20 p-4">
            <div class="flex items-center gap-3">
              <div class="rounded-xl bg-amber-500/10 p-2.5">
                <WalletCards class="h-4 w-4 text-amber-600 dark:text-amber-400" />
              </div>

              <div>
                <p class="text-xs text-muted-foreground">Remaining balance</p>

                <p class="mt-1 text-lg font-semibold text-foreground">
                  {{ formatCurrency(remainingAmount) }}
                </p>
              </div>
            </div>
          </article>

          <article class="rounded-2xl border border-border bg-muted/20 p-4">
            <div class="flex items-center gap-3">
              <div class="rounded-xl bg-blue-500/10 p-2.5">
                <CalendarDays class="h-4 w-4 text-blue-600 dark:text-blue-400" />
              </div>

              <div>
                <p class="text-xs text-muted-foreground">Purchase date</p>

                <p class="mt-1 text-sm font-semibold text-foreground">
                  {{ formatDate(invoice.created_at) }}
                </p>
              </div>
            </div>
          </article>
        </section>

        <!-- PAYMENT PROGRESS -->
        <section class="rounded-2xl border border-border bg-card p-5 shadow-sm">
          <div class="flex items-end justify-between gap-4">
            <div>
              <p class="text-sm font-medium text-foreground">Payment progress</p>

              <p class="mt-1 text-xs text-muted-foreground">
                {{ paymentProgress }}% of the invoice has been paid.
              </p>
            </div>

            <p class="text-sm font-semibold text-foreground">{{ paymentProgress }}%</p>
          </div>

          <div class="mt-4 h-2 overflow-hidden rounded-full bg-muted">
            <div
              class="h-full rounded-full bg-primary transition-all duration-500"
              :style="{
                width: `${paymentProgress}%`,
              }"
            />
          </div>

          <div class="mt-3 flex items-center justify-between text-xs text-muted-foreground">
            <span> {{ formatCurrency(invoice.paid_amount) }} paid </span>

            <span> {{ formatCurrency(remainingAmount) }} remaining </span>
          </div>
        </section>

        <!-- FINANCIAL BREAKDOWN -->
        <section class="rounded-2xl border border-border bg-card p-5 shadow-sm">
          <div class="mb-4 flex items-center gap-2">
            <ReceiptText class="h-4 w-4 text-muted-foreground" />

            <h3 class="text-sm font-semibold text-foreground">Financial summary</h3>
          </div>

          <div class="space-y-3">
            <div class="flex items-center justify-between">
              <span class="text-sm text-muted-foreground"> Subtotal </span>

              <span class="text-sm font-medium text-foreground">
                {{ formatCurrency(invoice.subtotal) }}
              </span>
            </div>

            <div class="flex items-center justify-between">
              <span class="text-sm text-muted-foreground"> Taxes </span>

              <span class="text-sm font-medium text-foreground">
                {{ formatCurrency(invoice.taxes) }}
              </span>
            </div>

            <div class="flex items-center justify-between border-t border-border pt-4">
              <span class="text-sm font-semibold text-foreground"> Total </span>

              <span class="text-xl font-semibold tracking-tight text-foreground">
                {{ formatCurrency(invoice.total_amount) }}
              </span>
            </div>
          </div>
        </section>

        <!-- ACTIONS -->
        <section class="grid gap-3 sm:grid-cols-2">
          <Button type="button" variant="outline" @click="handleViewItems">
            <PackageSearch class="mr-2 h-4 w-4" />
            View purchased items
          </Button>

          <Button type="button" variant="outline" @click="handleViewPayments">
            <CreditCard class="mr-2 h-4 w-4" />
            Payment history
          </Button>

          <Button v-if="canPay" type="button" class="sm:col-span-2" @click="handlePay">
            <WalletCards class="mr-2 h-4 w-4" />
            Pay outstanding balance
          </Button>
        </section>
      </div>
    </SheetContent>
  </Sheet>
</template>
