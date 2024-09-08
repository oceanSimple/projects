package pattern

import (
	"server-v2/card"
	"server-v2/pattern/patternTool"
)

type Straight struct {
	weight int
}

func NewStraight() Straight {
	return Straight{
		weight: 4,
	}
}

// return 最大的顺子牌
func (s Straight) Judge(cards []card.Card) (bool, []card.Card) {
	flag, straights := patternTool.HasStraight(cards)
	if flag {
		return true, straights[0]
	} else {
		return false, nil
	}
}

func (s Straight) GetWeight() int {
	//TODO implement me
	panic("implement me")
}
