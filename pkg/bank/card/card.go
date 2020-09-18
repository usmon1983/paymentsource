package card

import (
	"bank/pkg/bank/types"
)

func PaymentSources(cards []types.Card) []types.PaymentSource  {
	paymentCard := []types.PaymentSource{}

	for _, card := range cards {
		if (!card.Active || card.Balance <= 0) {
			continue
		}
		paymentCard = append(paymentCard, types.PaymentSource{Number:string(card.PAN), Balance: card.Balance})
	}
	return paymentCard
}