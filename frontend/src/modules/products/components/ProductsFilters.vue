<script setup lang="ts">
import { Search } from 'lucide-vue-next';

import { computed, watch } from 'vue';

import { useCategories } from '@/modules/categories/composables/useCategories';
import { useProductTypes } from '@/modules/product-types/composables/useProductTypes';

import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';

import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';

defineProps<{
  total: number;
}>();

const kind = defineModel<string>('kind', { required: true });
const categoryId = defineModel<string>('categoryId', { required: true });
const typeId = defineModel<string>('typeId', { required: true });
const search = defineModel<string>('search', { required: true });
const minPrice = defineModel<string>('minPrice', { required: true });
const maxPrice = defineModel<string>('maxPrice', { required: true });

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

// Emits
const emit = defineEmits<{
  clear: [];
}>();

// Filters data
const { data: categoriesData, isLoading: categoriesLoading } = useCategories();

const productCategories = computed(() => {
  return categoriesData.value?.product_categories ?? [];
});

const selectedCategoryId = computed<number | null>(() => {
  return categoryId.value !== 'all' ? Number(categoryId.value) : null;
});

const { data: productTypesData, isLoading: productTypesLoading } =
  useProductTypes(selectedCategoryId);

const availableProductTypes = computed(() => {
  return productTypesData.value?.items ?? [];
});

watch(categoryId, () => {
  typeId.value = 'all';
});

const selectedCategory = computed(() => {
  if (categoryId.value === 'all') return null;

  return (
    productCategories.value.find((category) => String(category.id) === categoryId.value) ?? null
  );
});

const categoryTriggerLabel = computed(() => {
  if (categoriesLoading.value) return 'Loading categories...';
  if (categoryId.value === 'all') return 'All categories';

  return selectedCategory.value?.name ?? 'Product category';
});

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

  return selectedProductType.value?.name ?? 'Product type';
});
</script>

<template>
  <div class="space-y-4">
    <div class="grid gap-3 md:grid-cols-2 xl:grid-cols-[1.5fr_180px_220px_220px]">
      <div class="relative">
        <Search
          class="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
        />

        <Input v-model="search" class="pl-9" placeholder="Search products..." />
      </div>

      <Select v-model="kind">
        <SelectTrigger>
          <span
            class="truncate"
            :class="kind === 'all' ? 'text-muted-foreground' : 'text-foreground'"
          >
            {{ kind === 'all' ? 'All kinds' : kind === 'product' ? 'Product' : 'Service' }}
          </span>
        </SelectTrigger>

        <SelectContent>
          <SelectItem value="all"> All kinds </SelectItem>

          <SelectItem value="product"> Product </SelectItem>

          <SelectItem value="service"> Service </SelectItem>
        </SelectContent>
      </Select>

      <Select v-model="categoryId">
        <SelectTrigger>
          <span
            class="truncate"
            :class="categoryId === 'all' ? 'text-muted-foreground' : 'text-foreground'"
          >
            {{ categoryTriggerLabel }}
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

    <div class="grid gap-3 md:grid-cols-[1fr_1fr_auto]">
      <Input
        v-model="minPrice"
        type="text"
        inputmode="decimal"
        placeholder="Min price"
        @keydown="validateNumberInput"
      />

      <Input
        v-model="maxPrice"
        type="text"
        inputmode="decimal"
        placeholder="Max price"
        @keydown="validateNumberInput"
      />

      <Button variant="ghost" size="sm" @click="emit('clear')"> Clear filters </Button>
    </div>

    <p class="text-sm text-muted-foreground">{{ total }} results</p>
  </div>
</template>
