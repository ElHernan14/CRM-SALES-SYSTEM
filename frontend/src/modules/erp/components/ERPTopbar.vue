<script setup lang="ts">
import { LogOut } from "lucide-vue-next"
import { Button } from "@/components/ui/button"
import ThemeToggle from "@/shared/components/ThemeToggle.vue"
import { logout } from "@/modules/auth/composables/useAuth"
</script>

<template>
  <header class="flex h-16 items-center justify-between border-b border-border bg-background/80 px-6 backdrop-blur">
    <div>
      <p class="text-sm font-medium text-foreground">
        Main Workspace
      </p>
      <p class="text-xs text-muted-foreground">
        Multi-tenant ERP dashboard
      </p>
    </div>

    <div class="flex items-center gap-2">
      <ThemeToggle />

      <Button
        variant="outline"
        size="sm"
        @click="logout"
      >
        <LogOut class="mr-2 h-4 w-4" />
        Logout
      </Button>
    </div>
  </header>
</template>