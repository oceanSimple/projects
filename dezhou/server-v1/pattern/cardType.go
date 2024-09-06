package pattern

import "server-v1/card"

type CardType interface {
	Judge(cards []*card.Card) (bool, []*card.Card)
	// When two cards type are same, compare them by this method
	// 0: cards1 = cards2, 1: cards1 > cards2, -1: cards1 < cards2
	Compare(cards1 []*card.Card, cards2 []*card.Card) int
	GetWeight() int
}
