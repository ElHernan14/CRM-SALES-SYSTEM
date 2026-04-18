INSERT INTO permissions (name) VALUES
-- clients
('client:create'),
('client:read'),
('client:update'),
('client:delete'),

-- companies
('company:create'),
('company:read'),
('company:update'),
('company:delete'),

-- products
('product:create'),
('product:read'),
('product:update'),
('product:delete'),

-- invoices
('invoice:create'),
('invoice:read'),
('invoice:update'),
('invoice:delete'),
('invoice:pay'),
('invoice:cancel'),

-- invoice items
('invoice_item:create'),
('invoice_item:read'),
('invoice_item:update'),
('invoice_item:delete');

INSERT INTO roles (name) VALUES
('admin'),
('manager'),
('user');


/*Asignar todos los permisos al rol admin*/
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
JOIN permissions p ON 1=1
WHERE r.name = 'admin';

/*Asigar permisos al rol manager*/
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
JOIN permissions p ON p.name IN (
    'product:create','product:read','product:update',
    'invoice:create','invoice:read','invoice:update','invoice:pay',
    'invoice_item:create','invoice_item:read','invoice_item:update'
)
WHERE r.name = 'manager';