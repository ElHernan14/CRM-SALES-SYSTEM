<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import { Package, RefreshCw, MoreHorizontal } from 'lucide-vue-next';

import { Button } from '@/components/ui/button';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';

import PageContainer from '@/shared/components/erp/PageContainer.vue';
import PageHeader from '@/shared/components/erp/PageHeader.vue';
import SectionCard from '@/shared/components/erp/SectionCard.vue';
import EmptyState from '@/shared/components/erp/EmptyState.vue';
import DataTable from '@/shared/components/erp/DataTable.vue';
import ProductsFilters from '../components/ProductsFilters.vue';
import DataPagination from '@/shared/components/erp/DataPagination.vue';

import { storeToRefs } from 'pinia';
import { useAuthStore } from '@/modules/auth/stores/auth.store';

import { useProducts } from '../composables/useProducts';

// User store
const auth = useAuthStore();
const { user } = storeToRefs(auth);

// Filters
const search = ref('');
const page = ref(1);
const limit = ref(10);
const type = ref('');
const minPrice = ref('');
const maxPrice = ref('');

const productParams = computed(() => ({
  search: search.value || undefined,
  type: type.value || undefined,
  min_price: minPrice.value ? Number(minPrice.value) : undefined,
  max_price: maxPrice.value ? Number(maxPrice.value) : undefined,
  company_id: user.value?.company_id,
  page: page.value,
  limit: limit.value,
}));

watch([search, type, minPrice, maxPrice], () => {
  page.value = 1;
});

function clearFilters() {
  search.value = '';
  type.value = '';
  minPrice.value = '';
  maxPrice.value = '';
  page.value = 1;
}

// Fetch products
const { data, isLoading, isError, refetch } = useProducts(productParams);

// Table columns and rows
const columns = [
  { key: 'id', label: 'ID' },
  { key: 'name', label: 'Name' },
  { key: 'type', label: 'Type' },
  { key: 'price', label: 'Price' },
  { key: 'stock', label: 'Stock' },
  { key: 'status', label: 'Status' },
  { key: 'actions', label: '' },
];

const rows = computed(() => {
  return (
    data.value?.data.map((product) => ({
      id: product.id,
      name: product.name,
      type: product.type,
      price: product.price,
      stock: product.stock,
      status: product.status,
      actions: product.id,
    })) ?? []
  );
});

// Pagination
const totalPages = computed(() => {
  return data.value?.meta.total_pages ?? 1;
});
</script>

<template>
  <PageContainer>
    <PageHeader title="Products" description="Manage product catalog, pricing and inventory." />

    <SectionCard
      title="Product catalog"
      description="Search and review products available for this tenant."
    >
      <template #actions>
        <Button variant="outline" size="sm" @click="refetch()">
          <RefreshCw class="mr-2 h-4 w-4" />
          Refresh
        </Button>
      </template>

      <ProductsFilters
        v-model:search="search"
        v-model:type="type"
        v-model:min-price="minPrice"
        v-model:max-price="maxPrice"
        :total="data?.meta.total ?? 0"
        @clear="clearFilters"
      />

      <div
        v-if="isLoading"
        class="rounded-xl border border-border bg-muted/30 p-8 text-center text-sm text-muted-foreground"
      >
        Loading products...
      </div>

      <div v-else-if="isError">
        <EmptyState
          title="Unable to load products"
          description="There was a problem connecting to the products API."
          :icon="Package"
        />
      </div>

      <div v-else-if="rows.length === 0">
        <EmptyState
          title="No products found"
          description="Try changing the search filters or create a new product."
          :icon="Package"
        />
      </div>

      <DataTable v-else :columns="columns" :rows="rows">
        <template #cell-price="{ value }"> ${{ Number(value).toFixed(2) }} </template>

        <template #cell-type="{ value }">
          <span
            class="rounded-full border border-border bg-muted px-2 py-1 text-xs font-medium text-muted-foreground"
          >
            {{ value }}
          </span>
        </template>

        <template #cell-status="{ value }">
          <span
            class="rounded-full px-2 py-1 text-xs font-medium"
            :class="
              value === 1
                ? 'bg-emerald-500/10 text-emerald-600 dark:text-emerald-400'
                : 'bg-muted text-muted-foreground'
            "
          >
            {{ value === 1 ? 'Active' : 'Inactive' }}
          </span>
        </template>

        <template #cell-actions="{ row }">
          <div class="flex justify-end">
            <DropdownMenu>
              <DropdownMenuTrigger as-child>
                <Button variant="ghost" size="icon">
                  <MoreHorizontal class="h-4 w-4" />
                </Button>
              </DropdownMenuTrigger>

              <DropdownMenuContent align="end">
                <DropdownMenuItem> View details </DropdownMenuItem>

                <DropdownMenuItem> Edit product </DropdownMenuItem>

                <DropdownMenuItem class="text-destructive"> Delete product </DropdownMenuItem>
              </DropdownMenuContent>
            </DropdownMenu>
          </div>
        </template>
      </DataTable>

      <DataPagination
        v-if="data?.meta"
        class="mt-4"
        :page="page"
        :total-pages="totalPages"
        :total="data.meta.total"
        @previous="page--"
        @next="page++"
      />
    </SectionCard>
  </PageContainer>
</template>
