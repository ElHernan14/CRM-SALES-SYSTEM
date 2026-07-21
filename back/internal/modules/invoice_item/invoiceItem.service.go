package invoiceitem

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	errorHandler "crm-system-sales/internal/core/error"
	tenantHelper "crm-system-sales/internal/core/tenant"
	tenantctx "crm-system-sales/internal/core/tenant"
	core "crm-system-sales/internal/core/transaction"

	"crm-system-sales/internal/modules/client"
	"crm-system-sales/internal/modules/inventory"
	invoiceaccess "crm-system-sales/internal/modules/invoice/access"
	"crm-system-sales/internal/modules/invoice/constants"
	invoice "crm-system-sales/internal/modules/invoice/repository"
	"crm-system-sales/internal/modules/product"

	metadto "crm-system-sales/internal/core/dto"
	invoiceItemDTO "crm-system-sales/internal/modules/invoice_item/dto"
	invoiceitemdto "crm-system-sales/internal/modules/invoice_item/dto"
	invoiceItemModel "crm-system-sales/internal/modules/invoice_item/models"
)

type InvoiceItemService interface {
	Create(ctx context.Context, invoiceID int, req *invoiceItemDTO.CreateInvoiceItemRequest) (*invoiceItemDTO.InvoiceItemResponse, error)
	Update(ctx context.Context, invoiceID int, itemID int, req *invoiceItemDTO.UpdateInvoiceItemRequest) (*invoiceItemDTO.InvoiceItemResponse, error)
	Delete(ctx context.Context, invoiceID int, itemID int) error
	GetInvoiceItems(ctx context.Context, invoiceID int, req *invoiceItemDTO.GetInvoiceItemsRequest) (*invoiceItemDTO.GetInvoiceItemsResponse, error)
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
		// Comprobar si ya existe un item para este producto en la invoice
		existingItem, err := s.repo.GetByInvoiceAndProduct(
			ctx,
			invoice.ID,
			product.ID,
		)

		if err != nil {
			log.Println("error checking existing invoice item:", err)
			return errorHandler.NewAppError(
				http.StatusInternalServerError,
				"No se pudo verificar si el item ya existe",
			)
		}

		// Compruebo si el item ya existe para esta invoice y producto, si existe sumo la cantidad, sino creo una nueva
		if existingItem != nil {
			// Si el item ya existe, actualizar la cantidad sumando la nueva cantidad a la existente
			newQty := existingItem.Quantity + req.Quantity

			diff := newQty - existingItem.Quantity

			err = s.inventoryService.AdjustReservedStock(
				ctx,
				tx,
				product.ID,
				diff,
			)
			existingItem.Price = product.Price
			existingItem.Quantity = newQty
			existingItem.Subtotal = product.Price * float64(newQty)

			// Actualizo el item existente con la nueva cantidad y subtotal junto con la invoice
			err = s.repo.Update(
				ctx,
				tx,
				existingItem,
			)

			if err != nil {
				log.Println("error updating invoice item:", err)
				return errorHandler.NewAppError(
					http.StatusInternalServerError,
					"No se pudo actualizar el item ya existente de invoice",
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
					"No se pudo recalcular el nuevo total de la invoice",
				)
			}

			response = &invoiceItemDTO.InvoiceItemResponse{
				ID:          existingItem.ID,
				InvoiceID:   &existingItem.InvoiceID,
				ProductID:   existingItem.ProductID,
				ProductName: existingItem.ProductName,
				Quantity:    existingItem.Quantity,
				Price:       existingItem.Price,
				Subtotal:    existingItem.Subtotal,
			}

			return nil
		}

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

		err = s.repo.Create(ctx, tx, item)
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
			InvoiceID: &item.InvoiceID,

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

	//  product
	product, err := s.productRepo.GetByID(ctx, item.ProductID)
	if err != nil {
		return nil, errorHandler.NewAppError(
			http.StatusNotFound,
			"Producto no encontrado",
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

		// update price item
		item.Price = product.Price

		//  update quantity
		item.Quantity = req.Quantity

		//  subtotal
		item.Subtotal = product.Price * float64(item.Quantity)

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
		InvoiceID: &item.InvoiceID,

		ProductID:   item.ProductID,
		ProductName: item.ProductName,

		Quantity: item.Quantity,

		Price:    item.Price,
		Subtotal: item.Subtotal,
	}, nil
}

func (s *invoiceItemService) Delete(
	ctx context.Context,
	invoiceID int,
	itemID int,
) error {

	var err error

	tenant := tenantctx.GetTenant(ctx)

	item, err := s.repo.GetByID(ctx, itemID)
	if err != nil {
		return errorHandler.NewAppError(
			http.StatusNotFound,
			"invoice item no encontrado",
		)
	}

	if item.InvoiceID != invoiceID {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			"invoice item inválido",
		)
	}

	invoice, err := s.invoiceRepo.GetByID(ctx, invoiceID)
	if err != nil {
		return errorHandler.NewAppError(
			http.StatusNotFound,
			"Invoice no encontrada",
		)
	}

	buyer, err := s.clientRepo.GetByID(ctx, invoice.BuyerClientID)
	if err != nil {
		return errorHandler.NewAppError(
			http.StatusNotFound,
			"Cliente no encontrado",
		)
	}

	// ownership
	err = invoiceaccess.CanMutateDraftInvoice(
		tenant,
		invoice,
		buyer,
	)

	if err != nil {
		return err
	}

	// draft only
	if invoice.StatusInvoice != constants.InvoiceDraft {
		return errorHandler.NewAppError(
			http.StatusForbidden,
			"solo se pueden eliminar items de invoices draft",
		)
	}

	err = core.RunInTransaction(ctx, s.db, func(tx *sql.Tx) error {

		// liberar stock
		err = s.inventoryService.ReleaseStock(
			ctx,
			tx,
			item.ProductID,
			item.Quantity,
		)

		if err != nil {
			log.Println("error releasing reserved stock:", err)
			return errorHandler.NewAppError(
				http.StatusInternalServerError,
				"No se pudo liberar el stock reservado",
			)
		}

		//  eliminar item
		err = s.repo.Delete(
			ctx,
			tx,
			item.ID,
		)

		if err != nil {
			log.Println("error deleting invoice item:", err)
			return errorHandler.NewAppError(
				http.StatusInternalServerError,
				"No se pudo eliminar el item de invoice",
			)
		}

		// recalcular invoice
		return s.invoiceRepo.RecalculateInvoiceTotals(
			ctx,
			tx,
			invoice.ID,
		)
	})

	return err
}

func (s *invoiceItemService) GetInvoiceItems(
	ctx context.Context,
	invoiceID int,
	req *invoiceItemDTO.GetInvoiceItemsRequest,
) (*invoiceItemDTO.GetInvoiceItemsResponse, error) {
	var err error

	tenant := tenantHelper.GetTenant(ctx)

	invoice, err := s.invoiceRepo.GetByID(ctx, invoiceID)
	if err != nil {
		log.Println("error fetching invoice:", err)
		return nil, errorHandler.NewAppError(
			http.StatusNotFound,
			"Invoice no encontrada",
		)
	}

	buyer, err := s.clientRepo.GetByID(ctx, invoice.BuyerClientID)
	if err != nil {
		log.Println("error fetching buyer client:", err)
		return nil, errorHandler.NewAppError(
			http.StatusNotFound,
			"Cliente comprador no encontrado",
		)
	}

	err = invoiceaccess.CanViewInvoice(tenant, invoice, buyer)
	if err != nil {
		return nil, err
	}

	items, total, err := s.repo.GetByInvoiceID(ctx, invoiceID, req.Page, req.Limit)
	if err != nil {
		log.Println("error fetching invoice items:", err)
		return nil, errorHandler.NewAppError(
			http.StatusInternalServerError,
			"Error obteniendo items de la invoice",
		)
	}

	resp := &invoiceItemDTO.GetInvoiceItemsResponse{
		Items: []invoiceItemDTO.InvoiceItemResponse{},
		Meta:  metadto.NewMeta(req.Page, req.Limit, total),
	}

	// calcular total pages
	if total > 0 {
		resp.Meta.TotalPages = (total + req.Limit - 1) / req.Limit
	}

	for _, item := range items {
		resp.Items = append(resp.Items, invoiceItemDTO.InvoiceItemResponse{
			ID:          item.ID,
			ProductID:   item.ProductID,
			ProductName: item.ProductName,
			Quantity:    item.Quantity,
			Price:       item.Price,
			Subtotal:    item.Subtotal,
		})
	}

	return resp, nil
}
