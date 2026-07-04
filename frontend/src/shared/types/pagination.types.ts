import { z } from 'zod';

export const PaginationMetaSchema = z.object({
  page: z.number(),
  limit: z.number(),
  total: z.number(),
  total_pages: z.number().optional(),
});

export type PaginationMeta = z.infer<typeof PaginationMetaSchema>;
