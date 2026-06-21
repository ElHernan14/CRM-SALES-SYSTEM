<script setup lang="ts">
import LoginForm from "../components/LoginForm.vue"
</script>

<template>
  <div
    class="min-h-screen bg-stone-50"
  >
    <div
      class="grid min-h-screen lg:grid-cols-2"
    >
      <!-- LEFT PANEL -->

      <div
        class="lg:flex flex-col justify-between p-16"
      >
        <div>
          <span
            class="text-sm font-medium text-stone-500"
          >
            CRM Commerce SaaS
          </span>

          <h1
            class="mt-8 max-w-lg text-5xl font-semibold tracking-tight text-stone-900"
          >
            Manage your business from a single platform.
          </h1>

          <p
            class="mt-6 max-w-md text-lg text-stone-600"
          >
            CRM, ERP, invoicing, inventory and commerce
            unified in one modern workspace.
          </p>
        </div>

        <div
          class="flex flex-wrap gap-3"
        >
          <span
            class="rounded-full border border-stone-200 bg-white px-4 py-2 text-sm text-stone-600"
          >
            Multi Tenant
          </span>

          <span
            class="rounded-full border border-stone-200 bg-white px-4 py-2 text-sm text-stone-600"
          >
            ERP
          </span>

          <span
            class="rounded-full border border-stone-200 bg-white px-4 py-2 text-sm text-stone-600"
          >
            CRM
          </span>

          <span
            class="rounded-full border border-stone-200 bg-white px-4 py-2 text-sm text-stone-600"
          >
            B2B / B2C
          </span>
        </div>
      </div>

      <!-- RIGHT PANEL -->

      <div
        class="flex items-center justify-center p-6 lg:p-16"
      >
        <LoginForm />
      </div>
    </div>
  </div>
</template>