<script setup lang="ts">
import { ShoppingBag } from 'lucide-vue-next';

import PageContainer from '@/shared/components/erp/PageContainer.vue';
import PageHeader from '@/shared/components/erp/PageHeader.vue';
import EmptyState from '@/shared/components/erp/EmptyState.vue';
</script>

<template>
  <PageContainer>
    <PageHeader
      title="Purchases"
      description="Track invoices where your company acts as the buyer."
    />

    <EmptyState
      title="Purchases module coming soon"
      description="Purchase workflows will be connected after the B2B marketplace flow."
      :icon="ShoppingBag"
    />
  </PageContainer>
</template>
