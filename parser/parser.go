package parser

import (
	"encoding/json"
	"fmt"
	"os"

	"rentalgen/config"
	"rentalgen/logger"
	"rentalgen/types"
)

func Parse(flags *config.Flags) (*types.Tmpl, error) {
	file, err := os.Open(flags.FilePath)
	if err != nil {
		return nil, fmt.Errorf("✗ erro ao abrir arquivo de alugueres.")
	}
	defer file.Close()

	var tmplData types.Tmpl
	err = json.NewDecoder(file).Decode(&tmplData)
	if err != nil {
		return nil, fmt.Errorf("✗ erro ao processar arquivo.")
	}

	rentals := []types.Rental{}
	for _, rental := range tmplData.Rentals {
		rental.ReferenceMonth = &tmplData.ReferenceDate.Month
		rental.ReferenceYear = &tmplData.ReferenceDate.Year
		rental.ExpirationMonth = &tmplData.ExpirationDate.Month
		rental.ExpirationYear = &tmplData.ExpirationDate.Year

		rental.Owner.Qualification = rental.Owner.GenerateQualification()

		tenants := []types.Person{}
		for _, tenant := range rental.Tenants {
			tenant.Qualification = tenant.GenerateQualification()

			tenants = append(tenants, tenant)
		}

		guarantors := []types.Person{}
		if rental.Guarantors != nil {
			for _, guarantor := range *rental.Guarantors {
				guarantor.Qualification = guarantor.GenerateQualification()

				guarantors = append(guarantors, guarantor)
			}
		}

		rental.Tenants = tenants
		rental.Guarantors = &guarantors
		rentals = append(rentals, rental)
	}

	logger.Log(logger.ColorGreen, "✓ arquivo de alugueres selecionado: ")
	logger.Logln(logger.ColorDefault, flags.FilePath)
	logger.Logln(logger.ColorDefault, "✓ para realizar a impressão dos recibos, acesse http://localhost:8080\n")
	logger.Logln(logger.ColorDefault, "Para encerrar, pressione Ctrl + C")

	tmplData.Rentals = rentals
	return &tmplData, nil
}
