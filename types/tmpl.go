package types

type Tmpl struct {
	ReferenceDate  MonthYear `json:"referenceDate"`
	ExpirationDate MonthYear `json:"expirationDate"`
	Rentals        []Rental  `json:"rentals"`
}

type MonthYear struct {
	Month string `json:"month"`
	Year  string `json:"year"`
}
