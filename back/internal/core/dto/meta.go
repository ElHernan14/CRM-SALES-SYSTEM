package dto

type Meta struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages,omitempty"`
}

func NewMeta(page int, limit int, total int) Meta {
	return Meta{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: CalculateTotalPages(total, limit),
	}
}

func CalculateTotalPages(total int, limit int) int {
	if total <= 0 || limit <= 0 {
		return 0
	}

	return (total + limit - 1) / limit
}
