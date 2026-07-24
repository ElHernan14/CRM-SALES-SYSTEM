import { useMutation, useQueryClient } from '@tanstack/vue-query';

import { updateClientProfile } from '../api/account.api';

import type { UpdateClientProfileRequest } from '../types/account.types';

export function useUpdateClientProfile() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({
      clientId,
      payload,
    }: {
      clientId: number;
      payload: UpdateClientProfileRequest;
    }) => updateClientProfile(clientId, payload),

    onSuccess: async () => {
      await queryClient.invalidateQueries({
        queryKey: ['client-profile'],
      });
    },
  });
}
