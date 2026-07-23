import { http } from '@/shared/api/http';

import { unwrapResponse } from '@/shared/utils/response';

import {
  GetStorePurchasesResponseSchema,
  type GetStorePurchasesRequest,
  type GetStorePurchasesResponse,
} from '../types/store-purchase.types';

export async function getStorePurchases(
  params: GetStorePurchasesRequest
): Promise<GetStorePurchasesResponse> {
  const response = await http.get('/store/purchases', {
    params,
  });

  return unwrapResponse(GetStorePurchasesResponseSchema, response.data);
}
