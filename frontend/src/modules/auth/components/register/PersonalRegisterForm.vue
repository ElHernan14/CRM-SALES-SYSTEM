<script setup lang="ts">
import { computed, reactive, ref } from 'vue';

import {
  ArrowRight,
  Eye,
  EyeOff,
  Loader2,
  LockKeyhole,
  Mail,
  Phone,
  UserRound,
} from 'lucide-vue-next';

import { toast } from 'vue-sonner';

import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

import { useRegister } from '../../composables/useRegister';

import { PersonalRegisterRequestSchema } from '../../types/register.types';

const { personalRegisterMutation } = useRegister();

const form = reactive({
  firstName: '',
  lastName: '',
  email: '',
  phone: '',
  password: '',
  confirmPassword: '',
  acceptTerms: false,
});

const errors = reactive({
  firstName: '',
  lastName: '',
  email: '',
  phone: '',
  password: '',
  confirmPassword: '',
  terms: '',
});

const showPassword = ref(false);
const showConfirmPassword = ref(false);

const isSubmitting = computed(() => {
  return personalRegisterMutation.isPending.value;
});

function clearErrors() {
  errors.firstName = '';
  errors.lastName = '';
  errors.email = '';
  errors.phone = '';
  errors.password = '';
  errors.confirmPassword = '';
  errors.terms = '';
}

function validateForm() {
  clearErrors();

  const parsed = PersonalRegisterRequestSchema.safeParse({
    first_name: form.firstName,
    last_name: form.lastName,
    email: form.email,
    password: form.password,
    phone: form.phone.trim() || undefined,
  });

  if (!parsed.success) {
    for (const issue of parsed.error.issues) {
      const field = issue.path[0];

      if (field === 'first_name') {
        errors.firstName = issue.message;
      }

      if (field === 'last_name') {
        errors.lastName = issue.message;
      }

      if (field === 'email') {
        errors.email = issue.message;
      }

      if (field === 'phone') {
        errors.phone = issue.message;
      }

      if (field === 'password') {
        errors.password = issue.message;
      }
    }
  }

  if (form.password !== form.confirmPassword) {
    errors.confirmPassword = 'Passwords do not match';
  }

  if (!form.acceptTerms) {
    errors.terms = 'You must accept the terms to continue';
  }

  return parsed.success && !errors.confirmPassword && !errors.terms;
}

async function submit() {
  if (!validateForm()) return;

  try {
    await personalRegisterMutation.mutateAsync({
      first_name: form.firstName.trim(),
      last_name: form.lastName.trim(),
      email: form.email.trim(),
      password: form.password,
      phone: form.phone.trim() || undefined,
    });
  } catch (error: any) {
    const message =
      error?.response?.data?.errorMessage ??
      error?.response?.data?.message ??
      'Unable to create your account';

    toast.error(message);
  }
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
  <form class="space-y-5" @submit.prevent="submit">
    <div class="grid gap-4 sm:grid-cols-2">
      <div class="space-y-2">
        <label for="personal-first-name" class="text-sm font-medium text-foreground">
          First name
        </label>

        <div class="relative">
          <UserRound
            class="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
          />

          <Input
            id="personal-first-name"
            v-model="form.firstName"
            class="h-11 pl-9"
            autocomplete="given-name"
            placeholder="Hernán"
            :disabled="isSubmitting"
          />
        </div>

        <p v-if="errors.firstName" class="text-xs text-destructive">
          {{ errors.firstName }}
        </p>
      </div>

      <div class="space-y-2">
        <label for="personal-last-name" class="text-sm font-medium text-foreground">
          Last name
        </label>

        <Input
          id="personal-last-name"
          v-model="form.lastName"
          class="h-11"
          autocomplete="family-name"
          placeholder="Constante"
          :disabled="isSubmitting"
        />

        <p v-if="errors.lastName" class="text-xs text-destructive">
          {{ errors.lastName }}
        </p>
      </div>
    </div>

    <div class="space-y-2">
      <label for="personal-email" class="text-sm font-medium text-foreground">
        Email address
      </label>

      <div class="relative">
        <Mail
          class="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
        />

        <Input
          id="personal-email"
          v-model="form.email"
          type="email"
          class="h-11 pl-9"
          autocomplete="email"
          placeholder="you@example.com"
          :disabled="isSubmitting"
        />
      </div>

      <p v-if="errors.email" class="text-xs text-destructive">
        {{ errors.email }}
      </p>
    </div>

    <div class="space-y-2">
      <label for="personal-phone" class="text-sm font-medium text-foreground">
        Phone
        <span class="font-normal text-muted-foreground"> (optional) </span>
      </label>

      <div class="relative">
        <Phone
          class="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
        />

        <Input
          id="personal-phone"
          v-model="form.phone"
          class="h-11 pl-9"
          autocomplete="tel"
          placeholder="+54 266..."
          :disabled="isSubmitting"
          @keydown="validateNumberInput"
        />
      </div>

      <p v-if="errors.phone" class="text-xs text-destructive">
        {{ errors.phone }}
      </p>
    </div>

    <div class="grid gap-4 sm:grid-cols-2">
      <div class="space-y-2">
        <label for="personal-password" class="text-sm font-medium text-foreground">
          Password
        </label>

        <div class="relative">
          <LockKeyhole
            class="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
          />

          <Input
            id="personal-password"
            v-model="form.password"
            :type="showPassword ? 'text' : 'password'"
            class="h-11 pl-9 pr-10"
            autocomplete="new-password"
            placeholder="Minimum 8 characters"
            :disabled="isSubmitting"
          />

          <button
            type="button"
            class="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground transition hover:text-foreground"
            :disabled="isSubmitting"
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
        <label for="personal-confirm-password" class="text-sm font-medium text-foreground">
          Confirm password
        </label>

        <div class="relative">
          <Input
            id="personal-confirm-password"
            v-model="form.confirmPassword"
            :type="showConfirmPassword ? 'text' : 'password'"
            class="h-11 pr-10"
            autocomplete="new-password"
            placeholder="Repeat password"
            :disabled="isSubmitting"
          />

          <button
            type="button"
            class="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground transition hover:text-foreground"
            :disabled="isSubmitting"
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
          I agree to the Nexora terms of service and acknowledge the privacy policy.
        </span>
      </label>

      <p v-if="errors.terms" class="mt-2 text-xs text-destructive">
        {{ errors.terms }}
      </p>
    </div>

    <Button type="submit" size="lg" class="w-full rounded-full" :disabled="isSubmitting">
      <Loader2 v-if="isSubmitting" class="mr-2 h-4 w-4 animate-spin" />

      <UserRound v-else class="mr-2 h-4 w-4" />

      {{ isSubmitting ? 'Creating your account...' : 'Create personal account' }}

      <ArrowRight v-if="!isSubmitting" class="ml-2 h-4 w-4" />
    </Button>
  </form>
</template>
