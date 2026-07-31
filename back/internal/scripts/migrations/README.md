# Nexora SQL migrations

These files define a reproducible order for creating a fresh production database.

Run order for a new empty Supabase database:

1. `000001_init.up.sql`
2. `000002_seed_roles_permissions.up.sql`
3. `000003_seed_categories_product_types.up.sql`
4. `000004_seed_companies_clients.up.sql`
5. `000005_seed_products.up.sql`

Repair migration:

- `000006_repair_seed_company_clients_roles.up.sql` is only needed if `000004_seed_companies_clients.up.sql` was executed before the client/role fix and created companies/users without their associated clients or `company_user` role.
- Do not run `000006` in a fresh database where the corrected `000004` already created clients and roles, although it is idempotent and should not duplicate rows.

Notes:

- `INIT.sql` is absorbed by `000001_init.up.sql`.
- `init_permisos_roles.sql` is absorbed by `000002_seed_roles_permissions.up.sql`.
- `nexora_seed_categories_and_product_types.sql` is absorbed by `000003_seed_categories_product_types.up.sql`.
- `seed_companys_clients-company.sql` is absorbed by `000004_seed_companies_clients.up.sql`.
- `seed_products.sql` is absorbed by `000005_seed_products.up.sql`.
- Incremental scripts such as `category_split_incremental.sql`, `product_type_incremental.sql`, `invoice_source_incremental.sql`, `marketplace_b2b_incremental.sql`, `product_image_incremental.sql` and `company_cover_image_incremental.sql` are for existing databases. They are already represented in the current init migration.
- `product_type_detach_products.sql` is a manual maintenance script, not part of the normal production bootstrap.
