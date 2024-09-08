package pattern

import (
	"server-v2/card"
	"server-v2/pattern/patternTool"
)

type FourOfKind struct {
	weight int
}

func NewFourOfKind() FourOfKind {
	return FourOfKind{
		weight: 7,
	}
}

// return：[四条的数, 单牌的数]
func (f FourOfKind) Judge(cards []card.Card) (bool, []card.Card) {
	flag, four := patternTool.HasNSame(cards, 4)
	if flag { // 如果有四条，找出非四条的最大的一张牌，一起返回
		for i := 0; i < len(cards); i++ {
			if cards[i].Number.Weight != four[0][0].Number.Weight {
				return true, append(four[0], cards[i])
			}
		}
	} else {
		return false, nil
	}
	return false, nil
}

func (f FourOfKind) GetWeight() int {
	return f.weight
}
