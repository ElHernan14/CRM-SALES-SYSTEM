import { createRouter, createWebHistory } from 'vue-router';

import { storeRoutes } from '@/modules/store/router/store.routes';
import { publicRoutes } from '@/app/router/routes/public.routes';
import { erpRoutes } from '@/modules/erp/router/erp.routes';

import { brand } from '@/shared/config/brand';

const routes = [...storeRoutes, ...publicRoutes, ...erpRoutes];

export const router = createRouter({
  history: createWebHistory(),

  routes,

  scrollBehavior(to, _from, savedPosition) {
    if (savedPosition) {
      return savedPosition;
    }

    if (to.hash) {
      return {
        el: to.hash,
        behavior: 'smooth',
      };
    }

    return {
      top: 0,
      behavior: 'smooth',
    };
  },
});

router.afterEach((to) => {
  const routeTitle = typeof to.meta.title === 'string' ? to.meta.title : '';

  document.title = routeTitle ? `${routeTitle} | ${brand.name}` : brand.browserTitle;
});
