package cardSample

import "server-v1/card"

// 皇家同花顺：♠A-♥A-♠K-♥K-♠Q-♠J-♠10
func RoyalFlushStraight() []*card.Card {
	c := []*card.Card{
		card.GetCardByString("1-1"),
		card.GetCardByString("2-1"),
		card.GetCardByString("1-13"),
		card.GetCardByString("2-13"),
		card.GetCardByString("1-12"),
		card.GetCardByString("1-11"),
		card.GetCardByString("1-10"),
	}
	return c
}

// 皇家同花顺：♠A-♠K-♥K-♥Q-♠Q-♠J-♠10
func RoyalFlushStraight1() []*card.Card {
	c := []*card.Card{
		card.GetCardByString("1-1"),
		card.GetCardByString("1-13"),
		card.GetCardByString("2-13"),
		card.GetCardByString("2-12"),
		card.GetCardByString("1-12"),
		card.GetCardByString("1-11"),
		card.GetCardByString("1-10"),
	}
	return c
}

// 最小同花顺：黑桃♠A-♥A-♠5-♥4-♠4-♠3-♠2
func StraightFlushMin() []*card.Card {
	c := []*card.Card{
		card.GetCardByString("1-1"),
		card.GetCardByString("2-1"),
		card.GetCardByString("1-5"),
		card.GetCardByString("2-4"),
		card.GetCardByString("1-4"),
		card.GetCardByString("1-3"),
		card.GetCardByString("1-2"),
	}
	return c
}

// 同花顺：♠9-♠8-♠7-♠6-♠5-♠4-♠3
func StraightFlush() []*card.Card {
	c := []*card.Card{
		card.GetCardByString("1-9"),
		card.GetCardByString("1-8"),
		card.GetCardByString("1-7"),
		card.GetCardByString("1-6"),
		card.GetCardByString("1-5"),
		card.GetCardByString("1-4"),
		card.GetCardByString("1-3"),
	}
	return c
}

// 四条： ♠A-♥A-♥K-♠K-♦K-♣K-♠Q
func FourOfAKind() []*card.Card {
	c := []*card.Card{
		card.GetCardByString("1-1"),
		card.GetCardByString("2-1"),
		card.GetCardByString("2-13"),
		card.GetCardByString("1-13"),
		card.GetCardByString("3-13"),
		card.GetCardByString("4-13"),
		card.GetCardByString("1-12"),
	}
	return c
}

// 葫芦：♠A-♥A-♦A-♥K-♠K-♦K-♣Q
func FullHouse() []*card.Card {
	c := []*card.Card{
		card.GetCardByString("1-1"),
		card.GetCardByString("2-1"),
		card.GetCardByString("3-1"),
		card.GetCardByString("2-13"),
		card.GetCardByString("1-13"),
		card.GetCardByString("3-13"),
		card.GetCardByString("4-12"),
	}
	return c
}

// 非顺子的同花：♠K-♠J-♥10-♠8-♦5-♠4-♠3
func Flush() []*card.Card {
	c := []*card.Card{
		card.GetCardByString("1-13"),
		card.GetCardByString("1-11"),
		card.GetCardByString("2-10"),
		card.GetCardByString("1-8"),
		card.GetCardByString("3-5"),
		card.GetCardByString("1-4"),
		card.GetCardByString("1-3"),
	}
	return c
}

// 非同花的顺子：♠A-♥K-♠Q-♦J-♥10-♣9-♠8
func Straight() []*card.Card {
	c := []*card.Card{
		card.GetCardByString("1-1"),
		card.GetCardByString("2-13"),
		card.GetCardByString("1-12"),
		card.GetCardByString("3-11"),
		card.GetCardByString("2-10"),
		card.GetCardByString("4-9"),
		card.GetCardByString("1-8"),
	}
	return c
}

// 三条：♠A-♥A-♦A-♠K-♣Q-♠J-♠7
func ThreeOfAKind() []*card.Card {
	c := []*card.Card{
		card.GetCardByString("1-1"),
		card.GetCardByString("2-1"),
		card.GetCardByString("3-1"),
		card.GetCardByString("1-13"),
		card.GetCardByString("4-12"),
		card.GetCardByString("1-11"),
		card.GetCardByString("1-7"),
	}
	return c
}

// 两对：♠A-♥A-♦K-♠K-♣Q-♠J-♠7
func TwoPairs() []*card.Card {
	c := []*card.Card{
		card.GetCardByString("1-1"),
		card.GetCardByString("2-1"),
		card.GetCardByString("3-13"),
		card.GetCardByString("1-13"),
		card.GetCardByString("4-12"),
		card.GetCardByString("1-11"),
		card.GetCardByString("1-7"),
	}
	return c
}

// 一对：♠A-♥A-♦K-♠Q-♥J-♠4-♠3
func OnePair() []*card.Card {
	c := []*card.Card{
		card.GetCardByString("1-1"),
		card.GetCardByString("2-1"),
		card.GetCardByString("3-13"),
		card.GetCardByString("1-12"),
		card.GetCardByString("2-11"),
		card.GetCardByString("1-4"),
		card.GetCardByString("1-3"),
	}
	return c
}

// 散排：♠A-♥Q-♦J-♠8-♦6-♠3-♠2
func HighCard() []*card.Card {
	c := []*card.Card{
		card.GetCardByString("1-1"),
		card.GetCardByString("2-12"),
		card.GetCardByString("3-11"),
		card.GetCardByString("1-8"),
		card.GetCardByString("4-6"),
		card.GetCardByString("1-3"),
		card.GetCardByString("1-2"),
	}
	return c
}
