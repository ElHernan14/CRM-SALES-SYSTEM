<script setup lang="ts">
import StoreHeader from '../components/layout/StoreHeader.vue';
import StoreFooter from '../components/layout/StoreFooter.vue';
import StoreMobileNavigation from '../components/layout/StoreMobileNavigation.vue';
import StoreCartDrawer from '../components/cart/StoreCartDrawer.vue';
</script>

<template>
  <div class="min-h-screen bg-background text-foreground">
    <StoreHeader />

    <main class="min-h-[calc(100vh-4rem)] pb-20 md:pb-0">
      <RouterView v-slot="{ Component, route }">
        <Transition
          mode="out-in"
          enter-active-class="transition duration-300 ease-out"
          enter-from-class="opacity-0 translate-y-1"
          enter-to-class="opacity-100 translate-y-0"
          leave-active-class="transition duration-150 ease-in"
          leave-from-class="opacity-100"
          leave-to-class="opacity-0"
        >
          <component :is="Component" :key="String(route.name)" />
        </Transition>
      </RouterView>
    </main>

    <StoreFooter />

    <StoreMobileNavigation />

    <StoreCartDrawer />
  </div>
</template>
