import { useMutation } from '@tanstack/vue-query';

import { ensurePurchaseCart } from '@/modules/marketplace/api/purchase-cart.api';

export function useEnsureStoreCart() {
  return useMutation({
    mutationFn: (sellerCompanyId: number) =>
      ensurePurchaseCart({
        seller_company_id: sellerCompanyId,
      }),
  });
}
