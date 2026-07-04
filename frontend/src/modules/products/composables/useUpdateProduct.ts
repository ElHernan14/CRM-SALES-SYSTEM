import { useMutation, useQueryClient } from '@tanstack/vue-query';

import { updateProduct } from '../api/products.api';
import type { UpdateProductRequest } from '../types/product.types';

export function useUpdateProduct() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, payload }: { id: number; payload: UpdateProductRequest }) =>
      updateProduct(id, payload),

    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ['products'],
      });
    },
  });
}
