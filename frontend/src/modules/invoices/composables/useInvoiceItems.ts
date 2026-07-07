import { computed, type Ref } from 'vue';
import { useQuery } from '@tanstack/vue-query';

import { getInvoiceItems } from '../api/invoice-items.api';
import type { GetInvoiceItemsRequest } from '../types/invoice-item.types';

export function useInvoiceItems(
  invoiceId: Ref<number | null>,
  params: Ref<GetInvoiceItemsRequest>
) {
  return useQuery({
    queryKey: computed(() => ['invoice-items', invoiceId.value, params.value]),

    queryFn: () => getInvoiceItems(invoiceId.value as number, params.value),

    enabled: computed(() => !!invoiceId.value),
  });
}
