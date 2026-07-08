import { http } from '@/shared/api/http';
import { unwrapResponse } from '@/shared/utils/response';

import {
  GetClientsResponseSchema,
  type GetClientsRequest,
  type GetClientsResponse,
} from '../types/client.types';

export async function getClients(params: GetClientsRequest): Promise<GetClientsResponse> {
  const response = await http.get('/clients/clients', {
    params,
  });

  return unwrapResponse(GetClientsResponseSchema, response.data);
}
