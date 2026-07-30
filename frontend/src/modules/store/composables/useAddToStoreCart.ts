import { computed, ref } from 'vue';
import { useRouter } from 'vue-router';
import { useQueryClient } from '@tanstack/vue-query';
import { storeToRefs } from 'pinia';
import { toast } from 'vue-sonner';

import { useAuthStore } from '@/modules/auth/stores/auth.store';
import { getErrorMessage } from '@/shared/utils/error-handler';

import { useStoreUiStore } from '../stores/store-ui.store';

import { useEnsureStoreCart } from './useEnsureStoreCart';
import { useAddStoreCartItem } from './useAddStoreCartItem';

import type { StoreProduct } from '../types/store-product.types';

export function useAddToStoreCart() {
  const router = useRouter();
  const queryClient = useQueryClient();

  const auth = useAuthStore();
  const { user } = storeToRefs(auth);

  const storeUi = useStoreUiStore();

  const ensureCartMutation = useEnsureStoreCart();

  const addItemMutation = useAddStoreCartItem();

  const addingProductId = ref<number | null>(null);

  const isAdding = computed(() => {
    return (
      addingProductId.value !== null ||
      ensureCartMutation.isPending.value ||
      addItemMutation.isPending.value
    );
  });

  async function addToCart(product: StoreProduct, quantity = 1, openCart = true): Promise<boolean> {
    if (!auth.isAuthenticated) {
      toast.info('Sign in to add products', {
        description: 'Your cart is connected to your Nexora account.',
      });

      await router.push({
        path: '/login',
        query: {
          redirect: `/store/products/${product.id}`,
        },
      });

      return false;
    }

    if (!user.value) {
      toast.error('Unable to identify your account');
      return false;
    }

    if (user.value.company_id && product.company_id === user.value.company_id) {
      toast.error('Your business cannot purchase its own products');
      return false;
    }

    if (!Number.isInteger(quantity) || quantity < 1) {
      toast.error('Select a valid quantity');
      return false;
    }

    if (product.available_stock <= 0) {
      toast.error('This product is currently unavailable');

      return false;
    }

    if (quantity > product.available_stock) {
      toast.error('Quantity exceeds available stock');

      return false;
    }

    try {
      addingProductId.value = product.id;

      const ensuredCart = await ensureCartMutation.mutateAsync(product.company_id);

      await addItemMutation.mutateAsync({
        invoiceId: ensuredCart.invoice_id,

        sellerCompanyId: product.company_id,

        productId: product.id,

        quantity,
      });

      /*
       * Cancelamos cualquier respuesta vieja del carrito
       * antes de traer el estado definitivo.
       */
      await queryClient.cancelQueries({
        queryKey: ['b2c-store-carts'],
      });

      /*
       * Al invalidar y refetchear la query activa,
       * el badge y el drawer reciben el estado nuevo.
       */
      await Promise.all([
        queryClient.refetchQueries({
          queryKey: ['b2c-store-carts'],
          type: 'active',
        }),

        queryClient.refetchQueries({
          queryKey: ['b2c-store-products'],
          type: 'active',
        }),

        queryClient.invalidateQueries({
          queryKey: ['b2c-store-product', product.id],
        }),
      ]);

      toast.success('Added to cart', {
        description: `${quantity} × ${product.name}`,
      });

      if (openCart) {
        storeUi.openCart();
      }

      return true;
    } catch (error) {
      toast.error(getErrorMessage(error, 'Failed to add product to cart'));

      return false;
    } finally {
      addingProductId.value = null;
    }
  }

  return {
    addToCart,
    addingProductId,
    isAdding,
  };
}
