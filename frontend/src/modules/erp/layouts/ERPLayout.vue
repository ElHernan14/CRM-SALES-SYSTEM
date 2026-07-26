<script setup lang="ts">
import { ref } from 'vue';

import ERPSidebar from '../components/ERPSidebar.vue';
import ERPTopbar from '../components/ERPTopbar.vue';

const mobileNavigationOpen = ref(false);

function openMobileNavigation() {
  mobileNavigationOpen.value = true;
}

function updateMobileNavigation(value: boolean) {
  mobileNavigationOpen.value = value;
}
</script>

<template>
  <div class="min-h-screen bg-background text-foreground">
    <div class="flex min-h-screen">
      <ERPSidebar
        :mobile-open="mobileNavigationOpen"
        @update:mobile-open="updateMobileNavigation"
      />

      <div class="flex min-w-0 flex-1 flex-col">
        <ERPTopbar @open-navigation="openMobileNavigation" />

        <main class="min-w-0 flex-1 overflow-x-hidden px-4 py-5 sm:px-6 sm:py-6">
          <RouterView />
        </main>
      </div>
    </div>
  </div>
</template>
