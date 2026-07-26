<script setup lang="ts">
import { computed, ref, watch } from 'vue';

import {
  ArrowDownUp,
  Building2,
  ChevronDown,
  Grid2X2,
  Loader2,
  RotateCcw,
  Search,
  Store,
} from 'lucide-vue-next';

import { useRoute, useRouter } from 'vue-router';

import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

import { Select, SelectContent, SelectItem, SelectTrigger } from '@/components/ui/select';

import EmptyState from '@/shared/components/erp/EmptyState.vue';
import DataPagination from '@/shared/components/erp/DataPagination.vue';

import { useCategories } from '@/modules/categories/composables/useCategories';

import { useStoreBusinesses } from '../composables/useStoreBusinesses';

import StoreBusinessCard from '../components/businesses/StoreBusinessCard.vue';
import StoreBusinessDrawer from '../components/businesses/StoreBusinessDrawer.vue';

import type { StoreBusiness } from '../types/store-business.types';

const router = useRouter();
const route = useRoute();

function queryString(value: unknown) {
  return typeof value === 'string' ? value : '';
}

const search = ref(queryString(route.query.search));

const debouncedSearch = ref(queryString(route.query.search).trim());

const categoryId = ref(queryString(route.query.category_id) || 'all');

const page = ref(Number(route.query.page) > 0 ? Number(route.query.page) : 1);

const limit = ref(9);

const sortColumn = ref<'name' | 'created_at' | 'total_products'>(
  queryString(route.query.sort) === 'name'
    ? 'name'
    : queryString(route.query.sort) === 'created_at'
      ? 'created_at'
      : 'total_products'
);

const order = ref<'asc' | 'desc'>(queryString(route.query.order) === 'asc' ? 'asc' : 'desc');

let searchTimer: ReturnType<typeof setTimeout> | undefined;

watch(search, (value) => {
  if (searchTimer) {
    clearTimeout(searchTimer);
  }

  searchTimer = setTimeout(() => {
    debouncedSearch.value = value.trim();
  }, 300);
});

watch(
  [debouncedSearch, categoryId, sortColumn, order],
  () => {
    if (page.value !== 1) {
      page.value = 1;
    }
  },
  {
    flush: 'sync',
  }
);

watch([debouncedSearch, categoryId, sortColumn, order, page], () => {
  const query: Record<string, string> = {};

  if (debouncedSearch.value.length >= 2) {
    query.search = debouncedSearch.value;
  }

  if (categoryId.value !== 'all') {
    query.category_id = categoryId.value;
  }

  if (sortColumn.value !== 'total_products') {
    query.sort = sortColumn.value;
  }

  if (order.value !== 'desc') {
    query.order = order.value;
  }

  if (page.value > 1) {
    query.page = String(page.value);
  }

  router.replace({
    name: 'store-businesses',
    query,
  });
});

const params = computed(() => ({
  search: debouncedSearch.value.length >= 2 ? debouncedSearch.value : undefined,

  category_id: categoryId.value !== 'all' ? Number(categoryId.value) : undefined,

  page: page.value,
  limit: limit.value,

  sort_column: sortColumn.value,

  order: order.value,
}));

const { data, isLoading, isFetching, isError } = useStoreBusinesses(params);

const { data: categoriesData, isLoading: categoriesLoading } = useCategories();

const companyCategories = computed(() => {
  return categoriesData.value?.company_categories ?? [];
});

const businesses = computed(() => {
  return data.value?.items ?? [];
});

const totalPages = computed(() => {
  return data.value?.meta.total_pages ?? 1;
});

const isRefreshing = computed(() => {
  return isFetching.value && !isLoading.value;
});

const selectedCategory = computed(() => {
  if (categoryId.value === 'all') {
    return null;
  }

  return (
    companyCategories.value.find((category) => String(category.id) === categoryId.value) ?? null
  );
});

const activeTitle = computed(() => {
  if (debouncedSearch.value) {
    return `Businesses matching “${debouncedSearch.value}”`;
  }

  if (selectedCategory.value) {
    return selectedCategory.value.name;
  }

  return 'Connected businesses';
});

const selectedBusiness = ref<StoreBusiness | null>(null);

const businessDrawerOpen = ref(false);

function openBusiness(business: StoreBusiness) {
  selectedBusiness.value = business;
  businessDrawerOpen.value = true;
}

function openBusinessCatalog(business: StoreBusiness) {
  businessDrawerOpen.value = false;

  router.push({
    name: 'store-catalog',
    query: {
      company_id: String(business.id),
    },
  });
}

function clearFilters() {
  search.value = '';
  debouncedSearch.value = '';

  categoryId.value = 'all';

  sortColumn.value = 'total_products';

  order.value = 'desc';

  page.value = 1;
}
</script>

<template>
  <div class="pb-20">
    <!-- HERO -->
    <section
      class="relative overflow-hidden border-b border-border bg-gradient-to-br from-background via-background to-primary/5"
    >
      <div
        class="pointer-events-none absolute -right-32 -top-48 h-[500px] w-[500px] rounded-full bg-primary/10 blur-3xl"
      />

      <div class="relative mx-auto w-full max-w-[1500px] px-6 py-14 lg:px-8">
        <div class="flex flex-col justify-between gap-8 lg:flex-row lg:items-end">
          <div class="max-w-3xl">
            <p class="text-sm font-semibold text-primary">Nexora businesses</p>

            <h1 class="mt-4 text-4xl font-semibold tracking-[-0.045em] sm:text-5xl">
              Discover the companies behind
              <span class="text-muted-foreground"> every product. </span>
            </h1>

            <p class="mt-4 max-w-2xl text-base leading-7 text-muted-foreground">
              Explore real businesses, understand what they offer and enter their connected Store
              catalog.
            </p>
          </div>

          <div
            class="rounded-2xl border border-border bg-background/80 px-5 py-4 shadow-sm backdrop-blur"
          >
            <p class="text-xs text-muted-foreground">Connected businesses</p>

            <p class="mt-1 text-2xl font-semibold tracking-tight">
              {{ data?.meta.total ?? 0 }}
            </p>

            <p class="text-xs text-muted-foreground">active sellers</p>
          </div>
        </div>
      </div>
    </section>

    <main class="mx-auto w-full max-w-[1500px] space-y-8 px-6 py-8 lg:px-8">
      <!-- BUSINESS FILTERS -->
      <section class="rounded-[1.75rem] border border-border bg-card p-5 shadow-sm">
        <div class="grid gap-3 lg:grid-cols-[minmax(0,1fr)_260px_200px_170px_auto]">
          <div class="relative">
            <Search
              class="pointer-events-none absolute left-4 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
            />

            <Input
              v-model="search"
              class="h-11 rounded-full bg-muted/30 pl-11"
              placeholder="Search businesses..."
            />

            <Loader2
              v-if="isRefreshing"
              class="absolute right-4 top-1/2 h-4 w-4 -translate-y-1/2 animate-spin text-muted-foreground"
            />
          </div>

          <Select v-model="categoryId">
            <SelectTrigger class="h-11 rounded-full">
              <span
                class="truncate"
                :class="categoryId === 'all' ? 'text-muted-foreground' : 'text-foreground'"
              >
                {{
                  categoriesLoading
                    ? 'Loading categories...'
                    : (selectedCategory?.name ?? 'All industries')
                }}
              </span>
            </SelectTrigger>

            <SelectContent>
              <SelectItem value="all"> All industries </SelectItem>

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
            <SelectTrigger class="h-11 rounded-full">
              <span>
                {{
                  sortColumn === 'total_products'
                    ? 'Largest catalog'
                    : sortColumn === 'created_at'
                      ? 'Recently joined'
                      : 'Business name'
                }}
              </span>
            </SelectTrigger>

            <SelectContent>
              <SelectItem value="total_products"> Largest catalog </SelectItem>

              <SelectItem value="created_at"> Recently joined </SelectItem>

              <SelectItem value="name"> Business name </SelectItem>
            </SelectContent>
          </Select>

          <Select v-model="order">
            <SelectTrigger class="h-11 rounded-full">
              <span>
                {{ order === 'desc' ? 'Descending' : 'Ascending' }}
              </span>
            </SelectTrigger>

            <SelectContent>
              <SelectItem value="desc"> Descending </SelectItem>

              <SelectItem value="asc"> Ascending </SelectItem>
            </SelectContent>
          </Select>

          <Button variant="ghost" class="h-11 rounded-full" @click="clearFilters">
            <RotateCcw class="mr-2 h-4 w-4" />
            Reset
          </Button>
        </div>
      </section>

      <!-- RESULTS HEADING -->
      <div class="flex flex-col gap-4 sm:flex-row sm:items-end sm:justify-between">
        <div>
          <p class="text-xs font-semibold uppercase tracking-[0.16em] text-primary">
            Business directory
          </p>

          <h2 class="mt-2 text-2xl font-semibold tracking-tight">
            {{ activeTitle }}
          </h2>

          <p class="mt-2 text-sm text-muted-foreground">
            {{ data?.meta.total ?? 0 }}
            {{ data?.meta.total === 1 ? 'business available' : 'businesses available' }}
          </p>
        </div>

        <div class="flex items-center gap-2 text-xs text-muted-foreground">
          <Store class="h-4 w-4 text-primary" />
          Active Nexora sellers
        </div>
      </div>

      <!-- LOADING -->
      <div v-if="isLoading" class="grid gap-5 md:grid-cols-2 xl:grid-cols-3">
        <div
          v-for="index in 6"
          :key="index"
          class="h-[430px] animate-pulse rounded-[1.75rem] bg-muted"
        />
      </div>

      <!-- ERROR -->
      <EmptyState
        v-else-if="isError"
        title="Unable to load businesses"
        description="There was a problem connecting to the Nexora business directory."
        :icon="Building2"
      />

      <!-- EMPTY -->
      <EmptyState
        v-else-if="businesses.length === 0"
        title="No businesses found"
        description="Try changing the company category or search term."
        :icon="Search"
      />

      <!-- BUSINESS GRID -->
      <template v-else>
        <div class="grid gap-5 md:grid-cols-2 xl:grid-cols-3">
          <StoreBusinessCard
            v-for="business in businesses"
            :key="business.id"
            :business="business"
            @view="openBusiness"
            @catalog="openBusinessCatalog"
          />
        </div>

        <DataPagination
          class="mt-8"
          :page="page"
          :total-pages="totalPages"
          :total="data?.meta.total"
          @previous="page--"
          @next="page++"
        />
      </template>
    </main>

    <StoreBusinessDrawer
      v-model:open="businessDrawerOpen"
      :business="selectedBusiness"
      @catalog="openBusinessCatalog"
    />
  </div>
</template>
