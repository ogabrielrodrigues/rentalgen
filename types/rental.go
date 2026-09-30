package types

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
	OnlyVoucher          *bool     `json:"onlyVoucher"`
	Skip                 *bool     `json:"skip"`
}
