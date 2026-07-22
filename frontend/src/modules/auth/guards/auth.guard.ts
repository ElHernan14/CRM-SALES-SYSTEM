import type { RouteLocationNormalized, Router } from 'vue-router';

import { useAuthStore } from '../stores/auth.store';

function getAuthenticatedHome(auth: ReturnType<typeof useAuthStore>) {
  return auth.user?.company_id ? '/erp/dashboard' : '/store';
}

export function setupAuthGuard(router: Router) {
  router.beforeEach((to: RouteLocationNormalized) => {
    const auth = useAuthStore();

    if (!auth.initialized) {
      return false;
    }

    const requiresAuth = to.matched.some((route) => route.meta.requiresAuth === true);

    const guestOnly = to.matched.some((route) => route.meta.guestOnly === true);

    const requiresCompany = to.matched.some((route) => route.meta.requiresCompany === true);

    /*
     * Ruta privada sin sesión.
     * Guardamos el destino para retomarlo después.
     */
    if (requiresAuth && !auth.isAuthenticated) {
      return {
        path: '/login',
        query: {
          redirect: to.fullPath,
        },
      };
    }

    /*
     * Login/register con una sesión ya activa.
     */
    if (guestOnly && auth.isAuthenticated) {
      return getAuthenticatedHome(auth);
    }

    /*
     * Un comprador personal no puede entrar al ERP.
     */
    if (requiresCompany && auth.isAuthenticated && !auth.user?.company_id) {
      return '/store';
    }

    return true;
  });
}
