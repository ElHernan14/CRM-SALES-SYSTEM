import type { RouteRecordRaw } from 'vue-router';

const StoreLayout = () => import('@/modules/store/layouts/StoreLayout.vue');

const StoreHomePage = () => import('@/modules/store/pages/StoreHomePage.vue');

const StoreCatalogPage = () => import('@/modules/store/pages/StoreCatalogPage.vue');

const StoreProductPage = () => import('@/modules/store/pages/StoreProductPage.vue');

const StoreCheckoutPage = () => import('@/modules/store/pages/StoreCheckoutPage.vue');

const StoreOrderSuccessPage = () => import('@/modules/store/pages/StoreOrderSuccessPage.vue');

const StorePurchasesPage = () => import('@/modules/store/pages/StorePurchasesPage.vue');

export const storeRoutes: RouteRecordRaw[] = [
  {
    path: '/store',
    component: StoreLayout,
    children: [
      {
        path: '',
        name: 'store-home',
        component: StoreHomePage,
      },
      {
        path: 'catalog',
        name: 'store-catalog',
        component: StoreCatalogPage,
      },
      {
        path: 'products/:productId',
        name: 'store-product',
        component: StoreProductPage,
        props: true,
      },
      {
        path: 'checkout',
        name: 'store-checkout',
        component: StoreCheckoutPage,
        meta: {
          authOnly: true,
        },
      },
      {
        path: 'order-success/:invoiceId',
        name: 'store-order-success',
        component: StoreOrderSuccessPage,
        props: true,
        meta: {
          authOnly: true,
        },
      },
      {
        path: 'purchases',
        name: 'store-purchases',
        component: StorePurchasesPage,
        meta: {
          authOnly: true,
        },
      },
    ],
  },
];
