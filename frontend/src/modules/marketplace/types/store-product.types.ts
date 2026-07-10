import { z } from 'zod';
import { PaginationMetaSchema } from '@/shared/types/pagination.types';

export const StoreProductTypeSchema = z.enum(['product', 'service']);

export const StoreProductSchema = z.object({
  id: z.number(),
  company_id: z.number(),
  company_name: z.string().optional(),
  name: z.string(),
  type: StoreProductTypeSchema,
  description: z.string(),
  price: z.number(),
  available_stock: z.number(),
});

export type StoreProduct = z.infer<typeof StoreProductSchema>;

export const GetStoreProductsRequestSchema = z.object({
  page: z.number(),
  limit: z.number(),
  company_id: z.number().optional(),
  search: z.string().optional(),
  name: z.string().optional(),
  type: StoreProductTypeSchema.optional(),
  min_price: z.number().optional(),
  max_price: z.number().optional(),
  sort_column: z.enum(['name', 'created_at', 'type', 'price', 'stock']).optional(),
  order: z.enum(['asc', 'desc']).optional(),
});

export type GetStoreProductsRequest = z.infer<typeof GetStoreProductsRequestSchema>;

export const GetStoreProductsResponseSchema = z.object({
  items: z.array(StoreProductSchema),
  meta: PaginationMetaSchema,
});

export type GetStoreProductsResponse = z.infer<typeof GetStoreProductsResponseSchema>;
