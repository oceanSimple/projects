package pattern

import (
	"server-v2/card"
	"server-v2/pattern/patternTool"
)

type OnePair struct {
	weight int
}

func NewOnePair() OnePair {
	return OnePair{
		weight: 1,
	}
}

func (o OnePair) Judge(cards []card.Card) (bool, []card.Card) {
	// 获取所有的对子
	flag, two := patternTool.HasNSame(cards, 2)
	if !flag {
		return false, nil
	}
	// 从对子中获取最大的对子
	c := two[0][0]
	// 找出除对子外三张最大的单牌
	var c1, c2, c3 card.Card
	for i := 0; i < len(cards); i++ {
		if cards[i].Number.Weight != c.Number.Weight {
			c1 = cards[i]
		}
	}
	for i := 0; i < len(cards); i++ {
		if cards[i].Number.Weight != c.Number.Weight &&
			cards[i].Number.Weight != c1.Number.Weight {
			c2 = cards[i]
		}
	}
	for i := 0; i < len(cards); i++ {
		if cards[i].Number.Weight != c.Number.Weight &&
			cards[i].Number.Weight != c1.Number.Weight &&
			cards[i].Number.Weight != c2.Number.Weight {
			c3 = cards[i]
		}
	}
	return true, []card.Card{c, c, c1, c2, c3}
}

func (o OnePair) GetWeight() int {
	return o.weight
}
