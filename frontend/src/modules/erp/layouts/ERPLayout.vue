<script setup lang="ts">
import ERPSidebar from "../components/ERPSidebar.vue"
import ERPTopbar from "../components/ERPTopbar.vue"
</script>

<template>
  <div class="min-h-screen bg-background text-foreground">
    <div class="flex min-h-screen">
      <ERPSidebar />

      <div class="flex min-w-0 flex-1 flex-col">
        <ERPTopbar />

        <main class="flex-1 p-6">
          <RouterView />
        </main>
      </div>
    </div>
  </div>
</template>