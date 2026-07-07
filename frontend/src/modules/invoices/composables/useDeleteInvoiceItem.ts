import { useMutation, useQueryClient } from '@tanstack/vue-query';

import { deleteInvoiceItem } from '../api/invoice-items.api';

export function useDeleteInvoiceItem() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ invoiceId, itemId }: { invoiceId: number; itemId: number }) =>
      deleteInvoiceItem(invoiceId, itemId),

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
