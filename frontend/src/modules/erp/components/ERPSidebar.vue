<script setup lang="ts">
import { computed, watch } from 'vue';
import { useRoute } from 'vue-router';

import {
  BarChart3,
  Boxes,
  CreditCard,
  FileText,
  LayoutDashboard,
  Package,
  Settings,
  ShoppingBag,
  Store,
  Users,
} from 'lucide-vue-next';

import NexoraBrand from '@/shared/components/brand/NexoraBrand.vue';

import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
} from '@/components/ui/sheet';

const props = defineProps<{
  mobileOpen: boolean;
}>();

const emit = defineEmits<{
  'update:mobile-open': [value: boolean];
}>();

const route = useRoute();

const navItems = [
  {
    label: 'Dashboard',
    description: 'Business overview and attention center',
    to: '/erp/dashboard',
    icon: LayoutDashboard,
  },
  {
    label: 'Products',
    description: 'Catalog, pricing and product availability',
    to: '/erp/products',
    icon: Package,
  },
  {
    label: 'Sales',
    description: 'Customer invoices and seller operations',
    to: '/erp/sales',
    icon: FileText,
  },
  {
    label: 'Marketplace',
    description: 'Discover connected Nexora suppliers',
    to: '/erp/marketplace',
    icon: Store,
  },
  {
    label: 'Purchases',
    description: 'Organizational purchasing and suppliers',
    to: '/erp/purchases',
    icon: ShoppingBag,
  },
  {
    label: 'Inventory',
    description: 'Stock, reservations and availability',
    to: '/erp/inventory',
    icon: Boxes,
  },
  {
    label: 'Customers',
    description: 'Business clients and relationships',
    to: '/erp/customers',
    icon: Users,
  },
  {
    label: 'Analytics',
    description: 'Performance and operational insights',
    to: '/erp/analytics',
    icon: BarChart3,
  },
  {
    label: 'Payments',
    description: 'Payment records and financial activity',
    to: '/erp/payments',
    icon: CreditCard,
  },
  {
    label: 'Settings',
    description: 'Company workspace and personal profile',
    to: '/erp/settings',
    icon: Settings,
  },
] as const;

const currentPath = computed(() => {
  return route.path;
});

watch(
  () => route.fullPath,
  () => {
    if (props.mobileOpen) {
      emit('update:mobile-open', false);
    }
  }
);

function isItemActive(path: string) {
  if (path === '/erp/dashboard') {
    return currentPath.value === path;
  }

  return currentPath.value.startsWith(path);
}

function closeMobileNavigation() {
  emit('update:mobile-open', false);
}
</script>

<template>
  <!-- DESKTOP SIDEBAR -->
  <aside
    class="sticky top-0 hidden h-screen w-64 shrink-0 border-r border-border bg-card/60 lg:flex lg:flex-col"
  >
    <div class="flex h-16 shrink-0 items-center border-b border-border px-5">
      <NexoraBrand />
    </div>

    <nav
      class="min-h-0 flex-1 space-y-1 overflow-y-auto overscroll-contain p-3"
      aria-label="ERP navigation"
    >
      <RouterLink
        v-for="item in navItems"
        :key="item.to"
        :to="item.to"
        class="group flex items-center gap-3 rounded-xl px-3 py-2.5 text-sm font-medium text-muted-foreground transition-colors duration-200 hover:bg-muted hover:text-foreground"
        :class="{
          'bg-primary text-primary-foreground hover:bg-primary hover:text-primary-foreground':
            isItemActive(item.to),
        }"
      >
        <component
          :is="item.icon"
          class="h-4 w-4 shrink-0 transition-colors duration-200"
          :class="isItemActive(item.to) ? 'text-primary-foreground' : 'group-hover:text-foreground'"
        />

        <span class="truncate">
          {{ item.label }}
        </span>
      </RouterLink>
    </nav>

    <div class="shrink-0 border-t border-border p-3">
      <div class="rounded-xl border border-border bg-background p-3">
        <p class="text-xs font-semibold text-foreground">Nexora Workspace</p>

        <p class="mt-1 text-xs leading-5 text-muted-foreground">
          Multi-tenant commerce and business operations.
        </p>
      </div>
    </div>
  </aside>

  <!-- MOBILE / TABLET SIDEBAR -->
  <Sheet :open="mobileOpen" @update:open="emit('update:mobile-open', $event)">
    <SheetContent
      side="left"
      class="flex h-full w-[min(88vw,360px)] flex-col overflow-hidden p-0 lg:hidden"
    >
      <div
        class="shrink-0 border-b border-border bg-gradient-to-br from-primary/5 via-background to-background px-5 py-5"
      >
        <SheetHeader class="pr-8 text-left">
          <NexoraBrand />

          <SheetTitle class="sr-only"> Nexora ERP navigation </SheetTitle>

          <SheetDescription class="mt-3 text-left">
            Navigate your business workspace.
          </SheetDescription>
        </SheetHeader>
      </div>

      <nav
        class="min-h-0 flex-1 space-y-1 overflow-y-auto overscroll-contain px-3 py-4"
        aria-label="ERP mobile navigation"
      >
        <RouterLink
          v-for="item in navItems"
          :key="item.to"
          :to="item.to"
          class="group flex min-w-0 items-start gap-3 rounded-xl px-3 py-3 transition-colors duration-200"
          :class="
            isItemActive(item.to)
              ? 'bg-primary text-primary-foreground shadow-sm'
              : 'text-muted-foreground hover:bg-muted hover:text-foreground'
          "
          @click="closeMobileNavigation"
        >
          <div
            class="mt-0.5 flex h-9 w-9 shrink-0 items-center justify-center rounded-xl"
            :class="isItemActive(item.to) ? 'bg-primary-foreground/15' : 'bg-muted'"
          >
            <component :is="item.icon" class="h-4 w-4" />
          </div>

          <div class="min-w-0 flex-1">
            <p class="truncate text-sm font-semibold">
              {{ item.label }}
            </p>

            <p
              class="mt-1 line-clamp-2 text-xs leading-5"
              :class="
                isItemActive(item.to) ? 'text-primary-foreground/75' : 'text-muted-foreground'
              "
            >
              {{ item.description }}
            </p>
          </div>
        </RouterLink>
      </nav>

      <div
        class="shrink-0 border-t border-border bg-background px-4 pb-[calc(1rem+env(safe-area-inset-bottom))] pt-4"
      >
        <div class="rounded-xl border border-border bg-muted/20 p-4">
          <p class="text-sm font-semibold">Nexora Business OS</p>

          <p class="mt-1 text-xs leading-5 text-muted-foreground">
            ERP, Store and connected commerce in one workspace.
          </p>
        </div>
      </div>
    </SheetContent>
  </Sheet>
</template>
