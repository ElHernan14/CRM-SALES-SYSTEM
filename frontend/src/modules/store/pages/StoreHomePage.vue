<script setup lang="ts">
import { ArrowRight, Search, ShieldCheck, Sparkles, Truck } from 'lucide-vue-next';

import { Button } from '@/components/ui/button';
</script>

<template>
  <div>
    <section
      class="relative overflow-hidden border-b border-border bg-gradient-to-br from-background via-background to-primary/5"
    >
      <div
        class="pointer-events-none absolute -right-32 -top-32 h-96 w-96 rounded-full bg-primary/10 blur-3xl"
      />

      <div
        class="mx-auto grid min-h-[620px] max-w-[1500px] items-center gap-12 px-6 py-16 lg:grid-cols-[1.1fr_0.9fr] lg:px-8"
      >
        <div class="relative">
          <div
            class="inline-flex items-center gap-2 rounded-full border border-primary/20 bg-primary/10 px-4 py-2 text-xs font-semibold text-primary"
          >
            <Sparkles class="h-4 w-4" />
            A smarter way to discover commerce
          </div>

          <h1
            class="mt-6 max-w-3xl text-5xl font-semibold tracking-[-0.045em] text-foreground sm:text-6xl lg:text-7xl"
          >
            Products worth discovering.
            <span class="text-muted-foreground"> Stores worth trusting. </span>
          </h1>

          <p class="mt-6 max-w-2xl text-lg leading-8 text-muted-foreground">
            Explore products and services from companies connected through a unified commerce
            network.
          </p>

          <div class="mt-8 flex flex-col gap-3 sm:flex-row">
            <Button size="lg" class="rounded-full px-7" @click="$router.push('/store/catalog')">
              Explore catalog
              <ArrowRight class="ml-2 h-4 w-4" />
            </Button>

            <Button size="lg" variant="outline" class="rounded-full px-7">
              <Search class="mr-2 h-4 w-4" />
              Search products
            </Button>
          </div>

          <div class="mt-10 flex flex-wrap gap-x-8 gap-y-4 text-sm text-muted-foreground">
            <div class="flex items-center gap-2">
              <ShieldCheck class="h-4 w-4 text-primary" />
              Verified companies
            </div>

            <div class="flex items-center gap-2">
              <Truck class="h-4 w-4 text-primary" />
              Real inventory
            </div>

            <div class="flex items-center gap-2">
              <Sparkles class="h-4 w-4 text-primary" />
              Curated discovery
            </div>
          </div>
        </div>

        <div
          class="relative min-h-[420px] overflow-hidden rounded-[2rem] border border-border bg-card shadow-2xl"
        >
          <div
            class="absolute inset-0 bg-gradient-to-br from-primary/15 via-muted/40 to-background"
          />

          <div class="relative flex h-full min-h-[420px] items-center justify-center p-8">
            <p class="max-w-sm text-center text-sm text-muted-foreground">
              Featured product experience will be rendered here once the catalog API is connected.
            </p>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>
