import { useMutation, useQueryClient } from '@tanstack/vue-query';

import { checkoutPurchaseCart } from '@/modules/marketplace/api/purchase-cart.api';

export function useCheckoutPurchase() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (invoiceId: number) =>
      checkoutPurchaseCart({
        invoice_id: invoiceId,
      }),

    onSuccess: async (_, invoiceId) => {
      await Promise.all([
        queryClient.invalidateQueries({
          queryKey: ['invoice', invoiceId],
        }),

        queryClient.invalidateQueries({
          queryKey: ['purchases'],
        }),

        queryClient.invalidateQueries({
          queryKey: ['purchase-items', invoiceId],
        }),

        queryClient.invalidateQueries({
          queryKey: ['store-products'],
        }),

        queryClient.invalidateQueries({
          queryKey: ['purchase-cart'],
        }),
      ]);
    },
  });
}
