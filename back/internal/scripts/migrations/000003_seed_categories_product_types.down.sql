-- Removes catalog data inserted by 000003_seed_categories_product_types.up.sql.
-- Run only after products using these categories/types were removed.

DELETE FROM product_type;
DELETE FROM category_product;
DELETE FROM category_company;
