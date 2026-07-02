<script setup lang="ts">
defineProps<{
  columns: string[]
  rows: Record<string, string | number>[]
}>()
</script>

<template>
  <div class="overflow-hidden rounded-xl border border-border bg-card shadow-sm">
    <table class="w-full border-collapse text-sm">
      <thead class="bg-muted/50">
        <tr>
          <th
            v-for="column in columns"
            :key="column"
            class="border-b border-border px-4 py-3 text-left font-medium text-muted-foreground"
          >
            {{ column }}
          </th>
        </tr>
      </thead>

      <tbody>
        <tr
          v-for="(row, index) in rows"
          :key="index"
          class="transition hover:bg-muted/40"
        >
          <td
            v-for="column in columns"
            :key="column"
            class="border-b border-border px-4 py-3 text-foreground last:border-b-0"
          >
            {{ row[column] }}
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>