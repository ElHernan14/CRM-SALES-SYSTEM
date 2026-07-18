import { http } from '@/shared/api/http';
import { unwrapResponse } from '@/shared/utils/response';

import {
  PayPurchaseResponseSchema,
  type PayPurchaseRequest,
  type PayPurchaseResponse,
} from '../types/purchase-payment.types';

export async function payPurchase(
  invoiceId: number,
  payload: PayPurchaseRequest
): Promise<PayPurchaseResponse> {
  const response = await http.post(`/invoice/${invoiceId}/pay`, payload);

  return unwrapResponse(PayPurchaseResponseSchema, response.data);
}
