package types

import (
	"fmt"
	"strings"
)

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

type Person struct {
	Fullname      string `json:"fullname"`
	Nationality   string `json:"nationality"`
	Denomination  string `json:"denomination,omitempty"`
	MaritalStatus string `json:"maritalStatus"`
	Occupation    string `json:"occupation"`
	Rg            string `json:"rg"`
	Cpf           string `json:"cpf"`
	Address       string `json:"address"`
	Qualification string
}

func (p *Person) GenerateQualification() string {
	genderStr := "%s, %s, %s, portador da Carteira da Cédula de Identidade RG n.° %s, inscrito no CPF sob o n.° %s, residente e domiciliado a %s"

	if p.Nationality == "brasileira" {
		genderStr = "%s, %s, %s, portadora da Cédula de Identidade RG n.° %s, inscrita no CPF sob o n.° %s, residente e domiciliada a %s"
	}

	return fmt.Sprintf(genderStr, strings.ToLower(p.Nationality), strings.ToLower(p.MaritalStatus), strings.ToLower(p.Occupation), p.Rg, p.Cpf, p.Address)
}

type Rental struct {
	ReferenceMonth       *string   `json:"-"`
	ReferenceYear        *string   `json:"-"`
	ExpirationMonth      *string   `json:"-"`
	ExpirationYear       *string   `json:"-"`
	Owner                Person    `json:"owner"`
	Tenants              []Person  `json:"tenants"`
	Guarantors           *[]Person `json:"guarantors,omitempty"`
	Guarantee            string    `json:"guarantee"`
	City                 string    `json:"city"`
	PropertyAddress      string    `json:"propertyAddress"`
	Kind                 string    `json:"kind"`
	Amount               string    `json:"amount"`
	LongAmount           string    `json:"longAmount"`
	AmountWithoutTax     string    `json:"amountWithoutTax"`
	DueDay               string    `json:"dueDay"`
	SignedIn             string    `json:"signedIn"`
	ExtendedIn           string    `json:"extendedIn"`
	ExpiresIn            string    `json:"expiresIn"`
	VoucherObservation   string    `json:"voucherObservation,omitempty"`
	BordereauObservation string    `json:"bordereauObservation,omitempty"`
}
