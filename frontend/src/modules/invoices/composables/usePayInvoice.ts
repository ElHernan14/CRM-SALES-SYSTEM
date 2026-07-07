import { useMutation, useQueryClient } from '@tanstack/vue-query';

import { payInvoice } from '../api/invoices.api';
import type { PayInvoiceRequest } from '../types/invoice.types';

export function usePayInvoice() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, payload }: { id: number; payload: PayInvoiceRequest }) =>
      payInvoice(id, payload),

    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({
        queryKey: ['invoices'],
      });

      queryClient.invalidateQueries({
        queryKey: ['invoice', variables.id],
      });
    },
  });
}
