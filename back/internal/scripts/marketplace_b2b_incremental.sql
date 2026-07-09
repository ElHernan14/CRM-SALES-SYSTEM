-- Marketplace B2B / Company branding incremental

ALTER TABLE company
ADD COLUMN IF NOT EXISTS description TEXT,
ADD COLUMN IF NOT EXISTS logo_path TEXT;

-- Necesario para que company_user pueda actualizar su logo si el entorno usa permisos sembrados en DB.
INSERT INTO role_permission (role_id, permission_id)
SELECT r.id, p.id
FROM rol r
JOIN permission p ON p.name IN ('company:read', 'company:update', 'company:view')
WHERE r.name = 'company_user'
ON CONFLICT DO NOTHING;
