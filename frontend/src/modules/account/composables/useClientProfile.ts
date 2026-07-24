import { computed } from 'vue';
import { storeToRefs } from 'pinia';
import { useQuery } from '@tanstack/vue-query';

import { useAuthStore } from '@/modules/auth/stores/auth.store';

import { getClientProfile } from '../api/account.api';

export function useClientProfile() {
  const auth = useAuthStore();
  const { user } = storeToRefs(auth);

  const clientId = computed(() => {
    return user.value?.client_id ?? null;
  });

  return useQuery({
    queryKey: computed(() => ['client-profile', user.value?.userID ?? null, clientId.value]),

    queryFn: () => {
      if (!clientId.value) {
        throw new Error('Authenticated user has no client profile');
      }

      return getClientProfile(clientId.value);
    },

    enabled: computed(() => {
      return auth.isAuthenticated && Boolean(clientId.value);
    }),

    staleTime: 60_000,
  });
}
