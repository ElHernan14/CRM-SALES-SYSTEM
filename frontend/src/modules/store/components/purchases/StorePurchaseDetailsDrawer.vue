<script setup lang="ts">
import PurchaseDetailsDrawer from '@/modules/purchases/components/PurchaseDetailsDrawer.vue';

import type { StorePurchase } from '../../types/store-purchase.types';

defineProps<{
  open: boolean;
  purchase: StorePurchase | null;
  initialAction?: 'overview' | 'payment';
}>();

const emit = defineEmits<{
  'update:open': [value: boolean];
}>();
</script>

<template>
  <PurchaseDetailsDrawer
    :open="open"
    :purchase="purchase"
    :initial-action="initialAction === 'payment' ? 'payment' : 'overview'"
    variant="store"
    @update:open="emit('update:open', $event)"
  />
</template>
