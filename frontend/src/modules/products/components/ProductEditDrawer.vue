<script setup lang="ts">
import { computed, reactive, watch } from 'vue';
import { toast } from 'vue-sonner';

import { useCategories } from '@/modules/categories/composables/useCategories';
import { useProductTypes } from '@/modules/product-types/composables/useProductTypes';
import { useUpdateProduct } from '../composables/useUpdateProduct';

import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
} from '@/components/ui/sheet';

import { Select, SelectContent, SelectItem, SelectTrigger } from '@/components/ui/select';

import type { ProductListItem } from '../types/product.types';

const props = defineProps<{
  open: boolean;
  product: ProductListItem | null;
}>();

const emit = defineEmits<{
  'update:open': [value: boolean];
}>();

const form = reactive({
  name: '',
  description: '',

  kind: 'product' as 'product' | 'service',

  category_id: 'all',
  type_id: 'all',

  price: '',
  stock: '',
  status: '1',
});

const { data: categoriesData, isLoading: categoriesLoading } = useCategories();

const productCategories = computed(() => {
  return categoriesData.value?.product_categories ?? [];
});

const selectedCategoryId = computed<number | null>(() => {
  return form.category_id !== 'all' ? Number(form.category_id) : null;
});

const { data: productTypesData, isLoading: productTypesLoading } =
  useProductTypes(selectedCategoryId);

const availableProductTypes = computed(() => {
  return productTypesData.value?.items ?? [];
});

const selectedCategory = computed(() => {
  if (form.category_id === 'all') return null;

  return (
    productCategories.value.find((category) => String(category.id) === form.category_id) ?? null
  );
});

const categoryTriggerLabel = computed(() => {
  if (categoriesLoading.value) {
    return 'Loading categories...';
  }

  return selectedCategory.value?.name ?? 'Select category';
});

const selectedProductType = computed(() => {
  if (form.type_id === 'all') return null;

  return (
    availableProductTypes.value.find((productType) => String(productType.id) === form.type_id) ??
    null
  );
});

const productTypeTriggerLabel = computed(() => {
  if (form.category_id === 'all') {
    return 'Select category first';
  }

  if (productTypesLoading.value) {
    return 'Loading types...';
  }

  return selectedProductType.value?.name ?? 'Select product type';
});

const updateMutation = useUpdateProduct();

const isSubmitting = computed(() => updateMutation.isPending.value);

let hydratingProduct = false;

watch(
  () => props.product,
  (product) => {
    if (!product) return;

    hydratingProduct = true;

    form.name = product.name;
    form.description = product.description ?? '';
    form.kind = product.kind;

    form.category_id = String(product.category_id);
    form.type_id = String(product.type_id);

    form.price = String(product.price);
    form.stock = String(product.stock);
    form.status = String(product.status);

    queueMicrotask(() => {
      hydratingProduct = false;
    });
  },
  {
    immediate: true,
  }
);

watch(
  () => form.category_id,
  () => {
    if (hydratingProduct) return;

    form.type_id = 'all';
  }
);

async function onSubmit() {
  if (!props.product) return;

  const name = form.name.trim();
  const description = form.description.trim();

  if (name.length < 2) {
    toast.error('Enter a valid product name');
    return;
  }

  if (form.category_id === 'all') {
    toast.error('Select a product category');
    return;
  }

  if (form.type_id === 'all') {
    toast.error('Select a product type');
    return;
  }

  const price = Number(form.price);
  const stock = Number(form.stock);

  if (Number.isNaN(price) || price < 0) {
    toast.error('Enter a valid price');
    return;
  }

  if (!Number.isInteger(stock) || stock < 0) {
    toast.error('Enter a valid stock');
    return;
  }

  try {
    await updateMutation.mutateAsync({
      id: props.product.id,

      payload: {
        name,
        description: description || undefined,

        kind: form.kind,

        category_id: Number(form.category_id),
        type_id: Number(form.type_id),

        price,
        stock,

        status: Number(form.status) as 0 | 1,
      },
    });

    toast.success('Product updated');
    emit('update:open', false);
  } catch (error: any) {
    const message =
      error?.response?.data?.errorMessage ??
      error?.response?.data?.message ??
      'Failed to update product';

    toast.error(message);
  }
}

function validateDecimalInput(event: KeyboardEvent) {
  const allowedKeys = [
    'Backspace',
    'Delete',
    'ArrowLeft',
    'ArrowRight',
    'Home',
    'End',
    'Tab',
    'Enter',
  ];

  if (allowedKeys.includes(event.key) || event.ctrlKey || event.metaKey) {
    return;
  }

  if (!/^[0-9.]$/.test(event.key)) {
    event.preventDefault();
    return;
  }

  const input = event.currentTarget as HTMLInputElement;

  if (event.key === '.' && input.value.includes('.')) {
    event.preventDefault();
  }
}

function validateIntegerInput(event: KeyboardEvent) {
  const allowedKeys = [
    'Backspace',
    'Delete',
    'ArrowLeft',
    'ArrowRight',
    'Home',
    'End',
    'Tab',
    'Enter',
  ];

  if (allowedKeys.includes(event.key) || event.ctrlKey || event.metaKey) {
    return;
  }

  if (!/^[0-9]$/.test(event.key)) {
    event.preventDefault();
  }
}
</script>

<template>
  <Sheet :open="open" @update:open="emit('update:open', $event)">
    <SheetContent class="w-full overflow-y-auto sm:max-w-xl">
      <SheetHeader>
        <SheetTitle>Edit product</SheetTitle>

        <SheetDescription> Update catalog, pricing and inventory information. </SheetDescription>
      </SheetHeader>

      <form class="mt-6 space-y-5" @submit.prevent="onSubmit">
        <div class="space-y-2">
          <label class="text-sm font-medium text-foreground"> Name </label>

          <Input v-model="form.name" placeholder="Product name" />
        </div>

        <div class="space-y-2">
          <label class="text-sm font-medium text-foreground"> Description </label>

          <textarea
            v-model="form.description"
            rows="4"
            maxlength="500"
            placeholder="Describe the product..."
            class="flex w-full resize-none rounded-md border border-input bg-background px-3 py-2 text-sm text-foreground shadow-sm outline-none transition placeholder:text-muted-foreground focus-visible:ring-2 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50"
            :disabled="isSubmitting"
          />

          <p class="text-right text-xs text-muted-foreground">{{ form.description.length }}/500</p>
        </div>

        <div class="grid gap-4 sm:grid-cols-2">
          <div class="space-y-2">
            <label class="text-sm font-medium text-foreground"> Resource kind </label>

            <Select v-model="form.kind">
              <SelectTrigger>
                <span class="capitalize">
                  {{ form.kind }}
                </span>
              </SelectTrigger>

              <SelectContent>
                <SelectItem value="product"> Product </SelectItem>

                <SelectItem value="service"> Service </SelectItem>
              </SelectContent>
            </Select>
          </div>

          <div class="space-y-2">
            <label class="text-sm font-medium text-foreground"> Status </label>

            <Select v-model="form.status">
              <SelectTrigger>
                <span>
                  {{ form.status === '1' ? 'Active' : 'Inactive' }}
                </span>
              </SelectTrigger>

              <SelectContent>
                <SelectItem value="1"> Active </SelectItem>

                <SelectItem value="0"> Inactive </SelectItem>
              </SelectContent>
            </Select>
          </div>
        </div>

        <div class="grid gap-4 sm:grid-cols-2">
          <div class="space-y-2">
            <label class="text-sm font-medium text-foreground"> Category </label>

            <Select v-model="form.category_id" :disabled="isSubmitting || categoriesLoading">
              <SelectTrigger>
                <span class="truncate">
                  {{ categoryTriggerLabel }}
                </span>
              </SelectTrigger>

              <SelectContent>
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
          </div>

          <div class="space-y-2">
            <label class="text-sm font-medium text-foreground"> Product type </label>

            <Select
              v-model="form.type_id"
              :disabled="isSubmitting || form.category_id === 'all' || productTypesLoading"
            >
              <SelectTrigger>
                <span class="truncate">
                  {{ productTypeTriggerLabel }}
                </span>
              </SelectTrigger>

              <SelectContent>
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
        </div>

        <div
          v-if="selectedCategory && selectedProductType"
          class="rounded-xl border border-border bg-muted/30 p-4"
        >
          <p class="text-xs font-medium text-muted-foreground">Product classification</p>

          <div class="mt-2 flex flex-wrap items-center gap-2">
            <span class="rounded-full bg-primary/10 px-3 py-1 text-xs font-medium text-primary">
              {{ selectedCategory.name }}
            </span>

            <span class="text-xs text-muted-foreground"> → </span>

            <span
              class="rounded-full border border-border bg-card px-3 py-1 text-xs font-medium text-foreground"
            >
              {{ selectedProductType.name }}
            </span>
          </div>
        </div>

        <div class="grid gap-4 sm:grid-cols-2">
          <div class="space-y-2">
            <label class="text-sm font-medium text-foreground"> Price </label>

            <Input
              v-model="form.price"
              type="text"
              inputmode="decimal"
              placeholder="0.00"
              :disabled="isSubmitting"
              @keydown="validateDecimalInput"
            />
          </div>

          <div class="space-y-2">
            <label class="text-sm font-medium text-foreground"> Stock </label>

            <Input
              v-model="form.stock"
              type="text"
              inputmode="numeric"
              placeholder="0"
              :disabled="isSubmitting"
              @keydown="validateIntegerInput"
            />
          </div>
        </div>

        <div class="grid grid-cols-2 gap-2 pt-4 sm:flex sm:justify-end">
          <Button
            type="button"
            variant="outline"
            class="w-full sm:w-auto"
            :disabled="isSubmitting"
            @click="emit('update:open', false)"
          >
            Cancel
          </Button>

          <Button type="submit" class="w-full sm:w-auto" :disabled="isSubmitting">
            {{ isSubmitting ? 'Saving...' : 'Save changes' }}
          </Button>
        </div>
      </form>
    </SheetContent>
  </Sheet>
</template>
