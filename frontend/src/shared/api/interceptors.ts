import { http } from "./http";
import { useAuthStore } from "@/modules/auth/stores/auth.store"
import { router } from "@/app/router";

export function setupInterceptors() {
  http.interceptors.request.use((config) => {
    const token = localStorage.getItem("access_token");

    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }

    return config;
  });

  http.interceptors.response.use(
    (response) => response,
    (error) => {
      const auth = useAuthStore()
      console.log(error)

      if (error.response?.status === 401) {
        // limpiar sesión
        auth.logout()

        // redirigir al login
        if (router.currentRoute.value.path !== "/login") {
          router.push("/login")
        }
      }

      return Promise.reject(error)
    }
  );
}