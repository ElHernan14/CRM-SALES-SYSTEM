-- Removes demo companies, their associated clients and users.
-- Run only after products/invoices for these companies were removed.

WITH seed(email, company_name) AS (
    VALUES
        ('globant@mail.com', 'Globant'),
        ('mercadolibre@mail.com', 'Mercado Libre'),
        ('stripe@mail.com', 'Stripe'),
        ('shopify@mail.com', 'Shopify'),
        ('salesforce@mail.com', 'Salesforce'),
        ('notion@mail.com', 'Notion'),
        ('slack@mail.com', 'Slack'),
        ('uber@mail.com', 'Uber'),
        ('airbnb@mail.com', 'Airbnb'),
        ('netflix@mail.com', 'Netflix')
),
seed_users AS (
    SELECT u.id
    FROM users u
    JOIN seed s ON LOWER(TRIM(s.email)) = LOWER(TRIM(u.email))
),
seed_companies AS (
    SELECT c.id
    FROM company c
    JOIN seed s ON LOWER(TRIM(s.company_name)) = LOWER(TRIM(c.name))
)
DELETE FROM user_rol
WHERE user_id IN (SELECT id FROM seed_users);

WITH seed(email, company_name) AS (
    VALUES
        ('globant@mail.com', 'Globant'),
        ('mercadolibre@mail.com', 'Mercado Libre'),
        ('stripe@mail.com', 'Stripe'),
        ('shopify@mail.com', 'Shopify'),
        ('salesforce@mail.com', 'Salesforce'),
        ('notion@mail.com', 'Notion'),
        ('slack@mail.com', 'Slack'),
        ('uber@mail.com', 'Uber'),
        ('airbnb@mail.com', 'Airbnb'),
        ('netflix@mail.com', 'Netflix')
)
DELETE FROM client
WHERE LOWER(TRIM(email)) IN (
    SELECT LOWER(TRIM(email)) FROM seed
);

WITH seed(email, company_name) AS (
    VALUES
        ('globant@mail.com', 'Globant'),
        ('mercadolibre@mail.com', 'Mercado Libre'),
        ('stripe@mail.com', 'Stripe'),
        ('shopify@mail.com', 'Shopify'),
        ('salesforce@mail.com', 'Salesforce'),
        ('notion@mail.com', 'Notion'),
        ('slack@mail.com', 'Slack'),
        ('uber@mail.com', 'Uber'),
        ('airbnb@mail.com', 'Airbnb'),
        ('netflix@mail.com', 'Netflix')
)
DELETE FROM company
WHERE LOWER(TRIM(name)) IN (
    SELECT LOWER(TRIM(company_name)) FROM seed
);

WITH seed(email, company_name) AS (
    VALUES
        ('globant@mail.com', 'Globant'),
        ('mercadolibre@mail.com', 'Mercado Libre'),
        ('stripe@mail.com', 'Stripe'),
        ('shopify@mail.com', 'Shopify'),
        ('salesforce@mail.com', 'Salesforce'),
        ('notion@mail.com', 'Notion'),
        ('slack@mail.com', 'Slack'),
        ('uber@mail.com', 'Uber'),
        ('airbnb@mail.com', 'Airbnb'),
        ('netflix@mail.com', 'Netflix')
)
DELETE FROM users
WHERE LOWER(TRIM(email)) IN (
    SELECT LOWER(TRIM(email)) FROM seed
);
