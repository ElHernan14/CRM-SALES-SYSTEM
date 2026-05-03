package constants

const (
	RoleAdmin   = "super_admin"
	RoleManager = "company_user"
	RoleUser    = "individual_user"
)

var RolePermissions = map[string][]string{
	RoleAdmin: {
		// todo
		ClientCreate, ClientRead, ClientUpdate, ClientDelete,
		CompanyCreate, CompanyRead, CompanyUpdate, CompanyDelete,
		ProductCreate, ProductRead, ProductUpdate, ProductDelete,
		InvoiceCreate, InvoiceRead, InvoiceUpdate, InvoiceDelete,
		InvoicePay, InvoiceCancel,
		InvoiceItemCreate, InvoiceItemRead, InvoiceItemUpdate, InvoiceItemDelete,
	},

	RoleManager: {
		ClientRead, ClientUpdate,
		ProductCreate, ProductRead, ProductUpdate,
		InvoiceCreate, InvoiceRead, InvoiceUpdate,
		InvoicePay,
		InvoiceItemCreate, InvoiceItemRead, InvoiceItemUpdate,
	},

	RoleUser: {
		ProductRead,
		InvoiceCreate,
		InvoiceRead,
	},
}
