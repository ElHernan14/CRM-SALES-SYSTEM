import { z } from 'zod';

export const PaymentMethodSchema = z.enum(['cash', 'card', 'transfer']);

export type PaymentMethod = z.infer<typeof PaymentMethodSchema>;

export const PayPurchaseRequestSchema = z.object({
  amount: z.number().positive(),
  payment_method: PaymentMethodSchema,
});

export type PayPurchaseRequest = z.infer<typeof PayPurchaseRequestSchema>;

export const PayPurchaseResponseSchema = z.object({
  payment_id: z.number(),
  status_invoice: z.string(),
  total_amount: z.number(),
  paid_amount: z.number(),
  remaining_amount: z.number(),
});

export type PayPurchaseResponse = z.infer<typeof PayPurchaseResponseSchema>;
