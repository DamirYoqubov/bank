package card

import (
	"bank/pkg/bank/types"
)

func Total(cards []types.Card) types.Money {
	sum := types.Money(0)
	for _, card := range cards {
		if !card.Active {
			continue
		}

		if card.Balance <= 0 {
			continue
		}
		
		sum += card.Balance
	}
	
	return sum
}

func PaymentSource(cards []types.Card) []types.PaymentSource {
	paySources:=make([]types.PaymentSource, 0, len(cards))
	for _, card := range cards {
		if card.Balance <= 0 {
			continue
		}
		if !card.Active {
			continue
		}	
		paySources = append(paySources, types.PaymentSource{
			Type: "card",
			Number: string(card.PAN),
			Balance: card.Balance,
		})
	}
	return paySources
}