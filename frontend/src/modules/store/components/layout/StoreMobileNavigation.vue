<script setup lang="ts">
import { Compass, Home, ShoppingBag, UserRound } from 'lucide-vue-next';
</script>

<template>
  <nav
    class="fixed inset-x-0 bottom-0 z-40 border-t border-border bg-background/95 px-3 py-2 backdrop-blur-xl md:hidden"
  >
    <div class="grid grid-cols-4 gap-1">
      <RouterLink
        to="/store"
        class="flex flex-col items-center gap-1 rounded-xl px-2 py-2 text-xs text-muted-foreground transition hover:bg-muted hover:text-foreground"
      >
        <Home class="h-5 w-5" />
        Home
      </RouterLink>

      <RouterLink
        to="/store/catalog"
        class="flex flex-col items-center gap-1 rounded-xl px-2 py-2 text-xs text-muted-foreground transition hover:bg-muted hover:text-foreground"
      >
        <Compass class="h-5 w-5" />
        Explore
      </RouterLink>

      <button
        class="flex flex-col items-center gap-1 rounded-xl px-2 py-2 text-xs text-muted-foreground transition hover:bg-muted hover:text-foreground"
      >
        <ShoppingBag class="h-5 w-5" />
        Cart
      </button>

      <RouterLink
        to="/store/purchases"
        class="flex flex-col items-center gap-1 rounded-xl px-2 py-2 text-xs text-muted-foreground transition hover:bg-muted hover:text-foreground"
      >
        <UserRound class="h-5 w-5" />
        Account
      </RouterLink>
    </div>
  </nav>
</template>
