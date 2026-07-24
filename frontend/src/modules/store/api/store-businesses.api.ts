import { http } from '@/shared/api/http';

import { unwrapResponse } from '@/shared/utils/response';

import {
  StoreBusinessesResponseSchema,
  type StoreBusinessesParams,
  type StoreBusinessesResponse,
} from '../types/store-business.types';

export async function getStoreBusinesses(
  params: StoreBusinessesParams
): Promise<StoreBusinessesResponse> {
  const response = await http.get('/marketplace/suppliers', {
    params,
  });

  return unwrapResponse(StoreBusinessesResponseSchema, response.data);
}
