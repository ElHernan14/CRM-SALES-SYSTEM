import { useMutation, useQueryClient } from '@tanstack/vue-query';

import { createInvoiceItem } from '../api/invoice-items.api';
import type { CreateInvoiceItemRequest } from '../types/invoice-item.types';

export function useCreateInvoiceItem() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({
      invoiceId,
      payload,
    }: {
      invoiceId: number;
      payload: CreateInvoiceItemRequest;
    }) => createInvoiceItem(invoiceId, payload),

    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({
        queryKey: ['invoice-items', variables.invoiceId],
      });

      queryClient.invalidateQueries({
        queryKey: ['invoice', variables.invoiceId],
      });

      queryClient.invalidateQueries({
        queryKey: ['invoices'],
      });
    },
  });
}
