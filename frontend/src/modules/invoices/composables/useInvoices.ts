import { computed, type Ref } from 'vue';
import { useQuery } from '@tanstack/vue-query';

import { getCompanyInvoices } from '../api/invoices.api';
import type { GetCompanyInvoicesRequest } from '../types/invoice.types';

export function useInvoices(params: Ref<GetCompanyInvoicesRequest>) {
  return useQuery({
    queryKey: computed(() => ['invoices', params.value]),
    queryFn: () => getCompanyInvoices(params.value),
  });
}
