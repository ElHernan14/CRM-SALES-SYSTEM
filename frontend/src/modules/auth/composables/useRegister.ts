import { useMutation } from '@tanstack/vue-query';
import { useRouter } from 'vue-router';
import { toast } from 'vue-sonner';

import { queryClient } from '@/app/providers/query-client';

import { useAuthStore } from '../stores/auth.store';

import { registerBusinessApi, registerPersonalApi } from '../api/register.api';

import type {
  BusinessRegisterRequest,
  PersonalRegisterRequest,
  RegisterResponse,
} from '../types/register.types';

export function useRegister() {
  const router = useRouter();
  const auth = useAuthStore();

  async function establishSession(response: RegisterResponse) {
    await queryClient.cancelQueries();
    queryClient.clear();

    localStorage.setItem('access_token', response.token);

    auth.setToken(response.token);
    auth.setUser(response.user);
    auth.setInitialized(true);
  }

  const personalRegisterMutation = useMutation({
    mutationFn: (payload: PersonalRegisterRequest) => registerPersonalApi(payload),

    onSuccess: async (response) => {
      await establishSession(response);

      toast.success('Welcome to Nexora', {
        description: 'Your personal account was created successfully.',
      });

      await router.replace('/store');
    },
  });

  /*
   * Business no redirige automáticamente.
   * El formulario debe mostrar el paso de branding.
   */
  const businessRegisterMutation = useMutation({
    mutationFn: (payload: BusinessRegisterRequest) => registerBusinessApi(payload),
  });

  return {
    establishSession,

    personalRegisterMutation,
    businessRegisterMutation,
  };
}
