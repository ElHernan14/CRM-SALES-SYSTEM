import { z } from 'zod';
import { ApiResponseSchema } from '../types/api.types';
import type { ApiResponse } from '../types/api.types';

export function unwrapResponse<T extends z.ZodTypeAny>(schema: T, response: unknown): z.output<T> {
  const parsed = ApiResponseSchema(schema).parse(response) as ApiResponse<T>;
  return parsed.data;
}
