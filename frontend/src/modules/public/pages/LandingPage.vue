<script setup lang="ts">
import { useRouter } from 'vue-router';

import { Building2, ShoppingBag, Sparkles } from 'lucide-vue-next';

import { Button } from '@/components/ui/button';

import { brand } from '@/shared/config/brand';

import PublicHeader from '../components/PublicHeader.vue';

const router = useRouter();
</script>

<template>
  <div class="min-h-screen bg-background">
    <PublicHeader />

    <main>
      <section class="relative overflow-hidden border-b border-border">
        <div
          class="pointer-events-none absolute -left-48 -top-56 h-[600px] w-[600px] rounded-full bg-primary/10 blur-3xl"
        />

        <div
          class="pointer-events-none absolute -right-48 top-20 h-[520px] w-[520px] rounded-full bg-blue-500/10 blur-3xl"
        />

        <div
          class="relative mx-auto grid min-h-[680px] w-full max-w-[1500px] items-center gap-14 px-6 py-20 lg:grid-cols-[1.08fr_0.92fr] lg:px-8"
        >
          <div>
            <div
              class="inline-flex items-center gap-2 rounded-full border border-primary/20 bg-primary/10 px-4 py-2 text-xs font-semibold text-primary"
            >
              <Sparkles class="h-4 w-4" />
              {{ brand.shortDescription }}
            </div>

            <h1
              class="mt-7 max-w-4xl text-5xl font-semibold tracking-[-0.05em] text-foreground sm:text-6xl lg:text-7xl"
            >
              Commerce connected to
              <span class="text-muted-foreground"> real operations. </span>
            </h1>

            <p class="mt-6 max-w-2xl text-lg leading-8 text-muted-foreground">
              Nexora connects product discovery, B2C commerce, B2B purchasing and business
              operations through a single multi-tenant platform.
            </p>

            <div class="mt-9 flex flex-col gap-3 sm:flex-row">
              <Button size="lg" class="rounded-full px-7" @click="router.push('/store')">
                <ShoppingBag class="mr-2 h-4 w-4" />
                Explore Store
              </Button>

              <Button
                size="lg"
                variant="outline"
                class="rounded-full px-7"
                @click="router.push('/login')"
              >
                <Building2 class="mr-2 h-4 w-4" />
                Access workspace
              </Button>
            </div>
          </div>

          <div
            class="relative min-h-[480px] overflow-hidden rounded-[2rem] border border-border bg-card shadow-2xl"
          >
            <div
              class="absolute inset-0 bg-gradient-to-br from-primary/15 via-background to-blue-500/10"
            />

            <div
              class="absolute left-8 right-20 top-10 rounded-2xl border border-border bg-background/90 p-5 shadow-lg backdrop-blur"
            >
              <p class="text-xs text-muted-foreground">Connected commerce</p>

              <p class="mt-2 text-xl font-semibold">Store → Purchase → Payment</p>
            </div>

            <div
              class="absolute bottom-12 left-20 right-8 rounded-2xl border border-border bg-background/90 p-5 shadow-xl backdrop-blur"
            >
              <p class="text-xs text-muted-foreground">Business operations</p>

              <p class="mt-2 text-xl font-semibold">Products → Sales → ERP</p>
            </div>
          </div>
        </div>
      </section>
    </main>
  </div>
</template>
