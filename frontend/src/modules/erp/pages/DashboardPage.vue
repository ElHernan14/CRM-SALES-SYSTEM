<script setup lang="ts">
import {
  CreditCard,
  FileText,
  Package,
  Wallet,
  FileSearch,
  Plus,
  RefreshCw,
} from 'lucide-vue-next';

import PageContainer from '@/shared/components/erp/PageContainer.vue';
import PageHeader from '@/shared/components/erp/PageHeader.vue';
import StatCard from '@/shared/components/erp/StatCard.vue';
import EmptyState from '@/shared/components/erp/EmptyState.vue';
import SectionCard from '@/shared/components/erp/SectionCard.vue';
import PageActionBarTable from '@/shared/components/erp/PageActionBarTable.vue';
import DataPreviewTable from '@/shared/components/erp/DataPreviewTable.vue';
import { Button } from '@/components/ui/button';

const recentInvoiceColumns = ['Invoice', 'Client', 'Status', 'Amount'];

const recentInvoiceRows = [
  {
    Invoice: '#INV-1024',
    Client: 'Acme Corp',
    Status: 'Pending',
    Amount: '$1,250',
  },
  {
    Invoice: '#INV-1023',
    Client: 'Northwind',
    Status: 'Paid',
    Amount: '$820',
  },
  {
    Invoice: '#INV-1022',
    Client: 'Globex',
    Status: 'Draft',
    Amount: '$430',
  },
];
</script>

<template>
  <PageContainer>
    <PageHeader
      title="Dashboard"
      description="Overview of your commerce, invoices, payments and inventory."
    />

    <div class="grid gap-4 md:grid-cols-2 xl:grid-cols-4">
      <StatCard
        title="Revenue"
        value="$24,520"
        description="Total revenue this month"
        :icon="Wallet"
      />

      <StatCard
        title="Invoices"
        value="128"
        description="Active commercial documents"
        :icon="FileText"
      />

      <StatCard title="Products" value="342" description="Available in inventory" :icon="Package" />

      <StatCard
        title="Payments"
        value="86"
        description="Registered payment records"
        :icon="CreditCard"
      />
    </div>

    <SectionCard
      title="Recent activity"
      description="Latest commercial operations across your workspace."
    >
      <EmptyState
        title="No recent activity yet"
        description="Once invoices, payments or products start moving, activity will appear here."
        :icon="FileSearch"
      >
        <Button variant="outline" size="sm"> View documentation </Button>
      </EmptyState>
    </SectionCard>

    <PageActionBarTable
      title="Workspace actions"
      description="Quick actions for daily ERP operations."
    >
      <Button variant="outline" size="sm">
        <RefreshCw class="mr-2 h-4 w-4" />
        Refresh
      </Button>

      <Button size="sm">
        <Plus class="mr-2 h-4 w-4" />
        New invoice
      </Button>
    </PageActionBarTable>

    <SectionCard
      title="Recent invoices"
      description="Latest invoice activity across your workspace."
    >
      <DataPreviewTable :columns="recentInvoiceColumns" :rows="recentInvoiceRows" />
    </SectionCard>
  </PageContainer>
</template>
