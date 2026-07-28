<script setup lang="ts">
import type { Component } from "vue"

defineProps<{
  title: string
  description?: string
  icon?: Component
}>()
</script>

<template>
  <div class="nexora-surface flex min-h-64 flex-col items-center justify-center rounded-xl border border-dashed border-border bg-card p-10 text-center">
    <div
      v-if="icon"
      class="mb-4 rounded-xl border border-border bg-muted p-3 text-muted-foreground"
    >
      <component :is="icon" class="h-5 w-5" />
    </div>

    <h3 class="text-sm font-medium text-foreground">
      {{ title }}
    </h3>

    <p
      v-if="description"
      class="mt-1 max-w-sm text-sm text-muted-foreground"
    >
      {{ description }}
    </p>

    <div
      v-if="$slots.default"
      class="mt-4"
    >
      <slot />
    </div>
  </div>
</template>
