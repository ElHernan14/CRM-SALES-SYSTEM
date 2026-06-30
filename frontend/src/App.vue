<script setup lang="ts">
import { Toaster } from "vue-sonner"
import GlobalLoader from "@/shared/components/GlobalLoader.vue"
import GlobalConfirmDialog from "@/shared/components/GlobalConfirmDialog.vue"
</script>

<template>
  <router-view />

  <!-- Global loader que muestra el estado de carga en toda la aplicación -->
  <GlobalLoader />

  <!-- Diálogo de confirmación global reutilizable -->
  <GlobalConfirmDialog />

  <Toaster
    position="top-center"
    richColors
    :duration="2000"
  />
</template>