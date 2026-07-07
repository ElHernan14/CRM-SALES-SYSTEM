import { useMutation, useQueryClient } from '@tanstack/vue-query';

import { cancelInvoice } from '../api/invoices.api';

export function useCancelInvoice() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: number) => cancelInvoice(id),

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
