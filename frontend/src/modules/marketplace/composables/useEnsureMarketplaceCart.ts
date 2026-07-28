import { useMutation } from '@tanstack/vue-query';

import { ensureMarketplaceCart } from '../api/marketplace-cart.api';

export function useEnsureMarketplaceCart() {
  return useMutation({
    mutationFn: (sellerCompanyId: number) => {
      return ensureMarketplaceCart(sellerCompanyId);
    },
  });
}
