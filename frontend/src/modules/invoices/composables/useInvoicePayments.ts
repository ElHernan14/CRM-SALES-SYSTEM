import { computed, type Ref } from 'vue';
import { useQuery } from '@tanstack/vue-query';

import { getInvoicePayments } from '../api/invoice-payments.api';
import type { GetInvoicePaymentsRequest } from '../types/invoice-payment.types';

export function useInvoicePayments(
  invoiceId: Ref<number | null>,
  params: Ref<GetInvoicePaymentsRequest>
) {
  return useQuery({
    queryKey: computed(() => ['invoice-payments', invoiceId.value, params.value]),

    queryFn: () => getInvoicePayments(invoiceId.value as number, params.value),

    enabled: computed(() => !!invoiceId.value),
  });
}
