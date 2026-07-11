package producttype

type ProductType struct {
	ID          int    `db:"id"`
	CategoryID  int    `db:"category_id"`
	Category    string `db:"category"`
	Name        string `db:"name"`
	Description string `db:"description"`
}
