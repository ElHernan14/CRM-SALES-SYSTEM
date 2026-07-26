<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import {
  ArrowLeft,
  Building2,
  PackageSearch,
  Search,
  ShoppingCart,
  Filter,
  RotateCcw,
  SlidersHorizontal,
  ChevronDown,
  Loader2,
} from 'lucide-vue-next';
import { useQueryClient } from '@tanstack/vue-query';
import { toast } from 'vue-sonner';

import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

import {
  Select,
  SelectContent,
  SelectItem,
  SelectItemText,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';

import PageContainer from '@/shared/components/erp/PageContainer.vue';
import EmptyState from '@/shared/components/erp/EmptyState.vue';
import DataPagination from '@/shared/components/erp/DataPagination.vue';
import MarketplaceProductDrawer from '../components/MarketplaceProductDrawer.vue';
import PurchaseCartDrawer from '../components/PurchaseCartDrawer.vue';

import { useStoreProducts } from '../composables/useStoreProducts';
import { useCategories } from '@/modules/categories/composables/useCategories';
import { useEnsurePurchaseCart } from '../composables/useEnsurePurchaseCart';
import { useAddPurchaseCartItem } from '../composables/useAddPurchaseCartItem';

import type { StoreProduct } from '../types/store-product.types';
import type { Supplier } from '../types/supplier.types';
import { useProductTypes } from '@/modules/product-types/composables/useProductTypes';

import { getCompanyCoverUrl, getCompanyLogoUrl } from '@/shared/utils/assets';
import { getProductImageUrl } from '@/shared/utils/assets';

import { getPurchaseCart } from '../api/purchase-cart.api';

const route = useRoute();
const router = useRouter();

const supplierId = computed(() => Number(route.params.supplierId));

const purchaseCartOpen = ref(false);

const activeSellerCompanyId = ref<number | null>(null);
const addingProductId = ref<number | null>(null);
const filtersOpen = ref(false);
const categoryId = ref('all');
const typeId = ref('all');
const kind = ref('all');
const search = ref('');
const minPrice = ref('');
const maxPrice = ref('');
const sortColumn = ref<'name' | 'created_at' | 'type' | 'price' | 'stock'>('name');
const order = ref<'asc' | 'desc'>('asc');

const page = ref(1);
const limit = ref(12);

const selectedCategoryId = computed<number | null>(() => {
  return categoryId.value !== 'all' ? Number(categoryId.value) : null;
});

watch(categoryId, () => {
  typeId.value = 'all';
  page.value = 1;
});

watch([search, kind, categoryId, typeId, minPrice, maxPrice, sortColumn, order], () => {
  page.value = 1;
});

const { data: categoriesData, isLoading: categoriesLoading } = useCategories();
const { data: productTypesData, isLoading: productTypesLoading } =
  useProductTypes(selectedCategoryId);

const availableProductTypes = computed(() => {
  return productTypesData.value?.items ?? [];
});

const productCategories = computed(() => {
  return categoriesData.value?.product_categories ?? [];
});

const productParams = computed(() => ({
  company_id: supplierId.value,

  search: search.value || undefined,

  kind: kind.value !== 'all' ? (kind.value as 'product' | 'service') : undefined,

  category_id: categoryId.value !== 'all' ? Number(categoryId.value) : undefined,

  type_id: typeId.value !== 'all' ? Number(typeId.value) : undefined,

  min_price: minPrice.value ? Number(minPrice.value) : undefined,

  max_price: maxPrice.value ? Number(maxPrice.value) : undefined,

  sort_column: sortColumn.value,
  order: order.value,

  page: page.value,
  limit: limit.value,
}));

const advancedFiltersCount = computed(() => {
  return [
    kind.value !== 'all' ? kind.value : '',
    typeId.value !== 'all' ? typeId.value : '',
    minPrice.value,
    maxPrice.value,
    sortColumn.value !== 'name' ? sortColumn.value : '',
    order.value !== 'asc' ? order.value : '',
  ].filter(Boolean).length;
});

const { data, isLoading, isError } = useStoreProducts(productParams);

const products = computed(() => data.value?.items ?? []);
const totalPages = computed(() => data.value?.meta.total_pages ?? 1);

function clearFilters() {
  search.value = '';

  kind.value = 'all';
  categoryId.value = 'all';
  typeId.value = 'all';

  minPrice.value = '';
  maxPrice.value = '';

  sortColumn.value = 'name';
  order.value = 'asc';

  page.value = 1;
}

function formatCurrency(value: number) {
  return `$${value.toFixed(2)}`;
}

async function addProductToCart(
  product: StoreProduct,
  quantity = 1,
  openCart = true
): Promise<boolean> {
  if (quantity < 1) {
    toast.error('Enter a valid quantity');
    return false;
  }

  if (product.available_stock <= 0) {
    toast.error('This product is currently unavailable');
    return false;
  }

  if (quantity > product.available_stock) {
    toast.error('Quantity exceeds available stock');
    return false;
  }

  const cartQueryKey = ['purchase-cart', product.company_id] as const;

  try {
    addingProductId.value = product.id;

    /*
     * Si había una consulta vieja o en curso, la cancelamos.
     * Evita que una respuesta anterior sobrescriba el cart nuevo.
     */
    await queryClient.cancelQueries({
      queryKey: cartQueryKey,
    });

    const ensuredCart = await ensureCartMutation.mutateAsync(product.company_id);

    await addCartItemMutation.mutateAsync({
      invoiceId: ensuredCart.invoice_id,
      sellerCompanyId: product.company_id,
      productId: product.id,
      quantity,
    });

    /*
     * Consultamos explícitamente el estado final,
     * después de que el item ya fue agregado.
     */
    const updatedCart = await getPurchaseCart(product.company_id);

    queryClient.setQueryData(cartQueryKey, updatedCart);

    /*
     * Esperamos el catálogo actualizado antes
     * de considerar terminada la operación.
     */
    await queryClient.refetchQueries({
      queryKey: ['store-products'],
      type: 'active',
    });

    toast.success('Product added to purchase', {
      description: `${quantity} × ${product.name}`,
    });

    if (openCart) {
      activeSellerCompanyId.value = product.company_id;

      purchaseCartOpen.value = true;
    }

    return true;
  } catch (error: any) {
    const message =
      error?.response?.data?.errorMessage ??
      error?.response?.data?.message ??
      'Failed to add product to purchase';

    toast.error(message);

    return false;
  } finally {
    addingProductId.value = null;
  }
}

const queryClient = useQueryClient();

const supplier = computed<Supplier | null>(() => {
  const cachedQueries = queryClient.getQueriesData<{
    items: Supplier[];
  }>({
    queryKey: ['marketplace-suppliers'],
  });

  for (const [, cachedData] of cachedQueries) {
    const found = cachedData?.items.find((item) => item.id === supplierId.value);

    if (found) return found;
  }

  return null;
});

const supplierName = computed(() => {
  return supplier.value?.name ?? products.value[0]?.company_name ?? 'Supplier catalog';
});

const supplierCover = computed(() => {
  return getCompanyCoverUrl(supplier.value?.cover_image);
});

const supplierLogo = computed(() => {
  return getCompanyLogoUrl(supplier.value?.logo);
});

const selectedProduct = ref<StoreProduct | null>(null);
const productDetailsOpen = ref(false);

function openProductDetails(product: StoreProduct) {
  selectedProduct.value = product;
  productDetailsOpen.value = true;
}

function addToPurchase(product: StoreProduct) {
  return addProductToCart(product, 1, true);
}

async function handleAddFromDrawer(payload: { product: StoreProduct; quantity: number }) {
  const added = await addProductToCart(payload.product, payload.quantity, false);

  if (!added) return;

  productDetailsOpen.value = false;

  requestAnimationFrame(() => {
    requestAnimationFrame(() => {
      activeSellerCompanyId.value = payload.product.company_id;

      purchaseCartOpen.value = true;
    });
  });
}

const selectedCategory = computed(() =>
  productCategories.value.find((c) => String(c.id) === categoryId.value)
);

const selectedProductType = computed(() => {
  if (typeId.value === 'all') return null;

  return (
    availableProductTypes.value.find((productType) => String(productType.id) === typeId.value) ??
    null
  );
});

const productTypeTriggerLabel = computed(() => {
  if (categoryId.value === 'all') {
    return 'Select category first';
  }

  if (productTypesLoading.value) {
    return 'Loading types...';
  }

  if (typeId.value === 'all') {
    return 'All product types';
  }

  return selectedProductType.value?.name ?? 'Select product type';
});

const ensureCartMutation = useEnsurePurchaseCart();
const addCartItemMutation = useAddPurchaseCartItem();

const isAddingToCart = computed(() => {
  return (
    addingProductId.value !== null ||
    ensureCartMutation.isPending.value ||
    addCartItemMutation.isPending.value
  );
});

function handleCheckoutCompleted(invoiceId: number) {
  console.log('Purchase checkout completed:', invoiceId);

  activeSellerCompanyId.value = null;
}
</script>

<template>
  <PageContainer>
    <div class="space-y-6">
      <Button variant="ghost" size="sm" class="-ml-2" @click="router.push('/erp/marketplace')">
        <ArrowLeft class="mr-2 h-4 w-4" />
        Back to suppliers
      </Button>

      <!-- HERO -->
      <section
        class="relative min-h-64 overflow-hidden rounded-2xl border border-border bg-card shadow-sm"
      >
        <img
          v-if="supplierCover"
          :src="supplierCover"
          :alt="`${supplierName} cover`"
          class="absolute inset-0 h-full w-full object-cover"
        />

        <div
          v-else
          class="absolute inset-0 bg-gradient-to-br from-primary/20 via-muted/50 to-background"
        />

        <div
          class="absolute inset-0 bg-gradient-to-r from-background/95 via-background/75 to-background/25"
        />

        <div
          class="relative flex min-h-64 flex-col justify-between gap-8 p-5 sm:p-8 md:flex-row md:items-end"
        >
          <div class="max-w-2xl">
            <div class="mb-5 flex items-center gap-4">
              <div
                class="flex h-14 w-14 shrink-0 items-center justify-center overflow-hidden rounded-2xl border border-border bg-background shadow-lg sm:h-16 sm:w-16"
              >
                <img
                  v-if="supplierLogo"
                  :src="supplierLogo"
                  :alt="supplierName"
                  class="h-full w-full object-contain p-2"
                />

                <Building2 v-else class="h-7 w-7 text-muted-foreground" />
              </div>

              <span
                class="rounded-full border border-primary/20 bg-primary/10 px-3 py-1 text-xs font-semibold text-primary"
              >
                {{ supplier?.category ?? 'B2B Supplier' }}
              </span>
            </div>

            <h1 class="text-3xl font-semibold tracking-tight text-foreground md:text-4xl">
              {{ supplierName }}
            </h1>

            <p class="mt-3 max-w-xl text-sm leading-6 text-muted-foreground">
              {{
                supplier?.description ??
                'Browse available products, review inventory and create a purchase flow for your company.'
              }}
            </p>
          </div>

          <div
            class="rounded-2xl border border-border bg-background/85 px-5 py-4 shadow-lg backdrop-blur-md"
          >
            <p class="text-xs text-muted-foreground">Available catalog</p>

            <p class="mt-1 text-2xl font-semibold text-foreground">
              {{ data?.meta.total ?? 0 }}
            </p>

            <p class="text-xs text-muted-foreground">products and services</p>
          </div>
        </div>
      </section>

      <!-- FILTERS: AHORA FUERA DEL HERO -->
      <section class="rounded-2xl border border-border bg-card p-5 shadow-sm">
        <div class="grid gap-3 sm:grid-cols-2 xl:grid-cols-[minmax(0,1fr)_240px_auto]">
          <div class="relative">
            <Search
              class="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
            />

            <Input v-model="search" class="pl-9" placeholder="Search supplier products..." />
          </div>

          <Select v-model="categoryId">
            <SelectTrigger>
              <span class="truncate">
                {{ selectedCategory?.name ?? 'Product category' }}
              </span>
            </SelectTrigger>

            <SelectContent>
              <SelectItem value="all"> All categories </SelectItem>

              <SelectItem
                v-for="category in productCategories"
                :key="category.id"
                :value="String(category.id)"
              >
                <div class="flex flex-col">
                  <SelectItemText>
                    {{ category.name }}
                  </SelectItemText>

                  <span v-if="category.description" class="text-xs text-muted-foreground">
                    {{ category.description }}
                  </span>
                </div>
              </SelectItem>
            </SelectContent>
          </Select>

          <Button
            variant="outline"
            class="w-full justify-between sm:col-span-2 xl:col-span-1"
            @click="filtersOpen = !filtersOpen"
          >
            <div class="flex items-center">
              <SlidersHorizontal class="mr-2 h-4 w-4" />
              Filters

              <span
                v-if="advancedFiltersCount > 0"
                class="ml-2 rounded-full bg-primary px-2 py-0.5 text-[10px] font-semibold text-primary-foreground"
              >
                {{ advancedFiltersCount }}
              </span>
            </div>

            <ChevronDown
              class="ml-3 h-4 w-4 transition-transform"
              :class="{ 'rotate-180': filtersOpen }"
            />
          </Button>
        </div>

        <Transition
          enter-active-class="transition-all duration-200 ease-out"
          enter-from-class="max-h-0 opacity-0"
          enter-to-class="max-h-96 opacity-100"
          leave-active-class="transition-all duration-150 ease-in"
          leave-from-class="max-h-96 opacity-100"
          leave-to-class="max-h-0 opacity-0"
        >
          <div
            v-if="filtersOpen"
            class="mt-4 overflow-hidden rounded-xl border border-border bg-muted/20"
          >
            <div class="grid gap-4 p-4 sm:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-6">
              <div class="space-y-2">
                <label class="text-xs font-medium text-muted-foreground"> Resource kind </label>

                <Select v-model="kind">
                  <SelectTrigger>
                    <SelectValue placeholder="All kinds" />
                  </SelectTrigger>

                  <SelectContent>
                    <SelectItem value="all"> All kinds </SelectItem>

                    <SelectItem value="product"> Product </SelectItem>

                    <SelectItem value="service"> Service </SelectItem>
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
                      {{ productTypeTriggerLabel }}
                    </span>
                  </SelectTrigger>

                  <SelectContent>
                    <SelectItem value="all"> All product types </SelectItem>

                    <SelectItem
                      v-for="productType in availableProductTypes"
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
                    <SelectValue />
                  </SelectTrigger>

                  <SelectContent>
                    <SelectItem value="name"> Name </SelectItem>

                    <SelectItem value="price"> Price </SelectItem>

                    <SelectItem value="stock"> Stock </SelectItem>

                    <SelectItem value="created_at"> Created date </SelectItem>
                  </SelectContent>
                </Select>
              </div>

              <div class="space-y-2">
                <label class="text-xs font-medium text-muted-foreground"> Order </label>

                <Select v-model="order">
                  <SelectTrigger>
                    <SelectValue />
                  </SelectTrigger>

                  <SelectContent>
                    <SelectItem value="asc"> Ascending </SelectItem>

                    <SelectItem value="desc"> Descending </SelectItem>
                  </SelectContent>
                </Select>
              </div>
            </div>

            <div class="flex justify-end border-t border-border px-4 py-3">
              <Button variant="ghost" size="sm" @click="clearFilters">
                <RotateCcw class="mr-2 h-4 w-4" />
                Clear all filters
              </Button>
            </div>
          </div>
        </Transition>

        <div class="mt-4 flex flex-col gap-2 sm:flex-row sm:items-center sm:justify-between">
          <p class="text-sm text-muted-foreground">{{ data?.meta.total ?? 0 }} catalog results</p>

          <div class="flex items-center gap-2 text-xs text-muted-foreground">
            <Filter class="h-3.5 w-3.5" />
            Marketplace catalog
          </div>
        </div>
      </section>

      <!-- ESTADOS Y PRODUCTOS -->
      <div
        v-if="isLoading"
        class="rounded-xl border border-border bg-muted/30 p-8 text-center text-sm text-muted-foreground"
      >
        Loading supplier catalog...
      </div>

      <EmptyState
        v-else-if="isError"
        title="Unable to load supplier catalog"
        description="There was a problem connecting to the marketplace products API."
        :icon="PackageSearch"
      />

      <EmptyState
        v-else-if="products.length === 0"
        title="No products found"
        description="Try changing your catalog filters."
        :icon="PackageSearch"
      />

      <template v-else>
        <div class="grid gap-5 md:grid-cols-2 xl:grid-cols-3">
          <article
            v-for="product in products"
            :key="product.id"
            class="group relative flex flex-col overflow-hidden rounded-2xl border border-border bg-card shadow-sm transition duration-300 hover:-translate-y-1 hover:border-primary/30 hover:shadow-xl"
          >
            <div
              class="pointer-events-none absolute inset-0 opacity-0 transition duration-300 group-hover:opacity-100"
            >
              <div
                class="absolute inset-x-0 top-0 h-px bg-gradient-to-r from-transparent via-primary/70 to-transparent"
              />
            </div>

            <button type="button" class="text-left" @click="openProductDetails(product)">
              <div class="relative h-56 overflow-hidden bg-muted/30">
                <img
                  v-if="getProductImageUrl(product.image_path)"
                  :src="getProductImageUrl(product.image_path)!"
                  :alt="product.name"
                  class="h-full w-full object-cover transition duration-500 group-hover:scale-105"
                />

                <div
                  v-else
                  class="flex h-full w-full items-center justify-center bg-gradient-to-br from-muted to-background"
                >
                  <PackageSearch class="h-12 w-12 text-muted-foreground/50" />
                </div>

                <div
                  class="absolute inset-0 bg-gradient-to-t from-background/90 via-background/10 to-transparent"
                />

                <div class="absolute left-4 top-4 flex max-w-[75%] flex-wrap gap-2">
                  <span
                    class="rounded-full bg-primary/90 px-3 py-1 text-xs font-semibold text-primary-foreground shadow-sm"
                  >
                    {{ product.category }}
                  </span>

                  <span
                    class="rounded-full bg-background/90 px-3 py-1 text-xs font-medium text-foreground shadow-sm"
                  >
                    {{ product.type }}
                  </span>
                </div>

                <div class="absolute bottom-4 left-4 right-4 flex items-end justify-between gap-4">
                  <div class="min-w-0">
                    <p class="text-xs font-medium text-muted-foreground">
                      {{ product.company_name }}
                    </p>

                    <h2 class="mt-1 line-clamp-2 text-lg font-semibold text-foreground">
                      {{ product.name }}
                    </h2>
                  </div>

                  <div
                    class="shrink-0 rounded-xl border border-border bg-background/90 px-3 py-2 text-right shadow-md backdrop-blur"
                  >
                    <p class="text-[10px] uppercase tracking-wide text-muted-foreground">From</p>

                    <p class="text-lg font-semibold tracking-tight text-foreground">
                      {{ formatCurrency(product.price) }}
                    </p>
                  </div>
                </div>
              </div>
            </button>

            <div class="flex flex-1 flex-col p-5">
              <div class="flex items-center justify-between gap-3">
                <div class="flex items-center gap-2">
                  <span
                    class="h-2.5 w-2.5 rounded-full"
                    :class="
                      product.available_stock > 10
                        ? 'bg-emerald-500'
                        : product.available_stock > 0
                          ? 'bg-amber-500'
                          : 'bg-destructive'
                    "
                  />

                  <p class="text-sm text-muted-foreground">
                    {{
                      product.available_stock > 0
                        ? `${product.available_stock} units available`
                        : 'Currently unavailable'
                    }}
                  </p>
                </div>

                <span
                  v-if="product.available_stock > 0 && product.available_stock <= 10"
                  class="rounded-full bg-amber-500/10 px-2 py-1 text-[11px] font-medium text-amber-600 dark:text-amber-400"
                >
                  Low stock
                </span>
              </div>

              <div class="mt-5 grid grid-cols-2 gap-2 border-t border-border pt-4">
                <Button
                  variant="ghost"
                  class="w-full rounded-full"
                  size="sm"
                  @click="openProductDetails(product)"
                >
                  View details
                </Button>

                <Button
                  size="sm"
                  class="w-full rounded-full"
                  :disabled="product.available_stock <= 0 || isAddingToCart"
                  @click="addToPurchase(product)"
                >
                  <Loader2
                    v-if="addingProductId === product.id"
                    class="mr-2 h-4 w-4 animate-spin"
                  />

                  <ShoppingCart v-else class="mr-2 h-4 w-4" />

                  {{ addingProductId === product.id ? 'Adding...' : 'Add' }}
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
    </div>

    <MarketplaceProductDrawer
      v-model:open="productDetailsOpen"
      :product="selectedProduct"
      :adding="selectedProduct ? addingProductId === selectedProduct.id : false"
      @add="handleAddFromDrawer"
    />

    <PurchaseCartDrawer
      v-model:open="purchaseCartOpen"
      :seller-company-id="activeSellerCompanyId"
      @checkout-completed="handleCheckoutCompleted"
    />
  </PageContainer>
</template>
