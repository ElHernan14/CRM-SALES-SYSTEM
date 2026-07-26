<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import { Building2, Search, SlidersHorizontal, ArrowRight } from 'lucide-vue-next';

import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

import PageContainer from '@/shared/components/erp/PageContainer.vue';
import PageHeader from '@/shared/components/erp/PageHeader.vue';
import EmptyState from '@/shared/components/erp/EmptyState.vue';
import DataPagination from '@/shared/components/erp/DataPagination.vue';

import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';

import { getCompanyLogoUrl } from '@/shared/utils/assets';
import { useSuppliers } from '../composables/useSuppliers';
import { useCategories } from '@/modules/categories/composables/useCategories';

import { storeToRefs } from 'pinia';
import { useAuthStore } from '@/modules/auth/stores/auth.store';

const auth = useAuthStore();
const { user } = storeToRefs(auth);

const categoryId = ref('all');
const search = ref('');
const page = ref(1);
const limit = ref(9);

const sortColumn = ref<'name' | 'total_products' | 'created_at'>('total_products');
const order = ref<'asc' | 'desc'>('desc');

watch([search, categoryId, sortColumn, order], () => {
  page.value = 1;
});

const supplierParams = computed(() => ({
  search: search.value || undefined,

  category_id: categoryId.value !== 'all' ? Number(categoryId.value) : undefined,

  page: page.value,
  limit: limit.value,

  sort_column: sortColumn.value,
  order: order.value,
}));

const { data, isLoading, isError } = useSuppliers(supplierParams);

const suppliers = computed(() => {
  return data.value?.items.filter((supplier) => supplier.id !== user.value?.company_id) ?? [];
});
const totalPages = computed(() => data.value?.meta.total_pages ?? 1);

const { data: categoriesData, isLoading: categoriesLoading } = useCategories();

const companyCategories = computed(() => {
  return categoriesData.value?.company_categories ?? [];
});

function getSupplierLogo(path?: string | null) {
  return getCompanyLogoUrl(path);
}

const selectedCompanyCategory = computed(() => {
  if (categoryId.value === 'all') return null;

  return (
    companyCategories.value.find((category) => String(category.id) === categoryId.value) ?? null
  );
});

const companyCategoryTriggerLabel = computed(() => {
  if (categoriesLoading.value) {
    return 'Loading categories...';
  }

  if (categoryId.value === 'all') {
    return 'All categories';
  }

  return selectedCompanyCategory.value?.name ?? 'Company category';
});
</script>

<template>
  <PageContainer>
    <PageHeader
      title="B2B Marketplace"
      description="Discover suppliers, browse their catalogs and create purchase flows for your company."
    />

    <section class="rounded-2xl border border-border bg-card p-4 shadow-sm sm:p-6">
      <div class="flex flex-col gap-4 lg:flex-row lg:items-center lg:justify-between">
        <div>
          <h2 class="text-lg font-semibold tracking-tight text-foreground">Verified suppliers</h2>

          <p class="mt-1 text-sm text-muted-foreground">
            Explore companies available in the multi-tenant commerce network.
          </p>
        </div>

        <div
          class="flex items-center gap-2 rounded-full border border-border bg-muted/40 px-3 py-1 text-xs text-muted-foreground"
        >
          <Building2 class="h-3.5 w-3.5" />
          {{ data?.meta.total ?? 0 }} suppliers
        </div>
      </div>

      <div class="mt-6 grid gap-3 sm:grid-cols-2 xl:grid-cols-[minmax(0,1fr)_220px_220px_160px]">
        <div class="relative">
          <Search
            class="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
          />

          <Input v-model="search" class="pl-9" placeholder="Search suppliers..." />
        </div>

        <Select v-model="categoryId">
          <SelectTrigger>
            <span
              class="truncate"
              :class="categoryId === 'all' ? 'text-muted-foreground' : 'text-foreground'"
            >
              {{ companyCategoryTriggerLabel }}
            </span>
          </SelectTrigger>

          <SelectContent>
            <SelectItem value="all"> All categories </SelectItem>

            <SelectItem
              v-for="category in companyCategories"
              :key="category.id"
              :value="String(category.id)"
            >
              <div class="flex flex-col py-0.5">
                <span class="text-sm font-medium">
                  {{ category.name }}
                </span>

                <span v-if="category.description" class="text-xs text-muted-foreground">
                  {{ category.description }}
                </span>
              </div>
            </SelectItem>
          </SelectContent>
        </Select>

        <Select v-model="sortColumn">
          <SelectTrigger>
            <SelectValue placeholder="Sort by" />
          </SelectTrigger>

          <SelectContent>
            <SelectItem value="total_products"> Total products </SelectItem>

            <SelectItem value="name"> Name </SelectItem>

            <SelectItem value="created_at"> Created date </SelectItem>
          </SelectContent>
        </Select>

        <Select v-model="order">
          <SelectTrigger>
            <SelectValue placeholder="Order" />
          </SelectTrigger>

          <SelectContent>
            <SelectItem value="desc"> Descending </SelectItem>

            <SelectItem value="asc"> Ascending </SelectItem>
          </SelectContent>
        </Select>
      </div>
    </section>

    <div
      v-if="isLoading"
      class="rounded-xl border border-border bg-muted/30 p-8 text-center text-sm text-muted-foreground"
    >
      Loading suppliers...
    </div>

    <EmptyState
      v-else-if="isError"
      title="Unable to load suppliers"
      description="There was a problem connecting to the marketplace API."
      :icon="Building2"
    />

    <EmptyState
      v-else-if="suppliers.length === 0"
      title="No suppliers found"
      description="Try changing your search criteria."
      :icon="SlidersHorizontal"
    />

    <template v-else>
      <div class="grid gap-4 sm:grid-cols-2 xl:grid-cols-3">
        <article
          v-for="supplier in suppliers"
          :key="supplier.id"
          class="group flex flex-col overflow-hidden rounded-2xl border border-border bg-card shadow-sm transition hover:-translate-y-0.5 hover:shadow-md"
        >
          <div class="relative h-40 overflow-hidden border-b border-border bg-muted/40">
            <img
              v-if="getSupplierLogo(supplier.logo)"
              :src="getSupplierLogo(supplier.logo)!"
              :alt="supplier.name"
              class="h-full w-full object-cover transition duration-300 group-hover:scale-[1.03]"
            />

            <div v-else class="flex h-full w-full items-center justify-center">
              <Building2 class="h-10 w-10 text-muted-foreground" />
            </div>

            <div
              class="absolute inset-0 bg-linear-to-t from-background/75 via-transparent to-transparent"
            />

            <span
              class="absolute right-4 top-4 rounded-full border border-primary/20 bg-background/90 px-3 py-1 text-xs font-semibold text-primary shadow-sm"
            >
              {{ supplier.category }}
            </span>
          </div>

          <div class="flex flex-1 flex-col p-5">
            <div class="flex-1">
              <div class="flex items-start justify-between gap-3">
                <div class="min-w-0">
                  <h3 class="truncate text-lg font-semibold tracking-tight text-foreground">
                    {{ supplier.name }}
                  </h3>

                  <p v-if="supplier.email" class="mt-1 truncate text-xs text-muted-foreground">
                    {{ supplier.email }}
                  </p>
                </div>

                <span
                  class="rounded-full bg-muted px-2 py-1 text-[11px] font-medium text-muted-foreground"
                >
                  B2B
                </span>
              </div>

              <p class="mt-4 line-clamp-3 min-h-15 text-sm leading-6 text-muted-foreground">
                {{ supplier.description ?? 'No supplier description provided yet.' }}
              </p>
            </div>

            <div
              class="mt-5 flex flex-col gap-4 border-t border-border pt-4 min-[390px]:flex-row min-[390px]:items-center min-[390px]:justify-between"
            >
              <div>
                <p class="text-xs text-muted-foreground">Available products</p>

                <p class="mt-1 text-lg font-semibold text-foreground">
                  {{ supplier.total_products }}
                </p>
              </div>

              <Button
                size="sm"
                class="w-full rounded-full px-4 min-[390px]:w-auto"
                @click="$router.push(`/erp/marketplace/suppliers/${supplier.id}`)"
              >
                Browse catalog
                <ArrowRight class="ml-2 h-4 w-4" />
              </Button>
            </div>
          </div>
        </article>
      </div>

      <DataPagination
        :page="page"
        :total-pages="totalPages"
        :total="data?.meta.total"
        @previous="page--"
        @next="page++"
      />
    </template>
  </PageContainer>
</template>
