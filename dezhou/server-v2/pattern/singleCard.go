package pattern

import "server-v2/card"

type SingleCard struct {
	weight int
}

func NewSingleCard() SingleCard {
	return SingleCard{
		weight: 0,
	}
}

func (s SingleCard) Judge(cards []card.Card) (bool, []card.Card) {
	return true, cards[0:5]
}

func (s SingleCard) GetWeight() int {
	return s.weight
}
