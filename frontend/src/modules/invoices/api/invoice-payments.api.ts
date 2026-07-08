import { http } from '@/shared/api/http';
import { unwrapResponse } from '@/shared/utils/response';

import {
  GetInvoicePaymentsResponseSchema,
  type GetInvoicePaymentsRequest,
  type GetInvoicePaymentsResponse,
} from '../types/invoice-payment.types';

export async function getInvoicePayments(
  invoiceId: number,
  params: GetInvoicePaymentsRequest
): Promise<GetInvoicePaymentsResponse> {
  const response = await http.get(`/invoice/${invoiceId}/payments`, { params });

  return unwrapResponse(GetInvoicePaymentsResponseSchema, response.data);
}
