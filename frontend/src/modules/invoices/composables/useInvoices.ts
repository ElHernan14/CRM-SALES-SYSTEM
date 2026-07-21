import { computed, type Ref } from 'vue';
import { useQuery } from '@tanstack/vue-query';

import { getCompanyInvoices } from '../api/invoices.api';
import type { GetCompanyInvoicesRequest } from '../types/invoice.types';
import { useTenantQueryScope } from '@/modules/auth/composables/useTenantQueryScope';

export function useInvoices(params: Ref<GetCompanyInvoicesRequest>) {
  const tenantScope = useTenantQueryScope();

  return useQuery({
    queryKey: computed(() => ['invoices', tenantScope.value.companyId, params.value]),

    queryFn: () => getCompanyInvoices(params.value),

    enabled: computed(() => Boolean(tenantScope.value.companyId)),
  });
}
