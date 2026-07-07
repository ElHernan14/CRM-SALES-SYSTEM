<script setup lang="ts">
import { computed, ref, toRef } from 'vue';
import { Minus, Plus, Trash2, PackagePlus, Loader2 } from 'lucide-vue-next';
import { toast } from 'vue-sonner';

import { useUiStore } from '@/shared/stores/ui.store';
import { useAuthStore } from '@/modules/auth/stores/auth.store';
import { storeToRefs } from 'pinia';

import type { InvoiceItem } from '../types/invoice-item.types';

import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
} from '@/components/ui/sheet';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';

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

const page = ref(1);
const limit = ref(10);

const params = computed(() => ({
  page: page.value,
  limit: limit.value,
}));

const { data, isLoading, isError } = useInvoiceItems(invoiceIdRef, params);

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

const isMutatingItems = computed(() => {
  return (
    createItemMutation.isPending.value ||
    updateItemMutation.isPending.value ||
    deleteItemMutation.isPending.value
  );
});

const selectedProductId = ref('');
const newItemQuantity = ref('1');

const productParams = computed(() => ({
  company_id: user.value?.company_id,
  page: 1,
  limit: 100,
}));

const { data: productsData, isLoading: productsLoading } = useProducts(productParams);

const availableProducts = computed(() => {
  return productsData.value?.data ?? [];
});

const selectedProduct = computed(() => {
  return availableProducts.value.find((product) => product.id === Number(selectedProductId.value));
});

const selectedProductAvailableStock = computed(() => {
  if (!selectedProduct.value) return 0;

  return selectedProduct.value.available_stock ?? selectedProduct.value.stock;
});

async function addItem() {
  if (!props.invoiceId) return;

  const productId = Number(selectedProductId.value);
  const quantity = Number(newItemQuantity.value);

  if (!productId) {
    toast.error('Select a product');
    return;
  }

  if (!quantity || quantity <= 0) {
    toast.error('Enter a valid quantity');
    return;
  }

  if (selectedProduct.value && selectedProductAvailableStock.value <= 0) {
    toast.error('This product has no available stock');
    return;
  }

  if (selectedProduct.value && quantity > selectedProductAvailableStock.value) {
    toast.error('Quantity exceeds available stock');
    return;
  }

  try {
    await createItemMutation.mutateAsync({
      invoiceId: props.invoiceId,
      payload: {
        product_id: productId,
        quantity,
      },
    });

    toast.success('Item added');

    selectedProductId.value = '';
    newItemQuantity.value = '1';
  } catch {
    toast.error('Failed to add item');
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
            <Select v-model="selectedProductId">
              <SelectTrigger class="min-w-0">
                <SelectValue
                  :placeholder="productsLoading ? 'Loading products...' : 'Select product'"
                />
              </SelectTrigger>

              <SelectContent>
                <SelectItem
                  v-for="product in availableProducts"
                  :key="product.id"
                  :value="String(product.id)"
                >
                  {{ product.name }} · ${{ product.price.toFixed(2) }}
                </SelectItem>
              </SelectContent>
            </Select>

            <Input v-model="newItemQuantity" type="text" inputmode="numeric" placeholder="Qty" />

            <Button type="button" :disabled="createItemMutation.isPending.value" @click="addItem">
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
            class="mt-3 flex items-center justify-between rounded-lg border border-border bg-muted/40 px-3 py-2"
          >
            <span class="text-sm text-muted-foreground"> Available stock </span>

            <span
              class="rounded-full px-2 py-1 text-xs font-medium"
              :class="
                selectedProductAvailableStock > 0
                  ? 'bg-emerald-500/10 text-emerald-600 dark:text-emerald-400'
                  : 'bg-destructive/10 text-destructive'
              "
            >
              {{ selectedProductAvailableStock }} units
            </span>
          </div>
        </div>

        <div
          v-if="isLoading"
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
            v-if="isMutatingItems"
            class="absolute inset-0 z-10 rounded-xl bg-background/60 backdrop-blur-[1px]"
          />

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

        <div class="relative rounded-xl border border-border bg-muted/40 p-4">
          <div
            v-if="isMutatingItems"
            class="absolute inset-0 rounded-xl bg-background/50 backdrop-blur-[1px]"
          />

          <div class="relative flex items-center justify-between">
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
