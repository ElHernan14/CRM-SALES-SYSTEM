<script setup lang="ts">
import { computed, reactive, ref, watch } from 'vue';

import {
  Building2,
  CalendarDays,
  CircleDollarSign,
  AlertCircle,
  CheckCircle2,
  ShieldCheck,
  Loader2,
  PackageSearch,
  ReceiptText,
  WalletCards,
  Boxes,
  Package,
  Banknote,
  ChevronDown,
  ChevronUp,
  CircleCheck,
  CreditCard,
  Landmark,
  Receipt,
  RefreshCw,
  ArrowRight,
  Send,
} from 'lucide-vue-next';

import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Select, SelectContent, SelectItem, SelectTrigger } from '@/components/ui/select';

import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
} from '@/components/ui/sheet';

import EmptyState from '@/shared/components/erp/EmptyState.vue';

import { usePayPurchase } from '../composables/usePayPurchase';
import { usePurchasePayments } from '../composables/usePurchasePayments';
import { usePurchaseItems } from '../composables/usePurchaseItems';
import { useInvoiceDetail } from '@/modules/invoices/composables/useInvoiceDetail';
import { useCheckoutPurchase } from '../composables/useCheckoutPurchase';
import { useCheckoutMarketplacePurchaseCart } from '@/modules/marketplace/composables/useCheckoutMarketplacePurchaseCart';

import type { PaymentMethod } from '../types/purchase-payment.types';
import type { PurchaseListItem } from '../types/purchase.types';
import type { StorePurchase } from '@/modules/store/types/store-purchase.types';

import { toast } from 'vue-sonner';

const props = withDefaults(
  defineProps<{
    open: boolean;
    purchase: PurchaseListItem | StorePurchase | null;
    initialAction?: 'overview' | 'checkout' | 'payment';
    variant?: 'erp' | 'store';
  }>(),
  {
    initialAction: 'overview',
    variant: 'erp',
  }
);

const emit = defineEmits<{
  'update:open': [value: boolean];
}>();

const paymentBalanceLimitCents = ref(0);
const paymentExpanded = ref(false);

const paymentForm = reactive({
  amount: '',
  method: 'transfer' as PaymentMethod,
});

const payPurchaseMutation = usePayPurchase();

const isPaying = computed(() => {
  return payPurchaseMutation.isPending.value;
});

const paymentAmount = computed(() => {
  return parseMoneyInput(paymentForm.amount);
});

const paymentAmountCents = computed(() => {
  return toMoneyCents(paymentAmount.value);
});

const paymentAmountValid = computed(() => {
  return paymentAmountCents.value > 0 && paymentAmountCents.value <= paymentBalanceLimitCents.value;
});

const balanceAfterPaymentCents = computed(() => {
  return Math.max(paymentBalanceLimitCents.value - paymentAmountCents.value, 0);
});

const balanceAfterPayment = computed(() => {
  return fromMoneyCents(balanceAfterPaymentCents.value);
});

const willCompletePurchase = computed(() => {
  return paymentAmountValid.value && balanceAfterPaymentCents.value === 0;
});

const purchaseId = computed<number | null>(() => {
  return props.purchase?.id ?? null;
});

const paymentsExpanded = ref(false);

const shouldLoadPayments = computed(() => {
  return props.open && paymentsExpanded.value && !!purchaseId.value;
});

const {
  data: paymentsData,
  isLoading: paymentsLoading,
  isFetching: paymentsFetching,
  isError: paymentsError,
  refetch: refetchPayments,
} = usePurchasePayments(purchaseId, shouldLoadPayments);

const itemsExpanded = ref(true);

const shouldLoadItems = computed(() => {
  return props.open && itemsExpanded.value && !!purchaseId.value;
});

const purchasePayments = computed(() => {
  return paymentsData.value?.items ?? [];
});

const paymentsTotal = computed(() => {
  return purchasePayments.value.reduce((total, payment) => total + payment.amount, 0);
});

const paymentsRefreshing = computed(() => {
  return paymentsFetching.value && !paymentsLoading.value;
});

watch(
  () => props.purchase?.id,
  () => {
    itemsExpanded.value = true;
    paymentsExpanded.value = false;
  }
);

const {
  data: itemsData,
  isLoading: itemsLoading,
  isFetching: itemsFetching,
  isError: itemsError,
  refetch: refetchItems,
} = usePurchaseItems(purchaseId, shouldLoadItems);

const purchaseItems = computed(() => {
  return itemsData.value?.items ?? [];
});

const totalUnits = computed(() => {
  return purchaseItems.value.reduce((total, item) => total + item.quantity, 0);
});

const itemsRefreshing = computed(() => {
  return itemsFetching.value && !itemsLoading.value;
});

watch(
  () => props.purchase?.id,
  () => {
    itemsExpanded.value = true;
  }
);

const { data: invoice, isLoading, isFetching, isError, refetch } = useInvoiceDetail(purchaseId);

const isRefreshing = computed(() => {
  return isFetching.value && !isLoading.value;
});

const remainingAmountCents = computed(() => {
  if (!invoice.value) {
    return 0;
  }

  const totalCents = toMoneyCents(invoice.value.total_amount);
  const paidCents = toMoneyCents(invoice.value.paid_amount);

  return Math.max(totalCents - paidCents, 0);
});

const remainingAmount = computed(() => {
  return fromMoneyCents(remainingAmountCents.value);
});

const paymentProgress = computed(() => {
  if (!invoice.value) return 0;

  if (invoice.value.total_amount <= 0) {
    return 0;
  }

  return Math.min(Math.round((invoice.value.paid_amount / invoice.value.total_amount) * 100), 100);
});

const canPay = computed(() => {
  if (!invoice.value) {
    return false;
  }

  return invoice.value.status_invoice === 'pending' && remainingAmountCents.value > 0;
});

function toMoneyCents(value: number) {
  if (!Number.isFinite(value)) {
    return 0;
  }

  return Math.round((value + Number.EPSILON) * 100);
}

function fromMoneyCents(value: number) {
  return value / 100;
}

function parseMoneyInput(value: string) {
  const normalized = value.trim().replace(/\s/g, '').replace(',', '.');

  if (!normalized) {
    return 0;
  }

  const parsed = Number(normalized);

  return Number.isFinite(parsed) ? parsed : 0;
}

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

function handlePay() {
  if (!invoice.value || !canPay.value) {
    return;
  }

  paymentBalanceLimitCents.value = remainingAmountCents.value;

  paymentForm.amount = fromMoneyCents(paymentBalanceLimitCents.value).toFixed(2);

  paymentForm.method = 'transfer';
  paymentExpanded.value = true;
}

watch(
  () => props.purchase?.id,
  () => {
    itemsExpanded.value = true;
    paymentsExpanded.value = false;
    paymentExpanded.value = false;

    paymentForm.amount = '';
    paymentBalanceLimitCents.value = 0;
    paymentForm.method = 'transfer';
  }
);

function getSelectedPaymentMethodLabel() {
  const labels: Record<PaymentMethod, string> = {
    cash: 'Cash',
    card: 'Card',
    transfer: 'Bank transfer',
  };

  return labels[paymentForm.method];
}

function getSelectedPaymentMethodIcon() {
  const icons = {
    cash: Banknote,
    card: CreditCard,
    transfer: Landmark,
  };

  return icons[paymentForm.method];
}

async function refreshPurchaseWorkspace() {
  const requests: Promise<unknown>[] = [refetch()];

  if (itemsExpanded.value) {
    requests.push(refetchItems());
  }

  if (paymentsExpanded.value) {
    requests.push(refetchPayments());
  }

  await Promise.all(requests);
}

const isRefreshingWorkspace = computed(() => {
  return isRefreshing.value || itemsRefreshing.value || paymentsRefreshing.value;
});

function getPaymentMethodLabel(method: string) {
  const labels: Record<string, string> = {
    cash: 'Cash',
    card: 'Card',
    transfer: 'Bank transfer',
  };

  return labels[method] ?? method;
}

function getPaymentMethodIcon(method: string) {
  const icons: Record<string, typeof CreditCard> = {
    cash: Banknote,
    card: CreditCard,
    transfer: Landmark,
  };

  return icons[method] ?? Receipt;
}

function setPaymentPercentage(percentage: number) {
  const amountCents = Math.round(remainingAmountCents.value * percentage);

  paymentForm.amount = fromMoneyCents(amountCents).toFixed(2);
}

function setFullPaymentAmount() {
  paymentForm.amount = fromMoneyCents(remainingAmountCents.value).toFixed(2);
}

async function submitPayment() {
  if (!invoice.value) {
    toast.error('Missing purchase information');
    return;
  }

  if (paymentAmountCents.value <= 0) {
    toast.error('Enter a valid payment amount');
    return;
  }

  if (paymentAmountCents.value > remainingAmountCents.value) {
    toast.error('Payment amount exceeds the outstanding balance');
    return;
  }

  try {
    const result = await payPurchaseMutation.mutateAsync({
      invoiceId: invoice.value.id,

      payload: {
        amount: fromMoneyCents(paymentAmountCents.value),
        payment_method: paymentForm.method,
      },
    });

    await Promise.all([refetch(), refetchPayments()]);

    paymentsExpanded.value = true;
    paymentExpanded.value = false;

    paymentForm.amount = '';

    if (toMoneyCents(result.remaining_amount) === 0) {
      toast.success('Purchase fully paid', {
        description: 'The invoice is now marked as paid.',
      });
    } else {
      toast.success('Payment registered', {
        description: `${formatCurrency(result.remaining_amount)} remains outstanding.`,
      });
    }
  } catch (error: any) {
    const message =
      error?.response?.data?.errorMessage ??
      error?.response?.data?.message ??
      'Failed to register payment';

    toast.error(message);
  }
}

const checkoutStorePurchaseMutation = useCheckoutPurchase();

const checkoutMarketplacePurchaseMutation = useCheckoutMarketplacePurchaseCart();

const isCheckingOut = computed(() => {
  return props.variant === 'store'
    ? checkoutStorePurchaseMutation.isPending.value
    : checkoutMarketplacePurchaseMutation.isPending.value;
});

const canCheckout = computed(() => {
  if (!invoice.value) return false;

  return invoice.value.status_invoice === 'draft' && purchaseItems.value.length > 0;
});

async function submitCheckout() {
  if (!invoice.value) {
    toast.error('Missing purchase information');
    return;
  }

  if (invoice.value.status_invoice !== 'draft') {
    toast.error('Only draft purchases can be submitted');
    return;
  }

  if (purchaseItems.value.length === 0) {
    toast.error('Add at least one item before checkout');
    return;
  }

  try {
    if (props.variant === 'store') {
      /*
       * Flujo personal:
       * POST /api/store/checkout
       * source = store
       */
      await checkoutStorePurchaseMutation.mutateAsync(invoice.value.id);
    } else {
      /*
       * Flujo organizacional:
       * POST /api/marketplace/checkout
       * source = erp
       */
      const result = await checkoutMarketplacePurchaseMutation.mutateAsync(invoice.value.id);

      if (result.source !== 'erp') {
        throw new Error('Marketplace checkout returned an invalid source');
      }
    }

    await Promise.all([refetch(), refetchItems()]);

    toast.success(props.variant === 'store' ? 'Order submitted' : 'Purchase submitted', {
      description:
        props.variant === 'store'
          ? 'The personal Store order is now awaiting payment.'
          : 'The B2B supplier invoice is now awaiting payment.',
    });
  } catch (error: any) {
    const message =
      error?.response?.data?.errorMessage ??
      error?.response?.data?.message ??
      (props.variant === 'store'
        ? 'Failed to submit Store order'
        : 'Failed to submit B2B purchase');

    toast.error(message);
  }
}

watch([() => props.open, () => props.purchase?.id], ([open]) => {
  if (!open) return;

  if (props.initialAction === 'payment') {
    requestAnimationFrame(() => {
      handlePay();
    });
  }
});

// Validate number input for minPrice and maxPrice
function validateNumberInput(event: KeyboardEvent) {
  const allowedKeys = ['Backspace', 'Delete', 'ArrowLeft', 'ArrowRight', 'Tab', 'Enter'];

  // Permitir números y punto decimal
  const isNumber = /^[0-9]$/.test(event.key);
  const isDot = event.key === '.';

  if (!isNumber && !isDot && !allowedKeys.includes(event.key)) {
    event.preventDefault();
  }
}
</script>

<template>
  <Sheet :open="open" @update:open="emit('update:open', $event)">
    <SheetContent
      class="flex w-full flex-col overflow-y-auto p-0"
      :class="variant === 'store' ? 'sm:max-w-3xl' : 'sm:max-w-2xl'"
    >
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
                <SheetTitle>
                  {{ variant === 'store' ? 'Order details' : 'Purchase details' }}
                </SheetTitle>

                <SheetDescription>
                  {{
                    variant === 'store'
                      ? 'Review products, payments and the current order balance.'
                      : 'Review buyer-side purchase information.'
                  }}
                </SheetDescription>
              </div>
            </div>

            <Button
              v-if="invoice"
              type="button"
              variant="ghost"
              size="icon"
              :disabled="isRefreshingWorkspace"
              @click="refreshPurchaseWorkspace"
            >
              <RefreshCw
                class="h-4 w-4"
                :class="{
                  'animate-spin': isRefreshingWorkspace,
                }"
              />
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
      <div v-else-if="invoice && purchase" class="space-y-5 p-4 sm:space-y-6 sm:p-6">
        <!-- SUPPLIER -->
        <section
          class="relative overflow-hidden rounded-2xl border border-border bg-card p-5 shadow-sm"
        >
          <div
            class="pointer-events-none absolute inset-0 bg-linear-to-br from-primary/5 via-transparent to-muted/40"
          />

          <div
            class="relative flex flex-col gap-4 min-[390px]:flex-row min-[390px]:items-start min-[390px]:justify-between"
          >
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

        <!-- PURCHASE ITEMS -->
        <section class="overflow-hidden rounded-2xl border border-border bg-card shadow-sm">
          <button
            type="button"
            class="flex w-full items-center justify-between gap-4 p-5 text-left transition hover:bg-muted/20"
            @click="itemsExpanded = !itemsExpanded"
          >
            <div class="flex min-w-0 items-center gap-3">
              <div
                class="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl bg-primary/10"
              >
                <Boxes class="h-5 w-5 text-primary" />
              </div>

              <div class="min-w-0">
                <h3 class="text-sm font-semibold text-foreground">Purchased items</h3>

                <p class="mt-1 text-xs text-muted-foreground">
                  <template v-if="itemsLoading"> Loading order contents... </template>

                  <template v-else>
                    {{ purchaseItems.length }}
                    {{ purchaseItems.length === 1 ? 'line item' : 'line items' }}

                    <template v-if="purchaseItems.length">
                      · {{ totalUnits }}
                      {{ totalUnits === 1 ? 'unit' : 'units' }}
                    </template>
                  </template>
                </p>
              </div>
            </div>

            <div class="flex shrink-0 items-center gap-3">
              <span
                v-if="purchaseItems.length"
                class="hidden text-sm font-semibold text-foreground sm:block"
              >
                {{ formatCurrency(invoice.subtotal) }}
              </span>

              <ChevronUp v-if="itemsExpanded" class="h-4 w-4 text-muted-foreground" />

              <ChevronDown v-else class="h-4 w-4 text-muted-foreground" />
            </div>
          </button>

          <Transition
            enter-active-class="transition-all duration-200 ease-out"
            enter-from-class="max-h-0 opacity-0"
            enter-to-class="max-h-[900px] opacity-100"
            leave-active-class="transition-all duration-150 ease-in"
            leave-from-class="max-h-[900px] opacity-100"
            leave-to-class="max-h-0 opacity-0"
          >
            <div v-if="itemsExpanded" class="overflow-hidden border-t border-border">
              <!-- LOADING -->
              <div v-if="itemsLoading" class="space-y-3 p-5">
                <div
                  v-for="index in 3"
                  :key="index"
                  class="animate-pulse rounded-xl border border-border bg-muted/20 p-4"
                >
                  <div class="h-4 w-2/3 rounded bg-muted" />

                  <div class="mt-4 grid grid-cols-3 gap-3">
                    <div class="h-8 rounded bg-muted" />
                    <div class="h-8 rounded bg-muted" />
                    <div class="h-8 rounded bg-muted" />
                  </div>
                </div>
              </div>

              <!-- ERROR -->
              <div v-else-if="itemsError" class="p-5">
                <div
                  class="rounded-xl border border-destructive/20 bg-destructive/5 p-4 text-center"
                >
                  <p class="text-sm font-medium text-foreground">Unable to load purchased items</p>

                  <p class="mt-1 text-xs text-muted-foreground">
                    The invoice is available, but its product lines could not be retrieved.
                  </p>

                  <Button
                    type="button"
                    variant="outline"
                    size="sm"
                    class="mt-4"
                    @click="refetchItems()"
                  >
                    <RefreshCw class="mr-2 h-4 w-4" />
                    Try again
                  </Button>
                </div>
              </div>

              <!-- EMPTY -->
              <div v-else-if="purchaseItems.length === 0" class="p-6 text-center">
                <div
                  class="mx-auto flex h-12 w-12 items-center justify-center rounded-xl border border-border bg-muted/30"
                >
                  <Package class="h-5 w-5 text-muted-foreground" />
                </div>

                <p class="mt-4 text-sm font-medium text-foreground">No items registered</p>

                <p class="mt-1 text-xs text-muted-foreground">
                  This purchase does not contain visible product lines.
                </p>
              </div>

              <!-- ITEMS -->
              <div v-else class="relative max-h-[430px] overflow-y-auto p-5">
                <div
                  v-if="itemsRefreshing"
                  class="absolute inset-x-5 top-5 z-10 flex items-center gap-3 rounded-xl border border-border bg-background/90 p-3 shadow-sm backdrop-blur"
                >
                  <Loader2 class="h-4 w-4 animate-spin text-muted-foreground" />

                  <p class="text-xs text-muted-foreground">Updating purchase items...</p>
                </div>

                <div class="space-y-3">
                  <article
                    v-for="item in purchaseItems"
                    :key="item.id"
                    class="group rounded-xl border border-border bg-background p-4 transition hover:border-primary/20 hover:shadow-sm"
                  >
                    <div class="flex items-start justify-between gap-4">
                      <div class="flex min-w-0 items-start gap-3">
                        <div
                          class="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl border border-border bg-muted/30"
                        >
                          <PackageSearch class="h-4 w-4 text-muted-foreground" />
                        </div>

                        <div class="min-w-0">
                          <p class="line-clamp-2 text-sm font-semibold text-foreground">
                            {{ item.product_name }}
                          </p>

                          <p class="mt-1 text-xs text-muted-foreground">
                            {{ formatCurrency(item.price) }}
                            per unit
                          </p>
                        </div>
                      </div>

                      <p class="shrink-0 text-sm font-semibold text-foreground">
                        {{ formatCurrency(item.subtotal) }}
                      </p>
                    </div>

                    <div
                      class="mt-4 grid grid-cols-2 gap-3 border-t border-border pt-4 sm:grid-cols-3"
                    >
                      <div>
                        <p class="text-[11px] uppercase tracking-wide text-muted-foreground">
                          Quantity
                        </p>

                        <p class="mt-1 text-sm font-semibold text-foreground">
                          {{ item.quantity }}
                        </p>
                      </div>

                      <div>
                        <p class="text-[11px] uppercase tracking-wide text-muted-foreground">
                          Unit price
                        </p>

                        <p class="mt-1 text-sm font-medium text-foreground">
                          {{ formatCurrency(item.price) }}
                        </p>
                      </div>

                      <div class="col-span-2 text-left sm:col-span-1 sm:text-right">
                        <p class="text-[11px] uppercase tracking-wide text-muted-foreground">
                          Subtotal
                        </p>

                        <p class="mt-1 text-sm font-semibold text-foreground">
                          {{ formatCurrency(item.subtotal) }}
                        </p>
                      </div>
                    </div>
                  </article>
                </div>

                <div
                  class="mt-4 flex items-center justify-between rounded-xl border border-border bg-muted/20 p-4"
                >
                  <div>
                    <p class="text-xs text-muted-foreground">Items subtotal</p>

                    <p class="mt-1 text-xs text-muted-foreground">
                      {{ totalUnits }}
                      {{ totalUnits === 1 ? 'unit purchased' : 'units purchased' }}
                    </p>
                  </div>

                  <p class="text-lg font-semibold text-foreground">
                    {{ formatCurrency(invoice.subtotal) }}
                  </p>
                </div>
              </div>
            </div>
          </Transition>
        </section>

        <!-- PAYMENT HISTORY -->
        <section class="overflow-hidden rounded-2xl border border-border bg-card shadow-sm">
          <button
            type="button"
            class="flex w-full items-center justify-between gap-4 p-5 text-left transition hover:bg-muted/20"
            @click="paymentsExpanded = !paymentsExpanded"
          >
            <div class="flex min-w-0 items-center gap-3">
              <div
                class="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl bg-emerald-500/10"
              >
                <CreditCard class="h-5 w-5 text-emerald-600 dark:text-emerald-400" />
              </div>

              <div class="min-w-0">
                <h3 class="text-sm font-semibold text-foreground">Payment history</h3>

                <p class="mt-1 text-xs text-muted-foreground">
                  <template v-if="paymentsLoading"> Loading registered payments... </template>

                  <template v-else-if="paymentsExpanded">
                    {{ purchasePayments.length }}
                    {{
                      purchasePayments.length === 1 ? 'payment registered' : 'payments registered'
                    }}
                  </template>

                  <template v-else> Review transactions associated with this purchase. </template>
                </p>
              </div>
            </div>

            <div class="flex shrink-0 items-center gap-3">
              <span
                v-if="paymentsExpanded && purchasePayments.length"
                class="hidden text-sm font-semibold text-foreground sm:block"
              >
                {{ formatCurrency(paymentsTotal) }}
              </span>

              <ChevronUp v-if="paymentsExpanded" class="h-4 w-4 text-muted-foreground" />

              <ChevronDown v-else class="h-4 w-4 text-muted-foreground" />
            </div>
          </button>

          <Transition
            enter-active-class="transition-all duration-200 ease-out"
            enter-from-class="max-h-0 opacity-0"
            enter-to-class="max-h-[800px] opacity-100"
            leave-active-class="transition-all duration-150 ease-in"
            leave-from-class="max-h-[800px] opacity-100"
            leave-to-class="max-h-0 opacity-0"
          >
            <div v-if="paymentsExpanded" class="overflow-hidden border-t border-border">
              <!-- LOADING -->
              <div v-if="paymentsLoading" class="space-y-3 p-5">
                <div
                  v-for="index in 2"
                  :key="index"
                  class="animate-pulse rounded-xl border border-border bg-muted/20 p-4"
                >
                  <div class="flex items-center gap-3">
                    <div class="h-10 w-10 rounded-xl bg-muted" />

                    <div class="flex-1">
                      <div class="h-4 w-1/2 rounded bg-muted" />
                      <div class="mt-2 h-3 w-1/3 rounded bg-muted" />
                    </div>

                    <div class="h-5 w-20 rounded bg-muted" />
                  </div>
                </div>
              </div>

              <!-- ERROR -->
              <div v-else-if="paymentsError" class="p-5">
                <div
                  class="rounded-xl border border-destructive/20 bg-destructive/5 p-4 text-center"
                >
                  <p class="text-sm font-medium text-foreground">Unable to load payment history</p>

                  <p class="mt-1 text-xs text-muted-foreground">
                    The purchase is available, but its transactions could not be retrieved.
                  </p>

                  <Button
                    type="button"
                    variant="outline"
                    size="sm"
                    class="mt-4"
                    @click="refetchPayments()"
                  >
                    <RefreshCw class="mr-2 h-4 w-4" />
                    Try again
                  </Button>
                </div>
              </div>

              <!-- EMPTY -->
              <div v-else-if="purchasePayments.length === 0" class="p-6 text-center">
                <div
                  class="mx-auto flex h-12 w-12 items-center justify-center rounded-xl border border-border bg-muted/30"
                >
                  <CreditCard class="h-5 w-5 text-muted-foreground" />
                </div>

                <p class="mt-4 text-sm font-medium text-foreground">No payments registered yet</p>

                <p class="mt-1 text-xs text-muted-foreground">
                  Payments made against this purchase will appear here.
                </p>

                <Button v-if="canPay" type="button" size="sm" class="mt-5" @click="handlePay">
                  <WalletCards class="mr-2 h-4 w-4" />
                  Register first payment
                </Button>
              </div>

              <!-- PAYMENTS -->
              <div v-else class="relative max-h-[380px] overflow-y-auto p-5">
                <div
                  v-if="paymentsRefreshing"
                  class="absolute inset-x-5 top-5 z-10 flex items-center gap-3 rounded-xl border border-border bg-background/90 p-3 shadow-sm backdrop-blur"
                >
                  <Loader2 class="h-4 w-4 animate-spin text-muted-foreground" />

                  <p class="text-xs text-muted-foreground">Updating payment history...</p>
                </div>

                <div class="space-y-3">
                  <article
                    v-for="payment in purchasePayments"
                    :key="payment.id"
                    class="rounded-xl border border-border bg-background p-4 transition hover:border-primary/20 hover:shadow-sm"
                  >
                    <div class="flex items-start justify-between gap-4">
                      <div class="flex min-w-0 items-center gap-3">
                        <div
                          class="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl border border-border bg-muted/30"
                        >
                          <component
                            :is="getPaymentMethodIcon(payment.payment_method)"
                            class="h-4 w-4 text-muted-foreground"
                          />
                        </div>

                        <div class="min-w-0">
                          <p class="text-sm font-semibold text-foreground">
                            {{ getPaymentMethodLabel(payment.payment_method) }}
                          </p>

                          <p class="mt-1 text-xs text-muted-foreground">
                            {{ formatDate(payment.created_at) }}
                          </p>
                        </div>
                      </div>

                      <div class="text-right">
                        <p class="text-sm font-semibold text-emerald-600 dark:text-emerald-400">
                          {{ formatCurrency(payment.amount) }}
                        </p>

                        <div
                          class="mt-1 flex items-center justify-end gap-1 text-xs text-muted-foreground"
                        >
                          <CircleCheck class="h-3.5 w-3.5" />
                          Registered
                        </div>
                      </div>
                    </div>
                  </article>
                </div>

                <div
                  class="mt-4 flex items-center justify-between rounded-xl border border-border bg-muted/20 p-4"
                >
                  <div>
                    <p class="text-xs text-muted-foreground">Registered payments</p>

                    <p class="mt-1 text-xs text-muted-foreground">
                      {{ purchasePayments.length }}
                      {{ purchasePayments.length === 1 ? 'transaction' : 'transactions' }}
                    </p>
                  </div>

                  <p class="text-lg font-semibold text-emerald-600 dark:text-emerald-400">
                    {{ formatCurrency(paymentsTotal) }}
                  </p>
                </div>
              </div>
            </div>
          </Transition>
        </section>

        <!-- DRAFT CHECKOUT -->
        <section
          v-if="invoice.status_invoice === 'draft'"
          class="overflow-hidden rounded-2xl border border-primary/20 bg-card shadow-lg"
        >
          <div
            class="border-b border-border bg-gradient-to-r from-primary/10 via-primary/5 to-transparent p-5"
          >
            <div class="flex items-start gap-3">
              <div
                class="flex h-11 w-11 shrink-0 items-center justify-center rounded-xl bg-primary text-primary-foreground shadow-sm"
              >
                <ShoppingBag class="h-5 w-5" />
              </div>

              <div>
                <h3 class="text-sm font-semibold text-foreground">
                  {{
                    variant === 'store'
                      ? 'Draft order ready for checkout'
                      : 'Draft purchase ready for checkout'
                  }}
                </h3>

                <p class="mt-1 text-xs leading-5 text-muted-foreground">
                  {{
                    variant === 'store'
                      ? 'Review the personal order before submitting it to the seller.'
                      : 'Review the items and financial summary before submitting this B2B order to the supplier.'
                  }}
                </p>
              </div>
            </div>
          </div>

          <div class="space-y-5 p-5">
            <div class="grid gap-3 sm:grid-cols-3">
              <div class="rounded-xl border border-border bg-muted/20 p-4">
                <p class="text-xs text-muted-foreground">Line items</p>

                <p class="mt-2 text-lg font-semibold text-foreground">
                  {{ purchaseItems.length }}
                </p>
              </div>

              <div class="rounded-xl border border-border bg-muted/20 p-4">
                <p class="text-xs text-muted-foreground">Total units</p>

                <p class="mt-2 text-lg font-semibold text-foreground">
                  {{ totalUnits }}
                </p>
              </div>

              <div class="rounded-xl border border-border bg-muted/20 p-4">
                <p class="text-xs text-muted-foreground">Order total</p>

                <p class="mt-2 text-lg font-semibold text-foreground">
                  {{ formatCurrency(invoice.total_amount) }}
                </p>
              </div>
            </div>

            <div
              class="flex items-start gap-3 rounded-xl border border-amber-500/20 bg-amber-500/10 p-4"
            >
              <ShieldCheck class="mt-0.5 h-4 w-4 shrink-0 text-amber-600 dark:text-amber-400" />

              <div>
                <p class="text-sm font-medium text-amber-700 dark:text-amber-400">
                  Checkout locks purchase editing
                </p>

                <p class="mt-1 text-xs leading-5 text-amber-700/80 dark:text-amber-400/80">
                  After submitting, the draft becomes pending and its items can no longer be changed
                  from the buyer flow.
                </p>
              </div>
            </div>

            <Button
              type="button"
              size="lg"
              class="w-full"
              :disabled="!canCheckout || isCheckingOut || itemsLoading || itemsRefreshing"
              @click="submitCheckout"
            >
              <Loader2 v-if="isCheckingOut" class="mr-2 h-4 w-4 animate-spin" />

              <Send v-else class="mr-2 h-4 w-4" />

              {{
                isCheckingOut
                  ? variant === 'store'
                    ? 'Submitting order...'
                    : 'Submitting purchase...'
                  : variant === 'store'
                    ? 'Submit Store order'
                    : 'Submit purchase to supplier'
              }}

              <ArrowRight v-if="!isCheckingOut" class="ml-2 h-4 w-4" />
            </Button>
          </div>
        </section>

        <!-- PAYMENT PANEL -->
        <Transition
          enter-active-class="transition-all duration-300 ease-out"
          enter-from-class="max-h-0 opacity-0 -translate-y-2"
          enter-to-class="max-h-[900px] opacity-100 translate-y-0"
          leave-active-class="transition-all duration-200 ease-in"
          leave-from-class="max-h-[900px] opacity-100 translate-y-0"
          leave-to-class="max-h-0 opacity-0 -translate-y-2"
        >
          <section
            v-if="paymentExpanded && canPay"
            class="overflow-hidden rounded-2xl border border-primary/20 bg-card shadow-lg"
          >
            <!-- PAYMENT HEADER -->
            <div
              class="border-b border-border bg-gradient-to-r from-primary/10 via-primary/5 to-transparent p-5"
            >
              <div class="flex items-start justify-between gap-4">
                <div class="flex items-center gap-3">
                  <div
                    class="flex h-11 w-11 items-center justify-center rounded-xl bg-primary text-primary-foreground shadow-sm"
                  >
                    <WalletCards class="h-5 w-5" />
                  </div>

                  <div>
                    <h3 class="text-sm font-semibold text-foreground">Register purchase payment</h3>

                    <p class="mt-1 text-xs text-muted-foreground">
                      Apply a partial or full payment to this supplier invoice.
                    </p>
                  </div>
                </div>

                <Button
                  type="button"
                  variant="ghost"
                  size="sm"
                  :disabled="isPaying"
                  @click="paymentExpanded = false"
                >
                  Cancel
                </Button>
              </div>
            </div>

            <div class="space-y-6 p-5">
              <!-- BALANCE CONTEXT -->
              <div class="grid gap-3 sm:grid-cols-2">
                <div class="rounded-xl border border-border bg-muted/20 p-4">
                  <p class="text-xs text-muted-foreground">Outstanding balance</p>

                  <p class="mt-2 text-xl font-semibold text-foreground">
                    {{ formatCurrency(remainingAmount) }}
                  </p>
                </div>

                <div class="rounded-xl border border-border bg-muted/20 p-4">
                  <p class="text-xs text-muted-foreground">Already paid</p>

                  <p class="mt-2 text-xl font-semibold text-emerald-600 dark:text-emerald-400">
                    {{ formatCurrency(invoice.paid_amount) }}
                  </p>
                </div>
              </div>

              <!-- FORM -->
              <div class="grid gap-4 sm:grid-cols-2">
                <div class="space-y-2">
                  <label class="text-sm font-medium text-foreground"> Payment amount </label>

                  <div class="relative">
                    <span
                      class="pointer-events-none absolute left-3 top-1/2 -translate-y-1/2 text-sm text-muted-foreground"
                    >
                      $
                    </span>

                    <Input
                      v-model="paymentForm.amount"
                      type="text"
                      inputmode="decimal"
                      class="pl-7"
                      placeholder="0.00"
                      :disabled="isPaying"
                      @keydown="validateNumberInput"
                    />
                  </div>

                  <p
                    v-if="
                      !isPaying && paymentExpanded && paymentAmountCents > paymentBalanceLimitCents
                    "
                    class="flex items-center gap-1 text-xs text-destructive"
                  >
                    <AlertCircle class="h-3.5 w-3.5" />

                    Amount exceeds outstanding balance.
                  </p>
                </div>

                <div class="space-y-2">
                  <label class="text-sm font-medium text-foreground"> Payment method </label>

                  <Select v-model="paymentForm.method" :disabled="isPaying">
                    <SelectTrigger>
                      <div class="flex items-center gap-2">
                        <component
                          :is="getSelectedPaymentMethodIcon()"
                          class="h-4 w-4 text-muted-foreground"
                        />

                        <span>
                          {{ getSelectedPaymentMethodLabel() }}
                        </span>
                      </div>
                    </SelectTrigger>

                    <SelectContent>
                      <SelectItem value="transfer">
                        <div class="flex items-center gap-2">
                          <Landmark class="h-4 w-4" />
                          Bank transfer
                        </div>
                      </SelectItem>

                      <SelectItem value="card">
                        <div class="flex items-center gap-2">
                          <CreditCard class="h-4 w-4" />
                          Card
                        </div>
                      </SelectItem>

                      <SelectItem value="cash">
                        <div class="flex items-center gap-2">
                          <Banknote class="h-4 w-4" />
                          Cash
                        </div>
                      </SelectItem>
                    </SelectContent>
                  </Select>
                </div>
              </div>

              <!-- QUICK AMOUNTS -->
              <div>
                <p class="text-xs font-medium text-muted-foreground">Quick amount</p>

                <div class="mt-2 flex flex-wrap gap-2">
                  <Button
                    type="button"
                    variant="outline"
                    size="sm"
                    :disabled="isPaying"
                    @click="setPaymentPercentage(0.25)"
                  >
                    25%
                  </Button>

                  <Button
                    type="button"
                    variant="outline"
                    size="sm"
                    :disabled="isPaying"
                    @click="setPaymentPercentage(0.5)"
                  >
                    50%
                  </Button>

                  <Button
                    type="button"
                    variant="outline"
                    size="sm"
                    :disabled="isPaying"
                    @click="setPaymentPercentage(0.75)"
                  >
                    75%
                  </Button>

                  <Button
                    type="button"
                    variant="secondary"
                    size="sm"
                    :disabled="isPaying"
                    @click="setFullPaymentAmount"
                  >
                    Full balance
                  </Button>
                </div>
              </div>

              <!-- TRANSACTION PREVIEW -->
              <div class="rounded-2xl border border-border bg-muted/20 p-5">
                <div class="mb-4 flex items-center gap-2">
                  <ShieldCheck class="h-4 w-4 text-primary" />

                  <p class="text-sm font-semibold text-foreground">Transaction preview</p>
                </div>

                <div class="space-y-3">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-muted-foreground"> Payment </span>

                    <span class="text-sm font-semibold text-foreground">
                      {{ formatCurrency(paymentAmount) }}
                    </span>
                  </div>

                  <div class="flex items-center justify-between">
                    <span class="text-sm text-muted-foreground"> Method </span>

                    <span class="text-sm font-medium text-foreground">
                      {{ getSelectedPaymentMethodLabel() }}
                    </span>
                  </div>

                  <div class="flex items-center justify-between border-t border-border pt-3">
                    <span class="text-sm font-medium text-foreground"> Balance after payment </span>

                    <span
                      class="text-lg font-semibold"
                      :class="
                        willCompletePurchase
                          ? 'text-emerald-600 dark:text-emerald-400'
                          : 'text-foreground'
                      "
                    >
                      {{ formatCurrency(balanceAfterPayment) }}
                    </span>
                  </div>
                </div>

                <div
                  v-if="willCompletePurchase"
                  class="mt-4 flex items-center gap-2 rounded-xl border border-emerald-500/20 bg-emerald-500/10 p-3 text-xs text-emerald-700 dark:text-emerald-400"
                >
                  <CheckCircle2 class="h-4 w-4 shrink-0" />
                  This payment will complete the purchase.
                </div>
              </div>

              <!-- CONFIRM -->
              <Button
                type="button"
                size="lg"
                class="w-full"
                :disabled="!paymentAmountValid || isPaying"
                @click="submitPayment"
              >
                <Loader2 v-if="isPaying" class="mr-2 h-4 w-4 animate-spin" />

                <ShieldCheck v-else class="mr-2 h-4 w-4" />

                {{
                  isPaying
                    ? 'Registering payment...'
                    : `Confirm ${formatCurrency(paymentAmount)} payment`
                }}
              </Button>
            </div>
          </section>
        </Transition>

        <!-- FINANCIAL BREAKDOWN -->
        <section class="rounded-2xl border border-border bg-card p-5 shadow-sm">
          <div class="mb-4 flex items-center gap-2">
            <ReceiptText class="h-4 w-4 text-muted-foreground" />

            <h3 class="text-sm font-semibold text-foreground">Financial summary</h3>
          </div>

          <div class="space-y-3">
            <div class="flex items-start justify-between gap-4">
              <span class="text-sm text-muted-foreground"> Subtotal </span>

              <span class="shrink-0 text-right text-sm font-medium text-foreground">
                {{ formatCurrency(invoice.subtotal) }}
              </span>
            </div>

            <div class="flex items-start justify-between gap-4">
              <span class="text-sm text-muted-foreground"> Taxes </span>

              <span class="shrink-0 text-right text-sm font-medium text-foreground">
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
        <section v-if="canPay && !paymentExpanded">
          <Button type="button" size="lg" class="w-full" @click="handlePay">
            <WalletCards class="mr-2 h-4 w-4" />
            Pay outstanding balance
          </Button>
        </section>
      </div>
    </SheetContent>
  </Sheet>
</template>
