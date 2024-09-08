package patternTool

import "server-v2/card"

func HasNSame(cards []card.Card, n int) (bool, [][]card.Card) {
	var result [][]card.Card

	// 寻找n张相同的牌，即cards[i] = cards[i+i+(n-1)]的number
	for i := 0; i < len(cards)-(n-1); i++ {
		if cards[i].Number.Weight == cards[i+(n-1)].Number.Weight {
			// 可以直接浅拷贝
			result = append(result, cards[i:i+n])
		}
	}

	return len(result) > 0, result
}
