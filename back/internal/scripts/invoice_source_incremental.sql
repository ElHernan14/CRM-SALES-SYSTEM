ALTER TABLE invoice
ADD COLUMN IF NOT EXISTS source VARCHAR(20);

UPDATE invoice
SET source = 'erp'
WHERE source IS NULL;

ALTER TABLE invoice
ALTER COLUMN source SET DEFAULT 'erp';

ALTER TABLE invoice
ALTER COLUMN source SET NOT NULL;

ALTER TABLE invoice
DROP CONSTRAINT IF EXISTS chk_invoice_source;

ALTER TABLE invoice
ADD CONSTRAINT chk_invoice_source
CHECK (source IN ('erp', 'store'));

CREATE INDEX IF NOT EXISTS idx_invoice_source
ON invoice(source);

CREATE INDEX IF NOT EXISTS idx_invoice_buyer_source_status
ON invoice(buyer_client_id, source, status_invoice, status);

CREATE INDEX IF NOT EXISTS idx_invoice_seller_source_status
ON invoice(seller_company_id, source, status_invoice, status);
