import { useMutation, useQueryClient } from '@tanstack/vue-query';

import { checkoutMarketplacePurchaseCart } from '../api/marketplace-purchase-cart.api';

export function useCheckoutMarketplacePurchaseCart() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (invoiceId: number) => {
      return checkoutMarketplacePurchaseCart(invoiceId);
    },

    onSuccess: async () => {
      await Promise.all([
        queryClient.invalidateQueries({
          queryKey: ['marketplace-purchase-cart'],
        }),

        queryClient.invalidateQueries({
          queryKey: ['purchases'],
        }),

        queryClient.invalidateQueries({
          queryKey: ['purchase'],
        }),

        queryClient.invalidateQueries({
          queryKey: ['dashboard-overview'],
        }),
      ]);
    },
  });
}
