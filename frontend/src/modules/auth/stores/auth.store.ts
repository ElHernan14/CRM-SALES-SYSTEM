import { defineStore } from 'pinia'
import type {TenantContext} from '@/modules/auth/types/auth.types'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: null as string | null,
    user: null as TenantContext | null,
    loading: false,
    initialized: false
  }),

  getters: {
    isAuthenticated: (state) => !!state.token && !!state.user,
    permissions: (state) => state.user?.permissions ?? [],
    roles: (state) => state.user?.roles ?? []
  },

  actions: {
    setToken(token: string | null) {
      this.token = token

      if (token) {
        localStorage.setItem("access_token", token)
      } else {
        localStorage.removeItem("access_token")
      }
    },
    setUser(user: TenantContext | null) {
      this.user = user
    },
    setLoading(value: boolean) {
      this.loading = value
    },
    setInitialized(value: boolean) {
      this.initialized = value
    },
    logout() {
      this.token = null
      this.user = null

      localStorage.removeItem("access_token")
    }
  }
})
