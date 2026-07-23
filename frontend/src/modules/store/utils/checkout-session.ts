import { CheckoutAllResponseSchema, type CheckoutAllResponse } from '../types/store-cart.types';

const CHECKOUT_RESULT_KEY = 'nexora_last_checkout';

export function saveCheckoutResult(result: CheckoutAllResponse) {
  sessionStorage.setItem(CHECKOUT_RESULT_KEY, JSON.stringify(result));
}

export function getCheckoutResult(): CheckoutAllResponse | null {
  const stored = sessionStorage.getItem(CHECKOUT_RESULT_KEY);

  if (!stored) return null;

  try {
    const parsed = JSON.parse(stored);

    return CheckoutAllResponseSchema.parse(parsed);
  } catch {
    sessionStorage.removeItem(CHECKOUT_RESULT_KEY);

    return null;
  }
}

export function clearCheckoutResult() {
  sessionStorage.removeItem(CHECKOUT_RESULT_KEY);
}
