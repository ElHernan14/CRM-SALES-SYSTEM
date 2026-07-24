import { computed, type ComputedRef } from 'vue';

import { useQuery } from '@tanstack/vue-query';

import { getStoreBusiness } from '../api/store-businesses.api';

export function useStoreBusiness(businessId: ComputedRef<number | null>) {
  return useQuery({
    queryKey: computed(() => ['store-business', businessId.value]),

    queryFn: () => {
      if (!businessId.value) {
        throw new Error('Business ID is required');
      }

      return getStoreBusiness(businessId.value);
    },

    enabled: computed(() => {
      return Boolean(businessId.value);
    }),

    staleTime: 60_000,
  });
}
