import { http } from '@/shared/api/http';
import { unwrapResponse } from '@/shared/utils/response';

import {
  // Get CompanyInvoices
  GetCompanyInvoicesResponseSchema,
  type GetCompanyInvoicesRequest,
  type GetCompanyInvoicesResponse,

  // Get CompanyInvoices details
  InvoiceDetailSchema,
  type InvoiceDetail,

  // Pay Invoice
  PayInvoiceResponseSchema,
  type PayInvoiceRequest,
  type PayInvoiceResponse,
} from '../types/invoice.types';

export async function getCompanyInvoices(
  params: GetCompanyInvoicesRequest
): Promise<GetCompanyInvoicesResponse> {
  const response = await http.get('/company/invoices', {
    params,
  });

  return unwrapResponse(GetCompanyInvoicesResponseSchema, response.data);
}

export async function getInvoiceById(id: number): Promise<InvoiceDetail> {
  const response = await http.get(`/invoice/${id}`);

  return unwrapResponse(InvoiceDetailSchema, response.data);
}

export async function submitInvoice(id: number): Promise<void> {
  await http.post(`/invoice/${id}/submit`);
}

export async function cancelInvoice(id: number): Promise<void> {
  await http.post(`/invoice/${id}/cancel`);
}

export async function payInvoice(
  id: number,
  payload: PayInvoiceRequest
): Promise<PayInvoiceResponse> {
  const response = await http.post(`/invoice/${id}/pay`, payload);

  return unwrapResponse(PayInvoiceResponseSchema, response.data);
}

import {
  InvoiceResponseSchema,
  type CreateInvoiceRequest,
  type InvoiceResponse,
} from '../types/invoice.types';

export async function createInvoice(payload: CreateInvoiceRequest): Promise<InvoiceResponse> {
  const response = await http.post('/invoice', payload);

  return unwrapResponse(InvoiceResponseSchema, response.data);
}
