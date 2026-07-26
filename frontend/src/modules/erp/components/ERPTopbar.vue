<script setup lang="ts">
import { Menu } from 'lucide-vue-next';

import { Button } from '@/components/ui/button';

import ThemeToggle from '@/shared/components/ThemeToggle.vue';

import WorkspaceSwitcher from './WorkspaceSwitcher.vue';

const emit = defineEmits<{
  'open-navigation': [];
}>();
</script>

<template>
  <header
    class="sticky top-0 z-40 flex h-16 shrink-0 items-center justify-between border-b border-border bg-background/85 px-3 backdrop-blur-xl sm:px-6"
  >
    <div class="flex min-w-0 items-center gap-2">
      <Button
        type="button"
        variant="ghost"
        size="icon"
        class="shrink-0 rounded-full lg:hidden"
        aria-label="Open ERP navigation"
        @click="emit('open-navigation')"
      >
        <Menu class="h-5 w-5" />
      </Button>

      <div class="min-w-0">
        <WorkspaceSwitcher />
      </div>
    </div>

    <div class="ml-2 flex shrink-0 items-center gap-2">
      <ThemeToggle />
    </div>
  </header>
</template>
