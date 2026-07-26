<script setup lang="ts">
import { ChevronLeft, ChevronRight } from 'lucide-vue-next';

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
    class="flex min-w-0 flex-col gap-3 border-t border-border pt-4 sm:flex-row sm:items-center sm:justify-between"
  >
    <p class="text-center text-sm text-muted-foreground sm:text-left">
      Page
      <span class="font-medium text-foreground">
        {{ page }}
      </span>
      of
      <span class="font-medium text-foreground">
        {{ totalPages }}
      </span>

      <span v-if="typeof total === 'number'">
        · {{ total }}
        {{ total === 1 ? 'result' : 'results' }}
      </span>
    </p>

    <div class="grid grid-cols-2 gap-2 sm:flex sm:items-center">
      <Button
        type="button"
        variant="outline"
        size="sm"
        class="w-full rounded-full sm:w-auto"
        :disabled="page <= 1"
        @click="emit('previous')"
      >
        <ChevronLeft class="mr-1 h-4 w-4" />
        Previous
      </Button>

      <Button
        type="button"
        variant="outline"
        size="sm"
        class="w-full rounded-full sm:w-auto"
        :disabled="page >= totalPages"
        @click="emit('next')"
      >
        Next
        <ChevronRight class="ml-1 h-4 w-4" />
      </Button>
    </div>
  </div>
</template>
