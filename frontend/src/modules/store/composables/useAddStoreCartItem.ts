import { useMutation } from '@tanstack/vue-query';

import { createInvoiceItem } from '@/modules/invoices/api/invoice-items.api';

export function useAddStoreCartItem() {
  return useMutation({
    mutationFn: ({
      invoiceId,
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
  });
}
