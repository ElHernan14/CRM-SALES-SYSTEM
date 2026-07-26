<script setup lang="ts">
import { computed, reactive, ref, watch } from 'vue';

import {
  BadgeCheck,
  Building2,
  CheckCircle2,
  Loader2,
  LockKeyhole,
  Mail,
  Pencil,
  Phone,
  RefreshCw,
  Save,
  ShoppingBag,
  UserRound,
  X,
} from 'lucide-vue-next';

import { toast } from 'vue-sonner';

import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

import EmptyState from '@/shared/components/erp/EmptyState.vue';

import { useClientProfile } from '../composables/useClientProfile';
import { useUpdateClientProfile } from '../composables/useUpdateClientProfile';

import { UpdateClientProfileRequestSchema } from '../types/account.types';

const props = withDefaults(
  defineProps<{
    embedded?: boolean;
    surface?: 'store' | 'erp';
  }>(),
  {
    embedded: false,
    surface: 'store',
  }
);

const pageEyebrow = computed(() => {
  return props.surface === 'erp' ? 'Nexora workspace' : 'Nexora account';
});

const pageTitle = computed(() => {
  return props.surface === 'erp'
    ? 'Your personal profile'
    : 'Your profile, purchases and identity.';
});

const pageDescription = computed(() => {
  return props.surface === 'erp'
    ? 'Manage the personal identity connected to your business workspace.'
    : 'Manage the personal information connected to your Nexora account and commerce activity.';
});

const { data: profile, isLoading, isFetching, isError, refetch } = useClientProfile();

const updateMutation = useUpdateClientProfile();

const editing = ref(false);

const form = reactive({
  firstName: '',
  lastName: '',
  email: '',
  phone: '',
});

const errors = reactive({
  firstName: '',
  lastName: '',
  email: '',
  phone: '',
});

const fullName = computed(() => {
  if (!profile.value) return 'Nexora customer';

  return [profile.value.first_name, profile.value.last_name].filter(Boolean).join(' ');
});

const initials = computed(() => {
  const first = profile.value?.first_name?.charAt(0).toUpperCase() ?? '';

  const last = profile.value?.last_name?.charAt(0).toUpperCase() ?? '';

  return `${first}${last}` || 'N';
});

const isBusinessProfile = computed(() => {
  return Boolean(profile.value?.company_id);
});

const accountType = computed(() => {
  return isBusinessProfile.value ? 'Business member' : 'Personal customer';
});

const isRefreshing = computed(() => {
  return isFetching.value && !isLoading.value;
});

const isSaving = computed(() => {
  return updateMutation.isPending.value;
});

watch(
  profile,
  (value) => {
    if (!value) return;

    form.firstName = value.first_name;
    form.lastName = value.last_name;
    form.email = value.email;
    form.phone = value.phone ?? '';
  },
  {
    immediate: true,
  }
);

function clearErrors() {
  errors.firstName = '';
  errors.lastName = '';
  errors.email = '';
  errors.phone = '';
}

function startEditing() {
  if (!profile.value) return;

  form.firstName = profile.value.first_name;
  form.lastName = profile.value.last_name;
  form.email = profile.value.email;
  form.phone = profile.value.phone ?? '';

  clearErrors();

  editing.value = true;
}

function cancelEditing() {
  editing.value = false;

  if (!profile.value) return;

  form.firstName = profile.value.first_name;
  form.lastName = profile.value.last_name;
  form.email = profile.value.email;
  form.phone = profile.value.phone ?? '';

  clearErrors();
}

function validate() {
  clearErrors();

  const parsed = UpdateClientProfileRequestSchema.safeParse({
    first_name: form.firstName.trim(),
    last_name: form.lastName.trim(),
    email: form.email.trim(),
    phone: form.phone.trim() || undefined,
  });

  if (parsed.success) {
    return parsed.data;
  }

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
  }

  return null;
}

async function saveProfile() {
  if (!profile.value) return;

  const payload = validate();

  if (!payload) return;

  try {
    await updateMutation.mutateAsync({
      clientId: profile.value.id,
      payload,
    });

    toast.success('Profile updated', {
      description: 'Your personal information was saved successfully.',
    });

    editing.value = false;
  } catch (error: any) {
    const message =
      error?.response?.data?.errorMessage ??
      error?.response?.data?.message ??
      'Unable to update your profile';

    toast.error(message);
  }
}

const contentContainerClass = computed(() => {
  return props.embedded
    ? 'space-y-6'
    : 'mx-auto w-full max-w-[1300px] space-y-6 px-4 py-8 sm:px-6 sm:py-10 lg:px-8';
});
</script>

<template>
  <div class="pb-20">
    <!-- HERO -->
    <section
      v-if="!embedded"
      class="relative overflow-hidden border-b border-border bg-gradient-to-br from-background via-background to-primary/5"
    >
      <div
        class="pointer-events-none absolute -right-48 -top-56 h-[520px] w-[520px] rounded-full bg-primary/10 blur-3xl"
      />

      <div class="relative mx-auto w-full max-w-[1300px] px-4 py-12 sm:px-6 sm:py-14 lg:px-8">
        <div class="flex flex-col justify-between gap-7 lg:flex-row lg:items-end">
          <div class="max-w-3xl">
            <p class="text-sm font-semibold text-primary">
              {{ pageEyebrow }}
            </p>

            <h1 class="mt-3 text-4xl font-semibold leading-[1.05] tracking-[-0.045em] sm:text-5xl">
              {{ pageTitle }}
            </h1>

            <p class="mt-4 max-w-2xl text-base leading-7 text-muted-foreground">
              {{ pageDescription }}
            </p>
          </div>

          <Button
            variant="outline"
            class="w-full self-start rounded-full sm:w-auto lg:self-auto"
            :disabled="isRefreshing"
            @click="refetch()"
          >
            <Loader2 v-if="isRefreshing" class="mr-2 h-4 w-4 animate-spin" />

            <RefreshCw v-else class="mr-2 h-4 w-4" />

            Refresh profile
          </Button>
        </div>
      </div>
    </section>

    <!-- LOADING -->
    <main
      v-if="isLoading"
      :class="
        embedded
          ? 'space-y-6'
          : 'mx-auto w-full max-w-[1300px] space-y-6 px-4 py-8 sm:px-6 sm:py-10 lg:px-8'
      "
    >
      <div class="h-80 animate-pulse rounded-[1.5rem] bg-muted sm:h-72 sm:rounded-[2rem]" />

      <div class="grid gap-6 lg:grid-cols-3">
        <div
          class="h-96 animate-pulse rounded-[1.5rem] bg-muted sm:h-72 sm:rounded-[1.75rem] lg:col-span-2"
        />

        <div class="h-72 animate-pulse rounded-[1.5rem] bg-muted sm:rounded-[1.75rem]" />
      </div>
    </main>

    <!-- ERROR -->
    <main v-else-if="isError" class="mx-auto max-w-3xl px-4 py-16 sm:px-6 sm:py-20">
      <EmptyState
        title="Unable to load your profile"
        description="There was a problem retrieving your Nexora account."
        :icon="UserRound"
      />

      <div class="mt-5 flex justify-center">
        <Button variant="outline" class="w-full rounded-full sm:w-auto" @click="refetch()">
          Try again
        </Button>
      </div>
    </main>

    <!-- PROFILE -->
    <main v-else-if="profile" :class="contentContainerClass">
      <!-- IDENTITY -->
      <section
        class="relative overflow-hidden rounded-[1.5rem] border border-border bg-card p-5 shadow-xl sm:rounded-[2rem] sm:p-8"
      >
        <div
          class="pointer-events-none absolute right-0 top-0 h-64 w-64 rounded-full bg-primary/10 blur-3xl"
        />

        <div class="relative flex flex-col gap-7 md:flex-row md:items-center md:justify-between">
          <div class="flex min-w-0 flex-col gap-5 sm:flex-row sm:items-center">
            <div
              class="flex h-20 w-20 shrink-0 items-center justify-center rounded-[1.5rem] bg-primary text-2xl font-semibold text-primary-foreground shadow-lg shadow-primary/15 sm:h-24 sm:w-24 sm:rounded-[1.75rem] sm:text-3xl"
            >
              {{ initials }}
            </div>

            <div class="min-w-0">
              <div class="flex flex-wrap items-center gap-2">
                <span
                  class="inline-flex items-center gap-1.5 rounded-full border border-primary/20 bg-primary/10 px-3 py-1 text-xs font-semibold text-primary"
                >
                  <BadgeCheck class="h-3.5 w-3.5" />

                  {{ accountType }}
                </span>

                <span
                  class="inline-flex items-center gap-1.5 rounded-full border border-emerald-500/20 bg-emerald-500/10 px-3 py-1 text-xs font-semibold text-emerald-600 dark:text-emerald-400"
                >
                  <CheckCircle2 class="h-3.5 w-3.5" />

                  Active account
                </span>
              </div>

              <h2 class="mt-4 break-words text-2xl font-semibold tracking-[-0.04em] sm:text-3xl">
                {{ fullName }}
              </h2>

              <p class="mt-2 break-all text-sm text-muted-foreground">
                {{ profile.email }}
              </p>

              <p
                v-if="profile.company_name"
                class="mt-2 flex min-w-0 items-center gap-2 text-sm font-medium"
              >
                <Building2 class="h-4 w-4 shrink-0 text-primary" />

                <span class="truncate">
                  {{ profile.company_name }}
                </span>
              </p>
            </div>
          </div>

          <Button
            v-if="!editing"
            class="w-full self-start rounded-full sm:w-auto md:self-auto"
            @click="startEditing"
          >
            <Pencil class="mr-2 h-4 w-4" />

            Edit profile
          </Button>

          <div v-else class="grid w-full gap-2 min-[390px]:grid-cols-2 md:w-auto">
            <Button
              variant="outline"
              class="w-full rounded-full"
              :disabled="isSaving"
              @click="cancelEditing"
            >
              <X class="mr-2 h-4 w-4" />
              Cancel
            </Button>

            <Button class="w-full rounded-full" :disabled="isSaving" @click="saveProfile">
              <Loader2 v-if="isSaving" class="mr-2 h-4 w-4 animate-spin" />

              <Save v-else class="mr-2 h-4 w-4" />

              {{ isSaving ? 'Saving...' : 'Save changes' }}
            </Button>
          </div>
        </div>
      </section>

      <section class="grid items-start gap-6 lg:grid-cols-[minmax(0,1.25fr)_minmax(320px,0.75fr)]">
        <!-- PROFILE INFO -->
        <article
          class="min-w-0 rounded-[1.5rem] border border-border bg-card p-5 shadow-sm sm:rounded-[1.75rem] sm:p-6"
        >
          <p class="text-xs font-semibold uppercase tracking-[0.16em] text-primary">
            Personal identity
          </p>

          <h2 class="mt-2 text-xl font-semibold">Contact information</h2>

          <p class="mt-2 text-sm leading-6 text-muted-foreground">
            Information used across Store purchases and your Nexora account.
          </p>

          <!-- VIEW -->
          <dl v-if="!editing" class="mt-7 divide-y divide-border">
            <div
              class="grid gap-2 py-4 sm:grid-cols-[150px_minmax(0,1fr)] lg:grid-cols-[170px_minmax(0,1fr)]"
            >
              <dt class="text-sm text-muted-foreground">First name</dt>

              <dd class="break-words text-sm font-semibold">
                {{ profile.first_name }}
              </dd>
            </div>

            <div
              class="grid gap-2 py-4 sm:grid-cols-[150px_minmax(0,1fr)] lg:grid-cols-[170px_minmax(0,1fr)]"
            >
              <dt class="text-sm text-muted-foreground">Last name</dt>

              <dd class="break-words text-sm font-semibold">
                {{ profile.last_name }}
              </dd>
            </div>

            <div
              class="grid gap-2 py-4 sm:grid-cols-[150px_minmax(0,1fr)] lg:grid-cols-[170px_minmax(0,1fr)]"
            >
              <dt class="text-sm text-muted-foreground">Email address</dt>

              <dd class="flex min-w-0 items-start gap-2 text-sm font-medium">
                <Mail class="mt-0.5 h-4 w-4 shrink-0 text-muted-foreground" />

                <span class="min-w-0 break-all">
                  {{ profile.email }}
                </span>
              </dd>
            </div>

            <div
              class="grid gap-2 py-4 sm:grid-cols-[150px_minmax(0,1fr)] lg:grid-cols-[170px_minmax(0,1fr)]"
            >
              <dt class="text-sm text-muted-foreground">Phone</dt>

              <dd class="flex min-w-0 items-start gap-2 text-sm font-medium">
                <Phone class="mt-0.5 h-4 w-4 shrink-0 text-muted-foreground" />

                <span class="break-words">
                  {{ profile.phone || 'No phone number added' }}
                </span>
              </dd>
            </div>
          </dl>

          <!-- EDIT -->
          <form v-else class="mt-7 space-y-5" @submit.prevent="saveProfile">
            <div class="grid gap-4 sm:grid-cols-2">
              <div class="space-y-2">
                <label for="account-first-name" class="text-sm font-medium"> First name </label>

                <Input
                  id="account-first-name"
                  v-model="form.firstName"
                  class="h-11"
                  :disabled="isSaving"
                />

                <p v-if="errors.firstName" class="text-xs text-destructive">
                  {{ errors.firstName }}
                </p>
              </div>

              <div class="space-y-2">
                <label for="account-last-name" class="text-sm font-medium"> Last name </label>

                <Input
                  id="account-last-name"
                  v-model="form.lastName"
                  class="h-11"
                  :disabled="isSaving"
                />

                <p v-if="errors.lastName" class="text-xs text-destructive">
                  {{ errors.lastName }}
                </p>
              </div>
            </div>

            <div class="space-y-2">
              <label for="account-email" class="text-sm font-medium"> Email address </label>

              <div class="relative">
                <Mail
                  class="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
                />

                <Input
                  id="account-email"
                  v-model="form.email"
                  type="email"
                  class="h-11 pl-9"
                  :disabled="isSaving"
                />
              </div>

              <p v-if="errors.email" class="text-xs text-destructive">
                {{ errors.email }}
              </p>

              <p class="text-xs leading-5 text-muted-foreground">
                Changing your email may affect how you access Nexora.
              </p>
            </div>

            <div class="space-y-2">
              <label for="account-phone" class="text-sm font-medium"> Phone </label>

              <div class="relative">
                <Phone
                  class="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
                />

                <Input
                  id="account-phone"
                  v-model="form.phone"
                  class="h-11 pl-9"
                  placeholder="+54 266..."
                  :disabled="isSaving"
                />
              </div>

              <p v-if="errors.phone" class="text-xs text-destructive">
                {{ errors.phone }}
              </p>
            </div>
          </form>
        </article>

        <!-- CONTEXT -->
        <aside class="min-w-0 space-y-4">
          <article
            v-if="surface === 'store'"
            class="rounded-[1.5rem] border border-border bg-card p-5 shadow-sm sm:rounded-[1.75rem] sm:p-6"
          >
            <div class="flex h-11 w-11 items-center justify-center rounded-xl bg-primary/10">
              <ShoppingBag class="h-5 w-5 text-primary" />
            </div>

            <h2 class="mt-5 text-lg font-semibold">Commerce account</h2>

            <p class="mt-2 text-sm leading-6 text-muted-foreground">
              Your identity connects carts, purchases, payments and seller orders throughout Nexora.
            </p>

            <RouterLink
              to="/store/purchases"
              class="mt-5 inline-flex text-sm font-semibold text-primary hover:underline"
            >
              View purchase history
            </RouterLink>
          </article>

          <article
            v-if="surface === 'erp'"
            class="rounded-[1.5rem] border border-border bg-card p-5 shadow-sm sm:rounded-[1.75rem] sm:p-6"
          >
            <div class="flex h-11 w-11 items-center justify-center rounded-xl bg-primary/10">
              <Building2 class="h-5 w-5 text-primary" />
            </div>

            <h2 class="mt-5 text-lg font-semibold">Workspace identity</h2>

            <p class="mt-2 text-sm leading-6 text-muted-foreground">
              Your personal profile identifies you within the active business workspace, permissions
              and tenant operations.
            </p>

            <RouterLink
              :to="{ name: 'erp-settings-company' }"
              class="mt-5 inline-flex text-sm font-semibold text-primary hover:underline"
            >
              View company workspace
            </RouterLink>
          </article>

          <article
            v-if="isBusinessProfile"
            class="rounded-[1.5rem] border border-border bg-card p-5 shadow-sm sm:rounded-[1.75rem] sm:p-6"
          >
            <div class="flex h-11 w-11 items-center justify-center rounded-xl bg-blue-500/10">
              <Building2 class="h-5 w-5 text-blue-600 dark:text-blue-400" />
            </div>

            <h2 class="mt-5 text-lg font-semibold">Business membership</h2>

            <p class="mt-2 break-words text-sm leading-6 text-muted-foreground">
              This account is connected to

              <strong class="font-semibold text-foreground">
                {{ profile.company_name }} </strong
              >.
            </p>

            <RouterLink
              to="/erp/dashboard"
              class="mt-5 inline-flex text-sm font-semibold text-primary hover:underline"
            >
              Open ERP workspace
            </RouterLink>
          </article>

          <article
            class="rounded-[1.5rem] border border-border bg-muted/20 p-5 sm:rounded-[1.75rem]"
          >
            <div class="flex items-start gap-3">
              <LockKeyhole class="mt-0.5 h-4 w-4 shrink-0 text-primary" />

              <div class="min-w-0">
                <p class="text-sm font-medium">Account security</p>

                <p class="mt-1 text-xs leading-5 text-muted-foreground">
                  Password changes and email verification will be handled through a dedicated
                  security flow.
                </p>
              </div>
            </div>
          </article>
        </aside>
      </section>
    </main>
  </div>
</template>
