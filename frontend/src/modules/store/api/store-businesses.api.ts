import { http } from '@/shared/api/http';

import { unwrapResponse } from '@/shared/utils/response';

import {
  StoreBusinessSchema,
  StoreBusinessesResponseSchema,
  type StoreBusiness,
  type StoreBusinessesParams,
  type StoreBusinessesResponse,
} from '../types/store-business.types';

export async function getStoreBusinesses(
  params: StoreBusinessesParams
): Promise<StoreBusinessesResponse> {
  const response = await http.get('/marketplace/suppliers', {
    params: {
      ...params,
      include_self: true,
    },
  });

  return unwrapResponse(StoreBusinessesResponseSchema, response.data);
}

export async function getStoreBusiness(businessId: number): Promise<StoreBusiness> {
  const response = await http.get(`/marketplace/suppliers/${businessId}`, {
    params: {
      include_self: true,
    },
  });

  return unwrapResponse(StoreBusinessSchema, response.data);
}
