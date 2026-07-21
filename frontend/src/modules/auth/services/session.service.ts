import { queryClient } from '@/app/providers/query-client';
import { useAuthStore } from '../stores/auth.store';

export async function clearSession() {
  const authStore = useAuthStore();

  await queryClient.cancelQueries();
  queryClient.clear();

  authStore.logout();
}
