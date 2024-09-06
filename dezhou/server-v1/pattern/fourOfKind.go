package pattern

import "server-v1/card"

// 四条：有四张同一点数的牌
type FourOfKind struct {
	weight int
}

func NewFourOfKind() *FourOfKind {
	return &FourOfKind{
		weight: 7,
	}
}

// Return 4 cards if it has four of kind
func (r *FourOfKind) Judge(cards []*card.Card) (bool, []*card.Card) {
	done, values := HasFourSame(cards)
	if done {
		return true, values[0]
	}
	return false, nil
}

func (r *FourOfKind) Compare(cards1 []*card.Card, cards2 []*card.Card) int {
	// Only need to compare the first card
	if cards1[0].Number.Weight > cards2[0].Number.Weight {
		return 1
	} else if cards1[0].Number.Weight < cards2[0].Number.Weight {
		return -1
	} else {
		return 0
	}
}

func (r *FourOfKind) GetWeight() int {
	return r.weight
}
