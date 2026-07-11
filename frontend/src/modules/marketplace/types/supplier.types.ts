import { z } from 'zod';
import { PaginationMetaSchema } from '@/shared/types/pagination.types';

export const SupplierSchema = z.object({
  id: z.number(),
  name: z.string(),

  email: z.string().email().nullable().optional(),

  logo: z.string().nullable().optional(),

  cover_image: z.string().nullable().optional(),

  description: z.string().nullable().optional(),

  category_id: z.number(),
  category: z.string(),

  total_products: z.number(),
});

export type Supplier = z.infer<typeof SupplierSchema>;

export const GetSuppliersRequestSchema = z.object({
  search: z.string().optional(),

  category_id: z.number().optional(),
  category: z.string().optional(),

  page: z.number(),
  limit: z.number(),

  sort_column: z.enum(['name', 'total_products', 'created_at']).optional(),

  order: z.enum(['asc', 'desc']).optional(),
});

export type GetSuppliersRequest = z.infer<typeof GetSuppliersRequestSchema>;

export const GetSuppliersResponseSchema = z.object({
  items: z.array(SupplierSchema),
  meta: PaginationMetaSchema,
});

export type GetSuppliersResponse = z.infer<typeof GetSuppliersResponseSchema>;
