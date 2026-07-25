<script setup lang="ts">
import {
  ArrowUpRight,
  BadgeCheck,
  Boxes,
  Building2,
  CheckCircle2,
  CircleDollarSign,
  PackageCheck,
  ReceiptText,
  ShoppingBag,
  TrendingUp,
  WalletCards,
} from 'lucide-vue-next';
</script>

<template>
  <div class="relative mx-auto min-h-[540px] w-full max-w-[670px]">
    <!-- GLOW -->
    <div class="pointer-events-none absolute inset-16 rounded-full bg-primary/15 blur-3xl" />

    <div
      class="pointer-events-none absolute right-0 top-8 h-72 w-72 rounded-full bg-violet-500/10 blur-3xl"
    />

    <!-- MAIN DASHBOARD SURFACE -->
    <div
      class="absolute inset-x-6 top-10 overflow-hidden rounded-[2rem] border border-border/80 bg-card/95 shadow-2xl backdrop-blur-xl sm:inset-x-12"
    >
      <div class="flex items-center justify-between border-b border-border bg-muted/20 px-5 py-4">
        <div class="flex items-center gap-3">
          <div
            class="flex h-9 w-9 items-center justify-center rounded-xl bg-primary text-primary-foreground"
          >
            N
          </div>

          <div>
            <p class="text-sm font-semibold">Nexora Workspace</p>

            <p class="mt-0.5 text-[11px] text-muted-foreground">Connected operations</p>
          </div>
        </div>

        <span
          class="inline-flex items-center gap-2 rounded-full border border-emerald-500/20 bg-emerald-500/10 px-3 py-1 text-[11px] font-semibold text-emerald-600 dark:text-emerald-400"
        >
          <span class="h-1.5 w-1.5 rounded-full bg-emerald-500" />
          Live
        </span>
      </div>

      <div class="p-5">
        <div class="grid grid-cols-3 gap-3">
          <div class="rounded-2xl border border-border bg-background p-4">
            <div class="flex h-9 w-9 items-center justify-center rounded-xl bg-emerald-500/10">
              <TrendingUp class="h-4 w-4 text-emerald-600 dark:text-emerald-400" />
            </div>

            <p class="mt-4 text-[11px] text-muted-foreground">Revenue</p>

            <p class="mt-1 text-lg font-semibold tracking-tight">$8.42M</p>
          </div>

          <div class="rounded-2xl border border-border bg-background p-4">
            <div class="flex h-9 w-9 items-center justify-center rounded-xl bg-primary/10">
              <WalletCards class="h-4 w-4 text-primary" />
            </div>

            <p class="mt-4 text-[11px] text-muted-foreground">Outstanding</p>

            <p class="mt-1 text-lg font-semibold tracking-tight">$640K</p>
          </div>

          <div class="rounded-2xl border border-border bg-background p-4">
            <div class="flex h-9 w-9 items-center justify-center rounded-xl bg-blue-500/10">
              <Boxes class="h-4 w-4 text-blue-600 dark:text-blue-400" />
            </div>

            <p class="mt-4 text-[11px] text-muted-foreground">Inventory</p>

            <p class="mt-1 text-lg font-semibold tracking-tight">1,248</p>
          </div>
        </div>

        <div class="mt-4 rounded-2xl border border-border bg-background p-4">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm font-semibold">Commercial activity</p>

              <p class="mt-1 text-xs text-muted-foreground">ERP and Store working together</p>
            </div>

            <ArrowUpRight class="h-4 w-4 text-primary" />
          </div>

          <div class="mt-5 flex h-24 items-end gap-2">
            <div class="h-[30%] flex-1 rounded-t-md bg-primary/25" />
            <div class="h-[48%] flex-1 rounded-t-md bg-primary/35" />
            <div class="h-[42%] flex-1 rounded-t-md bg-primary/30" />
            <div class="h-[70%] flex-1 rounded-t-md bg-primary/50" />
            <div class="h-[62%] flex-1 rounded-t-md bg-primary/45" />
            <div class="h-[88%] flex-1 rounded-t-md bg-primary" />
            <div class="h-[76%] flex-1 rounded-t-md bg-primary/70" />
          </div>
        </div>
      </div>
    </div>

    <!-- PRODUCT CARD -->
    <div
      class="absolute bottom-5 left-0 z-20 w-[260px] overflow-hidden rounded-[1.5rem] border border-border bg-card shadow-2xl sm:w-[290px]"
    >
      <div class="relative h-32 bg-gradient-to-br from-violet-500/20 via-primary/10 to-background">
        <div class="absolute inset-0 flex items-center justify-center">
          <div
            class="flex h-20 w-28 items-center justify-center rounded-2xl border border-border/70 bg-background/80 shadow-xl backdrop-blur"
          >
            <Boxes class="h-9 w-9 text-primary" />
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

        <p class="mt-2 text-sm font-semibold">Business workstation</p>

        <div class="mt-4 flex items-end justify-between">
          <div>
            <p class="text-[10px] text-muted-foreground">Current price</p>

            <p class="mt-1 text-lg font-semibold">$1.24M</p>
          </div>

          <div
            class="flex h-9 w-9 items-center justify-center rounded-full bg-primary text-primary-foreground"
          >
            <ShoppingBag class="h-4 w-4" />
          </div>
        </div>
      </div>
    </div>

    <!-- TRANSACTION CARD -->
    <div
      class="absolute bottom-0 right-0 z-30 w-[245px] rounded-[1.5rem] border border-border bg-card p-4 shadow-2xl sm:w-[270px]"
    >
      <div class="flex items-center justify-between gap-3">
        <div class="flex items-center gap-3">
          <div class="flex h-10 w-10 items-center justify-center rounded-xl bg-emerald-500/10">
            <CheckCircle2 class="h-5 w-5 text-emerald-600 dark:text-emerald-400" />
          </div>

          <div>
            <p class="text-sm font-semibold">Order received</p>

            <p class="mt-0.5 text-[11px] text-muted-foreground">#NX-1048</p>
          </div>
        </div>

        <BadgeCheck class="h-4 w-4 text-primary" />
      </div>

      <div class="mt-4 space-y-2.5">
        <div class="flex items-center justify-between text-xs">
          <span class="flex items-center gap-2 text-muted-foreground">
            <PackageCheck class="h-3.5 w-3.5" />
            Inventory
          </span>

          <span class="font-semibold text-emerald-600 dark:text-emerald-400"> Reserved </span>
        </div>

        <div class="flex items-center justify-between text-xs">
          <span class="flex items-center gap-2 text-muted-foreground">
            <ReceiptText class="h-3.5 w-3.5" />
            Invoice
          </span>

          <span class="font-semibold text-amber-600 dark:text-amber-400"> Pending </span>
        </div>

        <div class="flex items-center justify-between text-xs">
          <span class="flex items-center gap-2 text-muted-foreground">
            <CircleDollarSign class="h-3.5 w-3.5" />
            Payment
          </span>

          <span class="font-semibold"> Tracked </span>
        </div>
      </div>
    </div>

    <!-- CONNECTION LINE -->
    <div
      class="pointer-events-none absolute bottom-28 left-[245px] hidden h-px w-24 bg-gradient-to-r from-primary to-transparent sm:block"
    />
  </div>
</template>
