package categoriesdto

type CategoryItemResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProductTypeItemResponse struct {
	ID          int    `json:"id"`
	CategoryID  int    `json:"category_id"`
	Category    string `json:"category"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetCategoriesResponse struct {
	ProductCategories []CategoryItemResponse    `json:"product_categories"`
	CompanyCategories []CategoryItemResponse    `json:"company_categories"`
	ProductTypes      []ProductTypeItemResponse `json:"product_types"`
}

type GetProductTypesResponse struct {
	Items []ProductTypeItemResponse `json:"items"`
}
