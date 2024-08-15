package domain

type Car struct {
	Id    string `json:"id"`
	Year  int    `json:"year"`
	Model string `json:"model"`
	Brand string `json:"brand"`
	Color string `json:"color"`
}
