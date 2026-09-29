package types

type Tmpl struct {
	ReferenceDate  ReferenceDate  `json:"referenceDate"`
	ExpirationDate ExpirationDate `json:"expirationDate"`
	Rentals        []Rental       `json:"rentals"`
}

type ReferenceDate struct {
	Month string `json:"month"`
	Year  string `json:"year"`
}

type ExpirationDate struct {
	Month string `json:"month"`
	Year  string `json:"year"`
}
