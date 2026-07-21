<script setup lang="ts">
import { ShoppingBag } from 'lucide-vue-next';
import ModulePlaceholder from '@/shared/components/erp/ModulePlaceholder.vue';
</script>

<template>
  <ModulePlaceholder
    title="Inventory"
    description="Track invoices where your company acts as the buyer."
    emptyTitle="Inventory module coming soon"
    emptyDescription="Inventory workflows will be connected after the B2B marketplace flow."
    :icon="ShoppingBag"
  />
</template>
