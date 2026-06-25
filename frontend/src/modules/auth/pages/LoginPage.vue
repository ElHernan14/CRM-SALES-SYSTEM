<script setup lang="ts">
import LoginForm from "../components/LoginForm.vue"
import ThemeToggle from "@/shared/components/ThemeToggle.vue"
</script>

<template>
  <div class="min-h-screen bg-muted/40 text-foreground">
    <div class="fixed right-6 top-6 z-50">
      <ThemeToggle />
    </div>

    <div class="grid min-h-screen lg:grid-cols-2">
      <!-- LEFT PANEL -->
      <div class="hidden flex-col justify-between p-16 lg:flex">
        <div>
          <span class="text-sm font-medium text-muted-foreground">
            CRM Commerce SaaS
          </span>

          <h1 class="mt-8 max-w-lg text-5xl font-semibold tracking-tight text-foreground">
            Manage your business from a single platform.
          </h1>

          <p class="mt-6 max-w-md text-lg text-muted-foreground">
            CRM, ERP, invoicing, inventory and commerce
            unified in one modern workspace.
          </p>
        </div>

        <div class="flex flex-wrap gap-3">
          <span class="rounded-full border border-border bg-card px-4 py-2 text-sm text-muted-foreground shadow-sm">
            Multi Tenant
          </span>

          <span class="rounded-full border border-border bg-card px-4 py-2 text-sm text-muted-foreground shadow-sm">
            ERP
          </span>

          <span class="rounded-full border border-border bg-card px-4 py-2 text-sm text-muted-foreground shadow-sm">
            CRM
          </span>

          <span class="rounded-full border border-border bg-card px-4 py-2 text-sm text-muted-foreground shadow-sm">
            B2B / B2C
          </span>
        </div>
      </div>

      <!-- RIGHT PANEL -->
      <div class="flex items-center justify-center p-6 lg:p-16">
        <LoginForm />
      </div>
    </div>
  </div>
</template>