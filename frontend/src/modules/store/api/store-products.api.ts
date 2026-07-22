import { http } from '@/shared/api/http';
import { unwrapResponse } from '@/shared/utils/response';

import {
  GetStoreProductsResponseSchema,
  StoreProductSchema,
  type GetStoreProductsRequest,
  type GetStoreProductsResponse,
  type StoreProduct,
} from '../types/store-product.types';

export async function getStoreProducts(
  params: GetStoreProductsRequest
): Promise<GetStoreProductsResponse> {
  const response = await http.get('/store/products', {
    params,
  });

  return unwrapResponse(GetStoreProductsResponseSchema, response.data);
}

export async function getStoreProduct(productId: number): Promise<StoreProduct> {
  const response = await http.get(`/store/products/${productId}`);

  return unwrapResponse(StoreProductSchema, response.data);
}
