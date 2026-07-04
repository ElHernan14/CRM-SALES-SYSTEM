<script setup lang="ts">
import { Package } from 'lucide-vue-next';

import PageContainer from '@/shared/components/erp/PageContainer.vue';
import PageHeader from '@/shared/components/erp/PageHeader.vue';
import EmptyState from '@/shared/components/erp/EmptyState.vue';
</script>

<template>
  <PageContainer>
    <PageHeader title="Settings" description="Manage system settings and configurations." />

    <EmptyState
      title="Settings module coming soon"
      description="System settings management will be implemented in the next sprint."
      :icon="Package"
    />
  </PageContainer>
</template>
