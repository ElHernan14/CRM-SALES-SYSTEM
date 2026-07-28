import { http } from '@/shared/api/http';

import {
  MarketplaceCheckoutResultSchema,
  MarketplacePurchaseCartSchema,
  type MarketplaceCheckoutResult,
  type MarketplacePurchaseCart,
} from '../types/purchase-cart.types';

type ApiEnvelope<T> = {
  status: string;
  code: number;
  data: T;
};

export async function getMarketplacePurchaseCart(
  sellerCompanyId: number
): Promise<MarketplacePurchaseCart | null> {
  const response = await http.get<ApiEnvelope<unknown>>('/marketplace/cart', {
    params: {
      seller_company_id: sellerCompanyId,
    },
  });

  /*
   * El backend devuelve data=null cuando no existe un draft
   * o cuando el carrito no contiene items.
   */
  if (response.data.data === null) {
    return null;
  }

  return MarketplacePurchaseCartSchema.parse(response.data.data);
}

export async function checkoutMarketplacePurchaseCart(
  invoiceId: number
): Promise<MarketplaceCheckoutResult> {
  const response = await http.post<ApiEnvelope<unknown>>('/marketplace/checkout', {
    invoice_id: invoiceId,
  });

  return MarketplaceCheckoutResultSchema.parse(response.data.data);
}
