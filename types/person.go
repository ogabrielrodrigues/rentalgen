package types

import (
	"fmt"
	"strings"
)

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
	genderStr := `
  %s, %s, %s, portador da cédula de identidade RG n.° %s, inscrito no CPF sob o n.° %s,
  residente e domiciliado a %s`

	if p.Nationality == "brasileira" {
		genderStr = `
		%s, %s, %s, portadora da cédula de identidade RG n.º %s, inscrita no CPF sob o n.° %s,
		residente e domiciliada a %s`
	}

	return fmt.Sprintf(
		genderStr,
		strings.ToLower(p.Nationality),
		strings.ToLower(p.MaritalStatus),
		strings.ToLower(p.Occupation),
		p.Rg,
		p.Cpf,
		p.Address,
	)
}
