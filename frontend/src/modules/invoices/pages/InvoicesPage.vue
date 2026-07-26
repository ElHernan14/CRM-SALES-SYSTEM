<script setup lang="ts">
import { computed, ref, watch, unref } from 'vue';
import {
  FileText,
  RefreshCw,
  MoreHorizontal,
  Plus,
  Eye,
  PackageSearch,
  Send,
  CreditCard,
  Ban,
  Search,
} from 'lucide-vue-next';
import { toast } from 'vue-sonner';
import { useQueryClient } from '@tanstack/vue-query';

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

import {
  type InvoiceStatus,
  type SortColumn,
  SortColumnSchema,
} from '@/modules/invoices/types/invoice.types.ts';

import PageContainer from '@/shared/components/erp/PageContainer.vue';
import PageHeader from '@/shared/components/erp/PageHeader.vue';
import SectionCard from '@/shared/components/erp/SectionCard.vue';
import EmptyState from '@/shared/components/erp/EmptyState.vue';
import DataTable from '@/shared/components/erp/DataTable.vue';
import DataPagination from '@/shared/components/erp/DataPagination.vue';

import { Select, SelectContent, SelectItem, SelectTrigger } from '@/components/ui/select';
import InvoiceDetailsDrawer from '../components/InvoiceDetailsDrawer.vue';
import InvoiceItemsDrawer from '../components/InvoiceItemsDrawer.vue';
import InvoicePaymentsDrawer from '../components/InvoicePaymentsDrawer.vue';
import CreateInvoiceDrawer from '../components/CreateInvoiceDrawer.vue';

import { useSubmitInvoice } from '../composables/useSubmitInvoice';
import { useInvoices } from '../composables/useInvoices';
import { useCancelInvoice } from '../composables/useCancelInvoice';

import { useUiStore } from '@/shared/stores/ui.store';

const ui = useUiStore();

const queryClient = useQueryClient();

async function refreshInvoicesWorkspace() {
  await Promise.all([
    queryClient.invalidateQueries({ queryKey: ['invoices'] }),
    queryClient.invalidateQueries({ queryKey: ['invoice'] }),
    queryClient.invalidateQueries({ queryKey: ['invoice-items'] }),
    queryClient.invalidateQueries({ queryKey: ['invoice-payments'] }),
  ]);
}

const page = ref(1);
const limit = ref(10);

const statusInvoice = ref('all');
const buyerSearch = ref('');
const sortColumn = ref<SortColumn>('created_at');
const order = ref<'asc' | 'desc'>('desc');

function handleSort(columnKey: string) {
  const key = columnKey as SortColumn;
  if (sortColumn.value === key) {
    order.value = order.value === 'asc' ? 'desc' : 'asc';
    return;
  }

  sortColumn.value = key;
  order.value = 'desc';
}

watch([buyerSearch, statusInvoice, sortColumn, order], () => {
  page.value = 1;
});

const invoiceParams = computed(() => ({
  page: page.value,
  limit: limit.value,

  status_invoice:
    statusInvoice.value !== 'all' ? (statusInvoice.value as InvoiceStatus) : undefined,

  buyer_name: buyerSearch.value.trim().length >= 2 ? buyerSearch.value.trim() : undefined,

  sort_column: sortColumn.value,
  order: order.value,
}));

const { data, isLoading, isError } = useInvoices(invoiceParams);

const invoiceItems = computed(() => {
  return data.value?.items ?? [];
});

const totalPages = computed(() => {
  return data.value?.meta.total_pages ?? 1;
});

const columns = [
  {
    key: 'id',
    label: 'Invoice',
    cellClass: 'whitespace-nowrap',
  },
  {
    key: 'buyer_name',
    label: 'Buyer',
    cellClass: 'min-w-[220px]',
  },
  {
    key: 'status_invoice',
    label: 'Status',
    cellClass: 'whitespace-nowrap',
  },
  {
    key: 'total_amount',
    label: 'Total',
    cellClass: 'whitespace-nowrap',
  },
  {
    key: 'paid_amount',
    label: 'Paid',
    cellClass: 'whitespace-nowrap',
  },
  {
    key: 'created_at',
    label: 'Created',
    cellClass: 'whitespace-nowrap',
  },
  {
    key: 'actions',
    label: '',
    headerClass: 'sticky right-0 z-10 w-14 bg-muted/95',
    cellClass: 'sticky right-0 z-10 w-14 bg-card',
  },
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
  statusInvoice.value = 'all';
  sortColumn.value = 'created_at';
  buyerSearch.value = '';
  order.value = 'desc';
  page.value = 1;
}

function formatCurrency(value: unknown) {
  return `$${Number(value).toFixed(2)}`;
}

function formatDate(value: unknown) {
  return new Date(String(value)).toLocaleDateString();
}

function getStatusLabel(status: unknown) {
  const s = String(unref(status) ?? 'all');
  switch (s) {
    case 'draft':
      return 'Draft';
    case 'pending':
      return 'Pending';
    case 'paid':
      return 'Paid';
    case 'cancelled':
      return 'Cancelled';
    default:
      return 'All statuses';
  }
}

function canCancel(status: string) {
  return status === 'pending';
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

// Create invoice drawer state
const createOpen = ref(false);

function handleInvoiceCreated(invoiceId: number) {
  selectedItemsInvoiceId.value = invoiceId;
  itemsOpen.value = true;
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
        <div class="grid w-full grid-cols-2 gap-2 sm:flex sm:w-auto sm:items-center">
          <Button
            variant="outline"
            size="sm"
            class="w-full sm:w-auto"
            @click="refreshInvoicesWorkspace"
          >
            <RefreshCw class="mr-2 h-4 w-4" />
            Refresh
          </Button>

          <Button size="sm" class="w-full sm:w-auto" @click="createOpen = true">
            <Plus class="mr-2 h-4 w-4" />
            New invoice
          </Button>
        </div>
      </template>

      <div class="mb-5 space-y-4">
        <div class="grid gap-3 md:grid-cols-[minmax(260px,1fr)_220px]">
          <div class="relative">
            <Search
              class="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
            />

            <Input
              v-model="buyerSearch"
              class="pl-9"
              placeholder="Search buyer company or contact..."
            />
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
          table-min-width="940px"
          :columns="columns"
          :rows="rows"
          :sortable-columns="SortColumnSchema.options"
          :sort-column="sortColumn"
          :sort-order="order"
          @sort="handleSort"
        >
          <!-- MOBILE -->
          <template #mobile-card="{ row }">
            <div class="p-4">
              <div class="flex items-start justify-between gap-3">
                <div class="min-w-0">
                  <div class="flex flex-wrap items-center gap-2">
                    <span class="text-sm font-semibold text-foreground"> #INV-{{ row.id }} </span>

                    <span
                      class="rounded-full px-2.5 py-1 text-xs font-medium capitalize"
                      :class="{
                        'bg-muted text-muted-foreground': row.status_invoice === 'draft',
                        'bg-amber-500/10 text-amber-600 dark:text-amber-400':
                          row.status_invoice === 'pending',
                        'bg-emerald-500/10 text-emerald-600 dark:text-emerald-400':
                          row.status_invoice === 'paid',
                        'bg-destructive/10 text-destructive': row.status_invoice === 'cancelled',
                      }"
                    >
                      {{ getStatusLabel(row.status_invoice) }}
                    </span>
                  </div>

                  <p class="mt-3 break-words text-base font-semibold text-foreground">
                    {{ row.buyer_name }}
                  </p>

                  <p class="mt-1 text-xs text-muted-foreground">
                    Created {{ formatDate(row.created_at) }}
                  </p>
                </div>

                <DropdownMenu>
                  <DropdownMenuTrigger as-child>
                    <Button
                      type="button"
                      variant="ghost"
                      size="icon"
                      class="h-9 w-9 shrink-0 rounded-full"
                    >
                      <MoreHorizontal class="h-4 w-4" />
                    </Button>
                  </DropdownMenuTrigger>

                  <DropdownMenuContent align="end">
                    <DropdownMenuItem @click="openInvoiceDetails(row)">
                      <Eye class="mr-2 h-4 w-4" />
                      View invoice
                    </DropdownMenuItem>

                    <DropdownMenuItem
                      v-if="row.status_invoice === 'draft'"
                      @click="openManageItems(row)"
                    >
                      <PackageSearch class="mr-2 h-4 w-4" />
                      Manage items
                    </DropdownMenuItem>

                    <DropdownMenuItem
                      v-if="row.status_invoice === 'draft'"
                      @click="confirmSubmitInvoice(row)"
                    >
                      <Send class="mr-2 h-4 w-4" />
                      Submit invoice
                    </DropdownMenuItem>

                    <DropdownMenuItem @click="openInvoicePayments(row)">
                      <CreditCard class="mr-2 h-4 w-4" />
                      Payment history
                    </DropdownMenuItem>

                    <DropdownMenuItem
                      v-if="canCancel(String(row.status_invoice))"
                      class="text-destructive"
                      @click="confirmCancelInvoice(row)"
                    >
                      <Ban class="mr-2 h-4 w-4" />
                      Cancel invoice
                    </DropdownMenuItem>
                  </DropdownMenuContent>
                </DropdownMenu>
              </div>

              <div class="mt-4 grid grid-cols-2 gap-3 border-t border-border pt-4">
                <div>
                  <p class="text-[10px] font-medium uppercase tracking-wide text-muted-foreground">
                    Invoice total
                  </p>

                  <p class="mt-1 truncate text-lg font-semibold text-foreground">
                    {{ formatCurrency(row.total_amount) }}
                  </p>
                </div>

                <div class="text-right">
                  <p class="text-[10px] font-medium uppercase tracking-wide text-muted-foreground">
                    Paid
                  </p>

                  <p
                    class="mt-1 truncate text-lg font-semibold text-emerald-600 dark:text-emerald-400"
                  >
                    {{ formatCurrency(row.paid_amount) }}
                  </p>
                </div>
              </div>
            </div>
          </template>

          <!-- DESKTOP -->
          <template #cell-id="{ value }">
            <span class="font-medium"> #INV-{{ value }} </span>
          </template>

          <template #cell-status_invoice="{ value }">
            <span
              class="rounded-full px-2 py-1 text-xs font-medium capitalize"
              :class="{
                'bg-muted text-muted-foreground': value === 'draft',
                'bg-amber-500/10 text-amber-600 dark:text-amber-400': value === 'pending',
                'bg-emerald-500/10 text-emerald-600 dark:text-emerald-400': value === 'paid',
                'bg-destructive/10 text-destructive': value === 'cancelled',
              }"
            >
              {{ getStatusLabel(value) }}
            </span>
          </template>

          <template #cell-total_amount="{ value }">
            <span class="font-semibold">
              {{ formatCurrency(value) }}
            </span>
          </template>

          <template #cell-paid_amount="{ value }">
            <span class="text-emerald-600 dark:text-emerald-400">
              {{ formatCurrency(value) }}
            </span>
          </template>

          <template #cell-created_at="{ value }">
            {{ formatDate(value) }}
          </template>

          <template #cell-actions="{ row }">
            <div class="flex justify-end">
              <DropdownMenu>
                <DropdownMenuTrigger as-child>
                  <Button variant="ghost" size="icon" class="rounded-full">
                    <MoreHorizontal class="h-4 w-4" />
                  </Button>
                </DropdownMenuTrigger>

                <DropdownMenuContent align="end">
                  <DropdownMenuItem @click="openInvoiceDetails(row)">
                    <Eye class="mr-2 h-4 w-4" />
                    View invoice
                  </DropdownMenuItem>

                  <DropdownMenuItem
                    v-if="row.status_invoice === 'draft'"
                    @click="openManageItems(row)"
                  >
                    <PackageSearch class="mr-2 h-4 w-4" />
                    Manage items
                  </DropdownMenuItem>

                  <DropdownMenuItem
                    v-if="row.status_invoice === 'draft'"
                    @click="confirmSubmitInvoice(row)"
                  >
                    <Send class="mr-2 h-4 w-4" />
                    Submit invoice
                  </DropdownMenuItem>

                  <DropdownMenuItem @click="openInvoicePayments(row)">
                    <CreditCard class="mr-2 h-4 w-4" />
                    Payment history
                  </DropdownMenuItem>

                  <DropdownMenuItem
                    v-if="canCancel(String(row.status_invoice))"
                    class="text-destructive"
                    @click="confirmCancelInvoice(row)"
                  >
                    <Ban class="mr-2 h-4 w-4" />
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

    <!-- Create invoice drawer -->
    <CreateInvoiceDrawer v-model:open="createOpen" @created="handleInvoiceCreated" />
  </PageContainer>
</template>
