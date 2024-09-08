package patternTool

import (
	"server-v2/card"
	"server-v2/tool"
)

// 判断是否有同花
var HasFlush PatternToolInterface = func(cards []card.Card) (bool, [][]card.Card) {
	var result [][]card.Card

	// 首先将牌按花色排序（默认是按照Number.Weight排序的）
	tool.SortCardsByColorText(cards)
	defer tool.SortCardsByNumberWeight(cards) // 方法结束后恢复原来的排序

	// 判断cards[i]和cards[i+4]是否同花
	// 因此为了避免数组越界，我们只需要遍历到倒数第五张牌即可
	for i := 0; i < len(cards)-4; i++ {
		if cards[i].Color.Text == cards[i+4].Color.Text {
			// 因为最后我们要还原原来的排序，所以我们深拷贝同花的牌
			flush := make([]card.Card, 5)
			copy(flush, cards[i:i+5])
			result = append(result, flush)
		}
	}

	return len(result) > 0, result
}
