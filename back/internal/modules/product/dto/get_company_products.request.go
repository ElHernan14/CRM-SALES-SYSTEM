package productdto

type GetCompanyProductsRequest struct {
	Page  int `json:"page" validate:"required,min=1"`
	Limit int `json:"limit" validate:"required,min=1,max=100"`

	Name       string `json:"name" validate:"omitempty,max=50"`
	Kind       string `json:"kind" validate:"omitempty,product_type"`
	CategoryID *int   `json:"category_id" validate:"omitempty,gt=0"`
	Category   string `json:"category" validate:"omitempty,min=2,max=80"`
	TypeID     *int   `json:"type_id" validate:"omitempty,gt=0"`
	Type       string `json:"type" validate:"omitempty,min=2,max=80"`

	Status int `json:"status" validate:"omitempty,min=0"`

	SortColumn string `json:"sort_column" validate:"omitempty,max=50,oneof=name description kind type category price stock"`
	Order      string `json:"order" validate:"omitempty,oneof=asc desc"`
}
