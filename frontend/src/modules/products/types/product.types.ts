import { z } from 'zod';
import { PaginationMetaSchema } from '@/shared/types/pagination.types';

// Product Types
export const ProductTypeSchema = z.enum(['product', 'service']);

// Product list item
export const ProductListItemSchema = z.object({
  id: z.number(),
  name: z.string(),
  description: z.string(),
  type: ProductTypeSchema,
  price: z.number(),
  stock: z.number(),
  status: z.number(),
  company_id: z.number(),
  available_stock: z.number().optional(),
});

export type ProductListItem = z.infer<typeof ProductListItemSchema>;

export const GetProductsRequestSchema = z.object({
  search: z.string().optional(),
  type: ProductTypeSchema.optional(),
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

// Update Product request
export const UpdateProductRequestSchema = z.object({
  name: z.string().min(2).max(100).optional(),
  description: z.string().max(500).optional(),
  type: ProductTypeSchema.optional(),
  price: z.number().min(0).optional(),
  stock: z.number().int().min(0).optional(),
  status: z.union([z.literal(0), z.literal(1)]).optional(),
});

export type UpdateProductRequest = z.infer<typeof UpdateProductRequestSchema>;

export const ProductDetailResponseSchema = z.object({
  id: z.number(),
  name: z.string(),
  description: z.string(),
  type: ProductTypeSchema,
  price: z.number(),
  stock: z.number(),
  status: z.number(),
  company_id: z.number(),
});

export type ProductDetailResponse = z.infer<typeof ProductDetailResponseSchema>;
