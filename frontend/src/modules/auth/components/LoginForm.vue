<script setup lang="ts">
import { ref } from "vue"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { toast } from "vue-sonner"

import { login as loginAction } from "@/modules/auth/composables/useAuth"
// import { useAuthStore } from "@/modules/auth/stores/auth.store"

const email = ref("")
const password = ref("")

const loading = ref(false)

// const auth = useAuthStore()

async function onSubmit() {
  loading.value = true

  try {
    await loginAction({
      email: email.value,
      password: password.value
    })

    toast.success("Login successful")
    // el redirect ya lo maneja loginAction (ERP / STORE)
  } catch (e: any) {
    const errorMessage =
      e?.response?.data?.errorMessage ||
      "Invalid credentials. Please try again."
      toast.error(errorMessage)
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-background px-4">
    <Card class="w-full max-w-md shadow-xl border border-border">
      
      <CardHeader class="space-y-1">
        <CardTitle class="text-2xl font-semibold">
          Welcome back
        </CardTitle>

        <p class="text-sm text-muted-foreground">
          Sign in to your CRM & Commerce platform
        </p>
      </CardHeader>

      <CardContent>
        <form @submit.prevent="onSubmit" class="space-y-4">

          <!-- EMAIL -->
          <div class="space-y-2">
            <label class="text-sm font-medium">Email</label>
            <Input
              v-model="email"
              type="email"
              placeholder="you@company.com"
              autocomplete="email"
            />
          </div>

          <!-- PASSWORD -->
          <div class="space-y-2">
            <label class="text-sm font-medium">Password</label>
            <Input
              v-model="password"
              type="password"
              placeholder="••••••••"
              autocomplete="current-password"
            />
          </div>

          <!-- BUTTON -->
          <Button
            type="submit"
            class="w-full"
            :disabled="loading"
          >
            <span v-if="loading">Signing in...</span>
            <span v-else>Sign in</span>
          </Button>

        </form>
      </CardContent>
    </Card>
  </div>
</template>