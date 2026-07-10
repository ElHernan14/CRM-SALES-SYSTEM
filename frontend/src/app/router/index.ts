import { createRouter, createWebHistory } from 'vue-router';
import PublicLayout from '@/app/layouts/PublicLayout.vue';
import ERPLayout from '@/modules/erp/layouts/ERPLayout.vue';
import DashboardPage from '@/modules/erp/pages/DashboardPage.vue';
import LoginPage from '@/modules/auth/pages/LoginPage.vue';

const ProductsPage = () => import('@/modules/products/pages/ProductsPage.vue');
const InvoicesPage = () => import('@/modules/invoices/pages/InvoicesPage.vue');
const PaymentsPage = () => import('@/modules/erp/pages/PaymentsPage.vue');
const SettingsPage = () => import('@/modules/erp/pages/SettingsPage.vue');
const MarketplacePage = () => import('@/modules/marketplace/pages/MarketplacePage.vue');
const PurchasesPage = () => import('@/modules/erp/pages/PurchasesPage.vue');

const InventoryPage = () => import('@/modules/erp/pages/InventoryPage.vue');

const CustomersPage = () => import('@/modules/erp/pages/CustomersPage.vue');

const AnalyticsPage = () => import('@/modules/erp/pages/AnalyticsPage.vue');
const routes = [
  {
    path: '/',
    component: PublicLayout,
    children: [
      // {
      //   path: "",
      //   component: LandingPage
      // },
      {
        path: '',
        redirect: { name: 'login' },
      },
      {
        path: 'login',
        name: 'login',
        component: LoginPage,
        meta: {
          guestOnly: true,
        },
      },
    ],
  },
  {
    path: '/erp',
    component: ERPLayout,
    meta: {
      requiresAuth: true,
    },
    children: [
      {
        path: 'dashboard',
        name: 'erp-dashboard',
        component: DashboardPage,
      },
      {
        path: 'products',
        name: 'erp-products',
        component: ProductsPage,
      },
      {
        path: 'invoices',
        name: 'erp-invoices',
        component: InvoicesPage,
      },
      {
        path: 'payments',
        name: 'erp-payments',
        component: PaymentsPage,
      },
      {
        path: 'settings',
        name: 'erp-settings',
        component: SettingsPage,
      },
      {
        path: 'marketplace',
        name: 'erp-marketplace',
        component: MarketplacePage,
      },
      {
        path: 'purchases',
        name: 'erp-purchases',
        component: PurchasesPage,
      },
      {
        path: 'inventory',
        name: 'erp-inventory',
        component: InventoryPage,
      },
      {
        path: 'customers',
        name: 'erp-customers',
        component: CustomersPage,
      },
      {
        path: 'analytics',
        name: 'erp-analytics',
        component: AnalyticsPage,
      },
    ],
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
