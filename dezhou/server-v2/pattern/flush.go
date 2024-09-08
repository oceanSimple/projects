package pattern

import (
	"server-v2/card"
	"server-v2/pattern/patternTool"
)

type Flush struct {
	weight int
}

func NewFlush() Flush {
	return Flush{
		weight: 5,
	}
}

// return 最大的同花牌
func (f Flush) Judge(cards []card.Card) (bool, []card.Card) {
	flag, flushes := patternTool.HasFlush(cards)
	if flag {
		return true, flushes[0]
	} else {
		return false, nil
	}
}

func (f Flush) GetWeight() int {
	return f.weight
}
