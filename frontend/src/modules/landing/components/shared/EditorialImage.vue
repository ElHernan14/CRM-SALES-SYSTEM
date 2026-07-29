<script setup lang="ts">
import { ImageIcon } from 'lucide-vue-next';

withDefaults(
  defineProps<{
    src: string;
    alt: string;
    eyebrow?: string;
    title?: string;
    description?: string;
    position?: string;
    priority?: boolean;
    compact?: boolean;
  }>(),
  {
    eyebrow: '',
    title: '',
    description: '',
    position: 'center',
    priority: false,
    compact: false,
  }
);
</script>

<template>
  <figure
    class="nexora-card-interactive group relative isolate overflow-hidden border border-border bg-card shadow-xl"
    :class="compact ? 'min-h-[320px] rounded-[1.75rem]' : 'min-h-[420px] rounded-[2rem]'"
  >
    <img
      :src="src"
      :alt="alt"
      class="absolute inset-0 h-full w-full object-cover transition duration-700 ease-out group-hover:scale-[1.025]"
      :style="{
        objectPosition: position,
      }"
      :loading="priority ? 'eager' : 'lazy'"
      :fetchpriority="priority ? 'high' : 'auto'"
      decoding="async"
    />

    <!-- Unifica visualmente todas las fotografías -->
    <div
      class="absolute inset-0 bg-gradient-to-br from-slate-950/20 via-transparent to-indigo-950/35"
    />

    <!-- Asegura legibilidad inferior -->
    <div
      class="absolute inset-0 bg-gradient-to-t from-slate-950/90 via-slate-950/20 to-transparent"
    />

    <!-- Halo propio de Nexora -->
    <div
      class="pointer-events-none absolute -right-20 -top-20 h-56 w-56 rounded-full bg-violet-500/20 blur-3xl"
    />

    <figcaption
      v-if="eyebrow || title || description"
      class="absolute inset-x-0 bottom-0 p-5 text-white sm:p-7"
    >
      <p v-if="eyebrow" class="text-[11px] font-semibold uppercase tracking-[0.18em] text-blue-200">
        {{ eyebrow }}
      </p>

      <h3 v-if="title" class="mt-2 max-w-xl text-xl font-semibold tracking-tight sm:text-2xl">
        {{ title }}
      </h3>

      <p v-if="description" class="mt-2 max-w-xl text-sm leading-6 text-slate-200/90">
        {{ description }}
      </p>
    </figcaption>

    <div
      class="pointer-events-none absolute inset-0 rounded-[inherit] ring-1 ring-inset ring-white/10"
    />

    <noscript>
      <div class="flex min-h-[320px] items-center justify-center bg-muted">
        <ImageIcon class="h-10 w-10 text-muted-foreground" />
      </div>
    </noscript>
  </figure>
</template>
