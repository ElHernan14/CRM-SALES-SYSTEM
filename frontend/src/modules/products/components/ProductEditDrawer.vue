<script setup lang="ts">
import { computed, reactive, watch } from 'vue';
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

import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';

import { useUpdateProduct } from '../composables/useUpdateProduct';
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
  type: '',
  price: '',
  stock: '',
  status: '1',
});

const updateMutation = useUpdateProduct();

const isSubmitting = computed(() => updateMutation.isPending.value);

watch(
  () => props.product,
  (product) => {
    if (!product) return;

    form.name = product.name;
    form.description = product.description;
    form.type = product.type;
    form.price = String(product.price);
    form.stock = String(product.stock);
    form.status = String(product.status);
  },
  { immediate: true }
);

async function onSubmit() {
  if (!props.product) return;

  try {
    await updateMutation.mutateAsync({
      id: props.product.id,
      payload: {
        name: form.name,
        description: form.description,
        type: form.type as 'product' | 'service',
        price: Number(form.price),
        stock: Number(form.stock),
        status: Number(form.status) as 0 | 1,
      },
    });

    toast.success('Product updated');
    emit('update:open', false);
  } catch {
    toast.error('Failed to update product');
  }
}
</script>

<template>
  <Sheet :open="open" @update:open="emit('update:open', $event)">
    <SheetContent class="w-full sm:max-w-xl">
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

          <Input v-model="form.description" placeholder="Short description" />
        </div>

        <div class="grid gap-4 sm:grid-cols-2">
          <div class="space-y-2">
            <label class="text-sm font-medium text-foreground"> Type </label>

            <Select v-model="form.type">
              <SelectTrigger>
                <SelectValue placeholder="Type" />
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
                <SelectValue placeholder="Status" />
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
            <label class="text-sm font-medium text-foreground"> Price </label>

            <Input v-model="form.price" type="text" inputmode="decimal" placeholder="0.00" />
          </div>

          <div class="space-y-2">
            <label class="text-sm font-medium text-foreground"> Stock </label>

            <Input v-model="form.stock" type="text" inputmode="numeric" placeholder="0" />
          </div>
        </div>

        <div class="flex justify-end gap-2 pt-4">
          <Button type="button" variant="outline" @click="emit('update:open', false)">
            Cancel
          </Button>

          <Button type="submit" :disabled="isSubmitting">
            {{ isSubmitting ? 'Saving...' : 'Save changes' }}
          </Button>
        </div>
      </form>
    </SheetContent>
  </Sheet>
</template>
