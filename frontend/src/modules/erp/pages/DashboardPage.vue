<script setup lang="ts">
import {
  CreditCard,
  FileText,
  Package,
  Wallet,
} from "lucide-vue-next"

import PageContainer from "@/shared/components/erp/PageContainer.vue"
import PageHeader from "@/shared/components/erp/PageHeader.vue"
import StatCard from "@/shared/components/erp/StatCard.vue"
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
  </PageContainer>
</template>