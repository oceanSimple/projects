package pattern

import "server-v1/card"

type SingleCard struct {
	weight int
}

func (s *SingleCard) GetWeight() int {
	return s.weight
}

func NewSingleCard() *SingleCard {
	return &SingleCard{weight: 0}
}

func (s *SingleCard) Judge(cards []*card.Card) (bool, []*card.Card) {
	// return the first five cards
	return true, cards[0:5]
}

func (s *SingleCard) Compare(cards1 []*card.Card, cards2 []*card.Card) int {
	for i := 0; i < 5; i++ {
		if cards1[i].Number.Weight > cards2[i].Number.Weight {
			return 1
		} else if cards1[i].Number.Weight < cards2[i].Number.Weight {
			return -1
		}
	}
	return 0
}
