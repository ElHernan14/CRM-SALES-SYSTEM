<script setup lang="ts">
import { Settings } from 'lucide-vue-next';

import PageContainer from '@/shared/components/erp/PageContainer.vue';
import PageHeader from '@/shared/components/erp/PageHeader.vue';
import EmptyState from '@/shared/components/erp/EmptyState.vue';
</script>

<template>
  <PageContainer>
    <PageHeader
      title="Settings"
      description="Configure your company's preferences and system options."
    />

    <EmptyState
      title="Settings module coming soon"
      description="Configuration workflows will be connected after the core system setup."
      :icon="Settings"
    />
  </PageContainer>
</template>
