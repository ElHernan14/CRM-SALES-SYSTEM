import type { Router, RouteLocationNormalized } from "vue-router"
import { useAuthStore } from "../stores/auth.store"

export function setupAuthGuard(router: Router) {
  router.beforeEach((to: RouteLocationNormalized) => {
    const auth = useAuthStore()

    if (!auth.initialized) {
      return false
    }

    const requiresAuth = to.matched.some(route => route.meta.requiresAuth)
    const guestOnly = to.matched.some(route => route.meta.guestOnly)

    if (requiresAuth && !auth.isAuthenticated) {
      return "/login"
    }

    if (guestOnly && auth.isAuthenticated) {
      return "/erp/dashboard"
    }

    return true
  })
}
