<script setup lang="ts">
import { computed, reactive, watch } from 'vue';

import { FilePlus2, Loader2, PackagePlus } from 'lucide-vue-next';

import { toast } from 'vue-sonner';

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

import { useCategories } from '@/modules/categories/composables/useCategories';
import { useProductTypes } from '@/modules/product-types/composables/useProductTypes';

import { useCreateProduct } from '../composables/useCreateProduct';

const props = defineProps<{
  open: boolean;
}>();

const emit = defineEmits<{
  'update:open': [value: boolean];
  created: [productId: number];
}>();

const form = reactive({
  name: '',
  description: '',

  kind: 'product' as 'product' | 'service',

  category_id: 'all',
  type_id: 'all',

  price: '',
  stock: '',
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

const selectedProductType = computed(() => {
  if (form.type_id === 'all') return null;

  return (
    availableProductTypes.value.find((productType) => String(productType.id) === form.type_id) ??
    null
  );
});

const categoryTriggerLabel = computed(() => {
  if (categoriesLoading.value) {
    return 'Loading categories...';
  }

  return selectedCategory.value?.name ?? 'Select category';
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

const createMutation = useCreateProduct();

const isSubmitting = computed(() => createMutation.isPending.value);

watch(
  () => form.category_id,
  () => {
    form.type_id = 'all';
  }
);

watch(
  () => props.open,
  (open) => {
    if (open) return;

    resetForm();
  }
);

function resetForm() {
  form.name = '';
  form.description = '';

  form.kind = 'product';

  form.category_id = 'all';
  form.type_id = 'all';

  form.price = '';
  form.stock = '';
}

async function onSubmit() {
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
    const product = await createMutation.mutateAsync({
      name,
      description: description || undefined,

      kind: form.kind,

      category_id: Number(form.category_id),
      type_id: Number(form.type_id),

      price,
      stock,
    });

    toast.success('Product created');

    resetForm();

    emit('update:open', false);
    emit('created', product.id);
  } catch (error: any) {
    console.log(error);
    const message = error?.response?.data?.errorMessage ?? 'Failed to create product';

    toast.error(message);
  }
}

function validateDecimalInput(event: KeyboardEvent) {
  const allowedKeys = ['Backspace', 'Delete', 'ArrowLeft', 'ArrowRight', 'Tab', 'Enter'];

  const isNumber = /^[0-9]$/.test(event.key);
  const isDot = event.key === '.';

  if (!isNumber && !isDot && !allowedKeys.includes(event.key)) {
    event.preventDefault();
  }
}

function validateIntegerInput(event: KeyboardEvent) {
  const allowedKeys = ['Backspace', 'Delete', 'ArrowLeft', 'ArrowRight', 'Tab', 'Enter'];

  const isNumber = /^[0-9]$/.test(event.key);

  if (!isNumber && !allowedKeys.includes(event.key)) {
    event.preventDefault();
  }
}
</script>

<template>
  <Sheet :open="open" @update:open="emit('update:open', $event)">
    <SheetContent class="w-full overflow-y-auto sm:max-w-2xl">
      <SheetHeader>
        <SheetTitle class="flex items-center gap-2">
          <PackagePlus class="h-5 w-5" />
          Create product
        </SheetTitle>

        <SheetDescription> Add a product or service to your company catalog. </SheetDescription>
      </SheetHeader>

      <form class="mt-6 space-y-6" @submit.prevent="onSubmit">
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
            class="flex w-full rounded-md border border-input bg-background px-3 py-2 text-sm text-foreground shadow-sm outline-none transition placeholder:text-muted-foreground focus-visible:ring-2 focus-visible:ring-ring"
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
            <label class="text-sm font-medium text-foreground"> Category </label>

            <Select v-model="form.category_id">
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
        </div>

        <div class="space-y-2">
          <label class="text-sm font-medium text-foreground"> Product type </label>

          <Select
            v-model="form.type_id"
            :disabled="form.category_id === 'all' || productTypesLoading"
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
              @keydown="validateDecimalInput"
              v-model="form.price"
              type="text"
              inputmode="decimal"
              placeholder="0.00"
            />
          </div>

          <div class="space-y-2">
            <label class="text-sm font-medium text-foreground"> Initial stock </label>

            <Input
              @keydown="validateIntegerInput"
              v-model="form.stock"
              type="text"
              inputmode="numeric"
              placeholder="0"
            />
          </div>
        </div>

        <div class="rounded-xl border border-border bg-card p-4">
          <div class="flex items-start gap-3">
            <div class="rounded-lg border border-border bg-muted p-2 text-muted-foreground">
              <FilePlus2 class="h-4 w-4" />
            </div>

            <div>
              <p class="text-sm font-medium text-foreground">Image upload comes next</p>

              <p class="mt-1 text-xs leading-5 text-muted-foreground">
                After creating the product, you can upload its catalog image.
              </p>
            </div>
          </div>
        </div>

        <div class="grid grid-cols-2 gap-2 pt-2 sm:flex sm:justify-end">
          <Button
            type="button"
            variant="outline"
            class="w-full sm:w-auto"
            :disabled="isSubmitting"
            @click="emit('update:open', false)"
          >
            Cancel
          </Button>

          <Button type="submit" :disabled="isSubmitting" class="w-full sm:w-auto">
            <span class="mr-2 flex h-4 w-4 items-center justify-center">
              <Loader2 v-if="isSubmitting" class="h-4 w-4 animate-spin" />

              <PackagePlus v-else class="h-4 w-4" />
            </span>

            {{ isSubmitting ? 'Creating...' : 'Create product' }}
          </Button>
        </div>
      </form>
    </SheetContent>
  </Sheet>
</template>
