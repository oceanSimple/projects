package cardSample

import "server-v2/card"

// ♥K ♠A ♣10 ♦J ♠9 ♠8 ♠7
func TestSort() []card.Card {
	return []card.Card{
		card.GetCardByString("♥K"),
		card.GetCardByString("♠A"),
		card.GetCardByString("♣10"),
		card.GetCardByString("♦J"),
		card.GetCardByString("♠9"),
		card.GetCardByString("♠8"),
		card.GetCardByString("♠7"),
	}
}
