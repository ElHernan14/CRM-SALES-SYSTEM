import { useMutation, useQueryClient } from '@tanstack/vue-query';

import { deleteProduct } from '../api/products.api';

export function useDeleteProduct() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: number) => deleteProduct(id),

    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ['products'],
      });
    },
  });
}
