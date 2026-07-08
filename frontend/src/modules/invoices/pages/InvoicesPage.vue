<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import { FileText, RefreshCw, MoreHorizontal } from 'lucide-vue-next';
import { toast } from 'vue-sonner';

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
  DropdownMenuSeparator,
} from '@/components/ui/dropdown-menu';
import { Button } from '@/components/ui/button';

import {
  type InvoiceStatus,
  type SortOrder,
  type SortColumn,
  SortColumnSchema,
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

import InvoiceDetailsDrawer from '../components/InvoiceDetailsDrawer.vue';
import InvoiceItemsDrawer from '../components/InvoiceItemsDrawer.vue';
import InvoicePaymentsDrawer from '../components/InvoicePaymentsDrawer.vue';

import { useSubmitInvoice } from '../composables/useSubmitInvoice';
import { useInvoices } from '../composables/useInvoices';
import { useCancelInvoice } from '../composables/useCancelInvoice';

import { useUiStore } from '@/shared/stores/ui.store';

const ui = useUiStore();

const page = ref(1);
const limit = ref(10);

const statusInvoice = defineModel<InvoiceStatus>('statusInvoice');
const sortColumn = defineModel<SortColumn>('sortColumn');
const order = defineModel<SortOrder>('order');

function handleSort(columnKey: string) {
  const key = columnKey as SortColumn;
  if (sortColumn.value === key) {
    order.value = order.value === 'asc' ? 'desc' : 'asc';
    return;
  }

  sortColumn.value = key;
  order.value = 'desc';
}

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
  { key: 'actions', label: '' },
];

const rows = computed(() => {
  return invoiceItems.value.map((invoice) => ({
    id: invoice.id,
    buyer_name: invoice.buyer_name,
    status_invoice: invoice.status_invoice,
    total_amount: invoice.total_amount,
    paid_amount: invoice.paid_amount ?? 0,
    created_at: invoice.created_at,
    actions: invoice.id,
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

// Permissions and actions based on invoice status
function canManageItems(status: string) {
  return status === 'draft';
}

function canSubmit(status: string) {
  return status === 'draft';
}

function canCancel(status: string) {
  return status === 'pending';
}

function canViewPayments(status: string) {
  return status === 'paid' || status === 'pending';
}

// Invoice details drawer state
const detailsOpen = ref(false);
const selectedInvoiceId = ref<number | null>(null);

function openInvoiceDetails(row: Record<string, unknown>) {
  selectedInvoiceId.value = Number(row.id);
  detailsOpen.value = true;
}

// Submit invoice confirmation
const submitMutation = useSubmitInvoice();

function confirmSubmitInvoice(row: Record<string, unknown>) {
  const id = Number(row.id);

  ui.openConfirm({
    title: 'Submit invoice?',
    description: 'This will move the invoice from draft to pending and reserve product stock.',
    confirmText: 'Submit invoice',
    cancelText: 'Keep draft',
    variant: 'default',
    onConfirm: async () => {
      await submitMutation.mutateAsync(id);

      toast.success('Invoice submitted');
    },
  });
}

// Cancel invoice confirmation
const cancelMutation = useCancelInvoice();

function confirmCancelInvoice(row: Record<string, unknown>) {
  const id = Number(row.id);

  ui.openConfirm({
    title: 'Cancel invoice?',
    description: 'This will cancel the invoice and release reserved product stock.',
    confirmText: 'Cancel invoice',
    cancelText: 'Keep invoice',
    variant: 'destructive',
    onConfirm: async () => {
      await cancelMutation.mutateAsync(id);

      toast.success('Invoice cancelled');
    },
  });
}

// Invoice items drawer state
const itemsOpen = ref(false);
const selectedItemsInvoiceId = ref<number | null>(null);

function openManageItems(row: Record<string, unknown>) {
  selectedItemsInvoiceId.value = Number(row.id);
  itemsOpen.value = true;
}

// Invoice payments drawer state
const paymentsOpen = ref(false);
const selectedPaymentsInvoiceId = ref<number | null>(null);

function openInvoicePayments(row: Record<string, unknown>) {
  selectedPaymentsInvoiceId.value = Number(row.id);
  paymentsOpen.value = true;
}
</script>

<template>
  <PageContainer>
    <PageHeader
      title="Sales invoices"
      description="Manage invoices where your company acts as the seller."
    />

    <SectionCard
      title="Sales invoice records"
      description="Track sales invoices from draft to payment completion."
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
        <DataTable
          :columns="columns"
          :rows="rows"
          :sortable-columns="SortColumnSchema.options"
          :sort-column="sortColumn"
          :sort-order="order"
          @sort="handleSort"
        >
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

          <template #cell-actions="{ row }">
            <div class="flex justify-end">
              <DropdownMenu>
                <DropdownMenuTrigger as-child>
                  <Button variant="ghost" size="icon">
                    <MoreHorizontal class="h-4 w-4" />
                  </Button>
                </DropdownMenuTrigger>

                <DropdownMenuContent align="end">
                  <DropdownMenuItem @click="openInvoiceDetails(row)">
                    View invoice
                  </DropdownMenuItem>

                  <DropdownMenuItem
                    v-if="canManageItems(String(row.status_invoice))"
                    @click="openManageItems(row)"
                  >
                    Manage items
                  </DropdownMenuItem>

                  <DropdownMenuItem
                    v-if="canSubmit(String(row.status_invoice))"
                    @click="confirmSubmitInvoice(row)"
                  >
                    Submit invoice
                  </DropdownMenuItem>

                  <DropdownMenuItem
                    v-if="canViewPayments(String(row.status_invoice))"
                    @click="openInvoicePayments(row)"
                  >
                    View payments
                  </DropdownMenuItem>

                  <DropdownMenuSeparator v-if="canCancel(String(row.status_invoice))" />

                  <DropdownMenuItem
                    v-if="canCancel(String(row.status_invoice))"
                    class="text-destructive"
                    @click="confirmCancelInvoice(row)"
                  >
                    Cancel invoice
                  </DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>
            </div>
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

    <!-- Invoice details drawer -->
    <InvoiceDetailsDrawer v-model:open="detailsOpen" :invoice-id="selectedInvoiceId" />

    <!-- Invoice items drawer -->
    <InvoiceItemsDrawer v-model:open="itemsOpen" :invoice-id="selectedItemsInvoiceId" />

    <!-- Invoice payments drawer -->
    <InvoicePaymentsDrawer v-model:open="paymentsOpen" :invoice-id="selectedPaymentsInvoiceId" />
  </PageContainer>
</template>
