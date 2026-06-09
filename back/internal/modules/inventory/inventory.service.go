package inventory

import (
	"context"
	"database/sql"

	errorHandler "crm-system-sales/internal/core/error"
	productRepo "crm-system-sales/internal/modules/product"

	"net/http"
)

type InventoryService interface {
	ReserveStock(ctx context.Context, tx *sql.Tx, productID int, quantity int) error
	ConsolidateStock(ctx context.Context, tx *sql.Tx, productID int, quantity int) error
	IncrementReservedStock(ctx context.Context, tx *sql.Tx, productID int, quantity int) error
	AdjustReservedStock(ctx context.Context, tx *sql.Tx, productID int, diff int) error
	ReleaseStock(ctx context.Context, tx *sql.Tx, productID int, quantity int) error
	FinalizeReservedStock(ctx context.Context, tx *sql.Tx, productID int, quantity int) error
}

type inventoryService struct {
	productRepo productRepo.ProductRepository
}

func NewInventoryService(
	productRepo productRepo.ProductRepository,
) InventoryService {
	return &inventoryService{
		productRepo: productRepo,
	}
}

func (s *inventoryService) ReserveStock(
	ctx context.Context,
	tx *sql.Tx,
	productID int,
	quantity int,
) error {

	product, err := s.productRepo.GetByIDForUpdate(
		ctx,
		tx,
		productID,
	)

	if err != nil {
		return errorHandler.NewAppError(
			http.StatusNotFound,
			"Producto no encontrado",
		)
	}

	available := product.Stock - product.ReservedStock

	if quantity > available {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			"stock insuficiente",
		)
	}

	return s.IncrementReservedStock(
		ctx,
		tx,
		productID,
		quantity,
	)
}

func (s *inventoryService) ConsolidateStock(
	ctx context.Context,
	tx *sql.Tx,
	productID int,
	quantity int,
) error {

	query := `
		UPDATE product
		SET
			stock = stock - $1,
			reserved_stock = reserved_stock - $1
		WHERE id = $2
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		quantity,
		productID,
	)

	return err
}

func (r *inventoryService) IncrementReservedStock(
	ctx context.Context,
	tx *sql.Tx,
	productID int,
	quantity int,
) error {

	query := `
		UPDATE product
		SET reserved_stock = reserved_stock + $1
		WHERE id = $2
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		quantity,
		productID,
	)

	return err
}

func (s *inventoryService) AdjustReservedStock(
	ctx context.Context,
	tx *sql.Tx,
	productID int,
	diff int,
) error {

	if diff == 0 {
		return nil
	}

	if diff > 0 {
		return s.ReserveStock(
			ctx,
			tx,
			productID,
			diff,
		)
	}

	return s.ReleaseStock(
		ctx,
		tx,
		productID,
		-diff,
	)
}

func (s *inventoryService) ReleaseStock(
	ctx context.Context,
	tx *sql.Tx,
	productID int,
	quantity int,
) error {

	query := `
		UPDATE product
		SET reserved_stock = reserved_stock - $1
		WHERE id = $2
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		quantity,
		productID,
	)

	return err
}

func (s *inventoryService) FinalizeReservedStock(
	ctx context.Context,
	tx *sql.Tx,
	productID int,
	quantity int,
) error {
	query := `
		UPDATE product
		SET reserved_stock = reserved_stock - $1,
		    stock = stock - $1
		WHERE id = $2
	`
	_, err := tx.ExecContext(
		ctx,
		query,
		quantity,
		productID,
	)
	return err
}
