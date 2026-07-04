<script setup lang="ts">
import { Search } from 'lucide-vue-next';

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

const search = defineModel<string>('search', { required: true });
const type = defineModel<string>('type', { required: true });
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
</script>

<template>
  <div class="space-y-4">
    <div class="grid gap-3 md:grid-cols-[1.5fr_1fr_1fr_1fr]">
      <div class="relative">
        <Search
          class="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
        />

        <Input v-model="search" class="pl-9" placeholder="Search products..." />
      </div>

      <Select v-model="type">
        <SelectTrigger>
          <SelectValue placeholder="Product type" />
        </SelectTrigger>

        <SelectContent>
          <SelectItem value="product"> Product </SelectItem>

          <SelectItem value="service"> Service </SelectItem>
        </SelectContent>
      </Select>

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
    </div>

    <div class="flex items-center justify-between">
      <p class="text-sm text-muted-foreground">{{ total }} results</p>

      <Button variant="ghost" size="sm" @click="emit('clear')"> Clear filters </Button>
    </div>
  </div>
</template>
