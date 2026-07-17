import { useMutation } from '@tanstack/vue-query';

import { ensurePurchaseCart } from '../api/purchase-cart.api';

export function useEnsurePurchaseCart() {
  return useMutation({
    mutationFn: (sellerCompanyId: number) =>
      ensurePurchaseCart({
        seller_company_id: sellerCompanyId,
      }),
  });
}
