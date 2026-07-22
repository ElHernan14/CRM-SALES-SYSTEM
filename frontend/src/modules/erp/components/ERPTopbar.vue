<script setup lang="ts">
import ThemeToggle from '@/shared/components/ThemeToggle.vue';
import WorkspaceSwitcher from './WorkspaceSwitcher.vue';
</script>

<template>
  <header
    class="sticky top-0 z-40 flex h-16 items-center justify-between border-b border-border bg-background/85 px-4 backdrop-blur-xl sm:px-6"
  >
    <div class="min-w-0">
      <WorkspaceSwitcher />
    </div>

    <div class="ml-4 flex shrink-0 items-center gap-2">
      <ThemeToggle />
    </div>
  </header>
</template>
