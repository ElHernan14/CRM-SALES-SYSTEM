import { http } from "@/shared/api/http";
import type {
  LoginRequest,
  LoginResponse,
  TenantContext,
} from "../types/auth.types";

export async function login(payload: LoginRequest) {
  try {
    const response = await http.post(
      "/auth/login",
      payload
    )

    console.log("AXIOS RESPONSE", response)
    console.log("AXIOS DATA", response.data)

    return response.data

  } catch (error) {
    console.error("AXIOS ERROR", error)
    throw error
  }
}

export async function me() {
  const { data } =
    await http.get<TenantContext>(
      "/auth/me"
    );

  return data;
}