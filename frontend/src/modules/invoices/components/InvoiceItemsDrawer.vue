<script setup lang="ts">
import { computed, onUnmounted, ref, toRef, watch } from 'vue';
import { Minus, Plus, Trash2, PackagePlus, Loader2, Search, Check } from 'lucide-vue-next';
import { toast } from 'vue-sonner';

import { useUiStore } from '@/shared/stores/ui.store';
import { useAuthStore } from '@/modules/auth/stores/auth.store';
import { storeToRefs } from 'pinia';

import type { InvoiceItem } from '../types/invoice-item.types';
import type { ProductListItem } from '@/modules/products/types/product.types';

import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
} from '@/components/ui/sheet';
import SkeletonBlock from '@/shared/components/erp/SkeletonBlock.vue';

import { useInvoiceItems } from '../composables/useInvoiceItems';
import { useUpdateInvoiceItem } from '../composables/useUpdateInvoiceItem';
import { useDeleteInvoiceItem } from '../composables/useDeleteInvoiceItem';
import { useProducts } from '@/modules/products/composables/useProducts';
import { useCreateInvoiceItem } from '../composables/useCreateInvoiceItem';

// Stores
const ui = useUiStore();
const auth = useAuthStore();

const { user } = storeToRefs(auth);

// Props and emits
const props = defineProps<{
  open: boolean;
  invoiceId: number | null;
}>();

const emit = defineEmits<{
  'update:open': [value: boolean];
}>();

const invoiceIdRef = toRef(props, 'invoiceId');

const productSearch = ref('');
const selectedProduct = ref<ProductListItem | null>(null);
const suggestionsOpen = ref(false);

const page = ref(1);
const limit = ref(10);

const debouncedProductSearch = ref('');

let productSearchTimer: ReturnType<typeof setTimeout> | undefined;

watch(productSearch, (value) => {
  if (selectedProduct.value && value !== selectedProduct.value.name) {
    selectedProduct.value = null;
  }

  if (productSearchTimer) {
    clearTimeout(productSearchTimer);
  }

  productSearchTimer = setTimeout(() => {
    debouncedProductSearch.value = value.trim();
  }, 300);
});

onUnmounted(() => {
  if (productSearchTimer) {
    clearTimeout(productSearchTimer);
  }
});

const params = computed(() => ({
  page: page.value,
  limit: limit.value,
}));

const { data, isLoading, isError, isFetching } = useInvoiceItems(invoiceIdRef, params);

const items = computed(() => data.value?.items ?? []);

const total = computed(() => {
  return items.value.reduce((acc, item) => acc + item.subtotal, 0);
});

function formatCurrency(value: number) {
  return `$${value.toFixed(2)}`;
}

// Mutations
const updateItemMutation = useUpdateInvoiceItem();
const deleteItemMutation = useDeleteInvoiceItem();
const createItemMutation = useCreateInvoiceItem();

const isRefreshingItems = computed(() => {
  return isFetching.value && !isLoading.value;
});

const newItemQuantity = ref('1');

const productParams = computed(() => ({
  company_id: user.value?.company_id,

  search: debouncedProductSearch.value.length >= 2 ? debouncedProductSearch.value : undefined,

  status: 1 as const,

  page: 1,
  limit: 15,
}));

const { data: productsData, isLoading: productsLoading } = useProducts(productParams);

const productSuggestions = computed(() => {
  if (debouncedProductSearch.value.length < 2 || selectedProduct.value) {
    return [];
  }

  return productsData.value?.data ?? [];
});

const showProductSuggestions = computed(() => {
  return suggestionsOpen.value && productSearch.value.trim().length >= 2 && !selectedProduct.value;
});

function selectProduct(product: ProductListItem) {
  selectedProduct.value = product;
  productSearch.value = product.name;
  suggestionsOpen.value = false;
}

function clearSelectedProduct() {
  selectedProduct.value = null;
  productSearch.value = '';
  debouncedProductSearch.value = '';
  suggestionsOpen.value = false;
}

const selectedProductAvailableStock = computed(() => {
  if (!selectedProduct.value) return 0;

  return selectedProduct.value.available_stock ?? selectedProduct.value.stock;
});

async function addItem() {
  if (!props.invoiceId) return;

  const product = selectedProduct.value;
  const quantity = Number(newItemQuantity.value);

  if (!product) {
    toast.error('Select a product');
    return;
  }

  if (!Number.isInteger(quantity) || quantity <= 0) {
    toast.error('Enter a valid quantity');
    return;
  }

  if (selectedProductAvailableStock.value <= 0) {
    toast.error('This product has no available stock');
    return;
  }

  if (quantity > selectedProductAvailableStock.value) {
    toast.error('Quantity exceeds available stock');
    return;
  }

  try {
    await createItemMutation.mutateAsync({
      invoiceId: props.invoiceId,

      payload: {
        product_id: product.id,
        quantity,
      },
    });

    toast.success('Item added');

    clearSelectedProduct();
    newItemQuantity.value = '1';
  } catch (error: any) {
    const message =
      error?.response?.data?.errorMessage ?? error?.response?.data?.message ?? 'Failed to add item';

    toast.error(message);
  }
}

async function updateQuantity(item: InvoiceItem, nextQuantity: number) {
  if (!props.invoiceId) return;
  if (nextQuantity < 1) return;

  try {
    await updateItemMutation.mutateAsync({
      invoiceId: props.invoiceId,
      itemId: item.id,
      payload: {
        quantity: nextQuantity,
      },
    });

    toast.success('Item quantity updated');
  } catch {
    toast.error('Failed to update item');
  }
}

function confirmDeleteItem(item: InvoiceItem) {
  if (!props.invoiceId) return;

  ui.openConfirm({
    title: 'Remove invoice item?',
    description: `This will remove "${item.product_name}" from the invoice.`,
    confirmText: 'Remove item',
    cancelText: 'Cancel',
    variant: 'destructive',
    onConfirm: async () => {
      if (!props.invoiceId) return;

      await deleteItemMutation.mutateAsync({
        invoiceId: props.invoiceId,
        itemId: item.id,
      });

      toast.success('Item removed');
    },
  });
}
</script>

<template>
  <Sheet :open="open" @update:open="emit('update:open', $event)">
    <SheetContent class="w-full overflow-y-auto sm:max-w-3xl">
      <SheetHeader>
        <SheetTitle> Manage invoice items </SheetTitle>

        <SheetDescription>
          Add, review and adjust products before submitting the invoice.
        </SheetDescription>
      </SheetHeader>

      <div class="mt-6 space-y-6">
        <div class="rounded-xl border border-border bg-card p-4 shadow-sm">
          <div class="mb-4 flex items-center gap-2">
            <div class="rounded-lg border border-border bg-muted p-2 text-muted-foreground">
              <PackagePlus class="h-4 w-4" />
            </div>

            <div>
              <p class="text-sm font-medium text-foreground">Add product</p>
              <p class="text-xs text-muted-foreground">
                Select a product and quantity for this draft invoice.
              </p>
            </div>
          </div>

          <div class="grid gap-3 lg:grid-cols-[minmax(0,1fr)_120px_auto]">
            <div class="relative min-w-0">
              <div class="relative">
                <Search
                  class="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
                />

                <Input
                  v-model="productSearch"
                  class="pl-9"
                  autocomplete="off"
                  placeholder="Search product by name..."
                  :disabled="createItemMutation.isPending.value"
                  @focus="suggestionsOpen = true"
                  @input="suggestionsOpen = true"
                />
              </div>

              <div
                v-if="showProductSuggestions"
                class="absolute left-0 right-0 top-full z-50 mt-2 max-h-80 overflow-y-auto rounded-xl border border-border bg-popover p-1 shadow-xl"
              >
                <div
                  v-if="productsLoading"
                  class="flex items-center gap-2 px-3 py-4 text-sm text-muted-foreground"
                >
                  <Loader2 class="h-4 w-4 animate-spin" />
                  Searching products...
                </div>

                <button
                  v-for="product in productSuggestions"
                  v-else
                  :key="product.id"
                  type="button"
                  class="flex w-full items-center justify-between gap-4 rounded-lg px-3 py-3 text-left transition hover:bg-muted"
                  @mousedown.prevent="selectProduct(product)"
                >
                  <div class="min-w-0">
                    <p class="truncate text-sm font-medium text-foreground">
                      {{ product.name }}
                    </p>

                    <div
                      class="mt-1 flex flex-wrap items-center gap-2 text-xs text-muted-foreground"
                    >
                      <span>
                        {{ product.category }}
                      </span>

                      <span>→</span>

                      <span>
                        {{ product.type }}
                      </span>

                      <span>·</span>

                      <span>
                        {{ formatCurrency(product.price) }}
                      </span>
                    </div>
                  </div>

                  <div class="shrink-0 text-right">
                    <p
                      class="text-xs font-medium"
                      :class="
                        (product.available_stock ?? product.stock) > 0
                          ? 'text-emerald-600 dark:text-emerald-400'
                          : 'text-destructive'
                      "
                    >
                      {{ product.available_stock ?? product.stock }}
                      available
                    </p>
                  </div>
                </button>

                <div
                  v-if="!productsLoading && productSuggestions.length === 0"
                  class="px-3 py-6 text-center"
                >
                  <p class="text-sm font-medium text-foreground">No matching products</p>

                  <p class="mt-1 text-xs text-muted-foreground">Try another product name.</p>
                </div>
              </div>
            </div>

            <Input v-model="newItemQuantity" type="text" inputmode="numeric" placeholder="Qty" />

            <Button
              type="button"
              :disabled="!selectedProduct || createItemMutation.isPending.value"
              @click="addItem"
            >
              <Loader2
                v-if="createItemMutation.isPending.value"
                class="mr-2 h-4 w-4 animate-spin"
              />

              <PackagePlus v-else class="mr-2 h-4 w-4" />

              {{ createItemMutation.isPending.value ? 'Adding...' : 'Add' }}
            </Button>
          </div>

          <div
            v-if="selectedProduct"
            class="mt-3 rounded-xl border border-primary/20 bg-primary/5 p-4"
          >
            <div class="flex items-start justify-between gap-4">
              <div class="flex min-w-0 items-start gap-3">
                <div
                  class="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl border border-border bg-background"
                >
                  <Check class="h-4 w-4 text-primary" />
                </div>

                <div class="min-w-0">
                  <p class="truncate text-sm font-semibold text-foreground">
                    {{ selectedProduct.name }}
                  </p>

                  <p class="mt-1 text-xs text-muted-foreground">
                    {{ selectedProduct.category }}
                    →
                    {{ selectedProduct.type }}
                    ·
                    {{ formatCurrency(selectedProduct.price) }}
                  </p>
                </div>
              </div>

              <Button
                type="button"
                variant="ghost"
                size="sm"
                :disabled="createItemMutation.isPending.value"
                @click="clearSelectedProduct"
              >
                Change
              </Button>
            </div>

            <div class="mt-4 flex items-center justify-between border-t border-primary/10 pt-3">
              <span class="text-sm text-muted-foreground"> Available stock </span>

              <span
                class="rounded-full px-2.5 py-1 text-xs font-medium"
                :class="
                  selectedProductAvailableStock > 0
                    ? 'bg-emerald-500/10 text-emerald-600 dark:text-emerald-400'
                    : 'bg-destructive/10 text-destructive'
                "
              >
                {{ selectedProductAvailableStock }}
                units
              </span>
            </div>
          </div>
        </div>

        <div
          v-if="isLoading || isRefreshingItems"
          class="rounded-xl border border-border bg-muted/30 p-6 text-sm text-muted-foreground"
        >
          Loading invoice items...
        </div>

        <div
          v-else-if="isError"
          class="rounded-xl border border-border bg-muted/30 p-6 text-sm text-destructive"
        >
          Unable to load invoice items.
        </div>

        <div
          v-else-if="items.length === 0"
          class="rounded-xl border border-dashed border-border bg-card p-10 text-center"
        >
          <p class="text-sm font-medium text-foreground">No items added yet</p>

          <p class="mt-1 text-sm text-muted-foreground">
            Add products to build this draft invoice.
          </p>
        </div>

        <div v-else class="relative space-y-3">
          <div
            v-for="item in items"
            :key="item.id"
            class="rounded-xl border border-border bg-card p-4 shadow-sm"
          >
            <div class="flex flex-col gap-4 md:flex-row md:items-center md:justify-between">
              <div>
                <p class="font-medium text-foreground">
                  {{ item.product_name }}
                </p>

                <p class="mt-1 text-sm text-muted-foreground">
                  {{ formatCurrency(item.price) }} · Product #{{ item.product_id }}
                </p>
              </div>

              <div class="flex items-center gap-4">
                <div class="flex items-center rounded-full border border-border bg-muted/40 p-1">
                  <Button
                    size="icon"
                    variant="ghost"
                    class="h-8 w-8 rounded-full"
                    :disabled="item.quantity <= 1 || updateItemMutation.isPending.value"
                    @click="updateQuantity(item, item.quantity - 1)"
                  >
                    <Loader2
                      v-if="updateItemMutation.isPending.value"
                      class="h-4 w-4 animate-spin"
                    />
                    <Minus v-else class="h-4 w-4" />
                  </Button>

                  <span class="min-w-10 text-center text-sm font-medium">
                    {{ item.quantity }}
                  </span>

                  <Button
                    size="icon"
                    variant="ghost"
                    class="h-8 w-8 rounded-full"
                    :disabled="updateItemMutation.isPending.value"
                    @click="updateQuantity(item, item.quantity + 1)"
                  >
                    <Loader2
                      v-if="updateItemMutation.isPending.value"
                      class="h-4 w-4 animate-spin"
                    />
                    <Plus v-else class="h-4 w-4" />
                  </Button>
                </div>

                <div class="min-w-24 text-right">
                  <p class="text-sm font-semibold text-foreground">
                    {{ formatCurrency(item.subtotal) }}
                  </p>

                  <p class="text-xs text-muted-foreground">subtotal</p>
                </div>

                <Button
                  size="icon"
                  variant="ghost"
                  class="text-destructive"
                  :disabled="deleteItemMutation.isPending.value"
                  @click="confirmDeleteItem(item)"
                >
                  <Trash2 class="h-4 w-4" />
                </Button>
              </div>
            </div>
          </div>
        </div>

        <div class="relative overflow-hidden rounded-xl border border-border bg-muted/40 p-4">
          <div v-if="isRefreshingItems">
            <SkeletonBlock />
          </div>

          <div v-else class="flex items-center justify-between">
            <p class="text-sm text-muted-foreground">Items subtotal</p>

            <p class="text-lg font-semibold text-foreground">
              {{ formatCurrency(total) }}
            </p>
          </div>
        </div>
      </div>
    </SheetContent>
  </Sheet>
</template>
