<script setup lang="ts">
import { ref } from 'vue';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { useUiStore } from '@/shared/stores/ui.store';
import { toast } from 'vue-sonner';

import { login as loginAction } from '@/modules/auth/composables/useAuth';
// import { useAuthStore } from "@/modules/auth/stores/auth.store"

const email = ref('');
const password = ref('');

const loading = ref(false);
const ui = useUiStore();

// const auth = useAuthStore()

async function onSubmit() {
  loading.value = true;
  ui.setLoading(true, 'Loading dashboard...');

  // Mostrar toast de carga y guardar id
  const toastId = toast.loading('Login session...');

  try {
    await loginAction({
      email: email.value,
      password: password.value,
    });

    // Actualizar el mismo toast a success
    toast.success('Login successful', { id: toastId, duration: 2000 });
  } catch (e: any) {
    const errorMessage =
      e?.response?.data?.errorMessage || 'Invalid credentials. Please try again.';

    // Actualizar el mismo toast a error
    toast.error(errorMessage, { id: toastId, duration: 2000 });
  } finally {
    loading.value = false;
    ui.setLoading(false);
  }
}

function testConfirm() {
  ui.openConfirm({
    title: 'Cancel invoice?',
    description: 'This action will cancel the invoice and release reserved stock.',
    confirmText: 'Cancel invoice',
    cancelText: 'Keep invoice',
    variant: 'destructive',
    onConfirm: async () => {
      toast.success('Invoice cancelled');
    },
  });
}
</script>

<template>
  <Card class="w-full max-w-md border-border/60 shadow-lg backdrop-blur">
    <CardHeader class="space-y-1">
      <CardTitle class="text-3xl font-semibold tracking-tight"> Welcome back </CardTitle>

      <p class="text-sm text-muted-foreground">
        Sign in to access your commerce account or business workspace.
      </p>
    </CardHeader>

    <CardContent class="space-y-6">
      <form @submit.prevent="onSubmit" class="space-y-5">
        <!-- EMAIL -->
        <div class="space-y-2">
          <label class="text-sm font-medium text-foreground">Email</label>
          <Input
            class="h-11"
            v-model="email"
            type="email"
            placeholder="you@company.com"
            autocomplete="email"
          />
        </div>

        <!-- PASSWORD -->
        <div class="space-y-2">
          <label class="text-sm font-medium text-foreground">Password</label>
          <Input
            class="h-11"
            v-model="password"
            type="password"
            placeholder="••••••••"
            autocomplete="current-password"
          />
        </div>

        <!-- BUTTON -->
        <Button type="submit" class="h-11 w-full" :disabled="loading">
          <span v-if="loading">Signing in...</span>
          <span v-else>Sign in</span>
        </Button>

        <div class="pt-2 text-center">
          <span class="text-xs text-muted-foreground"> Secure authentication powered by JWT </span>
        </div>
      </form>
    </CardContent>
  </Card>
</template>
