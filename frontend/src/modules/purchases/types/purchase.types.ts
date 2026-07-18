import { z } from 'zod';

import { PaginationMetaSchema } from '@/shared/types/pagination.types';
import { InvoiceStatusSchema } from '@/modules/invoices/types/invoice.types';

export const PurchaseListItemSchema = z.object({
  id: z.number(),

  buyer_client_id: z.number(),
  seller_company_id: z.number(),

  buyer_name: z.string(),
  seller_company: z.string(),

  status_invoice: InvoiceStatusSchema,

  total_amount: z.number(),
  paid_amount: z.number(),

  created_at: z.string(),
});

export type PurchaseListItem = z.infer<typeof PurchaseListItemSchema>;

export const GetPurchasesRequestSchema = z.object({
  page: z.number().int().min(1),
  limit: z.number().int().min(1).max(100),

  status_invoice: InvoiceStatusSchema.optional(),
  status: z.number().int().min(0).optional(),

  sort_column: z.enum(['created_at', 'total_amount', 'paid_amount', 'status_invoice']).optional(),

  order: z.enum(['asc', 'desc']).optional(),
});

export type GetPurchasesRequest = z.infer<typeof GetPurchasesRequestSchema>;

export const GetPurchasesResponseSchema = z.object({
  items: z.array(PurchaseListItemSchema),
  meta: PaginationMetaSchema,
});

export type GetPurchasesResponse = z.infer<typeof GetPurchasesResponseSchema>;
