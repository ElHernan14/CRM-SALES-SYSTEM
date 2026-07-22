import { computed } from 'vue';

import { storeToRefs } from 'pinia';

import { useQuery } from '@tanstack/vue-query';

import { useAuthStore } from '@/modules/auth/stores/auth.store';

import { getStoreCarts } from '../api/store-cart.api';

export function useStoreCarts() {
  const auth = useAuthStore();

  const { user } = storeToRefs(auth);

  const authenticatedUserId = computed(() => {
    return user.value?.userID ?? null;
  });

  return useQuery({
    queryKey: computed(() => ['b2c-store-carts', authenticatedUserId.value]),

    queryFn: getStoreCarts,

    enabled: computed(() => {
      return auth.isAuthenticated;
    }),

    staleTime: 10_000,
  });
}
