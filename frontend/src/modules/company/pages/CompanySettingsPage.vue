<script setup lang="ts">
import { computed } from 'vue';

import {
  BadgeCheck,
  Building2,
  CalendarDays,
  ImageIcon,
  Loader2,
  Pencil,
  RefreshCw,
  ShieldCheck,
  Sparkles,
  Store,
  Upload,
} from 'lucide-vue-next';

import { Button } from '@/components/ui/button';

import PageContainer from '@/shared/components/erp/PageContainer.vue';
import EmptyState from '@/shared/components/erp/EmptyState.vue';

import { useCompanyMe } from '@/modules/company/composables/useCompanyMe';

import { getCompanyCoverUrl, getCompanyLogoUrl } from '@/shared/utils/assets';

const { data: company, isLoading, isFetching, isError, refetch } = useCompanyMe();

const companyLogo = computed(() => {
  return getCompanyLogoUrl(company.value?.logo);
});

const companyCover = computed(() => {
  return getCompanyCoverUrl(company.value?.cover_image);
});

const isRefreshing = computed(() => {
  return isFetching.value && !isLoading.value;
});

const companyStatus = computed(() => {
  return company.value?.status === 1 ? 'Active workspace' : 'Inactive workspace';
});

const companyInitial = computed(() => {
  return company.value?.name?.trim().charAt(0).toUpperCase() ?? 'N';
});

function formatDate(value?: string | null) {
  if (!value) {
    return 'Workspace member';
  }

  return new Intl.DateTimeFormat('en', {
    month: 'long',
    year: 'numeric',
  }).format(new Date(value));
}

function requestCompanyEdit() {
  /*
   * Próximo incremento:
   * abrir CompanyProfileDrawer.
   */
  console.log('Open company editor');
}

function requestLogoUpload() {
  /*
   * Próximo incremento:
   * abrir CompanyLogoDrawer.
   */
  console.log('Open logo upload');
}

function requestCoverUpload() {
  /*
   * Próximo incremento:
   * abrir CompanyCoverDrawer.
   */
  console.log('Open cover upload');
}
</script>

<template>
  <PageContainer>
    <!-- LOADING -->
    <div v-if="isLoading" class="space-y-6">
      <div class="h-[360px] animate-pulse rounded-[2rem] bg-muted" />

      <div class="grid gap-5 lg:grid-cols-3">
        <div class="h-64 animate-pulse rounded-[1.75rem] bg-muted lg:col-span-2" />

        <div class="h-64 animate-pulse rounded-[1.75rem] bg-muted" />
      </div>
    </div>

    <!-- ERROR -->
    <div v-else-if="isError" class="py-12">
      <EmptyState
        title="Unable to load your company"
        description="There was a problem retrieving the active workspace information."
        :icon="Building2"
      />

      <div class="mt-5 flex justify-center">
        <Button variant="outline" @click="refetch()">
          <RefreshCw class="mr-2 h-4 w-4" />
          Try again
        </Button>
      </div>
    </div>

    <!-- COMPANY EXPERIENCE -->
    <div v-else-if="company" class="space-y-6">
      <!-- COMPANY HERO -->
      <section
        class="group relative min-h-[360px] overflow-hidden rounded-[2rem] border border-border bg-card shadow-xl"
      >
        <img
          v-if="companyCover"
          :src="companyCover"
          :alt="`${company.name} cover`"
          class="absolute inset-0 h-full w-full object-cover transition duration-700 group-hover:scale-[1.015]"
        />

        <div
          v-else
          class="absolute inset-0 bg-gradient-to-br from-primary/20 via-muted/50 to-background"
        />

        <div
          class="absolute inset-0 bg-gradient-to-r from-background/95 via-background/75 to-background/20"
        />

        <div
          class="absolute inset-x-0 bottom-0 h-40 bg-gradient-to-t from-background/90 to-transparent"
        />

        <div class="relative flex min-h-[360px] flex-col justify-between p-6 sm:p-8">
          <!-- HERO ACTIONS -->
          <div class="flex items-start justify-between gap-4">
            <div
              class="inline-flex items-center gap-2 rounded-full border border-border/70 bg-background/75 px-3 py-1.5 text-xs font-medium text-muted-foreground shadow-sm backdrop-blur"
            >
              <ShieldCheck class="h-3.5 w-3.5 text-primary" />
              Nexora business workspace
            </div>

            <Button
              variant="secondary"
              size="sm"
              class="rounded-full bg-background/80 shadow-md backdrop-blur"
              @click="requestCoverUpload"
            >
              <ImageIcon class="mr-2 h-4 w-4" />

              {{ companyCover ? 'Change cover' : 'Add cover' }}
            </Button>
          </div>

          <!-- COMPANY IDENTITY -->
          <div class="flex flex-col gap-6 md:flex-row md:items-end md:justify-between">
            <div class="flex flex-col gap-5 sm:flex-row sm:items-end">
              <button
                type="button"
                class="group/logo relative flex h-28 w-28 shrink-0 items-center justify-center overflow-hidden rounded-[1.75rem] border border-border bg-background shadow-xl"
                @click="requestLogoUpload"
              >
                <img
                  v-if="companyLogo"
                  :src="companyLogo"
                  :alt="`${company.name} logo`"
                  class="h-full w-full object-contain p-3"
                />

                <span v-else class="text-4xl font-semibold text-primary">
                  {{ companyInitial }}
                </span>

                <div
                  class="absolute inset-0 flex items-center justify-center bg-background/85 opacity-0 backdrop-blur transition group-hover/logo:opacity-100"
                >
                  <Upload class="h-5 w-5" />
                </div>
              </button>

              <div class="max-w-2xl pb-1">
                <div class="flex flex-wrap items-center gap-2">
                  <span
                    class="rounded-full border border-primary/20 bg-primary/10 px-3 py-1 text-xs font-semibold text-primary"
                  >
                    {{ company.category }}
                  </span>

                  <span
                    class="inline-flex items-center gap-1.5 rounded-full border border-emerald-500/20 bg-emerald-500/10 px-3 py-1 text-xs font-semibold text-emerald-600 dark:text-emerald-400"
                  >
                    <BadgeCheck class="h-3.5 w-3.5" />
                    {{ companyStatus }}
                  </span>
                </div>

                <h1
                  class="mt-4 text-3xl font-semibold tracking-[-0.04em] text-foreground sm:text-4xl"
                >
                  {{ company.name }}
                </h1>

                <p
                  class="mt-3 line-clamp-3 max-w-xl text-sm leading-6 text-muted-foreground sm:text-base"
                >
                  {{
                    company.description ||
                    'Add a description to communicate what your company offers across Nexora.'
                  }}
                </p>
              </div>
            </div>

            <div class="flex flex-wrap gap-2 md:justify-end">
              <Button
                variant="outline"
                class="rounded-full bg-background/80 backdrop-blur"
                :disabled="isRefreshing"
                @click="refetch()"
              >
                <Loader2 v-if="isRefreshing" class="mr-2 h-4 w-4 animate-spin" />

                <RefreshCw v-else class="mr-2 h-4 w-4" />

                Refresh
              </Button>

              <Button class="rounded-full" @click="requestCompanyEdit">
                <Pencil class="mr-2 h-4 w-4" />
                Edit company
              </Button>
            </div>
          </div>
        </div>
      </section>

      <!-- WORKSPACE STATUS -->
      <section class="grid gap-4 md:grid-cols-3">
        <article class="rounded-[1.5rem] border border-border bg-card p-5 shadow-sm">
          <div class="flex items-start justify-between gap-4">
            <div>
              <p class="text-sm font-semibold">Workspace ready</p>

              <p class="mt-2 text-sm leading-6 text-muted-foreground">
                Your tenant is active and connected to Nexora operations.
              </p>
            </div>

            <div class="rounded-xl bg-emerald-500/10 p-3">
              <ShieldCheck class="h-5 w-5 text-emerald-600 dark:text-emerald-400" />
            </div>
          </div>

          <div
            class="mt-5 flex items-center gap-2 text-xs font-medium text-emerald-600 dark:text-emerald-400"
          >
            <BadgeCheck class="h-4 w-4" />
            Operational
          </div>
        </article>

        <article class="rounded-[1.5rem] border border-border bg-card p-5 shadow-sm">
          <div class="flex items-start justify-between gap-4">
            <div>
              <p class="text-sm font-semibold">Commerce identity</p>

              <p class="mt-2 text-sm leading-6 text-muted-foreground">
                Your profile represents the company throughout Store and Marketplace.
              </p>
            </div>

            <div class="rounded-xl bg-primary/10 p-3">
              <Store class="h-5 w-5 text-primary" />
            </div>
          </div>

          <div class="mt-5 flex items-center gap-2 text-xs text-muted-foreground">
            <Sparkles class="h-4 w-4 text-primary" />

            {{
              companyLogo && companyCover
                ? 'Brand identity complete'
                : 'Brand assets can be improved'
            }}
          </div>
        </article>

        <article class="rounded-[1.5rem] border border-border bg-card p-5 shadow-sm">
          <div class="flex items-start justify-between gap-4">
            <div>
              <p class="text-sm font-semibold">Workspace history</p>

              <p class="mt-2 text-sm leading-6 text-muted-foreground">
                Nexora business account for {{ company.name }}.
              </p>
            </div>

            <div class="rounded-xl bg-blue-500/10 p-3">
              <CalendarDays class="h-5 w-5 text-blue-600 dark:text-blue-400" />
            </div>
          </div>

          <p class="mt-5 text-xs font-medium text-muted-foreground">
            Since {{ formatDate(company.created_at) }}
          </p>
        </article>
      </section>

      <!-- INFORMATION + BRANDING -->
      <section class="grid items-start gap-6 xl:grid-cols-[minmax(0,1.2fr)_minmax(340px,0.8fr)]">
        <!-- COMPANY INFORMATION -->
        <article class="rounded-[1.75rem] border border-border bg-card p-6 shadow-sm">
          <div class="flex flex-col gap-4 sm:flex-row sm:items-start sm:justify-between">
            <div>
              <p class="text-xs font-semibold uppercase tracking-[0.16em] text-primary">
                Business profile
              </p>

              <h2 class="mt-2 text-xl font-semibold tracking-tight">Company information</h2>

              <p class="mt-2 text-sm leading-6 text-muted-foreground">
                Core details used across your ERP workspace and commerce presence.
              </p>
            </div>

            <Button variant="outline" size="sm" class="rounded-full" @click="requestCompanyEdit">
              <Pencil class="mr-2 h-4 w-4" />
              Edit
            </Button>
          </div>

          <dl class="mt-7 divide-y divide-border">
            <div class="grid gap-2 py-4 sm:grid-cols-[180px_minmax(0,1fr)]">
              <dt class="text-sm text-muted-foreground">Company name</dt>

              <dd class="text-sm font-semibold text-foreground">
                {{ company.name }}
              </dd>
            </div>

            <div class="grid gap-2 py-4 sm:grid-cols-[180px_minmax(0,1fr)]">
              <dt class="text-sm text-muted-foreground">Business category</dt>

              <dd>
                <span
                  class="inline-flex rounded-full bg-primary/10 px-3 py-1 text-xs font-semibold text-primary"
                >
                  {{ company.category }}
                </span>
              </dd>
            </div>

            <div class="grid gap-2 py-4 sm:grid-cols-[180px_minmax(0,1fr)]">
              <dt class="text-sm text-muted-foreground">Description</dt>

              <dd class="text-sm leading-6 text-foreground">
                {{ company.description || 'No company description has been added yet.' }}
              </dd>
            </div>

            <div class="grid gap-2 py-4 sm:grid-cols-[180px_minmax(0,1fr)]">
              <dt class="text-sm text-muted-foreground">Workspace status</dt>

              <dd>
                <span
                  class="inline-flex items-center gap-2 text-sm font-medium text-emerald-600 dark:text-emerald-400"
                >
                  <span class="h-2 w-2 rounded-full bg-emerald-500" />
                  {{ companyStatus }}
                </span>
              </dd>
            </div>
          </dl>
        </article>

        <!-- BRAND PREVIEW -->
        <article class="overflow-hidden rounded-[1.75rem] border border-border bg-card shadow-sm">
          <div class="border-b border-border bg-muted/20 p-6">
            <p class="text-xs font-semibold uppercase tracking-[0.16em] text-primary">
              Brand preview
            </p>

            <h2 class="mt-2 text-xl font-semibold tracking-tight">Nexora presence</h2>

            <p class="mt-2 text-sm leading-6 text-muted-foreground">
              Preview how your business identity appears throughout the platform.
            </p>
          </div>

          <div class="p-6">
            <div class="overflow-hidden rounded-2xl border border-border bg-background shadow-sm">
              <div
                class="relative h-28 overflow-hidden bg-gradient-to-br from-primary/20 via-muted to-background"
              >
                <img
                  v-if="companyCover"
                  :src="companyCover"
                  :alt="`${company.name} cover preview`"
                  class="h-full w-full object-cover"
                />

                <div class="absolute inset-0 bg-gradient-to-t from-background/80 to-transparent" />
              </div>

              <div class="relative px-5 pb-5">
                <div
                  class="-mt-9 flex h-18 w-18 items-center justify-center overflow-hidden rounded-2xl border-4 border-background bg-background shadow-md"
                >
                  <img
                    v-if="companyLogo"
                    :src="companyLogo"
                    :alt="`${company.name} logo preview`"
                    class="h-full w-full object-contain p-2"
                  />

                  <span v-else class="text-xl font-semibold text-primary">
                    {{ companyInitial }}
                  </span>
                </div>

                <p class="mt-4 text-base font-semibold">
                  {{ company.name }}
                </p>

                <p class="mt-1 text-xs text-muted-foreground">
                  {{ company.category }}
                </p>

                <p class="mt-3 line-clamp-2 text-xs leading-5 text-muted-foreground">
                  {{ company.description || 'Your company profile across Nexora.' }}
                </p>
              </div>
            </div>

            <div class="mt-5 grid gap-3 sm:grid-cols-2 xl:grid-cols-1">
              <Button variant="outline" class="justify-start rounded-xl" @click="requestLogoUpload">
                <Upload class="mr-3 h-4 w-4" />
                {{ companyLogo ? 'Replace company logo' : 'Upload company logo' }}
              </Button>

              <Button
                variant="outline"
                class="justify-start rounded-xl"
                @click="requestCoverUpload"
              >
                <ImageIcon class="mr-3 h-4 w-4" />
                {{ companyCover ? 'Replace cover image' : 'Upload cover image' }}
              </Button>
            </div>
          </div>
        </article>
      </section>
    </div>
  </PageContainer>
</template>
