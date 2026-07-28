import { http } from '@/shared/api/http';

import {
  EnsureMarketplaceCartRequestSchema,
  EnsureMarketplaceCartResponseSchema,
  type EnsureMarketplaceCart,
  type EnsureMarketplaceCartRequest,
} from '../types/marketplace-cart.types';

export async function ensureMarketplaceCart(
  sellerCompanyId: number
): Promise<EnsureMarketplaceCart> {
  const payload: EnsureMarketplaceCartRequest = EnsureMarketplaceCartRequestSchema.parse({
    seller_company_id: sellerCompanyId,
  });

  const response = await http.post('/marketplace/cart/ensure', payload);

  const parsed = EnsureMarketplaceCartResponseSchema.parse(response.data);

  return parsed.data;
}
