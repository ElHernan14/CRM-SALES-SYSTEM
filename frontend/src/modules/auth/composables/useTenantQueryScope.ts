import { computed } from 'vue';
import { storeToRefs } from 'pinia';

import { useAuthStore } from '../stores/auth.store';

export function useTenantQueryScope() {
  const authStore = useAuthStore();
  const { user } = storeToRefs(authStore);

  return computed(() => ({
    userId: user.value?.userID ?? null,
    companyId: user.value?.company_id ?? null,
    clientId: user.value?.client_id ?? null,
  }));
}
