<script setup lang="ts">
import { computed } from 'vue';
import { useRouter } from 'vue-router';
import { storeToRefs } from 'pinia';

import {
  AlertTriangle,
  ArrowRight,
  Boxes,
  Building2,
  CircleDollarSign,
  PackageCheck,
  ReceiptText,
  ShoppingCart,
  TrendingUp,
  WalletCards,
} from 'lucide-vue-next';

import { Button } from '@/components/ui/button';

import { useAuthStore } from '@/modules/auth/stores/auth.store';

const router = useRouter();

const auth = useAuthStore();
const { user } = storeToRefs(auth);

const isBusinessUser = computed(() => {
  return Boolean(user.value?.company_id);
});

function openBusinessExperience() {
  if (isBusinessUser.value) {
    router.push('/erp/dashboard');
    return;
  }

  router.push('/register/business');
}
</script>

<template>
  <section class="relative overflow-hidden border-b border-border bg-zinc-950 text-zinc-50">
    <!-- BACKGROUND -->
    <div
      class="pointer-events-none absolute -left-48 top-12 h-[520px] w-[520px] rounded-full bg-primary/20 blur-3xl"
    />

    <div
      class="pointer-events-none absolute -right-48 bottom-0 h-[520px] w-[520px] rounded-full bg-violet-500/15 blur-3xl"
    />

    <div
      class="pointer-events-none absolute inset-0 bg-[linear-gradient(to_right,rgba(255,255,255,0.035)_1px,transparent_1px),linear-gradient(to_bottom,rgba(255,255,255,0.035)_1px,transparent_1px)] bg-[size:64px_64px]"
    />

    <div
      class="relative mx-auto grid w-full max-w-[1500px] items-center gap-16 px-6 py-24 lg:grid-cols-[0.88fr_1.12fr] lg:px-8 lg:py-28"
    >
      <!-- COPY -->
      <div>
        <p class="text-sm font-semibold text-primary">Business Workspace</p>

        <h2 class="mt-4 text-4xl font-semibold tracking-[-0.055em] text-white sm:text-5xl">
          Your company deserves
          <span class="text-zinc-400"> its own operating space. </span>
        </h2>

        <p class="mt-6 max-w-xl text-base leading-7 text-zinc-400">
          Nexora gives every business a workspace shaped around its identity, products, financial
          activity and commercial operations.
        </p>

        <div class="mt-8 space-y-4">
          <div class="flex items-start gap-3">
            <div
              class="mt-0.5 flex h-8 w-8 shrink-0 items-center justify-center rounded-lg bg-primary/15"
            >
              <Building2 class="h-4 w-4 text-primary" />
            </div>

            <div>
              <p class="text-sm font-semibold text-white">Business identity</p>

              <p class="mt-1 text-sm leading-6 text-zinc-400">
                Company profile, category, branding, logo and cover image throughout the workspace.
              </p>
            </div>
          </div>

          <div class="flex items-start gap-3">
            <div
              class="mt-0.5 flex h-8 w-8 shrink-0 items-center justify-center rounded-lg bg-blue-500/15"
            >
              <Boxes class="h-4 w-4 text-blue-400" />
            </div>

            <div>
              <p class="text-sm font-semibold text-white">Operational control</p>

              <p class="mt-1 text-sm leading-6 text-zinc-400">
                Products, stock, purchases, sales and payments managed from one coherent
                environment.
              </p>
            </div>
          </div>

          <div class="flex items-start gap-3">
            <div
              class="mt-0.5 flex h-8 w-8 shrink-0 items-center justify-center rounded-lg bg-emerald-500/15"
            >
              <TrendingUp class="h-4 w-4 text-emerald-400" />
            </div>

            <div>
              <p class="text-sm font-semibold text-white">Business visibility</p>

              <p class="mt-1 text-sm leading-6 text-zinc-400">
                Financial summaries, operational alerts and recent activity translated into clear
                decisions.
              </p>
            </div>
          </div>
        </div>

        <Button size="lg" class="mt-9 rounded-full px-7" @click="openBusinessExperience">
          {{ isBusinessUser ? 'Open business workspace' : 'Create a business workspace' }}

          <ArrowRight class="ml-2 h-4 w-4" />
        </Button>
      </div>

      <!-- WORKSPACE COMPOSITION -->
      <div class="relative min-h-[620px]">
        <div class="pointer-events-none absolute inset-12 rounded-full bg-primary/20 blur-3xl" />

        <!-- MAIN DASHBOARD -->
        <div
          class="absolute inset-x-0 top-4 overflow-hidden rounded-[2rem] border border-white/10 bg-zinc-900/95 shadow-2xl backdrop-blur-xl"
        >
          <!-- TOPBAR -->
          <div class="flex items-center justify-between border-b border-white/10 px-5 py-4">
            <div class="flex items-center gap-3">
              <div
                class="nexora-gradient nexora-glow flex h-10 w-10 items-center justify-center rounded-xl bg-primary font-bold text-primary-foreground"
              >
                N
              </div>

              <div>
                <p class="text-sm font-semibold text-white">Nexora Operations</p>

                <p class="mt-0.5 text-[11px] text-zinc-500">Business overview</p>
              </div>
            </div>

            <div
              class="flex items-center gap-2 rounded-full border border-white/10 bg-white/5 px-3 py-1.5"
            >
              <span class="h-2 w-2 rounded-full bg-emerald-500" />

              <span class="text-[11px] font-medium text-zinc-300"> Workspace healthy </span>
            </div>
          </div>

          <div class="p-5">
            <!-- COMPANY CONTEXT -->
            <div
              class="relative overflow-hidden rounded-2xl border border-white/10 bg-gradient-to-r from-primary/20 via-zinc-900 to-zinc-900 p-5"
            >
              <div
                class="pointer-events-none absolute -right-16 -top-16 h-44 w-44 rounded-full bg-violet-500/20 blur-3xl"
              />

              <div class="relative flex items-center justify-between gap-5">
                <div class="flex min-w-0 items-center gap-4">
                  <div
                    class="flex h-14 w-14 shrink-0 items-center justify-center rounded-2xl border border-white/10 bg-white/10"
                  >
                    <Building2 class="h-6 w-6 text-primary" />
                  </div>

                  <div class="min-w-0">
                    <p class="text-xs font-medium text-zinc-400">Main workspace</p>

                    <p class="mt-1 truncate text-lg font-semibold text-white">Constante Labs</p>

                    <p class="mt-1 text-xs text-zinc-500">Technology & digital operations</p>
                  </div>
                </div>

                <div class="hidden text-right sm:block">
                  <p class="text-xs text-zinc-500">Operational status</p>

                  <p class="mt-1 text-sm font-semibold text-emerald-400">All systems active</p>
                </div>
              </div>
            </div>

            <!-- KPIS -->
            <div class="mt-4 grid grid-cols-2 gap-3 sm:grid-cols-4">
              <div class="rounded-2xl border border-white/10 bg-white/[0.035] p-4">
                <CircleDollarSign class="h-4 w-4 text-emerald-400" />

                <p class="mt-4 text-[10px] uppercase tracking-wide text-zinc-500">Sales</p>

                <p class="mt-1 text-lg font-semibold text-white">$8.4M</p>
              </div>

              <div class="rounded-2xl border border-white/10 bg-white/[0.035] p-4">
                <ShoppingCart class="h-4 w-4 text-blue-400" />

                <p class="mt-4 text-[10px] uppercase tracking-wide text-zinc-500">Purchases</p>

                <p class="mt-1 text-lg font-semibold text-white">$2.1M</p>
              </div>

              <div class="rounded-2xl border border-white/10 bg-white/[0.035] p-4">
                <Boxes class="h-4 w-4 text-primary" />

                <p class="mt-4 text-[10px] uppercase tracking-wide text-zinc-500">Inventory</p>

                <p class="mt-1 text-lg font-semibold text-white">1,248</p>
              </div>

              <div class="rounded-2xl border border-white/10 bg-white/[0.035] p-4">
                <WalletCards class="h-4 w-4 text-amber-400" />

                <p class="mt-4 text-[10px] uppercase tracking-wide text-zinc-500">Outstanding</p>

                <p class="mt-1 text-lg font-semibold text-white">$640K</p>
              </div>
            </div>

            <!-- LOWER GRID -->
            <div class="mt-4 grid gap-4 md:grid-cols-[1.15fr_0.85fr]">
              <!-- ACTIVITY -->
              <div class="rounded-2xl border border-white/10 bg-white/[0.025] p-4">
                <div class="flex items-center justify-between">
                  <div>
                    <p class="text-sm font-semibold text-white">Recent activity</p>

                    <p class="mt-1 text-[11px] text-zinc-500">Commercial operations</p>
                  </div>

                  <ReceiptText class="h-4 w-4 text-primary" />
                </div>

                <div class="mt-4 space-y-3">
                  <div
                    class="flex items-center justify-between gap-3 rounded-xl border border-white/5 bg-white/[0.025] p-3"
                  >
                    <div>
                      <p class="text-xs font-semibold text-zinc-200">Mercado Norte</p>

                      <p class="mt-1 text-[10px] text-zinc-500">Sales invoice #NX-1048</p>
                    </div>

                    <span
                      class="rounded-full bg-amber-500/10 px-2 py-1 text-[10px] font-semibold text-amber-400"
                    >
                      Pending
                    </span>
                  </div>

                  <div
                    class="flex items-center justify-between gap-3 rounded-xl border border-white/5 bg-white/[0.025] p-3"
                  >
                    <div>
                      <p class="text-xs font-semibold text-zinc-200">TechNova</p>

                      <p class="mt-1 text-[10px] text-zinc-500">Purchase payment</p>
                    </div>

                    <span class="text-xs font-semibold text-emerald-400"> $184K </span>
                  </div>

                  <div
                    class="flex items-center justify-between gap-3 rounded-xl border border-white/5 bg-white/[0.025] p-3"
                  >
                    <div>
                      <p class="text-xs font-semibold text-zinc-200">Store order</p>

                      <p class="mt-1 text-[10px] text-zinc-500">Consumer commerce</p>
                    </div>

                    <span
                      class="rounded-full bg-emerald-500/10 px-2 py-1 text-[10px] font-semibold text-emerald-400"
                    >
                      Paid
                    </span>
                  </div>
                </div>
              </div>

              <!-- ATTENTION -->
              <div class="rounded-2xl border border-white/10 bg-white/[0.025] p-4">
                <div class="flex items-center gap-3">
                  <div class="flex h-9 w-9 items-center justify-center rounded-xl bg-amber-500/10">
                    <AlertTriangle class="h-4 w-4 text-amber-400" />
                  </div>

                  <div>
                    <p class="text-sm font-semibold text-white">Attention center</p>

                    <p class="mt-1 text-[11px] text-zinc-500">Items requiring action</p>
                  </div>
                </div>

                <div class="mt-4 space-y-3">
                  <div class="flex items-center justify-between border-b border-white/10 pb-3">
                    <span class="text-xs text-zinc-400"> Pending sales </span>

                    <span class="text-sm font-semibold text-white"> 12 </span>
                  </div>

                  <div class="flex items-center justify-between border-b border-white/10 pb-3">
                    <span class="text-xs text-zinc-400"> Partial payments </span>

                    <span class="text-sm font-semibold text-white"> 4 </span>
                  </div>

                  <div class="flex items-center justify-between">
                    <span class="text-xs text-zinc-400"> Low stock </span>

                    <span class="text-sm font-semibold text-amber-400"> 7 </span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- FLOATING INVENTORY CARD -->
        <div
          class="absolute bottom-0 left-0 z-20 w-[230px] rounded-[1.5rem] border border-white/10 bg-zinc-900 p-4 shadow-2xl sm:w-[260px]"
        >
          <div class="flex items-center justify-between">
            <div class="flex h-10 w-10 items-center justify-center rounded-xl bg-primary/15">
              <PackageCheck class="h-5 w-5 text-primary" />
            </div>

            <span
              class="rounded-full bg-emerald-500/10 px-2.5 py-1 text-[10px] font-semibold text-emerald-400"
            >
              Synchronized
            </span>
          </div>

          <p class="mt-4 text-xs text-zinc-500">Available inventory</p>

          <p class="mt-1 text-2xl font-semibold text-white">4,820 units</p>

          <div class="mt-4 h-1.5 overflow-hidden rounded-full bg-white/10">
            <div class="h-full w-[78%] rounded-full bg-primary" />
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
