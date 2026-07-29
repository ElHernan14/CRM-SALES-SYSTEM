<script setup lang="ts">
import { Check, Monitor, Moon, Sun } from 'lucide-vue-next';

import { Button } from '@/components/ui/button';

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';

import { useTheme, type ThemePreference } from '@/shared/composables/useTheme';

const { theme, isDark, setTheme } = useTheme();

const options: {
  value: ThemePreference;
  label: string;
  description: string;
  icon: typeof Sun;
}[] = [
  {
    value: 'light',
    label: 'Light',
    description: 'Use the light Nexora theme',
    icon: Sun,
  },
  {
    value: 'dark',
    label: 'Dark',
    description: 'Use the dark Nexora theme',
    icon: Moon,
  },
  {
    value: 'system',
    label: 'System',
    description: 'Match your device settings',
    icon: Monitor,
  },
];
</script>

<template>
  <DropdownMenu>
    <DropdownMenuTrigger as-child>
      <Button
        type="button"
        variant="outline"
        size="icon"
        class="relative rounded-xl"
        aria-label="Change color theme"
      >
        <Moon v-if="isDark" class="h-4 w-4" />

        <Sun v-else class="h-4 w-4" />

        <span class="sr-only"> Change color theme </span>
      </Button>
    </DropdownMenuTrigger>

    <DropdownMenuContent align="end" :side-offset="8" class="w-64">
      <DropdownMenuLabel>
        <div>
          <p class="text-sm font-semibold">Appearance</p>

          <p class="mt-1 text-xs font-normal text-muted-foreground">
            Choose how Nexora looks on this device.
          </p>
        </div>
      </DropdownMenuLabel>

      <DropdownMenuSeparator />

      <DropdownMenuItem
        v-for="option in options"
        :key="option.value"
        class="gap-3 py-3"
        @select="setTheme(option.value)"
      >
        <div
          class="flex h-9 w-9 shrink-0 items-center justify-center rounded-xl border border-border bg-muted/40"
        >
          <component :is="option.icon" class="h-4 w-4" />
        </div>

        <div class="min-w-0 flex-1">
          <p class="text-sm font-medium">
            {{ option.label }}
          </p>

          <p class="mt-1 text-xs text-muted-foreground">
            {{ option.description }}
          </p>
        </div>

        <Check v-if="theme === option.value" class="h-4 w-4 shrink-0 text-primary" />
      </DropdownMenuItem>
    </DropdownMenuContent>
  </DropdownMenu>
</template>
