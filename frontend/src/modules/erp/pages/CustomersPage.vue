<script setup lang="ts">
import { Users } from 'lucide-vue-next';

import PageContainer from '@/shared/components/erp/PageContainer.vue';
import PageHeader from '@/shared/components/erp/PageHeader.vue';
import EmptyState from '@/shared/components/erp/EmptyState.vue';
</script>

<template>
  <PageContainer>
    <PageHeader title="Customers" description="Manage and track your company's customers." />

    <EmptyState
      title="Customers module coming soon"
      description="Customer management workflows will be connected after the CRM integration."
      :icon="Users"
    />
  </PageContainer>
</template>
