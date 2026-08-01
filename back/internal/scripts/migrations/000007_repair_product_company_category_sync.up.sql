-- Repairs denormalized category/type varchar fields and active soft-delete state.
-- Safe to run after the category split and product type migrations.

BEGIN;

UPDATE product p
SET
    category = cp.name,
    type = pt.name,
    deleted_at = CASE
        WHEN p.status = 1 THEN NULL
        ELSE p.deleted_at
    END
FROM category_product cp
JOIN product_type pt ON pt.category_id = cp.id
WHERE p.category_id = cp.id
  AND p.type_id = pt.id
  AND (
      p.category IS DISTINCT FROM cp.name
      OR p.type IS DISTINCT FROM pt.name
      OR (p.status = 1 AND p.deleted_at IS NOT NULL)
  );

UPDATE company c
SET category = cc.name
FROM category_company cc
WHERE c.category_id = cc.id
  AND c.category IS DISTINCT FROM cc.name;

COMMIT;
