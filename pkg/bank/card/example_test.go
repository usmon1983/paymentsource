package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {
	cards := []types.Card {
		{
			PAN: "5058 xxxx xxxx 8888",
			Balance: 20_00_00,
			Active: true,
		},
		{
			PAN: "5058 xxxx xxxx 7777",
			Balance: -10_000_00,
			Active: true,
		},
		{
			PAN: "5058 xxxx xxxx 9999",
			Balance: 15_000_00,
			Active: false,
		},
	}
	paymentSources := PaymentSources(cards)

	for _, payment := range paymentSources {
		fmt.Println(payment.Number)
	}

	//Output: 5058 xxxx xxxx 8888
}