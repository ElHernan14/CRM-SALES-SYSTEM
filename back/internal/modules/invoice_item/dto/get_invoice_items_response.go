package invoiceitemdto

import metaDto "crm-system-sales/internal/core/dto"

type GetInvoiceItemsResponse struct {
	Items []InvoiceItemResponse `json:"items"`
	Meta  metaDto.Meta          `json:"meta"`
}
