import { computed, type Ref } from 'vue';

import { useQuery } from '@tanstack/vue-query';

import { getInvoicePayments } from '@/modules/invoices/api/invoice-payments.api';

export function usePurchasePayments(invoiceId: Ref<number | null>, enabled: Ref<boolean>) {
  return useQuery({
    queryKey: computed(() => ['purchase-payments', invoiceId.value]),

    queryFn: () =>
      getInvoicePayments(invoiceId.value as number, {
        page: 1,
        limit: 100,
        sort_column: 'created_at',
        order: 'desc',
      }),

    enabled: computed(() => {
      return enabled.value && !!invoiceId.value;
    }),

    staleTime: 30_000,
  });
}
