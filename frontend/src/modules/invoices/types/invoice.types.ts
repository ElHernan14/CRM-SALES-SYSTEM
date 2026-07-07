import { z } from 'zod';
import { PaginationMetaSchema } from '@/shared/types/pagination.types';

// Status
export const InvoiceStatusSchema = z.enum(['draft', 'pending', 'paid', 'cancelled']);
export type InvoiceStatus = z.infer<typeof InvoiceStatusSchema>;

// Orden
export const SortOrderSchema = z.enum(['asc', 'desc']);
export type SortOrder = z.infer<typeof SortOrderSchema>;

// Columnas de orden
export const SortColumnSchema = z.enum([
  'created_at',
  'total_amount',
  'paid_amount',
  'status_invoice',
]);
export type SortColumn = z.infer<typeof SortColumnSchema>;

// Request completo
export const GetCompanyInvoicesRequestSchema = z.object({
  page: z.number(),
  limit: z.number(),
  status_invoice: InvoiceStatusSchema.optional(),
  status: z.number().optional(),
  sort_column: SortColumnSchema.optional(),
  order: SortOrderSchema.optional(),
});

export type GetCompanyInvoicesRequest = z.infer<typeof GetCompanyInvoicesRequestSchema>;

export const CompanyInvoiceListItemSchema = z.object({
  id: z.number(),
  buyer_client_id: z.number(),
  seller_company_id: z.number(),
  status_invoice: InvoiceStatusSchema,
  total_amount: z.number(),
  paid_amount: z.number(),
  created_at: z.string(),
  buyer_name: z.string(),
  seller_company: z.string(),
});

export type CompanyInvoiceListItem = z.infer<typeof CompanyInvoiceListItemSchema>;

export const GetCompanyInvoicesResponseSchema = z.object({
  items: z.array(CompanyInvoiceListItemSchema),
  meta: PaginationMetaSchema,
});

export type GetCompanyInvoicesResponse = z.infer<typeof GetCompanyInvoicesResponseSchema>;

export const InvoiceDetailSchema = z.object({
  id: z.number(),
  buyer_client_id: z.number(),
  seller_company_id: z.number(),
  created_by_user_id: z.number(),
  status_invoice: InvoiceStatusSchema,
  subtotal: z.number(),
  taxes: z.number(),
  total_amount: z.number(),
  paid_amount: z.number(),
  status: z.number(),
  created_at: z.string(),
  updated_at: z.string().nullable().optional(),
  deleted_at: z.string().nullable().optional(),
  buyer_name: z.string(),
  seller_company: z.string(),
});

export type InvoiceDetail = z.infer<typeof InvoiceDetailSchema>;

// Pay Invoice
export const PaymentMethodSchema = z.enum(['transfer', 'cash', 'card']);

export const PayInvoiceRequestSchema = z.object({
  amount: z.number().positive(),
  payment_method: PaymentMethodSchema,
});

export type PayInvoiceRequest = z.infer<typeof PayInvoiceRequestSchema>;

export const PayInvoiceResponseSchema = z.object({
  payment_id: z.number(),
  status_invoice: InvoiceStatusSchema,
  total_amount: z.number(),
  paid_amount: z.number(),
  remaining_amount: z.number(),
});

export type PayInvoiceResponse = z.infer<typeof PayInvoiceResponseSchema>;
