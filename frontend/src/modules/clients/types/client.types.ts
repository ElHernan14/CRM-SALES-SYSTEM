import { z } from 'zod';
import { PaginationMetaSchema } from '@/shared/types/pagination.types';

export const ClientResponseSchema = z.object({
  id: z.number(),
  first_name: z.string(),
  last_name: z.string(),
  email: z.string().email(),
  company_id: z.number().nullable().optional(),
  company_name: z.string().nullable().optional(),
  phone: z.string().optional(),
  status: z.number().optional(),
  deleted_at: z.string().nullable().optional(),
});

export type ClientResponse = z.infer<typeof ClientResponseSchema>;

export const GetClientsRequestSchema = z.object({
  company_id: z.number().optional(),
  search: z.string().optional(),
  email: z.string().optional(),
  page: z.number(),
  limit: z.number(),
});

export type GetClientsRequest = z.infer<typeof GetClientsRequestSchema>;

export const GetClientsResponseSchema = z.object({
  clients: z.array(ClientResponseSchema),
  meta: PaginationMetaSchema,
});

export type GetClientsResponse = z.infer<typeof GetClientsResponseSchema>;
