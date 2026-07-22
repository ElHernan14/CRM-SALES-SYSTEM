import { createRouter, createWebHistory } from 'vue-router';
import { storeRoutes } from '@/modules/store/router/store.routes';
import { publicRoutes } from '@/app/router/routes/public.routes';
import { erpRoutes } from '@/modules/erp/router/erp.routes';

const routes = [...storeRoutes, ...publicRoutes, ...erpRoutes];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
