package model

import "time"

type Product struct {
	ID            int        `db:"id"`
	CompanyID     int        `db:"company_id"`
	Name          string     `db:"name"`
	Description   string     `db:"description"`
	Type          string     `db:"type"`
	Price         float64    `db:"price"`
	Stock         int        `db:"stock"`
	ReservedStock int        `db:"reserved_stock"`
	Status        int        `db:"status"`
	CreatedAt     time.Time  `db:"created_at"`
	DeletedAt     *time.Time `db:"deleted_at"`
}
