import { useMutation, useQueryClient } from '@tanstack/vue-query';

import { ensurePurchaseCart } from '../api/purchase-cart.api';

export function useEnsurePurchaseCart() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (sellerCompanyId: number) =>
      ensurePurchaseCart({
        seller_company_id: sellerCompanyId,
      }),

    onSuccess: (_, sellerCompanyId) => {
      queryClient.invalidateQueries({
        queryKey: ['purchase-cart', sellerCompanyId],
      });
    },
  });
}
