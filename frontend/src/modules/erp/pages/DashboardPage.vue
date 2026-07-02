<script setup lang="ts">
import {
  CreditCard,
  FileText,
  Package,
  Wallet,
  FileSearch,
} from "lucide-vue-next"

import PageContainer from "@/shared/components/erp/PageContainer.vue"
import PageHeader from "@/shared/components/erp/PageHeader.vue"
import StatCard from "@/shared/components/erp/StatCard.vue"
import EmptyState from "@/shared/components/erp/EmptyState.vue"
import SectionCard from "@/shared/components/erp/SectionCard.vue"
import { Button } from "@/components/ui/button"
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

      <StatCard
        title="Products"
        value="342"
        description="Available in inventory"
        :icon="Package"
      />

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
        <Button variant="outline" size="sm">
          View documentation
        </Button>
      </EmptyState>
    </SectionCard>    
  </PageContainer>
</template>