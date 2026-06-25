<script setup lang="ts">
import { Loader2 } from "lucide-vue-next"
import { storeToRefs } from "pinia"

import { useUiStore } from "@/shared/stores/ui.store"

const ui = useUiStore()

const {
  loading,
  loadingMessage
} = storeToRefs(ui)
</script>

<template>
  <Transition name="fade">
    <div
      v-if="loading"
      class="fixed inset-0 z-9999 flex items-center justify-center bg-background/80 backdrop-blur-sm"
    >
      <div
        class="flex flex-col items-center gap-4"
      >
        <Loader2
          class="h-8 w-8 animate-spin text-primary"
        />

        <p
          v-if="loadingMessage"
          class="text-sm text-muted-foreground"
        >
          {{ loadingMessage }}
        </p>
      </div>
    </div>
  </Transition>
</template>