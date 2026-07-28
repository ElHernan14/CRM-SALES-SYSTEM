import { useMutation, useQueryClient } from '@tanstack/vue-query';

import { checkoutPurchaseCart } from '../api/purchase-cart.api';

export function useCheckoutPurchaseCart() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (variables: { invoiceId: number; sellerCompanyId: number }) =>
      checkoutPurchaseCart({
        invoice_id: variables.invoiceId,
      }),

    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({
        queryKey: ['purchase-cart', variables.sellerCompanyId],
      });

      queryClient.invalidateQueries({
        queryKey: ['purchases'],
      });

      queryClient.invalidateQueries({
        queryKey: ['store-products'],
      });

      queryClient.invalidateQueries({
        queryKey: ['products'],
      });
    },
  });
}
