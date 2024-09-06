package pattern

import "server-v1/card"

type OnePair struct {
	weight int
}

func (o *OnePair) GetWeight() int {
	return o.weight
}

func NewOnePair() *OnePair {
	return &OnePair{weight: 1}
}

func (o *OnePair) Judge(cards []*card.Card) (bool, []*card.Card) {
	var result []*card.Card
	oFlag, arr := HasTwoSame(cards)
	if oFlag {
		result = append(result, arr[0][0]) // add pair
		// add the rest three single cards
		var count = 0
		for _, c := range cards {
			if c.Number.Weight != arr[0][0].Number.Weight {
				result = append(result, c)
				count++
			}
			if count == 3 {
				return true, result
			}
		}
	}
	return false, nil
}

func (o *OnePair) Compare(cards1 []*card.Card, cards2 []*card.Card) int {
	// [pair, single card1, single card2, single card3]
	for i := 0; i < 4; i++ {
		if cards1[i].Number.Weight > cards2[i].Number.Weight {
			return 1
		} else if cards1[i].Number.Weight < cards2[i].Number.Weight {
			return -1
		}
	}
	return 0
}
