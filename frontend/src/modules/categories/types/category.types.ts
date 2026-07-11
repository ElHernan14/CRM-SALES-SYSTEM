import { z } from 'zod';

export const CategoryItemSchema = z.object({
  id: z.number(),
  name: z.string(),
  description: z
    .string()
    .nullable()
    .optional()
    .transform((value) => value ?? ''),
});

export type CategoryItem = z.infer<typeof CategoryItemSchema>;

export const CategoryProductTypeSchema = z.object({
  id: z.number(),
  category_id: z.number(),
  name: z.string(),
  description: z
    .string()
    .nullable()
    .optional()
    .transform((value) => value ?? ''),
});

export type CategoryProductType = z.infer<typeof CategoryProductTypeSchema>;

export const GetCategoriesResponseSchema = z.object({
  product_categories: z.array(CategoryItemSchema),
  company_categories: z.array(CategoryItemSchema),
  product_types: z.array(CategoryProductTypeSchema),
});

export type GetCategoriesResponse = z.infer<typeof GetCategoriesResponseSchema>;
