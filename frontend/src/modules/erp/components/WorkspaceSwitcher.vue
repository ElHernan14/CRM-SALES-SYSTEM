<script setup lang="ts">
import { Building2, ChevronsUpDown } from 'lucide-vue-next';
import { Button } from '@/components/ui/button';
</script>

<template>
  <Button variant="outline" class="h-10 justify-between gap-3 px-3">
    <div class="flex items-center gap-2">
      <div class="rounded-md border border-border bg-muted p-1.5">
        <Building2 class="h-4 w-4 text-muted-foreground" />
      </div>

      <div class="hidden text-left sm:block">
        <p class="text-sm font-medium leading-none">Main Workspace</p>
        <p class="mt-1 text-xs text-muted-foreground">ERP tenant</p>
      </div>
    </div>

    <ChevronsUpDown class="h-4 w-4 text-muted-foreground" />
  </Button>
</template>
