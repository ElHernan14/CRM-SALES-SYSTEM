import { computed, type Ref } from 'vue';

import { useQuery } from '@tanstack/vue-query';

import { getPurchaseCart } from '../api/purchase-cart.api';

export function usePurchaseCart(sellerCompanyId: Ref<number | null>) {
  return useQuery({
    queryKey: computed(() => ['purchase-cart', sellerCompanyId.value]),

    queryFn: () => getPurchaseCart(sellerCompanyId.value as number),

    enabled: computed(() => !!sellerCompanyId.value),
  });
}
