import { computed, type Ref } from 'vue';
import { useQuery } from '@tanstack/vue-query';

import { getSuppliers } from '../api/marketplace.api';
import type { GetSuppliersRequest } from '../types/supplier.types';

export function useSuppliers(params: Ref<GetSuppliersRequest>) {
  return useQuery({
    queryKey: computed(() => ['marketplace-suppliers', params.value]),

    queryFn: () => getSuppliers(params.value),
  });
}
