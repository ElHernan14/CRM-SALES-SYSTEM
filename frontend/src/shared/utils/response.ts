import { z } from "zod"
import { ApiResponseSchema } from "../types/api.types"

export function unwrapResponse<T extends z.ZodTypeAny>(
  schema: T,
  response: unknown
): z.infer<T> {
  // Validamos la respuesta completa con ApiResponseSchema
  const parsed = ApiResponseSchema(schema).parse(response) as { data: z.infer<T> }
  return parsed.data
}
