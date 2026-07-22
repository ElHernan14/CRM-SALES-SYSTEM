import { http } from '@/shared/api/http';

import { unwrapResponse } from '@/shared/utils/response';

import {
  CheckoutAllResponseSchema,
  GetStoreCartsResponseSchema,
  type CheckoutAllRequest,
  type CheckoutAllResponse,
  type GetStoreCartsResponse,
} from '../types/store-cart.types';

export async function getStoreCarts(): Promise<GetStoreCartsResponse> {
  const response = await http.get('/store/carts');

  return unwrapResponse(GetStoreCartsResponseSchema, response.data);
}

export async function checkoutAllStoreCarts(
  payload?: CheckoutAllRequest
): Promise<CheckoutAllResponse> {
  const response = await http.post('/store/checkout/all', payload ?? {});

  return unwrapResponse(CheckoutAllResponseSchema, response.data);
}
