<script setup lang="ts">
defineProps<{
  title?: string;
  description?: string;
}>();
</script>

<template>
  <div
    class="nexora-surface flex flex-col gap-6 rounded-xl border border-border bg-card px-6 py-5 shadow-sm md:flex-row md:items-center md:justify-between"
  >
    <div v-if="title || description">
      <h2 v-if="title" class="text-sm font-medium text-foreground">
        {{ title }}
      </h2>

      <p v-if="description" class="mt-1 text-sm text-muted-foreground">
        {{ description }}
      </p>
    </div>

    <div class="flex flex-wrap items-center gap-3">
      <slot />
    </div>
  </div>
</template>
