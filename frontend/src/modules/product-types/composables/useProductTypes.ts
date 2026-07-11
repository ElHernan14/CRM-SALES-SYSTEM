import { computed, type Ref } from 'vue';
import { useQuery } from '@tanstack/vue-query';

import { getProductTypes } from '../api/product-types.api';

export function useProductTypes(categoryId?: Ref<number | null | undefined>) {
  return useQuery({
    queryKey: computed(() => ['product-types', categoryId?.value ?? 'all']),

    queryFn: () =>
      getProductTypes({
        category_id: categoryId?.value ?? undefined,
      }),

    staleTime: 1000 * 60 * 30,

    enabled: computed(() => {
      if (!categoryId) return true;
      return !!categoryId.value;
    }),
  });
}
