<script setup lang="ts">
import { useRouter } from 'vue-router';

import { ArrowRight, Building2, ShoppingBag } from 'lucide-vue-next';

import { brand } from '@/shared/config/brand';

const router = useRouter();
</script>

<template>
  <main
    class="relative flex min-h-screen items-center justify-center overflow-hidden bg-background px-6 py-16"
  >
    <div
      class="pointer-events-none absolute -left-40 -top-40 h-96 w-96 rounded-full bg-primary/10 blur-3xl"
    />

    <div class="relative w-full max-w-4xl">
      <div class="text-center">
        <RouterLink to="/" class="inline-flex items-center gap-2">
          <div
            class="flex h-10 w-10 items-center justify-center rounded-xl bg-primary font-bold text-primary-foreground"
          >
            N
          </div>

          <span class="text-xl font-semibold">
            {{ brand.name }}
          </span>
        </RouterLink>

        <h1 class="mt-8 text-4xl font-semibold tracking-tight text-foreground">
          How will you use Nexora?
        </h1>

        <p class="mt-3 text-base text-muted-foreground">
          Choose the account experience that best matches your goals.
        </p>
      </div>

      <div class="mt-10 grid gap-5 md:grid-cols-2">
        <button
          type="button"
          class="group rounded-[1.75rem] border border-border bg-card p-7 text-left shadow-sm transition hover:-translate-y-1 hover:border-primary/30 hover:shadow-xl"
          @click="router.push('/register/personal')"
        >
          <div class="flex h-12 w-12 items-center justify-center rounded-2xl bg-primary/10">
            <ShoppingBag class="h-6 w-6 text-primary" />
          </div>

          <h2 class="mt-6 text-xl font-semibold text-foreground">Shop as an individual</h2>

          <p class="mt-3 text-sm leading-6 text-muted-foreground">
            Discover products, complete purchases and manage your personal order history.
          </p>

          <div class="mt-7 flex items-center text-sm font-semibold text-primary">
            Create personal account

            <ArrowRight class="ml-2 h-4 w-4 transition-transform group-hover:translate-x-1" />
          </div>
        </button>

        <button
          type="button"
          class="group rounded-[1.75rem] border border-border bg-card p-7 text-left shadow-sm transition hover:-translate-y-1 hover:border-primary/30 hover:shadow-xl"
          @click="router.push('/register/business')"
        >
          <div class="flex h-12 w-12 items-center justify-center rounded-2xl bg-blue-500/10">
            <Building2 class="h-6 w-6 text-blue-600 dark:text-blue-400" />
          </div>

          <h2 class="mt-6 text-xl font-semibold text-foreground">Create a business workspace</h2>

          <p class="mt-3 text-sm leading-6 text-muted-foreground">
            Manage products, sales, purchases, payments and company operations through the Nexora
            ERP.
          </p>

          <div class="mt-7 flex items-center text-sm font-semibold text-primary">
            Create business workspace

            <ArrowRight class="ml-2 h-4 w-4 transition-transform group-hover:translate-x-1" />
          </div>
        </button>
      </div>

      <p class="mt-8 text-center text-sm text-muted-foreground">
        Already have an account?

        <RouterLink to="/login" class="ml-1 font-semibold text-primary hover:underline">
          Sign in
        </RouterLink>
      </p>

      <p class="mt-3 text-center text-xs text-muted-foreground">
        Joining an existing company?

        <span class="ml-1"> Use the invitation link provided by its administrator. </span>
      </p>
    </div>
  </main>
</template>
