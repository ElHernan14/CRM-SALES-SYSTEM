import { http } from '@/shared/api/http';
import { unwrapResponse } from '@/shared/utils/response';

import {
  ClientProfileSchema,
  type ClientProfile,
  type UpdateClientProfileRequest,
} from '../types/account.types';

export async function getClientProfile(clientId: number): Promise<ClientProfile> {
  const response = await http.get(`/clients/clients/${clientId}`);

  return unwrapResponse(ClientProfileSchema, response.data);
}

export async function updateClientProfile(
  clientId: number,
  payload: UpdateClientProfileRequest
): Promise<ClientProfile> {
  const response = await http.patch(`/clients/clients/${clientId}`, payload);

  return unwrapResponse(ClientProfileSchema, response.data);
}
