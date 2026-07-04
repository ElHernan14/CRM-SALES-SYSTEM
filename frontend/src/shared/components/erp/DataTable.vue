<script setup lang="ts">
defineProps<{
  columns: {
    key: string;
    label: string;
  }[];
  rows: Record<string, unknown>[];
  selectable?: boolean;
  selectedRows?: unknown[];
  headerChecked?: boolean;
}>();

const emit = defineEmits<{
  toggleRow: [row: Record<string, unknown>];
  toggleAll: [];
}>();
</script>

<template>
  <div class="overflow-hidden rounded-xl border border-border bg-card shadow-sm">
    <table class="w-full border-collapse text-sm">
      <thead class="bg-muted/50">
        <tr>
          <th v-if="selectable" class="w-10 border-b border-border px-4 py-3">
            <input
              type="checkbox"
              class="h-4 w-4 rounded border-border"
              :checked="headerChecked"
              @change="emit('toggleAll')"
            />
          </th>

          <th
            v-for="column in columns"
            :key="column.key"
            class="border-b border-border px-4 py-3 text-left font-medium text-muted-foreground"
          >
            {{ column.label }}
          </th>
        </tr>
      </thead>

      <tbody>
        <tr v-for="(row, index) in rows" :key="index" class="transition hover:bg-muted/40">
          <td v-if="selectable" class="w-10 border-b border-border px-4 py-3">
            <input
              type="checkbox"
              class="h-4 w-4 rounded border-border"
              :checked="selectedRows?.includes(row.id)"
              @change="emit('toggleRow', row)"
            />
          </td>

          <td
            v-for="column in columns"
            :key="column.key"
            class="border-b border-border px-4 py-3 text-foreground last:border-b-0"
          >
            <slot :name="`cell-${column.key}`" :row="row" :value="row[column.key]">
              {{ row[column.key] }}
            </slot>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
