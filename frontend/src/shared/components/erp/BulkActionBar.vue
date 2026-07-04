<script setup lang="ts">
defineProps<{
  selectedCount: number;
}>();
</script>

<template>
  <div
    v-if="selectedCount > 0"
    class="mb-4 flex items-center justify-between rounded-xl border border-border bg-muted/40 px-4 py-3"
  >
    <p class="text-sm text-muted-foreground">{{ selectedCount }} selected</p>

    <div class="flex items-center gap-2">
      <slot />
    </div>
  </div>
</template>
