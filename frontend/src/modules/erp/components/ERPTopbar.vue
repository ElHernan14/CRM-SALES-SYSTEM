<script setup lang="ts">
import { LogOut } from 'lucide-vue-next';

import { Button } from '@/components/ui/button';
import ThemeToggle from '@/shared/components/ThemeToggle.vue';
import WorkspaceSwitcher from './WorkspaceSwitcher.vue';
import { logout } from '@/modules/auth/composables/useAuth';
</script>

<template>
  <header
    class="sticky top-0 z-40 flex h-16 items-center justify-between border-b border-border bg-background/80 px-6 backdrop-blur"
  >
    <WorkspaceSwitcher />

    <div class="flex items-center gap-2">
      <ThemeToggle />

      <Button variant="outline" size="sm" @click="logout">
        <LogOut class="mr-2 h-4 w-4" />
        Logout
      </Button>
    </div>
  </header>
</template>
