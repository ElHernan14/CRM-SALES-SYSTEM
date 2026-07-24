import { computed, type ComputedRef } from 'vue';

import { useQuery } from '@tanstack/vue-query';

import { getStoreBusinesses } from '../api/store-businesses.api';

import type { StoreBusinessesParams } from '../types/store-business.types';

export function useStoreBusinesses(params: ComputedRef<StoreBusinessesParams>) {
  return useQuery({
    queryKey: computed(() => ['store-businesses', params.value]),

    queryFn: () => getStoreBusinesses(params.value),

    staleTime: 30_000,
  });
}
