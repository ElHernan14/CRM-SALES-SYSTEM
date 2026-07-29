<script setup lang="ts">
import { useRouter } from 'vue-router';

import {
  ArrowRight,
  BadgeCheck,
  Building2,
  Grid2X2,
  ImageIcon,
  PackageCheck,
  Search,
  ShoppingBag,
  Store,
} from 'lucide-vue-next';

import { Button } from '@/components/ui/button';

const router = useRouter();
</script>

<template>
  <section
    id="nexora-commerce"
    class="relative scroll-mt-24 overflow-hidden border-b border-border bg-background"
  >
    <div
      class="pointer-events-none absolute -right-48 top-12 h-[520px] w-[520px] rounded-full bg-primary/8 blur-3xl"
    />

    <div
      class="pointer-events-none absolute -left-52 bottom-0 h-[460px] w-[460px] rounded-full bg-violet-500/10 blur-3xl"
    />

    <div class="relative mx-auto w-full max-w-[1500px] px-6 py-24 lg:px-8 lg:py-28">
      <!-- HEADING -->
      <div class="flex flex-col gap-8 lg:flex-row lg:items-end lg:justify-between">
        <div class="max-w-3xl">
          <p class="text-sm font-semibold text-primary">Connected Commerce</p>

          <h2 class="mt-4 text-4xl font-semibold tracking-[-0.055em] text-foreground sm:text-5xl">
            Commerce that starts
            <span class="text-muted-foreground"> inside the business. </span>
          </h2>
        </div>

        <div class="max-w-xl">
          <p class="text-base leading-7 text-muted-foreground">
            Products published from the ERP become part of a real commerce network, connected to
            seller identity, inventory and financial workflows.
          </p>

          <div class="mt-5 flex flex-wrap gap-3">
            <Button class="rounded-full" @click="router.push('/store')">
              Explore Store

              <ArrowRight class="ml-2 h-4 w-4" />
            </Button>

            <Button
              variant="outline"
              class="rounded-full"
              @click="router.push('/store/businesses')"
            >
              Browse businesses
            </Button>
          </div>
        </div>
      </div>

      <!-- CONNECTED OPERATIONS VISUAL -->
      <div
        class="nexora-surface nexora-glow relative mt-12 overflow-hidden rounded-[2rem] border border-border bg-card shadow-xl sm:mt-14"
      >
        <div class="grid min-h-[340px] lg:grid-cols-[1.15fr_0.85fr]">
          <!-- IMAGE -->
          <div class="relative min-h-[280px] overflow-hidden sm:min-h-[340px]">
            <img
              src="/images/landing/connected-commerce.jpg"
              alt="Modern logistics and distribution operations connected through Nexora"
              class="absolute inset-0 h-full w-full object-cover"
              loading="lazy"
              decoding="async"
            />

            <div
              class="absolute inset-0 bg-gradient-to-t from-slate-950/80 via-slate-950/5 to-transparent lg:bg-gradient-to-r lg:from-transparent lg:via-transparent lg:to-card"
            />

            <div
              class="absolute inset-0 bg-gradient-to-br from-violet-950/20 via-transparent to-primary/15"
            />

            <div
              class="pointer-events-none absolute -left-24 -top-24 h-72 w-72 rounded-full bg-violet-500/20 blur-3xl"
            />

            <div class="absolute inset-0 ring-1 ring-inset ring-white/10" />

            <div
              class="absolute bottom-5 left-5 right-5 rounded-2xl border border-white/10 bg-slate-950/55 p-4 text-white shadow-xl backdrop-blur-md sm:bottom-6 sm:left-6 sm:right-auto sm:max-w-sm"
            >
              <div class="flex items-start gap-3">
                <div
                  class="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl border border-white/10 bg-white/10"
                >
                  <PackageCheck class="h-5 w-5 text-blue-300" />
                </div>

                <div>
                  <p class="text-xs font-medium text-slate-300">Operational inventory</p>

                  <p class="mt-1 text-sm font-semibold">
                    Business activity becomes customer-facing commerce.
                  </p>
                </div>
              </div>
            </div>
          </div>

          <!-- EXPLANATION -->
          <div class="relative flex flex-col justify-center p-6 sm:p-8 lg:p-10">
            <div
              class="pointer-events-none absolute -right-20 -top-20 h-56 w-56 rounded-full bg-primary/10 blur-3xl"
            />

            <div class="relative">
              <p class="text-xs font-semibold uppercase tracking-[0.16em] text-primary">
                From operations to commerce
              </p>

              <h3 class="mt-3 text-3xl font-semibold tracking-[-0.04em] text-foreground">
                Published once.
                <span class="text-muted-foreground"> Available everywhere. </span>
              </h3>

              <p class="mt-4 text-sm leading-7 text-muted-foreground">
                Products created inside a business workspace carry their seller identity, pricing
                and available inventory directly into the Nexora Store.
              </p>

              <div class="mt-7 space-y-4">
                <div class="flex items-start gap-3">
                  <div
                    class="flex h-9 w-9 shrink-0 items-center justify-center rounded-xl bg-primary/10"
                  >
                    <Building2 class="h-4 w-4 text-primary" />
                  </div>

                  <div>
                    <p class="text-sm font-semibold">Business-owned catalogs</p>

                    <p class="mt-1 text-xs leading-5 text-muted-foreground">
                      Every resource remains connected to the company responsible for it.
                    </p>
                  </div>
                </div>

                <div class="flex items-start gap-3">
                  <div
                    class="flex h-9 w-9 shrink-0 items-center justify-center rounded-xl bg-primary/10"
                  >
                    <PackageCheck class="h-4 w-4 text-primary" />
                  </div>

                  <div>
                    <p class="text-sm font-semibold">Operational availability</p>

                    <p class="mt-1 text-xs leading-5 text-muted-foreground">
                      Store inventory reflects the same product state managed by the ERP.
                    </p>
                  </div>
                </div>

                <div class="flex items-start gap-3">
                  <div
                    class="flex h-9 w-9 shrink-0 items-center justify-center rounded-xl bg-primary/10"
                  >
                    <ShoppingBag class="h-4 w-4 text-primary" />
                  </div>

                  <div>
                    <p class="text-sm font-semibold">Structured purchases</p>

                    <p class="mt-1 text-xs leading-5 text-muted-foreground">
                      Customer carts become seller-specific transactional orders.
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- COMMERCE SURFACE -->
      <div
        class="mt-10 overflow-hidden rounded-[2.25rem] border border-border bg-card shadow-2xl sm:mt-12"
      >
        <!-- STORE HEADER -->
        <div
          class="flex flex-col gap-4 border-b border-border bg-background/80 px-5 py-4 sm:flex-row sm:items-center"
        >
          <div class="flex items-center gap-3">
            <div
              class="nexora-gradient nexora-glow flex h-10 w-10 items-center justify-center rounded-xl bg-primary font-bold text-primary-foreground"
            >
              N
            </div>

            <div>
              <p class="text-sm font-semibold">Nexora</p>

              <p class="mt-0.5 text-[11px] text-muted-foreground">Store</p>
            </div>
          </div>

          <div class="mx-auto hidden flex-1 justify-center md:flex">
            <div
              class="flex w-full max-w-xl items-center gap-3 rounded-full bg-muted/40 px-4 py-3 text-sm text-muted-foreground"
            >
              <Search class="h-4 w-4" />
              Search products, categories or businesses
            </div>
          </div>

          <div class="flex items-center gap-2">
            <Button variant="ghost" size="sm" class="rounded-full"> Catalog </Button>

            <Button variant="ghost" size="sm" class="rounded-full"> Businesses </Button>

            <Button size="icon" class="rounded-full">
              <ShoppingBag class="h-4 w-4" />
            </Button>
          </div>
        </div>

        <div class="grid min-h-[590px] lg:grid-cols-[0.72fr_1.28fr]">
          <!-- BUSINESS DIRECTORY -->
          <div class="border-b border-border bg-muted/15 p-5 sm:p-7 lg:border-b-0 lg:border-r">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-xs font-semibold uppercase tracking-[0.14em] text-primary">
                  Businesses
                </p>

                <p class="mt-2 text-xl font-semibold">Connected sellers</p>
              </div>

              <Store class="h-5 w-5 text-primary" />
            </div>

            <div class="mt-6 space-y-4">
              <!-- BUSINESS 1 -->
              <article
                class="overflow-hidden rounded-[1.5rem] border border-border bg-card shadow-sm"
              >
                <div
                  class="relative h-28 bg-gradient-to-br from-blue-500/20 via-primary/10 to-background"
                >
                  <div
                    class="absolute bottom-3 left-3 flex h-12 w-12 items-center justify-center rounded-xl border border-border bg-background shadow-md"
                  >
                    <Building2 class="h-5 w-5 text-primary" />
                  </div>
                </div>

                <div class="p-4">
                  <div class="flex items-center gap-2">
                    <p class="text-sm font-semibold">TechNova</p>

                    <BadgeCheck class="h-4 w-4 text-primary" />
                  </div>

                  <p class="mt-1 text-xs text-muted-foreground">Technology supplier</p>

                  <div class="mt-4 flex items-center justify-between border-t border-border pt-3">
                    <span class="text-xs text-muted-foreground"> 124 products </span>

                    <ArrowRight class="h-4 w-4 text-primary" />
                  </div>
                </div>
              </article>

              <!-- BUSINESS 2 -->
              <article
                class="overflow-hidden rounded-[1.5rem] border border-border bg-card shadow-sm"
              >
                <div
                  class="relative h-28 bg-gradient-to-br from-amber-500/20 via-muted/40 to-background"
                >
                  <div
                    class="absolute bottom-3 left-3 flex h-12 w-12 items-center justify-center rounded-xl border border-border bg-background shadow-md"
                  >
                    <Building2 class="h-5 w-5 text-amber-600 dark:text-amber-400" />
                  </div>
                </div>

                <div class="p-4">
                  <div class="flex items-center gap-2">
                    <p class="text-sm font-semibold">Constructa</p>

                    <BadgeCheck class="h-4 w-4 text-primary" />
                  </div>

                  <p class="mt-1 text-xs text-muted-foreground">Construction & infrastructure</p>

                  <div class="mt-4 flex items-center justify-between border-t border-border pt-3">
                    <span class="text-xs text-muted-foreground"> 86 products </span>

                    <ArrowRight class="h-4 w-4 text-primary" />
                  </div>
                </div>
              </article>
            </div>
          </div>

          <!-- PRODUCT CATALOG -->
          <div class="p-5 sm:p-7">
            <div class="flex flex-col gap-4 sm:flex-row sm:items-end sm:justify-between">
              <div>
                <p class="text-xs font-semibold uppercase tracking-[0.14em] text-primary">
                  Connected inventory
                </p>

                <p class="mt-2 text-2xl font-semibold">Products ready to discover.</p>
              </div>

              <div
                class="flex items-center gap-2 rounded-full border border-border bg-muted/20 px-3 py-1.5 text-xs text-muted-foreground"
              >
                <Grid2X2 class="h-3.5 w-3.5" />
                Live catalog
              </div>
            </div>

            <div class="mt-6 grid gap-4 sm:grid-cols-2">
              <!-- PRODUCT CARD -->
              <article
                class="group overflow-hidden rounded-[1.5rem] border border-border bg-card shadow-sm"
              >
                <div
                  class="relative h-44 overflow-hidden bg-gradient-to-br from-violet-500/20 via-primary/10 to-background"
                >
                  <div class="absolute inset-0 flex items-center justify-center">
                    <div
                      class="flex h-24 w-36 items-center justify-center rounded-2xl border border-border bg-background/80 shadow-xl backdrop-blur"
                    >
                      <ImageIcon class="h-9 w-9 text-primary" />
                    </div>
                  </div>

                  <span
                    class="absolute left-3 top-3 rounded-full border border-border bg-background/85 px-2.5 py-1 text-[10px] font-semibold backdrop-blur"
                  >
                    Technology
                  </span>
                </div>

                <div class="p-4">
                  <p class="flex items-center gap-1.5 text-[11px] text-muted-foreground">
                    <Building2 class="h-3 w-3" />
                    TechNova
                  </p>

                  <p class="mt-2 text-base font-semibold">Business workstation</p>

                  <p class="mt-1 text-xs text-muted-foreground">Professional equipment</p>

                  <div class="mt-5 flex items-end justify-between border-t border-border pt-4">
                    <div>
                      <p class="text-[10px] text-muted-foreground">Price</p>

                      <p class="mt-1 text-lg font-semibold">$1.240.000</p>
                    </div>

                    <div
                      class="flex h-9 w-9 items-center justify-center rounded-full bg-primary text-primary-foreground"
                    >
                      <ShoppingBag class="h-4 w-4" />
                    </div>
                  </div>
                </div>
              </article>

              <!-- PRODUCT CARD -->
              <article
                class="group overflow-hidden rounded-[1.5rem] border border-border bg-card shadow-sm"
              >
                <div
                  class="relative h-44 overflow-hidden bg-gradient-to-br from-blue-500/20 via-muted/30 to-background"
                >
                  <div class="absolute inset-0 flex items-center justify-center">
                    <div
                      class="flex h-24 w-36 items-center justify-center rounded-2xl border border-border bg-background/80 shadow-xl backdrop-blur"
                    >
                      <PackageCheck class="h-9 w-9 text-blue-600 dark:text-blue-400" />
                    </div>
                  </div>

                  <span
                    class="absolute left-3 top-3 rounded-full border border-border bg-background/85 px-2.5 py-1 text-[10px] font-semibold backdrop-blur"
                  >
                    Infrastructure
                  </span>
                </div>

                <div class="p-4">
                  <p class="flex items-center gap-1.5 text-[11px] text-muted-foreground">
                    <Building2 class="h-3 w-3" />
                    Constructa
                  </p>

                  <p class="mt-2 text-base font-semibold">Industrial equipment</p>

                  <p class="mt-1 text-xs text-muted-foreground">Professional machinery</p>

                  <div class="mt-5 flex items-end justify-between border-t border-border pt-4">
                    <div>
                      <p class="text-[10px] text-muted-foreground">Price</p>

                      <p class="mt-1 text-lg font-semibold">$3.480.000</p>
                    </div>

                    <div
                      class="flex h-9 w-9 items-center justify-center rounded-full bg-primary text-primary-foreground"
                    >
                      <ShoppingBag class="h-4 w-4" />
                    </div>
                  </div>
                </div>
              </article>
            </div>

            <!-- CART FLOW -->
            <div
              class="mt-5 flex flex-col gap-4 rounded-[1.5rem] border border-primary/20 bg-primary/5 p-5 sm:flex-row sm:items-center sm:justify-between"
            >
              <div class="flex items-center gap-4">
                <div
                  class="flex h-11 w-11 shrink-0 items-center justify-center rounded-xl bg-primary text-primary-foreground"
                >
                  <ShoppingBag class="h-5 w-5" />
                </div>

                <div>
                  <p class="text-sm font-semibold">One cart, multiple sellers</p>

                  <p class="mt-1 text-xs text-muted-foreground">
                    Nexora organizes independent orders automatically.
                  </p>
                </div>
              </div>

              <div class="flex items-center gap-2 text-xs font-semibold text-primary">
                Inventory connected

                <PackageCheck class="h-4 w-4" />
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- COMMERCE PRINCIPLES -->
      <div class="mt-8 grid gap-4 md:grid-cols-3">
        <article class="rounded-[1.5rem] border border-border bg-card p-5 shadow-sm">
          <Building2 class="h-5 w-5 text-primary" />

          <p class="mt-4 text-sm font-semibold">Seller identity</p>

          <p class="mt-2 text-sm leading-6 text-muted-foreground">
            Every product remains connected to the company behind it.
          </p>
        </article>

        <article class="rounded-[1.5rem] border border-border bg-card p-5 shadow-sm">
          <PackageCheck class="h-5 w-5 text-primary" />

          <p class="mt-4 text-sm font-semibold">Real availability</p>

          <p class="mt-2 text-sm leading-6 text-muted-foreground">
            Store inventory reflects the operational state of the ERP.
          </p>
        </article>

        <article class="rounded-[1.5rem] border border-border bg-card p-5 shadow-sm">
          <ShoppingBag class="h-5 w-5 text-primary" />

          <p class="mt-4 text-sm font-semibold">Structured checkout</p>

          <p class="mt-2 text-sm leading-6 text-muted-foreground">
            Multi-seller carts become independent transactional orders.
          </p>
        </article>
      </div>
    </div>
  </section>
</template>
