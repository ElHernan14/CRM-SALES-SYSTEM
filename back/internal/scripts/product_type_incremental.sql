CREATE TABLE IF NOT EXISTS product_type (
    id SERIAL PRIMARY KEY,
    category_id INT NOT NULL,
    name VARCHAR(80) NOT NULL,
    description TEXT,

    CONSTRAINT fk_product_type_category
        FOREIGN KEY (category_id)
        REFERENCES category_product(id)
        ON DELETE CASCADE,

    CONSTRAINT uq_product_type_category_name UNIQUE (category_id, name)
);

INSERT INTO product_type (category_id, name, description)
SELECT id, 'General', NULL
FROM category_product
ON CONFLICT (category_id, name) DO NOTHING;

ALTER TABLE product
ADD COLUMN IF NOT EXISTS kind VARCHAR(20);

UPDATE product
SET kind = COALESCE(NULLIF(TRIM(type), ''), 'product')
WHERE kind IS NULL;

ALTER TABLE product
ADD COLUMN IF NOT EXISTS type_id INT;

UPDATE product p
SET type_id = pt.id
FROM product_type pt
WHERE p.type_id IS NULL
  AND pt.category_id = p.category_id
  AND pt.name = 'General';

ALTER TABLE product
ALTER COLUMN kind SET NOT NULL;

ALTER TABLE product
ALTER COLUMN type_id SET NOT NULL;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_constraint
        WHERE conname = 'fk_product_type'
    ) THEN
        ALTER TABLE product
        ADD CONSTRAINT fk_product_type
        FOREIGN KEY (type_id)
        REFERENCES product_type(id);
    END IF;
END $$;

CREATE INDEX IF NOT EXISTS idx_product_type_id ON product(type_id);
CREATE INDEX IF NOT EXISTS idx_product_kind ON product(kind);
