package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"rentalgen/types"
)

func Parse(filePath string) (*types.Tmpl, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("✗ erro ao abrir arquivo de alugueres: %w", err)
	}
	defer file.Close()

	var tmplData types.Tmpl
	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&tmplData)
	if err != nil {
		return nil, fmt.Errorf("✗ erro ao processar arquivo: %w", err)
	}

	err = validate(&tmplData)
	if err != nil {
		return nil, fmt.Errorf("✗ arquivo inválido:\n%w", err)
	}

	for i := range tmplData.Rentals {
		rental := &tmplData.Rentals[i]

		rental.ReferenceMonth = tmplData.ReferenceDate.Month
		rental.ReferenceYear = tmplData.ReferenceDate.Year
		rental.ExpirationMonth = tmplData.ExpirationDate.Month
		rental.ExpirationYear = tmplData.ExpirationDate.Year
		rental.ReceiptDay = receiptDay(rental.DueDay, tmplData.ExpirationDate)

		rental.Owner.Qualification = rental.Owner.GenerateQualification()

		for j := range rental.Tenants {
			rental.Tenants[j].Qualification = rental.Tenants[j].GenerateQualification()
		}

		for j := range rental.Guarantors {
			rental.Guarantors[j].Qualification = rental.Guarantors[j].GenerateQualification()
		}
	}

	return &tmplData, nil
}

// receiptDay limita o dia de vencimento ao último dia do mês, evitando datas
// impossíveis como "31 de Fevereiro".
func receiptDay(dueDay string, date types.MonthYear) string {
	day, _ := strconv.Atoi(dueDay)
	year, _ := strconv.Atoi(date.Year)
	month, _ := parseMonth(date.Month)

	lastDay := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
	if day > lastDay {
		return strconv.Itoa(lastDay)
	}

	return dueDay
}
