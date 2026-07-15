<script setup lang="ts">
import { computed, ref, watch } from 'vue';

import { ImagePlus, Loader2, Upload, X } from 'lucide-vue-next';

import { toast } from 'vue-sonner';

import { Button } from '@/components/ui/button';

import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
} from '@/components/ui/sheet';

import { useUploadProductImage } from '../composables/useUploadProductImage';

const props = defineProps<{
  open: boolean;
  productId: number | null;
}>();

const emit = defineEmits<{
  'update:open': [value: boolean];
  uploaded: [];
}>();

const fileInput = ref<HTMLInputElement | null>(null);
const selectedFile = ref<File | null>(null);
const previewUrl = ref<string | null>(null);

const uploadMutation = useUploadProductImage();

const isUploading = computed(() => uploadMutation.isPending.value);

watch(
  () => props.open,
  (open) => {
    if (!open) {
      resetSelection();
    }
  }
);

function selectFile(file: File | null) {
  if (!file) return;

  if (!file.type.startsWith('image/')) {
    toast.error('Select a valid image file');
    return;
  }

  const maxSize = 5 * 1024 * 1024;

  if (file.size > maxSize) {
    toast.error('Image must be smaller than 5 MB');
    return;
  }

  if (previewUrl.value) {
    URL.revokeObjectURL(previewUrl.value);
  }

  selectedFile.value = file;
  previewUrl.value = URL.createObjectURL(file);
}

function handleInputChange(event: Event) {
  const target = event.target as HTMLInputElement;

  selectFile(target.files?.[0] ?? null);
}

function handleDrop(event: DragEvent) {
  event.preventDefault();

  selectFile(event.dataTransfer?.files?.[0] ?? null);
}

function resetSelection() {
  if (previewUrl.value) {
    URL.revokeObjectURL(previewUrl.value);
  }

  selectedFile.value = null;
  previewUrl.value = null;

  if (fileInput.value) {
    fileInput.value.value = '';
  }
}

async function uploadImage() {
  if (!props.productId) {
    toast.error('Missing product context');
    return;
  }

  if (!selectedFile.value) {
    toast.error('Select an image first');
    return;
  }

  try {
    await uploadMutation.mutateAsync({
      productId: props.productId,
      image: selectedFile.value,
    });

    toast.success('Product image uploaded');

    resetSelection();

    emit('uploaded');
    emit('update:open', false);
  } catch (error: any) {
    const message = error?.response?.data?.errorMessage ?? 'Failed to upload product image';

    toast.error(message);
  }
}
</script>

<template>
  <Sheet :open="open" @update:open="emit('update:open', $event)">
    <SheetContent class="w-full sm:max-w-xl">
      <SheetHeader>
        <SheetTitle class="flex items-center gap-2">
          <ImagePlus class="h-5 w-5" />
          Product image
        </SheetTitle>

        <SheetDescription>
          Upload the catalog image that will represent this product in the ERP and marketplace.
        </SheetDescription>
      </SheetHeader>

      <div class="mt-6 space-y-6">
        <input
          ref="fileInput"
          type="file"
          accept="image/*"
          class="hidden"
          @change="handleInputChange"
        />

        <button
          type="button"
          class="group relative flex min-h-64 w-full items-center justify-center overflow-hidden rounded-2xl border-2 border-dashed border-border bg-muted/20 transition hover:border-primary/40 hover:bg-muted/40"
          @click="fileInput?.click()"
          @dragover.prevent
          @drop="handleDrop"
        >
          <img
            v-if="previewUrl"
            :src="previewUrl"
            alt="Product preview"
            class="absolute inset-0 h-full w-full object-cover"
          />

          <div
            v-if="previewUrl"
            class="absolute inset-0 bg-gradient-to-t from-background/80 via-transparent to-transparent"
          />

          <div v-if="!previewUrl" class="flex flex-col items-center px-6 text-center">
            <div class="rounded-2xl border border-border bg-background p-4 shadow-sm">
              <Upload class="h-6 w-6 text-muted-foreground" />
            </div>

            <p class="mt-4 text-sm font-medium text-foreground">Drop an image here</p>

            <p class="mt-1 text-xs text-muted-foreground">
              Or click to browse. Maximum size: 5 MB.
            </p>
          </div>

          <div
            v-else
            class="absolute bottom-4 left-4 right-4 flex items-center justify-between rounded-xl border border-border bg-background/90 p-3 text-left shadow-md backdrop-blur"
          >
            <div class="min-w-0">
              <p class="truncate text-sm font-medium text-foreground">
                {{ selectedFile?.name }}
              </p>

              <p class="text-xs text-muted-foreground">Image ready to upload</p>
            </div>

            <Button
              type="button"
              variant="ghost"
              size="icon"
              class="shrink-0"
              @click.stop="resetSelection"
            >
              <X class="h-4 w-4" />
            </Button>
          </div>
        </button>

        <div class="rounded-xl border border-border bg-muted/30 p-4">
          <p class="text-sm font-medium text-foreground">Recommended format</p>

          <p class="mt-1 text-xs leading-5 text-muted-foreground">
            Use a landscape or square image with good lighting. The marketplace will crop it
            automatically when necessary.
          </p>
        </div>

        <div class="flex justify-end gap-2">
          <Button
            type="button"
            variant="outline"
            :disabled="isUploading"
            @click="emit('update:open', false)"
          >
            Skip for now
          </Button>

          <Button type="button" :disabled="!selectedFile || isUploading" @click="uploadImage">
            <Loader2 v-if="isUploading" class="mr-2 h-4 w-4 animate-spin" />

            <Upload v-else class="mr-2 h-4 w-4" />

            {{ isUploading ? 'Uploading...' : 'Upload image' }}
          </Button>
        </div>
      </div>
    </SheetContent>
  </Sheet>
</template>
