package pattern

import (
	"server-v1/card"
)

// 同花顺：同一花色的顺子
type StraightFlush struct {
	weight int
}

func NewStraightFlush() *StraightFlush {
	return &StraightFlush{
		weight: 8,
	}
}

func (r *StraightFlush) Judge(cards []*card.Card) (bool, []*card.Card) {
	// Has straight?
	sFlag, straights := HasStraight(cards)
	if !sFlag {
		return false, nil
	}
	// Judge straights is flush?
	flushes := r.IsFlush(straights)
	if len(flushes) == 0 {
		return false, nil
	}
	return true, flushes[0]
}

func (r *StraightFlush) Compare(cards1 []*card.Card, cards2 []*card.Card) int {
	// Only need to compare the first card
	if cards1[0].Number.Weight > cards2[0].Number.Weight {
		return 1
	} else if cards1[0].Number.Weight < cards2[0].Number.Weight {
		return -1
	} else {
		return 0
	}
}

func (r *StraightFlush) GetWeight() int {
	return r.weight
}
func (r *StraightFlush) IsFlush(cards [][]*card.Card) [][]*card.Card {
	var result [][]*card.Card
	for _, value := range cards {
		isFlush := true
		for i := 1; i < len(value); i++ {
			if value[i].Color.Text != value[0].Color.Text {
				isFlush = false
				break
			}
		}
		if isFlush {
			result = append(result, value)
		}
	}
	return result
}
