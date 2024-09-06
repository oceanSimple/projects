package pattern

import (
	"server-v1/card"
)

type TwoPairs struct {
	weight int
}

func NewTwoPairs() TwoPairs {
	return TwoPairs{weight: 2}
}

func (t TwoPairs) Judge(cards []*card.Card) (bool, []*card.Card) {
	tFlag, arr := HasTwoSame(cards)
	if tFlag {
		if len(arr) >= 2 {
			// find the single card that is not in the two same cards
			for _, c := range cards {
				if c.Number.Weight != arr[0][0].Number.Weight &&
					c.Number.Weight != arr[1][0].Number.Weight {
					return true, []*card.Card{arr[0][0], arr[1][0], c}
				}
			}
		}
	}
	return false, nil
}

func (t TwoPairs) Compare(cards1 []*card.Card, cards2 []*card.Card) int {
	// [first pair, second pair, single card]
	for i := 0; i < 3; i++ {
		if cards1[i].Number.Weight > cards2[i].Number.Weight {
			return 1
		} else if cards1[i].Number.Weight < cards2[i].Number.Weight {
			return -1
		}
	}
	return 0
}

func (t TwoPairs) GetWeight() int {
	return t.weight
}
