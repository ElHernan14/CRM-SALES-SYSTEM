import { useMutation, useQueryClient } from '@tanstack/vue-query';

import { bulkDeleteProducts } from '../api/products.api';

export function useBulkDeleteProducts() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (productIds: number[]) =>
      bulkDeleteProducts({
        product_ids: productIds,
      }),

    onSuccess: async () => {
      await Promise.all([
        queryClient.invalidateQueries({
          queryKey: ['products'],
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
