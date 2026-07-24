import type { RouteRecordRaw } from 'vue-router';

const StoreLayout = () => import('@/modules/store/layouts/StoreLayout.vue');

const StoreHomePage = () => import('@/modules/store/pages/StoreHomePage.vue');

const StoreCatalogPage = () => import('@/modules/store/pages/StoreCatalogPage.vue');

const StoreProductPage = () => import('@/modules/store/pages/StoreProductPage.vue');

const StoreCheckoutPage = () => import('@/modules/store/pages/StoreCheckoutPage.vue');

const StoreOrderSuccessPage = () => import('@/modules/store/pages/StoreOrderSuccessPage.vue');

const StorePurchasesPage = () => import('@/modules/store/pages/StorePurchasesPage.vue');

const StoreBusinessesPage = () => import('@/modules/store/pages/StoreBusinessesPage.vue');

export const storeRoutes: RouteRecordRaw[] = [
  {
    path: '/store',
    component: StoreLayout,
    children: [
      {
        path: '',
        name: 'store-home',
        component: StoreHomePage,
        meta: {
          surface: 'store',
          title: 'Nexora Store',
        },
      },
      {
        path: 'catalog',
        name: 'store-catalog',
        component: StoreCatalogPage,
        meta: {
          surface: 'store',
          title: 'Catalog',
        },
      },
      {
        path: 'products/:productId',
        name: 'store-product',
        component: StoreProductPage,
        props: true,
        meta: {
          surface: 'store',
          title: 'Product',
        },
      },
      {
        path: 'checkout',
        name: 'store-checkout',
        component: StoreCheckoutPage,
        meta: {
          requiresAuth: true,
        },
      },
      {
        path: 'purchases',
        name: 'store-purchases',
        component: StorePurchasesPage,
        meta: {
          requiresAuth: true,
        },
      },
      {
        path: 'order-success/:invoiceId',
        name: 'store-order-success',
        component: StoreOrderSuccessPage,
        props: true,
        meta: {
          requiresAuth: true,
        },
      },
      {
        path: 'account',
        name: 'store-account',
        component: () => import('@/modules/account/pages/StoreAccountPage.vue'),
        meta: {
          requiresAuth: true,
        },
      },
      {
        path: 'businesses',
        name: 'store-businesses',
        component: StoreBusinessesPage,
        meta: {
          surface: 'store',
          title: 'Businesses',
        },
      },
    ],
  },
];
