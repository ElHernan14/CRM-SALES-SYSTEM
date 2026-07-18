import { computed, type Ref } from 'vue';

import { useQuery } from '@tanstack/vue-query';

import { getInvoiceItems } from '@/modules/invoices/api/invoice-items.api';

export function usePurchaseItems(invoiceId: Ref<number | null>, enabled: Ref<boolean>) {
  return useQuery({
    queryKey: computed(() => ['purchase-items', invoiceId.value]),

    queryFn: () =>
      getInvoiceItems(invoiceId.value as number, {
        page: 1,
        limit: 100,
      }),

    enabled: computed(() => {
      return enabled.value && !!invoiceId.value;
    }),

    staleTime: 30_000,
  });
}
