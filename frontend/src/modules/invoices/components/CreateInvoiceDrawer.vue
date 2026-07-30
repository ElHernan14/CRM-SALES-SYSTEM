<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import { storeToRefs } from 'pinia';
import { toast } from 'vue-sonner';
import { FilePlus2, Loader2, UserRound, Check, ChevronsUpDown, Search } from 'lucide-vue-next';

import { Button } from '@/components/ui/button';
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
} from '@/components/ui/sheet';
import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/popover';
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from '@/components/ui/command';

import { useAuthStore } from '@/modules/auth/stores/auth.store';
import { useClients } from '@/modules/clients/composables/useClients';
import { useCreateInvoice } from '../composables/useCreateInvoice';
import { getErrorMessage } from '@/shared/utils/error-handler';

const props = defineProps<{
  open: boolean;
}>();

const emit = defineEmits<{
  'update:open': [value: boolean];
  created: [invoiceId: number];
}>();

const auth = useAuthStore();
const { user } = storeToRefs(auth);

// State
const selectedClientId = ref('');
const clientSearch = ref('');
const clientsPopoverOpen = ref(false);

function selectClient(clientId: number) {
  selectedClientId.value = String(clientId);
  clientsPopoverOpen.value = false;
}

watch(
  () => props.open,
  (open) => {
    if (!open) {
      clientSearch.value = '';
      selectedClientId.value = '';
    }
  }
);

const clientsParams = computed(() => ({
  page: 1,
  limit: 20,
  search: clientSearch.value || undefined,
}));

const { data: clientsData, isLoading: clientsLoading } = useClients(clientsParams);

const clients = computed(() => {
  return clientsData.value?.clients ?? [];
});

const selectedClient = computed(() => {
  return clients.value.find((client) => client.id === Number(selectedClientId.value));
});

const createInvoiceMutation = useCreateInvoice();

const isSubmitting = computed(() => createInvoiceMutation.isPending.value);

async function onSubmit() {
  const buyerClientId = Number(selectedClientId.value);
  const sellerCompanyId = user.value?.company_id;

  if (!buyerClientId) {
    toast.error('Select a buyer client');
    return;
  }

  if (!sellerCompanyId) {
    toast.error('Missing seller company context');
    return;
  }

  try {
    const invoice = await createInvoiceMutation.mutateAsync({
      buyer_client_id: buyerClientId,
      seller_company_id: sellerCompanyId,
    });

    toast.success('Invoice created');

    selectedClientId.value = '';

    emit('update:open', false);
    emit('created', invoice.id);
  } catch (error) {
    toast.error(getErrorMessage(error, 'Failed to create invoice'));
  }
}
</script>

<template>
  <Sheet :open="props.open" @update:open="emit('update:open', $event)">
    <SheetContent class="w-full sm:max-w-xl">
      <SheetHeader>
        <SheetTitle class="flex items-center gap-2">
          <FilePlus2 class="h-5 w-5" />
          Create sales invoice
        </SheetTitle>

        <SheetDescription>
          Select the buyer client and create a draft invoice for this workspace.
        </SheetDescription>
      </SheetHeader>

      <form class="mt-6 space-y-6" @submit.prevent="onSubmit">
        <div class="space-y-2">
          <label class="text-sm font-medium text-foreground"> Buyer client </label>

          <Popover v-model:open="clientsPopoverOpen">
            <PopoverTrigger as-child>
              <Button
                type="button"
                variant="outline"
                role="combobox"
                class="h-auto min-h-10 w-full justify-between px-3 py-2 text-left"
              >
                <div v-if="selectedClient" class="min-w-0">
                  <p class="truncate text-sm font-medium text-foreground">
                    {{ selectedClient.first_name }} {{ selectedClient.last_name }}
                  </p>

                  <p class="truncate text-xs text-muted-foreground">
                    {{ selectedClient.email }}
                    <template v-if="selectedClient.company_name">
                      · {{ selectedClient.company_name }}
                    </template>
                  </p>
                </div>

                <span v-else class="text-sm text-muted-foreground"> Search buyer client... </span>

                <ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 text-muted-foreground" />
              </Button>
            </PopoverTrigger>

            <PopoverContent class="w-[--radix-popover-trigger-width] p-0" align="start">
              <Command>
                <div class="flex items-center border-b border-border px-3">
                  <Search class="mr-2 h-4 w-4 shrink-0 text-muted-foreground" />

                  <CommandInput
                    v-model="clientSearch"
                    placeholder="Search by name or email..."
                    class="h-11"
                  />
                </div>

                <CommandList>
                  <CommandEmpty v-if="clientsLoading">
                    {{ 'Loading clients...' }}
                  </CommandEmpty>

                  <CommandEmpty v-if="!clientsLoading && clients.length === 0">
                    No clients found.
                  </CommandEmpty>

                  <CommandGroup>
                    <CommandItem
                      v-for="client in clients"
                      :key="client.id"
                      :value="`${client.first_name} ${client.last_name} ${client.email}`"
                      @select="selectClient(client.id)"
                    >
                      <div class="flex w-full items-center justify-between gap-3">
                        <div class="min-w-0">
                          <p class="truncate text-sm font-medium">
                            {{ client.first_name }} {{ client.last_name }}
                          </p>

                          <p class="truncate text-xs text-muted-foreground">
                            {{ client.email }}
                            <template v-if="client.company_name">
                              · {{ client.company_name }}
                            </template>
                          </p>
                        </div>

                        <Check
                          v-if="selectedClientId === String(client.id)"
                          class="h-4 w-4 shrink-0 text-primary"
                        />
                      </div>
                    </CommandItem>
                  </CommandGroup>
                </CommandList>
              </Command>
            </PopoverContent>
          </Popover>
        </div>

        <div v-if="selectedClient" class="rounded-xl border border-border bg-muted/40 p-4">
          <div class="mb-3 flex items-center gap-2">
            <div class="rounded-lg border border-border bg-card p-2 text-muted-foreground">
              <UserRound class="h-4 w-4" />
            </div>

            <div>
              <p class="text-sm font-medium text-foreground">
                {{ selectedClient.first_name }} {{ selectedClient.last_name }}
              </p>

              <p class="text-xs text-muted-foreground">
                {{ selectedClient.email }}
              </p>
            </div>
          </div>

          <div class="grid gap-3 sm:grid-cols-2">
            <div class="rounded-lg border border-border bg-card p-3">
              <p class="text-xs text-muted-foreground">Status</p>

              <p class="mt-1 text-sm font-medium text-foreground">
                {{ selectedClient.status === 1 ? 'Active' : 'Inactive' }}
              </p>
            </div>

            <div
              v-if="selectedClient.company_id"
              class="rounded-lg border border-border bg-card p-3"
            >
              <p class="text-xs text-muted-foreground">Company</p>

              <p class="mt-1 text-sm font-medium text-foreground">
                {{ selectedClient.company_name ?? `Company #${selectedClient.company_id}` }}
              </p>
            </div>
          </div>
        </div>

        <div class="rounded-xl border border-border bg-card p-4">
          <p class="text-sm font-medium text-foreground">What happens next?</p>

          <p class="mt-1 text-sm text-muted-foreground">
            The invoice will be created as a draft. You can add products, review totals and submit
            it when ready.
          </p>
        </div>

        <div class="grid grid-cols-2 gap-2 pt-2 sm:flex sm:justify-end">
          <Button
            type="button"
            variant="outline"
            class="w-full sm:w-auto"
            @click="emit('update:open', false)"
          >
            Cancel
          </Button>

          <Button type="submit" class="w-full sm:w-auto" :disabled="isSubmitting">
            <span class="mr-2 flex h-4 w-4 items-center justify-center">
              <Loader2 v-if="isSubmitting" class="h-4 w-4 animate-spin" />

              <FilePlus2 v-else class="h-4 w-4" />
            </span>

            {{ isSubmitting ? 'Creating...' : 'Create invoice' }}
          </Button>
        </div>
      </form>
    </SheetContent>
  </Sheet>
</template>
