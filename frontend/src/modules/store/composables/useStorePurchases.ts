import { computed, type Ref } from 'vue';

import { storeToRefs } from 'pinia';

import { useQuery } from '@tanstack/vue-query';

import { useAuthStore } from '@/modules/auth/stores/auth.store';

import { getStorePurchases } from '../api/store-purchases.api';

import type { GetStorePurchasesRequest } from '../types/store-purchase.types';

export function useStorePurchases(params: Ref<GetStorePurchasesRequest>) {
  const auth = useAuthStore();
  const { user } = storeToRefs(auth);

  return useQuery({
    queryKey: computed(() => [
      'store-purchases',
      user.value?.userID ?? null,
      user.value?.client_id ?? null,
      params.value,
    ]),

    queryFn: () => getStorePurchases(params.value),

    enabled: computed(() => {
      return auth.isAuthenticated && Boolean(user.value?.client_id);
    }),

    staleTime: 15_000,
  });
}
