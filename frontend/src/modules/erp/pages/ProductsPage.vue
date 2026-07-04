<script setup lang="ts">
import { Package } from 'lucide-vue-next';

import PageContainer from '@/shared/components/erp/PageContainer.vue';
import PageHeader from '@/shared/components/erp/PageHeader.vue';
import EmptyState from '@/shared/components/erp/EmptyState.vue';
</script>

<template>
  <PageContainer>
    <PageHeader title="Products" description="Manage product catalog, pricing and inventory." />

    <EmptyState
      title="Products module coming soon"
      description="Product management will be implemented in the next sprint."
      :icon="Package"
    />
  </PageContainer>
</template>
