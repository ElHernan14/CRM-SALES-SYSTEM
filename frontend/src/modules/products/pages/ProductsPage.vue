<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import { Package, MoreHorizontal, Plus } from 'lucide-vue-next';
import { toast } from 'vue-sonner';

import { Button } from '@/components/ui/button';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';

import CreateProductDrawer from '../components/CreateProductDrawer.vue';
import PageContainer from '@/shared/components/erp/PageContainer.vue';
import PageHeader from '@/shared/components/erp/PageHeader.vue';
import SectionCard from '@/shared/components/erp/SectionCard.vue';
import EmptyState from '@/shared/components/erp/EmptyState.vue';
import DataTable from '@/shared/components/erp/DataTable.vue';
import DataPagination from '@/shared/components/erp/DataPagination.vue';
import BulkActionBar from '@/shared/components/erp/BulkActionBar.vue';
import ProductImageDrawer from '../components/ProductImageDrawer.vue';

import ProductsFilters from '../components/ProductsFilters.vue';
import ProductDetailsDrawer from '../components/ProductDetailsDrawer.vue';
import ProductEditDrawer from '../components/ProductEditDrawer.vue';
import type { ProductListItem } from '../types/product.types';

import { storeToRefs } from 'pinia';
import { useAuthStore } from '@/modules/auth/stores/auth.store';
import { useUiStore } from '@/shared/stores/ui.store';

import { useProducts } from '../composables/useProducts';
import { useDeleteProduct } from '../composables/useDeleteProduct';

// Store
const auth = useAuthStore();
const { user } = storeToRefs(auth);
const ui = useUiStore();

// State image
const imageOpen = ref(false);
const imageProductId = ref<number | null>(null);

// State for product details drawer
const selectedProduct = ref<ProductListItem | null>(null);
const detailsOpen = ref(false);

function openProductDetails(row: Record<string, unknown>) {
  const product = data.value?.data.find((item) => item.id === Number(row.id));

  if (!product) return;

  selectedProduct.value = product;
  detailsOpen.value = true;
}

// State Create Product
const createOpen = ref(false);

function handleProductCreated(productId: number) {
  imageProductId.value = productId;
  imageOpen.value = true;
}

// State for product edit drawer
const editOpen = ref(false);
const editingProduct = ref<ProductListItem | null>(null);

function openEditProduct(row: Record<string, unknown>) {
  const product = data.value?.data.find((item) => item.id === Number(row.id));

  if (!product) return;

  editingProduct.value = product;
  editOpen.value = true;
}

// Filters
const search = ref('');
const page = ref(1);
const limit = ref(10);
const kind = ref('all');
const categoryId = ref('all');
const typeId = ref('all');
const minPrice = ref('');
const maxPrice = ref('');

const productParams = computed(() => ({
  search: search.value || undefined,

  kind: kind.value !== 'all' ? (kind.value as 'product' | 'service') : undefined,

  category_id: categoryId.value !== 'all' ? Number(categoryId.value) : undefined,

  type_id: typeId.value !== 'all' ? Number(typeId.value) : undefined,

  min_price: minPrice.value ? Number(minPrice.value) : undefined,

  max_price: maxPrice.value ? Number(maxPrice.value) : undefined,

  company_id: user.value?.company_id,

  page: page.value,
  limit: limit.value,
}));

watch([search, kind, categoryId, typeId, minPrice, maxPrice], () => {
  page.value = 1;
});

function clearFilters() {
  search.value = '';

  kind.value = 'all';
  categoryId.value = 'all';
  typeId.value = 'all';

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
  { key: 'kind', label: 'Kind' },
  { key: 'category', label: 'Category' },
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

      kind: product.kind,

      category: product.category,
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

// Selected product IDs
const selectedProductIds = ref<number[]>([]);

function toggleProduct(row: Record<string, unknown>) {
  const id = Number(row.id);

  if (selectedProductIds.value.includes(id)) {
    selectedProductIds.value = selectedProductIds.value.filter((productId) => productId !== id);
    return;
  }

  selectedProductIds.value.push(id);
}

function toggleAllProducts() {
  const currentIds = rows.value.map((row) => Number(row.id));

  const allSelected = currentIds.every((id) => selectedProductIds.value.includes(id));

  if (allSelected) {
    selectedProductIds.value = selectedProductIds.value.filter((id) => !currentIds.includes(id));
    return;
  }

  selectedProductIds.value = Array.from(new Set([...selectedProductIds.value, ...currentIds]));
}

const allSelected = computed(() => {
  if (rows.value.length === 0) return false;

  return rows.value.every((row) => selectedProductIds.value.includes(Number(row.id)));
});

function confirmBulkDelete() {
  ui.openConfirm({
    title: 'Delete selected products?',
    description: `This action will delete ${selectedProductIds.value.length} selected products.`,
    confirmText: 'Delete products',
    cancelText: 'Cancel',
    variant: 'destructive',
    onConfirm: async () => {
      toast.success('Selected products deleted');
      selectedProductIds.value = [];
    },
  });
}

// Confirm delete product
function confirmDeleteProduct(row: Record<string, unknown>) {
  const id = Number(row.id);
  const name = String(row.name);

  ui.openConfirm({
    title: 'Delete product?',
    description: `This will delete "${name}" from the product catalog.`,
    confirmText: 'Delete product',
    cancelText: 'Cancel',
    variant: 'destructive',
    onConfirm: async () => {
      await deleteMutation.mutateAsync(id);

      toast.success('Product deleted');

      selectedProductIds.value = selectedProductIds.value.filter((productId) => productId !== id);
    },
  });
}

const deleteMutation = useDeleteProduct();

function openProductImage(row: Record<string, unknown>) {
  imageProductId.value = Number(row.id);
  imageOpen.value = true;
}
</script>

<template>
  <PageContainer>
    <PageHeader title="Products" description="Manage product catalog, pricing and inventory." />

    <SectionCard
      title="Product catalog"
      description="Search and review products available for this tenant."
    >
      <template #actions>
        <div class="flex items-center gap-2">
          <Button variant="outline" size="sm" @click="refetch()"> Refresh </Button>

          <Button size="sm" @click="createOpen = true">
            <Plus class="mr-2 h-4 w-4" />
            New product
          </Button>
        </div>
      </template>

      <ProductsFilters
        v-model:search="search"
        v-model:kind="kind"
        v-model:category-id="categoryId"
        v-model:type-id="typeId"
        v-model:min-price="minPrice"
        v-model:max-price="maxPrice"
        :total="data?.meta.total ?? 0"
        @clear="clearFilters"
      />

      <BulkActionBar :selected-count="selectedProductIds.length">
        <Button variant="outline" size="sm"> Export </Button>

        <Button variant="destructive" size="sm" @click="confirmBulkDelete">
          Delete selected
        </Button>
      </BulkActionBar>

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

      <DataTable
        v-else
        selectable
        :columns="columns"
        :rows="rows"
        :selected-rows="selectedProductIds"
        :header-checked="allSelected"
        @toggle-row="toggleProduct"
        @toggle-all="toggleAllProducts"
      >
        <template #cell-price="{ value }"> ${{ Number(value).toFixed(2) }} </template>

        <template #cell-kind="{ value }">
          <span
            class="rounded-full border border-border bg-muted px-2 py-1 text-xs font-medium capitalize text-muted-foreground"
          >
            {{ value }}
          </span>
        </template>

        <template #cell-category="{ value }">
          <span class="rounded-full bg-primary/10 px-2 py-1 text-xs font-medium text-primary">
            {{ value }}
          </span>
        </template>

        <template #cell-type="{ value }">
          <span class="text-sm font-medium text-foreground">
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
                <DropdownMenuItem @click="openProductImage(row)"> Change image </DropdownMenuItem>

                <DropdownMenuItem @click="openProductDetails(row)"> View details </DropdownMenuItem>

                <DropdownMenuItem @click="openEditProduct(row)"> Edit product </DropdownMenuItem>

                <DropdownMenuItem class="text-destructive" @click="confirmDeleteProduct(row)">
                  Delete product
                </DropdownMenuItem>
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

    <!-- Product Details Drawer -->
    <ProductDetailsDrawer v-model:open="detailsOpen" :product="selectedProduct" />

    <!-- Product Create Drawer -->
    <CreateProductDrawer v-model:open="createOpen" @created="handleProductCreated" />

    <!-- Product Edit Drawer -->
    <ProductEditDrawer v-model:open="editOpen" :product="editingProduct" />

    <ProductImageDrawer v-model:open="imageOpen" :product-id="imageProductId" />
  </PageContainer>
</template>
