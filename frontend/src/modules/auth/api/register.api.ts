import { http } from '@/shared/api/http';
import { unwrapResponse } from '@/shared/utils/response';

import {
  RegisterResponseSchema,
  type BusinessRegisterRequest,
  type PersonalRegisterRequest,
  type RegisterResponse,
} from '../types/register.types';

export async function registerPersonalApi(
  payload: PersonalRegisterRequest
): Promise<RegisterResponse> {
  const response = await http.post('/auth/register/personal', payload);

  return unwrapResponse(RegisterResponseSchema, response.data);
}

export async function registerBusinessApi(
  payload: BusinessRegisterRequest
): Promise<RegisterResponse> {
  const response = await http.post('/auth/register/business', payload);

  return unwrapResponse(RegisterResponseSchema, response.data);
}
