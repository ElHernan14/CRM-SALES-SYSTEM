package productdto

type GetCompanyProductsRequest struct {
	Page  int `json:"page" validate:"required,min=1"`
	Limit int `json:"limit" validate:"required,min=1,max=100"`

	Name string `json:"name" validate:"omitempty,max=50"`
	Type string `json:"type" validate:"omitempty,product_type"`

	Status int `json:"status" validate:"omitempty,min=0"`

	SortColumn string `json:"sort_column" validate:"omitempty,max=50,oneof=name description type price stock"`
	Order      string `json:"order" validate:"omitempty,oneof=asc desc"`
}
