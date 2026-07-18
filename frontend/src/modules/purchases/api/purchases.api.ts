import { http } from '@/shared/api/http';
import { unwrapResponse } from '@/shared/utils/response';

import {
  GetPurchasesResponseSchema,
  type GetPurchasesRequest,
  type GetPurchasesResponse,
} from '../types/purchase.types';

export async function getCompanyPurchases(
  params: GetPurchasesRequest
): Promise<GetPurchasesResponse> {
  const response = await http.get('/company/purchases', { params });

  return unwrapResponse(GetPurchasesResponseSchema, response.data);
}
