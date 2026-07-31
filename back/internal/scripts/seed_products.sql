-- NEXORA - SEED DE PRODUCTOS DEMO
-- Requiere ejecutar antes:
-- 1) INIT.sql
-- 2) nexora_seed_categories_and_product_types.sql
-- 3) seed_companys_clients-company.sql
--
-- Script idempotente: evita duplicar productos por empresa + nombre activo.

BEGIN;

WITH seed(
    company_name,
    name,
    description,
    kind,
    price,
    stock,
    category_name,
    type_name
) AS (
    VALUES
        -- Globant
        ('Globant', 'Staff Augmentation', 'Desarrollo de equipos dedicados', 'service', 12000, 9999, 'Desarrollo de Software', 'Equipos Dedicados'),
        ('Globant', 'Software Factory', 'Desarrollo de software empresarial', 'service', 25000, 9999, 'Desarrollo de Software', 'Software Empresarial'),
        ('Globant', 'Cloud Migration', 'Migración de infraestructura cloud', 'service', 18000, 9999, 'Cloud y DevOps', 'Migración Cloud'),
        ('Globant', 'UX/UI Design', 'Diseño de experiencias digitales', 'service', 4500, 9999, 'Desarrollo de Software', 'Diseño UX/UI'),
        ('Globant', 'QA Automation', 'Automatización de pruebas', 'service', 6200, 9999, 'IA y Automatización', 'Automatización QA'),
        ('Globant', 'DevOps Consulting', 'Consultoría DevOps', 'service', 9000, 9999, 'Cloud y DevOps', 'Consultoría DevOps'),
        ('Globant', 'AI Solutions', 'Soluciones de Inteligencia Artificial', 'service', 22000, 9999, 'IA y Automatización', 'Inteligencia Artificial'),
        ('Globant', 'Mobile Development', 'Desarrollo Mobile', 'service', 10000, 9999, 'Desarrollo de Software', 'Desarrollo Mobile'),
        ('Globant', 'Web Development', 'Desarrollo Web', 'service', 8500, 9999, 'Desarrollo de Software', 'Desarrollo Web'),
        ('Globant', 'IT Consulting', 'Consultoría Tecnológica', 'service', 7000, 9999, 'Desarrollo de Software', 'Consultoría Tecnológica'),

        -- Mercado Libre
        ('Mercado Libre', 'Mercado Pago', 'Billetera virtual', 'service', 0, 9999, 'Pagos y Fintech', 'Transferencias Digitales'),
        ('Mercado Libre', 'Mercado Envíos', 'Servicio logístico', 'service', 12.50, 9999, 'Movilidad y Delivery', 'Logística Empresarial'),
        ('Mercado Libre', 'Mercado Shops', 'Tienda online', 'service', 39.99, 9999, 'Comercio Electrónico', 'Tienda Online'),
        ('Mercado Libre', 'Mercado Ads', 'Publicidad digital', 'service', 120, 9999, 'Comercio Electrónico', 'Publicidad Digital'),
        ('Mercado Libre', 'Mercado Crédito', 'Financiación', 'service', 0, 9999, 'Pagos y Fintech', 'Crédito y Financiación'),
        ('Mercado Libre', 'Mercado Point', 'Lector POS', 'product', 79.99, 250, 'Pagos y Fintech', 'Terminal de Pago'),
        ('Mercado Libre', 'Mercado Full', 'Logística Full', 'service', 29.99, 9999, 'Comercio Electrónico', 'Fulfillment'),
        ('Mercado Libre', 'Mercado Libre Premium', 'Suscripción Premium', 'service', 14.99, 9999, 'Comercio Electrónico', 'Marketplace'),
        ('Mercado Libre', 'Mercado Pago QR', 'Cobro QR', 'service', 0, 9999, 'Pagos y Fintech', 'Pasarela de Pagos'),
        ('Mercado Libre', 'Mercado API', 'API para desarrolladores', 'service', 49.99, 9999, 'Desarrollo de Software', 'API e Integraciones'),

        -- Stripe
        ('Stripe', 'Payments API', 'API de pagos', 'service', 29.99, 9999, 'Pagos y Fintech', 'Pasarela de Pagos'),
        ('Stripe', 'Stripe Billing', 'Facturación recurrente', 'service', 25, 9999, 'Pagos y Fintech', 'Facturación Recurrente'),
        ('Stripe', 'Stripe Checkout', 'Checkout Online', 'service', 15, 9999, 'Comercio Electrónico', 'Checkout'),
        ('Stripe', 'Stripe Connect', 'Marketplace Payments', 'service', 40, 9999, 'Pagos y Fintech', 'Pasarela de Pagos'),
        ('Stripe', 'Stripe Atlas', 'Creación de empresas', 'service', 500, 9999, 'Pagos y Fintech', 'Gestión Fiscal'),
        ('Stripe', 'Stripe Terminal', 'POS físico', 'product', 299, 40, 'Pagos y Fintech', 'Terminal de Pago'),
        ('Stripe', 'Stripe Radar', 'Prevención de fraude', 'service', 18, 9999, 'Pagos y Fintech', 'Prevención de Fraude'),
        ('Stripe', 'Stripe Identity', 'Validación de identidad', 'service', 10, 9999, 'Pagos y Fintech', 'Prevención de Fraude'),
        ('Stripe', 'Stripe Issuing', 'Tarjetas virtuales', 'service', 19, 9999, 'Pagos y Fintech', 'Tarjetas y Emisión'),
        ('Stripe', 'Stripe Tax', 'Cálculo de impuestos', 'service', 12, 9999, 'Pagos y Fintech', 'Gestión Fiscal'),

        -- Shopify
        ('Shopify', 'Basic Shopify', 'Plan básico', 'service', 39, 9999, 'Comercio Electrónico', 'Tienda Online'),
        ('Shopify', 'Shopify Plan', 'Plan profesional', 'service', 105, 9999, 'Comercio Electrónico', 'Tienda Online'),
        ('Shopify', 'Advanced Shopify', 'Plan avanzado', 'service', 399, 9999, 'Comercio Electrónico', 'Tienda Online'),
        ('Shopify', 'POS Lite', 'Sistema POS', 'service', 0, 9999, 'Comercio Electrónico', 'Checkout'),
        ('Shopify', 'POS Pro', 'Sistema POS Pro', 'service', 89, 9999, 'Comercio Electrónico', 'Checkout'),
        ('Shopify', 'Themes Premium', 'Tema premium', 'product', 180, 100, 'Comercio Electrónico', 'Gestión de Catálogo'),
        ('Shopify', 'Shopify Email', 'Email Marketing', 'service', 10, 9999, 'CRM y Marketing', 'Email Marketing'),
        ('Shopify', 'Shopify Flow', 'Automatizaciones', 'service', 25, 9999, 'Productividad y Colaboración', 'Automatización de Tareas'),
        ('Shopify', 'Shopify Markets', 'Ventas internacionales', 'service', 50, 9999, 'Comercio Electrónico', 'Marketplace'),
        ('Shopify', 'Hydrogen', 'Framework Headless', 'service', 0, 9999, 'Desarrollo de Software', 'Desarrollo Web'),

        -- Salesforce
        ('Salesforce', 'Sales Cloud', 'CRM Comercial', 'service', 75, 9999, 'CRM y Marketing', 'CRM de Ventas'),
        ('Salesforce', 'Service Cloud', 'Atención al Cliente', 'service', 80, 9999, 'CRM y Marketing', 'Atención al Cliente'),
        ('Salesforce', 'Marketing Cloud', 'Marketing Digital', 'service', 120, 9999, 'CRM y Marketing', 'Automatización de Marketing'),
        ('Salesforce', 'Commerce Cloud', 'E-commerce', 'service', 150, 9999, 'Comercio Electrónico', 'Tienda Online'),
        ('Salesforce', 'Tableau', 'Business Intelligence', 'service', 70, 9999, 'CRM y Marketing', 'Analítica Comercial'),
        ('Salesforce', 'Slack Enterprise', 'Colaboración', 'service', 25, 9999, 'Productividad y Colaboración', 'Mensajería Empresarial'),
        ('Salesforce', 'Einstein AI', 'IA Empresarial', 'service', 95, 9999, 'IA y Automatización', 'Inteligencia Artificial'),
        ('Salesforce', 'Experience Cloud', 'Portales', 'service', 110, 9999, 'CRM y Marketing', 'Customer Data Platform'),
        ('Salesforce', 'Platform Plus', 'Low Code', 'service', 100, 9999, 'Desarrollo de Software', 'Plataforma Low-Code'),
        ('Salesforce', 'Data Cloud', 'Gestión de Datos', 'service', 130, 9999, 'CRM y Marketing', 'Customer Data Platform'),

        -- Notion
        ('Notion', 'Notion Free', 'Plan gratuito', 'service', 0, 9999, 'Productividad y Colaboración', 'Espacio de Trabajo Digital'),
        ('Notion', 'Notion Plus', 'Plan Plus', 'service', 10, 9999, 'Productividad y Colaboración', 'Espacio de Trabajo Digital'),
        ('Notion', 'Notion Business', 'Plan Business', 'service', 18, 9999, 'Productividad y Colaboración', 'Espacio de Trabajo Digital'),
        ('Notion', 'Notion AI', 'Asistente IA', 'service', 8, 9999, 'IA y Automatización', 'Asistente Inteligente'),
        ('Notion', 'Notion Calendar', 'Calendario', 'service', 0, 9999, 'Productividad y Colaboración', 'Calendario y Agenda'),
        ('Notion', 'Enterprise', 'Plan Enterprise', 'service', 30, 9999, 'Productividad y Colaboración', 'Espacio de Trabajo Digital'),
        ('Notion', 'Forms', 'Formularios', 'service', 5, 9999, 'Productividad y Colaboración', 'Documentación Colaborativa'),
        ('Notion', 'Docs', 'Documentación', 'service', 7, 9999, 'Productividad y Colaboración', 'Documentación Colaborativa'),
        ('Notion', 'Wiki', 'Wiki Empresarial', 'service', 12, 9999, 'Productividad y Colaboración', 'Gestión del Conocimiento'),
        ('Notion', 'Projects', 'Gestión de Proyectos', 'service', 15, 9999, 'Productividad y Colaboración', 'Gestión de Proyectos'),

        -- Slack
        ('Slack', 'Slack Free', 'Plan gratuito', 'service', 0, 9999, 'Productividad y Colaboración', 'Mensajería Empresarial'),
        ('Slack', 'Slack Pro', 'Plan Pro', 'service', 8.75, 9999, 'Productividad y Colaboración', 'Mensajería Empresarial'),
        ('Slack', 'Slack Business+', 'Plan Business', 'service', 15, 9999, 'Productividad y Colaboración', 'Mensajería Empresarial'),
        ('Slack', 'Slack Enterprise', 'Plan Enterprise', 'service', 25, 9999, 'Productividad y Colaboración', 'Mensajería Empresarial'),
        ('Slack', 'Slack AI', 'IA para Slack', 'service', 10, 9999, 'IA y Automatización', 'Asistente Inteligente'),
        ('Slack', 'Huddles', 'Videollamadas', 'service', 0, 9999, 'Productividad y Colaboración', 'Videoconferencia'),
        ('Slack', 'Canvas', 'Documentación', 'service', 5, 9999, 'Productividad y Colaboración', 'Documentación Colaborativa'),
        ('Slack', 'Workflow Builder', 'Automatización', 'service', 12, 9999, 'Productividad y Colaboración', 'Automatización de Tareas'),
        ('Slack', 'Slack Connect', 'Conexión entre empresas', 'service', 20, 9999, 'Productividad y Colaboración', 'Mensajería Empresarial'),
        ('Slack', 'API Premium', 'Integraciones', 'service', 30, 9999, 'Desarrollo de Software', 'API e Integraciones'),

        -- Uber
        ('Uber', 'UberX', 'Viaje', 'service', 5, 9999, 'Movilidad y Delivery', 'Transporte Urbano'),
        ('Uber', 'Uber Comfort', 'Viaje Comfort', 'service', 8, 9999, 'Movilidad y Delivery', 'Transporte Urbano'),
        ('Uber', 'Uber Black', 'Viaje Premium', 'service', 15, 9999, 'Movilidad y Delivery', 'Transporte Urbano'),
        ('Uber', 'Uber Moto', 'Viaje Moto', 'service', 3, 9999, 'Movilidad y Delivery', 'Transporte Urbano'),
        ('Uber', 'Uber Reserve', 'Reserva anticipada', 'service', 20, 9999, 'Movilidad y Delivery', 'Transporte Programado'),
        ('Uber', 'Uber Eats', 'Delivery', 'service', 6, 9999, 'Movilidad y Delivery', 'Entrega de Última Milla'),
        ('Uber', 'Uber One', 'Suscripción', 'service', 9.99, 9999, 'Movilidad y Delivery', 'Transporte Urbano'),
        ('Uber', 'Uber Package', 'Envíos', 'service', 7, 9999, 'Movilidad y Delivery', 'Distribución de Paquetes'),
        ('Uber', 'Uber Business', 'Empresas', 'service', 30, 9999, 'Movilidad y Delivery', 'Logística Empresarial'),
        ('Uber', 'Uber Shuttle', 'Transporte Compartido', 'service', 4, 9999, 'Movilidad y Delivery', 'Transporte Programado'),

        -- Airbnb
        ('Airbnb', 'Apartamento', 'Alojamiento', 'service', 80, 9999, 'Turismo y Alojamiento', 'Reserva de Alojamiento'),
        ('Airbnb', 'Casa Completa', 'Alojamiento', 'service', 150, 9999, 'Turismo y Alojamiento', 'Reserva de Alojamiento'),
        ('Airbnb', 'Habitación Privada', 'Alojamiento', 'service', 45, 9999, 'Turismo y Alojamiento', 'Reserva de Alojamiento'),
        ('Airbnb', 'Experiencias', 'Experiencias Locales', 'service', 35, 9999, 'Turismo y Alojamiento', 'Experiencias Turísticas'),
        ('Airbnb', 'AirCover', 'Seguro', 'service', 0, 9999, 'Turismo y Alojamiento', 'Servicios para Huéspedes'),
        ('Airbnb', 'Luxury Stay', 'Lujo', 'service', 350, 9999, 'Turismo y Alojamiento', 'Reserva de Alojamiento'),
        ('Airbnb', 'Business Travel', 'Viajes Corporativos', 'service', 120, 9999, 'Turismo y Alojamiento', 'Viajes Corporativos'),
        ('Airbnb', 'Monthly Stay', 'Alquiler Mensual', 'service', 900, 9999, 'Turismo y Alojamiento', 'Alquiler Temporal'),
        ('Airbnb', 'Cabins', 'Cabañas', 'service', 110, 9999, 'Turismo y Alojamiento', 'Alquiler Temporal'),
        ('Airbnb', 'Beach House', 'Casa de Playa', 'service', 250, 9999, 'Turismo y Alojamiento', 'Alquiler Temporal'),

        -- Netflix
        ('Netflix', 'Plan Estándar', 'Streaming', 'service', 12.99, 9999, 'Streaming y Entretenimiento', 'Suscripción Digital'),
        ('Netflix', 'Plan Premium', 'Streaming', 'service', 19.99, 9999, 'Streaming y Entretenimiento', 'Suscripción Digital'),
        ('Netflix', 'Plan con Anuncios', 'Streaming', 'service', 6.99, 9999, 'Streaming y Entretenimiento', 'Suscripción Digital'),
        ('Netflix', 'Netflix Games', 'Juegos', 'service', 0, 9999, 'Streaming y Entretenimiento', 'Entretenimiento Interactivo'),
        ('Netflix', 'Gift Card', 'Tarjeta Regalo', 'product', 50, 500, 'Streaming y Entretenimiento', 'Suscripción Digital'),
        ('Netflix', 'Ultra HD', 'Complemento', 'service', 5, 9999, 'Streaming y Entretenimiento', 'Video bajo Demanda'),
        ('Netflix', 'Family Plan', 'Plan Familiar', 'service', 24.99, 9999, 'Streaming y Entretenimiento', 'Suscripción Digital'),
        ('Netflix', 'Mobile Plan', 'Plan Mobile', 'service', 4.99, 9999, 'Streaming y Entretenimiento', 'Suscripción Digital'),
        ('Netflix', 'Netflix Originals', 'Contenido Original', 'service', 0, 9999, 'Streaming y Entretenimiento', 'Producción Original'),
        ('Netflix', 'Kids Profile', 'Perfil Infantil', 'service', 0, 9999, 'Streaming y Entretenimiento', 'Contenido Infantil')
)
INSERT INTO product (
    name,
    description,
    kind,
    category_id,
    category,
    type_id,
    type,
    price,
    stock,
    reserved_stock,
    company_id,
    status
)
SELECT
    s.name,
    s.description,
    s.kind,
    cp.id,
    cp.name,
    pt.id,
    pt.name,
    s.price,
    s.stock,
    0,
    c.id,
    1
FROM seed s
JOIN company c
  ON LOWER(TRIM(c.name)) = LOWER(TRIM(s.company_name))
JOIN category_product cp
  ON LOWER(TRIM(cp.name)) = LOWER(TRIM(s.category_name))
JOIN product_type pt
  ON pt.category_id = cp.id
 AND LOWER(TRIM(pt.name)) = LOWER(TRIM(s.type_name))
WHERE c.deleted_at IS NULL
  AND NOT EXISTS (
      SELECT 1
      FROM product p
      WHERE p.company_id = c.id
        AND LOWER(TRIM(p.name)) = LOWER(TRIM(s.name))
        AND p.deleted_at IS NULL
  );

COMMIT;
