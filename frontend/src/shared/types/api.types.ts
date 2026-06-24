import { z } from "zod"

export const ApiResponseSchema = <T extends z.ZodTypeAny>(dataSchema: T) =>
  z.object({
    status: z.string(),
    code: z.number(),
    data: dataSchema,
  })

export type ApiResponse<T> = {
  status: string
  code: number
  data: z.infer<T>
}
