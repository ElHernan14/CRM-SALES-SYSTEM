import { z } from 'zod';

export const CartItemSchema = z.object({
  id: z.number(),
  invoice_id: z.number(),
  product_id: z.number(),
  product_name: z.string(),
  quantity: z.number(),
  price: z.number(),
  subtotal: z.number(),
});

export type CartItem = z.infer<typeof CartItemSchema>;

export const CartSchema = z.object({
  invoice_id: z.number(),
  buyer_client_id: z.number(),
  seller_company_id: z.number(),
  seller_company: z.string(),
  status_invoice: z.literal('draft'),
  subtotal: z.number(),
  taxes: z.number(),
  total_amount: z.number(),
  items: z.array(CartItemSchema),
});

export type Cart = z.infer<typeof CartSchema>;

export const EnsureCartRequestSchema = z.object({
  seller_company_id: z.number().int().positive(),
});

export type EnsureCartRequest = z.infer<typeof EnsureCartRequestSchema>;

export const EnsureCartResponseSchema = z.object({
  invoice_id: z.number(),
  buyer_client_id: z.number(),
  seller_company_id: z.number(),
  status_invoice: z.literal('draft'),
});

export type EnsureCartResponse = z.infer<typeof EnsureCartResponseSchema>;

export const CheckoutRequestSchema = z
  .object({
    seller_company_id: z.number().int().positive().optional(),
    invoice_id: z.number().int().positive().optional(),
  })
  .refine(
    (request) => request.seller_company_id !== undefined || request.invoice_id !== undefined,
    {
      message: 'seller_company_id or invoice_id is required',
    }
  );

export type CheckoutRequest = z.infer<typeof CheckoutRequestSchema>;

export const CheckoutResponseSchema = z.object({
  invoice_id: z.number(),
  status: z.string(),
});

export type CheckoutResponse = z.infer<typeof CheckoutResponseSchema>;

export const NullableCartSchema = CartSchema.nullable();

// MARKETPLACE TYPES
export const MarketplaceCartItemSchema = z.object({
  id: z.number().int().positive(),
  product_id: z.number().int().positive(),
  product_name: z.string(),
  price: z.number(),
  quantity: z.number().int().positive(),
  subtotal: z.number(),
});

export const MarketplacePurchaseCartSchema = z.object({
  invoice_id: z.number().int().positive(),
  buyer_client_id: z.number().int().positive(),
  seller_company_id: z.number().int().positive(),
  seller_company: z.string(),
  status_invoice: z.literal('draft'),
  source: z.literal('erp'),

  subtotal: z.number(),
  taxes: z.number(),
  total_amount: z.number(),

  items: z.array(MarketplaceCartItemSchema),
});

export const MarketplaceCheckoutResultSchema = z.object({
  invoice_id: z.number().int().positive(),
  status: z.literal('pending'),
  source: z.literal('erp'),
});

export type MarketplacePurchaseCart = z.infer<typeof MarketplacePurchaseCartSchema>;

export type MarketplaceCartItem = z.infer<typeof MarketplaceCartItemSchema>;

export type MarketplaceCheckoutResult = z.infer<typeof MarketplaceCheckoutResultSchema>;
