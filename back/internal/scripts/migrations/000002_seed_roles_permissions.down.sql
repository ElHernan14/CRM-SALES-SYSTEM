-- Removes the roles and permissions inserted by 000002_seed_roles_permissions.up.sql.
-- Run only after dependent seed users/roles were removed.

DELETE FROM role_permission
WHERE role_id IN (
    SELECT id FROM rol
    WHERE name IN ('super_admin', 'company_user', 'individual_user')
);

DELETE FROM user_rol
WHERE role_id IN (
    SELECT id FROM rol
    WHERE name IN ('super_admin', 'company_user', 'individual_user')
);

DELETE FROM rol
WHERE name IN ('super_admin', 'company_user', 'individual_user');

DELETE FROM permission
WHERE name IN (
    'user:create',
    'company:create',
    'company:view',
    'company:read',
    'company:update',
    'company:delete',
    'client:view',
    'client:create',
    'client:read',
    'client:update',
    'client:delete',
    'client:view:all',
    'client:view:company',
    'product:view',
    'product:create',
    'product:read',
    'product:update',
    'product:delete',
    'invoice:view',
    'invoice:create',
    'invoice:read',
    'invoice:update',
    'invoice:delete',
    'invoice:submit',
    'invoice:pay',
    'invoice:cancel',
    'invoice_item:create',
    'invoice_item:read',
    'invoice_item:update',
    'invoice_item:delete',
    'invoice_payment:read',
    'store:checkout'
);
