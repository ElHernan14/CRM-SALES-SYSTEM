import { z } from 'zod';

import { PaginationMetaSchema } from '@/shared/types/pagination.types';

import { InvoiceStatusSchema } from '@/modules/invoices/types/invoice.types';

export const StorePurchaseSchema = z.object({
  id: z.number(),

  seller_company_id: z.number(),
  seller_company: z.string(),

  status_invoice: InvoiceStatusSchema,

  subtotal: z.number(),
  taxes: z.number(),

  total_amount: z.number(),
  paid_amount: z.number(),
  remaining_amount: z.number(),

  item_count: z.number(),

  created_at: z.string(),
});

export type StorePurchase = z.infer<typeof StorePurchaseSchema>;

export const GetStorePurchasesRequestSchema = z.object({
  page: z.number().int().min(1),
  limit: z.number().int().min(1).max(100),

  search: z.string().optional(),

  status_invoice: InvoiceStatusSchema.optional(),

  seller_company_id: z.number().int().positive().optional(),

  sort_column: z.enum(['created_at', 'total_amount', 'paid_amount', 'status_invoice']).optional(),

  order: z.enum(['asc', 'desc']).optional(),
});

export type GetStorePurchasesRequest = z.infer<typeof GetStorePurchasesRequestSchema>;

export const GetStorePurchasesResponseSchema = z.object({
  items: z.array(StorePurchaseSchema),
  meta: PaginationMetaSchema,
});

export type GetStorePurchasesResponse = z.infer<typeof GetStorePurchasesResponseSchema>;
