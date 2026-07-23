import { z } from 'zod';

import { TenantContextSchema } from './auth.types';

export const PersonalRegisterRequestSchema = z.object({
  first_name: z
    .string()
    .trim()
    .min(2, 'First name must contain at least 2 characters')
    .max(100, 'First name is too long'),

  last_name: z
    .string()
    .trim()
    .min(2, 'Last name must contain at least 2 characters')
    .max(100, 'Last name is too long'),

  email: z.string().trim().email('Enter a valid email address'),

  password: z.string().min(8, 'Password must contain at least 8 characters'),

  phone: z.string().trim().max(50, 'Phone number is too long').optional(),
});

export type PersonalRegisterRequest = z.infer<typeof PersonalRegisterRequestSchema>;

export const BusinessRegisterRequestSchema = z.object({
  first_name: z.string().trim().min(2, 'First name must contain at least 2 characters').max(100),

  last_name: z.string().trim().min(2, 'Last name must contain at least 2 characters').max(100),

  email: z.string().trim().email('Enter a valid email address'),

  password: z.string().min(8, 'Password must contain at least 8 characters'),

  phone: z.string().trim().max(50).optional(),

  company: z.object({
    name: z.string().trim().min(2, 'Company name must contain at least 2 characters').max(150),

    email: z.string().trim().email('Enter a valid company email'),

    category_id: z.number().int().positive('Select a company category'),

    description: z.string().trim().max(1000, 'Description is too long').optional(),
  }),
});

export type BusinessRegisterRequest = z.infer<typeof BusinessRegisterRequestSchema>;

export const RegisterResponseSchema = z.object({
  token: z.string().min(1),
  user: TenantContextSchema,
});

export type RegisterResponse = z.infer<typeof RegisterResponseSchema>;
