<script setup lang="ts">
import {
  ArrowDown,
  ArrowRight,
  Ban,
  Boxes,
  Check,
  CheckCircle2,
  CircleDollarSign,
  Clock3,
  FilePenLine,
  PackageCheck,
  ReceiptText,
  RefreshCcw,
  ShieldCheck,
  WalletCards,
} from 'lucide-vue-next';

const primaryStates = [
  {
    key: 'draft',
    title: 'Draft',
    description: 'Items and quantities can still be changed.',
    icon: FilePenLine,
    stateClass: 'border-border bg-muted/40 text-muted-foreground',
    iconClass: 'bg-muted text-muted-foreground',
    result: 'Editable',
  },

  {
    key: 'pending',
    title: 'Pending',
    description: 'The invoice becomes an active commercial obligation.',
    icon: Clock3,
    stateClass: 'border-amber-500/20 bg-amber-500/10 text-amber-600 dark:text-amber-400',
    iconClass: 'bg-amber-500/10 text-amber-600 dark:text-amber-400',
    result: 'Stock reserved',
  },

  {
    key: 'partial',
    title: 'Partially paid',
    description: 'Payments accumulate while the balance remains open.',
    icon: WalletCards,
    stateClass: 'border-blue-500/20 bg-blue-500/10 text-blue-600 dark:text-blue-400',
    iconClass: 'bg-blue-500/10 text-blue-600 dark:text-blue-400',
    result: 'Balance updated',
  },

  {
    key: 'paid',
    title: 'Paid',
    description: 'The full invoice amount has been completed.',
    icon: CheckCircle2,
    stateClass: 'border-emerald-500/20 bg-emerald-500/10 text-emerald-600 dark:text-emerald-400',
    iconClass: 'bg-emerald-500/10 text-emerald-600 dark:text-emerald-400',
    result: 'Completed',
  },
] as const;

const lifecycleEffects = [
  {
    title: 'Inventory reserved',
    description: 'Submitting an invoice protects the units required by the transaction.',
    icon: Boxes,
    iconClass: 'bg-primary/10 text-primary',
  },

  {
    title: 'Payments accumulated',
    description: 'Every payment contributes to the authoritative paid and remaining balances.',
    icon: CircleDollarSign,
    iconClass: 'bg-blue-500/10 text-blue-600 dark:text-blue-400',
  },

  {
    title: 'Buyer and seller synchronized',
    description: 'The same invoice appears from the correct perspective on both sides.',
    icon: RefreshCcw,
    iconClass: 'bg-violet-500/10 text-violet-600 dark:text-violet-400',
  },

  {
    title: 'Ownership validated',
    description: 'Only authorized tenant participants may view or modify the transaction.',
    icon: ShieldCheck,
    iconClass: 'bg-emerald-500/10 text-emerald-600 dark:text-emerald-400',
  },
] as const;
</script>

<template>
  <section
    id="nexora-platform"
    class="relative scroll-mt-24 overflow-hidden border-b border-border bg-background"
  >
    <div
      class="pointer-events-none absolute -right-52 top-8 h-[520px] w-[520px] rounded-full bg-amber-500/5 blur-3xl"
    />

    <div
      class="pointer-events-none absolute -left-48 bottom-0 h-[480px] w-[480px] rounded-full bg-primary/5 blur-3xl"
    />

    <div class="relative mx-auto w-full max-w-[1500px] px-6 py-24 lg:px-8 lg:py-28">
      <!-- HEADING -->
      <div class="flex flex-col gap-8 lg:flex-row lg:items-end lg:justify-between">
        <div class="max-w-3xl">
          <p class="text-sm font-semibold text-primary">Invoice lifecycle</p>

          <h2 class="mt-4 text-4xl font-semibold tracking-[-0.055em] text-foreground sm:text-5xl">
            Every transaction follows
            <span class="text-muted-foreground"> a real business lifecycle. </span>
          </h2>
        </div>

        <p class="max-w-xl text-base leading-7 text-muted-foreground">
          Nexora prevents invoices from behaving like simple records. Every state changes what users
          can do, how inventory behaves and how balances are calculated.
        </p>
      </div>

      <!-- STATE MACHINE -->
      <div class="relative mt-16 grid gap-5 lg:grid-cols-4">
        <article
          v-for="(state, index) in primaryStates"
          :key="state.key"
          class="relative flex min-h-[270px] flex-col rounded-[1.75rem] border border-border bg-card p-6 shadow-sm"
        >
          <div class="flex items-start justify-between gap-4">
            <div
              class="flex h-12 w-12 items-center justify-center rounded-2xl"
              :class="state.iconClass"
            >
              <component :is="state.icon" class="h-5 w-5" />
            </div>

            <span
              class="rounded-full border px-3 py-1 text-xs font-semibold"
              :class="state.stateClass"
            >
              {{ state.title }}
            </span>
          </div>

          <div class="mt-7 flex-1">
            <p class="text-xs font-semibold uppercase tracking-[0.14em] text-muted-foreground">
              State {{ index + 1 }}
            </p>

            <h3 class="mt-2 text-xl font-semibold">
              {{ state.title }}
            </h3>

            <p class="mt-3 text-sm leading-6 text-muted-foreground">
              {{ state.description }}
            </p>
          </div>

          <div class="mt-6 flex items-center gap-2 border-t border-border pt-4">
            <Check class="h-4 w-4 text-primary" />

            <span class="text-xs font-semibold">
              {{ state.result }}
            </span>
          </div>

          <!-- DESKTOP CONNECTOR -->
          <div
            v-if="index < primaryStates.length - 1"
            class="absolute -right-[19px] top-1/2 z-20 hidden h-9 w-9 -translate-y-1/2 items-center justify-center rounded-full border border-border bg-background shadow-md lg:flex"
          >
            <ArrowRight class="h-4 w-4 text-primary" />
          </div>

          <!-- MOBILE CONNECTOR -->
          <div
            v-if="index < primaryStates.length - 1"
            class="absolute -bottom-[19px] left-1/2 z-20 flex h-9 w-9 -translate-x-1/2 items-center justify-center rounded-full border border-border bg-background shadow-md lg:hidden"
          >
            <ArrowDown class="h-4 w-4 text-primary" />
          </div>
        </article>
      </div>

      <!-- CANCELLATION PATH -->
      <div
        class="mt-8 overflow-hidden rounded-[1.75rem] border border-destructive/20 bg-destructive/5"
      >
        <div class="flex flex-col gap-5 p-6 sm:flex-row sm:items-center sm:justify-between">
          <div class="flex items-start gap-4">
            <div
              class="flex h-12 w-12 shrink-0 items-center justify-center rounded-2xl bg-destructive/10"
            >
              <Ban class="h-5 w-5 text-destructive" />
            </div>

            <div>
              <p class="text-xs font-semibold uppercase tracking-[0.14em] text-destructive">
                Controlled alternative path
              </p>

              <h3 class="mt-2 text-xl font-semibold">Cancellation is a business event.</h3>

              <p class="mt-2 max-w-3xl text-sm leading-6 text-muted-foreground">
                Eligible invoices can move to Cancelled, releasing reserved inventory and preventing
                additional financial operations.
              </p>
            </div>
          </div>

          <div
            class="flex shrink-0 items-center gap-3 rounded-2xl border border-destructive/20 bg-background px-4 py-3"
          >
            <PackageCheck class="h-5 w-5 text-destructive" />

            <div>
              <p class="text-xs text-muted-foreground">Inventory result</p>

              <p class="mt-1 text-sm font-semibold">Reserved stock released</p>
            </div>
          </div>
        </div>
      </div>

      <!-- PAYMENT VISUAL -->
      <div class="mt-8 grid gap-6 lg:grid-cols-[0.86fr_1.14fr]">
        <!-- FINANCIAL CARD -->
        <article
          class="relative overflow-hidden rounded-[2rem] border border-border bg-card p-6 shadow-sm sm:p-8"
        >
          <div
            class="pointer-events-none absolute -right-20 -top-20 h-52 w-52 rounded-full bg-primary/10 blur-3xl"
          />

          <div class="relative">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-xs font-semibold uppercase tracking-[0.14em] text-primary">
                  Financial snapshot
                </p>

                <h3 class="mt-2 text-2xl font-semibold">Authoritative balances</h3>
              </div>

              <ReceiptText class="h-6 w-6 text-primary" />
            </div>

            <div class="mt-8 space-y-4">
              <div class="flex items-center justify-between border-b border-border pb-4">
                <span class="text-sm text-muted-foreground"> Invoice total </span>

                <span class="text-lg font-semibold"> $1.200.000,00 </span>
              </div>

              <div class="flex items-center justify-between border-b border-border pb-4">
                <span class="text-sm text-muted-foreground"> Paid amount </span>

                <span class="text-lg font-semibold text-emerald-600 dark:text-emerald-400">
                  $450.000,00
                </span>
              </div>

              <div class="flex items-end justify-between">
                <div>
                  <p class="text-sm font-semibold">Remaining balance</p>

                  <p class="mt-1 text-xs text-muted-foreground">
                    Calculated from recorded payments
                  </p>
                </div>

                <span
                  class="text-2xl font-semibold tracking-[-0.035em] text-amber-600 dark:text-amber-400"
                >
                  $750.000,00
                </span>
              </div>
            </div>

            <div class="mt-7 h-2 overflow-hidden rounded-full bg-muted">
              <div class="h-full w-[37.5%] rounded-full bg-primary" />
            </div>

            <div class="mt-3 flex items-center justify-between text-xs text-muted-foreground">
              <span>37.5% paid</span>

              <span>Payment history preserved</span>
            </div>
          </div>
        </article>

        <!-- EFFECTS -->
        <div class="grid gap-4 sm:grid-cols-2">
          <article
            v-for="effect in lifecycleEffects"
            :key="effect.title"
            class="rounded-[1.5rem] border border-border bg-card p-5 shadow-sm"
          >
            <div
              class="flex h-11 w-11 items-center justify-center rounded-xl"
              :class="effect.iconClass"
            >
              <component :is="effect.icon" class="h-5 w-5" />
            </div>

            <h3 class="mt-5 text-base font-semibold">
              {{ effect.title }}
            </h3>

            <p class="mt-2 text-sm leading-6 text-muted-foreground">
              {{ effect.description }}
            </p>
          </article>
        </div>
      </div>

      <!-- FINAL STATEMENT -->
      <div
        class="mt-8 flex flex-col gap-5 rounded-[1.75rem] border border-border bg-muted/15 p-6 sm:flex-row sm:items-center sm:justify-between"
      >
        <div class="flex items-start gap-4">
          <div
            class="flex h-11 w-11 shrink-0 items-center justify-center rounded-xl bg-primary text-primary-foreground"
          >
            <ShieldCheck class="h-5 w-5" />
          </div>

          <div>
            <p class="text-sm font-semibold">Business rules remain authoritative.</p>

            <p class="mt-1 text-sm leading-6 text-muted-foreground">
              The interface reflects what the workflow allows; it does not invent the commercial
              state.
            </p>
          </div>
        </div>

        <span
          class="self-start rounded-full border border-primary/20 bg-primary/10 px-3 py-1.5 text-xs font-semibold text-primary sm:self-auto"
        >
          State-driven experience
        </span>
      </div>
    </div>
  </section>
</template>
