package docs

import (
	"crm-system-sales/internal/core/response"
	invoicedto "crm-system-sales/internal/modules/invoice/dto"
	invoiceitemdto "crm-system-sales/internal/modules/invoice_item/dto"
	productdto "crm-system-sales/internal/modules/product/dto"
	storedto "crm-system-sales/internal/modules/store/dto"
)

type SuccessResponse = response.APIResponse
type ErrorResponse = response.APIResponse

// * ----------------------- PRODUCTS ----------------------- *//
type ProductDetailSuccessResponse struct {
	Status string                           `json:"status"`
	Code   int                              `json:"code"`
	Data   productdto.ProductDetailResponse `json:"data"`
}

type GetProductsSuccessResponse struct {
	Status string                         `json:"status"`
	Code   int                            `json:"code"`
	Data   productdto.GetProductsResponse `json:"data"`
}

// * ----------------------- INVOICES ----------------------- *//
type CreateInvoiceSuccessResponse struct {
	Status string                     `json:"status" example:"success"`
	Code   int                        `json:"code" example:"200"`
	Data   invoicedto.InvoiceResponse `json:"data"`
}

type SubmitInvoiceSuccessResponse struct {
	Status string `json:"status" example:"success"`
	Code   int    `json:"code" example:"200"`
	Data   string `json:"data" example:"invoice enviada correctamente"`
}

type PayInvoiceSuccessResponse struct {
	Status string                        `json:"status" example:"success"`
	Code   int                           `json:"code" example:"200"`
	Data   invoicedto.PayInvoiceResponse `json:"data"`
}

type GetInvoiceSuccessResponse struct {
	Status string                        `json:"status" example:"success"`
	Code   int                           `json:"code" example:"200"`
	Data   invoicedto.GetInvoiceResponse `json:"data"`
}

// * ----------------------- INVOICE_ITEM ----------------------- *//
type CreateInvoiceItemSuccessResponse struct {
	Status string                             `json:"status" example:"success"`
	Code   int                                `json:"code" example:"200"`
	Data   invoiceitemdto.InvoiceItemResponse `json:"data"`
}

type UpdateInvoiceItemSuccessResponse struct {
	Status string                             `json:"status" example:"success"`
	Code   int                                `json:"code" example:"200"`
	Data   invoiceitemdto.InvoiceItemResponse `json:"data"`
}

type DeleteInvoiceItemSuccessResponse struct {
	Status string `json:"status" example:"success"`
	Code   int    `json:"code" example:"200"`
	Data   string `json:"data" example:"item eliminado correctamente"`
}

// * ----------------------- STORE ----------------------- *//
type GetStoreProductsSuccessResponse struct {
	Status string                            `json:"status" example:"success"`
	Code   int                               `json:"code" example:"200"`
	Data   storedto.GetStoreProductsResponse `json:"data"`
}

type CheckoutSuccessResponse struct {
	Status string                    `json:"status" example:"success"`
	Code   int                       `json:"code" example:"200"`
	Data   storedto.CheckoutResponse `json:"data"`
}
