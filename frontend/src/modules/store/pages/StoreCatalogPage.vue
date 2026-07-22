<script setup lang="ts">
import { computed, ref, watch } from 'vue';

import { useRouter } from 'vue-router';

import {
  ArrowDownUp,
  ChevronDown,
  Grid2X2,
  Loader2,
  RotateCcw,
  Search,
  SlidersHorizontal,
  Sparkles,
} from 'lucide-vue-next';

import { toast } from 'vue-sonner';

import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

import { Select, SelectContent, SelectItem, SelectTrigger } from '@/components/ui/select';

import EmptyState from '@/shared/components/erp/EmptyState.vue';
import DataPagination from '@/shared/components/erp/DataPagination.vue';

import StoreProductCard from '../components/catalog/StoreProductCard.vue';

import { useStoreCatalog } from '../composables/useStoreCatalog';

import { useCategories } from '@/modules/categories/composables/useCategories';
import { useProductTypes } from '@/modules/product-types/composables/useProductTypes';

import type { StoreProduct } from '../types/store-product.types';

const router = useRouter();

const search = ref('');
const debouncedSearch = ref('');

const categoryId = ref('all');
const typeId = ref('all');
const kind = ref('all');

const minPrice = ref('');
const maxPrice = ref('');

const sortColumn = ref<'name' | 'created_at' | 'type' | 'price' | 'stock'>('created_at');

const order = ref<'asc' | 'desc'>('desc');

const filtersOpen = ref(false);

const page = ref(1);
const limit = ref(12);

let searchTimer: ReturnType<typeof setTimeout> | undefined;

watch(search, (value) => {
  if (searchTimer) {
    clearTimeout(searchTimer);
  }

  searchTimer = setTimeout(() => {
    debouncedSearch.value = value.trim();
  }, 300);
});

watch(categoryId, () => {
  typeId.value = 'all';
});

watch([debouncedSearch, categoryId, typeId, kind, minPrice, maxPrice, sortColumn, order], () => {
  page.value = 1;
});

const selectedCategoryId = computed<number | null>(() => {
  return categoryId.value !== 'all' ? Number(categoryId.value) : null;
});

const { data: categoriesData, isLoading: categoriesLoading } = useCategories();

const productCategories = computed(() => {
  return categoriesData.value?.product_categories ?? [];
});

const { data: productTypesData, isLoading: productTypesLoading } =
  useProductTypes(selectedCategoryId);

const productTypes = computed(() => {
  return productTypesData.value?.items ?? [];
});

const selectedCategory = computed(() => {
  if (categoryId.value === 'all') return null;

  return (
    productCategories.value.find((category) => String(category.id) === categoryId.value) ?? null
  );
});

const selectedProductType = computed(() => {
  if (typeId.value === 'all') return null;

  return productTypes.value.find((type) => String(type.id) === typeId.value) ?? null;
});

const categoryLabel = computed(() => {
  if (categoriesLoading.value) {
    return 'Loading categories...';
  }

  return selectedCategory.value?.name ?? 'All categories';
});

const typeLabel = computed(() => {
  if (categoryId.value === 'all') {
    return 'Select category first';
  }

  if (productTypesLoading.value) {
    return 'Loading types...';
  }

  return selectedProductType.value?.name ?? 'All product types';
});

const productParams = computed(() => ({
  page: page.value,
  limit: limit.value,

  search: debouncedSearch.value.length >= 2 ? debouncedSearch.value : undefined,

  kind: kind.value !== 'all' ? (kind.value as 'product' | 'service') : undefined,

  category_id: categoryId.value !== 'all' ? Number(categoryId.value) : undefined,

  type_id: typeId.value !== 'all' ? Number(typeId.value) : undefined,

  min_price: minPrice.value !== '' ? Number(minPrice.value) : undefined,

  max_price: maxPrice.value !== '' ? Number(maxPrice.value) : undefined,

  sort_column: sortColumn.value,
  order: order.value,
}));

const { data, isLoading, isFetching, isError } = useStoreCatalog(productParams);

const products = computed(() => data.value?.items ?? []);

const totalPages = computed(() => data.value?.meta.total_pages ?? 1);

const isRefreshing = computed(() => isFetching.value && !isLoading.value);

const advancedFiltersCount = computed(() => {
  return [
    kind.value !== 'all' ? kind.value : '',
    typeId.value !== 'all' ? typeId.value : '',
    minPrice.value,
    maxPrice.value,
    sortColumn.value !== 'created_at' ? sortColumn.value : '',
    order.value !== 'desc' ? order.value : '',
  ].filter(Boolean).length;
});

function clearFilters() {
  search.value = '';
  debouncedSearch.value = '';

  categoryId.value = 'all';
  typeId.value = 'all';
  kind.value = 'all';

  minPrice.value = '';
  maxPrice.value = '';

  sortColumn.value = 'created_at';
  order.value = 'desc';

  page.value = 1;
}

function openProduct(product: StoreProduct) {
  router.push({
    name: 'store-product',
    params: {
      productId: product.id,
    },
  });
}

function addProduct(product: StoreProduct) {
  toast.success('Cart foundation ready', {
    description: `${product.name} will be connected to the multi-seller cart next.`,
  });
}

function favoriteProduct(product: StoreProduct) {
  toast.success('Saved for later', {
    description: product.name,
  });
}
</script>

<template>
  <div class="pb-16">
    <!-- CATALOG HERO -->
    <section
      class="relative overflow-hidden border-b border-border bg-gradient-to-br from-background via-background to-primary/5"
    >
      <div
        class="pointer-events-none absolute -right-32 -top-48 h-[500px] w-[500px] rounded-full bg-primary/10 blur-3xl"
      />

      <div class="relative mx-auto max-w-[1500px] px-6 py-14 lg:px-8">
        <div class="flex flex-col justify-between gap-8 lg:flex-row lg:items-end">
          <div class="max-w-3xl">
            <div
              class="inline-flex items-center gap-2 rounded-full border border-primary/20 bg-primary/10 px-4 py-2 text-xs font-semibold text-primary"
            >
              <Sparkles class="h-4 w-4" />
              Discover products from trusted stores.
            </div>

            <h1 class="mt-5 text-4xl font-semibold tracking-[-0.035em] text-foreground sm:text-5xl">
              Discover products built for
              <span class="text-muted-foreground"> everyday life. </span>
            </h1>

            <p class="mt-4 max-w-2xl text-base leading-7 text-muted-foreground">
              Browse verified stores, compare categories and find products backed by real inventory.
            </p>
          </div>

          <div
            class="rounded-2xl border border-border bg-background/80 px-5 py-4 shadow-sm backdrop-blur"
          >
            <p class="text-xs text-muted-foreground">Available catalog</p>

            <p class="mt-1 text-2xl font-semibold tracking-tight text-foreground">
              {{ data?.meta.total ?? 0 }}
            </p>

            <p class="text-xs text-muted-foreground">products and services</p>
          </div>
        </div>
      </div>
    </section>

    <div class="mx-auto max-w-[1500px] space-y-8 px-6 py-8 lg:px-8">
      <!-- PRIMARY FILTERS -->
      <section class="rounded-[1.75rem] border border-border bg-card p-5 shadow-sm">
        <div class="grid gap-3 lg:grid-cols-[minmax(0,1fr)_240px_auto]">
          <div class="relative">
            <Search
              class="pointer-events-none absolute left-4 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
            />

            <Input
              v-model="search"
              class="h-11 rounded-full bg-muted/30 pl-11"
              placeholder="Search products, services or stores..."
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
                {{ categoryLabel }}
              </span>
            </SelectTrigger>

            <SelectContent>
              <SelectItem value="all"> All categories </SelectItem>

              <SelectItem
                v-for="category in productCategories"
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

          <Button
            variant="outline"
            class="h-11 justify-between rounded-full px-5"
            @click="filtersOpen = !filtersOpen"
          >
            <span class="flex items-center">
              <SlidersHorizontal class="mr-2 h-4 w-4" />
              Filters

              <span
                v-if="advancedFiltersCount > 0"
                class="ml-2 rounded-full bg-primary px-2 py-0.5 text-[10px] font-semibold text-primary-foreground"
              >
                {{ advancedFiltersCount }}
              </span>
            </span>

            <ChevronDown
              class="ml-3 h-4 w-4 transition-transform"
              :class="{
                'rotate-180': filtersOpen,
              }"
            />
          </Button>
        </div>

        <Transition
          enter-active-class="transition-all duration-250 ease-out"
          enter-from-class="max-h-0 opacity-0"
          enter-to-class="max-h-[500px] opacity-100"
          leave-active-class="transition-all duration-200 ease-in"
          leave-from-class="max-h-[500px] opacity-100"
          leave-to-class="max-h-0 opacity-0"
        >
          <div
            v-if="filtersOpen"
            class="mt-5 overflow-hidden rounded-2xl border border-border bg-muted/15"
          >
            <div class="grid gap-4 p-5 md:grid-cols-2 xl:grid-cols-6">
              <div class="space-y-2">
                <label class="text-xs font-medium text-muted-foreground"> Resource kind </label>

                <Select v-model="kind">
                  <SelectTrigger>
                    <span>
                      {{
                        kind === 'all' ? 'All kinds' : kind === 'product' ? 'Products' : 'Services'
                      }}
                    </span>
                  </SelectTrigger>

                  <SelectContent>
                    <SelectItem value="all"> All kinds </SelectItem>

                    <SelectItem value="product"> Products </SelectItem>

                    <SelectItem value="service"> Services </SelectItem>
                  </SelectContent>
                </Select>
              </div>

              <div class="space-y-2">
                <label class="text-xs font-medium text-muted-foreground"> Product type </label>

                <Select v-model="typeId" :disabled="categoryId === 'all' || productTypesLoading">
                  <SelectTrigger>
                    <span
                      class="truncate"
                      :class="typeId === 'all' ? 'text-muted-foreground' : 'text-foreground'"
                    >
                      {{ typeLabel }}
                    </span>
                  </SelectTrigger>

                  <SelectContent>
                    <SelectItem value="all"> All product types </SelectItem>

                    <SelectItem
                      v-for="productType in productTypes"
                      :key="productType.id"
                      :value="String(productType.id)"
                    >
                      <div class="flex flex-col py-0.5">
                        <span class="text-sm font-medium">
                          {{ productType.name }}
                        </span>

                        <span v-if="productType.description" class="text-xs text-muted-foreground">
                          {{ productType.description }}
                        </span>
                      </div>
                    </SelectItem>
                  </SelectContent>
                </Select>
              </div>

              <div class="space-y-2">
                <label class="text-xs font-medium text-muted-foreground"> Minimum price </label>

                <Input v-model="minPrice" inputmode="decimal" placeholder="0.00" />
              </div>

              <div class="space-y-2">
                <label class="text-xs font-medium text-muted-foreground"> Maximum price </label>

                <Input v-model="maxPrice" inputmode="decimal" placeholder="0.00" />
              </div>

              <div class="space-y-2">
                <label class="text-xs font-medium text-muted-foreground"> Sort by </label>

                <Select v-model="sortColumn">
                  <SelectTrigger>
                    <span>
                      {{
                        sortColumn === 'created_at'
                          ? 'Newest'
                          : sortColumn === 'price'
                            ? 'Price'
                            : sortColumn === 'stock'
                              ? 'Availability'
                              : 'Name'
                      }}
                    </span>
                  </SelectTrigger>

                  <SelectContent>
                    <SelectItem value="created_at"> Newest </SelectItem>

                    <SelectItem value="name"> Name </SelectItem>

                    <SelectItem value="price"> Price </SelectItem>

                    <SelectItem value="stock"> Availability </SelectItem>
                  </SelectContent>
                </Select>
              </div>

              <div class="space-y-2">
                <label class="text-xs font-medium text-muted-foreground"> Order </label>

                <Select v-model="order">
                  <SelectTrigger>
                    <span>
                      {{ order === 'desc' ? 'Descending' : 'Ascending' }}
                    </span>
                  </SelectTrigger>

                  <SelectContent>
                    <SelectItem value="desc"> Descending </SelectItem>

                    <SelectItem value="asc"> Ascending </SelectItem>
                  </SelectContent>
                </Select>
              </div>
            </div>

            <div class="flex items-center justify-between border-t border-border px-5 py-4">
              <div class="flex items-center gap-2 text-xs text-muted-foreground">
                <ArrowDownUp class="h-3.5 w-3.5" />
                Refine your discovery experience
              </div>

              <Button variant="ghost" size="sm" @click="clearFilters">
                <RotateCcw class="mr-2 h-4 w-4" />
                Reset filters
              </Button>
            </div>
          </div>
        </Transition>
      </section>

      <div class="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
        <div>
          <h2 class="text-xl font-semibold tracking-tight text-foreground">Explore the catalog</h2>

          <p class="mt-1 text-sm text-muted-foreground">
            {{ data?.meta.total ?? 0 }}
            available results
          </p>
        </div>
      </div>

      <div v-if="isLoading" class="grid gap-5 md:grid-cols-2 xl:grid-cols-3">
        <div
          v-for="index in 6"
          :key="index"
          class="animate-pulse overflow-hidden rounded-[1.75rem] border border-border bg-card"
        >
          <div class="h-64 bg-muted" />

          <div class="space-y-4 p-5">
            <div class="h-5 w-3/4 rounded bg-muted" />
            <div class="h-4 w-full rounded bg-muted" />
            <div class="h-10 rounded bg-muted" />
          </div>
        </div>
      </div>

      <EmptyState
        v-else-if="isError"
        title="Unable to load the catalog"
        description="There was a problem connecting to the commerce network."
        :icon="Grid2X2"
      />

      <EmptyState
        v-else-if="products.length === 0"
        title="No products found"
        description="Try adjusting your search or filters."
        :icon="Search"
      />

      <template v-else>
        <div class="grid gap-5 md:grid-cols-2 xl:grid-cols-3">
          <StoreProductCard
            v-for="product in products"
            :key="product.id"
            :product="product"
            @view="openProduct"
            @add="addProduct"
            @favorite="favoriteProduct"
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
    </div>
  </div>
</template>
