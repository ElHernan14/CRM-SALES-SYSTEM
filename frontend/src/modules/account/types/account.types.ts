import { z } from 'zod';

export const ClientProfileSchema = z.object({
  id: z.number(),

  first_name: z.string(),
  last_name: z.string(),

  email: z.string().email(),

  phone: z.string().nullable().optional(),

  company_id: z.number().nullable().optional(),

  company_name: z.string().nullable().optional(),

  status: z.number(),

  deleted_at: z.string().nullable().optional(),
});

export type ClientProfile = z.infer<typeof ClientProfileSchema>;

export const UpdateClientProfileRequestSchema = z.object({
  first_name: z
    .string()
    .trim()
    .min(2, 'First name must contain at least 2 characters')
    .max(100)
    .optional(),

  last_name: z
    .string()
    .trim()
    .min(2, 'Last name must contain at least 2 characters')
    .max(100)
    .optional(),

  email: z.string().trim().email('Enter a valid email address').optional(),

  phone: z.string().trim().max(50, 'Phone number is too long').optional(),
});

export type UpdateClientProfileRequest = z.infer<typeof UpdateClientProfileRequestSchema>;
