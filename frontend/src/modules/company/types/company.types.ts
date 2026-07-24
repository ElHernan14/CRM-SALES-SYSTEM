import { z } from 'zod';

export const CompanyMeSchema = z.object({
  id: z.number(),

  name: z.string(),

  category_id: z.number(),

  category: z.string(),

  description: z.string().nullable().optional(),

  logo: z.string().nullable().optional(),

  cover_image: z.string().nullable().optional(),

  created_at: z.string().nullable().optional(),

  deleted_at: z.string().nullable().optional(),

  status: z.number(),
});

export type CompanyMe = z.infer<typeof CompanyMeSchema>;

export const UpdateCompanyMeRequestSchema = z.object({
  name: z.string().trim().min(2).max(150).optional(),

  category_id: z.number().int().positive().optional(),

  description: z.string().trim().max(1000).optional(),
});

export type UpdateCompanyMeRequest = z.infer<typeof UpdateCompanyMeRequestSchema>;

export const UploadCompanyLogoResponseSchema = z.object({
  logo: z.string(),
});

export type UploadCompanyLogoResponse = z.infer<typeof UploadCompanyLogoResponseSchema>;

export const UploadCompanyCoverResponseSchema = z.object({
  cover_image: z.string(),
});

export type UploadCompanyCoverResponse = z.infer<typeof UploadCompanyCoverResponseSchema>;
