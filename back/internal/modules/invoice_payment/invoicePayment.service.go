package invoicepayment

import (
	"context"
	metadto "crm-system-sales/internal/core/dto"
	tenantHelper "crm-system-sales/internal/core/tenant"
	clientrepo "crm-system-sales/internal/modules/client"
	invoiceAccess "crm-system-sales/internal/modules/invoice/access"
	invoicerepo "crm-system-sales/internal/modules/invoice/repository"
	invoicepaymentdto "crm-system-sales/internal/modules/invoice_payment/dto"
	invoicepaymentrepo "crm-system-sales/internal/modules/invoice_payment/repository"
	"database/sql"
)

type InvoicePaymentService interface {
	GetByInvoice(
		ctx context.Context,
		invoiceID int,
		req *invoicepaymentdto.GetInvoicePaymentsRequest,
	) (*invoicepaymentdto.GetInvoicePaymentsResponse, error)
}

type invoicePaymentService struct {
	db                 *sql.DB
	InvoiceRepo        invoicerepo.InvoiceRepository
	ClientRepo         clientrepo.ClientRepository
	InvoicePaymentRepo invoicepaymentrepo.InvoicePaymentRepository
}

func NewInvoicePaymentService(
	db *sql.DB,
	invoiceRepo invoicerepo.InvoiceRepository,
	clientRepo clientrepo.ClientRepository,
	invoicePaymentRepo invoicepaymentrepo.InvoicePaymentRepository,
) *invoicePaymentService {
	return &invoicePaymentService{
		db:                 db,
		InvoiceRepo:        invoiceRepo,
		ClientRepo:         clientRepo,
		InvoicePaymentRepo: invoicePaymentRepo,
	}
}

func (s *invoicePaymentService) GetByInvoice(
	ctx context.Context,
	invoiceID int,
	req *invoicepaymentdto.GetInvoicePaymentsRequest,
) (*invoicepaymentdto.GetInvoicePaymentsResponse, error) {

	var err error

	tenant := tenantHelper.GetTenant(ctx)

	invoice, err := s.InvoiceRepo.GetByID(
		ctx,
		invoiceID,
	)

	if err != nil {
		return nil, err
	}

	buyer, err := s.ClientRepo.GetByID(
		ctx,
		invoice.BuyerClientID,
	)

	if err != nil {
		return nil, err
	}

	err = invoiceAccess.CanViewInvoice(
		tenant,
		invoice,
		buyer,
	)

	if err != nil {
		return nil, err
	}

	payments, total, err := s.InvoicePaymentRepo.ListByInvoiceID(
		ctx,
		invoiceID,
		req,
	)

	if err != nil {
		return nil, err
	}

	items := make(
		[]invoicepaymentdto.PaymentResponse,
		0,
		len(payments),
	)

	for _, payment := range payments {

		items = append(items,
			invoicepaymentdto.PaymentResponse{
				ID:            payment.ID,
				InvoiceID:     payment.InvoiceID,
				Amount:        payment.Amount,
				PaymentMethod: payment.PaymentMethod,
				PaidByUserID:  payment.PaidByUserID,
				CreatedAt:     payment.CreatedAt,
			},
		)
	}

	res := &invoicepaymentdto.GetInvoicePaymentsResponse{
		Items: items,
		Meta: metadto.Meta{
			Page:  req.Page,
			Limit: req.Limit,
			Total: total,
		},
	}

	return res, nil
}
