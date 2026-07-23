import { useMutation, useQueryClient } from '@tanstack/vue-query';

import { payPurchase } from '../api/purchase-payments.api';

import type { PayPurchaseRequest } from '../types/purchase-payment.types';

export function usePayPurchase() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ invoiceId, payload }: { invoiceId: number; payload: PayPurchaseRequest }) =>
      payPurchase(invoiceId, payload),

    onSuccess: async (_, variables) => {
      await Promise.all([
        queryClient.invalidateQueries({
          queryKey: ['invoice', variables.invoiceId],
        }),

        queryClient.invalidateQueries({
          queryKey: ['purchase-payments', variables.invoiceId],
        }),

        queryClient.invalidateQueries({
          queryKey: ['purchases'],
        }),

        queryClient.invalidateQueries({
          queryKey: ['store-products'],
        }),

        queryClient.invalidateQueries({
          queryKey: ['store-purchases'],
        }),
      ]);
    },
  });
}
