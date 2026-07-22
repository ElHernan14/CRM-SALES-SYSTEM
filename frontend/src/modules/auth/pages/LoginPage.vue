<script setup lang="ts">
import {
  ArrowLeft,
  Building2,
  CheckCircle2,
  Layers3,
  ShoppingBag,
  Sparkles,
} from 'lucide-vue-next';

import LoginForm from '../components/LoginForm.vue';

import ThemeToggle from '@/shared/components/ThemeToggle.vue';

import { brand } from '@/shared/config/brand';
</script>

<template>
  <main class="relative min-h-screen overflow-hidden bg-background text-foreground">
    <!-- DECORATION -->
    <div
      class="pointer-events-none absolute -left-52 -top-52 h-[540px] w-[540px] rounded-full bg-primary/10 blur-3xl"
    />

    <div
      class="pointer-events-none absolute -bottom-52 right-0 h-[480px] w-[480px] rounded-full bg-blue-500/10 blur-3xl"
    />

    <!-- TOP NAVIGATION -->
    <div class="absolute inset-x-0 top-0 z-30 flex items-center justify-between px-5 py-5 sm:px-8">
      <RouterLink
        to="/"
        class="inline-flex items-center gap-2 text-sm font-medium text-muted-foreground transition hover:text-foreground"
      >
        <ArrowLeft class="h-4 w-4" />
        Back to Nexora
      </RouterLink>

      <ThemeToggle />
    </div>

    <div class="relative grid min-h-screen lg:grid-cols-[1.05fr_0.95fr]">
      <!-- LEFT PANEL -->
      <section
        class="relative hidden min-h-screen overflow-hidden border-r border-border/70 lg:flex lg:flex-col lg:justify-between lg:p-14 xl:p-20"
      >
        <div
          class="absolute inset-0 bg-gradient-to-br from-primary/10 via-background to-blue-500/5"
        />

        <div
          class="pointer-events-none absolute right-12 top-32 h-80 w-80 rounded-full border border-primary/10"
        />

        <div
          class="pointer-events-none absolute right-24 top-44 h-56 w-56 rounded-full border border-primary/15"
        />

        <div class="relative">
          <RouterLink to="/" class="inline-flex items-center gap-3">
            <div
              class="flex h-11 w-11 items-center justify-center rounded-2xl bg-primary font-bold text-primary-foreground shadow-lg shadow-primary/15"
            >
              N
            </div>

            <div>
              <p class="text-xl font-semibold tracking-[-0.03em]">
                {{ brand.name }}
              </p>

              <p class="text-xs text-muted-foreground">
                {{ brand.shortDescription }}
              </p>
            </div>
          </RouterLink>

          <div class="mt-20 max-w-xl">
            <div
              class="inline-flex items-center gap-2 rounded-full border border-primary/20 bg-primary/10 px-4 py-2 text-xs font-semibold text-primary"
            >
              <Sparkles class="h-4 w-4" />
              Connected commerce and operations
            </div>

            <h1 class="mt-7 text-5xl font-semibold tracking-[-0.05em] text-foreground xl:text-6xl">
              One account.
              <span class="text-muted-foreground"> Every side of commerce. </span>
            </h1>

            <p class="mt-6 max-w-lg text-lg leading-8 text-muted-foreground">
              Access your personal purchases or enter the business workspace where your company
              manages products, sales, payments and operations.
            </p>
          </div>

          <div class="mt-12 grid max-w-xl gap-3 sm:grid-cols-2">
            <div
              class="rounded-2xl border border-border bg-background/70 p-4 shadow-sm backdrop-blur"
            >
              <div class="flex items-center gap-3">
                <div class="rounded-xl bg-primary/10 p-2.5">
                  <ShoppingBag class="h-5 w-5 text-primary" />
                </div>

                <div>
                  <p class="text-sm font-semibold">Personal commerce</p>

                  <p class="mt-1 text-xs text-muted-foreground">Store, checkout and purchases</p>
                </div>
              </div>
            </div>

            <div
              class="rounded-2xl border border-border bg-background/70 p-4 shadow-sm backdrop-blur"
            >
              <div class="flex items-center gap-3">
                <div class="rounded-xl bg-blue-500/10 p-2.5">
                  <Building2 class="h-5 w-5 text-blue-600 dark:text-blue-400" />
                </div>

                <div>
                  <p class="text-sm font-semibold">Business workspace</p>

                  <p class="mt-1 text-xs text-muted-foreground">ERP, B2B and operations</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="relative flex flex-wrap gap-x-6 gap-y-3">
          <div class="flex items-center gap-2 text-xs text-muted-foreground">
            <CheckCircle2 class="h-4 w-4 text-primary" />
            Multi-tenant security
          </div>

          <div class="flex items-center gap-2 text-xs text-muted-foreground">
            <Layers3 class="h-4 w-4 text-primary" />
            Unified commerce
          </div>

          <div class="flex items-center gap-2 text-xs text-muted-foreground">
            <Building2 class="h-4 w-4 text-primary" />
            Business-ready workflows
          </div>
        </div>
      </section>

      <!-- LOGIN PANEL -->
      <section class="flex min-h-screen items-center justify-center px-5 py-24 sm:px-8 lg:px-12">
        <div class="w-full max-w-md">
          <!-- MOBILE BRAND -->
          <RouterLink to="/" class="mb-10 inline-flex items-center gap-3 lg:hidden">
            <div
              class="flex h-10 w-10 items-center justify-center rounded-xl bg-primary font-bold text-primary-foreground"
            >
              N
            </div>

            <span class="text-xl font-semibold">
              {{ brand.name }}
            </span>
          </RouterLink>

          <div class="mb-8">
            <p class="text-sm font-semibold text-primary">Welcome back</p>

            <h2 class="mt-3 text-3xl font-semibold tracking-[-0.035em] text-foreground">
              Sign in to Nexora
            </h2>

            <p class="mt-3 text-sm leading-6 text-muted-foreground">
              Access your commerce account or continue managing your business workspace.
            </p>
          </div>

          <LoginForm />

          <div class="mt-7 border-t border-border pt-6 text-center">
            <p class="text-sm text-muted-foreground">
              New to Nexora?

              <RouterLink
                to="/register"
                class="ml-1 font-semibold text-primary transition hover:underline"
              >
                Create an account
              </RouterLink>
            </p>
          </div>
        </div>
      </section>
    </div>
  </main>
</template>
