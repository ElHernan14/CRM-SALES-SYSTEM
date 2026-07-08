<script setup lang="ts">
import { computed, ref, toRef, watch } from 'vue';
import { CreditCard } from 'lucide-vue-next';

import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
} from '@/components/ui/sheet';

import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';

import DataTable from '@/shared/components/erp/DataTable.vue';
import DataPagination from '@/shared/components/erp/DataPagination.vue';
import EmptyState from '@/shared/components/erp/EmptyState.vue';

import { useInvoicePayments } from '../composables/useInvoicePayments';

import type { PaymentMethod } from '../types/invoice.types';

const props = defineProps<{
  open: boolean;
  invoiceId: number | null;
}>();

const emit = defineEmits<{
  'update:open': [value: boolean];
}>();

const invoiceIdRef = toRef(props, 'invoiceId');

const page = ref(1);
const limit = ref(10);
const paymentMethod = defineModel<PaymentMethod>('paymentMethod');
const sortColumn = ref('created_at');
const order = ref<'asc' | 'desc'>('desc');

watch([paymentMethod, sortColumn, order], () => {
  page.value = 1;
});

const params = computed(() => ({
  page: page.value,
  limit: limit.value,
  payment_method: paymentMethod.value || undefined,
  sort_column: sortColumn.value as 'created_at' | 'amount',
  order: order.value,
}));

const { data, isLoading, isError } = useInvoicePayments(invoiceIdRef, params);

const rows = computed(() => {
  return (
    data.value?.items.map((payment) => ({
      id: payment.id,
      amount: payment.amount,
      payment_method: payment.payment_method,
      paid_by_user_id: payment.paid_by_user_id,
      created_at: payment.created_at,
    })) ?? []
  );
});

const columns = [
  { key: 'id', label: 'Payment' },
  { key: 'amount', label: 'Amount' },
  { key: 'payment_method', label: 'Method' },
  { key: 'paid_by_user_id', label: 'Paid by' },
  { key: 'created_at', label: 'Created' },
];

const totalPages = computed(() => data.value?.meta.total_pages ?? 1);

function formatCurrency(value: unknown) {
  return `$${Number(value).toFixed(2)}`;
}

function formatDate(value: unknown) {
  return new Date(String(value)).toLocaleDateString();
}

function handleSort(columnKey: string) {
  if (sortColumn.value === columnKey) {
    order.value = order.value === 'asc' ? 'desc' : 'asc';
    return;
  }

  sortColumn.value = columnKey;
  order.value = 'desc';
}
</script>

<template>
  <Sheet :open="open" @update:open="emit('update:open', $event)">
    <SheetContent class="w-full overflow-y-auto sm:max-w-3xl">
      <SheetHeader>
        <SheetTitle> Invoice payments </SheetTitle>

        <SheetDescription> Payments registered for this sales invoice. </SheetDescription>
      </SheetHeader>

      <div class="mt-6 space-y-5">
        <div class="grid gap-3 sm:grid-cols-3">
          <Select v-model="paymentMethod">
            <SelectTrigger>
              <SelectValue placeholder="Payment method" />
            </SelectTrigger>

            <SelectContent>
              <SelectItem value="cash"> Cash </SelectItem>

              <SelectItem value="transfer"> Transfer </SelectItem>

              <SelectItem value="card"> Card </SelectItem>
            </SelectContent>
          </Select>

          <Select v-model="sortColumn">
            <SelectTrigger>
              <SelectValue placeholder="Sort by" />
            </SelectTrigger>

            <SelectContent>
              <SelectItem value="created_at"> Created date </SelectItem>

              <SelectItem value="amount"> Amount </SelectItem>
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
        </div>

        <div
          v-if="isLoading"
          class="rounded-xl border border-border bg-muted/30 p-6 text-sm text-muted-foreground"
        >
          Loading payments...
        </div>

        <EmptyState
          v-else-if="isError"
          title="Unable to load payments"
          description="There was a problem connecting to the payments API."
          :icon="CreditCard"
        />

        <EmptyState
          v-else-if="rows.length === 0"
          title="No payments found"
          description="Payments registered for this invoice will appear here."
          :icon="CreditCard"
        />

        <template v-else>
          <DataTable
            :columns="columns"
            :rows="rows"
            :sortable-columns="['created_at', 'amount']"
            :sort-column="sortColumn"
            :sort-order="order"
            @sort="handleSort"
          >
            <template #cell-id="{ value }">
              <span class="font-medium"> #PAY-{{ value }} </span>
            </template>

            <template #cell-amount="{ value }">
              {{ formatCurrency(value) }}
            </template>

            <template #cell-payment_method="{ value }">
              <span
                class="rounded-full border border-border bg-muted px-2 py-1 text-xs font-medium text-muted-foreground"
              >
                {{ value }}
              </span>
            </template>

            <template #cell-paid_by_user_id="{ value }"> User #{{ value }} </template>

            <template #cell-created_at="{ value }">
              {{ formatDate(value) }}
            </template>
          </DataTable>

          <DataPagination
            class="mt-4"
            :page="page"
            :total-pages="totalPages"
            :total="data?.meta.total"
            @previous="page--"
            @next="page++"
          />
        </template>
      </div>
    </SheetContent>
  </Sheet>
</template>
