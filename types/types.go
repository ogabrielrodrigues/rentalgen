package types

type Tmpl struct {
	ReferenceDate  ReferenceDate  `json:"referenceDate"`
	ExpirationDate ExpirationDate `json:"expirationDate"`
	Contracts      []Contract     `json:"contracts"`
}

type ReferenceDate struct {
	Month string `json:"month"`
	Year  string `json:"year"`
}

type ExpirationDate struct {
	Month string `json:"month"`
	Year  string `json:"year"`
}

type Contract struct {
	Owner                Person    `json:"owner"`
	Tenants              []Person  `json:"tenants"`
	Guarantors           *[]Person `json:"guarantors,omitempty"`
	Guarantee            string    `json:"guarantee"`
	Amount               string    `json:"amount"`
	LongAmount           string    `json:"longAmount"`
	AmountWithoutTax     string    `json:"amountWithoutTax"`
	Kind                 string    `json:"kind"`
	PropertyAddress      string    `json:"propertyAddress"`
	DueDay               string    `json:"dueDay"`
	SignedIn             string    `json:"signedIn"`
	ExpiresIn            string    `json:"expiresIn"`
	ExtendedIn           string    `json:"extendedIn"`
	VoucherObservation   string    `json:"voucherObservation,omitempty"`
	BordereauObservation string    `json:"bordereauObservation,omitempty"`
}

type Person struct {
	Fullname      string `json:"fullname"`
	Nationality   string `json:"nationality"`
	Denomination  string `json:"denomination,omitempty"`
	MaritalStatus string `json:"maritalStatus"`
	Occupation    string `json:"occupation"`
	RG            string `json:"rg"`
	CPF           string `json:"cpf"`
	Address       string `json:"address"`
}
