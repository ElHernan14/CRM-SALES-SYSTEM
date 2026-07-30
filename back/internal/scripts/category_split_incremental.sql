CREATE TABLE IF NOT EXISTS category_company (
    id SERIAL PRIMARY KEY,
    name VARCHAR(80) NOT NULL UNIQUE,
    description TEXT
);

CREATE TABLE IF NOT EXISTS category_product (
    id SERIAL PRIMARY KEY,
    name VARCHAR(80) NOT NULL UNIQUE,
    description TEXT
);

ALTER TABLE company
ADD COLUMN IF NOT EXISTS category VARCHAR(80) NOT NULL DEFAULT 'general';

ALTER TABLE product
ADD COLUMN IF NOT EXISTS category VARCHAR(80) NOT NULL DEFAULT 'general';

INSERT INTO category_company (name, description)
SELECT DISTINCT COALESCE(NULLIF(TRIM(category), ''), 'general'), NULL
FROM company
WHERE category IS NOT NULL
ON CONFLICT (name) DO NOTHING;

INSERT INTO category_company (name, description)
VALUES ('general', NULL)
ON CONFLICT (name) DO NOTHING;

INSERT INTO category_product (name, description)
SELECT DISTINCT COALESCE(NULLIF(TRIM(category), ''), 'general'), NULL
FROM product
WHERE category IS NOT NULL
ON CONFLICT (name) DO NOTHING;

INSERT INTO category_product (name, description)
VALUES ('general', NULL)
ON CONFLICT (name) DO NOTHING;

DO $$
BEGIN
    IF to_regclass('public.category') IS NOT NULL THEN
        EXECUTE 'INSERT INTO category_product (name, description) SELECT name, description FROM category ON CONFLICT (name) DO NOTHING';
    END IF;
END $$;

ALTER TABLE company
ADD COLUMN IF NOT EXISTS category_id INT;

ALTER TABLE product
ADD COLUMN IF NOT EXISTS category_id INT;

UPDATE company c
SET category_id = cc.id
FROM category_company cc
WHERE c.category_id IS NULL
  AND LOWER(cc.name) = LOWER(COALESCE(NULLIF(TRIM(c.category), ''), 'general'));

UPDATE company
SET category_id = (SELECT id FROM category_company WHERE name = 'general')
WHERE category_id IS NULL;

DO $$
BEGIN
    IF to_regclass('public.category') IS NOT NULL THEN
        EXECUTE '
            UPDATE product p
            SET category_id = cp.id
            FROM category old_cat
            INNER JOIN category_product cp ON LOWER(cp.name) = LOWER(old_cat.name)
            WHERE p.category_id = old_cat.id
        ';
    END IF;
END $$;

UPDATE product p
SET category_id = cp.id
FROM category_product cp
WHERE p.category_id IS NULL
  AND LOWER(cp.name) = LOWER(COALESCE(NULLIF(TRIM(p.category), ''), 'general'));

UPDATE product
SET category_id = (SELECT id FROM category_product WHERE name = 'general')
WHERE category_id IS NULL;

ALTER TABLE company
ALTER COLUMN category_id SET NOT NULL;

ALTER TABLE product
ALTER COLUMN category_id SET NOT NULL;

ALTER TABLE product
DROP CONSTRAINT IF EXISTS fk_product_category;

ALTER TABLE company
DROP CONSTRAINT IF EXISTS fk_company_category;

ALTER TABLE product
ADD CONSTRAINT fk_product_category
FOREIGN KEY (category_id)
REFERENCES category_product(id);

ALTER TABLE company
ADD CONSTRAINT fk_company_category
FOREIGN KEY (category_id)
REFERENCES category_company(id);

DROP INDEX IF EXISTS idx_product_category;
DROP INDEX IF EXISTS idx_company_category;

CREATE INDEX IF NOT EXISTS idx_product_category_id ON product(category_id);
CREATE INDEX IF NOT EXISTS idx_company_category_id ON company(category_id);

INSERT INTO category_company (
    name,
    description
) VALUES (
    'Tecnología e Ingeniería',
    'Empresas dedicadas al desarrollo tecnológico, ingeniería, innovación y soluciones industriales.'
)
ON CONFLICT (name) DO NOTHING;
