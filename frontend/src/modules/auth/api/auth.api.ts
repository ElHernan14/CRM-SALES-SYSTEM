// auth.api.ts
import { http } from "@/shared/api/http"
import { unwrapResponse } from "@/shared/utils/response"
import {
  LoginResponseSchema,
  TenantContextSchema,
} from "../types/auth.types"
import type {
  LoginRequest,
  LoginResponse,
  TenantContext,
} from "../types/auth.types"

export async function login(payload: LoginRequest) {
  const response = await http.post("/auth/login", payload)

  // Validar y extraer solo data
  const data: LoginResponse = unwrapResponse(LoginResponseSchema, response.data)
  return data
}

export async function me() {
  const response = await http.get("/me")

  // Validar y extraer solo data
  const data: TenantContext = unwrapResponse(TenantContextSchema, response.data)
  return data
}
