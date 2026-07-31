-- Repairs the demo company seed when users/companies already exist but
-- their associated clients and company_user roles were not created.
--
-- Safe to run after:
-- 000001_init.up.sql
-- 000002_seed_roles_permissions.up.sql
-- 000003_seed_categories_product_types.up.sql
-- 000004_seed_companies_clients.up.sql

BEGIN;

WITH seed(
    company_name,
    company_email,
    first_name,
    last_name,
    phone
) AS (
    VALUES
        ('Globant', 'globant@mail.com', 'Globant', 'Client', '123123123123'),
        ('Mercado Libre', 'mercadolibre@mail.com', 'Mercado', 'Libre', '123123123123'),
        ('Stripe', 'stripe@mail.com', 'Stripe', 'Client', '123123123123'),
        ('Shopify', 'shopify@mail.com', 'Shopify', 'Client', '123123123123'),
        ('Salesforce', 'salesforce@mail.com', 'Salesforce', 'Client', '123123123123'),
        ('Notion', 'notion@mail.com', 'Notion', 'Client', '123123123123'),
        ('Slack', 'slack@mail.com', 'Slack', 'Client', '123123123123'),
        ('Uber', 'uber@mail.com', 'Uber', 'Client', '123123123123'),
        ('Airbnb', 'airbnb@mail.com', 'Airbnb', 'Client', '123123123123'),
        ('Netflix', 'netflix@mail.com', 'Netflix', 'Client', '123123123123')
),
resolved AS (
    SELECT
        u.id AS user_id,
        c.id AS company_id,
        s.company_email,
        s.first_name,
        s.last_name,
        s.phone
    FROM seed s
    JOIN users u
      ON LOWER(TRIM(u.email)) = LOWER(TRIM(s.company_email))
    JOIN company c
      ON LOWER(TRIM(c.name)) = LOWER(TRIM(s.company_name))
    WHERE u.deleted_at IS NULL
      AND COALESCE(u.status, 1) = 1
      AND c.deleted_at IS NULL
      AND COALESCE(c.status, 1) = 1
),
updated_clients AS (
    UPDATE client cl
    SET
        user_id = COALESCE(cl.user_id, r.user_id),
        company_id = COALESCE(cl.company_id, r.company_id),
        first_name = COALESCE(NULLIF(TRIM(cl.first_name), ''), r.first_name),
        last_name = COALESCE(NULLIF(TRIM(cl.last_name), ''), r.last_name),
        phone = COALESCE(NULLIF(TRIM(cl.phone), ''), r.phone),
        status = 1,
        deleted_at = NULL
    FROM resolved r
    WHERE LOWER(TRIM(cl.email)) = LOWER(TRIM(r.company_email))
    RETURNING cl.id
)
INSERT INTO client (
    user_id,
    company_id,
    first_name,
    last_name,
    email,
    phone,
    status
)
SELECT
    r.user_id,
    r.company_id,
    r.first_name,
    r.last_name,
    r.company_email,
    r.phone,
    1
FROM resolved r
WHERE NOT EXISTS (
    SELECT 1
    FROM client cl
    WHERE (
        LOWER(TRIM(cl.email)) = LOWER(TRIM(r.company_email))
        OR cl.user_id = r.user_id
    )
      AND cl.deleted_at IS NULL
);

WITH seed(company_email) AS (
    VALUES
        ('globant@mail.com'),
        ('mercadolibre@mail.com'),
        ('stripe@mail.com'),
        ('shopify@mail.com'),
        ('salesforce@mail.com'),
        ('notion@mail.com'),
        ('slack@mail.com'),
        ('uber@mail.com'),
        ('airbnb@mail.com'),
        ('netflix@mail.com')
)
INSERT INTO user_rol (user_id, role_id)
SELECT u.id, r.id
FROM seed s
JOIN users u
  ON LOWER(TRIM(u.email)) = LOWER(TRIM(s.company_email))
JOIN rol r
  ON r.name = 'company_user'
WHERE u.deleted_at IS NULL
  AND COALESCE(u.status, 1) = 1
ON CONFLICT DO NOTHING;

COMMIT;
