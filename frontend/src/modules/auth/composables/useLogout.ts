import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { toast } from 'vue-sonner';

import { clearSession } from '../services/session.service';

export function useLogout() {
  const router = useRouter();

  const isLoggingOut = ref(false);

  async function logout() {
    if (isLoggingOut.value) return;

    try {
      isLoggingOut.value = true;

      await clearSession();

      await router.replace('/store');

      toast.success('Signed out successfully');
    } catch (error) {
      console.error('Logout error:', error);

      toast.error('Unable to sign out');
    } finally {
      isLoggingOut.value = false;
    }
  }

  return {
    logout,
    isLoggingOut,
  };
}
