import { z } from 'zod';

// LoginRequest
export const LoginRequestSchema = z.object({
  email: z.string().email(),
  password: z.string(),
});
export type LoginRequest = z.infer<typeof LoginRequestSchema>;

// LoginResponse (solo data.token)
export const LoginResponseSchema = z.object({
  token: z.string(),
});
export type LoginResponse = z.infer<typeof LoginResponseSchema>;

// TenantContext
export const TenantContextSchema = z.object({
  userID: z.number(),
  email: z.string().email(),
  company_id: z.number().nullable().optional(),
  client_id: z.number().nullable().optional(),
  roles: z.array(z.string()),
  permissions: z.array(z.string()),
});

export type TenantContext = z.infer<typeof TenantContextSchema>;
