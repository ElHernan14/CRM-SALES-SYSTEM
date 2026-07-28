<script setup lang="ts">
defineProps<{
  title?: string
  description?: string
}>()
</script>

<template>
  <section class="nexora-surface rounded-xl border border-border bg-card shadow-sm">
    <div
      v-if="title || description || $slots.actions"
      class="flex items-start justify-between gap-4 border-b border-border p-5"
    >
      <div>
        <h2
          v-if="title"
          class="text-sm font-medium text-foreground"
        >
          {{ title }}
        </h2>

        <p
          v-if="description"
          class="mt-1 text-sm text-muted-foreground"
        >
          {{ description }}
        </p>
      </div>

      <div v-if="$slots.actions">
        <slot name="actions" />
      </div>
    </div>

    <div class="p-5">
      <slot />
    </div>
  </section>
</template>
