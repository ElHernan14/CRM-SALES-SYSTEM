package constants

// ===== CLIENTS =====
const (
	ClientView        = "client:view"
	ClientViewAll     = "client:view:all"
	ClientViewCompany = "client:view:company"
	ClientCreate      = "client:create"
	ClientRead        = "client:read"
	ClientUpdate      = "client:update"
	ClientDelete      = "client:delete"
)

// ===== COMPANIES =====
const (
	CompanyView   = "company:view"
	CompanyCreate = "company:create"
	CompanyRead   = "company:read"
	CompanyUpdate = "company:update"
	CompanyDelete = "company:delete"
)

// ===== PRODUCTS =====
const (
	ProductView   = "product:view"
	ProductCreate = "product:create"
	ProductRead   = "product:read"
	ProductUpdate = "product:update"
	ProductDelete = "product:delete"
)

// ===== INVOICES =====
const (
	InvoiceView   = "invoice:view"
	InvoiceCreate = "invoice:create"
	InvoiceRead   = "invoice:read"
	InvoiceUpdate = "invoice:update"
	InvoiceDelete = "invoice:delete"

	InvoicePay    = "invoice:pay"
	InvoiceCancel = "invoice:cancel"
)

// ===== INVOICE ITEMS =====
const (
	InvoiceItemView   = "invoice_item:view"
	InvoiceItemCreate = "invoice_item:create"
	InvoiceItemRead   = "invoice_item:read"
	InvoiceItemUpdate = "invoice_item:update"
	InvoiceItemDelete = "invoice_item:delete"
)
