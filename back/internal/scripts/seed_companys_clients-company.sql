-- NEXORA - SEED DE EMPRESAS, USUARIOS Y CLIENTES EMPRESARIALES
-- Requiere ejecutar antes:
-- 1) INIT.sql
-- 2) init_permisos_roles.sql
-- 3) nexora_seed_categories_and_product_types.sql
--
-- Password demo para todos los usuarios seed: Password123!
-- Script idempotente: evita duplicar users, companies, clients y roles.

BEGIN;

WITH seed(
    company_name,
    company_email,
    category_name,
    description,
    logo_path,
    first_name,
    last_name,
    phone
) AS (
    VALUES
        ('Globant', 'globant@mail.com', 'Tecnología y Software', 'Empresa tecnológica enfocada en transformación digital, software, inteligencia artificial y soluciones empresariales modernas.', 'https://logo.clearbit.com/globant.com', 'Globant', 'Client', '123123123123'),
        ('Mercado Libre', 'mercadolibre@mail.com', 'Comercio Electrónico', 'Plataforma líder de comercio electrónico, pagos digitales, logística y soluciones fintech en Latinoamérica.', 'https://logo.clearbit.com/mercadolibre.com', 'Mercado', 'Libre', '123123123123'),
        ('Stripe', 'stripe@mail.com', 'Fintech y Pagos', 'Empresa fintech que ofrece infraestructura de pagos online para negocios, startups y plataformas SaaS.', 'https://logo.clearbit.com/stripe.com', 'Stripe', 'Client', '123123123123'),
        ('Shopify', 'shopify@mail.com', 'Comercio Electrónico', 'Plataforma de comercio digital que permite crear tiendas online, gestionar productos, pagos y ventas.', 'https://logo.clearbit.com/shopify.com', 'Shopify', 'Client', '123123123123'),
        ('Salesforce', 'salesforce@mail.com', 'CRM y Ventas', 'Empresa de software CRM enfocada en ventas, marketing, atención al cliente y automatización empresarial.', 'https://logo.clearbit.com/salesforce.com', 'Salesforce', 'Client', '123123123123'),
        ('Notion', 'notion@mail.com', 'Productividad y Colaboración', 'Plataforma de productividad para organizar documentos, bases de datos, tareas y colaboración en equipos.', 'https://logo.clearbit.com/notion.so', 'Notion', 'Client', '123123123123'),
        ('Slack', 'slack@mail.com', 'Productividad y Colaboración', 'Herramienta moderna de comunicación empresarial para equipos, canales, integraciones y colaboración interna.', 'https://logo.clearbit.com/slack.com', 'Slack', 'Client', '123123123123'),
        ('Uber', 'uber@mail.com', 'Movilidad y Logística', 'Plataforma tecnológica de movilidad, entregas y servicios urbanos conectados mediante aplicaciones digitales.', 'https://logo.clearbit.com/uber.com', 'Uber', 'Client', '123123123123'),
        ('Airbnb', 'airbnb@mail.com', 'Turismo y Alojamiento', 'Plataforma global para reservas de alojamientos, experiencias y servicios turísticos digitales.', 'https://logo.clearbit.com/airbnb.com', 'Airbnb', 'Client', '123123123123'),
        ('Netflix', 'netflix@mail.com', 'Entretenimiento Digital', 'Empresa de entretenimiento digital enfocada en streaming, contenido audiovisual y suscripciones online.', 'https://logo.clearbit.com/netflix.com', 'Netflix', 'Client', '123123123123')
),
created_users AS (
    INSERT INTO users (email, password_hash, status)
    SELECT
        s.company_email,
        '$2a$10$ZeTFLBA1p25H8gN5kT9qHeO9bubqJ2Arilr5VQdOSp3pW3ZZ0.8GO',
        1
    FROM seed s
    WHERE NOT EXISTS (
        SELECT 1
        FROM users u
        WHERE LOWER(TRIM(u.email)) = LOWER(TRIM(s.company_email))
    )
    RETURNING id, email
),
all_users AS (
    SELECT u.id, u.email
    FROM users u
    JOIN seed s ON LOWER(TRIM(s.company_email)) = LOWER(TRIM(u.email))
    UNION
    SELECT cu.id, cu.email
    FROM created_users cu
),
created_companies AS (
    INSERT INTO company (
        name,
        category_id,
        category,
        description,
        logo_path,
        status
    )
    SELECT
        s.company_name,
        cc.id,
        cc.name,
        s.description,
        s.logo_path,
        1
    FROM seed s
    JOIN category_company cc
      ON LOWER(TRIM(cc.name)) = LOWER(TRIM(s.category_name))
    WHERE NOT EXISTS (
        SELECT 1
        FROM company c
        WHERE LOWER(TRIM(c.name)) = LOWER(TRIM(s.company_name))
          AND c.deleted_at IS NULL
    )
    RETURNING id, name
),
all_companies AS (
    SELECT c.id, c.name
    FROM company c
    JOIN seed s ON LOWER(TRIM(s.company_name)) = LOWER(TRIM(c.name))
    WHERE c.deleted_at IS NULL
    UNION
    SELECT cc.id, cc.name
    FROM created_companies cc
),
created_clients AS (
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
        u.id,
        c.id,
        s.first_name,
        s.last_name,
        s.company_email,
        s.phone,
        1
    FROM seed s
    JOIN all_users u ON LOWER(TRIM(u.email)) = LOWER(TRIM(s.company_email))
    JOIN all_companies c ON LOWER(TRIM(c.name)) = LOWER(TRIM(s.company_name))
    WHERE NOT EXISTS (
        SELECT 1
        FROM client cl
        WHERE LOWER(TRIM(cl.email)) = LOWER(TRIM(s.company_email))
          AND cl.deleted_at IS NULL
    )
    RETURNING id
)
INSERT INTO user_rol (user_id, role_id)
SELECT u.id, r.id
FROM all_users u
CROSS JOIN rol r
WHERE r.name = 'company_user'
ON CONFLICT DO NOTHING;

COMMIT;
