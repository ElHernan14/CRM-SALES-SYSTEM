<script setup lang="ts">
import {
  LayoutDashboard,
  Package,
  FileText,
  CreditCard,
  Settings,
  Store,
  ShoppingBag,
  Boxes,
  Users,
  BarChart3,
} from 'lucide-vue-next';

import NexoraBrand from '@/shared/components/brand/NexoraBrand.vue';

const navItems = [
  {
    label: 'Dashboard',
    to: '/erp/dashboard',
    icon: LayoutDashboard,
  },
  {
    label: 'Products',
    to: '/erp/products',
    icon: Package,
  },
  {
    label: 'Sales',
    to: '/erp/invoices',
    icon: FileText,
  },
  {
    label: 'Marketplace',
    to: '/erp/marketplace',
    icon: Store,
  },
  {
    label: 'Purchases',
    to: '/erp/purchases',
    icon: ShoppingBag,
  },
  {
    label: 'Inventory',
    to: '/erp/inventory',
    icon: Boxes,
  },
  {
    label: 'Customers',
    to: '/erp/customers',
    icon: Users,
  },
  {
    label: 'Analytics',
    to: '/erp/analytics',
    icon: BarChart3,
  },
  {
    label: 'Payments',
    to: '/erp/payments',
    icon: CreditCard,
  },
  {
    label: 'Settings',
    to: '/erp/settings',
    icon: Settings,
  },
];
</script>

<template>
  <aside class="hidden w-64 border-r border-border bg-card/60 lg:flex lg:flex-col">
    <div class="flex h-16 items-center border-b border-border px-6">
      <div>
        <NexoraBrand />
      </div>
    </div>

    <nav class="flex-1 space-y-1 p-3">
      <RouterLink
        v-for="item in navItems"
        :key="item.to"
        :to="item.to"
        class="group flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium text-muted-foreground transition hover:bg-muted hover:text-foreground"
        active-class="bg-primary text-primary-foreground hover:bg-primary hover:text-primary-foreground"
      >
        <component :is="item.icon" class="h-4 w-4 transition group-hover:text-foreground" />
        <span>{{ item.label }}</span>
      </RouterLink>
    </nav>

    <div class="border-t border-border p-3">
      <div class="rounded-lg border border-border bg-background p-3">
        <p class="text-xs font-medium text-foreground">Portfolio SaaS</p>
        <p class="mt-1 text-xs text-muted-foreground">Multi-tenant CRM & Commerce ERP.</p>
      </div>
    </div>
  </aside>
</template>
