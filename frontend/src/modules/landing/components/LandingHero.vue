<script setup lang="ts">
import { computed } from 'vue';
import { useRouter } from 'vue-router';
import { storeToRefs } from 'pinia';

import { ArrowRight, Boxes, Building2, CreditCard, Network, ShoppingBag } from 'lucide-vue-next';

import { Button } from '@/components/ui/button';

import ProductComposition from './ProductComposition.vue';

import { useAuthStore } from '@/modules/auth/stores/auth.store';

const router = useRouter();

const auth = useAuthStore();

const { user } = storeToRefs(auth);

const isAuthenticated = computed(() => {
  return Boolean(user.value);
});

const isBusinessUser = computed(() => {
  return Boolean(user.value?.company_id);
});

const primaryLabel = computed(() => {
  if (!isAuthenticated.value) {
    return 'Explore Nexora';
  }

  return isBusinessUser.value ? 'Open business workspace' : 'Continue to Store';
});

function openPrimaryExperience() {
  if (!isAuthenticated.value) {
    router.push('/register');
    return;
  }

  router.push(isBusinessUser.value ? '/erp/dashboard' : '/store');
}
</script>

<template>
  <section class="relative overflow-hidden border-b border-border bg-background pt-16">
    <!-- BACKGROUND -->
    <div
      class="pointer-events-none absolute -left-48 top-28 h-[540px] w-[540px] rounded-full bg-primary/10 blur-3xl"
    />

    <div
      class="pointer-events-none absolute -right-40 top-0 h-[620px] w-[620px] rounded-full bg-violet-500/10 blur-3xl"
    />

    <div
      class="pointer-events-none absolute inset-x-0 top-16 h-px bg-gradient-to-r from-transparent via-primary/40 to-transparent"
    />

    <div
      class="relative mx-auto grid min-h-[680px] w-full max-w-[1500px] items-center gap-12 px-4 py-14 sm:px-6 sm:py-20 lg:min-h-[760px] lg:grid-cols-[0.95fr_1.05fr] lg:gap-16 lg:px-8 lg:py-24"
    >
      <!-- COPY -->
      <div>
        <p class="text-sm font-semibold text-primary">Commerce and operations, working as one</p>

        <h1
          class="mt-5 max-w-3xl text-[2.65rem] font-semibold leading-[1.02] tracking-[-0.055em] text-foreground sm:mt-6 sm:text-6xl lg:text-[4.75rem] lg:leading-[0.98]"
        >
          Your business and commerce,
          <span class="text-muted-foreground"> finally connected. </span>
        </h1>

        <p
          class="mt-6 max-w-2xl text-base leading-7 text-muted-foreground sm:mt-7 sm:text-lg sm:leading-8"
        >
          Run products, inventory, sales and purchases from one workspace—then bring your catalog
          directly to customers through a connected Store.
        </p>

        <div class="mt-8 flex flex-col gap-3 sm:mt-9 sm:flex-row">
          <Button
            size="lg"
            class="w-full rounded-full px-7 sm:w-auto"
            @click="openPrimaryExperience"
          >
            {{ primaryLabel }}

            <ArrowRight class="ml-2 h-4 w-4" />
          </Button>

          <Button
            size="lg"
            variant="outline"
            class="w-full rounded-full px-7 sm:w-auto"
            @click="router.push('/store')"
          >
            <ShoppingBag class="mr-2 h-4 w-4" />
            Open Store
          </Button>
        </div>

        <!-- CORE CAPABILITIES -->
        <div
          class="mt-10 grid max-w-2xl grid-cols-2 gap-x-4 gap-y-6 border-t border-border pt-7 sm:mt-11 sm:grid-cols-4 sm:gap-x-6"
        >
          <div>
            <Network class="h-5 w-5 text-primary" />

            <p class="mt-3 text-sm font-semibold">Multi-tenant</p>

            <p class="mt-1 text-xs leading-5 text-muted-foreground">Isolated business contexts</p>
          </div>

          <div>
            <Building2 class="h-5 w-5 text-primary" />

            <p class="mt-3 text-sm font-semibold">B2B & B2C</p>

            <p class="mt-1 text-xs leading-5 text-muted-foreground">Two commerce experiences</p>
          </div>

          <div>
            <Boxes class="h-5 w-5 text-primary" />

            <p class="mt-3 text-sm font-semibold">Live inventory</p>

            <p class="mt-1 text-xs leading-5 text-muted-foreground">Connected to operations</p>
          </div>

          <div>
            <CreditCard class="h-5 w-5 text-primary" />

            <p class="mt-3 text-sm font-semibold">Payment lifecycle</p>

            <p class="mt-1 text-xs leading-5 text-muted-foreground">Balances and history</p>
          </div>
        </div>
      </div>

      <!-- PRODUCT VISUAL -->
      <ProductComposition />
    </div>

    <!-- BOTTOM TRUST BAND -->
    <div class="relative border-t border-border/70 bg-muted/15">
      <div
        class="mx-auto flex w-full max-w-[1500px] flex-col gap-4 px-4 py-5 sm:flex-row sm:items-center sm:justify-between sm:px-6 lg:px-8"
      >
        <p class="text-sm font-semibold">Built for connected operations</p>

        <div class="flex flex-wrap gap-x-6 gap-y-3 text-xs font-medium text-muted-foreground">
          <span>Tenant-scoped data</span>
          <span>Dynamic permissions</span>
          <span>Inventory rules</span>
          <span>Transactional checkout</span>
          <span>Invoice lifecycle</span>
        </div>
      </div>
    </div>
  </section>
</template>
