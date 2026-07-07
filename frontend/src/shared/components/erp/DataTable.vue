<script setup lang="ts">
import { ArrowDown, ArrowUp, ChevronsUpDown } from 'lucide-vue-next';

const props = defineProps<{
  columns: {
    key: string;
    label: string;
  }[];
  rows: Record<string, unknown>[];
  selectable?: boolean;
  selectedRows?: unknown[];
  headerChecked?: boolean;
  sortColumn?: string;
  sortOrder?: 'asc' | 'desc';
  sortableColumns?: string[];
}>();

const emit = defineEmits<{
  toggleRow: [row: Record<string, unknown>];
  toggleAll: [];
  sort: [columnKey: string];
}>();

function isSortable(columnKey: string) {
  return props.sortableColumns?.includes(columnKey);
}
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
            <button
              v-if="isSortable(column.key)"
              type="button"
              class="inline-flex items-center gap-2 transition hover:text-foreground"
              @click="emit('sort', column.key)"
            >
              <span>{{ column.label }}</span>

              <ArrowUp
                v-if="sortColumn === column.key && sortOrder === 'asc'"
                class="h-3.5 w-3.5"
              />

              <ArrowDown
                v-else-if="sortColumn === column.key && sortOrder === 'desc'"
                class="h-3.5 w-3.5"
              />

              <ChevronsUpDown v-else class="h-3.5 w-3.5 opacity-50" />
            </button>

            <span v-else>
              {{ column.label }}
            </span>
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
