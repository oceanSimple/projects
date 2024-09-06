package pattern

import "server-v1/card"

// 葫芦：三张同一点数的牌，加一对其他点数的牌
type FullHouse struct {
	weight int
}

func NewFullHouse() *FullHouse {
	return &FullHouse{
		weight: 6,
	}
}

func (f FullHouse) Judge(cards []*card.Card) (bool, []*card.Card) {
	// Has three same?
	threeFlag, threeSlices := HasThreeSame(cards)
	if !threeFlag {
		return false, nil
	}
	// Get two same: because has three same, so must have two same
	_, twoSlices := HasTwoSame(cards)
	// Assert the pair are not the same with three
	for i := 0; i < len(threeSlices); i++ {
		for j := 0; j < len(twoSlices); j++ {
			if threeSlices[i][0].Number.Weight != twoSlices[j][0].Number.Weight {
				return true, threeSlices[i]
			}
		}
	}
	return false, nil
}

func (f FullHouse) Compare(cards1 []*card.Card, cards2 []*card.Card) int {
	// Only need to compare the three same
	if cards1[0].Number.Weight > cards2[0].Number.Weight {
		return 1
	} else if cards1[0].Number.Weight < cards2[0].Number.Weight {
		return -1
	} else {
		return 0
	}
}

func (f FullHouse) GetWeight() int {
	return f.weight
}
