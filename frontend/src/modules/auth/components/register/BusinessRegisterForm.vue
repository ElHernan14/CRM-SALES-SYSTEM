<script setup lang="ts">
import { computed, reactive, ref, watch, onUnmounted } from 'vue';

import {
  ArrowLeft,
  ArrowRight,
  Building2,
  Check,
  Eye,
  EyeOff,
  Loader2,
  LockKeyhole,
  Mail,
  Phone,
  UserRound,
  ImagePlus,
  Upload,
} from 'lucide-vue-next';

import { toast } from 'vue-sonner';

import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

import { Select, SelectContent, SelectItem, SelectTrigger } from '@/components/ui/select';

import { useCategories } from '@/modules/categories/composables/useCategories';

import { useRegister } from '../../composables/useRegister';

import { BusinessRegisterRequestSchema } from '../../types/register.types';

import { z } from 'zod';

import { useRouter } from 'vue-router';

import { useUploadCompanyLogo } from '@/modules/company/composables/useUploadCompanyLogo';

import { getCompanyLogoUrl } from '@/shared/utils/assets';

const router = useRouter();

const { businessRegisterMutation, establishSession } = useRegister();

const uploadLogoMutation = useUploadCompanyLogo();

const currentStep = ref<1 | 2 | 3>(1);

const companyLogoFile = ref<File | null>(null);

const companyLogoPreview = ref<string | null>(null);

const logoInput = ref<HTMLInputElement | null>(null);

const isUploadingLogo = computed(() => {
  return uploadLogoMutation.isPending.value;
});

const isFinishingSetup = computed(() => {
  return isSubmitting.value || isUploadingLogo.value;
});

const zEmail = z.string().email();

const showPassword = ref(false);
const showConfirmPassword = ref(false);

function selectLogoFile(event: Event) {
  const input = event.target as HTMLInputElement;

  const file = input.files?.[0] ?? null;

  if (!file) return;

  if (!file.type.startsWith('image/')) {
    toast.error('Select a valid image file');
    return;
  }

  const maxSize = 5 * 1024 * 1024;

  if (file.size > maxSize) {
    toast.error('The logo cannot exceed 5 MB');

    return;
  }

  if (companyLogoPreview.value) {
    URL.revokeObjectURL(companyLogoPreview.value);
  }

  companyLogoFile.value = file;

  companyLogoPreview.value = URL.createObjectURL(file);
}

function openLogoPicker() {
  logoInput.value?.click();
}

function removeLogo() {
  if (companyLogoPreview.value) {
    URL.revokeObjectURL(companyLogoPreview.value);
  }

  companyLogoFile.value = null;
  companyLogoPreview.value = null;

  if (logoInput.value) {
    logoInput.value.value = '';
  }
}

onUnmounted(() => {
  if (companyLogoPreview.value) {
    URL.revokeObjectURL(companyLogoPreview.value);
  }
});

const form = reactive({
  firstName: '',
  lastName: '',
  email: '',
  phone: '',
  password: '',
  confirmPassword: '',

  companyName: '',
  companyEmail: '',
  categoryId: 'all',
  description: '',

  acceptTerms: false,
});

const errors = reactive({
  firstName: '',
  lastName: '',
  email: '',
  phone: '',
  password: '',
  confirmPassword: '',

  companyName: '',
  companyEmail: '',
  categoryId: '',
  description: '',

  terms: '',
});

const {
  data: categoriesData,
  isLoading: categoriesLoading,
  isError: categoriesError,
} = useCategories();

const companyCategories = computed(() => {
  return categoriesData.value?.company_categories ?? [];
});

const selectedCompanyCategory = computed(() => {
  if (form.categoryId === 'all') return null;

  return (
    companyCategories.value.find((category) => String(category.id) === form.categoryId) ?? null
  );
});

const isSubmitting = computed(() => {
  return businessRegisterMutation.isPending.value;
});

const canGoBack = computed(() => {
  return currentStep.value === 2 && !isSubmitting.value;
});

const descriptionLength = computed(() => {
  return form.description.length;
});

watch(
  () => form.email,
  (email) => {
    /*
     * Ayuda visual:
     * inicialmente usamos el mismo email como contacto.
     * Luego el usuario puede modificarlo.
     */
    if (!form.companyEmail) {
      form.companyEmail = email;
    }
  }
);

function clearAccountErrors() {
  errors.firstName = '';
  errors.lastName = '';
  errors.email = '';
  errors.phone = '';
  errors.password = '';
  errors.confirmPassword = '';
}

function clearCompanyErrors() {
  errors.companyName = '';
  errors.companyEmail = '';
  errors.categoryId = '';
  errors.description = '';
  errors.terms = '';
}

function validateAccountStep() {
  clearAccountErrors();

  let valid = true;

  if (form.firstName.trim().length < 2) {
    errors.firstName = 'First name must contain at least 2 characters';

    valid = false;
  }

  if (form.lastName.trim().length < 2) {
    errors.lastName = 'Last name must contain at least 2 characters';

    valid = false;
  }

  const emailResult = zEmail.safeParse(form.email.trim());

  if (!emailResult.success) {
    errors.email = 'Enter a valid work email address';

    valid = false;
  }

  if (form.phone.trim().length > 50) {
    errors.phone = 'Phone number is too long';

    valid = false;
  }

  if (form.password.length < 8) {
    errors.password = 'Password must contain at least 8 characters';

    valid = false;
  }

  if (form.password !== form.confirmPassword) {
    errors.confirmPassword = 'Passwords do not match';

    valid = false;
  }

  return valid;
}

function validateCompanyStep() {
  clearCompanyErrors();

  let valid = true;

  if (form.companyName.trim().length < 2) {
    errors.companyName = 'Company name must contain at least 2 characters';

    valid = false;
  }

  const emailResult = zEmail.safeParse(form.companyEmail.trim());

  if (!emailResult.success) {
    errors.companyEmail = 'Enter a valid company email';

    valid = false;
  }

  if (form.categoryId === 'all') {
    errors.categoryId = 'Select a company category';

    valid = false;
  }

  if (form.description.length > 1000) {
    errors.description = 'Description cannot exceed 1000 characters';

    valid = false;
  }

  if (!form.acceptTerms) {
    errors.terms = 'You must accept the terms to continue';

    valid = false;
  }

  return valid;
}

function goToCompanyStep() {
  if (!validateAccountStep()) return;

  currentStep.value = 2;
}

function goToAccountStep() {
  if (!canGoBack.value) return;

  currentStep.value = 1;
}

async function submit() {
  if (!validateAccountStep()) {
    currentStep.value = 1;
    return;
  }

  if (!validateCompanyStep()) {
    currentStep.value = 2;
    return;
  }

  const payload = {
    first_name: form.firstName.trim(),
    last_name: form.lastName.trim(),
    email: form.email.trim(),
    password: form.password,

    phone: form.phone.trim() || undefined,

    company: {
      name: form.companyName.trim(),
      email: form.companyEmail.trim(),

      category_id: Number(form.categoryId),

      description: form.description.trim() || undefined,
    },
  };

  const parsed = BusinessRegisterRequestSchema.safeParse(payload);

  if (!parsed.success) {
    toast.error(parsed.error.issues[0]?.message ?? 'Review the registration fields');

    return;
  }

  try {
    const response = await businessRegisterMutation.mutateAsync(parsed.data);

    await establishSession(response);

    toast.success('Business workspace created', {
      description: 'Add your company logo or continue to the ERP.',
    });

    currentStep.value = 3;
  } catch (error: any) {
    const message =
      error?.response?.data?.errorMessage ??
      error?.response?.data?.message ??
      'Unable to create the business workspace';

    toast.error(message);
  }
}

async function finishBusinessSetup() {
  try {
    if (companyLogoFile.value) {
      await uploadLogoMutation.mutateAsync(companyLogoFile.value);

      toast.success('Company logo uploaded');
    }

    await router.replace('/erp/dashboard');
  } catch (error: any) {
    const message =
      error?.response?.data?.errorMessage ??
      error?.response?.data?.message ??
      'Unable to upload the company logo';

    toast.error(message);
  }
}

async function skipCompanyLogo() {
  await router.replace('/erp/dashboard');
}

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
</script>

<template>
  <form class="space-y-7" @submit.prevent="currentStep === 1 ? goToCompanyStep() : submit()">
    <!-- PROGRESS -->
    <div class="grid grid-cols-3 gap-3">
      <div
        v-for="step in [
          {
            id: 1,
            label: 'Your account',
          },
          {
            id: 2,
            label: 'Business',
          },
          {
            id: 3,
            label: 'Branding',
          },
        ]"
        :key="step.id"
        class="space-y-2"
      >
        <div class="flex items-center gap-2">
          <div
            class="flex h-8 w-8 shrink-0 items-center justify-center rounded-full text-xs font-semibold"
            :class="
              currentStep >= step.id
                ? 'bg-primary text-primary-foreground'
                : 'bg-muted text-muted-foreground'
            "
          >
            <Check v-if="currentStep > step.id" class="h-4 w-4" />

            <span v-else>
              {{ step.id }}
            </span>
          </div>

          <span
            class="hidden truncate text-xs font-medium sm:block"
            :class="currentStep >= step.id ? 'text-foreground' : 'text-muted-foreground'"
          >
            {{ step.label }}
          </span>
        </div>

        <div class="h-1 overflow-hidden rounded-full bg-muted">
          <div
            class="h-full bg-primary transition-all"
            :class="currentStep >= step.id ? 'w-full' : 'w-0'"
          />
        </div>
      </div>
    </div>

    <!-- STEP 1 -->
    <div v-if="currentStep === 1" class="space-y-5">
      <div>
        <p class="text-sm font-semibold text-primary">Step 1</p>

        <h3 class="mt-2 text-xl font-semibold">Create the administrator account</h3>

        <p class="mt-2 text-sm leading-6 text-muted-foreground">
          This user will become the first member of the new business workspace.
        </p>
      </div>

      <div class="grid gap-4 sm:grid-cols-2">
        <div class="space-y-2">
          <label for="business-first-name" class="text-sm font-medium"> First name </label>

          <div class="relative">
            <UserRound
              class="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
            />

            <Input
              :disabled="isSubmitting"
              id="business-first-name"
              v-model="form.firstName"
              class="h-11 pl-9"
              autocomplete="given-name"
              placeholder="Hernán"
            />
          </div>

          <p v-if="errors.firstName" class="text-xs text-destructive">
            {{ errors.firstName }}
          </p>
        </div>

        <div class="space-y-2">
          <label for="business-last-name" class="text-sm font-medium"> Last name </label>

          <Input
            :disabled="isSubmitting"
            id="business-last-name"
            v-model="form.lastName"
            class="h-11"
            autocomplete="family-name"
            placeholder="Constante"
          />

          <p v-if="errors.lastName" class="text-xs text-destructive">
            {{ errors.lastName }}
          </p>
        </div>
      </div>

      <div class="space-y-2">
        <label for="business-user-email" class="text-sm font-medium"> Work email </label>

        <div class="relative">
          <Mail
            class="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
          />

          <Input
            :disabled="isSubmitting"
            id="business-user-email"
            v-model="form.email"
            type="email"
            class="h-11 pl-9"
            autocomplete="email"
            placeholder="you@company.com"
          />
        </div>

        <p v-if="errors.email" class="text-xs text-destructive">
          {{ errors.email }}
        </p>
      </div>

      <div class="space-y-2">
        <label for="business-phone" class="text-sm font-medium">
          Phone
          <span class="font-normal text-muted-foreground"> (optional) </span>
        </label>

        <div class="relative">
          <Phone
            class="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
          />

          <Input
            @keydown="validateNumberInput"
            :disabled="isSubmitting"
            id="business-phone"
            v-model="form.phone"
            class="h-11 pl-9"
            autocomplete="tel"
            placeholder="+54 266..."
          />
        </div>

        <p v-if="errors.phone" class="text-xs text-destructive">
          {{ errors.phone }}
        </p>
      </div>

      <div class="grid gap-4 sm:grid-cols-2">
        <div class="space-y-2">
          <label for="business-password" class="text-sm font-medium"> Password </label>

          <div class="relative">
            <LockKeyhole
              class="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
            />

            <Input
              :disabled="isSubmitting"
              id="business-password"
              v-model="form.password"
              :type="showPassword ? 'text' : 'password'"
              class="h-11 pl-9 pr-10"
              autocomplete="new-password"
              placeholder="Minimum 8 characters"
            />

            <button
              type="button"
              class="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground transition hover:text-foreground"
              @click="showPassword = !showPassword"
            >
              <EyeOff v-if="showPassword" class="h-4 w-4" />

              <Eye v-else class="h-4 w-4" />
            </button>
          </div>

          <p v-if="errors.password" class="text-xs text-destructive">
            {{ errors.password }}
          </p>
        </div>

        <div class="space-y-2">
          <label for="business-confirm-password" class="text-sm font-medium">
            Confirm password
          </label>

          <div class="relative">
            <Input
              :disabled="isSubmitting"
              id="business-confirm-password"
              v-model="form.confirmPassword"
              :type="showConfirmPassword ? 'text' : 'password'"
              class="h-11 pr-10"
              autocomplete="new-password"
              placeholder="Repeat password"
            />

            <button
              type="button"
              class="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground transition hover:text-foreground"
              @click="showConfirmPassword = !showConfirmPassword"
            >
              <EyeOff v-if="showConfirmPassword" class="h-4 w-4" />

              <Eye v-else class="h-4 w-4" />
            </button>
          </div>

          <p v-if="errors.confirmPassword" class="text-xs text-destructive">
            {{ errors.confirmPassword }}
          </p>
        </div>
      </div>

      <Button type="submit" size="lg" class="w-full rounded-full">
        Continue to business details

        <ArrowRight class="ml-2 h-4 w-4" />
      </Button>
    </div>

    <!-- STEP 2 -->
    <div v-else-if="currentStep === 2" class="space-y-5">
      <div>
        <p class="text-sm font-semibold text-primary">Step 2</p>

        <h3 class="mt-2 text-xl font-semibold">Configure your workspace</h3>

        <p class="mt-2 text-sm leading-6 text-muted-foreground">
          Nexora will create the company, administrator membership and ERP workspace in one
          transaction.
        </p>
      </div>

      <div class="space-y-2">
        <label for="business-company-name" class="text-sm font-medium"> Company name </label>

        <div class="relative">
          <Building2
            class="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
          />

          <Input
            id="business-company-name"
            v-model="form.companyName"
            class="h-11 pl-9"
            autocomplete="organization"
            placeholder="Constante Labs"
            :disabled="isSubmitting"
          />
        </div>

        <p v-if="errors.companyName" class="text-xs text-destructive">
          {{ errors.companyName }}
        </p>
      </div>

      <div class="space-y-2">
        <label for="business-company-email" class="text-sm font-medium">
          Company contact email
        </label>

        <div class="relative">
          <Mail
            class="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
          />

          <Input
            id="business-company-email"
            v-model="form.companyEmail"
            type="email"
            class="h-11 pl-9"
            autocomplete="email"
            placeholder="contact@company.com"
            :disabled="isSubmitting"
          />
        </div>

        <p class="text-xs text-muted-foreground">
          Currently used as the account and client contact. It is already prepared to become the
          company contact email.
        </p>

        <p v-if="errors.companyEmail" class="text-xs text-destructive">
          {{ errors.companyEmail }}
        </p>
      </div>

      <div class="space-y-2">
        <label class="text-sm font-medium"> Company category </label>

        <Select v-model="form.categoryId" :disabled="categoriesLoading || isSubmitting">
          <SelectTrigger class="h-11">
            <span
              class="truncate"
              :class="form.categoryId === 'all' ? 'text-muted-foreground' : 'text-foreground'"
            >
              {{
                categoriesLoading
                  ? 'Loading categories...'
                  : (selectedCompanyCategory?.name ?? 'Select company category')
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

        <p v-if="categoriesError" class="text-xs text-destructive">
          Unable to load company categories.
        </p>

        <p v-if="errors.categoryId" class="text-xs text-destructive">
          {{ errors.categoryId }}
        </p>
      </div>

      <div class="space-y-2">
        <div class="flex items-center justify-between gap-3">
          <label for="business-description" class="text-sm font-medium">
            Company description
            <span class="font-normal text-muted-foreground"> (optional) </span>
          </label>

          <span class="text-xs text-muted-foreground"> {{ descriptionLength }}/1000 </span>
        </div>

        <textarea
          id="business-description"
          v-model="form.description"
          rows="4"
          maxlength="1000"
          class="flex w-full rounded-xl border border-input bg-background px-3 py-3 text-sm shadow-sm outline-none transition placeholder:text-muted-foreground focus-visible:ring-2 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50"
          placeholder="Describe what your company offers..."
          :disabled="isSubmitting"
        />

        <p v-if="errors.description" class="text-xs text-destructive">
          {{ errors.description }}
        </p>
      </div>

      <!-- SUMMARY -->
      <div class="rounded-2xl border border-primary/20 bg-primary/5 p-4">
        <p class="text-sm font-semibold">Workspace summary</p>

        <div class="mt-3 grid gap-3 text-xs sm:grid-cols-2">
          <div>
            <p class="text-muted-foreground">Administrator</p>

            <p class="mt-1 font-medium">
              {{ form.firstName }}
              {{ form.lastName }}
            </p>
          </div>

          <div>
            <p class="text-muted-foreground">Work email</p>

            <p class="mt-1 truncate font-medium">
              {{ form.email }}
            </p>
          </div>

          <div>
            <p class="text-muted-foreground">Company</p>

            <p class="mt-1 font-medium">
              {{ form.companyName || 'Not entered' }}
            </p>
          </div>

          <div>
            <p class="text-muted-foreground">Category</p>

            <p class="mt-1 font-medium">
              {{ selectedCompanyCategory?.name ?? 'Not selected' }}
            </p>
          </div>
        </div>
      </div>

      <div>
        <label
          class="flex cursor-pointer items-start gap-3 rounded-xl border border-border bg-muted/20 p-4"
        >
          <input
            v-model="form.acceptTerms"
            type="checkbox"
            class="mt-0.5 h-4 w-4 rounded border-border"
            :disabled="isSubmitting"
          />

          <span class="text-xs leading-5 text-muted-foreground">
            I agree to the Nexora terms and confirm that I am authorized to create this business
            workspace.
          </span>
        </label>

        <p v-if="errors.terms" class="mt-2 text-xs text-destructive">
          {{ errors.terms }}
        </p>
      </div>

      <div class="grid gap-3 sm:grid-cols-[auto_1fr]">
        <Button
          type="button"
          size="lg"
          variant="outline"
          class="rounded-full"
          :disabled="!canGoBack"
          @click="goToAccountStep"
        >
          <ArrowLeft class="mr-2 h-4 w-4" />
          Back
        </Button>

        <Button
          type="submit"
          size="lg"
          class="w-full rounded-full"
          :disabled="isSubmitting || categoriesLoading"
        >
          <Loader2 v-if="isSubmitting" class="mr-2 h-4 w-4 animate-spin" />

          <Building2 v-else class="mr-2 h-4 w-4" />

          {{ isSubmitting ? 'Creating workspace...' : 'Create business workspace' }}

          <ArrowRight v-if="!isSubmitting" class="ml-2 h-4 w-4" />
        </Button>
      </div>
    </div>

    <!-- STEP 3 -->
    <div v-else-if="currentStep === 3" class="space-y-6">
      <div>
        <p class="text-sm font-semibold text-primary">Final step</p>

        <h3 class="mt-2 text-xl font-semibold">Personalize your workspace</h3>

        <p class="mt-2 text-sm leading-6 text-muted-foreground">
          Add a company logo to identify your business across the ERP, Marketplace and Nexora Store.
        </p>
      </div>

      <div class="rounded-[1.75rem] border border-border bg-muted/15 p-6">
        <input
          ref="logoInput"
          type="file"
          accept="image/*"
          class="hidden"
          @change="selectLogoFile"
        />

        <div class="flex flex-col items-center text-center">
          <button
            type="button"
            class="group relative flex h-32 w-32 items-center justify-center overflow-hidden rounded-[1.75rem] border border-dashed border-border bg-background shadow-sm transition hover:border-primary/40"
            :disabled="isFinishingSetup"
            @click="openLogoPicker"
          >
            <img
              v-if="companyLogoPreview"
              :src="companyLogoPreview"
              :alt="`${form.companyName} logo preview`"
              class="h-full w-full object-contain p-3"
            />

            <div v-else class="flex flex-col items-center gap-2 text-muted-foreground">
              <ImagePlus class="h-7 w-7" />

              <span class="text-xs font-medium"> Add logo </span>
            </div>

            <div
              v-if="companyLogoPreview"
              class="absolute inset-0 flex items-center justify-center bg-background/80 opacity-0 backdrop-blur transition group-hover:opacity-100"
            >
              <Upload class="h-5 w-5 text-foreground" />
            </div>
          </button>

          <h4 class="mt-5 text-base font-semibold">
            {{ form.companyName }}
          </h4>

          <p class="mt-1 text-sm text-muted-foreground">
            {{ selectedCompanyCategory?.name ?? 'Nexora business' }}
          </p>

          <p class="mt-4 max-w-sm text-xs leading-5 text-muted-foreground">
            PNG, JPG or WebP. Maximum recommended size: 5 MB. A square image works best.
          </p>

          <div class="mt-5 flex flex-wrap justify-center gap-2">
            <Button
              type="button"
              variant="outline"
              :disabled="isFinishingSetup"
              @click="openLogoPicker"
            >
              <Upload class="mr-2 h-4 w-4" />

              {{ companyLogoFile ? 'Choose another image' : 'Select image' }}
            </Button>

            <Button
              v-if="companyLogoFile"
              type="button"
              variant="ghost"
              :disabled="isFinishingSetup"
              @click="removeLogo"
            >
              Remove
            </Button>
          </div>
        </div>
      </div>

      <div class="rounded-xl border border-primary/20 bg-primary/5 p-4">
        <p class="text-sm font-medium">Your workspace is already active</p>

        <p class="mt-1 text-xs leading-5 text-muted-foreground">
          The logo is optional and can be changed later from Company Settings.
        </p>
      </div>

      <div class="grid gap-3 sm:grid-cols-[auto_1fr]">
        <Button
          type="button"
          size="lg"
          variant="outline"
          class="rounded-full"
          :disabled="isFinishingSetup"
          @click="skipCompanyLogo"
        >
          Skip for now
        </Button>

        <Button
          type="button"
          size="lg"
          class="w-full rounded-full"
          :disabled="isFinishingSetup"
          @click="finishBusinessSetup"
        >
          <Loader2 v-if="isFinishingSetup" class="mr-2 h-4 w-4 animate-spin" />

          <Building2 v-else class="mr-2 h-4 w-4" />

          {{ isUploadingLogo ? 'Uploading logo...' : 'Open business workspace' }}

          <ArrowRight v-if="!isFinishingSetup" class="ml-2 h-4 w-4" />
        </Button>
      </div>
    </div>
  </form>
</template>
