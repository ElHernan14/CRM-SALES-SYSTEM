<script setup lang="ts">
import { computed, ref, watch } from 'vue';

import {
  ChevronDown,
  Grid2X2,
  Loader2,
  PackageSearch,
  RotateCcw,
  Search,
  SlidersHorizontal,
  ArrowRight,
  BadgeCheck,
  Building2,
  ImageIcon,
  Store,
  X,
} from 'lucide-vue-next';

import { useRoute, useRouter } from 'vue-router';

import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

import { Select, SelectContent, SelectItem, SelectTrigger } from '@/components/ui/select';

import EmptyState from '@/shared/components/erp/EmptyState.vue';
import DataPagination from '@/shared/components/erp/DataPagination.vue';

import StoreProductCard from '../components/catalog/StoreProductCard.vue';

import { useStoreCatalog } from '../composables/useStoreCatalog';
import { useAddToStoreCart } from '../composables/useAddToStoreCart';

import { useCategories } from '@/modules/categories/composables/useCategories';
import { useProductTypes } from '@/modules/product-types/composables/useProductTypes';

import type { StoreProduct } from '../types/store-product.types';

import { getCompanyCoverUrl, getCompanyLogoUrl } from '@/shared/utils/assets';

import { useStoreBusiness } from '../composables/useStoreBusiness';

const route = useRoute();
const router = useRouter();

const search = ref(getQueryString(route.query.search));

const debouncedSearch = ref(getQueryString(route.query.search).trim());

const categoryId = ref(getQueryString(route.query.category_id) || 'all');

const typeId = ref(getQueryString(route.query.type_id) || 'all');

const queryKind = getQueryString(route.query.kind);

const kind = ref(queryKind === 'product' || queryKind === 'service' ? queryKind : 'all');

const filtersOpen = ref(false);

const page = ref(Number(route.query.page) > 0 ? Number(route.query.page) : 1);

const minPrice = ref(getQueryString(route.query.min_price));

const maxPrice = ref(getQueryString(route.query.max_price));

const companyId = ref(getQueryString(route.query.company_id));

const querySort = getQueryString(route.query.sort);

const sortColumn = ref<'name' | 'created_at' | 'type' | 'price' | 'stock'>(
  ['name', 'created_at', 'type', 'price', 'stock'].includes(querySort)
    ? (querySort as 'name' | 'created_at' | 'type' | 'price' | 'stock')
    : 'created_at'
);

const queryOrder = getQueryString(route.query.order);

const order = ref<'asc' | 'desc'>(queryOrder === 'asc' ? 'asc' : 'desc');
const limit = ref(12);

function getQueryString(value: unknown) {
  return typeof value === 'string' ? value : '';
}

const initialSearch = getQueryString(route.query.search);

const initialCategoryId = getQueryString(route.query.category_id);

const initialTypeId = getQueryString(route.query.type_id);

const initialKind = getQueryString(route.query.kind);

search.value = initialSearch;

debouncedSearch.value = initialSearch.trim();

categoryId.value = initialCategoryId || 'all';

typeId.value = initialTypeId || 'all';

kind.value = initialKind === 'product' || initialKind === 'service' ? initialKind : 'all';

let searchTimer: ReturnType<typeof setTimeout> | undefined;

const catalogFilterSources = [
  debouncedSearch,
  companyId,
  categoryId,
  typeId,
  kind,
  minPrice,
  maxPrice,
  sortColumn,
  order,
] as const;

watch(search, (value) => {
  if (searchTimer) {
    clearTimeout(searchTimer);
  }

  searchTimer = setTimeout(() => {
    debouncedSearch.value = value.trim();
  }, 300);
});

watch(categoryId, (newValue, oldValue) => {
  if (oldValue && newValue !== oldValue) {
    typeId.value = 'all';
  }
});

watch(
  catalogFilterSources,
  () => {
    if (page.value !== 1) {
      page.value = 1;
    }
  },
  {
    flush: 'sync',
  }
);

watch([...catalogFilterSources, page], () => {
  syncCatalogQuery();
});

function syncCatalogQuery() {
  const query: Record<string, string> = {};

  if (debouncedSearch.value.length >= 2) {
    query.search = debouncedSearch.value;
  }

  if (companyId.value !== '') {
    query.company_id = companyId.value;
  }

  if (categoryId.value !== 'all') {
    query.category_id = categoryId.value;
  }

  if (typeId.value !== 'all') {
    query.type_id = typeId.value;
  }

  if (kind.value !== 'all') {
    query.kind = kind.value;
  }

  if (minPrice.value !== '') {
    query.min_price = minPrice.value;
  }

  if (maxPrice.value !== '') {
    query.max_price = maxPrice.value;
  }

  if (sortColumn.value !== 'created_at') {
    query.sort = sortColumn.value;
  }

  if (order.value !== 'desc') {
    query.order = order.value;
  }

  if (page.value > 1) {
    query.page = String(page.value);
  }

  router.replace({
    name: 'store-catalog',
    query,
  });
}

const selectedBusinessId = computed<number | null>(() => {
  if (companyId.value === '') {
    return null;
  }

  const id = Number(companyId.value);

  return Number.isFinite(id) && id > 0 ? id : null;
});

const {
  data: selectedBusiness,
  isLoading: selectedBusinessLoading,
  isError: selectedBusinessError,
} = useStoreBusiness(selectedBusinessId);

const selectedBusinessLogo = computed(() => {
  return getCompanyLogoUrl(selectedBusiness.value?.logo);
});

const selectedBusinessCover = computed(() => {
  return getCompanyCoverUrl(selectedBusiness.value?.cover_image);
});

function clearBusinessFilter() {
  companyId.value = '';
  page.value = 1;
}

const featuredCategories = computed(() => {
  return productCategories.value.slice(0, 7);
});

function selectCategory(value: string) {
  categoryId.value = value;
  typeId.value = 'all';
}

const { addToCart, addingProductId, isAdding } = useAddToStoreCart();

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

const hasInvalidPriceRange = computed(() => {
  if (minPrice.value === '' || maxPrice.value === '') {
    return false;
  }

  return Number(minPrice.value) > Number(maxPrice.value);
});

const productParams = computed(() => ({
  page: page.value,
  limit: limit.value,

  search: debouncedSearch.value.length >= 2 ? debouncedSearch.value : undefined,

  kind: kind.value !== 'all' ? (kind.value as 'product' | 'service') : undefined,

  category_id: categoryId.value !== 'all' ? Number(categoryId.value) : undefined,

  company_id: selectedBusinessId.value ?? undefined,

  type_id: typeId.value !== 'all' ? Number(typeId.value) : undefined,

  min_price:
    !hasInvalidPriceRange.value && minPrice.value !== '' ? Number(minPrice.value) : undefined,

  max_price:
    !hasInvalidPriceRange.value && maxPrice.value !== '' ? Number(maxPrice.value) : undefined,

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

  companyId.value = '';
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
  return addToCart(product, 1, true);
}

const activeDiscoveryLabel = computed(() => {
  if (debouncedSearch.value) {
    return `Results for “${debouncedSearch.value}”`;
  }

  if (selectedProductType.value) {
    return selectedProductType.value.name;
  }

  if (selectedCategory.value) {
    return selectedCategory.value.name;
  }

  if (kind.value === 'product') {
    return 'Products';
  }

  if (kind.value === 'service') {
    return 'Services';
  }

  return 'All available resources';
});

// Validate number input for minPrice and maxPrice
function validateNumberInput(event: KeyboardEvent) {
  const allowedKeys = ['Backspace', 'Delete', 'ArrowLeft', 'ArrowRight', 'Tab', 'Enter'];

  // Permitir números y punto decimal
  const isNumber = /^[0-9]$/.test(event.key);
  const isDot = event.key === '.';

  if (!isNumber && !isDot && !allowedKeys.includes(event.key)) {
    event.preventDefault();
  }
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

      <div class="relative mx-auto w-full max-w-[1500px] px-6 py-14 lg:px-8">
        <div class="flex flex-col justify-between gap-8 lg:flex-row lg:items-end">
          <div class="max-w-3xl">
            <p class="text-sm font-semibold text-primary">Nexora catalog</p>

            <h1 class="mt-4 text-4xl font-semibold tracking-[-0.045em] text-foreground sm:text-5xl">
              Products and services from
              <span class="text-muted-foreground"> connected businesses. </span>
            </h1>

            <p class="mt-4 max-w-2xl text-base leading-7 text-muted-foreground">
              Search real inventory, explore business categories and buy from companies operating
              throughout Nexora.
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

    <div class="mx-auto w-full max-w-[1500px] space-y-8 px-6 py-8 lg:px-8">
      <!-- ACTIVE BUSINESS CONTEXT -->
      <section
        v-if="selectedBusinessId"
        class="relative overflow-hidden rounded-[1.75rem] border border-border bg-card shadow-sm"
      >
        <div v-if="selectedBusinessLoading" class="flex min-h-48 items-center justify-center">
          <div class="text-center">
            <Loader2 class="mx-auto h-6 w-6 animate-spin text-primary" />

            <p class="mt-3 text-sm text-muted-foreground">Loading business catalog...</p>
          </div>
        </div>

        <div
          v-else-if="selectedBusinessError"
          class="flex flex-col items-center justify-center px-6 py-10 text-center"
        >
          <Building2 class="h-7 w-7 text-muted-foreground" />

          <p class="mt-4 text-sm font-semibold">Business information unavailable</p>

          <p class="mt-2 text-sm text-muted-foreground">
            Products are still filtered by the selected seller.
          </p>

          <Button variant="ghost" size="sm" class="mt-4 rounded-full" @click="clearBusinessFilter">
            <X class="mr-2 h-4 w-4" />
            Browse all businesses
          </Button>
        </div>

        <template v-else-if="selectedBusiness">
          <img
            v-if="selectedBusinessCover"
            :src="selectedBusinessCover"
            :alt="`${selectedBusiness.name} cover`"
            class="absolute inset-0 h-full w-full object-cover"
          />

          <div
            v-else
            class="absolute inset-0 bg-gradient-to-br from-primary/15 via-muted/40 to-background"
          />

          <div
            class="absolute inset-0 bg-gradient-to-r from-background/95 via-background/85 to-background/45"
          />

          <div
            class="relative flex min-h-52 flex-col justify-between gap-7 p-6 md:flex-row md:items-end"
          >
            <div class="flex min-w-0 items-start gap-4">
              <div
                class="flex h-20 w-20 shrink-0 items-center justify-center overflow-hidden rounded-[1.5rem] border border-border bg-background shadow-xl"
              >
                <img
                  v-if="selectedBusinessLogo"
                  :src="selectedBusinessLogo"
                  :alt="`${selectedBusiness.name} logo`"
                  class="h-full w-full object-contain p-2"
                />

                <Building2 v-else class="h-7 w-7 text-primary" />
              </div>

              <div class="min-w-0">
                <div class="flex flex-wrap items-center gap-2">
                  <span
                    class="rounded-full border border-primary/20 bg-primary/10 px-3 py-1 text-xs font-semibold text-primary"
                  >
                    {{ selectedBusiness.category }}
                  </span>

                  <span
                    class="inline-flex items-center gap-1.5 rounded-full border border-border bg-background/70 px-3 py-1 text-xs font-medium text-muted-foreground backdrop-blur"
                  >
                    <BadgeCheck class="h-3.5 w-3.5 text-primary" />
                    Connected business
                  </span>
                </div>

                <p class="mt-4 text-xs font-semibold uppercase tracking-[0.16em] text-primary">
                  Business catalog
                </p>

                <h2 class="mt-2 truncate text-3xl font-semibold tracking-[-0.04em]">
                  {{ selectedBusiness.name }}
                </h2>

                <p class="mt-3 line-clamp-2 max-w-2xl text-sm leading-6 text-muted-foreground">
                  {{
                    selectedBusiness.description ||
                    'Explore products and services published by this Nexora business.'
                  }}
                </p>
              </div>
            </div>

            <div class="flex shrink-0 flex-col gap-3 sm:flex-row md:flex-col">
              <div
                class="rounded-2xl border border-border bg-background/75 px-4 py-3 shadow-sm backdrop-blur"
              >
                <p class="text-xs text-muted-foreground">Available catalog</p>

                <p class="mt-1 text-xl font-semibold">
                  {{ selectedBusiness.total_products }}
                  resources
                </p>
              </div>

              <Button
                variant="outline"
                class="rounded-full bg-background/80 backdrop-blur"
                @click="clearBusinessFilter"
              >
                <X class="mr-2 h-4 w-4" />
                View all sellers
              </Button>
            </div>
          </div>
        </template>
      </section>

      <!-- CATEGORY DISCOVERY -->
      <section v-if="categoriesLoading || featuredCategories.length > 0" class="space-y-4">
        <div class="flex items-center justify-between gap-4">
          <div>
            <p class="text-xs font-semibold uppercase tracking-[0.16em] text-primary">Categories</p>

            <h2 class="mt-2 text-xl font-semibold tracking-tight">Explore by what you need</h2>
          </div>

          <Button
            v-if="categoryId !== 'all'"
            variant="ghost"
            size="sm"
            class="rounded-full"
            @click="selectCategory('all')"
          >
            <X class="mr-2 h-4 w-4" />
            Clear category
          </Button>
        </div>

        <div v-if="categoriesLoading" class="flex gap-3 overflow-hidden">
          <div
            v-for="index in 6"
            :key="index"
            class="h-12 w-40 shrink-0 animate-pulse rounded-full bg-muted"
          />
        </div>

        <div v-else class="flex gap-3 overflow-x-auto pb-2">
          <button
            type="button"
            class="inline-flex shrink-0 items-center gap-2 rounded-full border px-4 py-2.5 text-sm font-medium transition"
            :class="
              categoryId === 'all'
                ? 'border-primary bg-primary text-primary-foreground shadow-sm'
                : 'border-border bg-card text-muted-foreground hover:border-primary/30 hover:text-foreground'
            "
            @click="selectCategory('all')"
          >
            <Grid2X2 class="h-4 w-4" />
            All products
          </button>

          <button
            v-for="category in featuredCategories"
            :key="category.id"
            type="button"
            class="inline-flex shrink-0 items-center gap-2 rounded-full border px-4 py-2.5 text-sm font-medium transition"
            :class="
              categoryId === String(category.id)
                ? 'border-primary bg-primary text-primary-foreground shadow-sm'
                : 'border-border bg-card text-muted-foreground hover:border-primary/30 hover:text-foreground'
            "
            @click="selectCategory(String(category.id))"
          >
            {{ category.name }}
          </button>
        </div>
      </section>
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

                <Input
                  v-model="minPrice"
                  inputmode="decimal"
                  placeholder="0.00"
                  @keydown="validateNumberInput"
                />
                <p
                  v-if="hasInvalidPriceRange"
                  class="md:col-span-2 xl:col-span-6 text-xs text-destructive"
                >
                  Minimum price cannot be greater than maximum price.
                </p>
              </div>

              <div class="space-y-2">
                <label class="text-xs font-medium text-muted-foreground"> Maximum price </label>

                <Input
                  v-model="maxPrice"
                  inputmode="decimal"
                  placeholder="0.00"
                  @keydown="validateNumberInput"
                />
                <p
                  v-if="hasInvalidPriceRange"
                  class="md:col-span-2 xl:col-span-6 text-xs text-destructive"
                >
                  Minimum price cannot be greater than maximum price.
                </p>
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

            <div class="flex justify-end border-t border-border px-5 py-4">
              <Button variant="ghost" size="sm" @click="clearFilters">
                <RotateCcw class="mr-2 h-4 w-4" />
                Reset filters
              </Button>
            </div>
          </div>
        </Transition>
      </section>

      <div class="flex flex-col gap-4 sm:flex-row sm:items-end sm:justify-between">
        <div>
          <p class="text-xs font-semibold uppercase tracking-[0.16em] text-primary">
            Catalog results
          </p>

          <h2 class="mt-2 text-2xl font-semibold tracking-tight text-foreground">
            {{ activeDiscoveryLabel }}
          </h2>

          <p class="mt-2 text-sm text-muted-foreground">
            {{ data?.meta.total ?? 0 }}
            {{ data?.meta.total === 1 ? 'result available' : 'results available' }}
          </p>
        </div>

        <div
          v-if="isRefreshing"
          class="inline-flex items-center gap-2 self-start rounded-full border border-border bg-card px-3 py-1.5 text-xs text-muted-foreground sm:self-auto"
        >
          <Loader2 class="h-3.5 w-3.5 animate-spin" />
          Updating catalog
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
            :adding="addingProductId === product.id"
            :actions-disabled="isAdding"
            @view="openProduct"
            @add="addProduct"
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
