<script setup lang="ts">
import { useRouter } from 'vue-router';

import { ArrowUpRight, Building2, ShoppingBag, Store, UserRound } from 'lucide-vue-next';

import { brand } from '@/shared/config/brand';

const router = useRouter();

function scrollToSection(sectionId: string) {
  document.getElementById(sectionId)?.scrollIntoView({
    behavior: 'smooth',
    block: 'start',
  });
}
</script>

<template>
  <footer class="border-t border-border bg-muted/10">
    <div class="mx-auto max-w-[1500px] px-6 py-14 lg:px-8">
      <div class="grid gap-12 lg:grid-cols-[1.3fr_0.7fr_0.7fr_0.7fr]">
        <!-- BRAND -->
        <div class="max-w-md">
          <RouterLink to="/" class="inline-flex items-center gap-3">
            <div
              class="flex h-11 w-11 items-center justify-center rounded-xl bg-primary font-bold text-primary-foreground"
            >
              N
            </div>

            <div>
              <p class="text-lg font-semibold tracking-[-0.03em]">
                {{ brand.name }}
              </p>

              <p class="mt-1 text-xs text-muted-foreground">
                {{ brand.shortDescription }}
              </p>
            </div>
          </RouterLink>

          <p class="mt-6 text-sm leading-7 text-muted-foreground">
            A connected platform for business operations, commercial discovery, invoices, inventory
            and payments.
          </p>

          <button
            type="button"
            class="mt-6 inline-flex items-center gap-2 text-sm font-semibold text-primary transition hover:gap-3"
            @click="router.push('/register')"
          >
            Start with Nexora

            <ArrowUpRight class="h-4 w-4" />
          </button>
        </div>

        <!-- PRODUCT -->
        <div>
          <p class="text-sm font-semibold">Product</p>

          <nav class="mt-5 space-y-3">
            <button
              type="button"
              class="block text-sm text-muted-foreground transition hover:text-foreground"
              @click="scrollToSection('nexora-product')"
            >
              Foundation
            </button>

            <button
              type="button"
              class="block text-sm text-muted-foreground transition hover:text-foreground"
              @click="scrollToSection('nexora-solutions')"
            >
              Business loop
            </button>

            <button
              type="button"
              class="block text-sm text-muted-foreground transition hover:text-foreground"
              @click="scrollToSection('nexora-commerce')"
            >
              Connected commerce
            </button>

            <button
              type="button"
              class="block text-sm text-muted-foreground transition hover:text-foreground"
              @click="scrollToSection('nexora-architecture')"
            >
              Architecture
            </button>
          </nav>
        </div>

        <!-- EXPERIENCES -->
        <div>
          <p class="text-sm font-semibold">Experiences</p>

          <nav class="mt-5 space-y-3">
            <button
              type="button"
              class="flex items-center gap-2 text-sm text-muted-foreground transition hover:text-foreground"
              @click="router.push('/store')"
            >
              <Store class="h-4 w-4" />
              Store
            </button>

            <button
              type="button"
              class="flex items-center gap-2 text-sm text-muted-foreground transition hover:text-foreground"
              @click="router.push('/store/businesses')"
            >
              <Building2 class="h-4 w-4" />
              Businesses
            </button>

            <button
              type="button"
              class="flex items-center gap-2 text-sm text-muted-foreground transition hover:text-foreground"
              @click="router.push('/store/catalog')"
            >
              <ShoppingBag class="h-4 w-4" />
              Catalog
            </button>

            <button
              type="button"
              class="flex items-center gap-2 text-sm text-muted-foreground transition hover:text-foreground"
              @click="router.push('/store/account')"
            >
              <UserRound class="h-4 w-4" />
              Account
            </button>
          </nav>
        </div>

        <!-- ACCESS -->
        <div>
          <p class="text-sm font-semibold">Access</p>

          <nav class="mt-5 space-y-3">
            <RouterLink
              to="/login"
              class="block text-sm text-muted-foreground transition hover:text-foreground"
            >
              Sign in
            </RouterLink>

            <RouterLink
              to="/register"
              class="block text-sm text-muted-foreground transition hover:text-foreground"
            >
              Choose account
            </RouterLink>

            <RouterLink
              to="/register/personal"
              class="block text-sm text-muted-foreground transition hover:text-foreground"
            >
              Personal account
            </RouterLink>

            <RouterLink
              to="/register/business"
              class="block text-sm text-muted-foreground transition hover:text-foreground"
            >
              Business workspace
            </RouterLink>
          </nav>
        </div>
      </div>

      <div
        class="mt-14 flex flex-col gap-4 border-t border-border pt-6 sm:flex-row sm:items-center sm:justify-between"
      >
        <p class="text-xs text-muted-foreground">
          {{ brand.copyright }}
        </p>

        <div class="flex flex-wrap gap-x-5 gap-y-2 text-xs text-muted-foreground">
          <span>Connected commerce</span>
          <span>Smarter operations</span>
          <span>Built as one product</span>
        </div>
      </div>
    </div>
  </footer>
</template>
