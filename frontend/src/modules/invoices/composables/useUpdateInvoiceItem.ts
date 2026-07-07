import { useMutation, useQueryClient } from '@tanstack/vue-query';

import { updateInvoiceItem } from '../api/invoice-items.api';
import type { UpdateInvoiceItemRequest } from '../types/invoice-item.types';

export function useUpdateInvoiceItem() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({
      invoiceId,
      itemId,
      payload,
    }: {
      invoiceId: number;
      itemId: number;
      payload: UpdateInvoiceItemRequest;
    }) => updateInvoiceItem(invoiceId, itemId, payload),

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
