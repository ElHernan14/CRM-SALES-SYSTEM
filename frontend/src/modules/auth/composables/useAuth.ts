import { useAuthStore } from "@/modules/auth/stores/auth.store"
import { login as loginApi, me } from "@/modules/auth/api/auth.api"
import { router } from "@/app/router"
import type { LoginRequest, TenantContext } from "@/modules/auth/types/auth.types"
import { env } from "@/shared/config/env";

export async function login(credentials: LoginRequest) {
  const auth = useAuthStore()

  try {
    console.log(env.apiUrl)
    // 1. POST /auth/login
    const response = await loginApi(credentials)

    // 2. Guardar token
    auth.setToken(response.token)

    // 3. GET /auth/me
    const user: TenantContext = await me()

    // 4. Guardar user
    auth.setUser(user)

    // 5. Redirección según TenantContext
    if (user.companyID) {
      router.push("/erp/dashboard")
    } else {
      router.push("/store/dashboard")
    }
  } catch (error: any) {
    console.error("Login error:", error.message)
    throw error
  }
}

export function logout() {
  const auth = useAuthStore()
  auth.logout()
  router.push("/login")
}

export async function bootstrap() {
  const auth = useAuthStore()

  // 1. Leer token
  const token = localStorage.getItem("access_token")

  if (!token) {
    // 2. No existe → inicializar
    auth.setInitialized(true)
    return
  }

  try {
    // 3. Setear token
    auth.setToken(token)

    // 4. Llamar /auth/me
    const user = await me()

    // 5. Guardar usuario
    auth.setUser(user)
  } catch (error) {
    console.error("Bootstrap error:", error)
    // Si falla, limpiar sesión
    auth.logout()
  } finally {
    // 6. Inicializar
    auth.setInitialized(true)
  }
}
