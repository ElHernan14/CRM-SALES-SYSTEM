import { http } from '@/shared/api/http';
import { unwrapResponse } from '@/shared/utils/response';

import {
  GetProductTypesResponseSchema,
  type GetProductTypesRequest,
  type GetProductTypesResponse,
} from '../types/product-type.types';

export async function getProductTypes(
  params: GetProductTypesRequest = {}
): Promise<GetProductTypesResponse> {
  const response = await http.get('/product-types', {
    params,
  });

  return unwrapResponse(GetProductTypesResponseSchema, response.data);
}
