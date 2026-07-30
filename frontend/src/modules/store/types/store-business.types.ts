import { z } from 'zod';

export const StoreBusinessSchema = z.object({
  id: z.number(),

  name: z.string(),

  category_id: z.number(),
  category: z.string(),

  description: z.string().nullable().optional(),

  logo: z.string().nullable().optional(),

  cover_image: z.string().nullable().optional(),

  total_products: z.number(),
});

export type StoreBusiness = z.infer<typeof StoreBusinessSchema>;

export const StoreBusinessesParamsSchema = z.object({
  search: z.string().optional(),

  category_id: z.number().optional(),

  page: z.number().int().positive().optional(),

  limit: z.number().int().positive().optional(),

  sort_column: z.enum(['name', 'created_at', 'total_products']).optional(),

  order: z.enum(['asc', 'desc']).optional(),

  include_self: z.boolean().optional(),
});

export type StoreBusinessesParams = z.infer<typeof StoreBusinessesParamsSchema>;

export const StoreBusinessesResponseSchema = z.object({
  items: z.array(StoreBusinessSchema),

  meta: z.object({
    page: z.number(),
    limit: z.number(),
    total: z.number(),
    total_pages: z.number(),
  }),
});

export type StoreBusinessesResponse = z.infer<typeof StoreBusinessesResponseSchema>;
