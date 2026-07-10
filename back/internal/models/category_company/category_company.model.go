package categorycompany

type CategoryCompany struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}
