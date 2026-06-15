package dto

type GetStoreProductsRequest struct {
	Page  int `json:"page" validate:"required,min=1" example:"1"`
	Limit int `json:"limit" validate:"required,min=1,max=100" example:"10"`

	Name string `json:"name" validate:"omitempty,max=50" example:"notebook"`
	Type string `json:"type" validate:"omitempty,product_type" example:"physical"`

	SortColumn string `json:"sort_column" validate:"omitempty,max=50,oneof=name created_at type price stock" example:"price"`
	Order      string `json:"order" validate:"omitempty,oneof=asc desc" example:"asc"`
}
