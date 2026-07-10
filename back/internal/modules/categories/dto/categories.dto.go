package categoriesdto

type CategoryItemResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetCategoriesResponse struct {
	ProductCategories []CategoryItemResponse `json:"product_categories"`
	CompanyCategories []CategoryItemResponse `json:"company_categories"`
}
