package productdto

import meta "crm-system-sales/internal/core/dto"

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required,min=2,max=100" example:"Notebook Lenovo ThinkPad"`
	Description string  `json:"description" validate:"omitempty,max=500" example:"Business laptop with Intel Core i7 processor"`
	Kind        string  `json:"kind" validate:"required,product_type" example:"product"`
	CategoryID  int     `json:"category_id" validate:"required,gt=0" example:"1"`
	TypeID      int     `json:"type_id" validate:"required,gt=0" example:"3"`
	Price       float64 `json:"price" validate:"required,gt=0" example:"1200.50"`
	Stock       int     `json:"stock" validate:"gte=0" example:"50"`
	CompanyID   *int    `json:"company_id,omitempty" example:"1"`
	ImagePath   *string `json:"image_path,omitempty" validate:"omitempty,max=500" example:"/uploads/product/images/product_1.png"`
}

type ProductResponse struct {
	ID             int     `json:"id" example:"1"`
	Name           string  `json:"name" example:"Notebook Lenovo ThinkPad"`
	Description    string  `json:"description,omitempty" example:"Business laptop with Intel Core i7 processor"`
	Kind           string  `json:"kind" example:"product"`
	CategoryID     int     `json:"category_id" example:"1"`
	Category       string  `json:"category" example:"Technology"`
	TypeID         int     `json:"type_id" example:"3"`
	Type           string  `json:"type" example:"Notebook"`
	Price          float64 `json:"price" example:"1200.50"`
	Stock          int     `json:"stock" example:"50"`
	ReservedStock  int     `json:"reserved_stock" example:"5"`
	AvailableStock int     `json:"available_stock" example:"45"`
	ImagePath      *string `json:"image_path,omitempty" example:"/uploads/product/images/product_1.png"`
}

type UploadProductImageResponse struct {
	ImagePath string `json:"image_path"`
}

type GetProductsRequest struct {
	Search     string  `validate:"omitempty,min=2,max=100" example:"notebook"`
	Kind       string  `validate:"omitempty,product_type" example:"product"`
	CategoryID *int    `validate:"omitempty,gt=0" example:"1"`
	Category   string  `validate:"omitempty,min=2,max=80" example:"Technology"`
	TypeID     *int    `validate:"omitempty,gt=0" example:"3"`
	Type       string  `validate:"omitempty,min=2,max=80" example:"Notebook"`
	MinPrice   float64 `validate:"omitempty,gte=0" example:"100"`
	MaxPrice   float64 `validate:"omitempty,gte=0" example:"5000"`
	CompanyID  *int    `validate:"omitempty,gt=0" example:"1"`
	Page       int     `validate:"gte=1" example:"1"`
	Limit      int     `validate:"gte=1,lte=100" example:"10"`
}

type ProductListItem struct {
	ID             int     `json:"id" example:"1"`
	Name           string  `json:"name" example:"Notebook Lenovo ThinkPad"`
	Description    string  `json:"description" example:"Business laptop with Intel Core i7 processor"`
	Kind           string  `json:"kind" example:"product"`
	CategoryID     int     `json:"category_id" example:"1"`
	Category       string  `json:"category" example:"Technology"`
	TypeID         int     `json:"type_id" example:"3"`
	Type           string  `json:"type" example:"Notebook"`
	Price          float64 `json:"price" example:"1200.50"`
	Stock          int     `json:"stock" example:"50"`
	Status         int     `json:"status" example:"1"`
	CompanyID      int     `json:"company_id" example:"1"`
	AvailableStock int     `json:"available_stock" example:"25"`
	ImagePath      *string `json:"image_path,omitempty" example:"/uploads/product/images/product_1.png"`
}

type GetProductsResponse struct {
	Data []ProductListItem `json:"data"`
	Meta meta.Meta         `json:"meta"`
}

type ProductDetailResponse struct {
	ID          int     `json:"id" example:"1"`
	Name        string  `json:"name" example:"Notebook Lenovo ThinkPad"`
	Description string  `json:"description" example:"Business laptop with Intel Core i7 processor"`
	Kind        string  `json:"kind" example:"product"`
	CategoryID  int     `json:"category_id" example:"1"`
	Category    string  `json:"category" example:"Technology"`
	TypeID      int     `json:"type_id" example:"3"`
	Type        string  `json:"type" example:"Notebook"`
	Price       float64 `json:"price" example:"1200.50"`
	Stock       int     `json:"stock" example:"50"`
	Status      int     `json:"status" example:"1"`
	CompanyID   int     `json:"company_id" example:"1"`
	ImagePath   *string `json:"image_path,omitempty" example:"/uploads/product/images/product_1.png"`
}

type UpdateProductRequest struct {
	Name        *string  `json:"name" validate:"omitempty,min=2,max=100" example:"Notebook Lenovo ThinkPad Gen 2"`
	Description *string  `json:"description" validate:"omitempty,max=500" example:"Updated business laptop"`
	Kind        *string  `json:"kind" validate:"omitempty,product_type" example:"product"`
	CategoryID  *int     `json:"category_id" validate:"omitempty,gt=0" example:"1"`
	TypeID      *int     `json:"type_id" validate:"omitempty,gt=0" example:"3"`
	Price       *float64 `json:"price" validate:"omitempty,gte=0" example:"1350.75"`
	Stock       *int     `json:"stock" validate:"omitempty,gte=0" example:"75"`
	Status      *int     `json:"status" validate:"omitempty,oneof=0 1" example:"1"`
}
