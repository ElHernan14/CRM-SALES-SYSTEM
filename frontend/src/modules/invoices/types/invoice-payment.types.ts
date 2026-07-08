import { z } from 'zod';
import { PaginationMetaSchema } from '@/shared/types/pagination.types';
import { PaymentMethodSchema } from './invoice.types';

export const PaymentResponseSchema = z.object({
  id: z.number(),
  invoice_id: z.number(),
  amount: z.number(),
  payment_method: PaymentMethodSchema,
  paid_by_user_id: z.number(),
  created_at: z.string(),
});

export type PaymentResponse = z.infer<typeof PaymentResponseSchema>;

export const GetInvoicePaymentsRequestSchema = z.object({
  page: z.number(),
  limit: z.number(),
  payment_method: PaymentMethodSchema.optional(),
  sort_column: z.enum(['created_at', 'amount']).optional(),
  order: z.enum(['asc', 'desc']).optional(),
});

export type GetInvoicePaymentsRequest = z.infer<typeof GetInvoicePaymentsRequestSchema>;

export const GetInvoicePaymentsResponseSchema = z.object({
  items: z.array(PaymentResponseSchema),
  meta: PaginationMetaSchema,
});

export type GetInvoicePaymentsResponse = z.infer<typeof GetInvoicePaymentsResponseSchema>;
