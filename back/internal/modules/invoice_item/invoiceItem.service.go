package invoiceitem

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	errorHandler "crm-system-sales/internal/core/error"
	tenantctx "crm-system-sales/internal/core/tenant"
	core "crm-system-sales/internal/core/transaction"
	"crm-system-sales/internal/core/utils"

	"crm-system-sales/internal/modules/client"
	"crm-system-sales/internal/modules/inventory"
	"crm-system-sales/internal/modules/invoice"
	invoiceaccess "crm-system-sales/internal/modules/invoice/access"
	"crm-system-sales/internal/modules/invoice/constants"
	"crm-system-sales/internal/modules/product"

	invoiceItemDTO "crm-system-sales/internal/modules/invoice_item/dto"
	invoiceitemdto "crm-system-sales/internal/modules/invoice_item/dto"
	invoiceItemModel "crm-system-sales/internal/modules/invoice_item/models"
)

type InvoiceItemService interface {
	Create(ctx context.Context, invoiceID int, req *invoiceItemDTO.CreateInvoiceItemRequest) (*invoiceItemDTO.InvoiceItemResponse, error)
	Update(ctx context.Context, invoiceID int, itemID int, req *invoiceItemDTO.UpdateInvoiceItemRequest) (*invoiceItemDTO.InvoiceItemResponse, error)
}

type invoiceItemService struct {
	db               *sql.DB
	repo             InvoiceItemRepository
	invoiceRepo      invoice.InvoiceRepository
	clientRepo       client.ClientRepository
	productRepo      product.ProductRepository
	inventoryService inventory.InventoryService
}

func NewInvoiceItemService(
	db *sql.DB,
	repo InvoiceItemRepository,
	invoiceRepo invoice.InvoiceRepository,
	clientRepo client.ClientRepository,
	productRepo product.ProductRepository,
	inventoryService inventory.InventoryService,
) *invoiceItemService {
	return &invoiceItemService{
		db:               db,
		repo:             repo,
		invoiceRepo:      invoiceRepo,
		clientRepo:       clientRepo,
		productRepo:      productRepo,
		inventoryService: inventoryService,
	}
}

func (s *invoiceItemService) Create(
	ctx context.Context,
	invoiceID int,
	req *invoiceItemDTO.CreateInvoiceItemRequest,
) (*invoiceItemDTO.InvoiceItemResponse, error) {

	tenant := tenantctx.GetTenant(ctx)

	//  invoice
	invoice, err := s.invoiceRepo.GetByID(ctx, invoiceID)
	if err != nil {
		return nil, errorHandler.NewAppError(
			http.StatusNotFound,
			"Invoice no encontrada",
		)
	}

	//  buyer
	buyer, err := s.clientRepo.GetByID(ctx, invoice.BuyerClientID)
	if err != nil {
		return nil, errorHandler.NewAppError(
			http.StatusNotFound,
			"Cliente no encontrado",
		)
	}

	//  access
	err = invoiceaccess.CanMutateDraftInvoice(
		tenant,
		invoice,
		buyer,
	)

	if err != nil {
		return nil, err
	}

	//  product
	product, err := s.productRepo.GetByID(ctx, req.ProductID)
	if err != nil {
		return nil, errorHandler.NewAppError(
			http.StatusNotFound,
			"Producto no encontrado",
		)
	}

	//  seller ownership
	if product.CompanyID != invoice.SellerCompanyID {
		return nil, errorHandler.NewAppError(
			http.StatusBadRequest,
			"Producto inválido para esta invoice",
		)
	}

	var response *invoiceItemDTO.InvoiceItemResponse

	err = core.RunInTransaction(ctx, s.db, func(tx *sql.Tx) error {

		err = s.inventoryService.ReserveStock(
			ctx,
			tx,
			product.ID,
			req.Quantity,
		)

		if err != nil {
			return err
		}

		subtotal := product.Price * float64(req.Quantity)

		item := &invoiceItemModel.InvoiceItem{
			InvoiceID:   invoice.ID,
			ProductID:   product.ID,
			ProductName: product.Name,

			Quantity: req.Quantity,

			Price:    product.Price,
			Subtotal: subtotal,
		}

		err := s.repo.Create(ctx, tx, item)
		if err != nil {
			log.Println("error creating invoice item:", err)
			return errorHandler.NewAppError(
				http.StatusInternalServerError,
				"No se pudo crear el item de invoice",
			)
		}

		//  recalcular total
		err = s.invoiceRepo.RecalculateInvoiceTotals(
			ctx,
			tx,
			invoice.ID,
		)

		if err != nil {
			log.Println("error recalculating invoice total:", err)
			return errorHandler.NewAppError(
				http.StatusInternalServerError,
				"No se pudo recalcular el total de la invoice",
			)
		}

		response = &invoiceItemDTO.InvoiceItemResponse{
			ID:        item.ID,
			InvoiceID: item.InvoiceID,

			ProductID:   item.ProductID,
			ProductName: item.ProductName,

			Quantity: item.Quantity,

			Price:    item.Price,
			Subtotal: item.Subtotal,
		}

		log.Println("invoice item created:", item.ID)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *invoiceItemService) Update(
	ctx context.Context,
	invoiceID int,
	itemID int,
	req *invoiceitemdto.UpdateInvoiceItemRequest,
) (*invoiceitemdto.InvoiceItemResponse, error) {

	var err error
	defer func() {
		utils.Trace(ctx, "SERVICE UpdateInvoiceItem")(err)
	}()

	tenant := tenantctx.GetTenant(ctx)

	item, err := s.repo.GetByID(ctx, itemID)
	if err != nil {
		return nil, errorHandler.NewAppError(
			http.StatusNotFound,
			"invoice item no encontrado",
		)
	}

	if item.InvoiceID != invoiceID {
		return nil, errorHandler.NewAppError(
			http.StatusBadRequest,
			"invoice item inválido",
		)
	}

	invoice, err := s.invoiceRepo.GetByID(ctx, invoiceID)
	if err != nil {
		return nil, errorHandler.NewAppError(
			http.StatusNotFound,
			"Invoice no encontrada",
		)
	}

	buyer, err := s.clientRepo.GetByID(ctx, invoice.BuyerClientID)
	if err != nil {
		return nil, errorHandler.NewAppError(
			http.StatusNotFound,
			"Cliente no encontrado",
		)
	}

	//  ownership
	err = invoiceaccess.CanMutateDraftInvoice(
		tenant,
		invoice,
		buyer,
	)

	if err != nil {
		return nil, err
	}

	//  draft validation
	if invoice.StatusInvoice != constants.InvoiceDraft {
		return nil, errorHandler.NewAppError(
			http.StatusForbidden,
			"solo se puede modificar invoices draft",
		)
	}

	err = core.RunInTransaction(ctx, s.db, func(tx *sql.Tx) error {

		//  stock diff
		diff := req.Quantity - item.Quantity

		err = s.inventoryService.AdjustReservedStock(
			ctx,
			tx,
			item.ProductID,
			diff,
		)

		if err != nil {
			log.Println("error adjusting reserved stock:", err)
			return errorHandler.NewAppError(
				http.StatusInternalServerError,
				"No se pudo ajustar el stock reservado",
			)
		}

		//  update quantity
		item.Quantity = req.Quantity

		//  subtotal
		item.Subtotal = item.Price * float64(item.Quantity)

		//  persist
		err = s.repo.Update(
			ctx,
			tx,
			item,
		)

		if err != nil {
			log.Println("error updating invoice item:", err)
			return errorHandler.NewAppError(
				http.StatusInternalServerError,
				"No se pudo actualizar el item de invoice",
			)
		}

		err = s.invoiceRepo.RecalculateInvoiceTotals(
			ctx,
			tx,
			invoice.ID,
		)
		if err != nil {
			log.Println("error recalculating invoice total:", err)
			return errorHandler.NewAppError(
				http.StatusInternalServerError,
				"No se pudo recalcular el total de la invoice",
			)
		}

		return err
	})

	if err != nil {
		return nil, err
	}

	return &invoiceitemdto.InvoiceItemResponse{
		ID:        item.ID,
		InvoiceID: item.InvoiceID,

		ProductID:   item.ProductID,
		ProductName: item.ProductName,

		Quantity: item.Quantity,

		Price:    item.Price,
		Subtotal: item.Subtotal,
	}, nil
}
