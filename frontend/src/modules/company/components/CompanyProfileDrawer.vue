<script setup lang="ts">
import { computed, reactive, watch } from 'vue';

import { Building2, Loader2, Save } from 'lucide-vue-next';

import { toast } from 'vue-sonner';

import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

import { Select, SelectContent, SelectItem, SelectTrigger } from '@/components/ui/select';

import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
} from '@/components/ui/sheet';

import { useCategories } from '@/modules/categories/composables/useCategories';

import { useUpdateCompanyMe } from '../composables/useUpdateCompanyMe';

import { UpdateCompanyMeRequestSchema, type CompanyMe } from '../types/company.types';

const props = defineProps<{
  open: boolean;
  company: CompanyMe | null;
}>();

const emit = defineEmits<{
  'update:open': [value: boolean];
}>();

const form = reactive({
  name: '',
  categoryId: '',
  description: '',
});

const errors = reactive({
  name: '',
  categoryId: '',
  description: '',
});

const updateMutation = useUpdateCompanyMe();

const { data: categoriesData, isLoading: categoriesLoading } = useCategories();

const companyCategories = computed(() => {
  return categoriesData.value?.company_categories ?? [];
});

const selectedCategory = computed(() => {
  return (
    companyCategories.value.find((category) => String(category.id) === form.categoryId) ?? null
  );
});

const isSubmitting = computed(() => {
  return updateMutation.isPending.value;
});

const descriptionLength = computed(() => {
  return form.description.length;
});

watch(
  () => [props.open, props.company],
  () => {
    if (!props.open || !props.company) return;

    form.name = props.company.name;
    form.categoryId = String(props.company.category_id);

    form.description = props.company.description ?? '';

    clearErrors();
  },
  {
    immediate: true,
  }
);

function clearErrors() {
  errors.name = '';
  errors.categoryId = '';
  errors.description = '';
}

function validate() {
  clearErrors();

  const parsed = UpdateCompanyMeRequestSchema.safeParse({
    name: form.name.trim(),

    category_id: Number(form.categoryId),

    description: form.description.trim(),
  });

  if (parsed.success) {
    return parsed.data;
  }

  for (const issue of parsed.error.issues) {
    const field = issue.path[0];

    if (field === 'name') {
      errors.name = issue.message;
    }

    if (field === 'category_id') {
      errors.categoryId = issue.message;
    }

    if (field === 'description') {
      errors.description = issue.message;
    }
  }

  return null;
}

async function submit() {
  const payload = validate();

  if (!payload) return;

  try {
    await updateMutation.mutateAsync(payload);

    toast.success('Company profile updated', {
      description: 'Your changes are now visible across Nexora.',
    });

    emit('update:open', false);
  } catch (error: any) {
    const message =
      error?.response?.data?.errorMessage ??
      error?.response?.data?.message ??
      'Unable to update company information';

    toast.error(message);
  }
}
</script>

<template>
  <Sheet :open="open" @update:open="emit('update:open', $event)">
    <SheetContent class="w-full overflow-y-auto sm:max-w-xl">
      <SheetHeader class="text-left">
        <div class="flex items-center gap-3">
          <div class="flex h-11 w-11 items-center justify-center rounded-xl bg-primary/10">
            <Building2 class="h-5 w-5 text-primary" />
          </div>

          <div>
            <SheetTitle> Edit company profile </SheetTitle>

            <SheetDescription class="mt-1">
              Update how your business appears across Nexora.
            </SheetDescription>
          </div>
        </div>
      </SheetHeader>

      <form class="mt-8 space-y-6" @submit.prevent="submit">
        <div class="space-y-2">
          <label for="company-name" class="text-sm font-medium"> Company name </label>

          <Input
            id="company-name"
            v-model="form.name"
            class="h-11"
            placeholder="Company name"
            :disabled="isSubmitting"
          />

          <p v-if="errors.name" class="text-xs text-destructive">
            {{ errors.name }}
          </p>
        </div>

        <div class="space-y-2">
          <label class="text-sm font-medium"> Business category </label>

          <Select v-model="form.categoryId" :disabled="categoriesLoading || isSubmitting">
            <SelectTrigger class="h-11">
              <span
                class="truncate"
                :class="selectedCategory ? 'text-foreground' : 'text-muted-foreground'"
              >
                {{
                  categoriesLoading
                    ? 'Loading categories...'
                    : (selectedCategory?.name ?? 'Select category')
                }}
              </span>
            </SelectTrigger>

            <SelectContent>
              <SelectItem
                v-for="category in companyCategories"
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

          <p v-if="errors.categoryId" class="text-xs text-destructive">
            {{ errors.categoryId }}
          </p>
        </div>

        <div class="space-y-2">
          <div class="flex items-center justify-between">
            <label for="company-description" class="text-sm font-medium"> Description </label>

            <span class="text-xs text-muted-foreground"> {{ descriptionLength }}/1000 </span>
          </div>

          <textarea
            id="company-description"
            v-model="form.description"
            rows="7"
            maxlength="1000"
            class="flex w-full resize-none rounded-xl border border-input bg-background px-3 py-3 text-sm shadow-sm outline-none transition placeholder:text-muted-foreground focus-visible:ring-2 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50"
            placeholder="Describe your business, products and services..."
            :disabled="isSubmitting"
          />

          <p v-if="errors.description" class="text-xs text-destructive">
            {{ errors.description }}
          </p>
        </div>

        <div class="rounded-2xl border border-primary/20 bg-primary/5 p-4">
          <p class="text-sm font-medium">Platform-wide identity</p>

          <p class="mt-1 text-xs leading-5 text-muted-foreground">
            These changes will update your ERP workspace, Marketplace supplier profile and Store
            presence.
          </p>
        </div>

        <Button type="submit" size="lg" class="w-full rounded-full" :disabled="isSubmitting">
          <Loader2 v-if="isSubmitting" class="mr-2 h-4 w-4 animate-spin" />

          <Save v-else class="mr-2 h-4 w-4" />

          {{ isSubmitting ? 'Saving company...' : 'Save company profile' }}
        </Button>
      </form>
    </SheetContent>
  </Sheet>
</template>
