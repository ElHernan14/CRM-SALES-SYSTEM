import { http } from './http';
import { router } from '@/app/router';
import { clearSession } from '@/modules/auth/services/session.service';

let handlingUnauthorized = false;

export function setupInterceptors() {
  http.interceptors.request.use(
    (config) => {
      const token = localStorage.getItem('access_token');

      if (token) {
        config.headers.Authorization = `Bearer ${token}`;
      }

      return config;
    },
    (error) => Promise.reject(error)
  );

  http.interceptors.response.use(
    (response) => response,

    async (error) => {
      if (error.response?.status === 401 && !handlingUnauthorized) {
        handlingUnauthorized = true;

        try {
          await clearSession();

          const currentRoute = router.currentRoute.value;

          const publicAuthRoutes = ['/login', '/register'];

          if (!publicAuthRoutes.some((path) => currentRoute.path.startsWith(path))) {
            await router.replace({
              path: '/login',
              query: {
                redirect: currentRoute.fullPath,
              },
            });
          }
        } finally {
          handlingUnauthorized = false;
        }
      }

      return Promise.reject(error);
    }
  );
}
