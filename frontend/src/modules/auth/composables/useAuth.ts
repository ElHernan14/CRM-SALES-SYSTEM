import { useAuthStore } from '@/modules/auth/stores/auth.store';
import { login as loginApi, me } from '@/modules/auth/api/auth.api';
import { router } from '@/app/router';
import type { LoginRequest, TenantContext } from '@/modules/auth/types/auth.types';
import { queryClient } from '@/app/providers/query-client';
import type { LocationQueryValue } from 'vue-router';

function resolveRedirect(redirect: LocationQueryValue | LocationQueryValue[]) {
  if (Array.isArray(redirect)) {
    return redirect[0] ?? null;
  }

  return redirect ?? null;
}

export async function login(credentials: LoginRequest) {
  const auth = useAuthStore();

  try {
    await queryClient.cancelQueries();
    queryClient.clear();

    const response = await loginApi(credentials);

    auth.setToken(response.token);

    const user: TenantContext = await me();

    auth.setUser(user);
    auth.setInitialized(true);

    const requestedRedirect = resolveRedirect(router.currentRoute.value.query.redirect);

    /*
     * Respetar el redirect sólo cuando coincide
     * con el tipo de usuario autenticado.
     */
    if (requestedRedirect) {
      const isErpDestination = requestedRedirect.startsWith('/erp');

      const canUseRequestedDestination = !isErpDestination || Boolean(user.company_id);

      if (canUseRequestedDestination) {
        await router.replace(requestedRedirect);

        return;
      }
    }

    await router.replace(user.company_id ? '/erp/dashboard' : '/store');
  } catch (error: unknown) {
    console.error('Login error:', error);
    throw error;
  }
}

export async function logout() {
  const auth = useAuthStore();
  /*
   * Detiene requests que todavía podrían escribir datos
   * del tenant anterior en el cache.
   */
  await queryClient.cancelQueries();
  /*
   * Elimina queries y mutations de la sesión anterior.
   */
  queryClient.clear();

  auth.logout();
  router.push('/login');
}

export async function bootstrap() {
  const auth = useAuthStore();

  // 1. Leer token
  const token = localStorage.getItem('access_token');

  if (!token) {
    // 2. No existe → inicializar
    auth.setInitialized(true);
    return;
  }

  try {
    // 3. Setear token
    auth.setToken(token);

    // 4. Llamar /auth/me
    const user = await me();

    // 5. Guardar usuario
    auth.setUser(user);
  } catch (error) {
    console.error('Bootstrap error:', error);
    // Si falla, limpiar sesión
    auth.logout();
  } finally {
    // 6. Inicializar
    auth.setInitialized(true);

    // Si ya estoy autenticado y estoy en /login, redirigir
    if (auth.isAuthenticated && router.currentRoute.value.name === 'login') {
      router.replace('/erp/dashboard');
    }
  }
}
