package pattern

import (
	"server-v2/card"
	"server-v2/pattern/patternTool"
)

type ThreeOfKind struct {
	weight int
}

func NewThreeOfKind() ThreeOfKind {
	return ThreeOfKind{
		weight: 3,
	}
}

// return [三条的值，单牌1，单牌2]
func (t ThreeOfKind) Judge(cards []card.Card) (bool, []card.Card) {
	// 找到三条
	flag, three := patternTool.HasNSame(cards, 3)
	if flag {
		// 找到与三条不相同的最大的两张牌
		var c1, c2 card.Card
		for i := 0; i < len(cards); i++ {
			if cards[i].Number.Weight != three[0][0].Number.Weight {
				c1 = cards[i]
			}
		}
		for i := 0; i < len(cards); i++ {
			if cards[i].Number.Weight != three[0][0].Number.Weight &&
				cards[i].Number.Weight != c1.Number.Weight {
				c2 = cards[i]
			}
		}
		return true, append(three[0], c1, c2)
	} else {
		return false, nil
	}
}

func (t ThreeOfKind) GetWeight() int {
	return t.weight
}
