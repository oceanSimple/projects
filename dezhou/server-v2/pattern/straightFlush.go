package pattern

import (
	"server-v2/card"
	"server-v2/pattern/patternTool"
)

// 同花顺
type StraightFlush struct {
	weight int
}

func NewStraightFlush() StraightFlush {
	return StraightFlush{
		weight: 8,
	}
}

// return 最大的同花顺
func (s StraightFlush) Judge(cards []card.Card) (bool, []card.Card) {
	// 拿到所有的顺子
	flag, straights := patternTool.HasStraight(cards)
	if !flag {
		return false, nil
	}
	// 判断这些顺子是否是同花
	for i := 0; i < len(straights); i++ {
		flag, flush := patternTool.HasFlush(straights[i])
		if flag { // 只需要找到第一个，也即最大的同花顺
			return true, flush[0]
		}
	}
	return false, nil
}

func (s StraightFlush) GetWeight() int {
	return s.weight
}
