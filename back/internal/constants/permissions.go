package constants

// ===== CLIENTS =====
const (
	ClientCreate = "client:create"
	ClientRead   = "client:read"
	ClientUpdate = "client:update"
	ClientDelete = "client:delete"
)

// ===== COMPANIES =====
const (
	CompanyCreate = "company:create"
	CompanyRead   = "company:read"
	CompanyUpdate = "company:update"
	CompanyDelete = "company:delete"
)

// ===== PRODUCTS =====
const (
	ProductCreate = "product:create"
	ProductRead   = "product:read"
	ProductUpdate = "product:update"
	ProductDelete = "product:delete"
)

// ===== INVOICES =====
const (
	InvoiceCreate = "invoice:create"
	InvoiceRead   = "invoice:read"
	InvoiceUpdate = "invoice:update"
	InvoiceDelete = "invoice:delete"

	// acciones de negocio
	InvoicePay    = "invoice:pay"
	InvoiceCancel = "invoice:cancel"
)

// ===== INVOICE ITEMS =====
const (
	InvoiceItemCreate = "invoice_item:create"
	InvoiceItemRead   = "invoice_item:read"
	InvoiceItemUpdate = "invoice_item:update"
	InvoiceItemDelete = "invoice_item:delete"
)

var AllPermissions = []string{
	// clients
	ClientCreate, ClientRead, ClientUpdate, ClientDelete,

	// companies
	CompanyCreate, CompanyRead, CompanyUpdate, CompanyDelete,

	// products
	ProductCreate, ProductRead, ProductUpdate, ProductDelete,

	// invoices
	InvoiceCreate, InvoiceRead, InvoiceUpdate, InvoiceDelete,
	InvoicePay, InvoiceCancel,

	// invoice items
	InvoiceItemCreate, InvoiceItemRead, InvoiceItemUpdate, InvoiceItemDelete,
}
