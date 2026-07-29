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

        <main class="flex-1 overflow-y-auto p-4 sm:p-6">
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
      </div>
    </div>
  </div>
</template>
