package helper

import (
	errorHandler "crm-system-sales/internal/core/error"
	validatorx "crm-system-sales/internal/core/validator"
	clientdto "crm-system-sales/internal/modules/client/dto"
	invoicedto "crm-system-sales/internal/modules/invoice/dto"
	invoiceitemdto "crm-system-sales/internal/modules/invoice_item/dto"
	invoicepaymentdto "crm-system-sales/internal/modules/invoice_payment/dto"
	marketplacedto "crm-system-sales/internal/modules/marketplace/dto"
	productdto "crm-system-sales/internal/modules/product/dto"
	storedto "crm-system-sales/internal/modules/store/dto"
	"net/http"
	"strconv"
)

func ParseGetClientsRequest(r *http.Request) (*clientdto.GetClientsRequest, error) {
	q := r.URL.Query()
	req := &clientdto.GetClientsRequest{Search: q.Get("search"), Email: q.Get("email")}

	if v := q.Get("company_id"); v != "" {
		id, err := strconv.Atoi(v)
		if err != nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "company_id inválido")
		}
		req.CompanyID = &id
	}
	if err := parsePagination(q.Get("page"), q.Get("limit"), &req.Page, &req.Limit); err != nil {
		return nil, err
	}

	return validateRequest(req)
}

func ParseGetInvoicesRequest(r *http.Request) (*invoiceitemdto.GetInvoiceItemsRequest, error) {
	q := r.URL.Query()
	req := &invoiceitemdto.GetInvoiceItemsRequest{}
	if err := parsePagination(q.Get("page"), q.Get("limit"), &req.Page, &req.Limit); err != nil {
		return nil, err
	}
	return validateRequest(req)
}

func ParseGetInvoicePaymentsRequest(r *http.Request) (*invoicepaymentdto.GetInvoicePaymentsRequest, error) {
	q := r.URL.Query()
	req := &invoicepaymentdto.GetInvoicePaymentsRequest{PaymentMethod: q.Get("payment_method"), SortColumn: q.Get("sort_column"), Order: q.Get("order")}
	if err := parsePagination(q.Get("page"), q.Get("limit"), &req.Page, &req.Limit); err != nil {
		return nil, err
	}
	return validateRequest(req)
}

func ParseGetCompanyInvoicesRequest(r *http.Request) (*invoicedto.GetCompanyInvoicesRequest, error) {
	q := r.URL.Query()
	req := &invoicedto.GetCompanyInvoicesRequest{StatusInvoice: q.Get("status_invoice"), BuyerName: q.Get("buyer_name"), SellerCompany: q.Get("seller_company"), SortColumn: q.Get("sort_column"), Order: q.Get("order")}

	if v := q.Get("status"); v != "" {
		status, err := strconv.Atoi(v)
		if err != nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "status inválido")
		}
		req.Status = status
	}
	if err := parsePagination(q.Get("page"), q.Get("limit"), &req.Page, &req.Limit); err != nil {
		return nil, err
	}
	return validateRequest(req)
}

func ParseGetCompanyProductsRequest(r *http.Request) (*productdto.GetCompanyProductsRequest, error) {
	q := r.URL.Query()
	req := &productdto.GetCompanyProductsRequest{Name: q.Get("name"), Kind: q.Get("kind"), Type: q.Get("type"), Category: q.Get("category"), SortColumn: q.Get("sort_column"), Order: q.Get("order")}
	if req.Kind == "" && (req.Type == "product" || req.Type == "service") {
		req.Kind = req.Type
		req.Type = ""
	}

	if v := q.Get("category_id"); v != "" {
		categoryID, err := strconv.Atoi(v)
		if err != nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "category_id invalido")
		}
		req.CategoryID = &categoryID
	}
	if v := q.Get("type_id"); v != "" {
		typeID, err := strconv.Atoi(v)
		if err != nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "type_id invalido")
		}
		req.TypeID = &typeID
	}
	if v := q.Get("status"); v != "" {
		status, err := strconv.Atoi(v)
		if err != nil || (status != 0 && status != 1) {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "status invalido")
		}
		req.Status = &status
	}
	if err := parsePagination(q.Get("page"), q.Get("limit"), &req.Page, &req.Limit); err != nil {
		return nil, err
	}
	return validateRequest(req)
}

func ParseGetCompanyCustomersRequest(r *http.Request) (*clientdto.GetCompanyCustomersRequest, error) {
	q := r.URL.Query()
	req := &clientdto.GetCompanyCustomersRequest{Name: q.Get("name"), SortColumn: q.Get("sort_column"), Order: q.Get("order")}
	if err := parsePagination(q.Get("page"), q.Get("limit"), &req.Page, &req.Limit); err != nil {
		return nil, err
	}
	return validateRequest(req)
}

func ParseGetStoreProductsRequest(r *http.Request) (*storedto.GetStoreProductsRequest, error) {
	q := r.URL.Query()
	req := &storedto.GetStoreProductsRequest{
		Search:     q.Get("search"),
		Name:       q.Get("name"),
		Kind:       q.Get("kind"),
		Type:       q.Get("type"),
		Category:   q.Get("category"),
		SortColumn: q.Get("sort_column"),
		Order:      q.Get("order"),
	}
	if req.Kind == "" && (req.Type == "product" || req.Type == "service") {
		req.Kind = req.Type
		req.Type = ""
	}

	if v := q.Get("company_id"); v != "" {
		companyID, err := strconv.Atoi(v)
		if err != nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "company_id inválido")
		}
		req.CompanyID = &companyID
	}
	if v := q.Get("category_id"); v != "" {
		categoryID, err := strconv.Atoi(v)
		if err != nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "category_id inválido")
		}
		req.CategoryID = &categoryID
	}
	if v := q.Get("type_id"); v != "" {
		typeID, err := strconv.Atoi(v)
		if err != nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "type_id invalido")
		}
		req.TypeID = &typeID
	}
	if v := q.Get("min_price"); v != "" {
		minPrice, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "min_price inválido")
		}
		req.MinPrice = minPrice
	}
	if v := q.Get("max_price"); v != "" {
		maxPrice, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "max_price inválido")
		}
		req.MaxPrice = maxPrice
	}
	if err := parsePagination(q.Get("page"), q.Get("limit"), &req.Page, &req.Limit); err != nil {
		return nil, err
	}
	return validateRequest(req)
}

func ParseGetSuppliersRequest(r *http.Request) (*marketplacedto.GetSuppliersRequest, error) {
	q := r.URL.Query()
	req := &marketplacedto.GetSuppliersRequest{Search: q.Get("search"), Category: q.Get("category"), SortColumn: q.Get("sort_column"), Order: q.Get("order")}
	if v := q.Get("category_id"); v != "" {
		categoryID, err := strconv.Atoi(v)
		if err != nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "category_id invalido")
		}
		req.CategoryID = &categoryID
	}
	if err := parsePagination(q.Get("page"), q.Get("limit"), &req.Page, &req.Limit); err != nil {
		return nil, err
	}
	return validateRequest(req)
}

func parsePagination(rawPage string, rawLimit string, page *int, limit *int) error {
	if rawPage != "" {
		parsed, err := strconv.Atoi(rawPage)
		if err != nil {
			return errorHandler.NewAppError(http.StatusBadRequest, "page inválido")
		}
		*page = parsed
	}
	if rawLimit != "" {
		parsed, err := strconv.Atoi(rawLimit)
		if err != nil {
			return errorHandler.NewAppError(http.StatusBadRequest, "limit inválido")
		}
		*limit = parsed
	}
	if *page == 0 {
		*page = 1
	}
	if *limit == 0 {
		*limit = 10
	}
	return nil
}

func validateRequest[T any](req *T) (*T, error) {
	msg, invalid := validatorx.ValidateStruct(req)
	if invalid {
		return nil, errorHandler.NewAppError(http.StatusBadRequest, msg)
	}
	return req, nil
}
