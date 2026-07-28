import { z } from 'zod';

export const EnsureMarketplaceCartRequestSchema = z.object({
  seller_company_id: z.number().int().positive(),
});

export const MarketplaceCartSourceSchema = z.literal('erp');

export const EnsureMarketplaceCartSchema = z.object({
  invoice_id: z.number().int().positive(),

  buyer_client_id: z.number().int().positive(),

  seller_company_id: z.number().int().positive(),

  status_invoice: z.literal('draft'),

  source: MarketplaceCartSourceSchema,
});

export const EnsureMarketplaceCartResponseSchema = z.object({
  status: z.string(),

  code: z.number(),

  data: EnsureMarketplaceCartSchema,
});

export type EnsureMarketplaceCartRequest = z.infer<typeof EnsureMarketplaceCartRequestSchema>;

export type EnsureMarketplaceCart = z.infer<typeof EnsureMarketplaceCartSchema>;

export type EnsureMarketplaceCartResponse = z.infer<typeof EnsureMarketplaceCartResponseSchema>;
