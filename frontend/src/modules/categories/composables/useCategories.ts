import { useQuery } from '@tanstack/vue-query';

import { getCategories } from '../api/categories.api';

export function useCategories() {
  return useQuery({
    queryKey: ['categories'],
    queryFn: getCategories,

    // Son datos maestros; no deberían cambiar constantemente.
    staleTime: 1000 * 60 * 30,
  });
}
