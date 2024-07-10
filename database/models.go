package database

type User struct {
	ID             int    `json:"id"`
	LastName       string `json:"surname"`
	FirstName      string `json:"name"`
	Patronymic     string `json:"patronymic"`
	Address        string `json:"address"`
	PassportSerie  string `json:"passport_serie"`
	PassportNumber string `json:"passport_number"`
}
