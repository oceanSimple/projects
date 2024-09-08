package pattern

import (
	"server-v2/card"
	"server-v2/pattern/patternTool"
)

type FullHouse struct {
	weight int
}

func NewFullHouse() FullHouse {
	return FullHouse{
		weight: 6,
	}
}

// return：[三条的数, 对子的数]
func (f FullHouse) Judge(cards []card.Card) (bool, []card.Card) {
	// 先拿到三条
	flag, three := patternTool.HasNSame(cards, 3)
	if !flag {
		return false, nil
	}
	// 再拿到对子(因为有三天了，所以肯定有对子)
	_, pair := patternTool.HasNSame(cards, 2)

	// 找到三条和对子不相同的，就是葫芦
	for i := 0; i < len(pair); i++ {
		if three[0][0].Number.Weight != pair[i][0].Number.Weight {
			return true, append(three[0], pair[i]...)
		}
	}
	return false, nil
}

func (f FullHouse) GetWeight() int {
	return f.weight
}
