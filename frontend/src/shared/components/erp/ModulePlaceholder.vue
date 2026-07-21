<script setup lang="ts">
import PageContainer from '@/shared/components/erp/PageContainer.vue';
import PageHeader from '@/shared/components/erp/PageHeader.vue';
import EmptyState from '@/shared/components/erp/EmptyState.vue';

interface Props {
  title: string;
  description: string;
  emptyTitle: string;
  emptyDescription: string;
  icon: any;
}

const props = defineProps<Props>();
</script>

<template>
  <PageContainer>
    <PageHeader :title="props.title" :description="props.description" />

    <EmptyState
      :title="props.emptyTitle"
      :description="props.emptyDescription"
      :icon="props.icon"
    />
  </PageContainer>
</template>
