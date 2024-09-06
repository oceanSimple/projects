package tool

import (
	"server-v1/card"
	"sort"
)

func SortCardsByNumber(cards []*card.Card) {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Number.Weight > cards[j].Number.Weight
	})
}

func SortCardsByNumberExtraWeight(cards []*card.Card) {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Number.ExtraWeight > cards[j].Number.ExtraWeight
	})
}

func SortCardsByColor(cards []*card.Card) {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Color.Text > cards[j].Color.Text
	})
}

func EqualCardSlice(c1, c2 []*card.Card) bool {
	if len(c1) != len(c2) {
		return false
	}
	for i := 0; i < len(c1); i++ {
		if !(c1[i].Number.Weight == c2[i].Number.Weight &&
			c1[i].Color.Text == c2[i].Color.Text) {
			return false
		}
	}
	return true
}
