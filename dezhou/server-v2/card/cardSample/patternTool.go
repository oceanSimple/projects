package cardSample

import "server-v2/card"

// ♠A ♥Q ♠10 ♥10 ♠9 ♠8 ♠7
func TestHasFlush() []card.Card {
	return []card.Card{
		card.GetCardByString("♠A"),
		card.GetCardByString("♥Q"),
		card.GetCardByString("♠10"),
		card.GetCardByString("♥10"),
		card.GetCardByString("♠9"),
		card.GetCardByString("♠8"),
		card.GetCardByString("♠7"),
	}
}

// ♠A ♥Q ♠10 ♥10 ♣10 ♦10 ♠7
func TestHasFourSame() []card.Card {
	return []card.Card{
		card.GetCardByString("♠A"),
		card.GetCardByString("♥Q"),
		card.GetCardByString("♠10"),
		card.GetCardByString("♥10"),
		card.GetCardByString("♣10"),
		card.GetCardByString("♦10"),
		card.GetCardByString("♠7"),
	}
}

// ♠A ♥Q ♠10 ♥10 ♣10 ♦9 ♠7
func TestHasThreeSame() []card.Card {
	return []card.Card{
		card.GetCardByString("♠A"),
		card.GetCardByString("♥Q"),
		card.GetCardByString("♠10"),
		card.GetCardByString("♥10"),
		card.GetCardByString("♣10"),
		card.GetCardByString("♦9"),
		card.GetCardByString("♠7"),
	}
}

// ♠A ♥A ♠K ♥K ♠Q ♠J ♠10
func TestHasStraight() []card.Card {
	return []card.Card{
		card.GetCardByString("♠A"),
		card.GetCardByString("♥A"),
		card.GetCardByString("♠K"),
		card.GetCardByString("♥K"),
		card.GetCardByString("♠Q"),
		card.GetCardByString("♠J"),
		card.GetCardByString("♠10"),
	}
}

// ♠A ♠6 ♠5 ♠4 ♠3 ♠2 ♥2
func TestHasSpecialStraight() []card.Card {
	return []card.Card{
		card.GetCardByString("♠A"),
		card.GetCardByString("♠6"),
		card.GetCardByString("♠5"),
		card.GetCardByString("♠4"),
		card.GetCardByString("♠3"),
		card.GetCardByString("♠2"),
		card.GetCardByString("♥2"),
	}
}
