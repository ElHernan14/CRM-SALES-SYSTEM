<script setup lang="ts">
import { storeToRefs } from 'pinia';
import { Button } from '@/components/ui/button';
import { Loader2 } from 'lucide-vue-next';

import {
  AlertDialog,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from '@/components/ui/alert-dialog';

import { useUiStore } from '@/shared/stores/ui.store';

const ui = useUiStore();

const {
  confirmLoading,
  confirmOpen,
  confirmTitle,
  confirmDescription,
  confirmText,
  cancelText,
  confirmVariant,
} = storeToRefs(ui);

async function handleConfirm() {
  await ui.confirm();
}
</script>

<template>
  <AlertDialog
    :open="confirmOpen"
    @update:open="
      (value) => {
        if (!value) ui.closeConfirm();
      }
    "
  >
    <AlertDialogContent class="z-10000">
      <AlertDialogHeader>
        <AlertDialogTitle>
          {{ confirmTitle }}
        </AlertDialogTitle>

        <AlertDialogDescription v-if="confirmDescription">
          {{ confirmDescription }}
        </AlertDialogDescription>
      </AlertDialogHeader>

      <AlertDialogFooter>
        <AlertDialogCancel :disabled="confirmLoading">
          {{ cancelText }}
        </AlertDialogCancel>

        <Button
          type="button"
          :variant="confirmVariant === 'destructive' ? 'destructive' : 'default'"
          :disabled="confirmLoading"
          @click="handleConfirm"
        >
          <Loader2 v-if="confirmLoading" class="mr-2 h-4 w-4 animate-spin" />

          {{ confirmLoading ? 'Processing...' : confirmText }}
        </Button>
      </AlertDialogFooter>
    </AlertDialogContent>
  </AlertDialog>
</template>
