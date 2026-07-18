import { computed, type Ref } from 'vue';

import { useQuery } from '@tanstack/vue-query';

import { getCompanyPurchases } from '../api/purchases.api';
import type { GetPurchasesRequest } from '../types/purchase.types';

export function usePurchases(params: Ref<GetPurchasesRequest>) {
  return useQuery({
    queryKey: computed(() => ['purchases', params.value]),

    queryFn: () => getCompanyPurchases(params.value),
  });
}
