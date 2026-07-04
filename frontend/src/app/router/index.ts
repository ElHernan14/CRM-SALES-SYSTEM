import { createRouter, createWebHistory } from 'vue-router';
import PublicLayout from '@/app/layouts/PublicLayout.vue';
import ERPLayout from '@/modules/erp/layouts/ERPLayout.vue';
import DashboardPage from '@/modules/erp/pages/DashboardPage.vue';
import LoginPage from '@/modules/auth/pages/LoginPage.vue';

const ProductsPage = () => import('@/modules/products/pages/ProductsPage.vue');
const InvoicesPage = () => import('@/modules/erp/pages/InvoicesPage.vue');
const PaymentsPage = () => import('@/modules/erp/pages/PaymentsPage.vue');
const SettingsPage = () => import('@/modules/erp/pages/SettingsPage.vue');
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
    ],
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
