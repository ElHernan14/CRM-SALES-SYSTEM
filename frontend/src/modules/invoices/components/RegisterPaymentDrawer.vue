<script setup lang="ts">
import { computed, reactive, watch } from 'vue';
import { toast } from 'vue-sonner';

import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
} from '@/components/ui/sheet';

import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';

import { usePayInvoice } from '../composables/usePayInvoice';
import type { CompanyInvoiceListItem } from '../types/invoice.types';

const props = defineProps<{
  open: boolean;
  invoice: CompanyInvoiceListItem | null;
}>();

const emit = defineEmits<{
  'update:open': [value: boolean];
}>();

const form = reactive({
  amount: '',
  payment_method: 'cash' as 'cash' | 'transfer' | 'card',
});

const payMutation = usePayInvoice();

const isSubmitting = computed(() => payMutation.isPending.value);

const remainingAmount = computed(() => {
  if (!props.invoice) return 0;

  return props.invoice.total_amount - props.invoice.paid_amount;
});

watch(
  () => props.invoice,
  (invoice) => {
    if (!invoice) return;

    form.amount = String(Math.max(invoice.total_amount - invoice.paid_amount, 0));
    form.payment_method = 'cash';
  },
  { immediate: true }
);

function formatCurrency(value: number) {
  return `$${value.toFixed(2)}`;
}

async function onSubmit() {
  if (!props.invoice) return;

  const amount = Number(form.amount);

  if (!amount || amount <= 0) {
    toast.error('Enter a valid payment amount');
    return;
  }

  if (amount > remainingAmount.value) {
    toast.error('Payment amount exceeds remaining balance');
    return;
  }

  try {
    await payMutation.mutateAsync({
      id: props.invoice.id,
      payload: {
        amount,
        payment_method: form.payment_method,
      },
    });

    toast.success('Payment registered');
    emit('update:open', false);
  } catch {
    toast.error('Failed to register payment');
  }
}
</script>

<template>
  <Sheet :open="open" @update:open="emit('update:open', $event)">
    <SheetContent class="w-full sm:max-w-xl">
      <SheetHeader>
        <SheetTitle> Register payment </SheetTitle>

        <SheetDescription> Record a received payment for this sales invoice. </SheetDescription>
      </SheetHeader>

      <div v-if="invoice" class="mt-6 space-y-6">
        <div class="rounded-xl border border-border bg-muted/40 p-4">
          <p class="text-sm text-muted-foreground">Invoice</p>

          <p class="mt-1 text-xl font-semibold">#INV-{{ invoice.id }}</p>
        </div>

        <div class="grid grid-cols-2 gap-3 sm:grid-cols-3">
          <div class="rounded-xl border border-border bg-card p-4">
            <p class="text-xs text-muted-foreground">Total</p>

            <p class="mt-2 text-sm font-semibold">
              {{ formatCurrency(invoice.total_amount) }}
            </p>
          </div>

          <div class="rounded-xl border border-border bg-card p-4">
            <p class="text-xs text-muted-foreground">Paid</p>

            <p class="mt-2 text-sm font-semibold">
              {{ formatCurrency(invoice.paid_amount) }}
            </p>
          </div>

          <div class="rounded-xl border border-border bg-card p-4">
            <p class="text-xs text-muted-foreground">Remaining</p>

            <p class="mt-2 text-sm font-semibold">
              {{ formatCurrency(remainingAmount) }}
            </p>
          </div>
        </div>

        <form class="space-y-5" @submit.prevent="onSubmit">
          <div class="space-y-2">
            <label class="text-sm font-medium text-foreground"> Amount </label>

            <Input v-model="form.amount" type="text" inputmode="decimal" placeholder="0.00" />
          </div>

          <div class="space-y-2">
            <label class="text-sm font-medium text-foreground"> Payment method </label>

            <Select v-model="form.payment_method">
              <SelectTrigger>
                <SelectValue placeholder="Payment method" />
              </SelectTrigger>

              <SelectContent>
                <SelectItem value="cash"> Cash </SelectItem>

                <SelectItem value="transfer"> Transfer </SelectItem>

                <SelectItem value="card"> Card </SelectItem>
              </SelectContent>
            </Select>
          </div>

          <div class="flex justify-end gap-2 pt-4">
            <Button type="button" variant="outline" @click="emit('update:open', false)">
              Cancel
            </Button>

            <Button type="submit" :disabled="isSubmitting">
              {{ isSubmitting ? 'Registering...' : 'Register payment' }}
            </Button>
          </div>
        </form>
      </div>
    </SheetContent>
  </Sheet>
</template>
