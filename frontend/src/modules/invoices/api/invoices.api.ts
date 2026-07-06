import { http } from '@/shared/api/http';
import { unwrapResponse } from '@/shared/utils/response';

import {
  GetCompanyInvoicesResponseSchema,
  type GetCompanyInvoicesRequest,
  type GetCompanyInvoicesResponse,
} from '../types/invoice.types';

export async function getCompanyInvoices(
  params: GetCompanyInvoicesRequest
): Promise<GetCompanyInvoicesResponse> {
  const response = await http.get('/company/invoices', {
    params,
  });

  return unwrapResponse(GetCompanyInvoicesResponseSchema, response.data);
}
