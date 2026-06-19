import { http } from "@/shared/api/http";
import type {
  LoginRequest,
  LoginResponse,
  TenantContext,
} from "../types/auth.types";

export async function login(
  payload: LoginRequest
) {
  const { data } = await http.post<LoginResponse>(
    "/auth/login",
    payload
  );

  return data;
}

export async function me() {
  const { data } =
    await http.get<TenantContext>(
      "/auth/me"
    );

  return data;
}