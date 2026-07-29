<template>
  <div class="min-h-screen bg-background text-foreground">
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
        <component :is="Component" :key="route.fullPath" />
      </Transition>
    </RouterView>
  </div>
</template>
