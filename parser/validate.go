package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"rentalgen/types"
)

var months = map[string]time.Month{
	"janeiro":   time.January,
	"fevereiro": time.February,
	"março":     time.March,
	"abril":     time.April,
	"maio":      time.May,
	"junho":     time.June,
	"julho":     time.July,
	"agosto":    time.August,
	"setembro":  time.September,
	"outubro":   time.October,
	"novembro":  time.November,
	"dezembro":  time.December,
}

type field struct {
	name  string
	value string
}

func parseMonth(name string) (time.Month, bool) {
	month, ok := months[strings.ToLower(name)]
	return month, ok
}

func validate(tmplData *types.Tmpl) error {
	errs := append(
		validateMonthYear("referenceDate", tmplData.ReferenceDate),
		validateMonthYear("expirationDate", tmplData.ExpirationDate)...,
	)

	if len(tmplData.Rentals) == 0 {
		errs = append(errs, errors.New("  - rentals: nenhum aluguel informado"))
	}

	for i, rental := range tmplData.Rentals {
		if rental.Skip {
			continue
		}

		errs = append(errs, validateRental(fmt.Sprintf("aluguel %d", i+1), rental)...)
	}

	return errors.Join(errs...)
}

func validateMonthYear(name string, date types.MonthYear) []error {
	var errs []error

	if _, ok := parseMonth(date.Month); !ok {
		errs = append(errs, fmt.Errorf("  - %s.month: mês inválido %q", name, date.Month))
	}

	if year, err := strconv.Atoi(date.Year); err != nil || year < 1 {
		errs = append(errs, fmt.Errorf("  - %s.year: ano inválido %q", name, date.Year))
	}

	return errs
}

func validateRental(label string, rental types.Rental) []error {
	errs := requireFields(label, []field{
		{"city", rental.City},
		{"propertyAddress", rental.PropertyAddress},
		{"kind", rental.Kind},
		{"amount", rental.Amount},
		{"longAmount", rental.LongAmount},
	})

	if day, err := strconv.Atoi(rental.DueDay); err != nil || day < 1 || day > 31 {
		errs = append(errs, fmt.Errorf("  - %s: dueDay inválido %q", label, rental.DueDay))
	}

	errs = append(errs, validatePerson(label+", owner", rental.Owner)...)

	if len(rental.Tenants) == 0 {
		errs = append(errs, fmt.Errorf("  - %s: nenhum inquilino informado", label))
	}

	for i, tenant := range rental.Tenants {
		errs = append(errs, validatePerson(fmt.Sprintf("%s, tenants[%d]", label, i), tenant)...)
	}

	// Estes campos aparecem somente no borderô.
	if rental.OnlyVoucher {
		return errs
	}

	errs = append(errs, requireFields(label, []field{
		{"guarantee", rental.Guarantee},
		{"amountWithoutTax", rental.AmountWithoutTax},
		{"signedIn", rental.SignedIn},
		{"extendedIn", rental.ExtendedIn},
		{"expiresIn", rental.ExpiresIn},
	})...)

	for i, guarantor := range rental.Guarantors {
		errs = append(errs, validatePerson(fmt.Sprintf("%s, guarantors[%d]", label, i), guarantor)...)
	}

	return errs
}

func validatePerson(label string, person types.Person) []error {
	return requireFields(label, []field{
		{"fullname", person.Fullname},
		{"nationality", person.Nationality},
		{"maritalStatus", person.MaritalStatus},
		{"occupation", person.Occupation},
		{"rg", person.Rg},
		{"cpf", person.Cpf},
		{"address", person.Address},
	})
}

func requireFields(label string, fields []field) []error {
	var errs []error

	for _, f := range fields {
		if strings.TrimSpace(f.value) == "" {
			errs = append(errs, fmt.Errorf("  - %s: campo obrigatório %s não preenchido", label, f.name))
		}
	}

	return errs
}
