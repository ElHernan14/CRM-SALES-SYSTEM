<script setup lang="ts">
import { Package } from 'lucide-vue-next';

import PageContainer from '@/shared/components/erp/PageContainer.vue';
import PageHeader from '@/shared/components/erp/PageHeader.vue';
import EmptyState from '@/shared/components/erp/EmptyState.vue';
</script>

<template>
  <PageContainer>
    <PageHeader title="Invoices" description="Manage invoices and billing." />

    <EmptyState
      title="Invoices module coming soon"
      description="Invoice management will be implemented in the next sprint."
      :icon="Package"
    />
  </PageContainer>
</template>
