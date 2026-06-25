<script setup lang="ts">
import { storeToRefs } from "pinia"
import { Button } from "@/components/ui/button"

import {
  AlertDialog,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog"

import { useUiStore } from "@/shared/stores/ui.store"

const ui = useUiStore()

const {
  confirmOpen,
  confirmTitle,
  confirmDescription,
  confirmText,
  cancelText,
  confirmVariant,
} = storeToRefs(ui)

async function handleConfirm() {
  await ui.confirm()
}
</script>

<template>
    <AlertDialog
        :open="confirmOpen"
        @update:open="(value) => {
            if (!value) ui.closeConfirm()
        }"
    >
    <AlertDialogContent>
      <AlertDialogHeader>
        <AlertDialogTitle>
          {{ confirmTitle }}
        </AlertDialogTitle>

        <AlertDialogDescription v-if="confirmDescription">
          {{ confirmDescription }}
        </AlertDialogDescription>
      </AlertDialogHeader>

      <AlertDialogFooter>
        <AlertDialogCancel>
          {{ cancelText }}
        </AlertDialogCancel>

        <Button
            type="button"
            :variant="confirmVariant === 'destructive' ? 'destructive' : 'default'"
            @click="handleConfirm"
        >
            {{ confirmText }}
        </Button>
      </AlertDialogFooter>
    </AlertDialogContent>
  </AlertDialog>
</template>