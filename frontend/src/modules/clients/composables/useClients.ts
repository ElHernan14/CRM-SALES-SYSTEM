import { computed, type Ref } from 'vue';
import { useQuery } from '@tanstack/vue-query';

import { getClients } from '../api/clients.api';
import type { GetClientsRequest } from '../types/client.types';

export function useClients(params: Ref<GetClientsRequest>) {
  return useQuery({
    queryKey: computed(() => ['clients', params.value]),

    queryFn: () => getClients(params.value),
  });
}
