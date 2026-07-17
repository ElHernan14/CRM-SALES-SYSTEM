import { useMutation } from '@tanstack/vue-query';

import { deleteInvoiceItem } from '@/modules/invoices/api/invoice-items.api';

export function useDeletePurchaseCartItem() {
  return useMutation({
    mutationFn: ({
      invoiceId,
      itemId,
    }: {
      invoiceId: number;
      sellerCompanyId: number;
      itemId: number;
    }) => deleteInvoiceItem(invoiceId, itemId),
  });
}
