<script setup lang="ts">
import { brand } from '@/shared/config/brand';

withDefaults(
  defineProps<{
    compact?: boolean;
  }>(),
  {
    compact: false,
  }
);
</script>

<template>
  <RouterLink to="/erp/dashboard" class="flex min-w-0 items-center gap-3">
    <div
      class="nexora-gradient nexora-glow flex h-10 w-10 shrink-0 items-center justify-center rounded-xl bg-primary font-bold text-primary-foreground shadow-sm"
    >
      N
    </div>

    <div v-if="!compact" class="min-w-0">
      <p class="nexora-gradient-text truncate text-lg font-semibold tracking-[-0.03em]">
        {{ brand.name }}
      </p>

      <p class="truncate text-[11px] text-muted-foreground">Commerce &amp; Operations</p>
    </div>
  </RouterLink>
</template>
