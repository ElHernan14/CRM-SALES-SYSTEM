import { computed, type Ref } from 'vue';
import { useQuery } from '@tanstack/vue-query';

import { getStoreProducts } from '../api/marketplace.api';
import type { GetStoreProductsRequest } from '../types/store-product.types';

export function useStoreProducts(params: Ref<GetStoreProductsRequest>) {
  return useQuery({
    queryKey: computed(() => ['store-products', params.value]),

    queryFn: () => getStoreProducts(params.value),
  });
}
