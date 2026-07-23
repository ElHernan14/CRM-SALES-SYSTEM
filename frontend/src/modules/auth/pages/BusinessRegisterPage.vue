<script setup lang="ts">
import { ArrowLeft, Building2, CheckCircle2, Layers3, ShieldCheck } from 'lucide-vue-next';

import ThemeToggle from '@/shared/components/ThemeToggle.vue';

import { brand } from '@/shared/config/brand';

import BusinessRegisterForm from '../components/register/BusinessRegisterForm.vue';
</script>

<template>
  <main class="relative min-h-screen overflow-hidden bg-background text-foreground">
    <div
      class="pointer-events-none absolute -left-52 -top-52 h-[560px] w-[560px] rounded-full bg-primary/10 blur-3xl"
    />

    <div
      class="pointer-events-none absolute -bottom-52 right-0 h-[520px] w-[520px] rounded-full bg-blue-500/10 blur-3xl"
    />

    <div class="absolute inset-x-0 top-0 z-30 flex items-center justify-between px-5 py-5 sm:px-8">
      <RouterLink
        to="/register"
        class="inline-flex items-center gap-2 text-sm font-medium text-muted-foreground transition hover:text-foreground"
      >
        <ArrowLeft class="h-4 w-4" />
        Account options
      </RouterLink>

      <ThemeToggle />
    </div>

    <div
      class="relative mx-auto grid min-h-screen max-w-[1500px] items-center gap-14 px-6 py-24 lg:grid-cols-[0.82fr_1.18fr] lg:px-8"
    >
      <!-- BUSINESS VALUE -->
      <section class="hidden lg:block">
        <RouterLink to="/" class="inline-flex items-center gap-3">
          <div
            class="flex h-11 w-11 items-center justify-center rounded-2xl bg-primary font-bold text-primary-foreground shadow-lg"
          >
            N
          </div>

          <div>
            <p class="text-xl font-semibold">
              {{ brand.name }}
            </p>

            <p class="text-xs text-muted-foreground">Business workspace</p>
          </div>
        </RouterLink>

        <div class="mt-16 max-w-lg">
          <div class="flex h-14 w-14 items-center justify-center rounded-2xl bg-primary/10">
            <Building2 class="h-7 w-7 text-primary" />
          </div>

          <h1 class="mt-7 text-5xl font-semibold tracking-[-0.05em]">
            Build your business
            <span class="text-muted-foreground"> on one operating system. </span>
          </h1>

          <p class="mt-6 text-lg leading-8 text-muted-foreground">
            Create a multi-tenant workspace connected to products, sales, purchases, inventory,
            payments and both B2B and B2C commerce.
          </p>
        </div>

        <div class="mt-12 space-y-4">
          <div class="flex items-center gap-3">
            <CheckCircle2 class="h-5 w-5 text-primary" />

            <span class="text-sm text-muted-foreground">
              Company and administrator created atomically
            </span>
          </div>

          <div class="flex items-center gap-3">
            <ShieldCheck class="h-5 w-5 text-primary" />

            <span class="text-sm text-muted-foreground"> Tenant ownership and permissions </span>
          </div>

          <div class="flex items-center gap-3">
            <Layers3 class="h-5 w-5 text-primary" />

            <span class="text-sm text-muted-foreground">
              ERP, Marketplace and Store integration
            </span>
          </div>
        </div>
      </section>

      <!-- FORM -->
      <section
        class="mx-auto w-full max-w-3xl rounded-[2rem] border border-border bg-card p-6 shadow-2xl sm:p-8 lg:p-10"
      >
        <div class="mb-8">
          <div class="flex items-center gap-3 lg:hidden">
            <div
              class="flex h-10 w-10 items-center justify-center rounded-xl bg-primary font-bold text-primary-foreground"
            >
              N
            </div>

            <span class="text-xl font-semibold">
              {{ brand.name }}
            </span>
          </div>

          <p class="mt-7 text-sm font-semibold text-primary lg:mt-0">Business account</p>

          <h2 class="mt-3 text-3xl font-semibold tracking-[-0.035em]">
            Create your Nexora workspace
          </h2>

          <p class="mt-3 text-sm leading-6 text-muted-foreground">
            Set up the administrator and company information. Your ERP workspace will be available
            immediately.
          </p>
        </div>

        <BusinessRegisterForm />

        <p class="mt-7 text-center text-sm text-muted-foreground">
          Already have a business account?

          <RouterLink to="/login" class="ml-1 font-semibold text-primary hover:underline">
            Sign in
          </RouterLink>
        </p>
      </section>
    </div>
  </main>
</template>
