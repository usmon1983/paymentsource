package types

type Money int64

type Currency string

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type PAN string

type Card struct {
	ID int
	PAN PAN
	Balance Money
	MinBalance Money
	Currency Currency
	Color string
	Name string
	Active bool
	Type string
}

type Payment struct {
	ID int
	Amount Money
}

type PaymentSource struct {
	Type string
	Number string
	Balance Money
}