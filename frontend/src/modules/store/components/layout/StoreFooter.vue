<script setup lang="ts">
import { brand } from '@/shared/config/brand';
</script>

<template>
  <footer class="border-t border-border bg-muted/20">
    <div class="mx-auto grid w-full max-w-[1500px] gap-10 px-6 py-12 md:grid-cols-4">
      <div class="md:col-span-2">
        <p class="text-xl font-semibold tracking-tight">{{ brand.name }}</p>

        <p class="mt-3 max-w-md text-sm leading-6 text-muted-foreground">
          {{ brand.storeDescription }}
        </p>
      </div>

      <div>
        <p class="text-sm font-semibold text-foreground">Marketplace</p>

        <div class="mt-4 space-y-3 text-sm text-muted-foreground">
          <RouterLink to="/store/catalog" class="block transition hover:text-foreground">
            Explore products
          </RouterLink>

          <button class="block transition hover:text-foreground">Categories</button>
        </div>
      </div>

      <div>
        <p class="text-sm font-semibold text-foreground">Account</p>

        <div class="mt-4 space-y-3 text-sm text-muted-foreground">
          <RouterLink to="/login" class="block transition hover:text-foreground">
            Sign in
          </RouterLink>

          <RouterLink to="/register" class="block transition hover:text-foreground">
            Create account
          </RouterLink>

          <RouterLink to="/store/purchases" class="block transition hover:text-foreground">
            My purchases
          </RouterLink>
        </div>
      </div>
    </div>

    <div class="border-t border-border px-6 py-5">
      <div
        class="mx-auto flex w-full max-w-[1500px] flex-col gap-2 text-xs text-muted-foreground sm:flex-row sm:items-center sm:justify-between"
      >
        <span> {{ brand.copyright }} Multi-tenant commerce and operations platform </span>

        <span> Secure multi-tenant commerce </span>
      </div>
    </div>
  </footer>
</template>
