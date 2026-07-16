import { http } from '@/shared/api/http';
import { unwrapResponse } from '@/shared/utils/response';

import {
  CheckoutResponseSchema,
  EnsureCartResponseSchema,
  NullableCartSchema,
  type Cart,
  type CheckoutRequest,
  type CheckoutResponse,
  type EnsureCartRequest,
  type EnsureCartResponse,
} from '../types/purchase-cart.types';

export async function getPurchaseCart(sellerCompanyId: number): Promise<Cart | null> {
  const response = await http.get('/store/cart', {
    params: {
      seller_company_id: sellerCompanyId,
    },
  });

  return unwrapResponse(NullableCartSchema, response.data);
}

export async function ensurePurchaseCart(payload: EnsureCartRequest): Promise<EnsureCartResponse> {
  const response = await http.post('/store/cart/ensure', payload);

  return unwrapResponse(EnsureCartResponseSchema, response.data);
}

export async function checkoutPurchaseCart(payload: CheckoutRequest): Promise<CheckoutResponse> {
  const response = await http.post('/store/checkout', payload);

  return unwrapResponse(CheckoutResponseSchema, response.data);
}
