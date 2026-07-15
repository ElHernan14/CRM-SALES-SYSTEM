import { useMutation, useQueryClient } from '@tanstack/vue-query';

import { uploadProductImage } from '../api/products.api';

export function useUploadProductImage() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ productId, image }: { productId: number; image: File }) =>
      uploadProductImage(productId, image),

    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({
        queryKey: ['products'],
      });

      queryClient.invalidateQueries({
        queryKey: ['store-products'],
      });

      queryClient.invalidateQueries({
        queryKey: ['product', variables.productId],
      });
    },
  });
}
