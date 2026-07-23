import { useMutation, useQueryClient } from '@tanstack/vue-query';

import { uploadCompanyCover } from '../api/company.api';

export function useUploadCompanyCover() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: uploadCompanyCover,

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
