package pattern

import "server-v1/card"

type ThreeOfKind struct {
	weight int
}

func NewThreeOfKind() *ThreeOfKind {
	return &ThreeOfKind{weight: 3}
}

func (t *ThreeOfKind) Judge(cards []*card.Card) (bool, []*card.Card) {
	tFlag, arr := HasThreeSame(cards)
	if tFlag {
		return true, arr[0]
	} else {
		return false, nil
	}
}

func (t *ThreeOfKind) Compare(cards1 []*card.Card, cards2 []*card.Card) int {
	// just compare the first card
	if cards1[0].Number.Weight > cards2[0].Number.Weight {
		return 1
	} else if cards1[0].Number.Weight < cards2[0].Number.Weight {
		return -1
	} else {
		return 0
	}
}

func (t *ThreeOfKind) GetWeight() int {
	return t.weight
}
