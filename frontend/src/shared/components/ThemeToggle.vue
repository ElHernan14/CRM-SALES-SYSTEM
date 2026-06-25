<script setup lang="ts">
import { Moon, Sun } from "lucide-vue-next"
import { Button } from "@/components/ui/button"
import { useTheme } from "@/shared/composables/useTheme"

const { isDark, toggleTheme } = useTheme()
</script>

<template>
  <Button
    type="button"
    variant="outline"
    size="icon"
    @click="toggleTheme"
  >
    <Sun v-if="isDark" class="h-4 w-4" />
    <Moon v-else class="h-4 w-4" />
  </Button>
</template>