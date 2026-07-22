import { useMutation, useQueryClient } from '@tanstack/vue-query';

import { checkoutAllStoreCarts } from '../api/store-cart.api';

import type { CheckoutAllRequest } from '../types/store-cart.types';

export function useCheckoutAllStoreCarts() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload?: CheckoutAllRequest) => checkoutAllStoreCarts(payload),

    onSuccess: async () => {
      await Promise.all([
        queryClient.invalidateQueries({
          queryKey: ['b2c-store-carts'],
        }),

        queryClient.invalidateQueries({
          queryKey: ['b2c-store-products'],
        }),

        queryClient.invalidateQueries({
          queryKey: ['b2c-store-product'],
        }),

        queryClient.invalidateQueries({
          queryKey: ['purchases'],
        }),

        queryClient.invalidateQueries({
          queryKey: ['store-purchases'],
        }),
      ]);
    },
  });
}
