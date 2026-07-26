<script setup lang="ts">
import { computed } from 'vue';
import { useRouter } from 'vue-router';
import { storeToRefs } from 'pinia';

import {
  ArrowRight,
  BadgeCheck,
  Boxes,
  Building2,
  Check,
  CircleDollarSign,
  CreditCard,
  FileCheck2,
  Landmark,
  PackageCheck,
  ReceiptText,
  ShieldCheck,
  ShoppingBag,
  Store,
  UserRound,
  UsersRound,
  WalletCards,
} from 'lucide-vue-next';

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

function openBusinessFlow() {
  if (isBusinessUser.value) {
    router.push('/erp/dashboard');
    return;
  }

  router.push('/register/business');
}

function openConsumerFlow() {
  router.push('/store');
}

function openRegistration() {
  if (isAuthenticated.value) {
    router.push(isBusinessUser.value ? '/erp/dashboard' : '/store');

    return;
  }

  router.push('/register');
}
</script>

<template>
  <section class="relative overflow-hidden border-b border-border bg-muted/10">
    <div
      class="pointer-events-none absolute -left-48 top-20 h-[520px] w-[520px] rounded-full bg-blue-500/5 blur-3xl"
    />

    <div
      class="pointer-events-none absolute -right-52 bottom-0 h-[520px] w-[520px] rounded-full bg-violet-500/5 blur-3xl"
    />

    <div class="relative mx-auto w-full max-w-[1500px] px-6 py-24 lg:px-8 lg:py-28">
      <!-- HEADING -->
      <div class="mx-auto max-w-4xl text-center">
        <p class="text-sm font-semibold text-primary">Two sides of commerce</p>

        <h2 class="mt-4 text-4xl font-semibold tracking-[-0.055em] text-foreground sm:text-5xl">
          Built for businesses.
          <span class="text-muted-foreground"> Designed for customers. </span>
        </h2>

        <p class="mx-auto mt-5 max-w-2xl text-base leading-7 text-muted-foreground">
          Nexora supports organizational commerce and personal purchasing through dedicated
          experiences built over the same transactional foundation.
        </p>
      </div>

      <!-- MODES -->
      <div class="mt-16 grid gap-6 xl:grid-cols-2">
        <!-- B2B -->
        <article
          class="group relative overflow-hidden rounded-[2rem] border border-border bg-card shadow-sm transition duration-300 hover:-translate-y-1 hover:border-primary/25 hover:shadow-2xl"
        >
          <div
            class="pointer-events-none absolute -right-28 -top-28 h-72 w-72 rounded-full bg-blue-500/10 blur-3xl"
          />

          <div class="relative p-6 sm:p-8">
            <div class="flex flex-col gap-5 sm:flex-row sm:items-start sm:justify-between">
              <div class="flex items-center gap-4">
                <div
                  class="flex h-14 w-14 shrink-0 items-center justify-center rounded-2xl bg-blue-500 text-white shadow-lg shadow-blue-500/15"
                >
                  <Building2 class="h-6 w-6" />
                </div>

                <div>
                  <p
                    class="text-xs font-semibold uppercase tracking-[0.16em] text-blue-600 dark:text-blue-400"
                  >
                    Business to business
                  </p>

                  <h3 class="mt-2 text-2xl font-semibold tracking-[-0.035em]">
                    Organizational commerce
                  </h3>
                </div>
              </div>

              <span
                class="self-start rounded-full border border-blue-500/20 bg-blue-500/10 px-3 py-1.5 text-xs font-semibold text-blue-600 dark:text-blue-400"
              >
                ERP workflow
              </span>
            </div>

            <p class="mt-6 max-w-xl text-sm leading-7 text-muted-foreground">
              Companies sell directly to clients, purchase from suppliers and maintain commercial
              activity inside their tenant workspace.
            </p>

            <!-- B2B FLOW -->
            <div class="mt-8 overflow-hidden rounded-[1.5rem] border border-border bg-muted/15">
              <div class="grid divide-y divide-border sm:grid-cols-3 sm:divide-x sm:divide-y-0">
                <div class="p-4">
                  <div class="flex h-10 w-10 items-center justify-center rounded-xl bg-blue-500/10">
                    <UsersRound class="h-5 w-5 text-blue-600 dark:text-blue-400" />
                  </div>

                  <p class="mt-4 text-sm font-semibold">Business clients</p>

                  <p class="mt-2 text-xs leading-5 text-muted-foreground">
                    Sales remain connected to organizational buyers.
                  </p>
                </div>

                <div class="p-4">
                  <div class="flex h-10 w-10 items-center justify-center rounded-xl bg-primary/10">
                    <ReceiptText class="h-5 w-5 text-primary" />
                  </div>

                  <p class="mt-4 text-sm font-semibold">Controlled invoices</p>

                  <p class="mt-2 text-xs leading-5 text-muted-foreground">
                    Draft, submit, cancel and payment operations.
                  </p>
                </div>

                <div class="p-4">
                  <div
                    class="flex h-10 w-10 items-center justify-center rounded-xl bg-emerald-500/10"
                  >
                    <Landmark class="h-5 w-5 text-emerald-600 dark:text-emerald-400" />
                  </div>

                  <p class="mt-4 text-sm font-semibold">Financial visibility</p>

                  <p class="mt-2 text-xs leading-5 text-muted-foreground">
                    Paid and outstanding balances from both sides.
                  </p>
                </div>
              </div>
            </div>

            <!-- BUSINESS TRANSACTION -->
            <div class="mt-5 rounded-[1.5rem] border border-border bg-background p-5">
              <div class="flex flex-col gap-5 sm:flex-row sm:items-center sm:justify-between">
                <div class="flex min-w-0 items-center gap-4">
                  <div
                    class="flex h-11 w-11 shrink-0 items-center justify-center rounded-xl bg-primary/10"
                  >
                    <Building2 class="h-5 w-5 text-primary" />
                  </div>

                  <div class="min-w-0">
                    <div class="flex items-center gap-2">
                      <p class="truncate text-sm font-semibold">Constante Labs</p>

                      <ArrowRight class="h-3.5 w-3.5 text-muted-foreground" />

                      <p class="truncate text-sm font-semibold">Mercado Norte</p>
                    </div>

                    <p class="mt-1 text-xs text-muted-foreground">Organizational sales invoice</p>
                  </div>
                </div>

                <div class="flex items-center gap-3 sm:text-right">
                  <span
                    class="rounded-full border border-amber-500/20 bg-amber-500/10 px-2.5 py-1 text-xs font-semibold text-amber-600 dark:text-amber-400"
                  >
                    Pending
                  </span>

                  <p class="text-base font-semibold">$480.000,00</p>
                </div>
              </div>
            </div>

            <!-- FEATURES -->
            <div class="mt-6 grid gap-3 sm:grid-cols-2">
              <div class="flex items-start gap-3">
                <Check class="mt-1 h-4 w-4 shrink-0 text-primary" />

                <p class="text-sm text-muted-foreground">Company-scoped purchasing and selling</p>
              </div>

              <div class="flex items-start gap-3">
                <Check class="mt-1 h-4 w-4 shrink-0 text-primary" />

                <p class="text-sm text-muted-foreground">Buyer and seller ownership validation</p>
              </div>

              <div class="flex items-start gap-3">
                <Check class="mt-1 h-4 w-4 shrink-0 text-primary" />

                <p class="text-sm text-muted-foreground">Inventory reservation and release</p>
              </div>

              <div class="flex items-start gap-3">
                <Check class="mt-1 h-4 w-4 shrink-0 text-primary" />

                <p class="text-sm text-muted-foreground">Partial and complete payments</p>
              </div>
            </div>

            <Button variant="outline" class="mt-8 rounded-full" @click="openBusinessFlow">
              {{ isBusinessUser ? 'Open ERP workflow' : 'Create business workspace' }}

              <ArrowRight class="ml-2 h-4 w-4" />
            </Button>
          </div>
        </article>

        <!-- B2C -->
        <article
          class="group relative overflow-hidden rounded-[2rem] border border-border bg-card shadow-sm transition duration-300 hover:-translate-y-1 hover:border-primary/25 hover:shadow-2xl"
        >
          <div
            class="pointer-events-none absolute -right-28 -top-28 h-72 w-72 rounded-full bg-violet-500/10 blur-3xl"
          />

          <div class="relative p-6 sm:p-8">
            <div class="flex flex-col gap-5 sm:flex-row sm:items-start sm:justify-between">
              <div class="flex items-center gap-4">
                <div
                  class="flex h-14 w-14 shrink-0 items-center justify-center rounded-2xl bg-violet-500 text-white shadow-lg shadow-violet-500/15"
                >
                  <UserRound class="h-6 w-6" />
                </div>

                <div>
                  <p
                    class="text-xs font-semibold uppercase tracking-[0.16em] text-violet-600 dark:text-violet-400"
                  >
                    Business to consumer
                  </p>

                  <h3 class="mt-2 text-2xl font-semibold tracking-[-0.035em]">Personal commerce</h3>
                </div>
              </div>

              <span
                class="self-start rounded-full border border-violet-500/20 bg-violet-500/10 px-3 py-1.5 text-xs font-semibold text-violet-600 dark:text-violet-400"
              >
                Store experience
              </span>
            </div>

            <p class="mt-6 max-w-xl text-sm leading-7 text-muted-foreground">
              Individual buyers discover real businesses and products, create multi-seller carts and
              track their personal orders from Store.
            </p>

            <!-- B2C FLOW -->
            <div class="mt-8 overflow-hidden rounded-[1.5rem] border border-border bg-muted/15">
              <div class="grid divide-y divide-border sm:grid-cols-3 sm:divide-x sm:divide-y-0">
                <div class="p-4">
                  <div
                    class="flex h-10 w-10 items-center justify-center rounded-xl bg-violet-500/10"
                  >
                    <Store class="h-5 w-5 text-violet-600 dark:text-violet-400" />
                  </div>

                  <p class="mt-4 text-sm font-semibold">Seller discovery</p>

                  <p class="mt-2 text-xs leading-5 text-muted-foreground">
                    Browse businesses, industries and products.
                  </p>
                </div>

                <div class="p-4">
                  <div class="flex h-10 w-10 items-center justify-center rounded-xl bg-primary/10">
                    <ShoppingBag class="h-5 w-5 text-primary" />
                  </div>

                  <p class="mt-4 text-sm font-semibold">Multi-seller cart</p>

                  <p class="mt-2 text-xs leading-5 text-muted-foreground">
                    Items remain grouped by their seller company.
                  </p>
                </div>

                <div class="p-4">
                  <div
                    class="flex h-10 w-10 items-center justify-center rounded-xl bg-emerald-500/10"
                  >
                    <CreditCard class="h-5 w-5 text-emerald-600 dark:text-emerald-400" />
                  </div>

                  <p class="mt-4 text-sm font-semibold">Personal payments</p>

                  <p class="mt-2 text-xs leading-5 text-muted-foreground">
                    Each order retains an independent balance.
                  </p>
                </div>
              </div>
            </div>

            <!-- STORE TRANSACTION -->
            <div class="mt-5 rounded-[1.5rem] border border-border bg-background p-5">
              <div class="flex flex-col gap-5 sm:flex-row sm:items-center sm:justify-between">
                <div class="flex min-w-0 items-center gap-4">
                  <div
                    class="flex h-11 w-11 shrink-0 items-center justify-center rounded-xl bg-violet-500/10"
                  >
                    <ShoppingBag class="h-5 w-5 text-violet-600 dark:text-violet-400" />
                  </div>

                  <div class="min-w-0">
                    <div class="flex items-center gap-2">
                      <p class="truncate text-sm font-semibold">One customer cart</p>

                      <BadgeCheck class="h-4 w-4 text-primary" />
                    </div>

                    <p class="mt-1 text-xs text-muted-foreground">
                      TechNova · Constructa · Mercado Norte
                    </p>
                  </div>
                </div>

                <div class="sm:text-right">
                  <p class="text-[10px] font-medium uppercase tracking-wide text-muted-foreground">
                    Seller orders
                  </p>

                  <p class="mt-1 text-base font-semibold">3 invoices</p>
                </div>
              </div>
            </div>

            <!-- FEATURES -->
            <div class="mt-6 grid gap-3 sm:grid-cols-2">
              <div class="flex items-start gap-3">
                <Check class="mt-1 h-4 w-4 shrink-0 text-primary" />

                <p class="text-sm text-muted-foreground">Public business and product discovery</p>
              </div>

              <div class="flex items-start gap-3">
                <Check class="mt-1 h-4 w-4 shrink-0 text-primary" />

                <p class="text-sm text-muted-foreground">Live product availability</p>
              </div>

              <div class="flex items-start gap-3">
                <Check class="mt-1 h-4 w-4 shrink-0 text-primary" />

                <p class="text-sm text-muted-foreground">Transactional checkout across sellers</p>
              </div>

              <div class="flex items-start gap-3">
                <Check class="mt-1 h-4 w-4 shrink-0 text-primary" />

                <p class="text-sm text-muted-foreground">Personal purchases and payment history</p>
              </div>
            </div>

            <Button variant="outline" class="mt-8 rounded-full" @click="openConsumerFlow">
              Explore personal commerce

              <ArrowRight class="ml-2 h-4 w-4" />
            </Button>
          </div>
        </article>
      </div>

      <!-- SHARED ENGINE -->
      <div
        class="relative mt-8 overflow-hidden rounded-[2rem] border border-primary/20 bg-primary/5 p-6 sm:p-8"
      >
        <div
          class="pointer-events-none absolute -right-20 -top-20 h-52 w-52 rounded-full bg-primary/10 blur-3xl"
        />

        <div class="relative flex flex-col gap-8 lg:flex-row lg:items-center lg:justify-between">
          <div class="max-w-2xl">
            <p class="text-xs font-semibold uppercase tracking-[0.16em] text-primary">
              Shared transactional foundation
            </p>

            <h3 class="mt-3 text-3xl font-semibold tracking-[-0.04em]">
              Different experiences.
              <span class="text-muted-foreground"> One invoice engine. </span>
            </h3>

            <p class="mt-4 text-sm leading-7 text-muted-foreground">
              ERP and Store preserve their own purchasing experiences while sharing consistent
              inventory rules, invoices, payments and ownership validation.
            </p>
          </div>

          <div class="grid shrink-0 grid-cols-2 gap-3 sm:grid-cols-4">
            <div
              class="flex min-w-28 flex-col items-center rounded-2xl border border-border bg-background p-4 text-center shadow-sm"
            >
              <ReceiptText class="h-5 w-5 text-primary" />

              <span class="mt-3 text-xs font-semibold"> Invoices </span>
            </div>

            <div
              class="flex min-w-28 flex-col items-center rounded-2xl border border-border bg-background p-4 text-center shadow-sm"
            >
              <Boxes class="h-5 w-5 text-primary" />

              <span class="mt-3 text-xs font-semibold"> Inventory </span>
            </div>

            <div
              class="flex min-w-28 flex-col items-center rounded-2xl border border-border bg-background p-4 text-center shadow-sm"
            >
              <WalletCards class="h-5 w-5 text-primary" />

              <span class="mt-3 text-xs font-semibold"> Payments </span>
            </div>

            <div
              class="flex min-w-28 flex-col items-center rounded-2xl border border-border bg-background p-4 text-center shadow-sm"
            >
              <ShieldCheck class="h-5 w-5 text-primary" />

              <span class="mt-3 text-xs font-semibold"> Ownership </span>
            </div>
          </div>
        </div>

        <div
          class="relative mt-7 flex flex-col gap-4 border-t border-primary/15 pt-6 sm:flex-row sm:items-center sm:justify-between"
        >
          <div class="flex items-center gap-3">
            <div
              class="flex h-10 w-10 items-center justify-center rounded-xl bg-primary text-primary-foreground"
            >
              <FileCheck2 class="h-5 w-5" />
            </div>

            <p class="text-sm font-semibold">
              Consistent business rules across every commercial origin.
            </p>
          </div>

          <Button class="self-start rounded-full sm:self-auto" @click="openRegistration">
            {{ isAuthenticated ? 'Open Nexora' : 'Choose your Nexora experience' }}

            <ArrowRight class="ml-2 h-4 w-4" />
          </Button>
        </div>
      </div>
    </div>
  </section>
</template>
