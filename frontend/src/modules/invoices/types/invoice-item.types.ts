import { z } from 'zod';
import { PaginationMetaSchema } from '@/shared/types/pagination.types';

export const InvoiceItemSchema = z.object({
  id: z.number(),
  invoice_id: z.number().nullable().optional(),
  product_id: z.number(),
  product_name: z.string(),
  quantity: z.number(),
  price: z.number(),
  subtotal: z.number(),
  product_image_path: z.string().nullable().optional(),
});

export type InvoiceItem = z.infer<typeof InvoiceItemSchema>;

export const GetInvoiceItemsRequestSchema = z.object({
  page: z.number(),
  limit: z.number(),
});

export type GetInvoiceItemsRequest = z.infer<typeof GetInvoiceItemsRequestSchema>;

export const GetInvoiceItemsResponseSchema = z.object({
  items: z.array(InvoiceItemSchema),
  meta: PaginationMetaSchema,
});

export type GetInvoiceItemsResponse = z.infer<typeof GetInvoiceItemsResponseSchema>;

export const CreateInvoiceItemRequestSchema = z.object({
  product_id: z.number(),
  quantity: z.number().min(1).max(999),
});

export type CreateInvoiceItemRequest = z.infer<typeof CreateInvoiceItemRequestSchema>;

export const UpdateInvoiceItemRequestSchema = z.object({
  quantity: z.number().min(1).max(9999),
});

export type UpdateInvoiceItemRequest = z.infer<typeof UpdateInvoiceItemRequestSchema>;
