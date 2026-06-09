INSERT INTO permission (name) VALUES

-- ===== CLIENTS =====
('client:view'),
('client:create'),
('client:read'),
('client:update'),
('client:delete'),

-- ===== COMPANIES =====
('company:view'),
('company:create'),
('company:read'),
('company:update'),
('company:delete'),

-- ===== PRODUCTS =====
('product:view'),
('product:create'),
('product:read'),
('product:update'),
('product:delete'),

-- ===== INVOICES =====
('invoice:view'),
('invoice:create'),
('invoice:read'),
('invoice:update'),
('invoice:delete'),
('invoice:pay'),
('invoice:cancel'),

-- ===== INVOICE ITEMS =====
('invoice_item:view'),
('invoice_item:create'),
('invoice_item:read'),
('invoice_item:update'),
('invoice_item:delete');

-- ===== rol =====
INSERT INTO rol (name) VALUES
('super_admin'),
('company_user'),
('individual_user');


-- ===== Asignar todos los permisos al rol super_admin =====
INSERT INTO role_permission (role_id, permission_id)
SELECT r.id, p.id
FROM rol r
JOIN permission p ON 1=1
WHERE r.name = 'super_admin';

-- ===== Asignar permisos al rol company_user =====
INSERT INTO role_permission (role_id, permission_id)
SELECT r.id, p.id
FROM rol r
JOIN permission p ON p.name IN (

    -- CLIENTS
    'client:view','client:create','client:read','client:update',

    -- PRODUCTS
    'product:view','product:create','product:read','product:update',

    -- INVOICES
    'invoice:view','invoice:create','invoice:read','invoice:update','invoice:pay',

    -- INVOICE ITEMS
    'invoice_item:view','invoice_item:create','invoice_item:read','invoice_item:update'

)
WHERE r.name = 'company_user';

-- ===== Asignar permisos al rol individual_user =====
INSERT INTO role_permission (role_id, permission_id)
SELECT r.id, p.id
FROM rol r
JOIN permission p ON p.name IN (

    -- PRODUCTS
    'product:view','product:read',

    -- INVOICES
    'invoice:view','invoice:create','invoice:read'

)
WHERE r.name = 'individual_user';

-- CLIENTS VIEW SCOPES
INSERT INTO permission (name) VALUES
('client:view:all'),
('client:view:company');

-- ADMIN → acceso total
INSERT INTO role_permission (role_id, permission_id)
SELECT r.id, p.id
FROM rol r
JOIN permission p ON p.name IN (
    'client:view:all',
    'client:view:company'
)
WHERE r.name = 'super_admin';


-- 🏢 MANAGER → solo su empresa
INSERT INTO role_permission (role_id, permission_id)
SELECT r.id, p.id
FROM rol r
JOIN permission p ON p.name IN (
    'client:view:company'
)
WHERE r.name = 'company_user';


-- Manager también puede eliminar productos
INSERT INTO role_permission (role_id, permission_id)
SELECT r.id, p.id
FROM rol r
JOIN permission p ON p.name IN (
    'product:delete',
)
WHERE r.name = 'company_user';

-- USER → sin acceso a clientes, solo ver productos e interactuar con sus facturas
INSERT INTO role_permission (role_id, permission_id)
SELECT r.id, p.id
FROM rol r
JOIN permission p
ON p.name = 'invoice:pay'
WHERE r.name = 'individual_user';

INSERT INTO permission (name, description)
VALUES
('invoice:submit', 'Submit draft invoice'),

INSERT INTO role_permission (role_id, permission_id)
SELECT r.id, p.id
FROM rol r
JOIN permission p ON p.name IN (

    -- INVOICES
    'invoice:submit',

    -- INVOICE ITEMS
    'invoice_item:delete','invoice_item:create','invoice_item:read','invoice_item:update'

)
WHERE r.name = 'individual_user';

-- submit invoice también para company_user
INSERT INTO role_permission (role_id, permission_id)
SELECT r.id, p.id
FROM rol r
JOIN permission p ON p.name IN (
    'invoice:submit'
)
WHERE r.name = 'company_user';