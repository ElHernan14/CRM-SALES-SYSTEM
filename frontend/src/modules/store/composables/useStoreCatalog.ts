import { computed, type Ref } from 'vue';

import { useQuery } from '@tanstack/vue-query';

import { getStoreProducts } from '../api/store-products.api';

import type { GetStoreProductsRequest } from '../types/store-product.types';

export function useStoreCatalog(params: Ref<GetStoreProductsRequest>) {
  return useQuery({
    queryKey: computed(() => ['b2c-store-products', params.value]),

    queryFn: () => getStoreProducts(params.value),

    staleTime: 30_000,
  });
}
