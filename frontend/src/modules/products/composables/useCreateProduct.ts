import { useMutation, useQueryClient } from '@tanstack/vue-query';

import { createProduct } from '../api/products.api';
import type { CreateProductRequest } from '../types/product.types';

export function useCreateProduct() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: CreateProductRequest) => createProduct(payload),

    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ['products'],
      });

      queryClient.invalidateQueries({
        queryKey: ['store-products'],
      });
    },
  });
}
