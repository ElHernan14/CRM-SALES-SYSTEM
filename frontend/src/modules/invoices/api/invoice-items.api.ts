import { http } from '@/shared/api/http';
import { unwrapResponse } from '@/shared/utils/response';

import {
  GetInvoiceItemsResponseSchema,
  InvoiceItemSchema,
  type CreateInvoiceItemRequest,
  type GetInvoiceItemsRequest,
  type GetInvoiceItemsResponse,
  type InvoiceItem,
  type UpdateInvoiceItemRequest,
} from '../types/invoice-item.types';

export async function getInvoiceItems(
  invoiceId: number,
  params: GetInvoiceItemsRequest
): Promise<GetInvoiceItemsResponse> {
  const response = await http.get(`/invoice/${invoiceId}/items`, { params });

  return unwrapResponse(GetInvoiceItemsResponseSchema, response.data);
}

export async function createInvoiceItem(
  invoiceId: number,
  payload: CreateInvoiceItemRequest
): Promise<InvoiceItem> {
  const response = await http.post(`/invoice/${invoiceId}/items`, payload);

  return unwrapResponse(InvoiceItemSchema, response.data);
}

export async function updateInvoiceItem(
  invoiceId: number,
  itemId: number,
  payload: UpdateInvoiceItemRequest
): Promise<InvoiceItem> {
  const response = await http.patch(`/invoice/${invoiceId}/items/${itemId}`, payload);

  return unwrapResponse(InvoiceItemSchema, response.data);
}

export async function deleteInvoiceItem(invoiceId: number, itemId: number): Promise<void> {
  await http.delete(`/invoice/${invoiceId}/items/${itemId}`);
}
