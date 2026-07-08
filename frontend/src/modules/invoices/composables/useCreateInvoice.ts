import { useMutation, useQueryClient } from '@tanstack/vue-query';

import { createInvoice } from '../api/invoices.api';
import type { CreateInvoiceRequest } from '../types/invoice.types';

export function useCreateInvoice() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: CreateInvoiceRequest) => createInvoice(payload),

    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ['invoices'],
      });
    },
  });
}
