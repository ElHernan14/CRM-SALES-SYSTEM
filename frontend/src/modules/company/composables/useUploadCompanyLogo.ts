import { useMutation, useQueryClient } from '@tanstack/vue-query';

import { uploadCompanyLogo } from '../api/company.api';

export function useUploadCompanyLogo() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: uploadCompanyLogo,

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
