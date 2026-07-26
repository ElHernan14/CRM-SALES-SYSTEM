<script setup lang="ts">
import { Boxes, Building2, Fingerprint, Network, ReceiptText, ShieldCheck } from 'lucide-vue-next';

const foundations = [
  {
    title: 'Multi-tenant by design',
    description:
      'Every business operates inside its own isolated context, with company-scoped data and ownership validation.',
    icon: Building2,
    accentClass: 'bg-primary/10 text-primary',
  },

  {
    title: 'Role-based access',
    description:
      'Dynamic roles and permissions determine what each user can see and perform throughout the platform.',
    icon: Fingerprint,
    accentClass: 'bg-violet-500/10 text-violet-600 dark:text-violet-400',
  },

  {
    title: 'Connected inventory',
    description:
      'Stock, reserved units and Store availability remain synchronized with real business operations.',
    icon: Boxes,
    accentClass: 'bg-blue-500/10 text-blue-600 dark:text-blue-400',
  },

  {
    title: 'Transactional workflows',
    description:
      'Critical operations use consistent validations and atomic transactions to prevent partial or invalid states.',
    icon: ShieldCheck,
    accentClass: 'bg-emerald-500/10 text-emerald-600 dark:text-emerald-400',
  },

  {
    title: 'Unified invoice engine',
    description:
      'B2B sales, organizational purchases and personal Store orders share one controlled financial lifecycle.',
    icon: ReceiptText,
    accentClass: 'bg-amber-500/10 text-amber-600 dark:text-amber-400',
  },

  {
    title: 'Connected commerce',
    description:
      'ERP, Marketplace and Store operate as different experiences over the same business ecosystem.',
    icon: Network,
    accentClass: 'bg-cyan-500/10 text-cyan-600 dark:text-cyan-400',
  },
] as const;
</script>

<template>
  <section
    id="nexora-product"
    class="relative scroll-mt-24 overflow-hidden border-b border-border bg-background"
  >
    <div
      class="pointer-events-none absolute -left-48 bottom-0 h-[460px] w-[460px] rounded-full bg-primary/5 blur-3xl"
    />

    <div class="mx-auto w-full max-w-[1500px] px-6 py-24 lg:px-8 lg:py-28">
      <!-- HEADING -->
      <div class="flex flex-col gap-8 lg:flex-row lg:items-end lg:justify-between">
        <div class="max-w-3xl">
          <p class="text-sm font-semibold text-primary">Built for real operations</p>

          <h2 class="mt-4 text-4xl font-semibold tracking-[-0.05em] text-foreground sm:text-5xl">
            More than a collection of modules.
            <span class="text-muted-foreground"> One operational foundation. </span>
          </h2>
        </div>

        <p class="max-w-xl text-base leading-7 text-muted-foreground">
          Nexora connects identity, inventory, commerce, invoices and payments through the same
          rules and business context.
        </p>
      </div>

      <!-- FOUNDATION GRID -->
      <div class="mt-14 grid gap-4 md:grid-cols-2 xl:grid-cols-3">
        <article
          v-for="foundation in foundations"
          :key="foundation.title"
          class="group relative overflow-hidden rounded-[1.75rem] border border-border bg-card p-6 shadow-sm transition duration-300 hover:-translate-y-1 hover:border-primary/25 hover:shadow-xl"
        >
          <div
            class="pointer-events-none absolute -right-12 -top-12 h-36 w-36 rounded-full bg-primary/5 blur-3xl transition group-hover:bg-primary/10"
          />

          <div class="relative">
            <div
              class="flex h-12 w-12 items-center justify-center rounded-2xl"
              :class="foundation.accentClass"
            >
              <component :is="foundation.icon" class="h-5 w-5" />
            </div>

            <h3 class="mt-6 text-lg font-semibold tracking-tight text-foreground">
              {{ foundation.title }}
            </h3>

            <p class="mt-3 text-sm leading-6 text-muted-foreground">
              {{ foundation.description }}
            </p>
          </div>
        </article>
      </div>

      <!-- ARCHITECTURE STRIP -->
      <div class="mt-8 overflow-hidden rounded-[1.75rem] border border-border bg-muted/15">
        <div class="grid divide-y divide-border md:grid-cols-4 md:divide-x md:divide-y-0">
          <div class="p-5">
            <p class="text-xs font-semibold uppercase tracking-[0.14em] text-muted-foreground">
              Context
            </p>

            <p class="mt-2 text-base font-semibold">Tenant scoped</p>
          </div>

          <div class="p-5">
            <p class="text-xs font-semibold uppercase tracking-[0.14em] text-muted-foreground">
              Access
            </p>

            <p class="mt-2 text-base font-semibold">Roles & permissions</p>
          </div>

          <div class="p-5">
            <p class="text-xs font-semibold uppercase tracking-[0.14em] text-muted-foreground">
              Commerce
            </p>

            <p class="mt-2 text-base font-semibold">B2B & B2C</p>
          </div>

          <div class="p-5">
            <p class="text-xs font-semibold uppercase tracking-[0.14em] text-muted-foreground">
              Finance
            </p>

            <p class="mt-2 text-base font-semibold">Stateful invoices</p>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
