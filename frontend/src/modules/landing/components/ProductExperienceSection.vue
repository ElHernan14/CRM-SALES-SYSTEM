<script setup lang="ts">
import {
  ArrowRight,
  Building2,
  CheckCircle2,
  LayoutDashboard,
  ReceiptText,
  ShoppingBag,
  Store,
  UserRound,
} from 'lucide-vue-next';

const experiences = [
  {
    label: 'Business',
    title: 'ERP Workspace',
    description: 'Employees operate products, sales, purchases, inventory and company identity.',
    icon: LayoutDashboard,
    accentClass: 'bg-primary text-primary-foreground',
    features: ['Company context', 'Operational control', 'Financial overview'],
  },

  {
    label: 'Commerce',
    title: 'Nexora Store',
    description:
      'Customers discover businesses, products and seller-specific purchase experiences.',
    icon: Store,
    accentClass: 'bg-violet-500 text-white',
    features: ['Public discovery', 'Multi-seller cart', 'Connected inventory'],
  },

  {
    label: 'Personal',
    title: 'Customer Account',
    description: 'Individual users manage profile information, purchases and payment history.',
    icon: UserRound,
    accentClass: 'bg-emerald-500 text-white',
    features: ['Personal identity', 'Purchase history', 'Payment tracking'],
  },
] as const;
</script>

<template>
  <section class="relative overflow-hidden border-b border-border bg-background">
    <div
      class="pointer-events-none absolute -right-48 top-0 h-[520px] w-[520px] rounded-full bg-violet-500/5 blur-3xl"
    />

    <div class="relative mx-auto w-full max-w-[1500px] px-6 py-24 lg:px-8 lg:py-28">
      <div class="mx-auto max-w-4xl text-center">
        <p class="text-sm font-semibold text-primary">One product experience</p>

        <h2 class="mt-4 text-4xl font-semibold tracking-[-0.055em] sm:text-5xl">
          Different users.
          <span class="text-muted-foreground"> One coherent product. </span>
        </h2>

        <p class="mx-auto mt-5 max-w-2xl text-base leading-7 text-muted-foreground">
          Every surface adapts to its user without losing Nexora’s identity, navigation language or
          commercial context.
        </p>
      </div>

      <div class="mt-16 grid gap-5 lg:grid-cols-3">
        <article
          v-for="experience in experiences"
          :key="experience.title"
          class="nexora-card-interactive nexora-surface group relative flex min-h-[430px] flex-col overflow-hidden rounded-[2rem] border border-border bg-card p-6 shadow-sm transition duration-300 hover:-translate-y-1 hover:border-primary/25 hover:shadow-2xl sm:p-7"
        >
          <div
            class="pointer-events-none absolute -right-20 -top-20 h-52 w-52 rounded-full bg-primary/5 blur-3xl transition group-hover:bg-primary/10"
          />

          <div class="relative">
            <div
              class="flex h-14 w-14 items-center justify-center rounded-2xl shadow-lg"
              :class="experience.accentClass"
            >
              <component :is="experience.icon" class="h-6 w-6" />
            </div>

            <p class="mt-7 text-xs font-semibold uppercase tracking-[0.16em] text-primary">
              {{ experience.label }}
            </p>

            <h3 class="mt-2 text-2xl font-semibold tracking-[-0.035em]">
              {{ experience.title }}
            </h3>

            <p class="mt-4 text-sm leading-7 text-muted-foreground">
              {{ experience.description }}
            </p>
          </div>

          <!-- MINI SURFACE -->
          <div
            class="relative mt-7 overflow-hidden rounded-[1.5rem] border border-border bg-muted/15 p-4"
          >
            <template v-if="experience.label === 'Business'">
              <div class="flex items-center justify-between">
                <div class="flex items-center gap-3">
                  <div class="flex h-9 w-9 items-center justify-center rounded-xl bg-primary/10">
                    <Building2 class="h-4 w-4 text-primary" />
                  </div>

                  <div>
                    <p class="text-xs font-semibold">Business overview</p>

                    <p class="mt-1 text-[10px] text-muted-foreground">Main workspace</p>
                  </div>
                </div>

                <span class="h-2 w-2 rounded-full bg-emerald-500" />
              </div>

              <div class="mt-4 grid grid-cols-2 gap-2">
                <div class="rounded-xl border border-border bg-card p-3">
                  <p class="text-[10px] text-muted-foreground">Sales</p>

                  <p class="mt-1 text-sm font-semibold">$8,4 M</p>
                </div>

                <div class="rounded-xl border border-border bg-card p-3">
                  <p class="text-[10px] text-muted-foreground">Inventory</p>

                  <p class="mt-1 text-sm font-semibold">1.248</p>
                </div>
              </div>
            </template>

            <template v-else-if="experience.label === 'Commerce'">
              <div
                class="h-24 rounded-xl bg-gradient-to-br from-violet-500/20 via-primary/10 to-background"
              />

              <div class="mt-3 flex items-end justify-between">
                <div>
                  <p class="text-[10px] text-muted-foreground">Connected seller</p>

                  <p class="mt-1 text-sm font-semibold">TechNova</p>
                </div>

                <div
                  class="flex h-9 w-9 items-center justify-center rounded-full bg-primary text-primary-foreground"
                >
                  <ShoppingBag class="h-4 w-4" />
                </div>
              </div>
            </template>

            <template v-else>
              <div class="flex items-center gap-3">
                <div
                  class="flex h-10 w-10 items-center justify-center rounded-full bg-emerald-500/10"
                >
                  <UserRound class="h-5 w-5 text-emerald-600 dark:text-emerald-400" />
                </div>

                <div>
                  <p class="text-xs font-semibold">Personal account</p>

                  <p class="mt-1 text-[10px] text-muted-foreground">Purchase activity</p>
                </div>
              </div>

              <div
                class="mt-4 flex items-center justify-between rounded-xl border border-border bg-card p-3"
              >
                <div class="flex items-center gap-2">
                  <ReceiptText class="h-4 w-4 text-primary" />

                  <span class="text-xs font-medium"> #NX-1048 </span>
                </div>

                <span
                  class="rounded-full bg-amber-500/10 px-2 py-1 text-[10px] font-semibold text-amber-600 dark:text-amber-400"
                >
                  Pending
                </span>
              </div>
            </template>
          </div>

          <div class="relative mt-6 space-y-3">
            <div
              v-for="feature in experience.features"
              :key="feature"
              class="flex items-center gap-3"
            >
              <CheckCircle2 class="h-4 w-4 shrink-0 text-primary" />

              <span class="text-sm text-muted-foreground">
                {{ feature }}
              </span>
            </div>
          </div>

          <div
            class="relative mt-auto flex items-center gap-2 border-t border-border pt-5 text-sm font-semibold text-primary"
          >
            Part of the Nexora ecosystem

            <ArrowRight class="h-4 w-4" />
          </div>
        </article>
      </div>

      <div
        class="mt-8 flex flex-col gap-5 rounded-[1.75rem] border border-primary/20 bg-primary/5 p-6 sm:flex-row sm:items-center sm:justify-between"
      >
        <div class="flex items-start gap-4">
          <div
            class="flex h-11 w-11 shrink-0 items-center justify-center rounded-xl bg-primary text-primary-foreground"
          >
            <Store class="h-5 w-5" />
          </div>

          <div>
            <p class="text-sm font-semibold">Identity travels across the product.</p>

            <p class="mt-1 text-sm leading-6 text-muted-foreground">
              A company configured inside ERP appears with the same identity throughout Businesses,
              Catalog, Cart and Purchases.
            </p>
          </div>
        </div>

        <span
          class="self-start rounded-full border border-primary/20 bg-background px-3 py-1.5 text-xs font-semibold text-primary sm:self-auto"
        >
          One connected ecosystem
        </span>
      </div>
    </div>
  </section>
</template>
