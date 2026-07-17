import { useMutation } from '@tanstack/vue-query';

import { updateInvoiceItem } from '@/modules/invoices/api/invoice-items.api';

export function useUpdatePurchaseCartItem() {
  return useMutation({
    mutationFn: ({
      invoiceId,
      itemId,
      quantity,
    }: {
      invoiceId: number;
      sellerCompanyId: number;
      itemId: number;
      quantity: number;
    }) => updateInvoiceItem(invoiceId, itemId, { quantity }),
  });
}
