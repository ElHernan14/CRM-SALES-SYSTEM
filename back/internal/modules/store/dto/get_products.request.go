package dto

type GetStoreProductsRequest struct {
	Page  int `json:"page" validate:"required,min=1" example:"1"`
	Limit int `json:"limit" validate:"required,min=1,max=100" example:"10"`

	CompanyID *int    `json:"company_id" validate:"omitempty,min=1" example:"2"`
	Search    string  `json:"search" validate:"omitempty,max=50" example:"notebook"`
	Name      string  `json:"name" validate:"omitempty,max=50" example:"notebook"`
	Type      string  `json:"type" validate:"omitempty,product_type" example:"physical"`
	MinPrice  float64 `json:"min_price" validate:"omitempty,gte=0" example:"100"`
	MaxPrice  float64 `json:"max_price" validate:"omitempty,gte=0" example:"1500"`

	SortColumn string `json:"sort_column" validate:"omitempty,max=50,oneof=name created_at type price stock" example:"price"`
	Order      string `json:"order" validate:"omitempty,oneof=asc desc" example:"asc"`
}
