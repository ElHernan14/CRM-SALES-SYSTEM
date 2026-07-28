import { computed, type Ref } from 'vue';
import { useQuery } from '@tanstack/vue-query';

import { getMarketplacePurchaseCart } from '../api/marketplace-purchase-cart.api';

export function useMarketplacePurchaseCart(sellerCompanyId: Ref<number | null>) {
  return useQuery({
    queryKey: computed(() => ['marketplace-purchase-cart', sellerCompanyId.value]),

    queryFn: () => {
      if (!sellerCompanyId.value) {
        throw new Error('Missing seller company ID');
      }

      return getMarketplacePurchaseCart(sellerCompanyId.value);
    },

    enabled: computed(() => Boolean(sellerCompanyId.value)),
  });
}
