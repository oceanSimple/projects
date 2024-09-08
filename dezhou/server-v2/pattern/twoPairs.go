package pattern

import (
	"server-v2/card"
	"server-v2/pattern/patternTool"
)

type TwoPairs struct {
	weight int
}

func NewTwoPairs() TwoPairs {
	return TwoPairs{
		weight: 2,
	}
}

func (t TwoPairs) Judge(cards []card.Card) (bool, []card.Card) {
	// 获取所有的对子
	flag, two := patternTool.HasNSame(cards, 2)
	if !flag || len(two) < 2 {
		return false, nil
	}
	// 从对子中获取最大的两个对子
	c1, c2 := two[0][0], two[1][0]
	// 找出除了两个对子之外的最大单牌
	for i := 0; i < len(cards); i++ {
		if cards[i].Number.Weight != c1.Number.Weight &&
			cards[i].Number.Weight != c2.Number.Weight {
			return true, []card.Card{c1, c1, c2, c2, cards[i]}
		}
	}
	return false, nil
}

func (t TwoPairs) GetWeight() int {
	return t.weight
}
