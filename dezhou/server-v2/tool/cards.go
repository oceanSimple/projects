package tool

import (
	"server-v2/card"
	"sort"
	"strings"
)

// 根据card的Number.Weight进行排序.
// 从大到小排序.
func SortCardsByNumberWeight(cards []card.Card) {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Number.Weight > cards[j].Number.Weight
	})
}

// 根据card的Number.ExtraWeight进行排序.
// 从大到小排序.(A最小)
func SortCardsByNumberExtraWeight(cards []card.Card) {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Number.ExtraWeight > cards[j].Number.ExtraWeight
	})
}

// 根据card的Color.Text进行排序.
// 我们只需要相同花色在一起即可，顺序不重要.
func SortCardsByColorText(cards []card.Card) {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Color.Text < cards[j].Color.Text
	})
}

// 判断cards是否相同
func EqualCards(cards1 []card.Card, cards2 []card.Card) bool {
	if len(cards1) != len(cards2) {
		return false
	}

	for i := 0; i < len(cards1); i++ {
		if !cards1[i].Equal(cards2[i]) {
			return false
		}
	}

	return true
}

// 比较两个符合当前牌型的牌的大小
// params： cards1 第一个符合当前牌型的牌
// params： cards2 第二个符合当前牌型的牌
// return： int 1：cards1大，-1：cards2大，0：相等, -2: 长度不同
func CompareCards(cards1 []card.Card, cards2 []card.Card) int {
	if len(cards1) != len(cards2) {
		return -2
	}

	for i := 0; i < len(cards1); i++ {
		if cards1[i].Number.Weight > cards2[i].Number.Weight {
			return 1
		} else if cards1[i].Number.Weight < cards2[i].Number.Weight {
			return -1
		}
	}

	return 0
}

func PrintCards(cards []card.Card) string {
	var strBuilder strings.Builder
	for i := 0; i < len(cards); i++ {
		strBuilder.WriteString(cards[i].String())
		strBuilder.WriteString(" ")
	}
	return strBuilder.String()
}
