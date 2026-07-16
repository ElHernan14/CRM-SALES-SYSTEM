import { useMutation, useQueryClient } from '@tanstack/vue-query';

import { createInvoiceItem } from '@/modules/invoices/api/invoice-items.api';

export function useAddPurchaseCartItem() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({
      invoiceId,
      sellerCompanyId,
      productId,
      quantity,
    }: {
      invoiceId: number;
      sellerCompanyId: number;
      productId: number;
      quantity: number;
    }) =>
      createInvoiceItem(invoiceId, {
        product_id: productId,
        quantity,
      }),

    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({
        queryKey: ['purchase-cart', variables.sellerCompanyId],
      });

      queryClient.invalidateQueries({
        queryKey: ['store-products'],
      });
    },
  });
}
