<script setup lang="ts">
import { computed, toRef } from 'vue';

import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
} from '@/components/ui/sheet';

import { useInvoiceDetail } from '../composables/useInvoiceDetail';

const props = defineProps<{
  open: boolean;
  invoiceId: number | null;
}>();

const emit = defineEmits<{
  'update:open': [value: boolean];
}>();

const invoiceIdRef = toRef(props, 'invoiceId');

const { data: invoice, isLoading, isError } = useInvoiceDetail(invoiceIdRef);

const remainingAmount = computed(() => {
  if (!invoice.value) return 0;
  return invoice.value.total_amount - invoice.value.paid_amount;
});

function formatCurrency(value: number) {
  return `$${value.toFixed(2)}`;
}
</script>

<template>
  <Sheet :open="open" @update:open="emit('update:open', $event)">
    <SheetContent class="w-full sm:max-w-2xl">
      <SheetHeader>
        <SheetTitle> Invoice details </SheetTitle>

        <SheetDescription> Financial snapshot and workflow status. </SheetDescription>
      </SheetHeader>

      <div
        v-if="isLoading"
        class="mt-6 rounded-xl border border-border bg-muted/30 p-6 text-sm text-muted-foreground"
      >
        Loading invoice...
      </div>

      <div
        v-else-if="isError"
        class="mt-6 rounded-xl border border-border bg-muted/30 p-6 text-sm text-destructive"
      >
        Unable to load invoice.
      </div>

      <div v-else-if="invoice" class="mt-6 space-y-6">
        <div class="flex items-start justify-between gap-4">
          <div>
            <p class="text-sm text-muted-foreground">Invoice</p>

            <h2 class="mt-1 text-2xl font-semibold tracking-tight">#INV-{{ invoice.id }}</h2>
          </div>

          <span class="rounded-full border border-border bg-muted px-3 py-1 text-xs font-medium">
            {{ invoice.status_invoice }}
          </span>
        </div>

        <div class="grid gap-4 sm:grid-cols-2">
          <div class="rounded-xl border border-border bg-card p-4">
            <p class="text-xs text-muted-foreground">Subtotal</p>
            <p class="mt-2 text-lg font-semibold">
              {{ formatCurrency(invoice.subtotal) }}
            </p>
          </div>

          <div class="rounded-xl border border-border bg-card p-4">
            <p class="text-xs text-muted-foreground">Taxes</p>
            <p class="mt-2 text-lg font-semibold">
              {{ formatCurrency(invoice.taxes) }}
            </p>
          </div>

          <div class="rounded-xl border border-border bg-card p-4">
            <p class="text-xs text-muted-foreground">Total</p>
            <p class="mt-2 text-lg font-semibold">
              {{ formatCurrency(invoice.total_amount) }}
            </p>
          </div>

          <div class="rounded-xl border border-border bg-card p-4">
            <p class="text-xs text-muted-foreground">Paid</p>
            <p class="mt-2 text-lg font-semibold">
              {{ formatCurrency(invoice.paid_amount) }}
            </p>
          </div>
        </div>

        <div class="rounded-xl border border-border bg-muted/40 p-4">
          <p class="text-xs text-muted-foreground">Remaining balance</p>

          <p class="mt-2 text-xl font-semibold">
            {{ formatCurrency(remainingAmount) }}
          </p>
        </div>

        <div class="grid gap-4 sm:grid-cols-2">
          <div class="rounded-xl border border-border bg-card p-4">
            <p class="text-xs text-muted-foreground">Buyer client</p>
            <p class="mt-2 text-sm">
              {{ invoice.buyer_name }}
            </p>
          </div>

          <div class="rounded-xl border border-border bg-card p-4">
            <p class="text-xs text-muted-foreground">Seller company</p>
            <p class="mt-2 text-sm">
              {{ invoice.seller_company }}
            </p>
          </div>
        </div>
      </div>
    </SheetContent>
  </Sheet>
</template>
