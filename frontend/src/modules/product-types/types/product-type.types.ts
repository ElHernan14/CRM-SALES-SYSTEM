import { z } from 'zod';

export const ProductTypeSchema = z.object({
  id: z.number(),
  category_id: z.number(),
  name: z.string(),
  description: z
    .string()
    .nullable()
    .optional()
    .transform((value) => value ?? ''),
});

export type ProductType = z.infer<typeof ProductTypeSchema>;

export const GetProductTypesRequestSchema = z.object({
  category_id: z.number().optional(),
});

export type GetProductTypesRequest = z.infer<typeof GetProductTypesRequestSchema>;

export const GetProductTypesResponseSchema = z.object({
  items: z.array(ProductTypeSchema),
});

export type GetProductTypesResponse = z.infer<typeof GetProductTypesResponseSchema>;
