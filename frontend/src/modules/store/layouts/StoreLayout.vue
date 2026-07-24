<script setup lang="ts">
import StoreHeader from '../components/layout/StoreHeader.vue';
import StoreFooter from '../components/layout/StoreFooter.vue';
import StoreMobileNavigation from '../components/layout/StoreMobileNavigation.vue';
import StoreCartDrawer from '../components/cart/StoreCartDrawer.vue';
</script>

<template>
  <div class="min-h-screen bg-background text-foreground">
    <StoreHeader />

    <main class="relative min-h-[calc(100vh-4rem)] overflow-hidden pb-20 md:pb-0">
      <div
        class="pointer-events-none fixed inset-0 -z-10 bg-[radial-gradient(circle_at_top_right,hsl(var(--primary)/0.04),transparent_32rem)]"
      />

      <RouterView />
    </main>

    <StoreFooter />

    <StoreMobileNavigation />

    <StoreCartDrawer />
  </div>
</template>
