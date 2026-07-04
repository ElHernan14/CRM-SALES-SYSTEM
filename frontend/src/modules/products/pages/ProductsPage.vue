<script setup lang="ts">
import { computed, ref } from 'vue';
import { Package, RefreshCw } from 'lucide-vue-next';

import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

import PageContainer from '@/shared/components/erp/PageContainer.vue';
import PageHeader from '@/shared/components/erp/PageHeader.vue';
import SectionCard from '@/shared/components/erp/SectionCard.vue';
import EmptyState from '@/shared/components/erp/EmptyState.vue';
import DataTable from '@/shared/components/erp/DataTable.vue';

import { storeToRefs } from 'pinia';
import { useAuthStore } from '@/modules/auth/stores/auth.store';

import { useProducts } from '../composables/useProducts';

const auth = useAuthStore();
const { user } = storeToRefs(auth);

const search = ref('');
const page = ref(1);
const limit = ref(10);

const productParams = computed(() => ({
  search: search.value || undefined,
  company_id: user.value?.company_id,
  page: page.value,
  limit: limit.value,
}));

const { data, isLoading, isError, refetch } = useProducts(productParams);

const columns = [
  { key: 'id', label: 'ID' },
  { key: 'name', label: 'Name' },
  { key: 'type', label: 'Type' },
  { key: 'price', label: 'Price' },
  { key: 'stock', label: 'Stock' },
  { key: 'status', label: 'Status' },
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
    })) ?? []
  );
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

      <div class="mb-5 flex flex-col gap-3 md:flex-row md:items-center md:justify-between">
        <Input v-model="search" class="max-w-sm" placeholder="Search products..." />

        <p class="text-sm text-muted-foreground">Total: {{ data?.meta.total ?? 0 }}</p>
      </div>

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
      </DataTable>
    </SectionCard>
  </PageContainer>
</template>
