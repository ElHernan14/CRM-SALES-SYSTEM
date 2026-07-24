import { z } from 'zod';

import { InvoiceStatusSchema } from '@/modules/invoices/types/invoice.types';

export const DashboardSalesSummarySchema = z.object({
  total_invoices: z.number(),
  draft: z.number(),
  pending: z.number(),
  paid: z.number(),
  cancelled: z.number(),

  total_amount: z.number(),
  paid_amount: z.number(),
  outstanding_amount: z.number(),
});

export type DashboardSalesSummary = z.infer<typeof DashboardSalesSummarySchema>;

export const DashboardPurchasesSummarySchema = DashboardSalesSummarySchema;

export type DashboardPurchasesSummary = z.infer<typeof DashboardPurchasesSummarySchema>;

export const DashboardInventorySummarySchema = z.object({
  active_products: z.number(),
  inactive_products: z.number(),

  available_units: z.number(),
  reserved_units: z.number(),

  low_stock_products: z.number(),
  out_of_stock_products: z.number(),
});

export type DashboardInventorySummary = z.infer<typeof DashboardInventorySummarySchema>;

export const DashboardRecentInvoiceSchema = z.object({
  id: z.number(),

  counterparty: z.string(),

  status_invoice: InvoiceStatusSchema,

  total_amount: z.number(),
  paid_amount: z.number(),
  remaining_amount: z.number(),

  created_at: z.string(),
});

export type DashboardRecentInvoice = z.infer<typeof DashboardRecentInvoiceSchema>;

export const DashboardLowStockProductSchema = z.object({
  id: z.number(),

  name: z.string(),

  category: z.string(),
  type: z.string(),

  stock: z.number(),
  reserved_stock: z.number(),
  available_stock: z.number(),

  image_path: z.string().nullable().optional(),
});

export type DashboardLowStockProduct = z.infer<typeof DashboardLowStockProductSchema>;

export const DashboardAttentionSchema = z.object({
  pending_sales: z.number(),
  draft_sales: z.number(),

  pending_purchases: z.number(),
  partially_paid_purchases: z.number(),

  low_stock_products: z.number(),
  out_of_stock_products: z.number(),
});

export type DashboardAttention = z.infer<typeof DashboardAttentionSchema>;

export const DashboardOverviewSchema = z.object({
  sales: DashboardSalesSummarySchema,

  purchases: DashboardPurchasesSummarySchema,

  inventory: DashboardInventorySummarySchema,

  attention: DashboardAttentionSchema,

  recent_sales: z.array(DashboardRecentInvoiceSchema),

  recent_purchases: z.array(DashboardRecentInvoiceSchema),

  low_stock_products: z.array(DashboardLowStockProductSchema),
});

export type DashboardOverview = z.infer<typeof DashboardOverviewSchema>;
