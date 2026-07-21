import { computed, type Ref } from 'vue';

import { useQuery } from '@tanstack/vue-query';

import { getCompanyPurchases } from '../api/purchases.api';
import type { GetPurchasesRequest } from '../types/purchase.types';
import { useTenantQueryScope } from '@/modules/auth/composables/useTenantQueryScope';

export function usePurchases(params: Ref<GetPurchasesRequest>) {
  const tenantScope = useTenantQueryScope();

  return useQuery({
    queryKey: computed(() => [
      'purchases',
      tenantScope.value.companyId,
      tenantScope.value.clientId,
      params.value,
    ]),

    queryFn: () => getCompanyPurchases(params.value),

    enabled: computed(() => Boolean(tenantScope.value.companyId)),
  });
}
