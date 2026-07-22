import { computed, type Ref } from 'vue';

import { useQuery } from '@tanstack/vue-query';

import { getStoreProduct } from '../api/store-products.api';

export function useStoreProduct(productId: Ref<number | null>) {
  return useQuery({
    queryKey: computed(() => ['b2c-store-product', productId.value]),

    queryFn: () => getStoreProduct(productId.value as number),

    enabled: computed(() => Boolean(productId.value)),

    staleTime: 30_000,
  });
}
