<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import { FileText, RefreshCw } from 'lucide-vue-next';

import { Button } from '@/components/ui/button';

import {
  type InvoiceStatus,
  type SortOrder,
  type SortColumn,
} from '@/modules/invoices/types/invoice.types.ts';

import PageContainer from '@/shared/components/erp/PageContainer.vue';
import PageHeader from '@/shared/components/erp/PageHeader.vue';
import SectionCard from '@/shared/components/erp/SectionCard.vue';
import EmptyState from '@/shared/components/erp/EmptyState.vue';
import DataTable from '@/shared/components/erp/DataTable.vue';
import DataPagination from '@/shared/components/erp/DataPagination.vue';

import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';

import { useInvoices } from '../composables/useInvoices';

const page = ref(1);
const limit = ref(10);

const statusInvoice = defineModel<InvoiceStatus>('statusInvoice');
const sortColumn = defineModel<SortColumn>('sortColumn');
const order = defineModel<SortOrder>('order');

watch([statusInvoice, sortColumn, order], () => {
  page.value = 1;
});

const invoiceParams = computed(() => ({
  page: page.value,
  limit: limit.value,
  status_invoice: statusInvoice.value || undefined,
  sort_column: sortColumn.value || undefined,
  order: order.value || undefined,
}));

const { data, isLoading, isError, refetch } = useInvoices(invoiceParams);

const invoiceItems = computed(() => {
  return data.value?.items ?? [];
});

const totalPages = computed(() => {
  return data.value?.meta.total_pages ?? 1;
});

const columns = [
  { key: 'id', label: 'Invoice' },
  { key: 'buyer_name', label: 'Buyer' },
  { key: 'status_invoice', label: 'Status' },
  { key: 'total_amount', label: 'Total' },
  { key: 'paid_amount', label: 'Paid' },
  { key: 'created_at', label: 'Created' },
];

const rows = computed(() => {
  return invoiceItems.value.map((invoice) => ({
    id: invoice.id,
    buyer_name: invoice.buyer_name,
    status_invoice: invoice.status_invoice,
    total_amount: invoice.total_amount,
    paid_amount: invoice.paid_amount ?? 0,
    created_at: invoice.created_at,
  }));
});

function clearFilters() {
  statusInvoice.value = undefined;
  sortColumn.value = 'created_at';
  order.value = 'desc';
  page.value = 1;
}

function formatCurrency(value: unknown) {
  return `$${Number(value).toFixed(2)}`;
}

function formatDate(value: unknown) {
  return new Date(String(value)).toLocaleDateString();
}
</script>

<template>
  <PageContainer>
    <PageHeader
      title="Invoices"
      description="Manage commercial documents, workflow states and payments."
    />

    <SectionCard
      title="Company invoices"
      description="Track invoices from draft to payment completion."
    >
      <template #actions>
        <Button variant="outline" size="sm" @click="refetch()">
          <RefreshCw class="mr-2 h-4 w-4" />
          Refresh
        </Button>
      </template>

      <div class="mb-5 space-y-4">
        <div class="grid gap-3 md:grid-cols-3">
          <Select v-model="statusInvoice">
            <SelectTrigger>
              <SelectValue placeholder="Invoice status" />
            </SelectTrigger>

            <SelectContent>
              <SelectItem value="draft"> Draft </SelectItem>

              <SelectItem value="pending"> Pending </SelectItem>

              <SelectItem value="paid"> Paid </SelectItem>

              <SelectItem value="cancelled"> Cancelled </SelectItem>
            </SelectContent>
          </Select>

          <Select v-model="sortColumn">
            <SelectTrigger>
              <SelectValue placeholder="Sort by" />
            </SelectTrigger>

            <SelectContent>
              <SelectItem value="created_at"> Created date </SelectItem>

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
        </div>

        <div class="flex items-center justify-between">
          <p class="text-sm text-muted-foreground">Total: {{ data?.meta.total ?? 0 }}</p>

          <Button variant="ghost" size="sm" @click="clearFilters"> Clear filters </Button>
        </div>
      </div>

      <div
        v-if="isLoading"
        class="rounded-xl border border-border bg-muted/30 p-8 text-center text-sm text-muted-foreground"
      >
        Loading invoices...
      </div>

      <EmptyState
        v-else-if="isError"
        title="Unable to load invoices"
        description="There was a problem connecting to the invoices API."
        :icon="FileText"
      />

      <EmptyState
        v-else-if="rows.length === 0"
        title="No invoices found"
        description="Invoices created in this workspace will appear here."
        :icon="FileText"
      />

      <template v-else>
        <DataTable :columns="columns" :rows="rows">
          <template #cell-id="{ value }">
            <span class="font-medium"> #INV-{{ value }} </span>
          </template>

          <template #cell-status_invoice="{ value }">
            <span
              class="rounded-full px-2 py-1 text-xs font-medium"
              :class="{
                'bg-muted text-muted-foreground': value === 'draft',
                'bg-amber-500/10 text-amber-600 dark:text-amber-400': value === 'pending',
                'bg-emerald-500/10 text-emerald-600 dark:text-emerald-400': value === 'paid',
                'bg-destructive/10 text-destructive': value === 'cancelled',
              }"
            >
              {{ value }}
            </span>
          </template>

          <template #cell-total_amount="{ value }">
            {{ formatCurrency(value) }}
          </template>

          <template #cell-paid_amount="{ value }">
            {{ formatCurrency(value) }}
          </template>

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
    </SectionCard>
  </PageContainer>
</template>
