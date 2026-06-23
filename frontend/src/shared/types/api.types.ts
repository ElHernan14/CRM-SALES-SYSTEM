import { z } from "zod"

// Esquema genérico de respuesta
export const ApiResponseSchema = <T extends z.ZodTypeAny>(dataSchema: T) =>
  z.object({
    status: z.string(),
    code: z.number(),
    data: dataSchema,
  })

// Tipo genérico inferido
export type ApiResponse<T> = {
  status: string
  code: number
  data: T
}
