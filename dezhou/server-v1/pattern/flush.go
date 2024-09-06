package pattern

import "server-v1/card"

type Flush struct {
	weight int
}

func (f *Flush) Judge(cards []*card.Card) (bool, []*card.Card) {
	fFlag, arr := HasFlush(cards)
	if fFlag {
		return true, arr[0]
	} else {
		return false, nil
	}
}

func (f *Flush) Compare(cards1 []*card.Card, cards2 []*card.Card) int {
	// from index 0 to 4, compare the weight of the cards
	for i := 0; i < len(cards1); i++ {
		if cards1[i].Number.Weight > cards2[i].Number.Weight {
			return 1
		} else if cards1[i].Number.Weight < cards2[i].Number.Weight {
			return -1
		} else {
			continue
		}
	}
	return 0
}

func NewFlush() *Flush {
	return &Flush{weight: 5}
}

func (f *Flush) GetWeight() int {
	return f.weight
}
