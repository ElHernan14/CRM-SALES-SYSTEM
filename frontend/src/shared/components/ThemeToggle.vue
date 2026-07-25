<script setup lang="ts">
import { Moon, Sun } from 'lucide-vue-next';

import { Button } from '@/components/ui/button';

import { useTheme } from '@/shared/composables/useTheme';

const { isDark, toggleTheme } = useTheme();
</script>

<template>
  <Button
    type="button"
    variant="ghost"
    size="icon"
    class="rounded-full"
    :aria-label="isDark ? 'Use light theme' : 'Use dark theme'"
    @click="toggleTheme"
  >
    <Sun v-if="isDark" class="h-4 w-4" />

    <Moon v-else class="h-4 w-4" />
  </Button>
</template>
