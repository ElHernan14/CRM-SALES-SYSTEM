export interface LoginRequest {
  email: string;
  password: string;
}

export interface LoginResponse {
  token: string;
}

export interface TenantContext {
  userID: number;
  email: string;
  companyID?: number;
  clientID?: number;
  roles: string[];
  permissions: string[];
}