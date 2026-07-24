<script setup lang="ts">
import { computed, onUnmounted, ref, watch } from 'vue';

import { ImageIcon, ImagePlus, Loader2, Upload } from 'lucide-vue-next';

import { toast } from 'vue-sonner';

import { Button } from '@/components/ui/button';

import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
} from '@/components/ui/sheet';

const props = withDefaults(
  defineProps<{
    open: boolean;

    title: string;
    description: string;

    assetType: 'logo' | 'cover';

    currentImageUrl?: string | null;

    uploading?: boolean;
  }>(),
  {
    currentImageUrl: null,
    uploading: false,
  }
);

const emit = defineEmits<{
  'update:open': [value: boolean];
  upload: [file: File];
}>();

const inputRef = ref<HTMLInputElement | null>(null);

const selectedFile = ref<File | null>(null);

const localPreview = ref<string | null>(null);

const previewUrl = computed(() => {
  return localPreview.value ?? props.currentImageUrl ?? null;
});

const recommendedText = computed(() => {
  return props.assetType === 'logo'
    ? 'A square PNG, JPG or WebP works best.'
    : 'Use a wide image with enough space around the main subject.';
});

watch(
  () => props.open,
  (open) => {
    if (!open) {
      resetSelection();
    }
  }
);

function openPicker() {
  inputRef.value?.click();
}

function selectFile(event: Event) {
  const input = event.target as HTMLInputElement;

  const file = input.files?.[0] ?? null;

  if (!file) return;

  if (!file.type.startsWith('image/')) {
    toast.error('Select a valid image file');
    return;
  }

  const maxSize = 5 * 1024 * 1024;

  if (file.size > maxSize) {
    toast.error('The image cannot exceed 5 MB');

    return;
  }

  revokePreview();

  selectedFile.value = file;

  localPreview.value = URL.createObjectURL(file);
}

function revokePreview() {
  if (!localPreview.value) return;

  URL.revokeObjectURL(localPreview.value);

  localPreview.value = null;
}

function resetSelection() {
  revokePreview();

  selectedFile.value = null;

  if (inputRef.value) {
    inputRef.value.value = '';
  }
}

function submit() {
  if (!selectedFile.value) {
    toast.error('Select an image first');
    return;
  }

  emit('upload', selectedFile.value);
}

onUnmounted(() => {
  revokePreview();
});

watch(
  () => props.open,
  (open) => {
    if (!open) {
      resetSelection();
    }
  }
);
</script>

<template>
  <Sheet :open="open" @update:open="emit('update:open', $event)">
    <SheetContent class="w-full overflow-y-auto sm:max-w-xl">
      <SheetHeader class="text-left">
        <div class="flex items-center gap-3">
          <div class="flex h-11 w-11 items-center justify-center rounded-xl bg-primary/10">
            <ImagePlus class="h-5 w-5 text-primary" />
          </div>

          <div>
            <SheetTitle>
              {{ title }}
            </SheetTitle>

            <SheetDescription class="mt-1">
              {{ description }}
            </SheetDescription>
          </div>
        </div>
      </SheetHeader>

      <div class="mt-8 space-y-6">
        <input
          ref="inputRef"
          type="file"
          accept="image/png,image/jpeg,image/webp"
          class="hidden"
          @change="selectFile"
        />

        <button
          type="button"
          class="group relative flex w-full items-center justify-center overflow-hidden rounded-[1.75rem] border border-dashed border-border bg-muted/20 transition hover:border-primary/40"
          :class="assetType === 'logo' ? 'min-h-72' : 'aspect-[16/7]'"
          :disabled="uploading"
          @click="openPicker"
        >
          <img
            v-if="previewUrl"
            :src="previewUrl"
            alt="Company asset preview"
            :class="
              assetType === 'logo'
                ? 'h-44 w-44 rounded-[2rem] border border-border bg-background object-contain p-4 shadow-xl'
                : 'absolute inset-0 h-full w-full object-cover'
            "
          />

          <div v-if="!previewUrl" class="flex flex-col items-center gap-3 text-muted-foreground">
            <div class="flex h-14 w-14 items-center justify-center rounded-2xl bg-primary/10">
              <ImageIcon class="h-6 w-6 text-primary" />
            </div>

            <div>
              <p class="text-sm font-semibold text-foreground">Select an image</p>

              <p class="mt-1 text-xs">PNG, JPG or WebP</p>
            </div>
          </div>

          <div
            v-if="previewUrl"
            class="absolute inset-0 flex items-center justify-center bg-background/70 opacity-0 backdrop-blur-sm transition group-hover:opacity-100"
          >
            <div
              class="rounded-full border border-border bg-background px-4 py-2 text-sm font-medium shadow-md"
            >
              <Upload class="mr-2 inline h-4 w-4" />
              Choose another image
            </div>
          </div>
        </button>

        <div class="rounded-xl border border-border bg-muted/20 p-4">
          <p class="text-sm font-medium">Image recommendations</p>

          <p class="mt-2 text-xs leading-5 text-muted-foreground">
            {{ recommendedText }}
            Maximum file size: 5 MB.
          </p>
        </div>

        <div class="grid gap-3 sm:grid-cols-2">
          <Button
            type="button"
            variant="outline"
            class="rounded-full"
            :disabled="uploading"
            @click="openPicker"
          >
            <Upload class="mr-2 h-4 w-4" />
            Select image
          </Button>

          <Button
            type="button"
            class="rounded-full"
            :disabled="!selectedFile || uploading"
            @click="submit"
          >
            <Loader2 v-if="uploading" class="mr-2 h-4 w-4 animate-spin" />

            <Upload v-else class="mr-2 h-4 w-4" />

            {{ uploading ? 'Uploading image...' : 'Save image' }}
          </Button>
        </div>
      </div>
    </SheetContent>
  </Sheet>
</template>
