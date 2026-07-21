<script setup lang="ts">
import { BarChart3 } from 'lucide-vue-next';

import PageContainer from '@/shared/components/erp/PageContainer.vue';
import PageHeader from '@/shared/components/erp/PageHeader.vue';
import EmptyState from '@/shared/components/erp/EmptyState.vue';
</script>

<template>
  <PageContainer>
    <PageHeader
      title="Analytics"
      description="Visualize performance metrics and business insights."
    />

    <EmptyState
      title="Analytics module coming soon"
      description="Data visualization and reporting will be connected after the BI integration."
      :icon="BarChart3"
    />
  </PageContainer>
</template>
