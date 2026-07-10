-- Company cover image incremental

ALTER TABLE company
ADD COLUMN IF NOT EXISTS cover_image_path TEXT;
