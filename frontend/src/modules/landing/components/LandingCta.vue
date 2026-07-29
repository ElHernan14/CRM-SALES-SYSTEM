<script setup lang="ts">
import { computed } from 'vue';
import { useRouter } from 'vue-router';
import { storeToRefs } from 'pinia';

import { ArrowRight, Building2, ShoppingBag } from 'lucide-vue-next';

import { Button } from '@/components/ui/button';

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
    return 'Create your Nexora account';
  }

  return isBusinessUser.value ? 'Open business workspace' : 'Continue to Store';
});

function openPrimary() {
  if (!isAuthenticated.value) {
    router.push('/register');
    return;
  }

  router.push(isBusinessUser.value ? '/erp/dashboard' : '/store');
}
</script>

<template>
  <section class="relative overflow-hidden bg-background">
    <div
      class="pointer-events-none absolute left-1/2 top-0 h-[600px] w-[900px] -translate-x-1/2 rounded-full bg-primary/10 blur-3xl"
    />

    <div class="relative mx-auto w-full max-w-[1500px] px-6 py-24 lg:px-8 lg:py-32">
      <div
        class="nexora-glow relative overflow-hidden rounded-[2.5rem] border border-border bg-zinc-950 px-6 py-16 text-center text-white shadow-2xl sm:px-10 lg:py-20"
      >
        <div
          class="pointer-events-none absolute -left-40 -top-40 h-[440px] w-[440px] rounded-full bg-primary/25 blur-3xl"
        />

        <div
          class="pointer-events-none absolute -right-40 bottom-0 h-[440px] w-[440px] rounded-full bg-violet-500/20 blur-3xl"
        />

        <div
          class="pointer-events-none absolute inset-0 bg-[linear-gradient(to_right,rgba(255,255,255,0.035)_1px,transparent_1px),linear-gradient(to_bottom,rgba(255,255,255,0.035)_1px,transparent_1px)] bg-[size:64px_64px]"
        />

        <div class="relative mx-auto max-w-4xl">
          <p class="text-sm font-semibold text-primary">
            Build the business. Connect the commerce.
          </p>

          <h2
            class="mt-5 text-4xl font-semibold tracking-[-0.06em] text-white sm:text-5xl lg:text-6xl"
          >
            Operations, products and transactions
            <span class="nexora-gradient-text"> working as one. </span>
          </h2>

          <p class="mx-auto mt-6 max-w-2xl text-base leading-7 text-zinc-400">
            Explore a platform where companies operate, customers discover and every commercial
            event remains connected.
          </p>

          <div class="mt-9 flex flex-col justify-center gap-3 sm:flex-row">
            <Button size="lg" class="group rounded-full px-8" @click="openPrimary">
              <Building2 v-if="isBusinessUser" class="mr-2 h-4 w-4" />

              {{ primaryLabel }}

              <ArrowRight
                class="ml-2 h-4 w-4 transition-transform duration-200 group-hover:translate-x-0.5"
              />
            </Button>

            <Button
              size="lg"
              variant="outline"
              class="rounded-full border-white/15 bg-white/5 px-8 text-white hover:bg-white/10 hover:text-white"
              @click="router.push('/store')"
            >
              <ShoppingBag class="mr-2 h-4 w-4" />
              Explore Store
            </Button>
          </div>

          <div
            class="mt-10 flex flex-wrap justify-center gap-x-7 gap-y-3 text-xs font-medium text-zinc-500"
          >
            <span>Multi-tenant</span>
            <span>B2B & B2C</span>
            <span>Connected inventory</span>
            <span>Invoice lifecycle</span>
            <span>Payment tracking</span>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
