-- Removes demo products inserted by 000005_seed_products.up.sql.

DELETE FROM product
WHERE company_id IN (
    SELECT id
    FROM company
    WHERE LOWER(TRIM(name)) IN (
        'globant',
        'mercado libre',
        'stripe',
        'shopify',
        'salesforce',
        'notion',
        'slack',
        'uber',
        'airbnb',
        'netflix'
    )
);
