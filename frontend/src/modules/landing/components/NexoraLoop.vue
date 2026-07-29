<script setup lang="ts">
import {
  ArrowDown,
  ArrowRight,
  Boxes,
  Building2,
  CircleDollarSign,
  LayoutDashboard,
  PackageCheck,
  ReceiptText,
  ShoppingBag,
  Store,
} from 'lucide-vue-next';

const loopSteps = [
  {
    number: '01',
    title: 'Operate',
    description:
      'Build the company workspace, manage products, clients, inventory and commercial information.',
    icon: Building2,
    result: 'Business workspace',
    accentClass: 'bg-primary text-primary-foreground',
  },

  {
    number: '02',
    title: 'Publish',
    description:
      'Make active products and services available throughout the connected Nexora Store.',
    icon: Store,
    result: 'Connected catalog',
    accentClass: 'bg-violet-500 text-white',
  },

  {
    number: '03',
    title: 'Transact',
    description: 'Create B2B invoices or receive consumer orders grouped automatically by seller.',
    icon: ShoppingBag,
    result: 'Invoice created',
    accentClass: 'bg-blue-500 text-white',
  },

  {
    number: '04',
    title: 'Track',
    description:
      'Follow pending balances, partial payments, completed orders and inventory impact.',
    icon: CircleDollarSign,
    result: 'Financial lifecycle',
    accentClass: 'bg-emerald-500 text-white',
  },
] as const;
</script>

<template>
  <section
    id="nexora-solutions"
    class="relative scroll-mt-24 overflow-hidden border-b border-border bg-muted/10"
  >
    <div
      class="pointer-events-none absolute -right-44 top-24 h-[500px] w-[500px] rounded-full bg-violet-500/5 blur-3xl"
    />

    <div class="mx-auto w-full max-w-[1500px] px-6 py-24 lg:px-8 lg:py-28">
      <!-- HEADING -->
      <div class="mx-auto max-w-4xl text-center">
        <p class="text-sm font-semibold text-primary">The Nexora loop</p>

        <h2 class="mt-4 text-4xl font-semibold tracking-[-0.05em] text-foreground sm:text-5xl">
          One connected business loop.
        </h2>

        <p class="mx-auto mt-5 max-w-2xl text-base leading-7 text-muted-foreground">
          What happens in the ERP becomes commerce. What happens in commerce returns as invoices,
          payments, inventory movement and business information.
        </p>
      </div>

      <!-- LOOP STEPS -->
      <div class="relative mt-16 grid gap-5 lg:grid-cols-4">
        <article
          v-for="(step, index) in loopSteps"
          :key="step.title"
          class="nexora-interactive nexora-shine group relative flex min-h-[310px] flex-col overflow-hidden rounded-[1.75rem] border border-border bg-card p-6 shadow-sm hover:border-primary/25 hover:shadow-xl"
        >
          <div
            class="pointer-events-none absolute -right-12 -top-12 h-32 w-32 rounded-full bg-primary/5 blur-3xl transition group-hover:bg-primary/10"
          />

          <div class="relative flex items-start justify-between">
            <div
              class="flex h-12 w-12 items-center justify-center rounded-2xl shadow-sm"
              :class="step.accentClass"
            >
              <component :is="step.icon" class="h-5 w-5" />
            </div>

            <span class="text-sm font-semibold text-muted-foreground/60">
              {{ step.number }}
            </span>
          </div>

          <div class="relative mt-7 flex-1">
            <h3 class="text-xl font-semibold tracking-tight">
              {{ step.title }}
            </h3>

            <p class="mt-3 text-sm leading-6 text-muted-foreground">
              {{ step.description }}
            </p>
          </div>

          <div class="relative mt-6 flex items-center gap-2 border-t border-border pt-4">
            <PackageCheck class="h-4 w-4 text-primary" />

            <span class="text-xs font-semibold text-foreground">
              {{ step.result }}
            </span>
          </div>

          <!-- DESKTOP CONNECTOR -->
          <div
            v-if="index < loopSteps.length - 1"
            class="absolute -right-[19px] top-1/2 z-20 hidden h-9 w-9 -translate-y-1/2 items-center justify-center rounded-full border border-border bg-background shadow-md lg:flex"
          >
            <ArrowRight class="h-4 w-4 text-primary" />
          </div>

          <!-- MOBILE CONNECTOR -->
          <div
            v-if="index < loopSteps.length - 1"
            class="absolute -bottom-[19px] left-1/2 z-20 flex h-9 w-9 -translate-x-1/2 items-center justify-center rounded-full border border-border bg-background shadow-md lg:hidden"
          >
            <ArrowDown class="h-4 w-4 text-primary" />
          </div>
        </article>
      </div>

      <!-- SYSTEM CONNECTION -->
      <div class="mt-10 overflow-hidden rounded-[2rem] border border-border bg-card shadow-sm">
        <div class="grid items-stretch lg:grid-cols-[1fr_auto_1fr_auto_1fr]">
          <!-- ERP -->
          <div class="p-6 sm:p-8">
            <div class="flex h-12 w-12 items-center justify-center rounded-2xl bg-primary/10">
              <LayoutDashboard class="h-5 w-5 text-primary" />
            </div>

            <p class="mt-5 text-xs font-semibold uppercase tracking-[0.14em] text-primary">
              ERP Workspace
            </p>

            <h3 class="mt-2 text-xl font-semibold">Business operations</h3>

            <p class="mt-3 text-sm leading-6 text-muted-foreground">
              Company identity, products, inventory, clients, sales and organizational purchases.
            </p>
          </div>

          <div
            class="hidden items-center justify-center border-x border-border bg-muted/15 px-5 lg:flex"
          >
            <ArrowRight class="h-5 w-5 text-primary" />
          </div>

          <!-- COMMERCE -->
          <div class="border-y border-border p-6 sm:p-8 lg:border-y-0">
            <div class="flex h-12 w-12 items-center justify-center rounded-2xl bg-violet-500/10">
              <ShoppingBag class="h-5 w-5 text-violet-600 dark:text-violet-400" />
            </div>

            <p
              class="mt-5 text-xs font-semibold uppercase tracking-[0.14em] text-violet-600 dark:text-violet-400"
            >
              Connected commerce
            </p>

            <h3 class="mt-2 text-xl font-semibold">Store and Marketplace</h3>

            <p class="mt-3 text-sm leading-6 text-muted-foreground">
              Businesses become sellers, products become discoverable and buyers create
              seller-specific orders.
            </p>
          </div>

          <div
            class="hidden items-center justify-center border-x border-border bg-muted/15 px-5 lg:flex"
          >
            <ArrowRight class="h-5 w-5 text-primary" />
          </div>

          <!-- FINANCE -->
          <div class="p-6 sm:p-8">
            <div class="flex h-12 w-12 items-center justify-center rounded-2xl bg-emerald-500/10">
              <ReceiptText class="h-5 w-5 text-emerald-600 dark:text-emerald-400" />
            </div>

            <p
              class="mt-5 text-xs font-semibold uppercase tracking-[0.14em] text-emerald-600 dark:text-emerald-400"
            >
              Financial lifecycle
            </p>

            <h3 class="mt-2 text-xl font-semibold">Invoices and payments</h3>

            <p class="mt-3 text-sm leading-6 text-muted-foreground">
              Every order follows controlled states, balance calculations and buyer-seller
              visibility.
            </p>
          </div>
        </div>

        <!-- LOOP RETURN -->
        <div
          class="border-t border-border bg-gradient-to-r from-primary/10 via-primary/5 to-transparent px-6 py-5 sm:px-8"
        >
          <div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
            <div class="flex items-center gap-3">
              <div
                class="flex h-10 w-10 items-center justify-center rounded-xl bg-primary text-primary-foreground"
              >
                <Boxes class="h-4 w-4" />
              </div>

              <div>
                <p class="text-sm font-semibold">The information returns to the business</p>

                <p class="mt-1 text-xs text-muted-foreground">
                  Dashboard, inventory and commercial activity remain synchronized.
                </p>
              </div>
            </div>

            <span
              class="inline-flex self-start rounded-full border border-primary/20 bg-primary/10 px-3 py-1.5 text-xs font-semibold text-primary sm:self-auto"
            >
              Closed operational loop
            </span>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
