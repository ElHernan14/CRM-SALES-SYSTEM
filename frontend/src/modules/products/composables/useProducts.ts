import { computed, type Ref } from 'vue';
import { useQuery } from '@tanstack/vue-query';

import { getProducts } from '../api/products.api';
import type { GetProductsRequest } from '../types/product.types';

export function useProducts(params: Ref<GetProductsRequest>) {
  return useQuery({
    queryKey: computed(() => ['products', params.value]),
    queryFn: () => getProducts(params.value),
    enabled: computed(() => !!params.value.company_id),
  });
}
