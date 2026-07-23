import { computed } from 'vue';

import { storeToRefs } from 'pinia';

import { useQuery } from '@tanstack/vue-query';

import { useAuthStore } from '@/modules/auth/stores/auth.store';

import { getCompanyMe } from '../api/company.api';

export function useCompanyMe() {
  const auth = useAuthStore();
  const { user } = storeToRefs(auth);

  return useQuery({
    queryKey: computed(() => ['company-me', user.value?.company_id ?? null]),

    queryFn: getCompanyMe,

    enabled: computed(() => {
      return auth.isAuthenticated && Boolean(user.value?.company_id);
    }),

    staleTime: 60_000,
  });
}
