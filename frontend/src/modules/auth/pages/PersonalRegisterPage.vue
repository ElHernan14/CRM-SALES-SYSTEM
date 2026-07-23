<script setup lang="ts">
import { ArrowLeft, CheckCircle2, ShieldCheck, ShoppingBag } from 'lucide-vue-next';

import ThemeToggle from '@/shared/components/ThemeToggle.vue';

import { brand } from '@/shared/config/brand';

import PersonalRegisterForm from '../components/register/PersonalRegisterForm.vue';
</script>

<template>
  <main class="relative min-h-screen overflow-hidden bg-background text-foreground">
    <div
      class="pointer-events-none absolute -left-48 -top-48 h-[520px] w-[520px] rounded-full bg-primary/10 blur-3xl"
    />

    <div
      class="pointer-events-none absolute -bottom-40 right-0 h-[460px] w-[460px] rounded-full bg-blue-500/10 blur-3xl"
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
      class="relative mx-auto grid min-h-screen max-w-[1400px] items-center gap-14 px-6 py-24 lg:grid-cols-[0.9fr_1.1fr] lg:px-8"
    >
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

            <p class="text-xs text-muted-foreground">Personal commerce account</p>
          </div>
        </RouterLink>

        <div class="mt-16 max-w-lg">
          <div class="flex h-14 w-14 items-center justify-center rounded-2xl bg-primary/10">
            <ShoppingBag class="h-7 w-7 text-primary" />
          </div>

          <h1 class="mt-7 text-5xl font-semibold tracking-[-0.05em]">
            Your purchases,
            <span class="text-muted-foreground"> all connected. </span>
          </h1>

          <p class="mt-6 text-lg leading-8 text-muted-foreground">
            Create your Nexora account to build carts, complete secure multi-seller checkouts and
            manage your personal purchase history.
          </p>
        </div>

        <div class="mt-12 space-y-4">
          <div class="flex items-center gap-3">
            <CheckCircle2 class="h-5 w-5 text-primary" />

            <span class="text-sm text-muted-foreground"> Unified multi-seller cart </span>
          </div>

          <div class="flex items-center gap-3">
            <ShieldCheck class="h-5 w-5 text-primary" />

            <span class="text-sm text-muted-foreground"> Secure purchases and payments </span>
          </div>

          <div class="flex items-center gap-3">
            <ShoppingBag class="h-5 w-5 text-primary" />

            <span class="text-sm text-muted-foreground"> Complete order history </span>
          </div>
        </div>
      </section>

      <section
        class="mx-auto w-full max-w-2xl rounded-[2rem] border border-border bg-card p-6 shadow-2xl sm:p-8 lg:p-10"
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

          <p class="mt-7 text-sm font-semibold text-primary lg:mt-0">Personal account</p>

          <h2 class="mt-3 text-3xl font-semibold tracking-[-0.035em]">
            Create your Nexora account
          </h2>

          <p class="mt-3 text-sm leading-6 text-muted-foreground">
            Start shopping and keep your purchases, payments and orders in one place.
          </p>
        </div>

        <PersonalRegisterForm />

        <p class="mt-7 text-center text-sm text-muted-foreground">
          Already have an account?

          <RouterLink to="/login" class="ml-1 font-semibold text-primary hover:underline">
            Sign in
          </RouterLink>
        </p>
      </section>
    </div>
  </main>
</template>
