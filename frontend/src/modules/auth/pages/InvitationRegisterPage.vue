<script setup lang="ts">
import { useRoute } from 'vue-router';

import { Building2, KeyRound } from 'lucide-vue-next';

import { Button } from '@/components/ui/button';

import ThemeToggle from '@/shared/components/ThemeToggle.vue';

import { brand } from '@/shared/config/brand';

const route = useRoute();

const token = String(route.params.token ?? '');
</script>

<template>
  <main
    class="relative flex min-h-screen items-center justify-center overflow-hidden bg-background px-6 py-20"
  >
    <div class="absolute right-6 top-6">
      <ThemeToggle />
    </div>

    <div
      class="relative w-full max-w-lg rounded-[2rem] border border-border bg-card p-8 text-center shadow-2xl"
    >
      <div class="mx-auto flex h-14 w-14 items-center justify-center rounded-2xl bg-primary/10">
        <KeyRound class="h-7 w-7 text-primary" />
      </div>

      <p class="mt-6 text-sm font-semibold text-primary">{{ brand.name }} invitation</p>

      <h1 class="mt-3 text-3xl font-semibold tracking-tight">Join a business workspace</h1>

      <p class="mt-4 text-sm leading-6 text-muted-foreground">
        Employee invitations are prepared as the next phase of company administration.
      </p>

      <div class="mt-6 rounded-xl border border-border bg-muted/20 p-4">
        <div class="flex items-center justify-center gap-2 text-xs text-muted-foreground">
          <Building2 class="h-4 w-4" />
          Invitation token detected
        </div>

        <p class="mt-2 truncate font-mono text-xs">
          {{ token }}
        </p>
      </div>

      <Button class="mt-7 rounded-full" @click="$router.push('/login')"> Return to sign in </Button>
    </div>
  </main>
</template>
