import { computed, type Ref } from 'vue';
import { useQuery } from '@tanstack/vue-query';

import { getInvoiceById } from '../api/invoices.api';

export function useInvoiceDetail(invoiceId: Ref<number | null>) {
  return useQuery({
    queryKey: computed(() => ['invoice', invoiceId.value]),

    queryFn: () => getInvoiceById(invoiceId.value as number),

    enabled: computed(() => !!invoiceId.value),
  });
}
