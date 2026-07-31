-- Idempotent roles and permissions seed.
-- Safe to run more than once after INIT.sql.

INSERT INTO permission (name, description) VALUES
('user:create', 'Create users'),

('client:view', 'View clients'),
('client:create', 'Create clients'),
('client:read', 'Read clients'),
('client:update', 'Update clients'),
('client:delete', 'Delete clients'),
('client:view:all', 'View all clients'),
('client:view:company', 'View company clients'),

('company:view', 'View companies'),
('company:create', 'Create companies'),
('company:read', 'Read company'),
('company:update', 'Update company'),
('company:delete', 'Delete company'),

('product:view', 'View products'),
('product:create', 'Create products'),
('product:read', 'Read products'),
('product:update', 'Update products'),
('product:delete', 'Delete products'),

('invoice:view', 'View invoices'),
('invoice:create', 'Create invoices'),
('invoice:read', 'Read invoices'),
('invoice:update', 'Update invoices'),
('invoice:delete', 'Delete invoices'),
('invoice:pay', 'Pay invoices'),
('invoice:cancel', 'Cancel invoices'),
('invoice:submit', 'Submit draft invoices'),

('invoice_item:view', 'View invoice items'),
('invoice_item:create', 'Create invoice items'),
('invoice_item:read', 'Read invoice items'),
('invoice_item:update', 'Update invoice items'),
('invoice_item:delete', 'Delete invoice items'),

('invoice_payment:read', 'Read invoice payments'),

('store:checkout', 'Checkout store carts')
ON CONFLICT (name) DO UPDATE
SET description = EXCLUDED.description;

INSERT INTO rol (name) VALUES
('super_admin'),
('company_user'),
('individual_user')
ON CONFLICT (name) DO NOTHING;

-- super_admin: full access.
INSERT INTO role_permission (role_id, permission_id)
SELECT r.id, p.id
FROM rol r
JOIN permission p ON TRUE
WHERE r.name = 'super_admin'
ON CONFLICT DO NOTHING;

-- company_user: ERP workspace, Marketplace B2B, Store preview/purchase basics.
INSERT INTO role_permission (role_id, permission_id)
SELECT r.id, p.id
FROM rol r
JOIN permission p ON p.name IN (
    'client:view',
    'client:create',
    'client:read',
    'client:update',
    'client:delete',
    'client:view:company',

    'company:view',
    'company:read',
    'company:update',

    'product:view',
    'product:create',
    'product:read',
    'product:update',
    'product:delete',

    'invoice:view',
    'invoice:create',
    'invoice:read',
    'invoice:update',
    'invoice:pay',
    'invoice:cancel',
    'invoice:submit',

    'invoice_item:view',
    'invoice_item:create',
    'invoice_item:read',
    'invoice_item:update',
    'invoice_item:delete',

    'invoice_payment:read',
    'store:checkout'
)
WHERE r.name = 'company_user'
ON CONFLICT DO NOTHING;

-- individual_user: Store buyer and personal purchase access.
INSERT INTO role_permission (role_id, permission_id)
SELECT r.id, p.id
FROM rol r
JOIN permission p ON p.name IN (
    'product:view',
    'product:read',

    'invoice:view',
    'invoice:create',
    'invoice:read',
    'invoice:pay',
    'invoice:submit',

    'invoice_item:view',
    'invoice_item:create',
    'invoice_item:read',
    'invoice_item:update',
    'invoice_item:delete',

    'invoice_payment:read',
    'store:checkout'
)
WHERE r.name = 'individual_user'
ON CONFLICT DO NOTHING;
