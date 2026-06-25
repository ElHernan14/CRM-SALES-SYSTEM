<script setup lang="ts">
import { Toaster } from "vue-sonner"
import GlobalLoader from "@/shared/components/GlobalLoader.vue"
</script>

<template>
  <router-view />

  <GlobalLoader />

  <Toaster
    position="top-center"
    richColors
    :duration="2000"
  />
</template>