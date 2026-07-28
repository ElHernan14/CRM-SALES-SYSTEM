<script setup lang="ts">
import { ArrowDown, ArrowUp, ChevronsUpDown } from 'lucide-vue-next';

import { computed, ref, useSlots, watchEffect } from 'vue';

type DataTableColumn = {
  key: string;
  label: string;
  headerClass?: string;
  cellClass?: string;
};

const props = withDefaults(
  defineProps<{
    columns: DataTableColumn[];
    rows: Record<string, unknown>[];

    selectable?: boolean;
    selectedRows?: unknown[];

    headerChecked?: boolean;

    sortColumn?: string;
    sortOrder?: 'asc' | 'desc';
    sortableColumns?: string[];

    selectionDisabled?: boolean;

    rowKey?: string;
    tableMinWidth?: string;
  }>(),
  {
    selectable: false,
    selectedRows: () => [],
    headerChecked: false,
    sortableColumns: () => [],
    selectionDisabled: false,
    rowKey: 'id',
    tableMinWidth: '900px',
  }
);

const emit = defineEmits<{
  toggleRow: [row: Record<string, unknown>];
  toggleAll: [];
  sort: [columnKey: string];
}>();

const slots = useSlots();

const hasMobileCards = computed(() => {
  return Boolean(slots['mobile-card']);
});

function isSortable(columnKey: string) {
  return props.sortableColumns.includes(columnKey);
}

function getRowKey(row: Record<string, unknown>, index: number) {
  return String(row[props.rowKey] ?? index);
}

function isRowSelected(row: Record<string, unknown>) {
  return props.selectedRows.includes(row[props.rowKey]);
}

const headerCheckboxRef = ref<HTMLInputElement | null>(null);

const someRowsSelected = computed(() => {
  const selectedCount = props.rows.filter((row) => isRowSelected(row)).length;

  return selectedCount > 0 && selectedCount < props.rows.length;
});

watchEffect(() => {
  if (!headerCheckboxRef.value) return;

  headerCheckboxRef.value.indeterminate = someRowsSelected.value;
});
</script>

<template>
  <div class="min-w-0">
    <!-- MOBILE CARDS -->
    <div v-if="hasMobileCards" class="space-y-3 md:hidden">
      <article
        v-for="(row, index) in rows"
        :key="getRowKey(row, index)"
        class="nexora-surface min-w-0 overflow-hidden rounded-2xl border border-border bg-card shadow-sm"
      >
        <div
          v-if="selectable"
          class="flex items-center justify-between border-b border-border bg-muted/15 px-4 py-3"
        >
          <label class="flex items-center gap-2 text-xs font-medium text-muted-foreground">
            <input
              type="checkbox"
              class="h-4 w-4 rounded border-border"
              :checked="isRowSelected(row)"
              :disabled="selectionDisabled"
              @change="emit('toggleRow', row)"
            />

            Select record
          </label>

          <slot name="mobile-card-header" :row="row" />
        </div>

        <slot name="mobile-card" :row="row" :index="index" />
      </article>
    </div>

    <!-- DESKTOP / TABLET TABLE -->
    <div
      class="nexora-surface max-w-full overflow-x-auto rounded-xl border border-border bg-card shadow-sm"
      :class="hasMobileCards ? 'hidden md:block' : 'block'"
    >
      <table
        class="w-full border-collapse text-sm"
        :style="{
          minWidth: tableMinWidth,
        }"
      >
        <thead class="bg-muted/50">
          <tr>
            <th v-if="selectable" class="w-12 border-b border-border px-4 py-3 text-center">
              <input
                ref="headerCheckboxRef"
                type="checkbox"
                class="h-4 w-4 rounded border-border"
                :checked="headerChecked"
                :disabled="selectionDisabled"
                aria-label="Select all visible rows"
                @change="emit('toggleAll')"
              />
            </th>

            <th
              v-for="column in columns"
              :key="column.key"
              class="whitespace-nowrap border-b border-border px-4 py-3 text-left font-medium text-muted-foreground"
              :class="column.headerClass"
            >
              <button
                v-if="isSortable(column.key)"
                type="button"
                class="inline-flex items-center gap-2 whitespace-nowrap transition hover:text-foreground"
                @click="emit('sort', column.key)"
              >
                <span>
                  {{ column.label }}
                </span>

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
          <tr
            v-for="(row, index) in rows"
            :key="getRowKey(row, index)"
            class="transition hover:bg-muted/40"
          >
            <td v-if="selectable" class="w-12 border-b border-border px-4 py-3 text-center">
              <input
                type="checkbox"
                class="h-4 w-4 rounded border-border"
                :checked="isRowSelected(row)"
                :disabled="selectionDisabled"
                :aria-label="`Select record ${getRowKey(row, index)}`"
                @change="emit('toggleRow', row)"
              />
            </td>

            <td
              v-for="column in columns"
              :key="column.key"
              class="border-b border-border px-4 py-3 align-middle text-foreground"
              :class="column.cellClass"
            >
              <slot :name="`cell-${column.key}`" :row="row" :value="row[column.key]">
                {{ row[column.key] }}
              </slot>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <p v-if="!hasMobileCards" class="mt-2 text-xs text-muted-foreground md:hidden">
      Swipe horizontally to review all columns.
    </p>
  </div>
</template>
