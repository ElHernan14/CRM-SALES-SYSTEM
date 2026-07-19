import { z } from 'zod';
import { PaginationMetaSchema } from '@/shared/types/pagination.types';

// Product Types
export const ProductTypeSchema = z.enum(['product', 'service']);

export const ProductKindSchema = z.enum(['product', 'service']);

export type ProductKind = z.infer<typeof ProductKindSchema>;

// Product list item
export const ProductListItemSchema = z.object({
  id: z.number(),
  company_id: z.number(),

  name: z.string(),
  description: z.string(),

  kind: ProductKindSchema,

  category_id: z.number(),
  category: z.string(),

  type_id: z.number(),
  type: z.string(),

  price: z.number(),

  stock: z.number(),
  available_stock: z.number().optional(),

  image_path: z.string().nullable().optional(),

  status: z.number(),
});

export type ProductListItem = z.infer<typeof ProductListItemSchema>;

export const GetProductsRequestSchema = z.object({
  search: z.string().optional(),

  kind: ProductKindSchema.optional(),

  category_id: z.number().optional(),
  category: z.string().optional(),

  type_id: z.number().optional(),
  type: z.string().optional(),

  min_price: z.number().optional(),
  max_price: z.number().optional(),

  status: z.union([z.literal(0), z.literal(1)]).optional(),

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

  kind: ProductKindSchema.optional(),

  category_id: z.number().int().positive().optional(),

  type_id: z.number().int().positive().optional(),

  price: z.number().min(0).optional(),

  stock: z.number().int().min(0).optional(),

  status: z.union([z.literal(0), z.literal(1)]).optional(),
});

export type UpdateProductRequest = z.infer<typeof UpdateProductRequestSchema>;

export const ProductDetailResponseSchema = z.object({
  id: z.number(),
  company_id: z.number().optional(),

  name: z.string(),
  description: z.string(),

  kind: ProductKindSchema,

  category_id: z.number(),
  category: z.string(),

  type_id: z.number(),
  type: z.string(),

  price: z.number(),
  stock: z.number().optional(),
  available_stock: z.number().optional(),

  image_path: z.string().nullable().optional(),

  status: z.number().optional(),
});

export type ProductDetailResponse = z.infer<typeof ProductDetailResponseSchema>;

export const CreateProductRequestSchema = z.object({
  name: z.string().min(2).max(100),

  description: z.string().max(500).optional(),

  kind: ProductKindSchema,

  category_id: z.number().int().positive(),

  type_id: z.number().int().positive(),

  price: z.number().min(0),

  stock: z.number().int().min(0),
});

export type CreateProductRequest = z.infer<typeof CreateProductRequestSchema>;

export const UploadProductImageResponseSchema = z.object({
  image_path: z.string(),
});

export type UploadProductImageResponse = z.infer<typeof UploadProductImageResponseSchema>;

export const BulkDeleteProductsRequestSchema = z.object({
  product_ids: z.array(z.number().int().positive()).min(1),
});

export type BulkDeleteProductsRequest = z.infer<typeof BulkDeleteProductsRequestSchema>;

export const BulkDeleteProductsResponseSchema = z.object({
  deleted_ids: z.array(z.number()),
  count: z.number(),
});

export type BulkDeleteProductsResponse = z.infer<typeof BulkDeleteProductsResponseSchema>;
