package model

import "time"

type Product struct {
	ID            int        `db:"id"`
	CompanyID     int        `db:"company_id"`
	CompanyName   string     `db:"company_name"`
	Name          string     `db:"name"`
	Description   string     `db:"description"`
	Kind          string     `db:"kind"`
	TypeID        int        `db:"type_id"`
	Type          string     `db:"type"`
	CategoryID    int        `db:"category_id"`
	Category      string     `db:"category"`
	Price         float64    `db:"price"`
	Stock         int        `db:"stock"`
	ReservedStock int        `db:"reserved_stock"`
	ImagePath     *string    `db:"image_path"`
	Status        int        `db:"status"`
	CreatedAt     time.Time  `db:"created_at"`
	DeletedAt     *time.Time `db:"deleted_at"`
}
