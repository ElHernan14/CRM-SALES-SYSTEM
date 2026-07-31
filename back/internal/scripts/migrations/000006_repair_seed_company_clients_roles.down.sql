-- Rollback for 000006 repair only.
-- It removes the company_user role relation and clients for the demo seed users.
-- It intentionally keeps users and companies because those were created by 000004.

BEGIN;

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
),
seed_users AS (
    SELECT u.id
    FROM users u
    JOIN seed s
      ON LOWER(TRIM(u.email)) = LOWER(TRIM(s.company_email))
)
DELETE FROM user_rol
WHERE user_id IN (SELECT id FROM seed_users)
  AND role_id = (SELECT id FROM rol WHERE name = 'company_user');

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
DELETE FROM client
WHERE LOWER(TRIM(email)) IN (
    SELECT LOWER(TRIM(company_email))
    FROM seed
);

COMMIT;
