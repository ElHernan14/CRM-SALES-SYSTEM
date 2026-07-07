import { useMutation, useQueryClient } from '@tanstack/vue-query';

import { submitInvoice } from '../api/invoices.api';

export function useSubmitInvoice() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: number) => submitInvoice(id),

    onSuccess: (_, id) => {
      queryClient.invalidateQueries({
        queryKey: ['invoices'],
      });

      queryClient.invalidateQueries({
        queryKey: ['invoice', id],
      });
    },
  });
}
