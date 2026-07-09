-- Product image incremental

ALTER TABLE product
ADD COLUMN IF NOT EXISTS image_path TEXT;
