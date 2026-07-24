<script setup lang="ts">
import { Building2, Settings2, UserRound } from 'lucide-vue-next';

const settingsNavigation = [
  {
    name: 'erp-settings-company',
    label: 'Company workspace',
    description: 'Business profile and branding',
    icon: Building2,
  },
  {
    name: 'erp-settings-profile',
    label: 'My profile',
    description: 'Personal identity and contact',
    icon: UserRound,
  },
];
</script>

<template>
  <div class="space-y-6">
    <!-- SETTINGS HEADER -->
    <section
      class="relative overflow-hidden rounded-[2rem] border border-border bg-card p-6 shadow-sm sm:p-8"
    >
      <div
        class="pointer-events-none absolute -right-24 -top-32 h-72 w-72 rounded-full bg-primary/10 blur-3xl"
      />

      <div class="relative">
        <div class="flex h-11 w-11 items-center justify-center rounded-xl bg-primary/10">
          <Settings2 class="h-5 w-5 text-primary" />
        </div>

        <p class="mt-5 text-xs font-semibold uppercase tracking-[0.16em] text-primary">
          Nexora settings
        </p>

        <h1 class="mt-2 text-3xl font-semibold tracking-[-0.04em] text-foreground">
          Identity and workspace
        </h1>

        <p class="mt-3 max-w-2xl text-sm leading-6 text-muted-foreground">
          Manage your personal account and the business identity represented throughout Nexora.
        </p>
      </div>
    </section>

    <!-- NAVIGATION -->
    <nav
      class="grid gap-3 rounded-[1.5rem] border border-border bg-card p-2 shadow-sm sm:grid-cols-2"
      aria-label="Settings sections"
    >
      <RouterLink
        v-for="item in settingsNavigation"
        :key="item.name"
        :to="{ name: item.name }"
        class="group flex items-center gap-3 rounded-xl border border-transparent p-3 transition hover:bg-muted/50"
        active-class="border-primary/20 bg-primary/5 shadow-sm"
      >
        <div
          class="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl bg-muted text-muted-foreground transition group-hover:text-foreground"
        >
          <component :is="item.icon" class="h-5 w-5" />
        </div>

        <div class="min-w-0">
          <p class="text-sm font-semibold text-foreground">
            {{ item.label }}
          </p>

          <p class="mt-1 truncate text-xs text-muted-foreground">
            {{ item.description }}
          </p>
        </div>
      </RouterLink>
    </nav>

    <!-- ACTIVE SETTINGS EXPERIENCE -->
    <RouterView />
  </div>
</template>
