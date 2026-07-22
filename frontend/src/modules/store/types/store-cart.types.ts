import { z } from 'zod';

export const StoreCartItemSchema = z.object({
  id: z.number(),
  invoice_id: z.number(),

  product_id: z.number(),
  product_name: z.string(),

  product_image_path: z.string().nullable().optional(),

  quantity: z.number(),
  price: z.number(),
  subtotal: z.number(),
});

export type StoreCartItem = z.infer<typeof StoreCartItemSchema>;

export const StoreSellerCartSchema = z.object({
  invoice_id: z.number(),
  buyer_client_id: z.number(),

  seller_company_id: z.number(),
  seller_company: z.string(),

  status_invoice: z.string(),

  subtotal: z.number(),
  taxes: z.number(),
  total_amount: z.number(),

  items: z.array(StoreCartItemSchema),
});

export type StoreSellerCart = z.infer<typeof StoreSellerCartSchema>;

export const StoreCartsSummarySchema = z.object({
  seller_count: z.number(),
  item_count: z.number(),

  subtotal: z.number(),
  taxes: z.number(),
  total_amount: z.number(),
});

export type StoreCartsSummary = z.infer<typeof StoreCartsSummarySchema>;

export const GetStoreCartsResponseSchema = z.object({
  carts: z.array(StoreSellerCartSchema),
  summary: StoreCartsSummarySchema,
});

export type GetStoreCartsResponse = z.infer<typeof GetStoreCartsResponseSchema>;

export const CheckoutAllRequestSchema = z.object({
  invoice_ids: z.array(z.number().int().positive()).optional(),
});

export type CheckoutAllRequest = z.infer<typeof CheckoutAllRequestSchema>;

export const CheckoutAllOrderSchema = z.object({
  invoice_id: z.number(),

  seller_company_id: z.number(),
  seller_company: z.string(),

  status: z.string(),
});

export const CheckoutAllResponseSchema = z.object({
  orders: z.array(CheckoutAllOrderSchema),
  count: z.number(),
  total_amount: z.number(),
});

export type CheckoutAllResponse = z.infer<typeof CheckoutAllResponseSchema>;
