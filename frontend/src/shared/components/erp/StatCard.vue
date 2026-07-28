<script setup lang="ts">
import type { Component } from "vue"

defineProps<{
  title: string
  value: string | number
  description?: string
  icon?: Component
}>()
</script>

<template>
  <div class="nexora-surface rounded-xl border border-border bg-card p-5 shadow-sm">
    <div class="flex items-start justify-between gap-4">
      <div>
        <p class="text-sm text-muted-foreground">
          {{ title }}
        </p>

        <p class="mt-2 text-2xl font-semibold tracking-tight text-foreground">
          {{ value }}
        </p>

        <p
          v-if="description"
          class="mt-1 text-xs text-muted-foreground"
        >
          {{ description }}
        </p>
      </div>

      <div
        v-if="icon"
        class="rounded-lg border border-border bg-muted p-2 text-muted-foreground"
      >
        <component :is="icon" class="h-4 w-4" />
      </div>
    </div>
  </div>
</template>
