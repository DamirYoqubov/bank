package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleTotal() {

	cards := []types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
		},
	}
	fmt.Println(Total(cards))
	//Output: 100000
}

func ExamplePaymentSource() {
	cards := []types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
			PAN:     "5058",
		},
		{
			Balance: 2_000_00,
			Active:  true,
			PAN:     "ery",
		},
	}

	paySource := PaymentSource(cards)
	fmt.Println(len(paySource))
	// Output: 2
}
