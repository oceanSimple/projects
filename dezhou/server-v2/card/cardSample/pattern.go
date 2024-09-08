package cardSample

import "server-v2/card"

// ♠A ♠K ♠Q ♠J ♠10 ♠9 ♥8
func RoyalStraightFlush() []card.Card {
	return []card.Card{
		card.GetCardByString("♠A"),
		card.GetCardByString("♠K"),
		card.GetCardByString("♠Q"),
		card.GetCardByString("♠J"),
		card.GetCardByString("♠10"),
		card.GetCardByString("♠9"),
		card.GetCardByString("♥8"),
	}
}

// ♠A ♥K ♠Q ♠5 ♠4 ♠3 ♠2
func StraightFlush() []card.Card {
	return []card.Card{
		card.GetCardByString("♠A"),
		card.GetCardByString("♥K"),
		card.GetCardByString("♠Q"),
		card.GetCardByString("♠5"),
		card.GetCardByString("♠4"),
		card.GetCardByString("♠3"),
		card.GetCardByString("♠2"),
	}
}
