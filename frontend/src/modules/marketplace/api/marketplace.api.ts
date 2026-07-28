import { http } from '@/shared/api/http';
import { unwrapResponse } from '@/shared/utils/response';

import {
  GetSuppliersResponseSchema,
  type GetSuppliersRequest,
  type GetSuppliersResponse,
} from '../types/supplier.types';

import {
  GetStoreProductsResponseSchema,
  type GetStoreProductsRequest,
  type GetStoreProductsResponse,
} from '../types/store-product.types';

export async function getSuppliers(params: GetSuppliersRequest): Promise<GetSuppliersResponse> {
  const response = await http.get('/marketplace/suppliers', {
    params,
  });

  return unwrapResponse(GetSuppliersResponseSchema, response.data);
}

export async function getStoreProducts(
  params: GetStoreProductsRequest
): Promise<GetStoreProductsResponse> {
  const response = await http.get('/store/products', {
    params,
  });

  return unwrapResponse(GetStoreProductsResponseSchema, response.data);
}
