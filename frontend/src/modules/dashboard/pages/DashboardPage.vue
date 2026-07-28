<script setup lang="ts">
import { computed } from 'vue';
import { useRouter } from 'vue-router';

import {
  AlertTriangle,
  ArrowRight,
  BadgeCheck,
  Building2,
  CircleDollarSign,
  FileText,
  Loader2,
  PackageCheck,
  PackagePlus,
  ReceiptText,
  RefreshCw,
  ShoppingBag,
  ShoppingCart,
  Store,
  TriangleAlert,
  WalletCards,
  Warehouse,
} from 'lucide-vue-next';

import { Button } from '@/components/ui/button';

import PageContainer from '@/shared/components/erp/PageContainer.vue';
import EmptyState from '@/shared/components/erp/EmptyState.vue';

import { useCompanyMe } from '@/modules/company/composables/useCompanyMe';
import { useDashboardOverview } from '../composables/useDashboardOverview';

import {
  getCompanyCoverUrl,
  getCompanyLogoUrl,
  getProductImageUrl,
} from '@/shared/utils/assets';

import type { DashboardRecentInvoice } from '../types/dashboard.types';

const router = useRouter();

const {
  data: company,
  isLoading: companyLoading,
  isError: companyError,
  refetch: refetchCompany,
} = useCompanyMe();

const {
  data: dashboard,
  isLoading: dashboardLoading,
  isFetching,
  isError: dashboardError,
  refetch: refetchDashboard,
} = useDashboardOverview();

const isError = computed(() => {
  return companyError.value || dashboardError.value;
});

async function refreshDashboard() {
  await Promise.all([refetchCompany(), refetchDashboard()]);
}

const isLoading = computed(() => {
  return companyLoading.value || dashboardLoading.value;
});

const isRefreshing = computed(() => {
  return isFetching.value && !dashboardLoading.value;
});

const companyLogo = computed(() => {
  return getCompanyLogoUrl(company.value?.logo);
});

const companyCover = computed(() => {
  return getCompanyCoverUrl(company.value?.cover_image);
});

const companyInitial = computed(() => {
  return company.value?.name?.trim().charAt(0).toUpperCase() ?? 'N';
});

const greeting = computed(() => {
  const hour = new Date().getHours();

  if (hour < 12) {
    return 'Good morning';
  }

  if (hour < 18) {
    return 'Good afternoon';
  }

  return 'Good evening';
});

const operationalMessage = computed(() => {
  if (!dashboard.value) {
    return 'Your business workspace is ready.';
  }

  const attention = dashboard.value.attention;

  const criticalCount =
    attention.out_of_stock_products + attention.pending_sales + attention.pending_purchases;

  if (criticalCount === 0) {
    return 'Everything is operating normally.';
  }

  if (criticalCount <= 3) {
    return 'Your business is moving. A few items need attention.';
  }

  return 'Your operation is active, with several items requiring review.';
});

const pulseCards = computed(() => {
  if (!dashboard.value) return [];

  return [
    {
      label: 'Revenue collected',

      value: formatCurrency(dashboard.value.sales.paid_amount),

      description: 'Payments received across company sales.',

      secondary: `${dashboard.value.sales.paid} paid invoices`,

      icon: CircleDollarSign,

      iconClass: 'bg-emerald-500/10 text-emerald-600 dark:text-emerald-400',

      actionLabel: 'Open sales',

      action: () => router.push('/erp/sales'),
    },
    {
      label: 'Outstanding sales',

      value: formatCurrency(dashboard.value.sales.outstanding_amount),

      description: 'Remaining amount expected from customers.',

      secondary: `${dashboard.value.sales.pending} pending invoices`,

      icon: WalletCards,

      iconClass: 'bg-amber-500/10 text-amber-600 dark:text-amber-400',

      actionLabel: 'Review receivables',

      action: () => router.push('/erp/sales'),
    },
    {
      label: 'Company purchases',

      value: formatCurrency(dashboard.value.purchases.total_amount),

      description: 'Organizational B2B purchasing activity.',

      secondary: `${dashboard.value.purchases.pending} awaiting payment`,

      icon: ShoppingBag,

      iconClass: 'bg-blue-500/10 text-blue-600 dark:text-blue-400',

      actionLabel: 'Open purchases',

      action: () => router.push('/erp/purchases'),
    },
    {
      label: 'Available inventory',

      value: dashboard.value.inventory.available_units.toLocaleString(),

      description: 'Units currently available for sale.',

      secondary: `${dashboard.value.inventory.reserved_units} reserved units`,

      icon: PackageCheck,

      iconClass: 'bg-primary/10 text-primary',

      actionLabel: 'Manage products',

      action: () => router.push('/erp/products'),
    },
  ];
});

const attentionItems = computed(() => {
  if (!dashboard.value) return [];

  const attention = dashboard.value.attention;

  return [
    {
      key: 'pending-sales',

      title: 'Sales awaiting payment',

      description: 'Customer invoices remain pending or partially paid.',

      count: attention.pending_sales,

      icon: WalletCards,

      iconClass: 'bg-amber-500/10 text-amber-600 dark:text-amber-400',

      countClass: 'border-amber-500/20 bg-amber-500/10 text-amber-600 dark:text-amber-400',

      actionLabel: 'Review sales',

      action: () =>
        router.push({
          path: '/erp/sales',
          query: {
            status: 'pending',
          },
        }),
    },

    {
      key: 'draft-sales',

      title: 'Draft sales not submitted',

      description: 'Draft invoices still require products or submission.',

      count: attention.draft_sales,

      icon: FileText,

      iconClass: 'bg-blue-500/10 text-blue-600 dark:text-blue-400',

      countClass: 'border-blue-500/20 bg-blue-500/10 text-blue-600 dark:text-blue-400',

      actionLabel: 'Open drafts',

      action: () =>
        router.push({
          path: '/erp/sales',
          query: {
            status: 'draft',
          },
        }),
    },

    {
      key: 'pending-purchases',

      title: 'Purchases awaiting payment',

      description: 'Organizational supplier orders have outstanding balances.',

      count: attention.pending_purchases,

      icon: ShoppingBag,

      iconClass: 'bg-violet-500/10 text-violet-600 dark:text-violet-400',

      countClass: 'border-violet-500/20 bg-violet-500/10 text-violet-600 dark:text-violet-400',

      actionLabel: 'Review purchases',

      action: () =>
        router.push({
          path: '/erp/purchases',
          query: {
            status: 'pending',
          },
        }),
    },

    {
      key: 'stock',

      title: 'Inventory requires attention',

      description: 'Products are running low or currently unavailable.',

      count: attention.low_stock_products + attention.out_of_stock_products,

      icon: TriangleAlert,

      iconClass: 'bg-destructive/10 text-destructive',

      countClass: 'border-destructive/20 bg-destructive/10 text-destructive',

      actionLabel: 'Review inventory',

      action: () => router.push('/erp/products'),
    },
  ];
});

const activeAttentionItems = computed(() => {
  return attentionItems.value.filter((item) => item.count > 0);
});

const attentionTotal = computed(() => {
  return attentionItems.value.reduce((total, item) => total + item.count, 0);
});

const quickActions = [
  {
    title: 'Create product',

    description: 'Add a new product or service to the company catalog.',

    icon: PackagePlus,

    iconClass: 'bg-primary/10 text-primary',

    action: () =>
      router.push({
        path: '/erp/products',
        query: {
          action: 'create',
        },
      }),
  },

  {
    title: 'Create sales invoice',

    description: 'Start a new customer invoice and build its item list.',

    icon: ReceiptText,

    iconClass: 'bg-emerald-500/10 text-emerald-600 dark:text-emerald-400',

    action: () =>
      router.push({
        path: '/erp/sales',
        query: {
          action: 'create',
        },
      }),
  },

  {
    title: 'Browse Marketplace',

    description: 'Discover suppliers and create organizational purchases.',

    icon: ShoppingCart,

    iconClass: 'bg-blue-500/10 text-blue-600 dark:text-blue-400',

    action: () => router.push('/erp/marketplace'),
  },

  {
    title: 'Company identity',

    description: 'Update business information, logo and workspace cover.',

    icon: Building2,

    iconClass: 'bg-violet-500/10 text-violet-600 dark:text-violet-400',

    action: () =>
      router.push({
        name: 'erp-settings-company',
      }),
  },
];

function formatDate(value: string) {
  return new Intl.DateTimeFormat('es-AR', {
    day: '2-digit',
    month: 'short',
    year: 'numeric',
  }).format(new Date(value));
}

function getStatusLabel(status: string) {
  const labels: Record<string, string> = {
    draft: 'Draft',
    pending: 'Awaiting payment',
    paid: 'Paid',
    cancelled: 'Cancelled',
  };

  return labels[status] ?? status;
}

function getStatusClass(status: string) {
  const classes: Record<string, string> = {
    draft: 'border-border bg-muted text-muted-foreground',

    pending: 'border-amber-500/20 bg-amber-500/10 text-amber-600 dark:text-amber-400',

    paid: 'border-emerald-500/20 bg-emerald-500/10 text-emerald-600 dark:text-emerald-400',

    cancelled: 'border-destructive/20 bg-destructive/10 text-destructive',
  };

  return classes[status] ?? classes.draft;
}

function openRecentSale(invoice: DashboardRecentInvoice) {
  router.push({
    path: '/erp/sales',
    query: {
      invoice: String(invoice.id),
    },
  });
}

function openRecentPurchase(invoice: DashboardRecentInvoice) {
  router.push({
    path: '/erp/purchases',
    query: {
      invoice: String(invoice.id),
    },
  });
}

function formatCurrency(value: number) {
  return new Intl.NumberFormat('es-AR', {
    style: 'currency',
    currency: 'ARS',
    maximumFractionDigits: 0,
  }).format(value);
}
</script>

<template>
  <PageContainer>
    <!-- LOADING -->
    <div v-if="isLoading" class="space-y-6">
      <div class="h-[330px] animate-pulse rounded-[2rem] bg-muted" />

      <div class="grid gap-4 md:grid-cols-2 xl:grid-cols-4">
        <div v-for="index in 4" :key="index" class="h-48 animate-pulse rounded-[1.5rem] bg-muted" />
      </div>
    </div>

    <!-- ERROR -->
    <div v-else-if="isError" class="py-12">
      <EmptyState
        title="Unable to load business overview"
        description="There was a problem retrieving the current workspace activity."
        :icon="Building2"
      />

      <div class="mt-5 flex justify-center">
        <Button variant="outline" @click="refreshDashboard">
          <RefreshCw class="mr-2 h-4 w-4" />
          Try again
        </Button>
      </div>
    </div>

    <div v-else-if="company && dashboard" class="space-y-6">
      <!-- COMPANY OPERATIONS HERO -->
      <section
        class="nexora-surface nexora-glow group relative min-h-[330px] overflow-hidden rounded-[2rem] border border-border bg-card shadow-xl"
      >
        <img
          v-if="companyCover"
          :src="companyCover"
          :alt="`${company.name} workspace cover`"
          class="absolute inset-0 h-full w-full object-cover transition duration-700 group-hover:scale-[1.01]"
        />

        <div
          v-else
          class="absolute inset-0 bg-gradient-to-br from-primary/20 via-muted/40 to-background"
        />

        <div
          class="absolute inset-0 bg-gradient-to-r from-background/95 via-background/80 to-background/30"
        />

        <div
          class="absolute inset-x-0 bottom-0 h-44 bg-gradient-to-t from-background/90 to-transparent"
        />

        <div class="relative flex min-h-[330px] flex-col justify-between p-6 sm:p-8">
          <div class="flex items-start justify-between gap-4">
            <div
              class="inline-flex items-center gap-2 rounded-full border border-border/70 bg-background/75 px-3 py-1.5 text-xs font-medium text-muted-foreground shadow-sm backdrop-blur"
            >
              <BadgeCheck class="h-3.5 w-3.5 text-primary" />
              Active business workspace
            </div>

            <Button
              variant="secondary"
              size="sm"
              class="rounded-full bg-background/80 shadow-sm backdrop-blur"
              :disabled="isRefreshing"
              @click="refreshDashboard""
            >
              <Loader2 v-if="isRefreshing" class="mr-2 h-4 w-4 animate-spin" />

              <RefreshCw v-else class="mr-2 h-4 w-4" />

              Refresh overview
            </Button>
          </div>

          <div class="flex flex-col gap-6 lg:flex-row lg:items-end lg:justify-between">
            <div class="flex flex-col gap-5 sm:flex-row sm:items-end">
              <div
                class="flex h-24 w-24 shrink-0 items-center justify-center overflow-hidden rounded-[1.75rem] border border-border bg-background shadow-xl"
              >
                <img
                  v-if="companyLogo"
                  :src="companyLogo"
                  :alt="`${company.name} logo`"
                  class="h-full w-full object-contain p-3"
                />

                <span v-else class="text-3xl font-semibold text-primary">
                  {{ companyInitial }}
                </span>
              </div>

              <div class="max-w-2xl pb-1">
                <p class="text-sm font-semibold text-primary">{{ greeting }}.</p>

                <h1 class="mt-2 text-3xl font-semibold tracking-[-0.045em] sm:text-4xl">
                  {{ company.name }} is operating.
                </h1>

                <p class="mt-3 text-base text-muted-foreground">
                  {{ operationalMessage }}
                </p>

                <div class="mt-4 flex flex-wrap items-center gap-2">
                  <span
                    class="rounded-full border border-primary/20 bg-primary/10 px-3 py-1 text-xs font-semibold text-primary"
                  >
                    {{ company.category }}
                  </span>

                  <span
                    class="rounded-full border border-border bg-background/70 px-3 py-1 text-xs font-medium text-muted-foreground backdrop-blur"
                  >
                    {{ dashboard.sales.total_invoices }}
                    sales ·
                    {{ dashboard.purchases.total_invoices }}
                    purchases
                  </span>
                </div>
              </div>
            </div>

            <div class="flex flex-wrap gap-2">
              <Button
                variant="outline"
                class="rounded-full bg-background/80 backdrop-blur"
                @click="
                  router.push({
                    name: 'erp-settings-company',
                  })
                "
              >
                <Building2 class="mr-2 h-4 w-4" />
                Company profile
              </Button>

              <Button class="rounded-full" @click="router.push('/store')">
                <Store class="mr-2 h-4 w-4" />
                Open Nexora Store
              </Button>
            </div>
          </div>
        </div>
      </section>

      <!-- BUSINESS PULSE -->
      <section>
        <div class="mb-4 flex items-end justify-between gap-4">
          <div>
            <p class="text-xs font-semibold uppercase tracking-[0.16em] text-primary">
              Business pulse
            </p>

            <h2 class="mt-2 text-xl font-semibold tracking-tight">Your operation at a glance</h2>
          </div>

          <p class="hidden items-center gap-2 text-xs text-muted-foreground sm:flex">
            <span class="h-2 w-2 rounded-full bg-emerald-500" />
            Live tenant information
          </p>
        </div>

        <div class="grid gap-4 md:grid-cols-2 xl:grid-cols-4">
          <article
            v-for="card in pulseCards"
            :key="card.label"
            class="nexora-card-interactive nexora-surface group flex min-h-52 flex-col rounded-[1.5rem] border border-border bg-card p-5 shadow-sm transition hover:-translate-y-0.5 hover:border-primary/20 hover:shadow-lg"
          >
            <div class="flex items-start justify-between gap-4">
              <div
                class="flex h-11 w-11 items-center justify-center rounded-xl"
                :class="card.iconClass"
              >
                <component :is="card.icon" class="h-5 w-5" />
              </div>
            </div>

            <div class="mt-5 flex-1">
              <p class="text-sm text-muted-foreground">
                {{ card.label }}
              </p>

              <p class="mt-2 truncate text-2xl font-semibold tracking-[-0.035em]">
                {{ card.value }}
              </p>

              <p class="mt-3 text-xs leading-5 text-muted-foreground">
                {{ card.description }}
              </p>
            </div>

            <div class="mt-5 border-t border-border pt-4">
              <p class="text-xs font-medium text-foreground">
                {{ card.secondary }}
              </p>

              <button
                type="button"
                class="mt-3 inline-flex items-center text-xs font-semibold text-primary transition hover:gap-2"
                @click="card.action"
              >
                {{ card.actionLabel }}
                <ArrowRight class="ml-1 h-3.5 w-3.5" />
              </button>
            </div>
          </article>
        </div>
      </section>

      <!-- ATTENTION CENTER -->
      <section class="nexora-surface overflow-hidden rounded-[1.75rem] border border-border bg-card shadow-sm">
        <div
          class="flex flex-col gap-4 border-b border-border bg-muted/15 p-6 sm:flex-row sm:items-start sm:justify-between"
        >
          <div class="flex items-start gap-4">
            <div
              class="flex h-12 w-12 shrink-0 items-center justify-center rounded-2xl"
              :class="attentionTotal > 0 ? 'bg-amber-500/10' : 'bg-emerald-500/10'"
            >
              <AlertTriangle
                v-if="attentionTotal > 0"
                class="h-5 w-5 text-amber-600 dark:text-amber-400"
              />

              <BadgeCheck v-else class="h-5 w-5 text-emerald-600 dark:text-emerald-400" />
            </div>

            <div>
              <p class="text-xs font-semibold uppercase tracking-[0.16em] text-primary">
                Attention center
              </p>

              <h2 class="mt-2 text-xl font-semibold tracking-tight">
                {{
                  attentionTotal > 0
                    ? 'Items that need your attention'
                    : 'Everything is under control'
                }}
              </h2>

              <p class="mt-2 max-w-2xl text-sm leading-6 text-muted-foreground">
                {{
                  attentionTotal > 0
                    ? 'Review pending financial activity and inventory risks before they affect daily operations.'
                    : 'There are currently no urgent financial or inventory issues in this workspace.'
                }}
              </p>
            </div>
          </div>

          <span
            class="self-start rounded-full border px-3 py-1 min-w-8 justify-center text-sm font-semibold"
            :class="
              attentionTotal > 0
                ? 'border-amber-500/20 bg-amber-500/10 text-amber-600 dark:text-amber-400'
                : 'border-emerald-500/20 bg-emerald-500/10 text-emerald-600 dark:text-emerald-400'
            "
          >
            {{ attentionTotal > 0 ? `${attentionTotal} pending signals` : 'Healthy operation' }}
          </span>
        </div>

        <div v-if="activeAttentionItems.length > 0" class="grid gap-3 p-4 md:grid-cols-2">
          <button
            v-for="item in activeAttentionItems"
            :key="item.key"
            type="button"
            class="group flex items-start gap-4 rounded-2xl border border-border bg-background/50 p-5 text-left transition hover:border-primary/20 hover:bg-muted/25"
            @click="item.action"
          >
            <div
              class="flex h-11 w-11 shrink-0 items-center justify-center rounded-xl"
              :class="item.iconClass"
            >
              <component :is="item.icon" class="h-5 w-5" />
            </div>

            <div class="min-w-0 flex-1">
              <div class="flex items-start justify-between gap-3">
                <div>
                  <p class="text-[15px] font-semibold">
                    {{ item.title }}
                  </p>

                  <p class="mt-1 text-sm leading-5 text-muted-foreground">
                    {{ item.description }}
                  </p>
                </div>

                <span
                  class="inline-flex min-w-8 shrink-0 justify-center rounded-full border px-2.5 py-1 text-sm font-semibold"
                  :class="item.countClass"
                >
                  {{ item.count }}
                </span>
              </div>

              <span class="mt-4 inline-flex items-center text-sm font-semibold text-primary">
                {{ item.actionLabel }}

                <ArrowRight
                  class="ml-1 h-3.5 w-3.5 transition-transform group-hover:translate-x-0.5"
                />
              </span>
            </div>
          </button>
        </div>

        <div v-else class="flex flex-col items-center justify-center px-6 py-12 text-center">
          <div class="flex h-14 w-14 items-center justify-center rounded-2xl bg-emerald-500/10">
            <BadgeCheck class="h-7 w-7 text-emerald-600 dark:text-emerald-400" />
          </div>

          <p class="mt-4 text-[15px] font-semibold">No urgent actions</p>

          <p class="mt-2 max-w-md text-[15px] leading-6 text-muted-foreground">
            Financial flows and inventory are currently operating within normal conditions.
          </p>
        </div>
      </section>

      <!-- RECENT COMMERCIAL ACTIVITY -->
      <section class="grid items-start gap-6 xl:grid-cols-2">
        <!-- RECENT SALES -->
        <article class="nexora-surface overflow-hidden rounded-[1.75rem] border border-border bg-card shadow-sm">
          <div class="flex items-start justify-between gap-4 border-b border-border p-6">
            <div>
              <p class="text-xs font-semibold uppercase tracking-[0.16em] text-primary">
                Sales activity
              </p>

              <h2 class="mt-2 text-xl font-semibold">Recent sales</h2>

              <p class="mt-2 text-sm text-muted-foreground">
                Latest invoices where your company acts as seller.
              </p>
            </div>

            <Button
              variant="ghost"
              size="sm"
              class="rounded-full"
              @click="router.push('/erp/sales')"
            >
              View all
              <ArrowRight class="ml-2 h-4 w-4" />
            </Button>
          </div>

          <div v-if="dashboard.recent_sales.length > 0" class="divide-y divide-border">
            <button
              v-for="sale in dashboard.recent_sales"
              :key="sale.id"
              type="button"
              class="group flex w-full items-center gap-4 p-5 text-left transition hover:bg-muted/25"
              @click="openRecentSale(sale)"
            >
              <div
                class="flex h-11 w-11 shrink-0 items-center justify-center rounded-xl bg-emerald-500/10"
              >
                <ReceiptText class="h-5 w-5 text-emerald-600 dark:text-emerald-400" />
              </div>

              <div class="min-w-0 flex-1">
                <div class="flex items-start justify-between gap-4">
                  <div class="min-w-0">
                    <p class="truncate text-base font-semibold text-foreground">
                      {{ sale.counterparty }}
                    </p>

                    <div class="mt-1.5 flex flex-wrap items-center gap-2">
                      <span class="text-sm font-semibold text-foreground"> #NX-{{ sale.id }} </span>

                      <span class="h-1 w-1 rounded-full bg-muted-foreground/50" />

                      <span class="text-xs text-muted-foreground">
                        {{ formatDate(sale.created_at) }}
                      </span>
                    </div>
                  </div>

                  <div class="shrink-0 text-right">
                    <p
                      class="text-[11px] font-medium uppercase tracking-wide text-muted-foreground"
                    >
                      Total
                    </p>

                    <p class="mt-1 text-base font-semibold tracking-tight text-foreground">
                      {{ formatCurrency(sale.total_amount) }}
                    </p>
                  </div>
                </div>

                <div class="mt-3 flex flex-wrap items-center justify-between gap-3">
                  <span
                    class="inline-flex rounded-full border px-2.5 py-1 text-[11px] font-semibold"
                    :class="getStatusClass(sale.status_invoice)"
                  >
                    {{ getStatusLabel(sale.status_invoice) }}
                  </span>

                  <div v-if="sale.remaining_amount > 0" class="text-right">
                    <p class="text-[11px] text-muted-foreground">Outstanding</p>

                    <p class="mt-0.5 text-sm font-semibold text-amber-600 dark:text-amber-400">
                      {{ formatCurrency(sale.remaining_amount) }}
                    </p>
                  </div>
                </div>
              </div>
            </button>
          </div>

          <div v-else class="flex flex-col items-center justify-center p-10 text-center">
            <ReceiptText class="h-7 w-7 text-muted-foreground/50" />

            <p class="mt-4 text-sm font-semibold">No sales activity yet</p>

            <p class="mt-2 text-xs text-muted-foreground">
              New customer invoices will appear here.
            </p>
          </div>
        </article>

        <!-- RECENT PURCHASES -->
        <article class="nexora-surface overflow-hidden rounded-[1.75rem] border border-border bg-card shadow-sm">
          <div class="flex items-start justify-between gap-4 border-b border-border p-6">
            <div>
              <p class="text-xs font-semibold uppercase tracking-[0.16em] text-primary">
                Purchasing activity
              </p>

              <h2 class="mt-2 text-xl font-semibold">Recent purchases</h2>

              <p class="mt-2 text-sm text-muted-foreground">
                Latest organizational purchases made through ERP.
              </p>
            </div>

            <Button
              variant="ghost"
              size="sm"
              class="rounded-full"
              @click="router.push('/erp/purchases')"
            >
              View all
              <ArrowRight class="ml-2 h-4 w-4" />
            </Button>
          </div>

          <div v-if="dashboard.recent_purchases.length > 0" class="divide-y divide-border">
            <button
              v-for="purchase in dashboard.recent_purchases"
              :key="purchase.id"
              type="button"
              class="group flex w-full items-center gap-4 p-5 text-left transition hover:bg-muted/25"
              @click="openRecentPurchase(purchase)"
            >
              <div
                class="flex h-11 w-11 shrink-0 items-center justify-center rounded-xl bg-blue-500/10"
              >
                <ShoppingBag class="h-5 w-5 text-blue-600 dark:text-blue-400" />
              </div>

              <div class="min-w-0 flex-1">
                <div class="flex items-start justify-between gap-4">
                  <div class="min-w-0">
                    <p class="truncate text-base font-semibold text-foreground">
                      {{ purchase.counterparty }}
                    </p>

                    <div class="mt-1.5 flex flex-wrap items-center gap-2">
                      <span class="text-sm font-semibold text-foreground">
                        #NX-{{ purchase.id }}
                      </span>

                      <span class="h-1 w-1 rounded-full bg-muted-foreground/50" />

                      <span class="text-xs text-muted-foreground">
                        {{ formatDate(purchase.created_at) }}
                      </span>
                    </div>
                  </div>

                  <div class="shrink-0 text-right">
                    <p
                      class="text-[11px] font-medium uppercase tracking-wide text-muted-foreground"
                    >
                      Total
                    </p>

                    <p class="mt-1 text-base font-semibold tracking-tight text-foreground">
                      {{ formatCurrency(purchase.total_amount) }}
                    </p>
                  </div>
                </div>

                <div class="mt-3 flex flex-wrap items-center justify-between gap-3">
                  <span
                    class="inline-flex rounded-full border px-2.5 py-1 text-[11px] font-semibold"
                    :class="getStatusClass(purchase.status_invoice)"
                  >
                    {{ getStatusLabel(purchase.status_invoice) }}
                  </span>

                  <div v-if="purchase.remaining_amount > 0" class="text-right">
                    <p class="text-[11px] text-muted-foreground">Outstanding</p>

                    <p class="mt-0.5 text-sm font-semibold text-amber-600 dark:text-amber-400">
                      {{ formatCurrency(purchase.remaining_amount) }}
                    </p>
                  </div>
                </div>
              </div>
            </button>
          </div>

          <div v-else class="flex flex-col items-center justify-center p-10 text-center">
            <ShoppingBag class="h-7 w-7 text-muted-foreground/50" />

            <p class="mt-4 text-sm font-semibold">No organizational purchases yet</p>

            <p class="mt-2 text-xs text-muted-foreground">
              ERP and Marketplace purchases will appear here.
            </p>
          </div>
        </article>
      </section>

      <!-- INVENTORY + QUICK ACTIONS -->
      <section class="grid items-start gap-6 xl:grid-cols-[minmax(0,1.2fr)_minmax(340px,0.8fr)]">
        <!-- INVENTORY SNAPSHOT -->
        <article class="nexora-surface overflow-hidden rounded-[1.75rem] border border-border bg-card shadow-sm">
          <div class="flex items-start justify-between gap-4 border-b border-border p-6">
            <div>
              <p class="text-xs font-semibold uppercase tracking-[0.16em] text-primary">
                Inventory snapshot
              </p>

              <h2 class="mt-2 text-xl font-semibold">Stock health</h2>

              <p class="mt-2 text-sm text-muted-foreground">
                Products with the lowest current availability.
              </p>
            </div>

            <Button
              variant="ghost"
              size="sm"
              class="rounded-full"
              @click="router.push('/erp/products')"
            >
              Manage inventory
              <ArrowRight class="ml-2 h-4 w-4" />
            </Button>
          </div>

          <div class="grid grid-cols-2 gap-px border-b border-border bg-border sm:grid-cols-4">
            <div class="bg-card p-4">
              <p class="text-xs text-muted-foreground">Active products</p>

              <p class="mt-2 text-xl font-semibold">
                {{ dashboard.inventory.active_products }}
              </p>
            </div>

            <div class="bg-card p-4">
              <p class="text-xs text-muted-foreground">Available units</p>

              <p class="mt-2 text-xl font-semibold">
                {{ dashboard.inventory.available_units }}
              </p>
            </div>

            <div class="bg-card p-4">
              <p class="text-xs text-muted-foreground">Low stock</p>

              <p class="mt-2 text-xl font-semibold text-amber-600 dark:text-amber-400">
                {{ dashboard.inventory.low_stock_products }}
              </p>
            </div>

            <div class="bg-card p-4">
              <p class="text-xs text-muted-foreground">Out of stock</p>

              <p class="mt-2 text-xl font-semibold text-destructive">
                {{ dashboard.inventory.out_of_stock_products }}
              </p>
            </div>
          </div>

          <div v-if="dashboard.low_stock_products.length > 0" class="divide-y divide-border">
            <button
              v-for="product in dashboard.low_stock_products"
              :key="product.id"
              type="button"
              class="flex w-full items-center gap-4 p-5 text-left transition hover:bg-muted/25"
              @click="router.push('/erp/products')"
            >
              <div
                class="flex h-12 w-12 shrink-0 items-center justify-center overflow-hidden rounded-xl border border-border bg-muted/30"
              >
                <img
                  v-if="getProductImageUrl(product.image_path)"
                  :src="getProductImageUrl(product.image_path)!"
                  :alt="product.name"
                  class="h-full w-full object-cover"
                />

                <Warehouse v-else class="h-5 w-5 text-muted-foreground" />
              </div>

              <div class="min-w-0 flex-1">
                <div class="flex items-start justify-between gap-4">
                  <div class="min-w-0">
                    <p class="truncate text-sm font-semibold">
                      {{ product.name }}
                    </p>

                    <p class="mt-1 truncate text-xs text-muted-foreground">
                      {{ product.category }} ·
                      {{ product.type }}
                    </p>
                  </div>

                  <span
                    class="shrink-0 rounded-full border px-2.5 py-1 text-xs font-semibold"
                    :class="
                      product.available_stock <= 0
                        ? 'border-destructive/20 bg-destructive/10 text-destructive'
                        : 'border-amber-500/20 bg-amber-500/10 text-amber-600 dark:text-amber-400'
                    "
                  >
                    {{
                      product.available_stock <= 0
                        ? 'Out of stock'
                        : `${product.available_stock} available`
                    }}
                  </span>
                </div>

                <div class="mt-3 flex items-center gap-4 text-xs text-muted-foreground">
                  <span> Stock: {{ product.stock }} </span>

                  <span>
                    Reserved:
                    {{ product.reserved_stock }}
                  </span>
                </div>
              </div>
            </button>
          </div>

          <div v-else class="flex flex-col items-center justify-center p-10 text-center">
            <PackageCheck class="h-7 w-7 text-emerald-600 dark:text-emerald-400" />

            <p class="mt-4 text-sm font-semibold">Inventory is healthy</p>

            <p class="mt-2 text-xs text-muted-foreground">
              No products currently require stock attention.
            </p>
          </div>
        </article>

        <!-- QUICK ACTIONS -->
        <aside class="nexora-surface overflow-hidden rounded-[1.75rem] border border-border bg-card shadow-sm">
          <div class="border-b border-border p-6">
            <p class="text-xs font-semibold uppercase tracking-[0.16em] text-primary">
              Quick actions
            </p>

            <h2 class="mt-2 text-xl font-semibold">Move your business forward</h2>

            <p class="mt-2 text-sm leading-6 text-muted-foreground">
              Start common workflows directly from your operations center.
            </p>
          </div>

          <div class="divide-y divide-border">
            <button
              v-for="action in quickActions"
              :key="action.title"
              type="button"
              class="group flex w-full items-start gap-4 p-5 text-left transition hover:bg-muted/25"
              @click="action.action"
            >
              <div
                class="flex h-11 w-11 shrink-0 items-center justify-center rounded-xl"
                :class="action.iconClass"
              >
                <component :is="action.icon" class="h-5 w-5" />
              </div>

              <div class="min-w-0 flex-1">
                <div class="flex items-center justify-between gap-3">
                  <p class="text-sm font-semibold">
                    {{ action.title }}
                  </p>

                  <ArrowRight
                    class="h-4 w-4 text-muted-foreground transition group-hover:translate-x-0.5 group-hover:text-primary"
                  />
                </div>

                <p class="mt-1 text-xs leading-5 text-muted-foreground">
                  {{ action.description }}
                </p>
              </div>
            </button>
          </div>
        </aside>
      </section>
    </div>
  </PageContainer>
</template>
