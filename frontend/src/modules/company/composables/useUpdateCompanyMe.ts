import { useMutation, useQueryClient } from '@tanstack/vue-query';

import { updateCompanyMe } from '../api/company.api';

export function useUpdateCompanyMe() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: updateCompanyMe,

    onSuccess: async () => {
      await Promise.all([
        queryClient.invalidateQueries({
          queryKey: ['company-me'],
        }),

        queryClient.invalidateQueries({
          queryKey: ['marketplace-suppliers'],
        }),
      ]);
    },
  });
}
