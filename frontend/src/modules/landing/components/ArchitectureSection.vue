<script setup lang="ts">
import {
  Building2,
  Database,
  Fingerprint,
  KeyRound,
  Network,
  Route,
  ShieldCheck,
  ShoppingBag,
  Split,
  Store,
  UserCheck,
} from 'lucide-vue-next';

const architecturePrinciples = [
  {
    title: 'Tenant-scoped context',
    description:
      'Every authenticated request carries the company, client, roles and permissions required to resolve its business context.',
    icon: Building2,
    iconClass: 'bg-primary/10 text-primary',
  },

  {
    title: 'Ownership validation',
    description:
      'Invoices, products and commercial operations are validated against the buyer and seller participating in the transaction.',
    icon: UserCheck,
    iconClass: 'bg-emerald-500/10 text-emerald-600 dark:text-emerald-400',
  },

  {
    title: 'Dynamic authorization',
    description:
      'Roles and permissions control access independently from the interface, keeping business rules authoritative.',
    icon: KeyRound,
    iconClass: 'bg-violet-500/10 text-violet-600 dark:text-violet-400',
  },

  {
    title: 'Separated purchase origins',
    description:
      'ERP and Store purchases remain explicitly separated while sharing the same invoice and payment infrastructure.',
    icon: Split,
    iconClass: 'bg-blue-500/10 text-blue-600 dark:text-blue-400',
  },
] as const;
</script>

<template>
  <section
    id="nexora-architecture"
    class="relative scroll-mt-24 overflow-hidden border-b border-border bg-zinc-950 text-zinc-50"
  >
    <div
      class="pointer-events-none absolute -left-48 top-0 h-[560px] w-[560px] rounded-full bg-primary/20 blur-3xl"
    />

    <div
      class="pointer-events-none absolute -right-52 bottom-0 h-[560px] w-[560px] rounded-full bg-blue-500/15 blur-3xl"
    />

    <div
      class="pointer-events-none absolute inset-0 bg-[linear-gradient(to_right,rgba(255,255,255,0.03)_1px,transparent_1px),linear-gradient(to_bottom,rgba(255,255,255,0.03)_1px,transparent_1px)] bg-[size:72px_72px]"
    />

    <div class="relative mx-auto w-full max-w-[1500px] px-6 py-24 lg:px-8 lg:py-28">
      <!-- HEADING -->
      <div class="flex flex-col gap-8 lg:flex-row lg:items-end lg:justify-between">
        <div class="max-w-3xl">
          <p class="text-sm font-semibold text-primary">Multi-tenant architecture</p>

          <h2 class="mt-4 text-4xl font-semibold tracking-[-0.055em] text-white sm:text-5xl">
            One platform.
            <span class="text-zinc-400"> Isolated business contexts. </span>
          </h2>
        </div>

        <p class="max-w-xl text-base leading-7 text-zinc-400">
          Nexora connects businesses without mixing their operations. Identity, ownership and
          authorization remain part of every commercial decision.
        </p>
      </div>

      <div class="mt-16 grid items-stretch gap-6 xl:grid-cols-[0.9fr_1.1fr]">
        <!-- CONTEXT VISUAL -->
        <div
          class="relative overflow-hidden rounded-[2rem] border border-white/10 bg-zinc-900/90 p-6 shadow-2xl backdrop-blur-xl sm:p-8"
        >
          <div
            class="pointer-events-none absolute -right-20 -top-20 h-52 w-52 rounded-full bg-primary/20 blur-3xl"
          />

          <div class="relative">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-xs font-semibold uppercase tracking-[0.16em] text-primary">
                  Tenant Context
                </p>

                <h3 class="mt-2 text-2xl font-semibold text-white">Request identity</h3>
              </div>

              <Fingerprint class="h-6 w-6 text-primary" />
            </div>

            <div class="mt-7 rounded-[1.5rem] border border-white/10 bg-black/20 p-5">
              <div class="space-y-4 font-mono text-xs">
                <div class="flex items-center justify-between border-b border-white/10 pb-3">
                  <span class="text-zinc-500">user_id</span>
                  <span class="text-zinc-200">42</span>
                </div>

                <div class="flex items-center justify-between border-b border-white/10 pb-3">
                  <span class="text-zinc-500">company_id</span>
                  <span class="text-primary">15</span>
                </div>

                <div class="flex items-center justify-between border-b border-white/10 pb-3">
                  <span class="text-zinc-500">client_id</span>
                  <span class="text-zinc-200">28</span>
                </div>

                <div class="flex items-center justify-between border-b border-white/10 pb-3">
                  <span class="text-zinc-500">role</span>
                  <span class="text-blue-400">company_user</span>
                </div>

                <div class="flex items-start justify-between gap-5">
                  <span class="text-zinc-500">permissions</span>

                  <div class="flex max-w-52 flex-wrap justify-end gap-1.5">
                    <span
                      class="rounded-md border border-white/10 bg-white/5 px-2 py-1 text-[10px] text-zinc-300"
                    >
                      product:read
                    </span>

                    <span
                      class="rounded-md border border-white/10 bg-white/5 px-2 py-1 text-[10px] text-zinc-300"
                    >
                      invoice:pay
                    </span>

                    <span
                      class="rounded-md border border-white/10 bg-white/5 px-2 py-1 text-[10px] text-zinc-300"
                    >
                      company:update
                    </span>
                  </div>
                </div>
              </div>
            </div>

            <!-- VALIDATION FLOW -->
            <div class="mt-6 space-y-3">
              <div
                class="flex items-center gap-3 rounded-xl border border-white/10 bg-white/[0.035] p-4"
              >
                <div
                  class="flex h-9 w-9 shrink-0 items-center justify-center rounded-xl bg-primary/15"
                >
                  <Fingerprint class="h-4 w-4 text-primary" />
                </div>

                <div class="min-w-0">
                  <p class="text-sm font-semibold text-white">Resolve identity</p>

                  <p class="mt-1 text-xs text-zinc-500">
                    Authentication creates the request context.
                  </p>
                </div>
              </div>

              <div
                class="flex items-center gap-3 rounded-xl border border-white/10 bg-white/[0.035] p-4"
              >
                <div
                  class="flex h-9 w-9 shrink-0 items-center justify-center rounded-xl bg-violet-500/15"
                >
                  <KeyRound class="h-4 w-4 text-violet-400" />
                </div>

                <div class="min-w-0">
                  <p class="text-sm font-semibold text-white">Validate permission</p>

                  <p class="mt-1 text-xs text-zinc-500">
                    The operation must be explicitly authorized.
                  </p>
                </div>
              </div>

              <div
                class="flex items-center gap-3 rounded-xl border border-emerald-500/20 bg-emerald-500/5 p-4"
              >
                <div
                  class="flex h-9 w-9 shrink-0 items-center justify-center rounded-xl bg-emerald-500/15"
                >
                  <ShieldCheck class="h-4 w-4 text-emerald-400" />
                </div>

                <div class="min-w-0">
                  <p class="text-sm font-semibold text-white">Enforce ownership</p>

                  <p class="mt-1 text-xs text-zinc-500">
                    Tenant and commercial participation are confirmed.
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- PLATFORM MAP -->
        <div
          class="relative overflow-hidden rounded-[2rem] border border-white/10 bg-zinc-900/90 p-6 shadow-2xl backdrop-blur-xl sm:p-8"
        >
          <div
            class="pointer-events-none absolute -left-20 -top-20 h-52 w-52 rounded-full bg-blue-500/15 blur-3xl"
          />

          <div class="relative">
            <div>
              <p class="text-xs font-semibold uppercase tracking-[0.16em] text-blue-400">
                Platform boundaries
              </p>

              <h3 class="mt-2 text-2xl font-semibold text-white">
                Shared engine, distinct experiences
              </h3>

              <p class="mt-3 max-w-2xl text-sm leading-6 text-zinc-400">
                ERP and Store use different access patterns and purchase origins without duplicating
                financial logic.
              </p>
            </div>

            <!-- EXPERIENCES -->
            <div class="mt-8 grid gap-4 sm:grid-cols-2">
              <div class="rounded-[1.5rem] border border-white/10 bg-white/[0.035] p-5">
                <div class="flex items-center justify-between">
                  <div class="flex h-11 w-11 items-center justify-center rounded-xl bg-primary/15">
                    <Building2 class="h-5 w-5 text-primary" />
                  </div>

                  <span
                    class="rounded-full border border-white/10 bg-white/5 px-2.5 py-1 font-mono text-[10px] text-zinc-400"
                  >
                    source = erp
                  </span>
                </div>

                <p class="mt-5 text-base font-semibold text-white">Business Workspace</p>

                <p class="mt-2 text-sm leading-6 text-zinc-400">
                  Organizational sales and purchases operate at company level.
                </p>
              </div>

              <div class="rounded-[1.5rem] border border-white/10 bg-white/[0.035] p-5">
                <div class="flex items-center justify-between">
                  <div
                    class="flex h-11 w-11 items-center justify-center rounded-xl bg-violet-500/15"
                  >
                    <Store class="h-5 w-5 text-violet-400" />
                  </div>

                  <span
                    class="rounded-full border border-white/10 bg-white/5 px-2.5 py-1 font-mono text-[10px] text-zinc-400"
                  >
                    source = store
                  </span>
                </div>

                <p class="mt-5 text-base font-semibold text-white">Personal Store</p>

                <p class="mt-2 text-sm leading-6 text-zinc-400">
                  Consumer carts and purchases remain scoped to the authenticated client.
                </p>
              </div>
            </div>

            <!-- SHARED DOMAIN -->
            <div class="mt-5 rounded-[1.5rem] border border-primary/20 bg-primary/5 p-5">
              <div class="flex flex-col gap-5 sm:flex-row sm:items-center sm:justify-between">
                <div class="flex items-center gap-4">
                  <div
                    class="flex h-12 w-12 shrink-0 items-center justify-center rounded-2xl bg-primary text-primary-foreground"
                  >
                    <Network class="h-5 w-5" />
                  </div>

                  <div>
                    <p class="text-sm font-semibold text-white">Shared commercial domain</p>

                    <p class="mt-1 text-xs text-zinc-400">
                      Products · inventory · invoices · payments
                    </p>
                  </div>
                </div>

                <span
                  class="self-start rounded-full border border-emerald-500/20 bg-emerald-500/10 px-3 py-1.5 text-xs font-semibold text-emerald-400 sm:self-auto"
                >
                  Consistent rules
                </span>
              </div>
            </div>

            <!-- REQUEST ROUTES -->
            <div class="mt-6">
              <p class="text-xs font-semibold uppercase tracking-[0.14em] text-zinc-500">
                Route resolution
              </p>

              <div class="mt-4 space-y-3">
                <div
                  class="flex items-center gap-4 rounded-xl border border-white/10 bg-black/20 p-4"
                >
                  <ShoppingBag class="h-4 w-4 shrink-0 text-violet-400" />

                  <span class="font-mono text-xs text-zinc-400"> /store/purchases </span>

                  <Route class="ml-auto h-4 w-4 text-zinc-600" />

                  <span class="text-xs font-semibold text-zinc-200"> client context </span>
                </div>

                <div
                  class="flex items-center gap-4 rounded-xl border border-white/10 bg-black/20 p-4"
                >
                  <Building2 class="h-4 w-4 shrink-0 text-primary" />

                  <span class="font-mono text-xs text-zinc-400"> /company/purchases </span>

                  <Route class="ml-auto h-4 w-4 text-zinc-600" />

                  <span class="text-xs font-semibold text-zinc-200"> company context </span>
                </div>

                <div
                  class="flex items-center gap-4 rounded-xl border border-white/10 bg-black/20 p-4"
                >
                  <Database class="h-4 w-4 shrink-0 text-emerald-400" />

                  <span class="font-mono text-xs text-zinc-400"> invoice </span>

                  <Route class="ml-auto h-4 w-4 text-zinc-600" />

                  <span class="text-xs font-semibold text-zinc-200"> shared engine </span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- PRINCIPLES -->
      <div class="mt-8 grid gap-4 md:grid-cols-2 xl:grid-cols-4">
        <article
          v-for="principle in architecturePrinciples"
          :key="principle.title"
          class="rounded-[1.5rem] border border-white/10 bg-white/[0.035] p-5"
        >
          <div
            class="flex h-11 w-11 items-center justify-center rounded-xl"
            :class="principle.iconClass"
          >
            <component :is="principle.icon" class="h-5 w-5" />
          </div>

          <h3 class="mt-5 text-base font-semibold text-white">
            {{ principle.title }}
          </h3>

          <p class="mt-2 text-sm leading-6 text-zinc-400">
            {{ principle.description }}
          </p>
        </article>
      </div>

      <div
        class="mt-8 flex flex-col gap-5 rounded-[1.75rem] border border-white/10 bg-white/[0.035] p-6 sm:flex-row sm:items-center sm:justify-between"
      >
        <div class="flex items-start gap-4">
          <div
            class="flex h-11 w-11 shrink-0 items-center justify-center rounded-xl bg-emerald-500/15"
          >
            <ShieldCheck class="h-5 w-5 text-emerald-400" />
          </div>

          <div>
            <p class="text-sm font-semibold text-white">Isolation is part of the business model.</p>

            <p class="mt-1 text-sm leading-6 text-zinc-400">
              Nexora does not rely on the interface to separate tenants; the context is enforced
              throughout the request lifecycle.
            </p>
          </div>
        </div>

        <span
          class="self-start rounded-full border border-emerald-500/20 bg-emerald-500/10 px-3 py-1.5 text-xs font-semibold text-emerald-400 sm:self-auto"
        >
          Tenant-aware platform
        </span>
      </div>
    </div>
  </section>
</template>
