package pattern

import "server-v1/card"

type Straight struct {
	weight int
}

func (s *Straight) GetWeight() int {
	return s.weight
}

func NewStraight() *Straight {
	return &Straight{weight: 4}
}

func (s *Straight) Judge(cards []*card.Card) (bool, []*card.Card) {
	sFlag, arr := HasStraight(cards)
	if sFlag {
		return true, arr[0]
	} else {
		return false, nil
	}
}

func (s *Straight) Compare(cards1 []*card.Card, cards2 []*card.Card) int {
	// Just compare the weight of the first card
	if cards1[0].Number.Weight > cards2[0].Number.Weight {
		return 1
	} else if cards1[0].Number.Weight < cards2[0].Number.Weight {
		return -1
	} else {
		return 0
	}
}
