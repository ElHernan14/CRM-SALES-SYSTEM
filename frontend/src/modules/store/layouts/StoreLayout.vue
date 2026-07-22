<script setup lang="ts">
import StoreHeader from '../components/layout/StoreHeader.vue';
import StoreFooter from '../components/layout/StoreFooter.vue';
import StoreMobileNavigation from '../components/layout/StoreMobileNavigation.vue';
</script>

<template>
  <div class="min-h-screen bg-background text-foreground">
    <StoreHeader />

    <main class="min-h-[calc(100vh-4rem)] pb-20 md:pb-0">
      <RouterView />
    </main>

    <StoreFooter />

    <StoreMobileNavigation />
  </div>
</template>
