import { http } from '@/shared/api/http';
import { unwrapResponse } from '@/shared/utils/response';

import { GetCategoriesResponseSchema, type GetCategoriesResponse } from '../types/category.types';

export async function getCategories(): Promise<GetCategoriesResponse> {
  const response = await http.get('/categories');

  return unwrapResponse(GetCategoriesResponseSchema, response.data);
}
