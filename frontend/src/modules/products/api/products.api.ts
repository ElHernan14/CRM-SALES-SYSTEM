import { http } from '@/shared/api/http';
import { unwrapResponse } from '@/shared/utils/response';

import {
  GetProductsResponseSchema,
  type GetProductsRequest,
  type GetProductsResponse,
} from '../types/product.types';

export async function getProducts(params: GetProductsRequest): Promise<GetProductsResponse> {
  const response = await http.get('/product', {
    params,
  });

  return unwrapResponse(GetProductsResponseSchema, response.data);
}
