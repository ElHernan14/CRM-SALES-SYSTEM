import { computed } from 'vue';

import { storeToRefs } from 'pinia';

import { useQuery } from '@tanstack/vue-query';

import { useAuthStore } from '@/modules/auth/stores/auth.store';

import { getDashboardOverview } from '../api/dashboard.api';

export function useDashboardOverview() {
  const auth = useAuthStore();
  const { user } = storeToRefs(auth);

  return useQuery({
    queryKey: computed(() => [
      'dashboard-overview',
      user.value?.userID ?? null,
      user.value?.company_id ?? null,
    ]),

    queryFn: getDashboardOverview,

    enabled: computed(() => {
      return auth.isAuthenticated && Boolean(user.value?.company_id);
    }),

    staleTime: 30_000,
  });
}
