-- Temporary maintenance script.
-- Use this only when you need to recreate/reseed product_type rows and
-- reassign products afterwards.

ALTER TABLE product
DROP CONSTRAINT IF EXISTS fk_product_type;

ALTER TABLE product
ALTER COLUMN type_id DROP NOT NULL;

UPDATE product
SET type_id = NULL;

-- Optional cleanup if you want to fully reseed product types.
-- Uncomment only if you really want to remove all current product_type rows.
--
-- DELETE FROM product_type;

