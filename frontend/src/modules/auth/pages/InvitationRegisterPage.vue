<script setup lang="ts">
import { brand } from '@/shared/config/brand';
</script>

<template>
  <main class="flex min-h-screen items-center justify-center bg-background px-6">
    <div class="max-w-md text-center">
      <p class="text-sm font-semibold text-primary">
        {{ brand.name }}
      </p>

      <h1 class="mt-4 text-3xl font-semibold">Personal registration</h1>

      <p class="mt-3 text-muted-foreground">
        Personal account creation will be connected to the customer registration API.
      </p>

      <RouterLink to="/register" class="mt-6 inline-block text-sm font-semibold text-primary">
        Back to account options
      </RouterLink>
    </div>
  </main>
</template>
