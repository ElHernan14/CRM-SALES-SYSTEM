import { z } from 'zod';
import { PaginationMetaSchema } from '@/shared/types/pagination.types';

export const ProductTypeSchema = z.enum(['product', 'service']);

export const ProductListItemSchema = z.object({
  id: z.number(),
  name: z.string(),
  description: z.string(),
  type: z.string(),
  price: z.number(),
  stock: z.number(),
  status: z.number(),
  company_id: z.number(),
});

export type ProductListItem = z.infer<typeof ProductListItemSchema>;

export const GetProductsRequestSchema = z.object({
  search: z.string().optional(),
  type: z.string().optional(),
  min_price: z.number().optional(),
  max_price: z.number().optional(),
  company_id: z.number().optional(),
  page: z.number().default(1),
  limit: z.number().default(10),
});

export type GetProductsRequest = z.infer<typeof GetProductsRequestSchema>;

export const GetProductsResponseSchema = z.object({
  data: z.array(ProductListItemSchema),
  meta: PaginationMetaSchema,
});

export type GetProductsResponse = z.infer<typeof GetProductsResponseSchema>;
