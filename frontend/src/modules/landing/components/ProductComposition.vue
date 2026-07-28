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
  <div class="relative mx-auto min-h-[500px] w-full max-w-[670px] sm:min-h-[540px]">
    <!-- GLOW -->
    <div class="pointer-events-none absolute inset-16 rounded-full bg-primary/15 blur-3xl" />

    <div
      class="pointer-events-none absolute right-0 top-8 h-72 w-72 rounded-full bg-violet-500/10 blur-3xl"
    />

    <!-- MAIN DASHBOARD SURFACE -->
    <div
      class="absolute inset-x-0 top-6 overflow-hidden rounded-[1.5rem] border border-border/80 bg-card/95 shadow-2xl backdrop-blur-xl sm:inset-x-12 sm:top-10 sm:rounded-[2rem]"
    >
      <div class="flex items-center justify-between border-b border-border bg-muted/20 px-5 py-4">
        <div class="flex min-w-0 items-center gap-2 sm:gap-3">
          <div
            class="nexora-gradient nexora-glow flex h-9 w-9 items-center justify-center rounded-xl bg-primary text-primary-foreground"
          >
            N
          </div>

          <div class="min-w-0">
            <p class="truncate text-xs font-semibold sm:text-sm">Nexora Workspace</p>

            <p class="mt-0.5 hidden text-[11px] text-muted-foreground sm:block">
              Connected operations
            </p>
          </div>
        </div>

        <span
          class="inline-flex shrink-0 items-center gap-1.5 rounded-full border border-emerald-500/20 bg-emerald-500/10 px-2 py-1 text-[10px] font-semibold text-emerald-600 dark:text-emerald-400 sm:gap-2 sm:px-3 sm:text-[11px]"
        >
          <span class="h-1.5 w-1.5 rounded-full bg-emerald-500" />
          Live
        </span>
      </div>

      <div class="p-3 sm:p-5">
        <div class="grid grid-cols-3 gap-1.5 sm:gap-3">
          <div
            class="min-w-0 rounded-xl border border-border bg-background p-2 sm:rounded-2xl sm:p-4"
          >
            <div
              class="flex h-7 w-7 items-center justify-center rounded-lg bg-emerald-500/10 sm:h-9 sm:w-9 sm:rounded-xl"
            >
              <TrendingUp class="h-4 w-4 text-emerald-600 dark:text-emerald-400" />
            </div>

            <p class="mt-2 truncate text-[9px] text-muted-foreground sm:mt-4 sm:text-[11px]">
              Revenue
            </p>

            <p class="mt-1 truncate text-sm font-semibold tracking-tight sm:text-lg">$8.42M</p>
          </div>

          <div
            class="min-w-0 rounded-xl border border-border bg-background p-2 sm:rounded-2xl sm:p-4"
          >
            <div
              class="flex h-7 w-7 items-center justify-center rounded-lg bg-primary/10 sm:h-9 sm:w-9 sm:rounded-xl"
            >
              <WalletCards class="h-4 w-4 text-primary" />
            </div>

            <p class="mt-2 truncate text-[9px] text-muted-foreground sm:mt-4 sm:text-[11px]">
              <span class="sm:hidden">Open</span>
              <span class="hidden sm:inline">Outstanding</span>
            </p>

            <p class="mt-1 truncate text-sm font-semibold tracking-tight sm:text-lg">$640K</p>
          </div>

          <div
            class="min-w-0 rounded-xl border border-border bg-background p-2 sm:rounded-2xl sm:p-4"
          >
            <div
              class="flex h-7 w-7 items-center justify-center rounded-lg bg-blue-500/10 sm:h-9 sm:w-9 sm:rounded-xl"
            >
              <Boxes class="h-4 w-4 text-blue-600 dark:text-blue-400" />
            </div>

            <p class="mt-2 truncate text-[9px] text-muted-foreground sm:mt-4 sm:text-[11px]">
              Inventory
            </p>

            <p class="mt-1 truncate text-sm font-semibold tracking-tight sm:text-lg">1,248</p>
          </div>
        </div>

        <div
          class="mt-3 rounded-xl border border-border bg-background p-3 sm:mt-4 sm:rounded-2xl sm:p-4"
        >
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm font-semibold">Commercial activity</p>

              <p class="mt-1 hidden text-xs text-muted-foreground sm:block">
                ERP and Store working together
              </p>
            </div>

            <ArrowUpRight class="h-4 w-4 text-primary" />
          </div>

          <div class="mt-4 flex h-16 items-end gap-1.5 sm:mt-5 sm:h-24 sm:gap-2">
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
      class="absolute bottom-6 left-0 z-20 w-[205px] overflow-hidden rounded-[1.25rem] border border-border bg-card shadow-2xl sm:bottom-5 sm:w-[290px] sm:rounded-[1.5rem]"
    >
      <div
        class="relative h-24 bg-gradient-to-br from-violet-500/20 via-primary/10 to-background sm:h-32"
      >
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

      <div class="p-3 sm:p-4">
        <p class="flex items-center gap-1.5 text-[11px] text-muted-foreground">
          <Building2 class="h-3 w-3" />
          TechNova
        </p>

        <p class="mt-2 truncate text-xs font-semibold sm:text-sm">Business workstation</p>

        <div class="mt-4 flex items-end justify-between">
          <div>
            <p class="text-[10px] text-muted-foreground">Current price</p>

            <p class="mt-1 text-base font-semibold sm:text-lg">$1.24M</p>
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
      class="absolute bottom-0 right-0 z-30 w-[190px] rounded-[1.25rem] border border-border bg-card p-3 shadow-2xl sm:w-[270px] sm:rounded-[1.5rem] sm:p-4"
    >
      <div class="flex min-w-0 items-center justify-between gap-2">
        <div class="flex min-w-0 items-center gap-2 sm:gap-3">
          <div
            class="flex h-8 w-8 shrink-0 items-center justify-center rounded-lg bg-emerald-500/10 sm:h-10 sm:w-10 sm:rounded-xl"
          >
            <CheckCircle2 class="h-5 w-5 text-emerald-600 dark:text-emerald-400" />
          </div>

          <div>
            <p class="truncate text-xs font-semibold sm:text-sm">Order received</p>

            <p class="mt-0.5 text-[11px] text-muted-foreground">#NX-1048</p>
          </div>
        </div>

        <BadgeCheck class="h-4 w-4 text-primary" />
      </div>

      <div class="mt-3 space-y-2 text-[10px] sm:mt-4 sm:space-y-2.5 sm:text-xs">
        <div class="flex items-center justify-between text-xs">
          <span class="flex min-w-0 items-center gap-1.5 text-muted-foreground sm:gap-2">
            <PackageCheck class="h-3.5 w-3.5" />
            Inventory
          </span>

          <span class="font-semibold text-emerald-600 dark:text-emerald-400"> Reserved </span>
        </div>

        <div class="flex items-center justify-between text-xs">
          <span class="flex min-w-0 items-center gap-1.5 text-muted-foreground sm:gap-2">
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
