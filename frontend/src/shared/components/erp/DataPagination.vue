<script setup lang="ts">
import { Button } from '@/components/ui/button';

defineProps<{
  page: number;
  totalPages: number;
  total?: number;
}>();

const emit = defineEmits<{
  previous: [];
  next: [];
}>();
</script>

<template>
  <div
    class="flex flex-col gap-3 border-t border-border pt-4 md:flex-row md:items-center md:justify-between"
  >
    <p class="text-sm text-muted-foreground">
      Page {{ page }} of {{ totalPages }}
      <span v-if="typeof total === 'number'"> · {{ total }} results </span>
    </p>

    <div class="flex items-center gap-2">
      <Button variant="outline" size="sm" :disabled="page <= 1" @click="emit('previous')">
        Previous
      </Button>

      <Button variant="outline" size="sm" :disabled="page >= totalPages" @click="emit('next')">
        Next
      </Button>
    </div>
  </div>
</template>
